package apigateway

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for acm group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_api_gateway_rest_api", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "policy")
	})
}
