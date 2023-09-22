package ds

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for the ds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_directory_service_directory", func(r *config.Resource) {
		r.References["vpc_settings.subnet_ids"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
		r.References["connect_settings.subnet_ids"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
	})
}
