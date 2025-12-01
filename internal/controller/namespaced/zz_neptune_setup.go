// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/cluster"
	clusterendpoint "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/clusterendpoint"
	clusterinstance "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/clusterinstance"
	clusterparametergroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/clusterparametergroup"
	clustersnapshot "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/clustersnapshot"
	eventsubscription "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/eventsubscription"
	globalcluster "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/globalcluster"
	parametergroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/parametergroup"
	subnetgroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/neptune/subnetgroup"
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

// SetupGated_neptune creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_neptune(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.SetupGated,
		clusterendpoint.SetupGated,
		clusterinstance.SetupGated,
		clusterparametergroup.SetupGated,
		clustersnapshot.SetupGated,
		eventsubscription.SetupGated,
		globalcluster.SetupGated,
		parametergroup.SetupGated,
		subnetgroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
