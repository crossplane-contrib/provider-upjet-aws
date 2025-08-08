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
}
