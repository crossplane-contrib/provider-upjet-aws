// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	listener "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/listener"
	resourceconfiguration "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/resourceconfiguration"
	resourcegateway "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/resourcegateway"
	service "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/service"
	servicenetwork "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/servicenetwork"
	targetgroup "github.com/upbound/provider-aws/internal/controller/namespaced/vpclattice/targetgroup"
)

// Setup_vpclattice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_vpclattice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		listener.Setup,
		resourceconfiguration.Setup,
		resourcegateway.Setup,
		service.Setup,
		servicenetwork.Setup,
		targetgroup.Setup,
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
		listener.SetupGated,
		resourceconfiguration.SetupGated,
		resourcegateway.SetupGated,
		service.SetupGated,
		servicenetwork.SetupGated,
		targetgroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
