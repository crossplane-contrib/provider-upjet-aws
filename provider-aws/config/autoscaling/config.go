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
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.References["vpc_zone_identifier"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.Subnet",
		}
		r.UseAsync = true

		// Managed by Attachment resource.
		if s, ok := r.TerraformResource.Schema["load_balancers"]; ok {
			s.Optional = false
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["target_group_arns"]; ok {
			s.Optional = false
			s.Computed = true
		}
	})
	p.AddResourceConfigurator("aws_autoscaling_attachment", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["autoscaling_group_name"] = config.Reference{
			Type: "AutoscalingGroup",
		}
		r.References["alb_target_group_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/elbv2/v1alpha2.LBTargetGroup",
			Extractor: common.PathARNExtractor,
		}
	})
}
