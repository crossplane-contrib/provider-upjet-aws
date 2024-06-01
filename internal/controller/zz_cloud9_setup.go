// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	environmentec2 "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloud9/environmentec2"
	environmentmembership "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/cloud9/environmentmembership"
)

// Setup_cloud9 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloud9(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environmentec2.Setup,
		environmentmembership.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
