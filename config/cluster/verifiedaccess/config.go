// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package verifiedaccess

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the verifiedaccess group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_verifiedaccess_trust_provider", func(r *config.Resource) {
		r.RemoveSingletonListConversion("device_options")
		r.RemoveSingletonListConversion("oidc_options")
	})
	p.AddResourceConfigurator("aws_verifiedaccess_endpoint", func(r *config.Resource) {
		r.References["load_balancer_options.subnet_ids"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["cidr_options.subnet_ids"] = config.Reference{
			TerraformName: "aws_subnet",
		}
	})
}
