// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appcookiestickinesspolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/appcookiestickinesspolicy"
	attachment "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/attachment"
	backendserverpolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/backendserverpolicy"
	elb "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/elb"
	lbcookiestickinesspolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/lbcookiestickinesspolicy"
	lbsslnegotiationpolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/lbsslnegotiationpolicy"
	listenerpolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/listenerpolicy"
	policy "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/policy"
	proxyprotocolpolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/elb/proxyprotocolpolicy"
)

// Setup_elb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_elb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appcookiestickinesspolicy.Setup,
		attachment.Setup,
		backendserverpolicy.Setup,
		elb.Setup,
		lbcookiestickinesspolicy.Setup,
		lbsslnegotiationpolicy.Setup,
		listenerpolicy.Setup,
		policy.Setup,
		proxyprotocolpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_elb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_elb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appcookiestickinesspolicy.SetupGated,
		attachment.SetupGated,
		backendserverpolicy.SetupGated,
		elb.SetupGated,
		lbcookiestickinesspolicy.SetupGated,
		lbsslnegotiationpolicy.SetupGated,
		listenerpolicy.SetupGated,
		policy.SetupGated,
		proxyprotocolpolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_elb registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_elb(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		appcookiestickinesspolicy.SetupWebhookWithManager,
		attachment.SetupWebhookWithManager,
		backendserverpolicy.SetupWebhookWithManager,
		elb.SetupWebhookWithManager,
		lbcookiestickinesspolicy.SetupWebhookWithManager,
		lbsslnegotiationpolicy.SetupWebhookWithManager,
		listenerpolicy.SetupWebhookWithManager,
		policy.SetupWebhookWithManager,
		proxyprotocolpolicy.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
