/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accesspoint "github.com/upbound/provider-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/upbound/provider-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/upbound/provider-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/upbound/provider-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/upbound/provider-aws/internal/controller/efs/mounttarget"
	replicationconfiguration "github.com/upbound/provider-aws/internal/controller/efs/replicationconfiguration"
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
