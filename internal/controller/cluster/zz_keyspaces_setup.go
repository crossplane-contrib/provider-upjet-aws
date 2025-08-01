// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	keyspace "github.com/upbound/provider-aws/internal/controller/cluster/keyspaces/keyspace"
	table "github.com/upbound/provider-aws/internal/controller/cluster/keyspaces/table"
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

// SetupGated_keyspaces creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_keyspaces(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		keyspace.SetupGated,
		table.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
