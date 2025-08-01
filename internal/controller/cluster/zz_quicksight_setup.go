// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	group "github.com/upbound/provider-aws/internal/controller/cluster/quicksight/group"
	user "github.com/upbound/provider-aws/internal/controller/cluster/quicksight/user"
)

// Setup_quicksight creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_quicksight(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
