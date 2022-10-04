package cloudsearch

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for ebs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudsearch_domain", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_cloudsearch_domain_service_access_policy", func(r *config.Resource) {
		r.UseAsync = true
	})
}
