// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package mwaa

import (
	"github.com/crossplane/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_mwaa_environment", func(r *config.Resource) {
		r.References["network_configuration.subnet_ids"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["network_configuration.security_group_ids"] = config.Reference{
			TerraformName: "aws_security_group",
		}
	})

}
