package efs

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for efs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_efs_mount_target", func(r *config.Resource) {
		r.References = config.References{
			"file_system_id": config.Reference{
				Type: "FileSystem",
			},
			"subnet_id": config.Reference{
				Type: "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
			},
		}
	})
}
