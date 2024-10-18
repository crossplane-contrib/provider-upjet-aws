// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elbv2

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the elbv2 group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_lb", func(r *config.Resource) {
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		r.References = config.References{
			"security_groups": {
				TerraformName:     "aws_security_group",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
			"subnets": {
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"access_logs.bucket": {
				TerraformName: "aws_s3_bucket",
			},
			"subnet_mapping.subnet_id": {
				TerraformName: "aws_subnet",
			},
		}
		r.UseAsync = true
		r.LateInitializer.IgnoredFields = []string{"access_logs"}
	})

	p.AddResourceConfigurator("aws_lb_listener", func(r *config.Resource) {
		r.References = config.References{
			"load_balancer_arn": {
				TerraformName: "aws_lb",
			},
			"default_action.target_group_arn": {
				TerraformName: "aws_lb_target_group",
			},
			"default_action.forward.target_group.arn": {
				TerraformName: "aws_lb_target_group",
			},
		}

		// lb_listener schema allows to configure "default_action" with type
		// "forward", in 2 different ways.
		// 1. you can specify, default_action.0.forward, which allows configuring
		//    multiple target groups.
		// 2. you can specify default_action.0.target_group_arn if you want
		//    to configure only one target group.
		// Former is a more general way, and latter is more of a shortcut for
		// a specific case, which can already be expressed with #1.
		// TF implementation instructs to specify either #1 or #2, not both.
		// However, they both end up in the state redundantly and cause
		// unnecessary diff.
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) { //nolint:gocyclo
			// skip no diff or destroy diffs
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			duplicatedAction := ""
			for k, attrDiff := range diff.Attributes {
				// user specified the "default_action.0.target_group_arn" and
				// "default_action.0.forward" is not specified in config.
				// In that case, default_action.0.forward is populated
				// by the AWS API, which can cause an unnecessary diff,
				// trying to remove that "auto-populated" element.
				if regexp.MustCompile(`^default_action\.\d+\.forward\.#$`).MatchString(k) &&
					attrDiff.New == "0" && attrDiff.Old == "1" {
					delete(diff.Attributes, k)
					// save this attribute path to remove remaining diffs
					// of its nested fields if any in a second pass
					duplicatedAction = strings.TrimSuffix(k, ".#")
				}
				// this is the same case as above
				if regexp.MustCompile(`^default_action\.\d+\.forward\.\d+\.target_group\.#$`).MatchString(k) &&
					attrDiff.New == "0" && attrDiff.Old == "1" {
					delete(diff.Attributes, k)
				}
				// this is the same case but vice versa. user specified
				// forward target via default_action.0.forward, and
				// default_action.0.target_group_arn is omitted.
				// In that case, default_action.0.target_group_arn is populated
				// by the AWS API and ends up in the state, causing
				// an unnecessary diff, which we omit here.
				if regexp.MustCompile(`^default_action\.\d+\.target_group_arn$`).MatchString(k) &&
					attrDiff.New == "" && attrDiff.Old != "" && attrDiff.NewRemoved {
					delete(diff.Attributes, k)
				}
			}
			// if we have caught an unnecessary diff for default_action.0.forward, remove
			// any sub-element diffs of it.
			if duplicatedAction != "" {
				for k := range diff.Attributes {
					if strings.HasPrefix(k, duplicatedAction) {
						delete(diff.Attributes, k)
					}
				}
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("aws_lb_target_group", func(r *config.Resource) {
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Optional = false
			s.ForceNew = true
			s.Computed = false
		}
		r.LateInitializer.IgnoredFields = []string{"target_failover"}
	})

	p.AddResourceConfigurator("aws_lb_target_group_attachment", func(r *config.Resource) {
		r.References = config.References{
			"target_group_arn": {
				TerraformName: "aws_lb_target_group",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_lb_trust_store", func(r *config.Resource) {
		r.ShortGroup = "elbv2"
		r.Kind = "LBTrustStore"
	})
}
