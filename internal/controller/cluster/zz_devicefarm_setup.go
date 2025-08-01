// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	devicepool "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/devicepool"
	instanceprofile "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/instanceprofile"
	networkprofile "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/networkprofile"
	project "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/project"
	testgridproject "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/testgridproject"
	upload "github.com/upbound/provider-aws/internal/controller/cluster/devicefarm/upload"
)

// Setup_devicefarm creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_devicefarm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		devicepool.Setup,
		instanceprofile.Setup,
		networkprofile.Setup,
		project.Setup,
		testgridproject.Setup,
		upload.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
