// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	directoryconfig "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appstream/directoryconfig"
	fleet "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appstream/fleet"
	fleetstackassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appstream/fleetstackassociation"
	imagebuilder "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appstream/imagebuilder"
	stack "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appstream/stack"
	user "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appstream/user"
	userstackassociation "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/appstream/userstackassociation"
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
