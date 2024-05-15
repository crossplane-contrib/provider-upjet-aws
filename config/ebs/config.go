// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ebs

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the ebs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ebs_volume", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"kms_key_id": {
				TerraformName: "aws_kms_key",
			},
		}
	})
}
