// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	lb "github.com/upbound/provider-aws/v2/internal/controller/cluster/elbv2/lb"
	lblistener "github.com/upbound/provider-aws/v2/internal/controller/cluster/elbv2/lblistener"
	lblistenercertificate "github.com/upbound/provider-aws/v2/internal/controller/cluster/elbv2/lblistenercertificate"
	lblistenerrule "github.com/upbound/provider-aws/v2/internal/controller/cluster/elbv2/lblistenerrule"
	lbtargetgroup "github.com/upbound/provider-aws/v2/internal/controller/cluster/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/upbound/provider-aws/v2/internal/controller/cluster/elbv2/lbtargetgroupattachment"
	lbtruststore "github.com/upbound/provider-aws/v2/internal/controller/cluster/elbv2/lbtruststore"
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
		lbtruststore.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_elbv2 creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_elbv2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		lb.SetupGated,
		lblistener.SetupGated,
		lblistenercertificate.SetupGated,
		lblistenerrule.SetupGated,
		lbtargetgroup.SetupGated,
		lbtargetgroupattachment.SetupGated,
		lbtruststore.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_elbv2 registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_elbv2(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		lb.SetupWebhookWithManager,
		lblistener.SetupWebhookWithManager,
		lblistenercertificate.SetupWebhookWithManager,
		lblistenerrule.SetupWebhookWithManager,
		lbtargetgroup.SetupWebhookWithManager,
		lbtargetgroupattachment.SetupWebhookWithManager,
		lbtruststore.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
