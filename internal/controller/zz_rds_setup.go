// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/cluster"
	clusteractivitystream "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/clusteractivitystream"
	clusterendpoint "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/clusterendpoint"
	clusterinstance "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/clusterinstance"
	clusterparametergroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshot "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/dbsnapshotcopy"
	eventsubscription "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/eventsubscription"
	globalcluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/globalcluster"
	instance "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/instance"
	instanceroleassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/optiongroup"
	parametergroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/parametergroup"
	proxy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/proxytarget"
	snapshot "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/snapshot"
	subnetgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/rds/subnetgroup"
)

// Setup_rds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_rds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusteractivitystream.Setup,
		clusterendpoint.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clusterroleassociation.Setup,
		clustersnapshot.Setup,
		dbinstanceautomatedbackupsreplication.Setup,
		dbsnapshotcopy.Setup,
		eventsubscription.Setup,
		globalcluster.Setup,
		instance.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergroup.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		snapshot.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
