// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accessentry "github.com/upbound/provider-aws/internal/controller/cluster/eks/accessentry"
	accesspolicyassociation "github.com/upbound/provider-aws/internal/controller/cluster/eks/accesspolicyassociation"
	addon "github.com/upbound/provider-aws/internal/controller/cluster/eks/addon"
	cluster "github.com/upbound/provider-aws/internal/controller/cluster/eks/cluster"
	fargateprofile "github.com/upbound/provider-aws/internal/controller/cluster/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/provider-aws/internal/controller/cluster/eks/identityproviderconfig"
	nodegroup "github.com/upbound/provider-aws/internal/controller/cluster/eks/nodegroup"
	podidentityassociation "github.com/upbound/provider-aws/internal/controller/cluster/eks/podidentityassociation"
	clusterauth "github.com/upbound/provider-aws/internal/controller/cluster/internal/controller/eks/clusterauth"
)

// Setup_eks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessentry.Setup,
		accesspolicyassociation.Setup,
		addon.Setup,
		cluster.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		podidentityassociation.Setup,
		clusterauth.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
