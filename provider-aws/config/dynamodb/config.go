/*
Copyright 2022 Upbound Inc.
*/

package dynamodb

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for the dynamodb group.
func Configure(p *config.Provider) {
	// currently needs an ARN reference for external name
	p.AddResourceConfigurator("aws_dynamodb_contributor_insights", func(r *config.Resource) {
		r.References = config.References{
			"table_name": config.Reference{
				Type: "Table",
			},
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_kinesis_streaming_destination", func(r *config.Resource) {
		r.References = config.References{
			"table_name": config.Reference{
				Type: "Table",
			},
			"stream_arn": config.Reference{
				Type: "github.com/upbound/official-providers/provider-aws/apis/kinesis/v1beta1.Stream",
			},
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_table_item", func(r *config.Resource) {
		r.References = config.References{
			"table_name": config.Reference{
				Type: "Table",
			},
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_tag", func(r *config.Resource) {
		r.References = config.References{
			"resource_arn": config.Reference{
				Type:      "Table",
				Extractor: common.PathARNExtractor,
			},
		}
	})

}
