package events

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for events group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudwatch_event_permission", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/events/v1beta1.Bus",
			RefFieldName:      "EventBusNameRefs",
			SelectorFieldName: "EventBusNameSelector",
		}
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_rule", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/events/v1beta1.Bus",
			RefFieldName:      "EventBusNameRefs",
			SelectorFieldName: "EventBusNameSelector",
		}
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_target", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/events/v1beta1.Bus",
			RefFieldName:      "EventBusNameRefs",
			SelectorFieldName: "EventBusNameSelector",
		}
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_permission", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_rule", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_target", func(r *config.Resource) {
		r.UseAsync = true
	})
}
