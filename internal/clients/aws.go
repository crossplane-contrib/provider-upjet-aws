// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"maps"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	awsrequest "github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/smithy-go/middleware"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/metrics"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/hashicorp/aws-sdk-go-base/v2/endpoints"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/provider-aws/apis/v1beta1"
)

const (
	keyAccountID        = "account_id"
	keyRegion           = "region"
	keyPartition        = "partition"
	localstackAccountID = "000000000"
)

type SetupConfig struct {
	TerraformProvider *schema.Provider
	Logger            logging.Logger
}

// iamRegions holds the region used for signing IAM credentials for each AWS partition.
var iamRegions = getIAMDefaultSigningRegions()

func SelectTerraformSetup(config *SetupConfig) terraform.SetupFn { // nolint:gocyclo
	credsCache := NewAWSCredentialsProviderCache(WithCacheLogger(config.Logger))
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		pc := &v1beta1.ProviderConfig{}
		if err := c.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
			return terraform.Setup{}, errors.Wrapf(err, "cannot get referenced ProviderConfig: %q", mg.GetProviderConfigReference().Name)
		}
		t := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return terraform.Setup{}, errors.Wrapf(err, "cannot track ProviderConfig usage for %q", mg.GetProviderConfigReference().Name)
		}

		ps := terraform.Setup{}
		awsCfg, err := getAWSConfigWithDefaultRegion(ctx, c, mg, pc)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot get aws config")
		} else if awsCfg == nil {
			return terraform.Setup{}, errors.Wrap(err, "obtained aws config cannot be nil")
		}

		// only IRSA auth credentials are currently cached, other auth methods
		// will skip the cache and call the downstream
		// CredentialsProvider.Retrieve().
		credCache, err := credsCache.RetrieveCredentials(ctx, pc, awsCfg.Region, awsCfg.Credentials, func(ctx context.Context) (string, error) {
			if pc.Spec.SkipCredsValidation {
				// then we do not try to resolve the account ID and instead,
				// return a constant value as before.
				return localstackAccountID, nil
			}
			o, err := sts.NewFromConfig(*awsCfg).GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
			if err != nil {
				return "", errors.Wrap(err, errGetCallerIdentityFailed)
			}
			return *o.Account, nil
		})
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cache manager failure")
		}

		// if we are to retrieve the AWS account ID and if we have not already
		// retrieved it via the credential cache, then we will utilize the
		// identity cache.
		// TODO: Replace the identity cache with the credential cache.
		if !pc.Spec.SkipCredsValidation && credCache.accountID == "" {
			credCache.accountID, err = getAccountId(ctx, awsCfg, credCache.creds)
			if err != nil {
				return terraform.Setup{}, errors.Wrap(err, "cannot get account id")
			}
		}
		// just in case the localstack implementation relies on this...
		if credCache.accountID == "" {
			credCache.accountID = localstackAccountID
		}
		ps.ClientMetadata = map[string]string{
			keyAccountID: credCache.accountID,
			keyPartition: "aws",
		}

		if err := setPartition(awsCfg, pc, &ps); err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot configure AWS partition")
		}

		// several external name configs depend on the setup.Configuration for templating region
		ps.Configuration = map[string]any{
			keyRegion: awsCfg.Region,
		}
		if config.TerraformProvider == nil {
			return terraform.Setup{}, errors.New("terraform provider cannot be nil")
		}
		return ps, errors.Wrap(configureNoForkAWSClient(ctx, &ps, config, awsCfg.Region, credCache.creds, pc), "could not configure the no-fork AWS client")
	}
}

func setPartition(awsCfg *aws.Config, pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	var partitionFromConfig string
	if pc.Spec.Endpoint != nil && pc.Spec.Endpoint.PartitionID != nil {
		partitionFromConfig = *pc.Spec.Endpoint.PartitionID
		ps.ClientMetadata[keyPartition] = partitionFromConfig
	}
	// region should never be empty, but defensively code to preserve existing behavior
	if awsCfg.Region == "" {
		return nil
	}

	// TODO(erhan): localstack environments with ALLOW_NONSTANDARD_REGIONS configuration
	// might fail this check. Consider introducing a config that opt-out from partition
	// resolution
	partitionFromRegion, ok := endpoints.PartitionForRegion(endpoints.DefaultPartitions(), awsCfg.Region)
	if !ok || partitionFromRegion.ID() == "" {
		// tolerate unknown region and honor when explicit partition config exists
		if partitionFromConfig != "" {
			return nil
		}
		return errors.Errorf("managed resource region %q does not belong to a known partition", awsCfg.Region)
	}

	if partitionFromConfig != "" && partitionFromConfig != partitionFromRegion.ID() {
		return errors.Errorf("conflicting partition config: managed resource region %q does not belong to configured partition %q at provider config", awsCfg.Region, *pc.Spec.Endpoint.PartitionID)
	}

	ps.ClientMetadata[keyPartition] = partitionFromRegion.ID()
	return nil
}

// getAccountId retrieves the account ID associated with the given credentials.
// Results are cached.
func getAccountId(ctx context.Context, cfg *aws.Config, creds aws.Credentials) (string, error) {
	identity, err := GlobalCallerIdentityCache.GetCallerIdentity(ctx, *cfg, creds)
	if err != nil {
		return "", errors.Wrap(err, "cannot get the caller identity")
	}
	return *identity.Account, nil
}

// getAWSConfigWithDefaultRegion is a utility function that wraps the
// GetAWSConfigWithoutTracking and fills empty region in the returned for
// "iam.aws.upbound.io" group with a default "us-east-1" region. Although
// this does not have an effect on the resource, as IAM group resources
// has no concept of region, this is done to conform with the TF AWS config
// which requires non-empty region
func getAWSConfigWithDefaultRegion(ctx context.Context, c client.Client, obj runtime.Object, pc *v1beta1.ProviderConfig) (*aws.Config, error) {
	cfg, err := GetAWSConfigWithoutTracking(ctx, c, obj, pc)
	if err != nil {
		return nil, err
	}
	if cfg.Region == "" && obj.GetObjectKind().GroupVersionKind().Group == "iam.aws.upbound.io" {
		cfg.Region = getIAMRegion(pc)
	}
	return cfg, nil
}

func getIAMRegion(pc *v1beta1.ProviderConfig) string {
	defaultRegion := "us-east-1"
	if pc == nil || pc.Spec.Endpoint == nil || pc.Spec.Endpoint.PartitionID == nil {
		return defaultRegion
	}
	if region, ok := iamRegions[*pc.Spec.Endpoint.PartitionID]; ok {
		return region
	}
	return defaultRegion
}

type metaOnlyPrimary struct {
	meta any
}

func (m *metaOnlyPrimary) Meta() any {
	return m.meta
}

// withExternalAPICallCounter configures an AWS SDK v2 stack (client)
// with an API call counter. AWS SDK v2 offers configuring
// "middlewares" to customize a request. Middlewares can be plugged
// into different steps of the stack. Middlewares can save and access
// metadata in the stack, such as ServiceID (EC2, IAM, etc.) and
// OperationName (DescribeVPCs, etc.). For documentation, see:
// https://aws.github.io/aws-sdk-go-v2/docs/middleware/
func withExternalAPICallCounter(stack *middleware.Stack) error {
	externalAPICallCounterMiddleware := middleware.DeserializeMiddlewareFunc("externalAPICallCounter",
		func(ctx context.Context, input middleware.DeserializeInput, next middleware.DeserializeHandler) (middleware.DeserializeOutput, middleware.Metadata, error) {
			serviceID := awsmiddleware.GetServiceID(ctx)
			operationName := awsmiddleware.GetOperationName(ctx)

			// next.HandleDeserialize() calls the next middleware function
			// in the stack, which in turn calls the next. Finally, the
			// request is performed. Each middleware function receives the
			// output from the middleware function it invoked, processes it,
			// and returns its result to the middleware function that
			// invoked itself.
			output, metadata, err := next.HandleDeserialize(ctx, input)
			if err == nil {
				metrics.ExternalAPICalls.WithLabelValues(serviceID, operationName).Inc()
			}
			return output, metadata, err
		},
	)

	// We register the call counter to the end of the deserialization
	// step, so that we're right next to Transport handler
	// (http.RoundTripper) in the stack (see
	// https://aws.github.io/aws-sdk-go-v2/docs/middleware/). In this
	// case, it's easy to distinguish API errors from connection
	// errors, because only connection errors cause a non-nil error
	// returned by next.HandleDeserialize() (see middleware
	// implementation above). If we were to register the call counter
	// to any other position (such as earlier stack steps (finalize,
	// build, etc.) or even the beginning of deserialization step), we
	// would have to implement a logic to distinguish between API
	// errors and connection errors.
	return stack.Deserialize.Add(externalAPICallCounterMiddleware, middleware.After)
}

// configureNoForkAWSClient populates the supplied *terraform.Setup with
// Terraform Plugin SDK style AWS client (Meta) and Terraform Plugin Framework
// style FrameworkProvider
func configureNoForkAWSClient(ctx context.Context, ps *terraform.Setup, config *SetupConfig, region string, creds aws.Credentials, pc *v1beta1.ProviderConfig) error { //nolint:gocyclo
	tfAwsConnsCfg := xpprovider.AWSConfig{
		AccessKey:               creds.AccessKeyID,
		Endpoints:               map[string]string{},
		Region:                  region,
		S3UsePathStyle:          pc.Spec.S3UsePathStyle,
		SecretKey:               creds.SecretAccessKey,
		SkipCredsValidation:     true, // disabled to prevent extra AWS STS call
		SkipRegionValidation:    pc.Spec.SkipRegionValidation,
		SkipRequestingAccountId: true, // disabled to prevent extra AWS STS call
		Token:                   creds.SessionToken,
	}

	if pc.Spec.SkipMetadataApiCheck {
		tfAwsConnsCfg.EC2MetadataServiceEnableState = imds.ClientDisabled
	}

	if pc.Spec.Endpoint != nil {
		if pc.Spec.Endpoint.URL.Static != nil {
			if len(pc.Spec.Endpoint.Services) > 0 && *pc.Spec.Endpoint.URL.Static == "" {
				return errors.New("endpoint.url.static cannot be empty")
			} else {
				for _, service := range pc.Spec.Endpoint.Services {
					tfAwsConnsCfg.Endpoints[service] = aws.ToString(pc.Spec.Endpoint.URL.Static)
				}
			}
		}
	}

	// only used for retrieving the ServicePackages from the singleton provider instance
	// so that we can reuse it in a newly instantiated aws client
	p := config.TerraformProvider.Meta()
	singletonServicePackages := maps.Collect(p.(*xpprovider.AWSClient).ServicePackages(ctx))
	newClient := xpprovider.AWSClient{}
	newClient.SetServicePackages(ctx, singletonServicePackages)
	tfAwsConnsClient, diags := tfAwsConnsCfg.GetClient(ctx, &newClient)
	if diags.HasError() {
		return errors.Errorf("cannot construct TF AWS Client from TF AWS Config, %v", diags)
	}
	// accountID is already calculated/retrieved from Caller ID cache while
	// obtaining AWS config. The terraform config is explicitly constructed
	// to skip requesting account ID to prevent the extra STS call. Therefore,
	// the resulting TF AWS Client has empty account ID.
	// Fill with previously calculated account ID.
	// No need for nil check on ps.ClientMetadata per golang spec.
	tfAwsConnsClient.SetAccountID(ps.ClientMetadata[keyAccountID])
	ps.Meta = tfAwsConnsClient
	fwProvider := xpprovider.GetFrameworkProviderWithMeta(&metaOnlyPrimary{meta: tfAwsConnsClient})
	ps.FrameworkProvider = fwProvider

	// Register AWS SDK v1 call counter. Unlike AWS SDK v2, v1 doesn't
	// store service ID (EC2, IAM, etc.) and operation name
	// (DescribeVPCs, etc.) in request context. Therefore, it's not
	// possible to implement this session handler's functionality with
	// an http.RoundTripper. To learn how SDK v1 session handler phases
	// map to SDK v2 middleware stack steps, see:
	// https://aws.github.io/aws-sdk-go-v2/docs/migrating/#handler-phases
	tfAwsConnsClient.Session().Handlers.Send.PushBack(func(r *awsrequest.Request) {
		// In case of API errors (or no errors), r.Error is nil.
		// In case of connection errors, r.Error is non-nil.
		if r.Error == nil {
			metrics.ExternalAPICalls.WithLabelValues(r.ClientInfo.ServiceID, r.Operation.Name).Inc()
		}
	})

	// Register AWS SDK v2 call counter
	tfAwsConnsClient.AppendAPIOptions(withExternalAPICallCounter)

	return nil
}
