package sagemaker

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for sagemaker group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_sagemaker_workforce", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"source_ip_config"},
		}
	})
}
