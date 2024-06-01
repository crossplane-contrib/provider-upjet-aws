// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apidestination "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloudwatchevents/apidestination"
	archive "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloudwatchevents/archive"
	bus "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloudwatchevents/bus"
	buspolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloudwatchevents/buspolicy"
	connection "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloudwatchevents/connection"
	permission "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloudwatchevents/permission"
	rule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloudwatchevents/rule"
	target "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloudwatchevents/target"
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
