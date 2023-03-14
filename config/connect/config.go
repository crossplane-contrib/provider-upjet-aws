/*
Copyright 2022 Upbound Inc.
*/

package connect

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for connect group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_connect_contact_flow", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_contact_flow_module", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_hours_of_operation", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_queue", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["hours_of_operation_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.HoursOfOperation",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("hours_of_operation_id",true)`,
		}
	})
	p.AddResourceConfigurator("aws_connect_quick_connect", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_routing_profile", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["default_outbound_queue_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Queue",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("queue_id",true)`,
		}
	})
	p.AddResourceConfigurator("aws_connect_security_profile", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_user_hierarchy_structure", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_vocabulary", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/connect/v1beta1.Instance",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
