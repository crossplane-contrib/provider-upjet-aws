package emr

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_emr_cluster", func(r *config.Resource) {
		r.References["ec2_attributes.additional_master_security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["ec2_attributes.additional_slave_security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["ec2_attributes.key_name"] = config.Reference{
			TerraformName: "aws_key_pair",
		}
		r.References["log_uri"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"master_instance_fleet", // Cannot be specified with master_instance_group
				"core_instance_fleet",   // Cannot be specified with core_instance_group
				"configurations_json",   // Alternative to configurations
			},
		}
	})
}
