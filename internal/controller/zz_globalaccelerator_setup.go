/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accelerator "github.com/upbound/provider-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/upbound/provider-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/upbound/provider-aws/internal/controller/globalaccelerator/listener"
)

// Setup_globalaccelerator creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_globalaccelerator(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
