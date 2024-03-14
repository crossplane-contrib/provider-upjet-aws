package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/upbound/provider-aws/internal/controller/apigateway/account"
	apikey "github.com/upbound/provider-aws/internal/controller/apigateway/apikey"
	authorizer "github.com/upbound/provider-aws/internal/controller/apigateway/authorizer"
	basepathmapping "github.com/upbound/provider-aws/internal/controller/apigateway/basepathmapping"
	clientcertificate "github.com/upbound/provider-aws/internal/controller/apigateway/clientcertificate"
	deployment "github.com/upbound/provider-aws/internal/controller/apigateway/deployment"
	documentationpart "github.com/upbound/provider-aws/internal/controller/apigateway/documentationpart"
	documentationversion "github.com/upbound/provider-aws/internal/controller/apigateway/documentationversion"
	domainname "github.com/upbound/provider-aws/internal/controller/apigateway/domainname"
	gatewayresponse "github.com/upbound/provider-aws/internal/controller/apigateway/gatewayresponse"
	integration "github.com/upbound/provider-aws/internal/controller/apigateway/integration"
	integrationresponse "github.com/upbound/provider-aws/internal/controller/apigateway/integrationresponse"
	method "github.com/upbound/provider-aws/internal/controller/apigateway/method"
	methodresponse "github.com/upbound/provider-aws/internal/controller/apigateway/methodresponse"
	methodsettings "github.com/upbound/provider-aws/internal/controller/apigateway/methodsettings"
	model "github.com/upbound/provider-aws/internal/controller/apigateway/model"
	requestvalidator "github.com/upbound/provider-aws/internal/controller/apigateway/requestvalidator"
	resource "github.com/upbound/provider-aws/internal/controller/apigateway/resource"
	restapi "github.com/upbound/provider-aws/internal/controller/apigateway/restapi"
	restapipolicy "github.com/upbound/provider-aws/internal/controller/apigateway/restapipolicy"
	stage "github.com/upbound/provider-aws/internal/controller/apigateway/stage"
	usageplan "github.com/upbound/provider-aws/internal/controller/apigateway/usageplan"
	usageplankey "github.com/upbound/provider-aws/internal/controller/apigateway/usageplankey"
	vpclink "github.com/upbound/provider-aws/internal/controller/apigateway/vpclink"
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
