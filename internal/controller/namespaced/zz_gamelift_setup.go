// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alias "github.com/upbound/provider-aws/internal/controller/namespaced/gamelift/alias"
	build "github.com/upbound/provider-aws/internal/controller/namespaced/gamelift/build"
	fleet "github.com/upbound/provider-aws/internal/controller/namespaced/gamelift/fleet"
	gamesessionqueue "github.com/upbound/provider-aws/internal/controller/namespaced/gamelift/gamesessionqueue"
	script "github.com/upbound/provider-aws/internal/controller/namespaced/gamelift/script"
)

// Setup_gamelift creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_gamelift(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		build.Setup,
		fleet.Setup,
		gamesessionqueue.Setup,
		script.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_gamelift creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_gamelift(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.SetupGated,
		build.SetupGated,
		fleet.SetupGated,
		gamesessionqueue.SetupGated,
		script.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
