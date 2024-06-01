// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appconfig/application"
	configurationprofile "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appconfig/configurationprofile"
	deployment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appconfig/deployment"
	deploymentstrategy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appconfig/deploymentstrategy"
	environment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appconfig/environment"
	extension "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appconfig/extension"
	extensionassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appconfig/extensionassociation"
	hostedconfigurationversion "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appconfig/hostedconfigurationversion"
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
