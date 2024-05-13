// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package devicefarm

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the devicefarm group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_devicefarm_test_grid_project", func(r *config.Resource) {
		r.References["vpc_config.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
	})

	p.AddResourceConfigurator("aws_devicefarm_test_grid_project", func(r *config.Resource) {
		r.References["vpc_config.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
	})
}
