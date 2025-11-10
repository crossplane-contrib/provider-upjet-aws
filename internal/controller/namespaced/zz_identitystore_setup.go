// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	group "github.com/upbound/provider-aws/v2/internal/controller/namespaced/identitystore/group"
	groupmembership "github.com/upbound/provider-aws/v2/internal/controller/namespaced/identitystore/groupmembership"
	user "github.com/upbound/provider-aws/v2/internal/controller/namespaced/identitystore/user"
)

// Setup_identitystore creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_identitystore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		groupmembership.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_identitystore creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_identitystore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.SetupGated,
		groupmembership.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
