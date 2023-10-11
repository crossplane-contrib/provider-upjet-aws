package dax

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the dax group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_dax_cluster", func(r *config.Resource) {
		r.UseAsync = true
	})
}
