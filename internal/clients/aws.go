/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"
)

const (
	// Terraform provider configuration keys for AWS credentials.
	keyRegion          = "region"
	keyAccountId       = "account_id"
	keySessionToken    = "token"
	keyAccessKeyID     = "access_key"
	keySecretAccessKey = "secret_key"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
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
		identity, err := GlobalCallerIdentityCache.GetCallerIdentity(ctx, *cfg, creds)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot get the caller identity")
		}

		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
			Configuration: map[string]any{
				keyRegion:          cfg.Region,
				keyAccessKeyID:     creds.AccessKeyID,
				keySecretAccessKey: creds.SecretAccessKey,
				keySessionToken:    creds.SessionToken,
			},
			// Account ID is not part of provider configuration schema, so it
			// needs to be given separately.
			ClientMetadata: map[string]string{
				keyAccountId: *identity.Account,
			},
		}
		return ps, err
	}
}
