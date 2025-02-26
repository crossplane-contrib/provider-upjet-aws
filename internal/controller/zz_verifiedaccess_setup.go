// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	trustprovider "github.com/upbound/provider-aws/internal/controller/verifiedaccess/trustprovider"
)

// Setup_verifiedaccess creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_verifiedaccess(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		trustprovider.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
