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
	})
}
