/*
Copyright 2021 Upbound Inc.
*/

package autoscaling

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for autoscaling group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_autoscaling_group", func(r *config.Resource) {
		// These are mutually exclusive with aws_autoscaling_attachment.
		common.MutuallyExclusiveFields(r.TerraformResource, "load_balancers", "target_group_arns")

		r.References["vpc_zone_identifier"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_autoscaling_attachment", func(r *config.Resource) {
		r.References["autoscaling_group_name"] = config.Reference{
			Type: "AutoscalingGroup",
		}
		r.References["alb_target_group_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/elbv2/v1alpha2.LBTargetGroup",
			Extractor: common.PathARNExtractor,
		}
	})
}
