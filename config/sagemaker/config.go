package sagemaker

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the sagemaker group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_sagemaker_workforce", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"source_ip_config"},
		}
	})
	p.AddResourceConfigurator("aws_sagemaker_device_fleet", func(r *config.Resource) {
		r.Path = "devicefleet"
	})
	p.AddResourceConfigurator("aws_sagemaker_endpoint", func(r *config.Resource) {
		r.References["endpoint_config_name"] = config.Reference{
			Type: "EndpointConfiguration",
		}
		r.UseAsync = true
	})
}
