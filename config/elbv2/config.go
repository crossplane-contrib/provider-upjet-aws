package elbv2

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for elbv2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lb", func(r *config.Resource) {
		r.UseAsync = true
		r.LateInitializer.IgnoredFields = []string{"access_logs"}
	})
	p.AddResourceConfigurator("aws_lb_target_group_attachment", func(r *config.Resource) {
		r.References = config.References{
			"target_id": config.Reference{
				Type: "github.com/upbound/provider-aws/apis/elbv2/v1beta1.LB",
			},
			"target_group_arn": config.Reference{
				Type: "github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup",
			},
		}
		r.UseAsync = true
	})
}
