// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	activation "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/activation"
	association "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/association"
	defaultpatchbaseline "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/defaultpatchbaseline"
	document "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/document"
	maintenancewindow "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/maintenancewindowtask"
	parameter "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/parameter"
	patchbaseline "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/patchbaseline"
	patchgroup "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/patchgroup"
	resourcedatasync "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/resourcedatasync"
	servicesetting "github.com/upbound/provider-aws/internal/controller/namespaced/ssm/servicesetting"
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
