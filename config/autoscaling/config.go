/*
Copyright 2021 Upbound Inc.
*/

package autoscaling

import (
	"strconv"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"

	"github.com/upbound/provider-aws/apis/autoscaling/v1beta1"
	"github.com/upbound/provider-aws/apis/autoscaling/v1beta2"
	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the autoscaling group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_autoscaling_group", func(r *config.Resource) {
		// These are mutually exclusive with aws_autoscaling_attachment.
		config.MoveToStatus(r.TerraformResource, "load_balancers", "target_group_arns")

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"availability_zones",
			},
		}

		r.References["vpc_zone_identifier"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
		delete(r.References, "launch_template.version")
		r.UseAsync = true

		r.Version = "v1beta2"
		r.Conversions = append(r.Conversions,
			conversion.NewCustomConverter("v1beta1", "v1beta2", autoScalingGroupConverterFromv1beta1Tov1beta2),
			conversion.NewCustomConverter("v1beta2", "v1beta1", autoScalingGroupConverterFromv1beta2Tov1beta1),
		)
	})
	p.AddResourceConfigurator("aws_autoscaling_attachment", func(r *config.Resource) {
		r.References["autoscaling_group_name"] = config.Reference{
			Type: "AutoscalingGroup",
		}
		r.References["alb_target_group_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup",
			Extractor: common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_autoscaling_group_tag", func(r *config.Resource) {
		r.References["autoscaling_group_name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/autoscaling/v1beta1.AutoscalingGroup",
		}
		r.OverrideFieldNames = map[string]string{
			"TagParameters":     "GroupTagTagParameters",
			"TagObservation":    "GroupTagTagObservation",
			"TagInitParameters": "GroupTagTagInitParameters",
		}
	})
}

func autoScalingGroupConverterFromv1beta1Tov1beta2(src, target xpresource.Managed) error { //nolint:gocyclo
	srcTyped := src.(*v1beta1.AutoscalingGroup)
	targetTyped := target.(*v1beta2.AutoscalingGroup)
	for _, e := range srcTyped.Spec.ForProvider.Tags {
		tp := v1beta2.TagParameters{
			Key:   e["key"],
			Value: e["value"],
		}
		if e["propagate_at_launch"] != nil {
			propagateAtLaunchStr := e["propagate_at_launch"]
			propagateAtLaunch, err := strconv.ParseBool(*propagateAtLaunchStr)
			cpPropagateAtLaunch := propagateAtLaunch
			if err != nil {
				return err
			}
			tp.PropagateAtLaunch = &cpPropagateAtLaunch
		}
		targetTyped.Spec.ForProvider.Tag = append(targetTyped.Spec.ForProvider.Tag, tp)
	}
	for _, e := range srcTyped.Spec.InitProvider.Tags {
		tp := v1beta2.TagInitParameters{
			Key:   e["key"],
			Value: e["value"],
		}
		if e["propagate_at_launch"] != nil {
			propagateAtLaunchStr := e["propagate_at_launch"]
			propagateAtLaunch, err := strconv.ParseBool(*propagateAtLaunchStr)
			cpPropagateAtLaunch := propagateAtLaunch
			if err != nil {
				return err
			}
			tp.PropagateAtLaunch = &cpPropagateAtLaunch
		}
		targetTyped.Spec.InitProvider.Tag = append(targetTyped.Spec.InitProvider.Tag, tp)
	}
	for _, e := range srcTyped.Status.AtProvider.Tags {
		tp := v1beta2.TagObservation{
			Key:   e["key"],
			Value: e["value"],
		}
		if e["propagate_at_launch"] != nil {
			propagateAtLaunchStr := e["propagate_at_launch"]
			propagateAtLaunch, err := strconv.ParseBool(*propagateAtLaunchStr)
			cpPropagateAtLaunch := propagateAtLaunch
			if err != nil {
				return err
			}
			tp.PropagateAtLaunch = &cpPropagateAtLaunch
		}
		targetTyped.Status.AtProvider.Tag = append(targetTyped.Status.AtProvider.Tag, tp)
	}
	return nil
}

func autoScalingGroupConverterFromv1beta2Tov1beta1(src, target xpresource.Managed) error { //nolint:gocyclo
	srcTyped := src.(*v1beta2.AutoscalingGroup)
	targetTyped := target.(*v1beta1.AutoscalingGroup)
	for _, e := range srcTyped.Spec.ForProvider.Tag {
		m := map[string]*string{}
		if e.Key != nil {
			m["key"] = e.Key
		}
		if e.Value != nil {
			m["value"] = e.Value
		}
		if e.PropagateAtLaunch != nil {
			propagateAtLaunch := strconv.FormatBool(*e.PropagateAtLaunch)
			cpPropagateAtLaunch := propagateAtLaunch
			m["propagate_at_launch"] = &cpPropagateAtLaunch
		}
		targetTyped.Spec.ForProvider.Tags = append(targetTyped.Spec.ForProvider.Tags, m)
	}
	for _, e := range srcTyped.Spec.InitProvider.Tag {
		m := map[string]*string{}
		if e.Key != nil {
			m["key"] = e.Key
		}
		if e.Value != nil {
			m["value"] = e.Value
		}
		if e.PropagateAtLaunch != nil {
			propagateAtLaunch := strconv.FormatBool(*e.PropagateAtLaunch)
			cpPropagateAtLaunch := propagateAtLaunch
			m["propagate_at_launch"] = &cpPropagateAtLaunch
		}
		targetTyped.Spec.InitProvider.Tags = append(targetTyped.Spec.InitProvider.Tags, m)
	}
	for _, e := range srcTyped.Status.AtProvider.Tag {
		m := map[string]*string{}
		if e.Key != nil {
			m["key"] = e.Key
		}
		if e.Value != nil {
			m["value"] = e.Value
		}
		if e.PropagateAtLaunch != nil {
			propagateAtLaunch := strconv.FormatBool(*e.PropagateAtLaunch)
			cpPropagateAtLaunch := propagateAtLaunch
			m["propagate_at_launch"] = &cpPropagateAtLaunch
		}
		targetTyped.Status.AtProvider.Tags = append(targetTyped.Status.AtProvider.Tags, m)
	}
	return nil
}
