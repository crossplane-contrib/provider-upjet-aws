/*
Copyright 2022 Upbound Inc.
*/

package connect

import (
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/upbound/provider-aws/apis/connect/v1beta1"
	"github.com/upbound/provider-aws/apis/connect/v1beta2"
)

// Configure adds configurations for the connect group.
func Configure(p *config.Provider) {
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

		r.Version = "v1beta2"
		r.Conversions = append(r.Conversions,
			conversion.NewCustomConverter("v1beta1", "v1beta2", func(src, target xpresource.Managed) error {
				srcTyped := src.(*v1beta1.RoutingProfile)
				targetTyped := target.(*v1beta2.RoutingProfile)
				for _, e := range srcTyped.Status.AtProvider.QueueConfigsAssociated {
					qc := v1beta2.QueueConfigsObservation{
						Channel:   e.Channel,
						Delay:     e.Delay,
						Priority:  e.Priority,
						QueueArn:  e.QueueArn,
						QueueID:   e.QueueID,
						QueueName: e.QueueName,
					}
					targetTyped.Status.AtProvider.QueueConfigs = append(targetTyped.Status.AtProvider.QueueConfigs, qc)
				}
				return nil
			}),
			conversion.NewCustomConverter("v1beta2", "v1beta1", func(src, target xpresource.Managed) error {
				srcTyped := src.(*v1beta2.RoutingProfile)
				targetTyped := target.(*v1beta1.RoutingProfile)
				for _, e := range srcTyped.Status.AtProvider.QueueConfigs {
					qca := v1beta1.QueueConfigsAssociatedObservation{
						Channel:   e.Channel,
						Delay:     e.Delay,
						Priority:  e.Priority,
						QueueArn:  e.QueueArn,
						QueueID:   e.QueueID,
						QueueName: e.QueueName,
					}
					targetTyped.Status.AtProvider.QueueConfigsAssociated = append(targetTyped.Status.AtProvider.QueueConfigsAssociated, qca)
				}
				return nil
			}))
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
