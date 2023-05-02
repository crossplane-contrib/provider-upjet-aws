/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accountassignment "github.com/upbound/provider-aws/internal/controller/ssoadmin/accountassignment"
	managedpolicyattachment "github.com/upbound/provider-aws/internal/controller/ssoadmin/managedpolicyattachment"
	permissionset "github.com/upbound/provider-aws/internal/controller/ssoadmin/permissionset"
	permissionsetinlinepolicy "github.com/upbound/provider-aws/internal/controller/ssoadmin/permissionsetinlinepolicy"
)

// Setup_ssoadmin creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ssoadmin(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountassignment.Setup,
		managedpolicyattachment.Setup,
		permissionset.Setup,
		permissionsetinlinepolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
