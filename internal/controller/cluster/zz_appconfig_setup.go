// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/upbound/provider-aws/v2/internal/controller/cluster/appconfig/application"
	configurationprofile "github.com/upbound/provider-aws/v2/internal/controller/cluster/appconfig/configurationprofile"
	deployment "github.com/upbound/provider-aws/v2/internal/controller/cluster/appconfig/deployment"
	deploymentstrategy "github.com/upbound/provider-aws/v2/internal/controller/cluster/appconfig/deploymentstrategy"
	environment "github.com/upbound/provider-aws/v2/internal/controller/cluster/appconfig/environment"
	extension "github.com/upbound/provider-aws/v2/internal/controller/cluster/appconfig/extension"
	extensionassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/appconfig/extensionassociation"
	hostedconfigurationversion "github.com/upbound/provider-aws/v2/internal/controller/cluster/appconfig/hostedconfigurationversion"
)

// Setup_appconfig creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appconfig(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		configurationprofile.Setup,
		deployment.Setup,
		deploymentstrategy.Setup,
		environment.Setup,
		extension.Setup,
		extensionassociation.Setup,
		hostedconfigurationversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_appconfig creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_appconfig(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		configurationprofile.SetupGated,
		deployment.SetupGated,
		deploymentstrategy.SetupGated,
		environment.SetupGated,
		extension.SetupGated,
		extensionassociation.SetupGated,
		hostedconfigurationversion.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
