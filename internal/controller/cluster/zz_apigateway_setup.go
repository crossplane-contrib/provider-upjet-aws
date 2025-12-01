// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/account"
	apikey "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/apikey"
	authorizer "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/authorizer"
	basepathmapping "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/basepathmapping"
	clientcertificate "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/clientcertificate"
	deployment "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/deployment"
	documentationpart "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/documentationpart"
	documentationversion "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/documentationversion"
	domainname "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/domainname"
	gatewayresponse "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/gatewayresponse"
	integration "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/integration"
	integrationresponse "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/integrationresponse"
	method "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/method"
	methodresponse "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/methodresponse"
	methodsettings "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/methodsettings"
	model "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/model"
	requestvalidator "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/requestvalidator"
	resource "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/resource"
	restapi "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/restapi"
	restapipolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/restapipolicy"
	stage "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/stage"
	usageplan "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/usageplan"
	usageplankey "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/usageplankey"
	vpclink "github.com/upbound/provider-aws/v2/internal/controller/cluster/apigateway/vpclink"
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
