// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	identityprovider "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/identityprovider"
	resourceserver "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/resourceserver"
	riskconfiguration "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/riskconfiguration"
	user "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/user"
	usergroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/usergroup"
	useringroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/useringroup"
	userpool "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/userpool"
	userpoolclient "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/userpoolclient"
	userpooldomain "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/userpooldomain"
	userpooluicustomization "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cognitoidp/userpooluicustomization"
)

// Setup_cognitoidp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cognitoidp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		identityprovider.Setup,
		resourceserver.Setup,
		riskconfiguration.Setup,
		user.Setup,
		usergroup.Setup,
		useringroup.Setup,
		userpool.Setup,
		userpoolclient.Setup,
		userpooldomain.Setup,
		userpooluicustomization.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cognitoidp creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cognitoidp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		identityprovider.SetupGated,
		resourceserver.SetupGated,
		riskconfiguration.SetupGated,
		user.SetupGated,
		usergroup.SetupGated,
		useringroup.SetupGated,
		userpool.SetupGated,
		userpoolclient.SetupGated,
		userpooldomain.SetupGated,
		userpooluicustomization.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
