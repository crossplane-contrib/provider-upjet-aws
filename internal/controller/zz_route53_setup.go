/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	delegationset "github.com/upbound/provider-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/upbound/provider-aws/internal/controller/route53/healthcheck"
	record "github.com/upbound/provider-aws/internal/controller/route53/record"
	zone "github.com/upbound/provider-aws/internal/controller/route53/zone"
)

// Setup_route53 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		delegationset.Setup,
		healthcheck.Setup,
		record.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
