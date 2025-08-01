// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/account"
	apikey "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/apikey"
	authorizer "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/authorizer"
	basepathmapping "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/basepathmapping"
	clientcertificate "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/clientcertificate"
	deployment "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/deployment"
	documentationpart "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/documentationpart"
	documentationversion "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/documentationversion"
	domainname "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/domainname"
	gatewayresponse "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/gatewayresponse"
	integration "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/integration"
	integrationresponse "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/integrationresponse"
	method "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/method"
	methodresponse "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/methodresponse"
	methodsettings "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/methodsettings"
	model "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/model"
	requestvalidator "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/requestvalidator"
	resource "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/resource"
	restapi "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/restapi"
	restapipolicy "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/restapipolicy"
	stage "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/stage"
	usageplan "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/usageplan"
	usageplankey "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/usageplankey"
	vpclink "github.com/upbound/provider-aws/internal/controller/namespaced/apigateway/vpclink"
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

// SetupGated_apigateway creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apigateway(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		apikey.SetupGated,
		authorizer.SetupGated,
		basepathmapping.SetupGated,
		clientcertificate.SetupGated,
		deployment.SetupGated,
		documentationpart.SetupGated,
		documentationversion.SetupGated,
		domainname.SetupGated,
		gatewayresponse.SetupGated,
		integration.SetupGated,
		integrationresponse.SetupGated,
		method.SetupGated,
		methodresponse.SetupGated,
		methodsettings.SetupGated,
		model.SetupGated,
		requestvalidator.SetupGated,
		resource.SetupGated,
		restapi.SetupGated,
		restapipolicy.SetupGated,
		stage.SetupGated,
		usageplan.SetupGated,
		usageplankey.SetupGated,
		vpclink.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
