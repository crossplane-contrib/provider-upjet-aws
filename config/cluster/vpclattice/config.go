// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package vpclattice

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the vpclattice group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_vpclattice_target_group_attachment", func(r *config.Resource) {
		delete(r.References, "target.id")
	})
}
