/*
Copyright 2022 Upbound Inc.
*/

package cognitoidp

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for acm group.
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
}
