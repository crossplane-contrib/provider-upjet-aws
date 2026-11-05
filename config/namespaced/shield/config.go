// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package shield

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

// Configure adds configurations for the shield group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_shield_protection", func(r *config.Resource) {
		r.References["resource_arn"] = config.Reference{
			TerraformName: "aws_cloudfront_distribution",
			Extractor:     common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_shield_protection_health_check_association", func(r *config.Resource) {
		r.References["shield_protection_id"] = config.Reference{
			TerraformName: "aws_shield_protection",
		}
		r.References["health_check_arn"] = config.Reference{
			TerraformName: "aws_route53_health_check",
			Extractor:     common.PathARNExtractor,
		}
	})
}
