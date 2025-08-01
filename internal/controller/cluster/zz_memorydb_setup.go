// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	acl "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/acl"
	cluster "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/cluster"
	parametergroup "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/parametergroup"
	snapshot "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/snapshot"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/subnetgroup"
	user "github.com/upbound/provider-aws/internal/controller/cluster/memorydb/user"
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
