// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accessentry "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/accessentry"
	accesspolicyassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/accesspolicyassociation"
	addon "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/addon"
	capability "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/capability"
	cluster "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/cluster"
	clusterauth "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/clusterauth"
	fargateprofile "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/identityproviderconfig"
	nodegroup "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/nodegroup"
	podidentityassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/eks/podidentityassociation"
)

// Setup_eks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessentry.Setup,
		accesspolicyassociation.Setup,
		addon.Setup,
		capability.Setup,
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
		capability.SetupGated,
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

// SetupWebhookWithManager_eks registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_eks(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		accessentry.SetupWebhookWithManager,
		accesspolicyassociation.SetupWebhookWithManager,
		addon.SetupWebhookWithManager,
		capability.SetupWebhookWithManager,
		cluster.SetupWebhookWithManager,
		clusterauth.SetupWebhookWithManager,
		fargateprofile.SetupWebhookWithManager,
		identityproviderconfig.SetupWebhookWithManager,
		nodegroup.SetupWebhookWithManager,
		podidentityassociation.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
