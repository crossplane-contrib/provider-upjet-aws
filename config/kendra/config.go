// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kendra

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the kendra group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kendra_thesaurus", func(r *config.Resource) {
		r.Path = "thesaurus"
	})
}
