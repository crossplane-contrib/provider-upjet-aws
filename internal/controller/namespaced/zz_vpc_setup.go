// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	networkperformancemetricsubscription "github.com/upbound/provider-aws/v2/internal/controller/namespaced/vpc/networkperformancemetricsubscription"
)

// Setup_vpc creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_vpc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		networkperformancemetricsubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_vpc creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_vpc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		networkperformancemetricsubscription.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
