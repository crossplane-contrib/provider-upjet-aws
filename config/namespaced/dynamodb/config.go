// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dynamodb

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

// Configure adds configurations for the dynamodb group.
func Configure(p *config.Provider) { //nolint:gocyclo
	// currently needs an ARN reference for external name
	p.AddResourceConfigurator("aws_dynamodb_contributor_insights", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_kinesis_streaming_destination", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
		}

		r.References["stream_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_table_item", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
		}
		delete(r.References, "hash_key")
	})

	p.AddResourceConfigurator("aws_dynamodb_resource_policy", func(r *config.Resource) {
		r.References["resource_arn"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
			Extractor:     common.PathARNExtractor,
		}
	})
}
