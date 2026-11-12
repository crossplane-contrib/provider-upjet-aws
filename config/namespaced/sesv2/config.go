// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sesv2

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for the sesv2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_sesv2_tenant_resource_association", func(r *config.Resource) {
		r.References["tenant_name"] = config.Reference{
			TerraformName: "aws_sesv2_tenant",
		}
	})
}
