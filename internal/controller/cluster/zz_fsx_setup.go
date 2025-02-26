// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	backup "github.com/upbound/provider-aws/internal/controller/fsx/backup"
	datarepositoryassociation "github.com/upbound/provider-aws/internal/controller/fsx/datarepositoryassociation"
	lustrefilesystem "github.com/upbound/provider-aws/internal/controller/fsx/lustrefilesystem"
	ontapfilesystem "github.com/upbound/provider-aws/internal/controller/fsx/ontapfilesystem"
	ontapstoragevirtualmachine "github.com/upbound/provider-aws/internal/controller/fsx/ontapstoragevirtualmachine"
	windowsfilesystem "github.com/upbound/provider-aws/internal/controller/fsx/windowsfilesystem"
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
