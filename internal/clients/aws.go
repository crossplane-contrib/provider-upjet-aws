/*
Copyright 2022 Upbound Inc.
*/

package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"k8s.io/apimachinery/pkg/types"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/terraform"
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
	keyWebIdentityTokenFile      = "web_identity_token_file"
	keyAssumeRole                = "assume_role"
	keyTags                      = "tags"
	keyTransitiveTagKeys         = "transitive_tag_keys"
)

func SelectTerraformSetup(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		pc := &v1beta1.ProviderConfig{}
		if err := c.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot get referenced Provider")
		}
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}
		account, err := getAccountId(ctx, c, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot get account id")
		}
		ps.ClientMetadata = map[string]string{
			keyAccountId: account,
		}

		var err2 error
		if (len(pc.Spec.AssumeRoleChain) > 1 || pc.Spec.Endpoint != nil) && (pc.Spec.Credentials.Source == authKeyIRSA || pc.Spec.Credentials.Source == authKeyWebIdentity || pc.Spec.Credentials.Source == authKeyUpbound) {
			ps, err2 = DefaultTerraformSetupBuilder(ctx, c, mg, ps)
			if err2 != nil {
				return terraform.Setup{}, errors.Wrap(err2, "cannot build terraform configuration")
			}
		} else {
			ps, err2 = OptimizedTerraformSetupBuilder(ctx, c, mg, pc, ps)
			if err2 != nil {
				return terraform.Setup{}, errors.Wrap(err2, "cannot build terraform configuration")
			}
		}
		return ps, nil
	}
}

func OptimizedTerraformSetupBuilder(ctx context.Context, c client.Client, mg resource.Managed, pc *v1beta1.ProviderConfig, ps terraform.Setup) (terraform.Setup, error) { //nolint:gocyclo
	cfg, err := getAWSConfig(ctx, c, mg)
	if err != nil {
		return terraform.Setup{}, errors.Wrap(err, "cannot get AWS config")
	}
	ps.Configuration = map[string]any{
		keyRegion: cfg.Region,
	}

	switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
	case authKeyWebIdentity:
		if pc.Spec.Credentials.WebIdentity == nil {
			return terraform.Setup{}, errors.New(`spec.credentials.webIdentity of ProviderConfig cannot be nil when the credential source is "WebIdentity"`)
		}
		ps.Configuration[keyAssumeRoleWithWebIdentity] = pc.Spec.Credentials.WebIdentity.RoleARN
		ps.Configuration[keyWebIdentityTokenFile] = defaultIdentityTokenFile
	case authKeyUpbound:
		if pc.Spec.Credentials.Upbound == nil || pc.Spec.Credentials.Upbound.WebIdentity == nil {
			return terraform.Setup{}, errors.New(`spec.credentials.upbound.webIdentity of ProviderConfig cannot be nil when the credential source is "Upbound"`)
		}
		ps.Configuration[keyAssumeRoleWithWebIdentity] = pc.Spec.Credentials.Upbound.WebIdentity.RoleARN
		ps.Configuration[keyWebIdentityTokenFile] = upboundProviderIdentityTokenFile
	case authKeySecret:
		data, err := resource.CommonCredentialExtractor(ctx, s, c, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot get credentials")
		}
		cfg, err = UseProviderSecret(ctx, data, DefaultSection, cfg.Region)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, errAWSConfig)
		}
		creds, err := cfg.Credentials.Retrieve(ctx)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
		}
		ps.Configuration = map[string]any{
			keyRegion:          cfg.Region,
			keyAccessKeyID:     creds.AccessKeyID,
			keySecretAccessKey: creds.SecretAccessKey,
			keySessionToken:    creds.SessionToken,
		}
	}

	if len(pc.Spec.AssumeRoleChain) != 0 {
		ps.Configuration[keyAssumeRole] = pc.Spec.AssumeRoleChain[0].RoleARN
		ps.Configuration[keyTags] = pc.Spec.AssumeRoleChain[0].Tags
		ps.Configuration[keyTransitiveTagKeys] = pc.Spec.AssumeRoleChain[0].TransitiveTagKeys
	}
	return ps, nil
}

func DefaultTerraformSetupBuilder(ctx context.Context, c client.Client, mg resource.Managed, ps terraform.Setup) (terraform.Setup, error) {
	cfg, err := getAWSConfig(ctx, c, mg)
	if err != nil {
		return terraform.Setup{}, errors.Wrap(err, "cannot get AWS config")
	}
	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return terraform.Setup{}, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
	}
	ps.Configuration = map[string]any{
		keyRegion:          cfg.Region,
		keyAccessKeyID:     creds.AccessKeyID,
		keySecretAccessKey: creds.SecretAccessKey,
		keySessionToken:    creds.SessionToken,
	}
	return ps, err
}

func getAccountId(ctx context.Context, c client.Client, mg resource.Managed) (string, error) {
	cfg, err := getAWSConfig(ctx, c, mg)
	if err != nil {
		return "", errors.Wrap(err, "cannot get AWS config")
	}
	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve aws credentials from aws config")
	}
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
