// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/clusterendpoint"
	clusterinstance "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshot "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/clustersnapshot"
	eventsubscription "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/eventsubscription"
	globalcluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/globalcluster"
	parametergroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/parametergroup"
	subnetgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/neptune/subnetgroup"
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
