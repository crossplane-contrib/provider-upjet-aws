// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	app "github.com/upbound/provider-aws/internal/controller/cluster/pinpoint/app"
	smschannel "github.com/upbound/provider-aws/internal/controller/cluster/pinpoint/smschannel"
)

// Setup_pinpoint creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_pinpoint(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
		smschannel.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
