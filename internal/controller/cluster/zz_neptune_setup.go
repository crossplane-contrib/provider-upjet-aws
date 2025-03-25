// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/upbound/provider-aws/internal/controller/cluster/neptune/cluster"
	clusterendpoint "github.com/upbound/provider-aws/internal/controller/cluster/neptune/clusterendpoint"
	clusterinstance "github.com/upbound/provider-aws/internal/controller/cluster/neptune/clusterinstance"
	clusterparametergroup "github.com/upbound/provider-aws/internal/controller/cluster/neptune/clusterparametergroup"
	clustersnapshot "github.com/upbound/provider-aws/internal/controller/cluster/neptune/clustersnapshot"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/cluster/neptune/eventsubscription"
	globalcluster "github.com/upbound/provider-aws/internal/controller/cluster/neptune/globalcluster"
	parametergroup "github.com/upbound/provider-aws/internal/controller/cluster/neptune/parametergroup"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/cluster/neptune/subnetgroup"
)

// Setup_neptune creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_neptune(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusterendpoint.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		eventsubscription.Setup,
		globalcluster.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
