/*
Copyright 2022 Upbound Inc.
*/

package directconnect

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for directconnect group.
func Configure(p *config.Provider) { // nolint:gocyclo
	p.AddResourceConfigurator("aws_dx_public_virtual_interface", func(r *config.Resource) {
		r.References["connection_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/directconnect/v1beta1.Connection",
		}
		r.UseAsync = true
	})
}
