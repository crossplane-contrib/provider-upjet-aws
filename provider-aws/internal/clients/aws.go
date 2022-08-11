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
	// Terraform provider configuration keys for AWS credentials
	keySessionToken    = "token"
	keyAccessKeyID     = "access_key"
	keySecretAccessKey = "secret_key"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		awsConf, err := GetAWSConfig(ctx, client, mg)
		if err != nil {
			return ps, errors.Wrap(err, "cannot get AWS config")
		}
		creds, err := awsConf.Credentials.Retrieve(ctx)

		if err != nil {
			return ps, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
		}

		// TODO(hasan): figure out what other values could be possible set here.
		//   e.g. what about setting an assume_role section: https://registry.terraform.io/providers/hashicorp/aws/latest/docs#argument-reference
		tfCfg := map[string]interface{}{}
		tfCfg["region"] = awsConf.Region
		if awsConf.Region == "" {
			// Some resources, like iam group, do not have a notion of region
			// hence we have no region in their schema. However, terraform still
			// attempts validating region in provider config and does not like
			// both empty string or not setting it at all. We need to skip
			// region validation in this case.
			tfCfg["skip_region_validation"] = true

			if mg.GetObjectKind().GroupVersionKind().Group == "iam.aws.upbound.io" {
				tfCfg["region"] = "us-east-1"
			}
		}

		// provider configuration for credentials
		tfCfg[keyAccessKeyID] = creds.AccessKeyID
		tfCfg[keySecretAccessKey] = creds.SecretAccessKey
		tfCfg[keySessionToken] = creds.SessionToken
		ps.Configuration = tfCfg
		return ps, err
	}
}
