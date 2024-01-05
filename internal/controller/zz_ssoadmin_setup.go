// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accountassignment "github.com/upbound/provider-aws/internal/controller/ssoadmin/accountassignment"
	customermanagedpolicyattachment "github.com/upbound/provider-aws/internal/controller/ssoadmin/customermanagedpolicyattachment"
	instanceaccesscontrolattributes "github.com/upbound/provider-aws/internal/controller/ssoadmin/instanceaccesscontrolattributes"
	managedpolicyattachment "github.com/upbound/provider-aws/internal/controller/ssoadmin/managedpolicyattachment"
	permissionsboundaryattachment "github.com/upbound/provider-aws/internal/controller/ssoadmin/permissionsboundaryattachment"
	permissionset "github.com/upbound/provider-aws/internal/controller/ssoadmin/permissionset"
	permissionsetinlinepolicy "github.com/upbound/provider-aws/internal/controller/ssoadmin/permissionsetinlinepolicy"
)

// Setup_ssoadmin creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ssoadmin(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountassignment.Setup,
		customermanagedpolicyattachment.Setup,
		instanceaccesscontrolattributes.Setup,
		managedpolicyattachment.Setup,
		permissionsboundaryattachment.Setup,
		permissionset.Setup,
		permissionsetinlinepolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
