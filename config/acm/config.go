// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package acm

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the acm group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_acm_certificate_validation", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"certificate_arn": {
				TerraformName: "aws_acm_certificate",
			},
		}
		// Deletion takes a while.
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_acm_certificate", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			// These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/crossplane-contrib/provider-upjet-aws/issues/464
			IgnoredFields: []string{
				"validation_method",
				"key_algorithm",
				"certificate_body",
				"options",
				"subject_alternative_names",
			},
		}
	})
}
