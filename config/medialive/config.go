// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package medialive

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the medialive group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_medialive_multiplex", func(r *config.Resource) {
		r.Path = "multiplices"
	})
}
