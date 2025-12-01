// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	profile "github.com/upbound/provider-aws/v2/internal/controller/namespaced/rolesanywhere/profile"
)

// Setup_rolesanywhere creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_rolesanywhere(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		profile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_rolesanywhere creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_rolesanywhere(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		profile.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
