package sqs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for sns group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_sqs_queue_policy", func(r *config.Resource) {
		r.References["queue_url"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue",
			Extractor: common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_sqs_queue", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name_prefix"},
		}
	})
}
