// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesspoint "github.com/upbound/provider-aws/internal/controller/cluster/efs/accesspoint"
	backuppolicy "github.com/upbound/provider-aws/internal/controller/cluster/efs/backuppolicy"
	filesystem "github.com/upbound/provider-aws/internal/controller/cluster/efs/filesystem"
	filesystempolicy "github.com/upbound/provider-aws/internal/controller/cluster/efs/filesystempolicy"
	mounttarget "github.com/upbound/provider-aws/internal/controller/cluster/efs/mounttarget"
	replicationconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/efs/replicationconfiguration"
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
