// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elbv2

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the elbv2 group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_lb", func(r *config.Resource) {
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		r.References = config.References{
			"security_groups": {
				TerraformName:     "aws_security_group",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
			"subnets": {
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"access_logs.bucket": {
				TerraformName: "aws_s3_bucket",
			},
			"subnet_mapping.subnet_id": {
				TerraformName: "aws_subnet",
			},
		}
		r.UseAsync = true
		r.LateInitializer.IgnoredFields = []string{"access_logs"}
	})

	p.AddResourceConfigurator("aws_lb_listener", func(r *config.Resource) {
		r.References = config.References{
			"load_balancer_arn": {
				TerraformName: "aws_lb",
			},
			"default_action.target_group_arn": {
				TerraformName: "aws_lb_target_group",
			},
			"default_action.forward.target_group.arn": {
				TerraformName: "aws_lb_target_group",
			},
		}
	})

	p.AddResourceConfigurator("aws_lb_target_group", func(r *config.Resource) {
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Optional = false
			s.ForceNew = true
			s.Computed = false
		}
		r.LateInitializer.IgnoredFields = []string{"target_failover"}
	})

	p.AddResourceConfigurator("aws_lb_target_group_attachment", func(r *config.Resource) {
		r.References = config.References{
			"target_group_arn": {
				TerraformName: "aws_lb_target_group",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_lb_trust_store", func(r *config.Resource) {
		r.ShortGroup = "elbv2"
		r.Kind = "LBTrustStore"
	})
}
