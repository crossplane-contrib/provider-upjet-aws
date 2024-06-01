// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	appcookiestickinesspolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/appcookiestickinesspolicy"
	attachment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/attachment"
	backendserverpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/backendserverpolicy"
	elb "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/elb"
	lbcookiestickinesspolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/lbcookiestickinesspolicy"
	lbsslnegotiationpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/lbsslnegotiationpolicy"
	listenerpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/listenerpolicy"
	policy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/policy"
	proxyprotocolpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/elb/proxyprotocolpolicy"
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
