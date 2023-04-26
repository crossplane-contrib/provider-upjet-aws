/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	database "github.com/upbound/provider-aws/internal/controller/timestreamwrite/database"
	table "github.com/upbound/provider-aws/internal/controller/timestreamwrite/table"
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
