// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	app "github.com/upbound/provider-aws/v2/internal/controller/namespaced/pinpoint/app"
	smschannel "github.com/upbound/provider-aws/v2/internal/controller/namespaced/pinpoint/smschannel"
)

// Setup_pinpoint creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_pinpoint(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
		smschannel.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_pinpoint creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_pinpoint(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.SetupGated,
		smschannel.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
