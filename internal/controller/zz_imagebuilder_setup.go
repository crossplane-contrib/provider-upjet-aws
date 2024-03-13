package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	component "github.com/upbound/provider-aws/internal/controller/imagebuilder/component"
	containerrecipe "github.com/upbound/provider-aws/internal/controller/imagebuilder/containerrecipe"
	distributionconfiguration "github.com/upbound/provider-aws/internal/controller/imagebuilder/distributionconfiguration"
	image "github.com/upbound/provider-aws/internal/controller/imagebuilder/image"
	imagepipeline "github.com/upbound/provider-aws/internal/controller/imagebuilder/imagepipeline"
	imagerecipe "github.com/upbound/provider-aws/internal/controller/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/upbound/provider-aws/internal/controller/imagebuilder/infrastructureconfiguration"
)

// Setup_imagebuilder creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_imagebuilder(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		component.Setup,
		containerrecipe.Setup,
		distributionconfiguration.Setup,
		image.Setup,
		imagepipeline.Setup,
		imagerecipe.Setup,
		infrastructureconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
