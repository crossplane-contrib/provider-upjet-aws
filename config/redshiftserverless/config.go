// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package redshiftserverless

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for redshiftserverless group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_redshiftserverless_namespace", func(r *config.Resource) {
		r.Kind = "RedshiftServerlessNamespace"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"kms_key_id"},
		}
	})

	p.AddResourceConfigurator("aws_redshiftserverless_workgroup", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"config_parameter"},
		}
	})
}
