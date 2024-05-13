// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elb

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the elb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elb", func(r *config.Resource) {
		r.References["instances"] = config.Reference{
			TerraformName: "aws_instance",
		}
		r.References["subnets"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"access_logs"},
		}
	})

	p.AddResourceConfigurator("aws_elb_attachment", func(r *config.Resource) {
		r.References["elb"] = config.Reference{
			TerraformName: "aws_elb",
		}
		r.References["instance"] = config.Reference{
			TerraformName: "aws_instance",
		}
	})
}
