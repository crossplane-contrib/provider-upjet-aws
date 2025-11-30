// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package autoscaling

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the autoscaling group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_autoscaling_group", func(r *config.Resource) {
		// These are mutually exclusive with aws_autoscaling_attachment.
		config.MoveToStatus(r.TerraformResource, "load_balancers", "target_group_arns")

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"availability_zones",
			},
		}

		r.References["vpc_zone_identifier"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		delete(r.References, "launch_template.version")
		r.UseAsync = true

	})
	p.AddResourceConfigurator("aws_autoscaling_attachment", func(r *config.Resource) {
		r.References["autoscaling_group_name"] = config.Reference{
			TerraformName: "aws_autoscaling_group",
		}
		r.References["alb_target_group_arn"] = config.Reference{
			TerraformName: "aws_lb_target_group",
			Extractor:     common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_autoscaling_group_tag", func(r *config.Resource) {
		r.References["autoscaling_group_name"] = config.Reference{
			TerraformName: "aws_autoscaling_group",
		}
	})
}
