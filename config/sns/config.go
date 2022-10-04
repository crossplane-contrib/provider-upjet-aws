package sns

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for sns group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_sns_topic_subscription", func(r *config.Resource) {
		r.References["endpoint"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue",
			Extractor: common.PathARNExtractor,
		}
		r.References["topic_arn"] = config.Reference{
			Type:      "Topic",
			Extractor: common.PathARNExtractor,
		}
	})
}
