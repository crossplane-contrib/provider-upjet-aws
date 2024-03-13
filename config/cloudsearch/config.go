// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudsearch

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the cloudsearch group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudsearch_domain", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_cloudsearch_domain_service_access_policy", func(r *config.Resource) {
		r.UseAsync = true
	})
}
