// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/fsx/backup"
	datarepositoryassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/fsx/datarepositoryassociation"
	lustrefilesystem "github.com/upbound/provider-aws/v2/internal/controller/namespaced/fsx/lustrefilesystem"
	ontapfilesystem "github.com/upbound/provider-aws/v2/internal/controller/namespaced/fsx/ontapfilesystem"
	ontapstoragevirtualmachine "github.com/upbound/provider-aws/v2/internal/controller/namespaced/fsx/ontapstoragevirtualmachine"
	windowsfilesystem "github.com/upbound/provider-aws/v2/internal/controller/namespaced/fsx/windowsfilesystem"
)

// Setup_fsx creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_fsx(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.Setup,
		datarepositoryassociation.Setup,
		lustrefilesystem.Setup,
		ontapfilesystem.Setup,
		ontapstoragevirtualmachine.Setup,
		windowsfilesystem.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_fsx creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_fsx(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.SetupGated,
		datarepositoryassociation.SetupGated,
		lustrefilesystem.SetupGated,
		ontapfilesystem.SetupGated,
		ontapstoragevirtualmachine.SetupGated,
		windowsfilesystem.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
