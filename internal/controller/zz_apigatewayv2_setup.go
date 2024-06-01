// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	api "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/api"
	apimapping "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/apimapping"
	authorizer "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/authorizer"
	deployment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/deployment"
	domainname "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/domainname"
	integration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/integration"
	integrationresponse "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/integrationresponse"
	model "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/model"
	route "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/route"
	routeresponse "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/routeresponse"
	stage "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/stage"
	vpclink "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigatewayv2/vpclink"
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
