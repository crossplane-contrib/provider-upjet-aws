// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connector "github.com/upbound/provider-aws/v2/internal/controller/cluster/kafkaconnect/connector"
	customplugin "github.com/upbound/provider-aws/v2/internal/controller/cluster/kafkaconnect/customplugin"
	workerconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/kafkaconnect/workerconfiguration"
)

// Setup_kafkaconnect creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kafkaconnect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connector.Setup,
		customplugin.Setup,
		workerconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_kafkaconnect creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_kafkaconnect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connector.SetupGated,
		customplugin.SetupGated,
		workerconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
