// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	datalakesettings "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/lakeformation/datalakesettings"
	permissions "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/lakeformation/permissions"
	resource "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/lakeformation/resource"
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
