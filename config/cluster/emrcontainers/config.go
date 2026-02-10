// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package emrcontainers

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the emrcontainers group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_emrcontainers_virtual_cluster", func(r *config.Resource) {
		r.References = config.References{
			"container_provider.id": {
				TerraformName: "aws_eks_cluster",
			},
		}
		r.UseAsync = true
	})
}
