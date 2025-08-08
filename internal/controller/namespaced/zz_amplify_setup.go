// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	app "github.com/upbound/provider-aws/internal/controller/namespaced/amplify/app"
	backendenvironment "github.com/upbound/provider-aws/internal/controller/namespaced/amplify/backendenvironment"
	branch "github.com/upbound/provider-aws/internal/controller/namespaced/amplify/branch"
	webhook "github.com/upbound/provider-aws/internal/controller/namespaced/amplify/webhook"
)

// Setup_amplify creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_amplify(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
		backendenvironment.Setup,
		branch.Setup,
		webhook.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_amplify creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_amplify(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.SetupGated,
		backendenvironment.SetupGated,
		branch.SetupGated,
		webhook.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
