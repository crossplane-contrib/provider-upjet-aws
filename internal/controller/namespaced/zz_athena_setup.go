// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	database "github.com/upbound/provider-aws/internal/controller/namespaced/athena/database"
	datacatalog "github.com/upbound/provider-aws/internal/controller/namespaced/athena/datacatalog"
	namedquery "github.com/upbound/provider-aws/internal/controller/namespaced/athena/namedquery"
	workgroup "github.com/upbound/provider-aws/internal/controller/namespaced/athena/workgroup"
)

// Setup_athena creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_athena(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		datacatalog.Setup,
		namedquery.Setup,
		workgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_athena creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_athena(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.SetupGated,
		datacatalog.SetupGated,
		namedquery.SetupGated,
		workgroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
