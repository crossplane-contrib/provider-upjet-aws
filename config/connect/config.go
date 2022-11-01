/*
Copyright 2022 Upbound Inc.
*/

package connect

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for acm group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_connect_contact_flow", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"instance_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
				Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			},
		}
	})
}
