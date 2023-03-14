/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/terraform"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/provider-aws/apis/v1beta1"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	// Terraform provider configuration keys for AWS credentials.
	keyRegion               = "region"
	keySkipCredsValidation  = "skip_credentials_validation"
	keyS3UsePathStyle       = "s3_use_path_style"
	keySkipMetadataApiCheck = "skip_metadata_api_check"
	keySkipReqAccountId     = "skip_requesting_account_id"
	keyAccountId            = "account_id"
	keySessionToken         = "token"
	keyAccessKeyID          = "access_key"
	keySecretAccessKey      = "secret_key"
	keyEndpoints            = "endpoints"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot get referenced Provider")
		}
		cfg, err := GetAWSConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot get AWS config")
		}
		if cfg.Region == "" && mg.GetObjectKind().GroupVersionKind().Group == "iam.aws.upbound.io" {
			cfg.Region = "us-east-1"
		}
		creds, err := cfg.Credentials.Retrieve(ctx)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
		}
		accountId := "000000000"
		if !pc.Spec.SkipCredsValidation {
			identity, err := GlobalCallerIdentityCache.GetCallerIdentity(ctx, *cfg, creds)
			if err != nil {
				return terraform.Setup{}, errors.Wrap(err, "cannot get the caller identity")
			}
			accountId = *identity.Account
		}

		if pc.Spec.Endpoint.URL.Static != nil {
			if len(pc.Spec.Endpoint.Services) > 0 && *pc.Spec.Endpoint.URL.Static == "" {
				return terraform.Setup{}, errors.Wrap(err, "endpoint is wrong")
			}
		}
		endpoints := make(map[string]string)
		for _, service := range pc.Spec.Endpoint.Services {
			endpoints[service] = aws.ToString(pc.Spec.Endpoint.URL.Static)
		}
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
			Configuration: map[string]any{
				keyRegion:               cfg.Region,
				keySkipCredsValidation:  pc.Spec.SkipCredsValidation,
				keyS3UsePathStyle:       pc.Spec.S3UsePathStyle,
				keySkipMetadataApiCheck: pc.Spec.SkipMetadataApiCheck,
				keySkipReqAccountId:     pc.Spec.SkipReqAccountId,
				keyAccessKeyID:          creds.AccessKeyID,
				keySecretAccessKey:      creds.SecretAccessKey,
				keySessionToken:         creds.SessionToken,
				keyEndpoints:            endpoints,
			},

			// Account ID is not part of provider configuration schema, so it
			// needs to be given separately.
			ClientMetadata: map[string]string{
				keyAccountId: accountId,
			},
		}
		return ps, err
	}
}
