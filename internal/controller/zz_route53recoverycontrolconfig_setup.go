// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/route53recoverycontrolconfig/cluster"
	controlpanel "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/route53recoverycontrolconfig/controlpanel"
	routingcontrol "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/route53recoverycontrolconfig/routingcontrol"
	safetyrule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/route53recoverycontrolconfig/safetyrule"
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
