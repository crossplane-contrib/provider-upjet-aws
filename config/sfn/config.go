package sfn

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for sfn group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_sfn_state_machine", func(r *config.Resource) {
		r.References["role_arn"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
		}
	})
}
