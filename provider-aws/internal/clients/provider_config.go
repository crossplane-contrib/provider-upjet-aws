/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	stscredstypesv2 "github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/upbound/official-providers/provider-aws/apis/v1beta1"
	"github.com/upbound/official-providers/provider-aws/internal/version"
)

const (
	// DefaultSection for INI files.
	DefaultSection = ini.DefaultSection

	envWebIdentityTokenFile = "AWS_WEB_IDENTITY_TOKEN_FILE"
)

// GlobalRegion is the region name used for AWS services that do not have a notion
// of region.
const GlobalRegion = "aws-global"

// Endpoint URL configuration types.
const (
	URLConfigTypeStatic  = "Static"
	URLConfigTypeDynamic = "Dynamic"
)

// userAgentV2 constructs the Crossplane user agent for AWS v2 clients
var userAgentV2 = config.WithAPIOptions([]func(*middleware.Stack) error{
	awsmiddleware.AddUserAgentKeyValue("upbound-provider-aws", version.Version),
})

func getRegion(obj runtime.Object) (string, error) {
	fromMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", errors.Wrap(err, "cannot convert to unstructured")
	}
	r, err := fieldpath.Pave(fromMap).GetString("spec.forProvider.region")
	if fieldpath.IsNotFound(err) {
		// Region is not required for all resources, e.g. resource in "iam"
		// group.
		return "", nil
	}
	return r, err
}

// GetAWSConfig to produce a config that can be used to authenticate to AWS.
func GetAWSConfig(ctx context.Context, c client.Client, mg resource.Managed) (*aws.Config, error) { // nolint:gocyclo
	if mg.GetProviderConfigReference() == nil {
		return nil, errors.New("no providerConfigRef provided")
	}
	region, err := getRegion(mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get region")
	}
	pc := &v1beta1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, "cannot get referenced Provider")
	}

	t := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, "cannot track ProviderConfig usage")
	}

	switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
	case xpv1.CredentialsSourceInjectedIdentity:
		if pc.Spec.AssumeRole != nil {
			cfg, err := UsePodServiceAccountAssumeRole(ctx, []byte{}, DefaultSection, region, pc)
			if err != nil {
				return nil, err
			}
			return SetResolver(pc, cfg), nil
		}
		if pc.Spec.AssumeRoleWithWebIdentity != nil && pc.Spec.AssumeRoleWithWebIdentity.RoleARN != nil {
			cfg, err := UsePodServiceAccountAssumeRoleWithWebIdentity(ctx, []byte{}, DefaultSection, region, pc)
			if err != nil {
				return nil, err
			}
			return SetResolver(pc, cfg), nil
		}
		cfg, err := UsePodServiceAccount(ctx, []byte{}, DefaultSection, region)
		if err != nil {
			return nil, err
		}
		return SetResolver(pc, cfg), nil
	default:
		data, err := resource.CommonCredentialExtractor(ctx, s, c, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get credentials")
		}
		if pc.Spec.AssumeRole != nil {
			cfg, err := UseProviderSecretAssumeRole(ctx, data, DefaultSection, region, pc)
			if err != nil {
				return nil, err
			}
			return SetResolver(pc, cfg), nil
		}
		cfg, err := UseProviderSecret(ctx, data, DefaultSection, region)
		if err != nil {
			return nil, err
		}
		return SetResolver(pc, cfg), nil
	}
}

type awsEndpointResolverAdaptorWithOptions func(service, region string, options interface{}) (aws.Endpoint, error)

func (a awsEndpointResolverAdaptorWithOptions) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return a(service, region, options)
}

// SetResolver parses annotations from the managed resource
// and returns a configuration accordingly.
func SetResolver(pc *v1beta1.ProviderConfig, cfg *aws.Config) *aws.Config { // nolint:gocyclo
	if pc.Spec.Endpoint == nil {
		return cfg
	}
	cfg.EndpointResolverWithOptions = awsEndpointResolverAdaptorWithOptions(func(service, region string, options interface{}) (aws.Endpoint, error) {
		fullURL := ""
		switch pc.Spec.Endpoint.URL.Type {
		case URLConfigTypeStatic:
			if pc.Spec.Endpoint.URL.Static == nil {
				return aws.Endpoint{}, errors.New("static type is chosen but static field does not have a value")
			}
			fullURL = aws.ToString(pc.Spec.Endpoint.URL.Static)
		case URLConfigTypeDynamic:
			if pc.Spec.Endpoint.URL.Dynamic == nil {
				return aws.Endpoint{}, errors.New("dynamic type is chosen but dynamic configuration is not given")
			}
			// NOTE(muvaf): IAM does not have any region.
			if service == "IAM" {
				fullURL = fmt.Sprintf("%s://%s.%s", pc.Spec.Endpoint.URL.Dynamic.Protocol, strings.ToLower(service), pc.Spec.Endpoint.URL.Dynamic.Host)
			} else {
				fullURL = fmt.Sprintf("%s://%s.%s.%s", pc.Spec.Endpoint.URL.Dynamic.Protocol, strings.ToLower(service), region, pc.Spec.Endpoint.URL.Dynamic.Host)
			}
		default:
			return aws.Endpoint{}, errors.New("unsupported url config type is chosen")
		}
		e := aws.Endpoint{
			URL:               fullURL,
			HostnameImmutable: aws.ToBool(pc.Spec.Endpoint.HostnameImmutable),
			PartitionID:       aws.ToString(pc.Spec.Endpoint.PartitionID),
			SigningName:       aws.ToString(pc.Spec.Endpoint.SigningName),
			SigningRegion:     aws.ToString(LateInitializeStringPtr(pc.Spec.Endpoint.SigningRegion, &region)),
			SigningMethod:     aws.ToString(pc.Spec.Endpoint.SigningMethod),
		}
		// Only IAM does not have a region parameter and "aws-global" is used in
		// SDK setup. However, signing region has to be us-east-1 and it needs
		// to be set.
		if region == "aws-global" {
			switch aws.ToString(pc.Spec.Endpoint.PartitionID) {
			case "aws-us-gov", "aws-cn", "aws-iso", "aws-iso-b":
				e.SigningRegion = aws.ToString(LateInitializeStringPtr(pc.Spec.Endpoint.SigningRegion, &region))
			default:
				e.SigningRegion = "us-east-1"
			}
		}
		if pc.Spec.Endpoint.Source != nil {
			switch *pc.Spec.Endpoint.Source {
			case "ServiceMetadata":
				e.Source = aws.EndpointSourceServiceMetadata
			case "Custom":
				e.Source = aws.EndpointSourceCustom
			}
		}
		return e, nil
	})
	return cfg
}

// CredentialsIDSecret retrieves AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY from the data which contains
// aws credentials under given profile
// Example:
// [default]
// aws_access_key_id = <YOUR_ACCESS_KEY_ID>
// aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
func CredentialsIDSecret(data []byte, profile string) (aws.Credentials, error) {
	awsConfig, err := ini.InsensitiveLoad(data)
	if err != nil {
		return aws.Credentials{}, errors.Wrap(err, "cannot parse credentials secret")
	}

	iniProfile, err := awsConfig.GetSection(profile)
	if err != nil {
		return aws.Credentials{}, errors.Wrap(err, fmt.Sprintf("cannot get %s profile in credentials secret", profile))
	}

	accessKeyID := iniProfile.Key("aws_access_key_id")
	secretAccessKey := iniProfile.Key("aws_secret_access_key")
	sessionToken := iniProfile.Key("aws_session_token")

	// NOTE(muvaf): Key function implementation never returns nil but still its
	// type is pointer so we check to make sure its next versions doesn't break
	// that implicit contract.
	if accessKeyID == nil || secretAccessKey == nil || sessionToken == nil {
		return aws.Credentials{}, errors.New("returned key can be empty but cannot be nil")
	}

	return aws.Credentials{
		AccessKeyID:     accessKeyID.Value(),
		SecretAccessKey: secretAccessKey.Value(),
		SessionToken:    sessionToken.Value(),
	}, nil
}

// AuthMethod is a method of authenticating to the AWS API
type AuthMethod func(context.Context, []byte, string, string) (*aws.Config, error)

// UseProviderSecret - AWS configuration which can be used to issue requests against AWS API
func UseProviderSecret(ctx context.Context, data []byte, profile, region string) (*aws.Config, error) {
	creds, err := CredentialsIDSecret(data, profile)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse credentials secret")
	}

	awsConfig, err := config.LoadDefaultConfig(
		ctx,
		userAgentV2,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: creds,
		}),
	)
	return &awsConfig, err
}

// UseProviderSecretAssumeRole - AWS configuration which can be used to issue requests against AWS API
// assume Cross account IAM roles
func UseProviderSecretAssumeRole(ctx context.Context, data []byte, profile, region string, pc *v1beta1.ProviderConfig) (*aws.Config, error) {
	creds, err := CredentialsIDSecret(data, profile)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse credentials secret")
	}

	awsConfig, err := config.LoadDefaultConfig(
		ctx,
		userAgentV2,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: creds,
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load default AWS config")
	}

	roleArn, err := GetAssumeRoleARN(pc.Spec.DeepCopy())
	if err != nil {
		return nil, err
	}

	stsSvc := sts.NewFromConfig(awsConfig)

	stsAssumeRoleOptions := SetAssumeRoleOptions(pc)
	stsAssume := stscreds.NewAssumeRoleProvider(
		stsSvc,
		aws.ToString(roleArn),
		stsAssumeRoleOptions,
	)
	awsConfig.Credentials = aws.NewCredentialsCache(stsAssume)

	return &awsConfig, err
}

// UsePodServiceAccountAssumeRole assumes an IAM role configured via a ServiceAccount
// assume Cross account IAM roles
// https://aws.amazon.com/blogs/containers/cross-account-iam-roles-for-kubernetes-service-accounts/
func UsePodServiceAccountAssumeRole(ctx context.Context, _ []byte, _, region string, pc *v1beta1.ProviderConfig) (*aws.Config, error) {
	cfg, err := UsePodServiceAccount(ctx, []byte{}, DefaultSection, region)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load default AWS config")
	}
	roleArn, err := GetAssumeRoleARN(pc.Spec.DeepCopy())
	if err != nil {
		return nil, err
	}
	stsclient := sts.NewFromConfig(*cfg)
	stsAssumeRoleOptions := SetAssumeRoleOptions(pc)
	cnf, err := config.LoadDefaultConfig(
		ctx,
		userAgentV2,
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(aws.NewCredentialsCache(
			stscreds.NewAssumeRoleProvider(
				stsclient,
				aws.ToString(roleArn),
				stsAssumeRoleOptions,
			)),
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load assumed role AWS config")
	}
	return &cnf, err
}

// UsePodServiceAccountAssumeRoleWithWebIdentity assumes an IAM role
// configured via a ServiceAccount assume Cross account IAM roles
// https://aws.amazon.com/blogs/containers/cross-account-iam-roles-for-kubernetes-service-accounts/
func UsePodServiceAccountAssumeRoleWithWebIdentity(ctx context.Context, _ []byte, _, region string, pc *v1beta1.ProviderConfig) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx, userAgentV2)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load default AWS config")
	}

	roleArn, err := GetAssumeRoleWithWebIdentityARN(pc.Spec.DeepCopy())
	if err != nil {
		return nil, err
	}

	stsclient := sts.NewFromConfig(cfg)
	webIdentityRoleOptions := SetWebIdentityRoleOptions(pc)

	cnf, err := config.LoadDefaultConfig(
		ctx,
		userAgentV2,
		config.WithRegion(region),
		config.WithCredentialsProvider(aws.NewCredentialsCache(
			stscreds.NewWebIdentityRoleProvider(
				stsclient,
				aws.ToString(roleArn),
				stscreds.IdentityTokenFile(os.Getenv(envWebIdentityTokenFile)),
				webIdentityRoleOptions,
			)),
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load assumed role AWS config")
	}
	return &cnf, err
}

// UsePodServiceAccount assumes an IAM role configured via a ServiceAccount.
// https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
func UsePodServiceAccount(ctx context.Context, _ []byte, _, region string) (*aws.Config, error) {
	if region == GlobalRegion {
		cfg, err := config.LoadDefaultConfig(
			ctx,
			userAgentV2,
		)
		return &cfg, errors.Wrap(err, "failed to load default AWS config")
	}
	cfg, err := config.LoadDefaultConfig(
		ctx,
		userAgentV2,
		config.WithRegion(region),
	)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to load default AWS config with region %s", region))
	}
	return &cfg, err
}

// GetAssumeRoleARN gets the AssumeRoleArn from a ProviderConfigSpec
func GetAssumeRoleARN(pcs *v1beta1.ProviderConfigSpec) (*string, error) {
	if pcs.AssumeRole != nil && aws.ToString(pcs.AssumeRole.RoleARN) != "" {
		return pcs.AssumeRole.RoleARN, nil
	}

	return nil, errors.New("a RoleARN must be set to assume an IAM Role")
}

// GetAssumeRoleWithWebIdentityARN gets the RoleArn from a ProviderConfigSpec
func GetAssumeRoleWithWebIdentityARN(pcs *v1beta1.ProviderConfigSpec) (*string, error) {
	if pcs.AssumeRoleWithWebIdentity != nil {
		if pcs.AssumeRoleWithWebIdentity.RoleARN != nil && aws.ToString(pcs.AssumeRoleWithWebIdentity.RoleARN) != "" {
			return pcs.AssumeRoleWithWebIdentity.RoleARN, nil
		}
	}

	return nil, errors.New("a RoleARN must be set to assume with web identity")
}

// SetAssumeRoleOptions sets options when Assuming an IAM Role
func SetAssumeRoleOptions(pc *v1beta1.ProviderConfig) func(*stscreds.AssumeRoleOptions) {
	if pc.Spec.AssumeRole != nil {
		return func(opt *stscreds.AssumeRoleOptions) {
			if pc.Spec.AssumeRole.ExternalID != nil {
				opt.ExternalID = pc.Spec.AssumeRole.ExternalID
			}

			if pc.Spec.AssumeRole.Tags != nil && len(pc.Spec.AssumeRole.Tags) > 0 {
				for _, t := range pc.Spec.AssumeRole.Tags {
					opt.Tags = append(
						opt.Tags,
						stscredstypesv2.Tag{Key: t.Key, Value: t.Value})
				}
			}

			if pc.Spec.AssumeRole.TransitiveTagKeys != nil && len(pc.Spec.AssumeRole.TransitiveTagKeys) > 0 {
				opt.TransitiveTagKeys = pc.Spec.AssumeRole.TransitiveTagKeys
			}
		}
	}
	return func(opt *stscreds.AssumeRoleOptions) {}
}

// SetWebIdentityRoleOptions sets options when exchanging a WebIdentity Token for a Role
func SetWebIdentityRoleOptions(pc *v1beta1.ProviderConfig) func(*stscreds.WebIdentityRoleOptions) {
	if pc.Spec.AssumeRoleWithWebIdentity != nil {
		return func(opt *stscreds.WebIdentityRoleOptions) {
			if pc.Spec.AssumeRoleWithWebIdentity.RoleSessionName != "" {
				opt.RoleSessionName = pc.Spec.AssumeRoleWithWebIdentity.RoleSessionName
			}
		}
	}

	return func(opt *stscreds.WebIdentityRoleOptions) {}
}

// LateInitializeStringPtr returns in if it's non-nil, otherwise returns from
// which is the backup for the cases in is nil.
func LateInitializeStringPtr(in *string, from *string) *string {
	if in != nil {
		return in
	}
	return from
}
