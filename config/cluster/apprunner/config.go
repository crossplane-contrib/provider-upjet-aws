// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apprunner

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the apprunner group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_apprunner_vpc_connector", func(r *config.Resource) {
		r.References["subnets"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetRefs",
			SelectorFieldName: "SubnetSelector",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.UseAsync = true
	})
}
