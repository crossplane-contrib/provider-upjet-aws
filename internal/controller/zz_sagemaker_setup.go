/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	app "github.com/upbound/provider-aws/internal/controller/sagemaker/app"
	appimageconfig "github.com/upbound/provider-aws/internal/controller/sagemaker/appimageconfig"
	coderepository "github.com/upbound/provider-aws/internal/controller/sagemaker/coderepository"
	device "github.com/upbound/provider-aws/internal/controller/sagemaker/device"
	devicefleet "github.com/upbound/provider-aws/internal/controller/sagemaker/devicefleet"
	domain "github.com/upbound/provider-aws/internal/controller/sagemaker/domain"
	endpointconfiguration "github.com/upbound/provider-aws/internal/controller/sagemaker/endpointconfiguration"
	featuregroup "github.com/upbound/provider-aws/internal/controller/sagemaker/featuregroup"
	image "github.com/upbound/provider-aws/internal/controller/sagemaker/image"
	imageversion "github.com/upbound/provider-aws/internal/controller/sagemaker/imageversion"
	model "github.com/upbound/provider-aws/internal/controller/sagemaker/model"
	modelpackagegroup "github.com/upbound/provider-aws/internal/controller/sagemaker/modelpackagegroup"
	modelpackagegrouppolicy "github.com/upbound/provider-aws/internal/controller/sagemaker/modelpackagegrouppolicy"
	notebookinstance "github.com/upbound/provider-aws/internal/controller/sagemaker/notebookinstance"
	notebookinstancelifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/sagemaker/notebookinstancelifecycleconfiguration"
	servicecatalogportfoliostatus "github.com/upbound/provider-aws/internal/controller/sagemaker/servicecatalogportfoliostatus"
	space "github.com/upbound/provider-aws/internal/controller/sagemaker/space"
	studiolifecycleconfig "github.com/upbound/provider-aws/internal/controller/sagemaker/studiolifecycleconfig"
	userprofile "github.com/upbound/provider-aws/internal/controller/sagemaker/userprofile"
	workforce "github.com/upbound/provider-aws/internal/controller/sagemaker/workforce"
	workteam "github.com/upbound/provider-aws/internal/controller/sagemaker/workteam"
)

// Setup_sagemaker creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sagemaker(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
		appimageconfig.Setup,
		coderepository.Setup,
		device.Setup,
		devicefleet.Setup,
		domain.Setup,
		endpointconfiguration.Setup,
		featuregroup.Setup,
		image.Setup,
		imageversion.Setup,
		model.Setup,
		modelpackagegroup.Setup,
		modelpackagegrouppolicy.Setup,
		notebookinstance.Setup,
		notebookinstancelifecycleconfiguration.Setup,
		servicecatalogportfoliostatus.Setup,
		space.Setup,
		studiolifecycleconfig.Setup,
		userprofile.Setup,
		workforce.Setup,
		workteam.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
