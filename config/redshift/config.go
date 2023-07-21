package redshift

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for redshift group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_redshift_cluster", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_redshiftserverless_namespace", func(r *config.Resource) {
		r.Kind = "RedshiftServerlessNamespace"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"kms_key_id"},
		}
	})

	p.AddResourceConfigurator("aws_redshiftserverless_workgroup", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"config_parameter"},
		}
	})
}
