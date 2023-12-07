package kinesisanalyticsv2

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the kinesisanalyticsv2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kinesisanalyticsv2_application", func(r *config.Resource) {
		r.References["application_configuration.application_code_configuration.code_content.s3_content_location.bucket_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
			Extractor: common.PathARNExtractor,
		}
		r.References["service_execution_role"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["application_configuration.sql_application_configuration.reference_data_source.s3_reference_data_source.bucket_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
			Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)`,
		}
		r.References["application_configuration.sql_application_configuration.input.kinesis_streams_input.resource_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
