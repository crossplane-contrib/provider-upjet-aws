// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	activation "github.com/upbound/provider-aws/internal/controller/cluster/ssm/activation"
	association "github.com/upbound/provider-aws/internal/controller/cluster/ssm/association"
	defaultpatchbaseline "github.com/upbound/provider-aws/internal/controller/cluster/ssm/defaultpatchbaseline"
	document "github.com/upbound/provider-aws/internal/controller/cluster/ssm/document"
	maintenancewindow "github.com/upbound/provider-aws/internal/controller/cluster/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/upbound/provider-aws/internal/controller/cluster/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/upbound/provider-aws/internal/controller/cluster/ssm/maintenancewindowtask"
	parameter "github.com/upbound/provider-aws/internal/controller/cluster/ssm/parameter"
	patchbaseline "github.com/upbound/provider-aws/internal/controller/cluster/ssm/patchbaseline"
	patchgroup "github.com/upbound/provider-aws/internal/controller/cluster/ssm/patchgroup"
	resourcedatasync "github.com/upbound/provider-aws/internal/controller/cluster/ssm/resourcedatasync"
	servicesetting "github.com/upbound/provider-aws/internal/controller/cluster/ssm/servicesetting"
)

// Setup_ssm creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ssm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activation.Setup,
		association.Setup,
		defaultpatchbaseline.Setup,
		document.Setup,
		maintenancewindow.Setup,
		maintenancewindowtarget.Setup,
		maintenancewindowtask.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		patchgroup.Setup,
		resourcedatasync.Setup,
		servicesetting.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
