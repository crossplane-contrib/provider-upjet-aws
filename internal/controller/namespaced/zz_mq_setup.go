// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	broker "github.com/upbound/provider-aws/internal/controller/namespaced/mq/broker"
	configuration "github.com/upbound/provider-aws/internal/controller/namespaced/mq/configuration"
	user "github.com/upbound/provider-aws/internal/controller/namespaced/mq/user"
)

// Setup_mq creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mq(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		broker.Setup,
		configuration.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_mq creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_mq(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		broker.SetupGated,
		configuration.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
