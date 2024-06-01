// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addon "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/eks/addon"
	cluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/eks/cluster"
	clusterauth "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/eks/clusterauth"
	fargateprofile "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/eks/nodegroup"
	podidentityassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/eks/podidentityassociation"
)

// Setup_eks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
