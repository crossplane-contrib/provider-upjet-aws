package wafv2

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for wafv2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_wafv2_web_acl_association", func(r *config.Resource) {
		r.References = config.References{
			"resource_arn": config.Reference{
				Type: "github.com/upbound/provider-aws/apis/elbv2/v1beta1.LB",
			},
			"web_acl_arn": config.Reference{
				Type: "github.com/upbound/provider-aws/apis/waf/v1beta1.WebACL",
			},
		}
		r.UseAsync = true
	})
}
