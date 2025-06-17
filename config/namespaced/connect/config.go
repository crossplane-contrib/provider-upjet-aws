// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package connect

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the connect group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_connect_contact_flow", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_contact_flow_module", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_hours_of_operation", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_queue", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["hours_of_operation_id"] = config.Reference{
			TerraformName: "aws_connect_hours_of_operation",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("hours_of_operation_id",true)`,
		}
	})
	p.AddResourceConfigurator("aws_connect_quick_connect", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_routing_profile", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["default_outbound_queue_id"] = config.Reference{
			TerraformName: "aws_connect_queue",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("queue_id",true)`,
		}

	})
	p.AddResourceConfigurator("aws_connect_security_profile", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_user_hierarchy_structure", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("aws_connect_vocabulary", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_connect_instance",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
