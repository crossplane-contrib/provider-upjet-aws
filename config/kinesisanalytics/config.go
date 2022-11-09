package kinesisanalytics

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for kinesisanalytics group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kinesis_analytics_application", func(r *config.Resource) {
		r.References["inputs.kinesis_stream.resource_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
		r.References["inputs.kinesis_stream.role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
	})
}
