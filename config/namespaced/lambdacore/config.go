// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package lambdacore

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

// Configure adds configurations for the lambdacore group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lambdacore_network_connector", func(r *config.Resource) {
		r.References["configuration.vpc_egress_configuration.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
		r.References["configuration.vpc_egress_configuration.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.References["operator_role"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("configuration", "configuration")
		r.AddSingletonListConversion("configuration[*].vpc_egress_configuration", "configuration[*].vpcEgressConfiguration")
	})
}
