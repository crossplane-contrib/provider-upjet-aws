/*
Copyright 2024 Upbound Inc.
*/

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
