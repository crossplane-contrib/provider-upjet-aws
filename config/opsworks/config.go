// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package opsworks

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the opsworks group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_opsworks_stack", func(r *config.Resource) {
		r.References["default_subnet_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
	})

	p.AddResourceConfigurator("aws_opsworks_instance", func(r *config.Resource) {
		r.References["layer_ids"] = config.Reference{
			Type: "CustomLayer",
		}
	})
}
