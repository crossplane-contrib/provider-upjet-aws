// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	directory "github.com/upbound/provider-aws/v2/internal/controller/namespaced/workspaces/directory"
	ipgroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/workspaces/ipgroup"
)

// Setup_workspaces creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_workspaces(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		directory.Setup,
		ipgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_workspaces creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_workspaces(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		directory.SetupGated,
		ipgroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
