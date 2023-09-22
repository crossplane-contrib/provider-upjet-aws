package opensearch

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for the opensearch group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_opensearch_domain", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "access_policies")
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_opensearch_domain_policy", func(r *config.Resource) {
		r.References["domain_name"] = config.Reference{
			Type: "Domain",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_opensearch_domain_saml_options", func(r *config.Resource) {
		r.UseAsync = true
	})
}
