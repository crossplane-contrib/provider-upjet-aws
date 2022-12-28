/*
Copyright 2022 Upbound Inc.
*/

package networkmanager

import (
	"github.com/upbound/upjet/pkg/config"
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
}
