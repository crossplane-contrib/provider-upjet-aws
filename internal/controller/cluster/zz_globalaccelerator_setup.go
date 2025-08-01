// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accelerator "github.com/upbound/provider-aws/internal/controller/cluster/globalaccelerator/accelerator"
	endpointgroup "github.com/upbound/provider-aws/internal/controller/cluster/globalaccelerator/endpointgroup"
	listener "github.com/upbound/provider-aws/internal/controller/cluster/globalaccelerator/listener"
)

// Setup_globalaccelerator creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_globalaccelerator(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_globalaccelerator creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_globalaccelerator(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accelerator.SetupGated,
		endpointgroup.SetupGated,
		listener.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
