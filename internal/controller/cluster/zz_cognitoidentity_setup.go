// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cognitoidentitypoolproviderprincipaltag "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidentity/cognitoidentitypoolproviderprincipaltag"
	pool "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidentity/pool"
	poolrolesattachment "github.com/upbound/provider-aws/internal/controller/cluster/cognitoidentity/poolrolesattachment"
)

// Setup_cognitoidentity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cognitoidentity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cognitoidentitypoolproviderprincipaltag.Setup,
		pool.Setup,
		poolrolesattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
