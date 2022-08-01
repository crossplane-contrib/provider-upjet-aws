package elb

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for elb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elb", func(r *config.Resource) {
		r.References["instances"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Instance",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"access_logs"},
		}
	})

	p.AddResourceConfigurator("aws_elb_attachment", func(r *config.Resource) {
		r.References["elb"] = config.Reference{
			Type: "ELB",
		}
		r.References["instance"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Instance",
		}
	})
}
