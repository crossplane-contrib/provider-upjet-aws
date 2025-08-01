// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	delegationset "github.com/upbound/provider-aws/internal/controller/cluster/route53/delegationset"
	healthcheck "github.com/upbound/provider-aws/internal/controller/cluster/route53/healthcheck"
	hostedzonednssec "github.com/upbound/provider-aws/internal/controller/cluster/route53/hostedzonednssec"
	querylog "github.com/upbound/provider-aws/internal/controller/cluster/route53/querylog"
	record "github.com/upbound/provider-aws/internal/controller/cluster/route53/record"
	resolverconfig "github.com/upbound/provider-aws/internal/controller/cluster/route53/resolverconfig"
	trafficpolicy "github.com/upbound/provider-aws/internal/controller/cluster/route53/trafficpolicy"
	trafficpolicyinstance "github.com/upbound/provider-aws/internal/controller/cluster/route53/trafficpolicyinstance"
	vpcassociationauthorization "github.com/upbound/provider-aws/internal/controller/cluster/route53/vpcassociationauthorization"
	zone "github.com/upbound/provider-aws/internal/controller/cluster/route53/zone"
	zoneassociation "github.com/upbound/provider-aws/internal/controller/cluster/route53/zoneassociation"
)

// Setup_route53 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		delegationset.Setup,
		healthcheck.Setup,
		hostedzonednssec.Setup,
		querylog.Setup,
		record.Setup,
		resolverconfig.Setup,
		trafficpolicy.Setup,
		trafficpolicyinstance.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
		zoneassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_route53 creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_route53(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		delegationset.SetupGated,
		healthcheck.SetupGated,
		hostedzonednssec.SetupGated,
		querylog.SetupGated,
		record.SetupGated,
		resolverconfig.SetupGated,
		trafficpolicy.SetupGated,
		trafficpolicyinstance.SetupGated,
		vpcassociationauthorization.SetupGated,
		zone.SetupGated,
		zoneassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
