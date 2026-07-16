// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elbv2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/crossplane/upjet/v2/pkg/config"
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

		r.ServerSideApplyMergeStrategies["default_action"] = config.MergeStrategy{
			ListMergeStrategy: config.ListMergeStrategy{
				ListMapKeys: config.ListMapKeys{
					InjectedKey: config.InjectedKey{
						Key:          "index",
						DefaultValue: "default",
					},
				},
				MergeStrategy: config.ListTypeMap,
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

	p.AddResourceConfigurator("aws_lb_listener_rule", func(r *config.Resource) {
		r.TerraformCustomDiff = lbListenerRuleCustomDiff
	})

	p.AddResourceConfigurator("aws_lb_target_group", func(r *config.Resource) {
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Optional = false
			s.ForceNew = true
			s.Computed = false
		}
		r.LateInitializer.IgnoredFields = []string{"target_failover"}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			// skip no diff or destroy diffs
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}

			// ignore diff due to defaulting in the TF schema
			udiDiff, ok := diff.Attributes["target_health_state.0.unhealthy_draining_interval"]
			if ok && udiDiff.Old == "" && udiDiff.New == "0" {
				delete(diff.Attributes, "target_health_state.0.unhealthy_draining_interval")
			}
			return diff, nil
		}
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

func lbListenerRuleCustomDiff(diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig) (*terraform.InstanceDiff, error) { //nolint:gocyclo // easier to follow as a unit
	if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
		return diff, nil
	}
	for k, attrDiff := range diff.Attributes {
		if attrDiff == nil || !strings.HasPrefix(k, "action.") || !strings.HasSuffix(k, ".target_group_arn") {
			continue
		}
		idx := strings.TrimPrefix(k, "action.")
		idx = strings.TrimSuffix(idx, ".target_group_arn")

		// Case A: arn absent from state, present in config — late-init artifact
		// when the provider's Read took the flattenForwardActionOneOf path because
		// forward is configured, stripping arn from state while late-init had
		// already written it to spec.
		if attrDiff.Old == "" && attrDiff.New != "" && configActionHasForward(cfg, idx) {
			delete(diff.Attributes, k)
		}

		// Case B: arn present in state, absent from config — the provider's Read
		// took the flattenForwardActionBoth path (null RawPlan after Update or
		// restart), writing arn to state, while spec only declares forward.
		if attrDiff.Old != "" && attrDiff.New == "" && attrDiff.NewRemoved {
			forwardKey := fmt.Sprintf("action.%s.forward.#", idx)
			if state != nil && state.Attributes[forwardKey] != "" && state.Attributes[forwardKey] != "0" {
				delete(diff.Attributes, k)
			}
		}
	}
	return diff, nil
}

func configActionHasForward(cfg *terraform.ResourceConfig, actionIdx string) bool {
	if cfg == nil || cfg.Config == nil {
		return false
	}
	actions, ok := cfg.Config["action"].([]interface{})
	if !ok {
		return false
	}
	idx, err := strconv.Atoi(actionIdx)
	if err != nil || idx < 0 || idx >= len(actions) {
		return false
	}
	actionMap, ok := actions[idx].(map[string]interface{})
	if !ok {
		return false
	}
	forwards, _ := actionMap["forward"].([]interface{})
	return len(forwards) > 0
}
