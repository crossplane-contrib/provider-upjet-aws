// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	compositealarm "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatch/compositealarm"
	dashboard "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatch/dashboard"
	metricalarm "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatch/metricalarm"
	metricstream "github.com/upbound/provider-aws/internal/controller/namespaced/cloudwatch/metricstream"
)

// Setup_cloudwatch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudwatch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		compositealarm.Setup,
		dashboard.Setup,
		metricalarm.Setup,
		metricstream.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
