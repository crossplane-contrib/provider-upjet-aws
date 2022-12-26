/*
Copyright 2022 Upbound Inc.
*/

package route53recoverycontrolconfig

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for route53recoverycontrolconfig group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53recoverycontrolconfig_control_panel", func(r *config.Resource) {
		r.References["cluster_arn"] = config.Reference{
			TerraformName: "aws_route53recoverycontrolconfig_cluster",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
