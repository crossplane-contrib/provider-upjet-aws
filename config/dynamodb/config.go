// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dynamodb

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the dynamodb group.
func Configure(p *config.Provider) {
	// currently needs an ARN reference for external name
	p.AddResourceConfigurator("aws_dynamodb_contributor_insights", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			Type: "Table",
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_kinesis_streaming_destination", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			Type: "Table",
		}

		r.References["stream_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_table_item", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			Type: "Table",
		}
		delete(r.References, "hash_key")
	})
}
