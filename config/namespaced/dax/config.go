// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dax

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for the dax group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_dax_cluster", func(r *config.Resource) {
		r.UseAsync = true
	})
}
