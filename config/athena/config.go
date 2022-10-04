package athena

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for athena group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_athena_workgroup", func(r *config.Resource) {
		r.References["configuration.result_configuration.encryption_configuration.kms_key_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
	})
}
