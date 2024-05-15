// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kinesisanalytics2

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the kinesisanalytics2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kinesisanalyticsv2_application", func(r *config.Resource) {
		r.References["application_configuration.application_code_configuration.code_content.s3_content_location.bucket_arn"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     common.PathARNExtractor,
		}
		r.References["service_execution_role"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["application_configuration.sql_application_configuration.reference_data_source.s3_reference_data_source.bucket_arn"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)`,
		}
		r.References["application_configuration.sql_application_configuration.input.kinesis_streams_input.resource_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
