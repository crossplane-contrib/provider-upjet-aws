// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
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

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/upbound/provider-aws/apis/v1beta1"
	"github.com/upbound/provider-aws/internal/version"
)

const (
	// DefaultSection for INI files.
	DefaultSection = ini.DefaultSection

	// authentication types
	authKeyIRSA        = "IRSA"
	authKeyWebIdentity = "WebIdentity"
	authKeyPodIdentity = "PodIdentity"
	authKeyUpbound     = "Upbound"
	// authKeySAML        = "SAML"

	envWebIdentityTokenFile = "AWS_WEB_IDENTITY_TOKEN_FILE"
	envWebIdentityRoleARN   = "AWS_ROLE_ARN"
	errRoleChainConfig      = "failed to load assumed role AWS config"
	errAWSConfig            = "failed to get AWS config"
	errAWSConfigIRSA        = "failed to get AWS config using IAM Roles for Service Accounts"
	errAWSConfigWebIdentity = "failed to get AWS config using web identity token"
	errAWSConfigPodIdentity = "failed to get AWS config using pod identity"
	errAWSConfigUpbound     = "failed to get AWS config using Upbound identity"

	upboundProviderIdentityTokenFile = "/var/run/secrets/upbound.io/provider/token"
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
	awsmiddleware.AddUserAgentKeyValue("crossplane-provider-aws", version.Version),
	withExternalAPICallCounter,
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

// GetAWSConfigWithoutTracking produces an AWS config from the specified
// v1beta1.ProviderConfig that can be used to authenticate to AWS.
// ProviderConfigUsage is not tracked when this function is called.
// The caller is responsible for tracking the usage if needed.
func GetAWSConfigWithoutTracking(ctx context.Context, c client.Client, obj runtime.Object, pc *v1beta1.ProviderConfig) (*aws.Config, error) { // nolint:gocyclo
	region, err := getRegion(obj)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get region")
	}
	var cfg *aws.Config
	switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
	case authKeyIRSA:
		cfg, err = UseDefault(ctx, region)
		if err != nil {
			return nil, errors.Wrap(err, errAWSConfigIRSA)
		}
	case authKeyPodIdentity:
		cfg, err = UseDefault(ctx, region)
		if err != nil {
			return nil, errors.Wrap(err, errAWSConfigPodIdentity)
		}
	case authKeyWebIdentity:
		cfg, err = UseWebIdentityToken(ctx, region, &pc.Spec, c)
		if err != nil {
			return nil, errors.Wrap(err, errAWSConfigWebIdentity)
		}
	case authKeyUpbound:
		cfg, err = UseUpbound(ctx, region, &pc.Spec)
		if err != nil {
			return nil, errors.Wrap(err, errAWSConfigUpbound)
		}
	default:
		data, err := resource.CommonCredentialExtractor(ctx, s, c, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get credentials")
		}
		cfg, err = UseProviderSecret(ctx, data, DefaultSection, region)
		if err != nil {
			return nil, errors.Wrap(err, errAWSConfig)
		}
	}

	cfg, err = GetRoleChainConfig(ctx, c, &pc.Spec, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get credentials")
	}
	return SetResolver(pc, cfg), nil
}

// GetAWSConfigWithTracking obtains the provider config referenced by the
// specified managed resource and produces a config that can be used to
// authenticate to AWS and tracks the ProviderConfigUsage. Useful for obtaining
// AWS config for non-upjet based MR controllers.
func GetAWSConfigWithTracking(ctx context.Context, c client.Client, mg resource.Managed) (*aws.Config, error) {
	if mg.GetProviderConfigReference() == nil {
		return nil, errors.New("no providerConfigRef provided")
	}
	pc := &v1beta1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, "cannot get referenced Provider")
	}

	t := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, "cannot track ProviderConfig usage")
	}
	return GetAWSConfigWithoutTracking(ctx, c, mg, pc)
}

// TODO: Update to use the new endpoint resolution method. SA1019: aws.Endpoint is deprecated.
type awsEndpointResolverAdaptorWithOptions func(service, region string, options interface{}) (aws.Endpoint, error) // nolint: staticcheck

func (a awsEndpointResolverAdaptorWithOptions) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) { // nolint: staticcheck
	return a(service, region, options)
}

// SetResolver parses annotations from the managed resource
// and returns a configuration accordingly.
func SetResolver(pc *v1beta1.ProviderConfig, cfg *aws.Config) *aws.Config { // nolint:gocyclo
	if pc.Spec.Endpoint == nil {
		return cfg
	}
	cfg.EndpointResolverWithOptions = awsEndpointResolverAdaptorWithOptions(func(service, region string, options interface{}) (aws.Endpoint, error) { //nolint:staticcheck
		fullURL := ""
		switch pc.Spec.Endpoint.URL.Type {
		case URLConfigTypeStatic:
			if pc.Spec.Endpoint.URL.Static == nil {
				return aws.Endpoint{}, errors.New("static type is chosen but static field does not have a value") // nolint: staticcheck
			}
			fullURL = aws.ToString(pc.Spec.Endpoint.URL.Static)
		case URLConfigTypeDynamic:
			if pc.Spec.Endpoint.URL.Dynamic == nil {
				return aws.Endpoint{}, errors.New("dynamic type is chosen but dynamic configuration is not given") // nolint: staticcheck
			}
			// NOTE(muvaf): IAM does not have any region.
			if service == "IAM" {
				fullURL = fmt.Sprintf("%s://%s.%s", pc.Spec.Endpoint.URL.Dynamic.Protocol, strings.ToLower(service), pc.Spec.Endpoint.URL.Dynamic.Host)
			} else {
				fullURL = fmt.Sprintf("%s://%s.%s.%s", pc.Spec.Endpoint.URL.Dynamic.Protocol, strings.ToLower(service), region, pc.Spec.Endpoint.URL.Dynamic.Host)
			}
		default:
			return aws.Endpoint{}, errors.New("unsupported url config type is chosen") // nolint: staticcheck
		}
		e := aws.Endpoint{ // nolint: staticcheck
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

// stsRegionOrDefault sets the STS client region to the passed region, or
// defaults to the global region.
func stsRegionOrDefault(region string) func(*sts.Options) {
	return func(o *sts.Options) {
		if region == "" {
			o.Region = GlobalRegion
		}
	}
}

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
	if err != nil {
		return nil, errors.Wrap(err, "cannot load default AWS config")
	}
	return &awsConfig, nil
}

// GetRoleChainConfig returns an aws.Config capable of doing role chaining with
// AssumeRoleWithWebIdentity & AssumeRoles.
func GetRoleChainConfig(ctx context.Context, kube client.Client, pcs *v1beta1.ProviderConfigSpec, cfg *aws.Config) (*aws.Config, error) {
	pCfg := cfg
	for _, aro := range pcs.AssumeRoleChain {
		stsAssume := stscreds.NewAssumeRoleProvider(
			sts.NewFromConfig(*pCfg, stsRegionOrDefault(cfg.Region)), //nolint:contextcheck
			aws.ToString(aro.RoleARN),
			SetAssumeRoleOptions(ctx, kube, aro),
		)
		cfgWithAssumeRole, err := config.LoadDefaultConfig(
			ctx,
			userAgentV2,
			config.WithRegion(cfg.Region),
			config.WithCredentialsProvider(aws.NewCredentialsCache(stsAssume)),
		)
		if err != nil {
			return nil, errors.Wrap(err, errRoleChainConfig)
		}
		pCfg = &cfgWithAssumeRole
	}
	return pCfg, nil
}

// GetAssumeRoleWithWebIdentityConfig returns an aws.Config capable of doing
// AssumeRoleWithWebIdentity.
func GetAssumeRoleWithWebIdentityConfig(ctx context.Context, cfg *aws.Config, webID v1beta1.AssumeRoleWithWebIdentityOptions, tokenFile string) (*aws.Config, error) {
	return GetAssumeRoleWithWebIdentityConfigViaTokenRetriever(ctx, cfg, webID, stscreds.IdentityTokenFile(filepath.Clean(tokenFile)))
}

// GetAssumeRoleWithWebIdentityConfigViaTokenRetriever returns an aws.Config capable of doing
// AssumeRoleWithWebIdentity using the token obtained from the supplied stscreds.IdentityTokenRetriever.
func GetAssumeRoleWithWebIdentityConfigViaTokenRetriever(ctx context.Context, cfg *aws.Config, webID v1beta1.AssumeRoleWithWebIdentityOptions, tokenRetriever stscreds.IdentityTokenRetriever) (*aws.Config, error) {
	stsclient := sts.NewFromConfig(*cfg, stsRegionOrDefault(cfg.Region))
	awsConfig, err := config.LoadDefaultConfig(
		ctx,
		userAgentV2,
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(aws.NewCredentialsCache(
			stscreds.NewWebIdentityRoleProvider(
				stsclient,
				aws.ToString(webID.RoleARN),
				tokenRetriever,
				SetWebIdentityRoleOptions(webID),
			)),
		),
	)
	return &awsConfig, errors.Wrap(err, "failed to assume role via web identity")
}

// UseDefault loads the default AWS config with the specified region.
func UseDefault(ctx context.Context, region string) (*aws.Config, error) {
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
	return &cfg, nil
}

type xpWebIdentityTokenRetriever struct {
	ctx           context.Context
	kube          client.Client
	tokenSource   v1.CredentialsSource
	tokenSelector v1.CommonCredentialSelectors
}

func (x *xpWebIdentityTokenRetriever) GetIdentityToken() ([]byte, error) {
	token, err := resource.CommonCredentialExtractor(x.ctx, x.tokenSource, x.kube, x.tokenSelector)
	return token, errors.Wrap(err, "could not extract token from tokenSource")
}

// UseWebIdentityToken calls sts.AssumeRoleWithWebIdentity using
// the configuration supplied in ProviderConfig's
// spec.credentials.assumeRoleWithWebIdentity.
func UseWebIdentityToken(ctx context.Context, region string, pcs *v1beta1.ProviderConfigSpec, kube client.Client) (*aws.Config, error) {
	if pcs.Credentials.WebIdentity == nil {
		return nil, errors.New(`spec.credentials.webIdentity of ProviderConfig cannot be nil when the credential source is "WebIdentity"`)
	}

	// this is to preserve backward compatibility with
	// 0.x providers working with >=1.x ProviderConfig API
	// TODO: when configuring via AWS environment variable support is removed
	// tokenConfig should be mandatory and this should return an error
	if pcs.Credentials.WebIdentity.TokenConfig == nil {
		cfg, err := UseDefault(ctx, region)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get default AWS config")
		}
		return GetAssumeRoleWithWebIdentityConfig(ctx, cfg, *pcs.Credentials.WebIdentity, os.Getenv(envWebIdentityTokenFile))
	}

	// new behavior with tokenConfig in
	// spec.credentials.webIdentity.tokenConfig
	// the new behavior with tokenConfig does not rely on
	// the AWS environment variables AWS_WEB_IDENTITY_TOKEN_FILE
	// and AWS_ROLE_ARN.
	// However, we start by constructing a default AWS config and
	// AWS SDK enforces that when AWS_WEB_IDENTITY_TOKEN_FILE environment
	// variable is set AWS_ROLE_ARN must be present.
	// Otherwise, constructing the default AWS config fails.
	// Hence, either both env vars must be set
	// (to support the case where the controller pod has extra AWS IRSA config
	// which should be automatically injecting AWS_WEB_IDENTITY_TOKEN_FILE
	// and AWS_ROLE_ARN environment variables already)
	// or AWS_WEB_IDENTITY_TOKEN_FILE must not exist at all.
	_, foundTokenEnv := os.LookupEnv(envWebIdentityTokenFile)
	_, foundRoleArnEnv := os.LookupEnv(envWebIdentityRoleARN)
	if foundTokenEnv && !foundRoleArnEnv {
		return nil, errors.Errorf("if you intend to use IRSA together with WebIdentity auth, environment variable %s must be set together with %s. If only WebIdentity auth without any IRSA configuration is intended, %s must be unset",
			envWebIdentityRoleARN, envWebIdentityTokenFile, envWebIdentityTokenFile)
	}

	cfg, err := UseDefault(ctx, region)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get default AWS config")
	}
	tokenRetriever := &xpWebIdentityTokenRetriever{
		ctx:         ctx,
		kube:        kube,
		tokenSource: pcs.Credentials.WebIdentity.TokenConfig.Source,
		tokenSelector: v1.CommonCredentialSelectors{
			Fs:        pcs.Credentials.WebIdentity.TokenConfig.Fs,
			SecretRef: pcs.Credentials.WebIdentity.TokenConfig.SecretRef,
		},
	}

	return GetAssumeRoleWithWebIdentityConfigViaTokenRetriever(ctx, cfg, *pcs.Credentials.WebIdentity, tokenRetriever)
}

// UseUpbound calls sts.AssumeRoleWithWebIdentity using the configuration
// supplied in ProviderConfig's spec.credentials.assumeRoleWithWebIdentity and
// the identity supplied by the injected Upbound OIDC token.
// NOTE(hasheddan): this is the same functionality used for generic web identity
// token role assumption, but uses fields under Upbound in the ProviderConfig
// spec and the dedicated Upbound token injection path. This allows for clear
// separation of intent by users when exercising the functionality, and allows
// for uniformity across ProviderConfigs from other providers.
func UseUpbound(ctx context.Context, region string, pcs *v1beta1.ProviderConfigSpec) (*aws.Config, error) {
	cfg, err := UseDefault(ctx, region)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get default AWS config ")
	}
	if pcs.Credentials.Upbound == nil || pcs.Credentials.Upbound.WebIdentity == nil {
		return nil, errors.New(`spec.credentials.upbound.webIdentity of ProviderConfig cannot be nil when the credential source is "Upbound"`)
	}
	return GetAssumeRoleWithWebIdentityConfig(ctx, cfg, *pcs.Credentials.Upbound.WebIdentity, upboundProviderIdentityTokenFile)
}

// ExtractSecretValue extracts the value from a secret using a v1.SecretKeySelector, to reuse resource.ExtractSecret that does it using v1.CommonCredentialSelectors
func ExtractSecretValue(ctx context.Context, kube client.Client, selector *v1.SecretKeySelector) ([]byte, error) {
	if selector == nil {
		return nil, errors.New("SecretKeySelector is nil")
	}

	// Wrap SecretKeySelector in CommonCredentialSelectors
	credSelectors := v1.CommonCredentialSelectors{
		SecretRef: selector,
	}

	// Reuse ExtractSecret function (https://github.com/crossplane/crossplane-runtime/blob/19d95a69cc03690c4b867ff91d89681fcf872a93/pkg/resource/providerconfig.go#L77-L87)
	return resource.ExtractSecret(ctx, kube, credSelectors)
}

// SetAssumeRoleOptions sets options when Assuming an IAM Role
func SetAssumeRoleOptions(ctx context.Context, kube client.Client, aro v1beta1.AssumeRoleOptions) func(*stscreds.AssumeRoleOptions) {
	return func(opt *stscreds.AssumeRoleOptions) {
		if aro.ExternalID != nil {
			opt.ExternalID = aro.ExternalID
		}
		if aro.ExternalIDSecretRef != nil {
			// Fetch secret value
			secretValue, err := ExtractSecretValue(ctx, kube, aro.ExternalIDSecretRef)
			if err != nil {
				fmt.Printf("Error fetching ExternalID from secret: %v\n", err)
				return
			}
			secretValueStr := string(secretValue)
			opt.ExternalID = &secretValueStr
		}

		for _, t := range aro.Tags {
			opt.Tags = append(
				opt.Tags,
				stscredstypesv2.Tag{
					Key:   t.Key,
					Value: t.Value,
				})
		}
		opt.TransitiveTagKeys = append(opt.TransitiveTagKeys, aro.TransitiveTagKeys...)
	}
}

// SetWebIdentityRoleOptions sets options when exchanging a WebIdentity Token for a Role
func SetWebIdentityRoleOptions(opts v1beta1.AssumeRoleWithWebIdentityOptions) func(*stscreds.WebIdentityRoleOptions) {
	return func(opt *stscreds.WebIdentityRoleOptions) {
		opt.RoleSessionName = opts.RoleSessionName
	}
}

// LateInitializeStringPtr returns in if it's non-nil, otherwise returns from
// which is the backup for the cases in is nil.
func LateInitializeStringPtr(in *string, from *string) *string {
	if in != nil {
		return in
	}
	return from
}
