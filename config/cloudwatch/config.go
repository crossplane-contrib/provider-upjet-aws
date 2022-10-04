package cloudwatch

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for cloudwatch group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudwatch_metric_stream", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "name")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name_prefix"},
		}
	})
}
