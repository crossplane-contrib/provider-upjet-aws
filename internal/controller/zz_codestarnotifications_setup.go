// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	notificationrule "github.com/upbound/provider-aws/internal/controller/codestarnotifications/notificationrule"
)

// Setup_codestarnotifications creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_codestarnotifications(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		notificationrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
