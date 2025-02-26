// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	association "github.com/upbound/provider-aws/internal/controller/licensemanager/association"
	licenseconfiguration "github.com/upbound/provider-aws/internal/controller/licensemanager/licenseconfiguration"
)

// Setup_licensemanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_licensemanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		association.Setup,
		licenseconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
