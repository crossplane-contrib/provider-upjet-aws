// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package route53

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the route53 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53_traffic_policy_instance", func(r *config.Resource) {
		r.References["hosted_zone_id"] = config.Reference{
			TerraformName: "aws_route53_zone",
		}
		r.References["traffic_policy_id"] = config.Reference{
			TerraformName: "aws_route53_traffic_policy",
		}
	})
	p.AddResourceConfigurator("aws_route53_hosted_zone_dnssec", func(r *config.Resource) {
		r.References["hosted_zone_id"] = config.Reference{
			TerraformName: "aws_route53_zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_record", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			TerraformName: "aws_route53_zone",
		}
		r.References["health_check_id"] = config.Reference{
			TerraformName: "aws_route53_health_check",
		}
		delete(r.References, "alias.name")
		delete(r.References, "alias.zone_id")
	})
	p.AddResourceConfigurator("aws_route53_vpc_association_authorization", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			TerraformName: "aws_route53_zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_zone", func(r *config.Resource) {
		r.References["delegation_set_id"] = config.Reference{
			TerraformName: "aws_route53_delegation_set",
		}
		r.UseAsync = true
	})
}
