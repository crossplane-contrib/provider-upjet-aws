// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connection "github.com/upbound/provider-aws/v2/internal/controller/cluster/codestarconnections/connection"
	host "github.com/upbound/provider-aws/v2/internal/controller/cluster/codestarconnections/host"
)

// Setup_codestarconnections creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_codestarconnections(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connection.Setup,
		host.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_codestarconnections creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_codestarconnections(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connection.SetupGated,
		host.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
