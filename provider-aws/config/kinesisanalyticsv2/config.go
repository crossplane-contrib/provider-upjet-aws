package kinesisanalytics2

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for kinesisanalytics2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kinesisanalyticsv2_application", func(r *config.Resource) {
		r.References["application_configuration.application_code_configuration.code_content.s3_content_location.bucket_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket",
			Extractor: common.PathARNExtractor,
		}
		r.References["service_execution_role"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
	})
}
