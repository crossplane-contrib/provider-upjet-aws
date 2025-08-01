// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	identityprovider "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/identityprovider"
	resourceserver "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/resourceserver"
	riskconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/riskconfiguration"
	user "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/user"
	usergroup "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/usergroup"
	useringroup "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/useringroup"
	userpool "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/userpool"
	userpoolclient "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/userpoolclient"
	userpooldomain "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/userpooldomain"
	userpooluicustomization "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidp/userpooluicustomization"
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
