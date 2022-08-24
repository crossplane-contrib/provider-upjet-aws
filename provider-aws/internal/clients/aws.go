/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sts"

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
		awsConf, err := GetAWSConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot get AWS config")
		}
		creds, err := awsConf.Credentials.Retrieve(ctx)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
		}
		// TODO(muvaf): Maybe some sort of cache so that we don't make this call
		// every time?
		identity, err := sts.NewFromConfig(*awsConf).GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
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
				keyRegion:          awsConf.Region,
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
		if ps.Configuration["region"] == "" {
			// Some resources, like iam group, do not have a notion of region
			// hence we have no region in their schema. However, terraform still
			// attempts validating region in provider config and does not like
			// both empty string or not setting it at all. We need to skip
			// region validation in this case.
			ps.Configuration["skip_region_validation"] = true

			if mg.GetObjectKind().GroupVersionKind().Group == "iam.aws.upbound.io" {
				ps.Configuration["region"] = "us-east-1"
			}
		}
		return ps, err
	}
}
