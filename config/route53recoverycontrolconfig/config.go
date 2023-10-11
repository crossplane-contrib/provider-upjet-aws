/*
Copyright 2022 Upbound Inc.
*/

package route53recoverycontrolconfig

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the route53recoverycontrolconfig group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53recoverycontrolconfig_control_panel", func(r *config.Resource) {
		r.References["cluster_arn"] = config.Reference{
			TerraformName: "aws_route53recoverycontrolconfig_cluster",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
	p.AddResourceConfigurator("aws_route53recoverycontrolconfig_routing_control", func(r *config.Resource) {
		r.References["cluster_arn"] = config.Reference{
			TerraformName: "aws_route53recoverycontrolconfig_cluster",
			Extractor:     common.PathTerraformIDExtractor,
		}
		r.References["control_panel_arn"] = config.Reference{
			TerraformName: "aws_route53recoverycontrolconfig_control_panel",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
	p.AddResourceConfigurator("aws_route53recoverycontrolconfig_safety_rule", func(r *config.Resource) {
		r.References["control_panel_arn"] = config.Reference{
			TerraformName: "aws_route53recoverycontrolconfig_control_panel",
			Extractor:     common.PathTerraformIDExtractor,
		}
		r.References["asserted_controls"] = config.Reference{
			TerraformName: "aws_route53recoverycontrolconfig_routing_control",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
