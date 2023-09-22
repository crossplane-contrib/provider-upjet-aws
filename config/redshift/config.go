package redshift

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for the redshift group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_redshift_cluster", func(r *config.Resource) {
		r.UseAsync = true
	})
}
