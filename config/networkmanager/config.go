/*
Copyright 2022 Upbound Inc.
*/

package networkmanager

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for networkmanager group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_networkmanager_link", func(r *config.Resource) {
		r.References["site_id"] = config.Reference{
			Type: "Site",
		}
	})
	p.AddResourceConfigurator("aws_networkmanager_link_association", func(r *config.Resource) {
		r.References["device_id"] = config.Reference{
			Type: "Device",
		}
	})
	p.AddResourceConfigurator("aws_networkmanager_site_to_site_vpn_attachment", func(r *config.Resource) {
		r.References["core_network_id"] = config.Reference{
			Type: "CoreNetwork",
		}
	})
	p.AddResourceConfigurator("aws_networkmanager_vpc_attachment", func(r *config.Resource) {
		r.References["subnet_arns"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
			Extractor: common.PathARNExtractor,
		}
		r.References["core_network_id"] = config.Reference{
			Type: "CoreNetwork",
		}
	})
	p.AddResourceConfigurator("aws_networkmanager_connect_attachment", func(r *config.Resource) {
		r.References["core_network_id"] = config.Reference{
			Type: "CoreNetwork",
		}
	})
}
