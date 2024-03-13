// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudformation

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the cloudformation group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudformation_stack_set_instance", func(r *config.Resource) {
		r.TerraformConfigurationInjector = func(jsonMap map[string]any, params map[string]any) error {
			params["region"] = jsonMap["region"]
			return nil
		}
	})
}
