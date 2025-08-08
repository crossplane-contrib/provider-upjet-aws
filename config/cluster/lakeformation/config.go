// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package lakeformation

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the lakeformation group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_lakeformation_data_lake_settings", func(r *config.Resource) {
		delete(r.References, "create_database_default_permissions.principal")
		delete(r.References, "create_table_default_permissions.principal")
	})

	p.AddResourceConfigurator("aws_lakeformation_permissions", func(r *config.Resource) {
		delete(r.References, "principal")
		delete(r.References, "table_with_columns.database_name")
	})
}
