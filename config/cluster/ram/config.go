// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ram

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the ram group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_ram_resource_association", func(r *config.Resource) {
		delete(r.References, "resource_arn")
	})
}
