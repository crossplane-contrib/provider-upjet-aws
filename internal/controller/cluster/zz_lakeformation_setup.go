// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	datalakesettings "github.com/upbound/provider-aws/internal/controller/cluster/lakeformation/datalakesettings"
	permissions "github.com/upbound/provider-aws/internal/controller/cluster/lakeformation/permissions"
	resource "github.com/upbound/provider-aws/internal/controller/cluster/lakeformation/resource"
)

// Setup_lakeformation creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lakeformation(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		datalakesettings.Setup,
		permissions.Setup,
		resource.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
