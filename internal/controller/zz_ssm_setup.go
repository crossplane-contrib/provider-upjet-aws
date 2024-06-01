// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	activation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/activation"
	association "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/association"
	defaultpatchbaseline "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/defaultpatchbaseline"
	document "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/document"
	maintenancewindow "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/maintenancewindowtask"
	parameter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/parameter"
	patchbaseline "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/patchbaseline"
	patchgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/patchgroup"
	resourcedatasync "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/resourcedatasync"
	servicesetting "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/ssm/servicesetting"
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
