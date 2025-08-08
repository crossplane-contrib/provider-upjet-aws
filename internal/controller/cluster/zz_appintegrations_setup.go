// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	eventintegration "github.com/upbound/provider-aws/internal/controller/cluster/appintegrations/eventintegration"
)

// Setup_appintegrations creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appintegrations(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		eventintegration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_appintegrations creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_appintegrations(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		eventintegration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
