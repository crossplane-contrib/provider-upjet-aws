// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package wafv2

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the wafv2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_wafv2_web_acl", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "rule")
		l := r.TFListConversionPaths()
		for _, e := range l {
			if strings.HasPrefix(e, "rule[*].") {
				r.RemoveSingletonListConversion(e)
			}
		}
		r.MetaResource.ArgumentDocs["rule_json"] = "A raw JSON string used to define the rules for allowing, blocking, or counting web requests. When this field is used, Crossplane cannot observe changes in the configuration through the AWS API; therefore, drift detection cannot be performed. Refer to the AWS documentation for the expected JSON structure: https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateWebACL.html"
		r.MetaResource.Description = "Creates a WAFv2 Web ACL resource. The 'rule' field is not supported due to Kubernetes CRD size limitations with deeply nested fields. Please use the 'ruleJson' field to define rules."
	})
	p.AddResourceConfigurator("aws_wafv2_rule_group", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "rule")
		l := r.TFListConversionPaths()
		for _, e := range l {
			if strings.HasPrefix(e, "rule[*].") {
				r.RemoveSingletonListConversion(e)
			}
		}
		r.MetaResource.Description = "Creates a WAFv2 rule group resource. The 'rule' field is not supported due to Kubernetes CRD size limitations with deeply nested fields. Please use the 'ruleJson' field to define rules."
		r.TerraformResource.Schema["rule_json"].Description = "A raw JSON string used to define the rules for allowing, blocking, or counting web requests. When this field is used, Crossplane cannot observe changes in the configuration through the AWS API; therefore, drift detection cannot be performed. Refer to the AWS documentation for the expected JSON structure: https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html"
	})
	p.AddResourceConfigurator("aws_wafv2_web_acl_rule_group_association", func(r *config.Resource) {
		registerWebACLRuleGroupAssociationSingletonListConversions(r)
	})
}

func registerWebACLRuleGroupAssociationSingletonListConversions(r *config.Resource) {
	r.AddSingletonListConversion("managed_rule_group", "managedRuleGroup")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs", "managedRuleGroup[*].managedRuleGroupConfigs")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].request_inspection", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].requestInspection")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].request_inspection[*].address_fields", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].requestInspection[*].addressFields")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].request_inspection[*].email_field", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].requestInspection[*].emailField")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].request_inspection[*].password_field", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].requestInspection[*].passwordField")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].request_inspection[*].phone_number_fields", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].requestInspection[*].phoneNumberFields")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].request_inspection[*].username_field", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].requestInspection[*].usernameField")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].response_inspection", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].responseInspection")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].response_inspection[*].body_contains", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].responseInspection[*].bodyContains")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].response_inspection[*].header", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].responseInspection[*].header")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].response_inspection[*].json", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].responseInspection[*].json")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_acfp_rule_set[*].response_inspection[*].status_code", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAcfpRuleSet[*].responseInspection[*].statusCode")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_anti_ddos_rule_set", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAntiDdosRuleSet")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_anti_ddos_rule_set[*].client_side_action_config", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAntiDdosRuleSet[*].clientSideActionConfig")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_anti_ddos_rule_set[*].client_side_action_config[*].challenge", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAntiDdosRuleSet[*].clientSideActionConfig[*].challenge")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set[*].request_inspection", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet[*].requestInspection")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set[*].request_inspection[*].password_field", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet[*].requestInspection[*].passwordField")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set[*].request_inspection[*].username_field", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet[*].requestInspection[*].usernameField")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set[*].response_inspection", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet[*].responseInspection")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set[*].response_inspection[*].body_contains", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet[*].responseInspection[*].bodyContains")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set[*].response_inspection[*].header", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet[*].responseInspection[*].header")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set[*].response_inspection[*].json", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet[*].responseInspection[*].json")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_atp_rule_set[*].response_inspection[*].status_code", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesAtpRuleSet[*].responseInspection[*].statusCode")
	r.AddSingletonListConversion("managed_rule_group[*].managed_rule_group_configs[*].aws_managed_rules_bot_control_rule_set", "managedRuleGroup[*].managedRuleGroupConfigs[*].awsManagedRulesBotControlRuleSet")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use", "managedRuleGroup[*].ruleActionOverride[*].actionToUse")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].allow", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].allow")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].allow[*].custom_request_handling", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].allow[*].customRequestHandling")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].block", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].block")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].block[*].custom_response", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].block[*].customResponse")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].captcha", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].captcha")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].captcha[*].custom_request_handling", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].captcha[*].customRequestHandling")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].challenge", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].challenge")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].challenge[*].custom_request_handling", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].challenge[*].customRequestHandling")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].count", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].count")
	r.AddSingletonListConversion("managed_rule_group[*].rule_action_override[*].action_to_use[*].count[*].custom_request_handling", "managedRuleGroup[*].ruleActionOverride[*].actionToUse[*].count[*].customRequestHandling")
	r.AddSingletonListConversion("rule_group_reference", "ruleGroupReference")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use", "ruleGroupReference[*].ruleActionOverride[*].actionToUse")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].allow", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].allow")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].allow[*].custom_request_handling", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].allow[*].customRequestHandling")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].block", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].block")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].block[*].custom_response", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].block[*].customResponse")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].captcha", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].captcha")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].captcha[*].custom_request_handling", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].captcha[*].customRequestHandling")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].challenge", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].challenge")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].challenge[*].custom_request_handling", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].challenge[*].customRequestHandling")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].count", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].count")
	r.AddSingletonListConversion("rule_group_reference[*].rule_action_override[*].action_to_use[*].count[*].custom_request_handling", "ruleGroupReference[*].ruleActionOverride[*].actionToUse[*].count[*].customRequestHandling")
	r.AddSingletonListConversion("visibility_config", "visibilityConfig")
}
