// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alias "github.com/upbound/provider-aws/internal/controller/gamelift/alias"
	build "github.com/upbound/provider-aws/internal/controller/gamelift/build"
	fleet "github.com/upbound/provider-aws/internal/controller/gamelift/fleet"
	gamesessionqueue "github.com/upbound/provider-aws/internal/controller/gamelift/gamesessionqueue"
	script "github.com/upbound/provider-aws/internal/controller/gamelift/script"
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
