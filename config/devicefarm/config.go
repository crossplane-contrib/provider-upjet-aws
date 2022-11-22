/*
Copyright 2022 Upbound Inc.
*/

package devicefarm

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for devicefarm group.
func Configure(p *config.Provider) { // nolint:gocyclo
	p.AddResourceConfigurator("aws_devicefarm_test_grid_project", func(r *config.Resource) {
		r.References["vpc_config.subnet_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
	})

	p.AddResourceConfigurator("aws_devicefarm_test_grid_project", func(r *config.Resource) {
		r.References["vpc_config.security_group_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
	})
}
