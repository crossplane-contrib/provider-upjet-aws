// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	apicache "github.com/upbound/provider-aws/internal/controller/namespaced/appsync/apicache"
	apikey "github.com/upbound/provider-aws/internal/controller/namespaced/appsync/apikey"
	datasource "github.com/upbound/provider-aws/internal/controller/namespaced/appsync/datasource"
	function "github.com/upbound/provider-aws/internal/controller/namespaced/appsync/function"
	graphqlapi "github.com/upbound/provider-aws/internal/controller/namespaced/appsync/graphqlapi"
	resolver "github.com/upbound/provider-aws/internal/controller/namespaced/appsync/resolver"
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

// SetupGated_appsync creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_appsync(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apicache.SetupGated,
		apikey.SetupGated,
		datasource.SetupGated,
		function.SetupGated,
		graphqlapi.SetupGated,
		resolver.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
