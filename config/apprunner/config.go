// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apprunner

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the apprunner group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_apprunner_vpc_connector", func(r *config.Resource) {
		r.References["subnets"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
			RefFieldName:      "SubnetRefs",
			SelectorFieldName: "SubnetSelector",
		}
		r.References["security_groups"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.UseAsync = true
	})
}
