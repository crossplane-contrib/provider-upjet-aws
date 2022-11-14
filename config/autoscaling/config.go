/*
Copyright 2021 Upbound Inc.
*/

package autoscaling

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for autoscaling group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_autoscaling_group", func(r *config.Resource) {
		// These are mutually exclusive with aws_autoscaling_attachment.
		config.MoveToStatus(r.TerraformResource, "load_balancers", "target_group_arns")

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"availability_zones",
			},
		}

		r.References["vpc_zone_identifier"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
		delete(r.References, "launch_template.version")
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_autoscaling_attachment", func(r *config.Resource) {
		r.References["autoscaling_group_name"] = config.Reference{
			Type: "AutoscalingGroup",
		}
		r.References["alb_target_group_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup",
			Extractor: common.PathARNExtractor,
		}
	})
}
