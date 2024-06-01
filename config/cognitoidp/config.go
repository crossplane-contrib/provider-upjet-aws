// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cognitoidp

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-upjet-aws/config/common"
)

// Configure adds configurations for the cognitoidp group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cognito_user_pool_client", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			TerraformName: "aws_cognito_user_pool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_user_pool_domain", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			TerraformName: "aws_cognito_user_pool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_user_group", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			TerraformName: "aws_cognito_user_pool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_resource_server", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			TerraformName: "aws_cognito_user_pool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_identity_provider", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			TerraformName: "aws_cognito_user_pool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_user_pool_ui_customization", func(r *config.Resource) {
		r.References["client_id"] = config.Reference{
			TerraformName: "aws_cognito_user_pool_client",
		}
		r.References["user_pool_id"] = config.Reference{
			TerraformName: "aws_cognito_user_pool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_user_pool", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"verification_message_template",
				"sms_verification_message",
				"email_verification_message",
				"email_verification_subject",
			},
		}
		r.References["lambda_config.create_auth_challenge"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.custom_email_sender.lambda_arn"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.custom_message"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.custom_sms_sender.lambda_arn"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.define_auth_challenge"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.post_authentication"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.post_confirmation"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.pre_authentication"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.pre_sign_up"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.pre_token_generation"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.user_migration"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["lambda_config.verify_auth_challenge_response"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.References["sms_configuration.sns_caller_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
	})
}
