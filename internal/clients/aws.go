// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"reflect"
	"unsafe"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/provider-aws/apis/v1beta1"
)

const (
	keyAccountID = "account_id"
	keyPartition = "partition"
	keyRegion    = "region"
)

type SetupConfig struct {
	TerraformProvider *schema.Provider
}

func SelectTerraformSetup(config *SetupConfig) terraform.SetupFn { // nolint:gocyclo
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
		creds, err := awsCfg.Credentials.Retrieve(ctx)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
		}

		account := "000000000"
		if !pc.Spec.SkipCredsValidation {
			account, err = getAccountId(ctx, awsCfg, creds)
			if err != nil {
				return terraform.Setup{}, errors.Wrap(err, "cannot get account id")
			}
		}
		ps.ClientMetadata = map[string]string{
			keyAccountID: account,
			keyPartition: "aws",
		}

		if pc.Spec.Endpoint != nil && pc.Spec.Endpoint.PartitionID != nil {
			ps.ClientMetadata[keyPartition] = *pc.Spec.Endpoint.PartitionID
		}

		// several external name configs depend on the setup.Configuration for templating region
		ps.Configuration = map[string]any{
			keyRegion: awsCfg.Region,
		}
		if config.TerraformProvider == nil {
			return terraform.Setup{}, errors.New("terraform provider cannot be nil")
		}
		return ps, errors.Wrap(configureNoForkAWSClient(ctx, &ps, config, awsCfg.Region, creds, pc), "could not configure the no-fork AWS client")
	}
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
		cfg.Region = "us-east-1"
	}
	return cfg, nil
}

type metaOnlyPrimary struct {
	meta any
}

func (m *metaOnlyPrimary) Meta() any {
	return m.meta
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
	p := config.TerraformProvider.Meta()
	tfAwsConnsClient, diags := tfAwsConnsCfg.GetClient(ctx, &xpprovider.AWSClient{
		// #nosec G103
		ServicePackages: (*xpprovider.AWSClient)(unsafe.Pointer(reflect.ValueOf(p).Pointer())).ServicePackages,
	})
	if diags.HasError() {
		return errors.Errorf("cannot construct TF AWS Client from TF AWS Config, %v", diags)
	}
	// accountID is already calculated/retrieved from Caller ID cache while
	// obtaining AWS config. The terraform config is explicitly constructed
	// to skip requesting account ID to prevent the extra STS call. Therefore,
	// the resulting TF AWS Client has empty account ID.
	// Fill with previously calculated account ID.
	// No need for nil check on ps.ClientMetadata per golang spec.
	tfAwsConnsClient.AccountID = ps.ClientMetadata[keyAccountID]
	ps.Meta = tfAwsConnsClient
	fwProvider := xpprovider.GetFrameworkProviderWithMeta(&metaOnlyPrimary{meta: tfAwsConnsClient})
	ps.FrameworkProvider = fwProvider
	return nil
}
