// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/docdb/cluster"
	clusterinstance "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/docdb/clusterinstance"
	clusterparametergroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/docdb/clusterparametergroup"
	clustersnapshot "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/docdb/clustersnapshot"
	eventsubscription "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/docdb/eventsubscription"
	globalcluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/docdb/globalcluster"
	subnetgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/docdb/subnetgroup"
)

// Setup_docdb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_docdb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		eventsubscription.Setup,
		globalcluster.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
