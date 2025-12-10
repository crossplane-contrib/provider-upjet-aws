// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	api "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/api"
	apimapping "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/apimapping"
	authorizer "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/authorizer"
	deployment "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/deployment"
	domainname "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/domainname"
	integration "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/integration"
	integrationresponse "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/integrationresponse"
	model "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/model"
	route "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/route"
	routeresponse "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/routeresponse"
	stage "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/stage"
	vpclink "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigatewayv2/vpclink"
)

// Setup_apigatewayv2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apigatewayv2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.Setup,
		apimapping.Setup,
		authorizer.Setup,
		deployment.Setup,
		domainname.Setup,
		integration.Setup,
		integrationresponse.Setup,
		model.Setup,
		route.Setup,
		routeresponse.Setup,
		stage.Setup,
		vpclink.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_apigatewayv2 creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apigatewayv2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.SetupGated,
		apimapping.SetupGated,
		authorizer.SetupGated,
		deployment.SetupGated,
		domainname.SetupGated,
		integration.SetupGated,
		integrationresponse.SetupGated,
		model.SetupGated,
		route.SetupGated,
		routeresponse.SetupGated,
		stage.SetupGated,
		vpclink.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
