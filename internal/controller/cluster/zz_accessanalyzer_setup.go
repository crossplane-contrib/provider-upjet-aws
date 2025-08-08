// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	analyzer "github.com/upbound/provider-aws/internal/controller/cluster/accessanalyzer/analyzer"
	archiverule "github.com/upbound/provider-aws/internal/controller/cluster/accessanalyzer/archiverule"
)

// Setup_accessanalyzer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_accessanalyzer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.Setup,
		archiverule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_accessanalyzer creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_accessanalyzer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.SetupGated,
		archiverule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
