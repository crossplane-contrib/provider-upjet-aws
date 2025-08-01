// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	app "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/app"
	appimageconfig "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/appimageconfig"
	coderepository "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/coderepository"
	device "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/device"
	devicefleet "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/devicefleet"
	domain "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/domain"
	endpoint "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/endpoint"
	endpointconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/endpointconfiguration"
	featuregroup "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/featuregroup"
	image "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/image"
	imageversion "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/imageversion"
	mlflowtrackingserver "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/mlflowtrackingserver"
	model "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/model"
	modelpackagegroup "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/modelpackagegroup"
	modelpackagegrouppolicy "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/modelpackagegrouppolicy"
	notebookinstance "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/notebookinstance"
	notebookinstancelifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/notebookinstancelifecycleconfiguration"
	servicecatalogportfoliostatus "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/servicecatalogportfoliostatus"
	space "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/space"
	studiolifecycleconfig "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/studiolifecycleconfig"
	userprofile "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/userprofile"
	workforce "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/workforce"
	workteam "github.com/upbound/provider-aws/internal/controller/cluster/sagemaker/workteam"
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
		endpoint.Setup,
		endpointconfiguration.Setup,
		featuregroup.Setup,
		image.Setup,
		imageversion.Setup,
		mlflowtrackingserver.Setup,
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

// SetupGated_sagemaker creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_sagemaker(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.SetupGated,
		appimageconfig.SetupGated,
		coderepository.SetupGated,
		device.SetupGated,
		devicefleet.SetupGated,
		domain.SetupGated,
		endpoint.SetupGated,
		endpointconfiguration.SetupGated,
		featuregroup.SetupGated,
		image.SetupGated,
		imageversion.SetupGated,
		mlflowtrackingserver.SetupGated,
		model.SetupGated,
		modelpackagegroup.SetupGated,
		modelpackagegrouppolicy.SetupGated,
		notebookinstance.SetupGated,
		notebookinstancelifecycleconfiguration.SetupGated,
		servicecatalogportfoliostatus.SetupGated,
		space.SetupGated,
		studiolifecycleconfig.SetupGated,
		userprofile.SetupGated,
		workforce.SetupGated,
		workteam.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
