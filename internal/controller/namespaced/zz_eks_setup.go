// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accessentry "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/accessentry"
	accesspolicyassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/accesspolicyassociation"
	addon "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/addon"
	cluster "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/cluster"
	clusterauth "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/clusterauth"
	fargateprofile "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/identityproviderconfig"
	nodegroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/nodegroup"
	podidentityassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/eks/podidentityassociation"
)

// Setup_eks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessentry.Setup,
		accesspolicyassociation.Setup,
		addon.Setup,
		cluster.Setup,
		clusterauth.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		podidentityassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_eks creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_eks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessentry.SetupGated,
		accesspolicyassociation.SetupGated,
		addon.SetupGated,
		cluster.SetupGated,
		clusterauth.SetupGated,
		fargateprofile.SetupGated,
		identityproviderconfig.SetupGated,
		nodegroup.SetupGated,
		podidentityassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
