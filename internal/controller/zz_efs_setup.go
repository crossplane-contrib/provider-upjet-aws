// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesspoint "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/efs/mounttarget"
	replicationconfiguration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/efs/replicationconfiguration"
)

// Setup_efs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_efs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspoint.Setup,
		backuppolicy.Setup,
		filesystem.Setup,
		filesystempolicy.Setup,
		mounttarget.Setup,
		replicationconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
