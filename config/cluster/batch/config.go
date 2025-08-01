// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package batch

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the batch group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_batch_compute_environment", func(r *config.Resource) {
		r.References["compute_resources.subnets"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["compute_resources.security_group_ids"] = config.Reference{
			TerraformName: "aws_security_group",
		}
	})
}
