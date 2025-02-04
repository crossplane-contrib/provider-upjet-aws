// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package firehose

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the firehose group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_kinesis_firehose_delivery_stream", func(r *config.Resource) {
		r.TerraformResource.Schema["splunk_configuration"].Elem.(*schema.Resource).Schema["hec_token"].Sensitive = true

		r.References["extended_s3_configuration.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["extended_s3_configuration.bucket_arn"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     common.PathARNExtractor,
		}

		r.References["s3_configuration.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["s3_configuration.bucket_arn"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     common.PathARNExtractor,
		}

		r.References["redshift_configuration.s3_backup_configuration.bucket_arn"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     common.PathARNExtractor,
		}

		config.MoveToStatus(r.TerraformResource, "arn")

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"server_side_encryption",
				"version_id",
			},
		}

		delete(r.References, "extended_s3_configuration.data_format_conversion_configuration.schema_configuration.database_name")
	})
}
