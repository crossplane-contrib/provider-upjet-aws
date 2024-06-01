// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	app "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/deploy/app"
	deploymentconfig "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/deploy/deploymentconfig"
	deploymentgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/deploy/deploymentgroup"
)

// Setup_deploy creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_deploy(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
		deploymentconfig.Setup,
		deploymentgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
