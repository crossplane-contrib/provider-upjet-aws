// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ecrpublic

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the ecrpublic group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_ecrpublic_repository", func(r *config.Resource) {
		// Deletion takes a while.
		r.UseAsync = true
	})
}
