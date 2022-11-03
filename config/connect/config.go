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
	p.AddResourceConfigurator("aws_connect_contact_flow_module", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"instance_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
				Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			},
		}
	})
	p.AddResourceConfigurator("aws_connect_hours_of_operation", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"instance_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
				Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			},
		}
	})
	p.AddResourceConfigurator("aws_connect_queue", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"instance_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
				Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			},
			"hours_of_operation_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.HoursOfOperation",
				Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("hours_of_operation_id",true)`,
			},
		}
	})
	p.AddResourceConfigurator("aws_connect_quick_connect", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"instance_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
				Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			},
		}
	})
	p.AddResourceConfigurator("aws_connect_routing_profile", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"instance_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
				Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			},
			"default_outbound_queue_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Queue",
				Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("queue_id",true)`,
			},
		}
	})
	p.AddResourceConfigurator("aws_connect_security_profile", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"instance_id": {
				Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
				Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			},
		}
	})
}
