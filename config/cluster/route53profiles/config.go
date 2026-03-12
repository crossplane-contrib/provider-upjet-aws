// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package route53profiles

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the route53resolver group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53profiles_resource_association", func(r *config.Resource) {
		// remove reference fields due to multiple resource types supported
		delete(r.References, "resource_arn")
	})
}
