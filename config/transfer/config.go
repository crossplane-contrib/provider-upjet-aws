package transfer

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for transfer group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_transfer_user", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			Type: "Server",
		}
		r.References["role"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
	})
}
