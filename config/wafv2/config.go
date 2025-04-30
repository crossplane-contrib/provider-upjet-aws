// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package wafv2

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the sfn group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_wafv2_web_acl", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "rule")
		l := r.TFListConversionPaths()
		for _, e := range l {
			if strings.HasPrefix(e, "rule[*].") {
				r.RemoveSingletonListConversion(e)
			}
		}
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
		r.TerraformResource.Schema["rule_json"].Description = "Raw JSON string to allow more than three nested statements. Conflicts with rule attribute. This is for advanced use cases where more than 3 levels of nested statements are required. There is no drift detection at this time. If you use this attribute instead of rule, you will be foregoing drift detection. See the AWS documentation for the JSON structure."
	})
}
