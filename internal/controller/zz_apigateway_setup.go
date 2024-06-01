// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/account"
	apikey "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/apikey"
	authorizer "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/authorizer"
	basepathmapping "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/basepathmapping"
	clientcertificate "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/clientcertificate"
	deployment "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/deployment"
	documentationpart "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/documentationpart"
	documentationversion "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/documentationversion"
	domainname "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/domainname"
	gatewayresponse "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/gatewayresponse"
	integration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/integration"
	integrationresponse "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/integrationresponse"
	method "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/method"
	methodresponse "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/methodresponse"
	methodsettings "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/methodsettings"
	model "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/model"
	requestvalidator "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/requestvalidator"
	resource "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/resource"
	restapi "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/restapi"
	restapipolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/restapipolicy"
	stage "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/stage"
	usageplan "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/usageplan"
	usageplankey "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/usageplankey"
	vpclink "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/apigateway/vpclink"
)

// Setup_apigateway creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apigateway(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		apikey.Setup,
		authorizer.Setup,
		basepathmapping.Setup,
		clientcertificate.Setup,
		deployment.Setup,
		documentationpart.Setup,
		documentationversion.Setup,
		domainname.Setup,
		gatewayresponse.Setup,
		integration.Setup,
		integrationresponse.Setup,
		method.Setup,
		methodresponse.Setup,
		methodsettings.Setup,
		model.Setup,
		requestvalidator.Setup,
		resource.Setup,
		restapi.Setup,
		restapipolicy.Setup,
		stage.Setup,
		usageplan.Setup,
		usageplankey.Setup,
		vpclink.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
