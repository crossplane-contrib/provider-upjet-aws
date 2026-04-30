// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package route53

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure adds configurations for the route53 group.
func Configure(p *config.Provider) { //nolint:gocyclo
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
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			nameDiff, ok := diff.Attributes["name"]
			if ok {
				if strings.TrimSuffix(nameDiff.New, ".") == strings.TrimSuffix(nameDiff.Old, ".") {
					delete(diff.Attributes, "name")
				}
			}
			// When `records` are being modified, AWS SDK expects `ttl` to be
			// present in the request body. The underlying TF provider checks
			// the RawPlan to find `ttl` and inject it to request body, but
			// it can't find since we are converting the diff to RawPlan during
			// diff calculation. So, if `ttl` isn't being changed by user, we
			// need to inject it manually from the current state for a successful
			// update.
			for k := range diff.Attributes {
				if strings.HasPrefix(k, "records.") {
					if _, exists := diff.Attributes["ttl"]; !exists {
						if state != nil && state.Attributes != nil {
							if ttlState, ok := state.Attributes["ttl"]; ok {
								diff.Attributes["ttl"] = &terraform.ResourceAttrDiff{
									Old: ttlState,
									New: ttlState,
								}
							}
						}
					}
					break
				}
			}
			return diff, nil
		}
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
