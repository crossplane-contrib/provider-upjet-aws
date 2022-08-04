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
		r.References = map[string]config.Reference{
			"user_pool_id": {
				Type: "UserPool",
			},
		}
	})
	p.AddResourceConfigurator("aws_cognito_user_pool_domain", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"user_pool_id": {
				Type: "UserPool",
			},
		}
	})
}
