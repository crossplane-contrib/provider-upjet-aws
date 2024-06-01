// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cognitoidentity

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-upjet-aws/config/common"
)

// Configure adds configurations for the cognitoidentity group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cognito_identity_pool", func(r *config.Resource) {
		r.References["saml_provider_arns"] = config.Reference{
			TerraformName: "aws_iam_saml_provider",
			Extractor:     common.PathARNExtractor,
		}
		r.References["cognito_identity_providers.client_id"] = config.Reference{
			TerraformName: "aws_cognito_user_pool_client",
		}
	})
}
