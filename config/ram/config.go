/*
Copyright 2021 Upbound Inc.
*/

package ram

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for the ram group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ram_resource_association", func(r *config.Resource) {
		delete(r.References, "resource_arn")
	})
}
