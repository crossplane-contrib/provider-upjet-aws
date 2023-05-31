/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	principalassociation "github.com/upbound/provider-aws/internal/controller/ram/principalassociation"
	resourceassociation "github.com/upbound/provider-aws/internal/controller/ram/resourceassociation"
	resourceshare "github.com/upbound/provider-aws/internal/controller/ram/resourceshare"
	resourceshareaccepter "github.com/upbound/provider-aws/internal/controller/ram/resourceshareaccepter"
)

// Setup_ram creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ram(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		principalassociation.Setup,
		resourceassociation.Setup,
		resourceshare.Setup,
		resourceshareaccepter.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
