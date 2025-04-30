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
	})
	p.AddResourceConfigurator("aws_wafv2_rule_group", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "rule")
		l := r.TFListConversionPaths()
		for _, e := range l {
			if strings.HasPrefix(e, "rule[*].") {
				r.RemoveSingletonListConversion(e)
			}
		}
	})
}
