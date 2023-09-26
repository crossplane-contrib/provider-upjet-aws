// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/upbound/provider-aws/internal/controller/opsworks/application"
	instance "github.com/upbound/provider-aws/internal/controller/opsworks/instance"
	permission "github.com/upbound/provider-aws/internal/controller/opsworks/permission"
	rdsdbinstance "github.com/upbound/provider-aws/internal/controller/opsworks/rdsdbinstance"
	stack "github.com/upbound/provider-aws/internal/controller/opsworks/stack"
	userprofile "github.com/upbound/provider-aws/internal/controller/opsworks/userprofile"
)

// Setup_opsworks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_opsworks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		instance.Setup,
		permission.Setup,
		rdsdbinstance.Setup,
		stack.Setup,
		userprofile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
