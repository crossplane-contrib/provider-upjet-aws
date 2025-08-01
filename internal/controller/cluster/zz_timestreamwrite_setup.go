// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	database "github.com/upbound/provider-aws/internal/controller/cluster/timestreamwrite/database"
	table "github.com/upbound/provider-aws/internal/controller/cluster/timestreamwrite/table"
)

// Setup_timestreamwrite creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_timestreamwrite(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		table.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_timestreamwrite creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_timestreamwrite(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.SetupGated,
		table.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
