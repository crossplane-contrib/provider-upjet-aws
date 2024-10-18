// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package globalaccelerator

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the globalaccelerator group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_globalaccelerator_endpoint_group", func(r *config.Resource) {
		r.References = config.References{
			"listener_arn": {
				TerraformName: "aws_globalaccelerator_listener",
			},
		}
	})

	p.AddResourceConfigurator("aws_globalaccelerator_listener", func(r *config.Resource) {
		r.References = config.References{
			"accelerator_arn": {
				TerraformName: "aws_globalaccelerator_accelerator",
			},
		}
	})
}
