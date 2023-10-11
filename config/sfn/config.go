package sfn

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the sfn group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_sfn_state_machine", func(r *config.Resource) {
		r.References["role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
	})
}
