package sqs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the sqs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_sqs_queue_policy", func(r *config.Resource) {
		r.References["queue_url"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue",
			Extractor: common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_sqs_queue", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["url"].(string); ok {
				conn["url"] = []byte(a)
			}
			return conn, nil
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name_prefix"},
		}
	})

	p.AddResourceConfigurator("aws_sqs_queue_redrive_policy", func(r *config.Resource) {
		r.References["queue_url"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue",
			Extractor: common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_sqs_queue_redrive_allow_policy", func(r *config.Resource) {
		r.References["queue_url"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue",
			Extractor: common.PathTerraformIDExtractor,
		}
	})
}
