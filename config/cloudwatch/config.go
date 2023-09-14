package cloudwatch

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the cloudwatch group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudwatch_metric_stream", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "name")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name_prefix"},
		}
	})

	p.AddResourceConfigurator("aws_cloudwatch_event_target", func(r *config.Resource) {
		r.References["arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
