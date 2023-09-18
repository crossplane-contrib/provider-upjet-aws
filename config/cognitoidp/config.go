/*
Copyright 2022 Upbound Inc.
*/

package cognitoidp

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for the cognitoidp group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cognito_user_pool_client", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			Type: "UserPool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_user_pool_domain", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			Type: "UserPool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_user_group", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			Type: "UserPool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_resource_server", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			Type: "UserPool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_identity_provider", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			Type: "UserPool",
		}
	})
	p.AddResourceConfigurator("aws_cognito_user_pool_ui_customization", func(r *config.Resource) {
		r.References["user_pool_id"] = config.Reference{
			Type: "UserPool",
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
	})
}
