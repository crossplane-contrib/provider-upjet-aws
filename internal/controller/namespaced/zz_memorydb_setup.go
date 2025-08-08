// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	acl "github.com/upbound/provider-aws/internal/controller/namespaced/memorydb/acl"
	cluster "github.com/upbound/provider-aws/internal/controller/namespaced/memorydb/cluster"
	parametergroup "github.com/upbound/provider-aws/internal/controller/namespaced/memorydb/parametergroup"
	snapshot "github.com/upbound/provider-aws/internal/controller/namespaced/memorydb/snapshot"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/namespaced/memorydb/subnetgroup"
	user "github.com/upbound/provider-aws/internal/controller/namespaced/memorydb/user"
)

// Setup_memorydb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_memorydb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		cluster.Setup,
		parametergroup.Setup,
		snapshot.Setup,
		subnetgroup.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_memorydb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_memorydb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.SetupGated,
		cluster.SetupGated,
		parametergroup.SetupGated,
		snapshot.SetupGated,
		subnetgroup.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
