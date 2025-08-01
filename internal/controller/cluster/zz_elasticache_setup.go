// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/cluster"
	globalreplicationgroup "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/globalreplicationgroup"
	parametergroup "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/parametergroup"
	replicationgroup "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/replicationgroup"
	serverlesscache "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/serverlesscache"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/subnetgroup"
	user "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/user"
	usergroup "github.com/upbound/provider-aws/internal/controller/cluster/elasticache/usergroup"
)

// Setup_elasticache creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_elasticache(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		globalreplicationgroup.Setup,
		parametergroup.Setup,
		replicationgroup.Setup,
		serverlesscache.Setup,
		subnetgroup.Setup,
		user.Setup,
		usergroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
