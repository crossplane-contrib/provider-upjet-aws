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
		r.References = map[string]config.Reference{
			"saml_provider_arns": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.SAMLProvider",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
