package emr

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_emr_cluster", func(r *config.Resource) {
		r.References["service_role"] = config.Reference{
			TerraformName: "aws_iam_role",
		}
		r.References["ec2_attributes.instance_profile"] = config.Reference{
			TerraformName: "aws_iam_instance_profile",
		}
		r.References["ec2_attributes.subnet_id"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["ec2_attributes.key_name"] = config.Reference{
			TerraformName: "aws_key_pair",
		}
		r.References["ec2_attributes.emr_managed_master_security_group"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["ec2_attributes.emr_managed_slave_security_group"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["ec2_attributes.additional_master_security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["ec2_attributes.additional_slave_security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["log_uri"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"core_instance_fleet",
				"master_instance_fleet",
			},
		}
	})
	p.AddResourceConfigurator("aws_emr_instance_group", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			TerraformName: "aws_emr_cluster",
		}
	})
	p.AddResourceConfigurator("aws_emr_instance_fleet", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			TerraformName: "aws_emr_cluster",
		}
	})
}
