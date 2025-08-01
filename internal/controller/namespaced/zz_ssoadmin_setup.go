// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accountassignment "github.com/upbound/provider-aws/internal/controller/namespaced/ssoadmin/accountassignment"
	customermanagedpolicyattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ssoadmin/customermanagedpolicyattachment"
	instanceaccesscontrolattributes "github.com/upbound/provider-aws/internal/controller/namespaced/ssoadmin/instanceaccesscontrolattributes"
	managedpolicyattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ssoadmin/managedpolicyattachment"
	permissionsboundaryattachment "github.com/upbound/provider-aws/internal/controller/namespaced/ssoadmin/permissionsboundaryattachment"
	permissionset "github.com/upbound/provider-aws/internal/controller/namespaced/ssoadmin/permissionset"
	permissionsetinlinepolicy "github.com/upbound/provider-aws/internal/controller/namespaced/ssoadmin/permissionsetinlinepolicy"
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

// SetupGated_ssoadmin creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ssoadmin(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountassignment.SetupGated,
		customermanagedpolicyattachment.SetupGated,
		instanceaccesscontrolattributes.SetupGated,
		managedpolicyattachment.SetupGated,
		permissionsboundaryattachment.SetupGated,
		permissionset.SetupGated,
		permissionsetinlinepolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
