/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	appcookiestickinesspolicy "github.com/upbound/provider-aws/internal/controller/elb/appcookiestickinesspolicy"
	attachment "github.com/upbound/provider-aws/internal/controller/elb/attachment"
	backendserverpolicy "github.com/upbound/provider-aws/internal/controller/elb/backendserverpolicy"
	elb "github.com/upbound/provider-aws/internal/controller/elb/elb"
	lbcookiestickinesspolicy "github.com/upbound/provider-aws/internal/controller/elb/lbcookiestickinesspolicy"
	lbsslnegotiationpolicy "github.com/upbound/provider-aws/internal/controller/elb/lbsslnegotiationpolicy"
	listenerpolicy "github.com/upbound/provider-aws/internal/controller/elb/listenerpolicy"
	policy "github.com/upbound/provider-aws/internal/controller/elb/policy"
	proxyprotocolpolicy "github.com/upbound/provider-aws/internal/controller/elb/proxyprotocolpolicy"
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
