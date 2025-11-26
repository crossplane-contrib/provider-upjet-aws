// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	autoscalingconfigurationversion "github.com/upbound/provider-aws/v2/internal/controller/namespaced/apprunner/autoscalingconfigurationversion"
	connection "github.com/upbound/provider-aws/v2/internal/controller/namespaced/apprunner/connection"
	observabilityconfiguration "github.com/upbound/provider-aws/v2/internal/controller/namespaced/apprunner/observabilityconfiguration"
	service "github.com/upbound/provider-aws/v2/internal/controller/namespaced/apprunner/service"
	vpcconnector "github.com/upbound/provider-aws/v2/internal/controller/namespaced/apprunner/vpcconnector"
)

// Setup_apprunner creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apprunner(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfigurationversion.Setup,
		connection.Setup,
		observabilityconfiguration.Setup,
		service.Setup,
		vpcconnector.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_apprunner creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apprunner(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfigurationversion.SetupGated,
		connection.SetupGated,
		observabilityconfiguration.SetupGated,
		service.SetupGated,
		vpcconnector.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
