/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"
	"os"
	"reflect"
	"unsafe"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/provider-aws/apis/v1beta1"
)

const (
	// Terraform provider configuration keys for AWS credentials.
	keyRegion                    = "region"
	keyAccountId                 = "account_id"
	keySessionToken              = "token"
	keyAccessKeyID               = "access_key"
	keySecretAccessKey           = "secret_key"
	keyAssumeRoleWithWebIdentity = "assume_role_with_web_identity"
	keyRoleArn                   = "role_arn"
	keySessionName               = "session_name"
	keyWebIdentityTokenFile      = "web_identity_token_file"
	keySkipCredsValidation       = "skip_credentials_validation"
	keyS3UsePathStyle            = "s3_use_path_style"
	keySkipMetadataApiCheck      = "skip_metadata_api_check"
	keySkipRegionValidation      = "skip_region_validation"
	keySkipReqAccountId          = "skip_requesting_account_id"
	keyEndpoints                 = "endpoints"
)

type SetupConfig struct {
	NativeProviderPath    *string
	NativeProviderSource  *string
	NativeProviderVersion *string
	TerraformVersion      *string
	DefaultScheduler      terraform.ProviderScheduler
	TerraformProvider     *schema.Provider
	AWSClient             *xpprovider.AWSClient
}

func SelectTerraformSetup(log logging.Logger, config *SetupConfig) terraform.SetupFn { // nolint:gocyclo
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		pc := &v1beta1.ProviderConfig{}
		var err error
		if err = c.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
			return terraform.Setup{}, errors.Wrapf(err, "cannot get referenced Provider: %s", mg.GetProviderConfigReference().Name)
		}
		ps := terraform.Setup{
			Version: *config.TerraformVersion,
			Requirement: terraform.ProviderRequirement{
				Source:  *config.NativeProviderSource,
				Version: *config.NativeProviderVersion,
			},
			Scheduler: config.DefaultScheduler,
		}
		awsCfg, err := getAWSConfig(ctx, c, mg)
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
			keyAccountId: account,
		}

		if len(pc.Spec.AssumeRoleChain) > 0 || pc.Spec.Endpoint != nil {
			err = DefaultTerraformSetupBuilder(ctx, pc, &ps, awsCfg, creds)
			if err != nil {
				return terraform.Setup{}, errors.Wrap(err, "cannot build terraform configuration")
			}
			// we cannot use the shared scheduler here.
			// We will force a workspace scheduler if we can configure one.
			if len(*config.NativeProviderPath) != 0 {
				ps.Scheduler = terraform.NewWorkspaceProviderScheduler(log, terraform.WithNativeProviderPath(*config.NativeProviderPath), terraform.WithNativeProviderName("registry.terraform.io/"+*config.NativeProviderSource))
			}
		} else {
			err = pushDownTerraformSetupBuilder(ctx, c, pc, &ps, awsCfg)
			if err != nil {
				return terraform.Setup{}, errors.Wrap(err, "cannot build terraform configuration")
			}
		}

		if config.TerraformProvider == nil {
			return terraform.Setup{}, errors.New("terraform provider cannot be nil")
		}

		return ps, errors.Wrap(configureNoForkAWSClient(ctx, &ps, config), "could not configure the no-fork AWS client")
	}
}

func pushDownTerraformSetupBuilder(ctx context.Context, c client.Client, pc *v1beta1.ProviderConfig, ps *terraform.Setup, cfg *aws.Config) error { //nolint:gocyclo
	if len(pc.Spec.AssumeRoleChain) > 0 || pc.Spec.Endpoint != nil {
		return errors.New("shared scheduler cannot be used because the length of assume role chain array " +
			"is more than 0 or endpoint configuration is not nil")
	}

	ps.Configuration = map[string]any{
		keyRegion: cfg.Region,
	}

	switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
	case authKeyWebIdentity:
		if pc.Spec.Credentials.WebIdentity == nil {
			return errors.New(`spec.credentials.webIdentity of ProviderConfig cannot be nil when the credential source is "WebIdentity"`)
		}
		webIdentityConfig := map[string]any{
			keyRoleArn:              aws.ToString(pc.Spec.Credentials.WebIdentity.RoleARN),
			keyWebIdentityTokenFile: os.Getenv(envWebIdentityTokenFile),
		}
		if pc.Spec.Credentials.WebIdentity.RoleSessionName != "" {
			webIdentityConfig[keySessionName] = pc.Spec.Credentials.WebIdentity.RoleSessionName
		}
		ps.Configuration[keyAssumeRoleWithWebIdentity] = []any{
			webIdentityConfig,
		}
	case authKeyUpbound:
		if pc.Spec.Credentials.Upbound == nil || pc.Spec.Credentials.Upbound.WebIdentity == nil {
			return errors.New(`spec.credentials.upbound.webIdentity of ProviderConfig cannot be nil when the credential source is "Upbound"`)
		}
		webIdentityConfig := map[string]any{
			keyRoleArn:              aws.ToString(pc.Spec.Credentials.Upbound.WebIdentity.RoleARN),
			keyWebIdentityTokenFile: upboundProviderIdentityTokenFile,
		}
		if pc.Spec.Credentials.Upbound.WebIdentity.RoleSessionName != "" {
			webIdentityConfig[keySessionName] = pc.Spec.Credentials.Upbound.WebIdentity.RoleSessionName
		}
		ps.Configuration[keyAssumeRoleWithWebIdentity] = []any{
			webIdentityConfig,
		}
	case authKeySecret:
		data, err := resource.CommonCredentialExtractor(ctx, s, c, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return errors.Wrap(err, "cannot get credentials")
		}
		cfg, err = UseProviderSecret(ctx, data, DefaultSection, cfg.Region)
		if err != nil {
			return errors.Wrap(err, errAWSConfig)
		}
		creds, err := cfg.Credentials.Retrieve(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to retrieve aws credentials from aws config")
		}
		ps.Configuration = map[string]any{
			keyRegion:          cfg.Region,
			keyAccessKeyID:     creds.AccessKeyID,
			keySecretAccessKey: creds.SecretAccessKey,
			keySessionToken:    creds.SessionToken,
		}
	}
	return nil
}

func DefaultTerraformSetupBuilder(_ context.Context, pc *v1beta1.ProviderConfig, ps *terraform.Setup, cfg *aws.Config, creds aws.Credentials) error {
	ps.Configuration = map[string]any{
		keyRegion:               cfg.Region,
		keyAccessKeyID:          creds.AccessKeyID,
		keySecretAccessKey:      creds.SecretAccessKey,
		keySessionToken:         creds.SessionToken,
		keySkipCredsValidation:  pc.Spec.SkipCredsValidation,
		keyS3UsePathStyle:       pc.Spec.S3UsePathStyle,
		keySkipRegionValidation: pc.Spec.SkipRegionValidation,
		keySkipMetadataApiCheck: pc.Spec.SkipMetadataApiCheck,
		keySkipReqAccountId:     pc.Spec.SkipReqAccountId,
	}

	if pc.Spec.Endpoint != nil {
		if pc.Spec.Endpoint.URL.Static != nil {
			if len(pc.Spec.Endpoint.Services) > 0 && *pc.Spec.Endpoint.URL.Static == "" {
				return errors.New("endpoint.url.static cannot be empty")
			} else {
				endpoints := make(map[string]string)
				for _, service := range pc.Spec.Endpoint.Services {
					endpoints[service] = aws.ToString(pc.Spec.Endpoint.URL.Static)
				}
				ps.Configuration[keyEndpoints] = []any{endpoints}
			}
		}
	}
	return nil
}

func getAccountId(ctx context.Context, cfg *aws.Config, creds aws.Credentials) (string, error) {
	identity, err := GlobalCallerIdentityCache.GetCallerIdentity(ctx, *cfg, creds)
	if err != nil {
		return "", errors.Wrap(err, "cannot get the caller identity")
	}
	return *identity.Account, nil
}

func getAWSConfig(ctx context.Context, c client.Client, mg resource.Managed) (*aws.Config, error) {
	cfg, err := GetAWSConfig(ctx, c, mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get AWS config")
	}
	if cfg.Region == "" && mg.GetObjectKind().GroupVersionKind().Group == "iam.aws.upbound.io" {
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

func configureNoForkAWSClient(ctx context.Context, ps *terraform.Setup, config *SetupConfig) error { //nolint:gocyclo
	p := *config.TerraformProvider
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	// #nosec G103
	(*xpprovider.AWSClient)(unsafe.Pointer(reflect.ValueOf(ps.Meta).Pointer())).ServicePackages = (*xpprovider.AWSClient)(unsafe.Pointer(reflect.ValueOf(config.AWSClient).Pointer())).ServicePackages
	fwProvider := xpprovider.GetFrameworkProviderWithMeta(&metaOnlyPrimary{meta: p.Meta()})
	ps.FrameworkProvider = fwProvider
	return nil
}
