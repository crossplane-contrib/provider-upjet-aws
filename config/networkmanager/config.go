// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package networkmanager

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the networkmanager group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_networkmanager_link", func(r *config.Resource) {
		r.References["site_id"] = config.Reference{
			TerraformName: "aws_networkmanager_site",
		}
	})
	p.AddResourceConfigurator("aws_networkmanager_link_association", func(r *config.Resource) {
		r.References["device_id"] = config.Reference{
			TerraformName: "aws_networkmanager_device",
		}
	})
	p.AddResourceConfigurator("aws_networkmanager_vpc_attachment", func(r *config.Resource) {
		r.References["subnet_arns"] = config.Reference{
			TerraformName: "aws_subnet",
			Extractor:     common.PathARNExtractor,
		}
		r.References["core_network_id"] = config.Reference{
			TerraformName: "aws_networkmanager_core_network",
		}
	})
	p.AddResourceConfigurator("aws_networkmanager_connect_attachment", func(r *config.Resource) {
		r.References["core_network_id"] = config.Reference{
			TerraformName: "aws_networkmanager_core_network",
		}
	})
}
