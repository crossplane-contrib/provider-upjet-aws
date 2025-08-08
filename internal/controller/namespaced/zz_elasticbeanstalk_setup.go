// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/upbound/provider-aws/internal/controller/namespaced/elasticbeanstalk/application"
	applicationversion "github.com/upbound/provider-aws/internal/controller/namespaced/elasticbeanstalk/applicationversion"
	configurationtemplate "github.com/upbound/provider-aws/internal/controller/namespaced/elasticbeanstalk/configurationtemplate"
)

// Setup_elasticbeanstalk creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_elasticbeanstalk(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationversion.Setup,
		configurationtemplate.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_elasticbeanstalk creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_elasticbeanstalk(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		applicationversion.SetupGated,
		configurationtemplate.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
