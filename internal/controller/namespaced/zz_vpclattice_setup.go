// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesslogsubscription "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/accesslogsubscription"
	authpolicy "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/authpolicy"
	listener "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/listener"
	listenerrule "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/listenerrule"
	resourceconfiguration "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/resourceconfiguration"
	resourcegateway "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/resourcegateway"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/resourcepolicy"
	service "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/service"
	servicenetwork "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/servicenetwork"
	servicenetworkresourceassociation "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/servicenetworkresourceassociation"
	servicenetworkserviceassociation "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/servicenetworkserviceassociation"
	servicenetworkvpcassociation "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/servicenetworkvpcassociation"
	targetgroup "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/targetgroup"
	targetgroupattachment "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/targetgroupattachment"
)

// Setup_vpclattice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_vpclattice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslogsubscription.Setup,
		authpolicy.Setup,
		listener.Setup,
		listenerrule.Setup,
		resourceconfiguration.Setup,
		resourcegateway.Setup,
		resourcepolicy.Setup,
		service.Setup,
		servicenetwork.Setup,
		servicenetworkresourceassociation.Setup,
		servicenetworkserviceassociation.Setup,
		servicenetworkvpcassociation.Setup,
		targetgroup.Setup,
		targetgroupattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_vpclattice creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_vpclattice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslogsubscription.SetupGated,
		authpolicy.SetupGated,
		listener.SetupGated,
		listenerrule.SetupGated,
		resourceconfiguration.SetupGated,
		resourcegateway.SetupGated,
		resourcepolicy.SetupGated,
		service.SetupGated,
		servicenetwork.SetupGated,
		servicenetworkresourceassociation.SetupGated,
		servicenetworkserviceassociation.SetupGated,
		servicenetworkvpcassociation.SetupGated,
		targetgroup.SetupGated,
		targetgroupattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
