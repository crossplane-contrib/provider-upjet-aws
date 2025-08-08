// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	apidestination "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/apidestination"
	archive "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/archive"
	bus "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/bus"
	buspolicy "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/buspolicy"
	connection "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/connection"
	permission "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/permission"
	rule "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/rule"
	target "github.com/upbound/provider-aws/internal/controller/cluster/cloudwatchevents/target"
)

// Setup_cloudwatchevents creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudwatchevents(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apidestination.Setup,
		archive.Setup,
		bus.Setup,
		buspolicy.Setup,
		connection.Setup,
		permission.Setup,
		rule.Setup,
		target.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudwatchevents creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudwatchevents(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apidestination.SetupGated,
		archive.SetupGated,
		bus.SetupGated,
		buspolicy.SetupGated,
		connection.SetupGated,
		permission.SetupGated,
		rule.SetupGated,
		target.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
