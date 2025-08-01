// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	definition "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatchlogs/definition"
	destination "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatchlogs/destination"
	destinationpolicy "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatchlogs/destinationpolicy"
	group "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatchlogs/group"
	metricfilter "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatchlogs/resourcepolicy"
	stream "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatchlogs/stream"
	subscriptionfilter "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatchlogs/subscriptionfilter"
)

// Setup_cloudwatchlogs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudwatchlogs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudwatchlogs creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudwatchlogs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		definition.SetupGated,
		destination.SetupGated,
		destinationpolicy.SetupGated,
		group.SetupGated,
		metricfilter.SetupGated,
		resourcepolicy.SetupGated,
		stream.SetupGated,
		subscriptionfilter.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
