// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	connector "github.com/upbound/provider-aws/internal/controller/kafkaconnect/connector"
	customplugin "github.com/upbound/provider-aws/internal/controller/kafkaconnect/customplugin"
	workerconfiguration "github.com/upbound/provider-aws/internal/controller/kafkaconnect/workerconfiguration"
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
