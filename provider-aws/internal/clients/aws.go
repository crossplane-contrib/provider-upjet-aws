/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	xpabeta1 "github.com/crossplane/provider-aws/apis/v1beta1"
	xpawsclient "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/upbound/official-providers/provider-aws/apis/v1beta1"
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

		if mg.GetProviderConfigReference() == nil {
			return ps, errors.New("no providerConfigRef provided")
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
			return ps, errors.Wrap(err, "cannot get referenced Provider")
		}

		region, err := getRegion(mg)
		if err != nil {
			return ps, errors.Wrap(err, "cannot get region")
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, "cannot track ProviderConfig usage")
		}

		var cfg *aws.Config
		var roleARN *string
		if pc.Spec.AssumeRole != nil {
			roleARN = pc.Spec.AssumeRole.RoleARN
		}
		xpapc := &xpabeta1.ProviderConfig{
			Spec: xpabeta1.ProviderConfigSpec{
				Credentials:   xpabeta1.ProviderCredentials(pc.Spec.Credentials),
				AssumeRoleARN: roleARN,
			},
		}
		switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
		case xpv1.CredentialsSourceInjectedIdentity:
			if roleARN != nil {
				if cfg, err = xpawsclient.UsePodServiceAccountAssumeRole(ctx, []byte{}, xpawsclient.DefaultSection, region, xpapc); err != nil {
					return ps, errors.Wrap(err, "failed to use pod service account assumeRoleARN")
				}
			} else {
				if cfg, err = xpawsclient.UsePodServiceAccount(ctx, []byte{}, xpawsclient.DefaultSection, region); err != nil {
					return ps, errors.Wrap(err, "failed to use pod service account")
				}
			}
		default:
			data, err := resource.CommonCredentialExtractor(ctx, s, client, pc.Spec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, "cannot get credentials")
			}
			if roleARN != nil {
				if cfg, err = xpawsclient.UseProviderSecretAssumeRole(ctx, data, xpawsclient.DefaultSection, region, xpapc); err != nil {
					return ps, errors.Wrap(err, "failed to use provider secret with assumeRoleARN")
				}
			} else {
				if cfg, err = xpawsclient.UseProviderSecret(ctx, data, xpawsclient.DefaultSection, region); err != nil {
					return ps, errors.Wrap(err, "failed to use provider secret")
				}
			}
		}
		awsConf := xpawsclient.SetResolver(xpapc, cfg)
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
		}

		// provider configuration for credentials
		tfCfg[keyAccessKeyID] = creds.AccessKeyID
		tfCfg[keySecretAccessKey] = creds.SecretAccessKey
		tfCfg[keySessionToken] = creds.SessionToken
		ps.Configuration = tfCfg
		return ps, err
	}
}

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
