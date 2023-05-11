/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/cluster"
	controlpanel "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/controlpanel"
	routingcontrol "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/routingcontrol"
	safetyrule "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/safetyrule"
)

// Setup_route53recoverycontrolconfig creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53recoverycontrolconfig(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		controlpanel.Setup,
		routingcontrol.Setup,
		safetyrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
