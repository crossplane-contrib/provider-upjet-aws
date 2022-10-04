/*
Copyright 2022 Upbound Inc.
*/

package cognitoidentity

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for acm group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cognito_identity_pool", func(r *config.Resource) {
		r.References["saml_provider_arns"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.SAMLProvider",
			Extractor: common.PathARNExtractor,
		}
		r.References["cognito_identity_providers.client_id"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-aws/apis/cognitoidp/v1beta1.UserPoolClient",
		}
	})
}
