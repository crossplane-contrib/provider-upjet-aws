// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	lb "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elbv2/lb"
	lblistener "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elbv2/lblistener"
	lblistenercertificate "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elbv2/lblistenercertificate"
	lblistenerrule "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elbv2/lblistenerrule"
	lbtargetgroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elbv2/lbtargetgroupattachment"
)

// Setup_elbv2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_elbv2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		lb.Setup,
		lblistener.Setup,
		lblistenercertificate.Setup,
		lblistenerrule.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
