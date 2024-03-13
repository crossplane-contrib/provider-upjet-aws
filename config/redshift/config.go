// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package redshift

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the redshift group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_redshift_cluster", func(r *config.Resource) {
		r.UseAsync = true
	})
}
