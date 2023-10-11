// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apicache "github.com/upbound/provider-aws/internal/controller/appsync/apicache"
	apikey "github.com/upbound/provider-aws/internal/controller/appsync/apikey"
	datasource "github.com/upbound/provider-aws/internal/controller/appsync/datasource"
	function "github.com/upbound/provider-aws/internal/controller/appsync/function"
	graphqlapi "github.com/upbound/provider-aws/internal/controller/appsync/graphqlapi"
	resolver "github.com/upbound/provider-aws/internal/controller/appsync/resolver"
)

// Setup_appsync creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appsync(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apicache.Setup,
		apikey.Setup,
		datasource.Setup,
		function.Setup,
		graphqlapi.Setup,
		resolver.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
