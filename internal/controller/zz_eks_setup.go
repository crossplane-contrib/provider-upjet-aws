/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	addon "github.com/upbound/provider-aws/internal/controller/eks/addon"
	cluster "github.com/upbound/provider-aws/internal/controller/eks/cluster"
	clusterauth "github.com/upbound/provider-aws/internal/controller/eks/clusterauth"
	nodegroup "github.com/upbound/provider-aws/internal/controller/eks/nodegroup"
)

// Setup_eks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addon.Setup,
		cluster.Setup,
		clusterauth.Setup,
		nodegroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
