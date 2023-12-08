package iot

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the iot group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_iot_topic_rule_destination", func(r *config.Resource) {
		r.References["vpc_configuration.security_groups"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["vpc_configuration.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
	})
}
