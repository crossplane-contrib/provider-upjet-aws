// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	directoryconfig "github.com/upbound/provider-aws/internal/controller/namespaced/appstream/directoryconfig"
	fleet "github.com/upbound/provider-aws/internal/controller/namespaced/appstream/fleet"
	fleetstackassociation "github.com/upbound/provider-aws/internal/controller/namespaced/appstream/fleetstackassociation"
	imagebuilder "github.com/upbound/provider-aws/internal/controller/namespaced/appstream/imagebuilder"
	stack "github.com/upbound/provider-aws/internal/controller/namespaced/appstream/stack"
	user "github.com/upbound/provider-aws/internal/controller/namespaced/appstream/user"
	userstackassociation "github.com/upbound/provider-aws/internal/controller/namespaced/appstream/userstackassociation"
)

// Setup_appstream creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appstream(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		directoryconfig.Setup,
		fleet.Setup,
		fleetstackassociation.Setup,
		imagebuilder.Setup,
		stack.Setup,
		user.Setup,
		userstackassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_appstream creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_appstream(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		directoryconfig.SetupGated,
		fleet.SetupGated,
		fleetstackassociation.SetupGated,
		imagebuilder.SetupGated,
		stack.SetupGated,
		user.SetupGated,
		userstackassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
