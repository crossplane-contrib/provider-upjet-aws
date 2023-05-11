/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	keyspace "github.com/upbound/provider-aws/internal/controller/keyspaces/keyspace"
	table "github.com/upbound/provider-aws/internal/controller/keyspaces/table"
)

// Setup_keyspaces creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_keyspaces(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		keyspace.Setup,
		table.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
