// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package s3tables

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the s3tables group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_s3tables_namespace", func(r *config.Resource) {
		r.Kind = "S3TablesNamespace"
	})
	p.AddResourceConfigurator("aws_s3tables_table", func(r *config.Resource) {
		r.AddSingletonListConversion("metadata", "metadata")
		r.AddSingletonListConversion("metadata[*].iceberg", "metadata[*].iceberg")
		r.AddSingletonListConversion("metadata[*].iceberg[*].schema", "metadata[*].iceberg[*].schema")
	})
	p.AddResourceConfigurator("aws_s3tables_table_bucket_replication", func(r *config.Resource) {
		r.AddSingletonListConversion("rule", "rule")
	})
	p.AddResourceConfigurator("aws_s3tables_table_replication", func(r *config.Resource) {
		r.AddSingletonListConversion("rule", "rule")
	})
}
