package efs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the efs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_efs_mount_target", func(r *config.Resource) {
		r.UseAsync = true
		r.References["file_system_id"] = config.Reference{
			Type: "FileSystem",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
		r.References["security_groups"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
		}
		/*r.MetaResource.Examples[0].Dependencies["aws_efs_file_system.foo"] = `{"creation_token": "my-product-foo", "region": "us-west-1"}`
		if err := r.MetaResource.Examples[0].Dependencies.SetPathValue("aws_subnet.alpha", "availability_zone", "us-west-1b"); err != nil {
			panic(err)
		}*/
	})
	p.AddResourceConfigurator("aws_efs_access_point", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			Type: "FileSystem",
		}
		// r.MetaResource.Examples[0].Dependencies["aws_efs_file_system.foo"] = `{"creation_token": "my-product-foo", "region": "us-west-1"}`
	})
	p.AddResourceConfigurator("aws_efs_backup_policy", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			Type: "FileSystem",
		}
	})
	p.AddResourceConfigurator("aws_efs_file_system_policy", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			Type: "FileSystem",
		}
	})

	p.AddResourceConfigurator("aws_efs_file_system", func(r *config.Resource) {
		r.References["kms_key_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
	})
}
