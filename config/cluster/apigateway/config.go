// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apigateway

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the apigateway group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_api_gateway_rest_api", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "policy")
	})

	p.AddResourceConfigurator("aws_api_gateway_domain_name_access_association", func(r *config.Resource) {
		r.References["domain_name_arn"] = config.Reference{
			TerraformName: "aws_api_gateway_domain_name",
			Extractor:     common.PathARNExtractor,
		}
		r.References["access_association_source"] = config.Reference{
			TerraformName: "aws_vpc_endpoint",
		}
	})

	p.AddResourceConfigurator("aws_api_gateway_vpc_link", func(r *config.Resource) {
		r.References["target_arns"] = config.Reference{
			TerraformName:     "aws_lb",
			RefFieldName:      "TargetArnRefs",
			SelectorFieldName: "TargetArnSelector",
			Extractor:         common.PathARNExtractor,
		}
		r.UseAsync = true
	})
}
