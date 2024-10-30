// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alternatecontact "github.com/upbound/provider-aws/internal/controller/account/alternatecontact"
	region "github.com/upbound/provider-aws/internal/controller/account/region"
)

// Setup_account creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_account(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alternatecontact.Setup,
		region.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
