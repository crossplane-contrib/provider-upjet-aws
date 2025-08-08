// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package qldb

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the qldb group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_qldb_stream", func(r *config.Resource) {
		r.References["kinesis_configuration.stream_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
		r.References["ledger_name"] = config.Reference{
			TerraformName: "aws_qldb_ledger",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
