package ssoadmin

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the ssoadmin group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ssoadmin_account_assignment", func(r *config.Resource) {
		r.References["permission_set_arn"] = config.Reference{
			TerraformName: "aws_ssoadmin_permission_set",
			Extractor: common.PathARNExtractor,
		}
	})
}
