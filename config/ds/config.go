// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ds

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the ds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_directory_service_directory", func(r *config.Resource) {
		r.References["vpc_settings.subnet_ids"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["connect_settings.subnet_ids"] = config.Reference{
			TerraformName: "aws_subnet",
		}
	})
}
