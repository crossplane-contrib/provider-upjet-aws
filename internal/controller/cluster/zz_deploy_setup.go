// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	app "github.com/upbound/provider-aws/v2/internal/controller/cluster/deploy/app"
	deploymentconfig "github.com/upbound/provider-aws/v2/internal/controller/cluster/deploy/deploymentconfig"
	deploymentgroup "github.com/upbound/provider-aws/v2/internal/controller/cluster/deploy/deploymentgroup"
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

// SetupGated_deploy creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_deploy(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.SetupGated,
		deploymentconfig.SetupGated,
		deploymentgroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
