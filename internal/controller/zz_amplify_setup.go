// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	app "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/amplify/app"
	backendenvironment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/amplify/backendenvironment"
	branch "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/amplify/branch"
	webhook "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/amplify/webhook"
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
