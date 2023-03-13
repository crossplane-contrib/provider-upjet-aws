package rolesanywhere

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for rolesanywhere group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_rolesanywhere_profile", func(r *config.Resource) {
		r.References["role_arns"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
	})
}
