// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package autoscaling

import (
	"strconv"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"

	"github.com/upbound/provider-aws/apis/cluster/autoscaling/v1beta1"
	"github.com/upbound/provider-aws/apis/cluster/autoscaling/v1beta2"
	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the autoscaling group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_autoscaling_group", func(r *config.Resource) {
		// These are mutually exclusive with aws_autoscaling_attachment.
		config.MoveToStatus(r.TerraformResource, "load_balancers", "target_group_arns")

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"availability_zones",
			},
		}

		r.References["vpc_zone_identifier"] = config.Reference{
			TerraformName: "aws_subnet",
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
			TerraformName: "aws_autoscaling_group",
		}
		r.References["alb_target_group_arn"] = config.Reference{
			TerraformName: "aws_lb_target_group",
			Extractor:     common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_autoscaling_group_tag", func(r *config.Resource) {
		r.References["autoscaling_group_name"] = config.Reference{
			TerraformName: "aws_autoscaling_group",
		}
	})
}

func autoScalingGroupConverterFromv1beta1Tov1beta2(src, target xpresource.Managed) error { //nolint:gocyclo
	srcTyped := src.(*v1beta1.AutoscalingGroup)
	targetTyped := target.(*v1beta2.AutoscalingGroup)

	if srcTyped.Spec.ForProvider.Tags != nil {
		var tl []v1beta2.TagParameters
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
			tl = append(tl, tp)
		}
		targetTyped.Spec.ForProvider.Tag = tl
	}

	if srcTyped.Spec.InitProvider.Tags != nil {
		var tl []v1beta2.TagInitParameters
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
			tl = append(tl, tp)
		}
		targetTyped.Spec.InitProvider.Tag = tl
	}

	if srcTyped.Status.AtProvider.Tags != nil {
		var tl []v1beta2.TagObservation
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
			tl = append(tl, tp)
		}
		targetTyped.Status.AtProvider.Tag = tl
	}

	return nil
}

func autoScalingGroupConverterFromv1beta2Tov1beta1(src, target xpresource.Managed) error { //nolint:gocyclo
	srcTyped := src.(*v1beta2.AutoscalingGroup)
	targetTyped := target.(*v1beta1.AutoscalingGroup)

	if srcTyped.Spec.ForProvider.Tag != nil {
		var tl []map[string]*string
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
			tl = append(tl, m)
		}
		targetTyped.Spec.ForProvider.Tags = tl
	}

	if srcTyped.Spec.InitProvider.Tag != nil {
		var tl []map[string]*string
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
			tl = append(tl, m)
		}
		targetTyped.Spec.InitProvider.Tags = tl
	}

	if srcTyped.Status.AtProvider.Tag != nil {
		var tl []map[string]*string
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
			tl = append(tl, m)
		}
		targetTyped.Status.AtProvider.Tags = tl
	}

	return nil
}
