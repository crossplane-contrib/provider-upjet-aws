// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	httpnamespace "github.com/upbound/provider-aws/v2/internal/controller/namespaced/servicediscovery/httpnamespace"
	privatednsnamespace "github.com/upbound/provider-aws/v2/internal/controller/namespaced/servicediscovery/privatednsnamespace"
	publicdnsnamespace "github.com/upbound/provider-aws/v2/internal/controller/namespaced/servicediscovery/publicdnsnamespace"
	service "github.com/upbound/provider-aws/v2/internal/controller/namespaced/servicediscovery/service"
)

// Setup_servicediscovery creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicediscovery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		httpnamespace.Setup,
		privatednsnamespace.Setup,
		publicdnsnamespace.Setup,
		service.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_servicediscovery creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_servicediscovery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		httpnamespace.SetupGated,
		privatednsnamespace.SetupGated,
		publicdnsnamespace.SetupGated,
		service.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
