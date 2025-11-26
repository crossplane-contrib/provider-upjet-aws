// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	component "github.com/upbound/provider-aws/v2/internal/controller/cluster/imagebuilder/component"
	containerrecipe "github.com/upbound/provider-aws/v2/internal/controller/cluster/imagebuilder/containerrecipe"
	distributionconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/imagebuilder/distributionconfiguration"
	image "github.com/upbound/provider-aws/v2/internal/controller/cluster/imagebuilder/image"
	imagepipeline "github.com/upbound/provider-aws/v2/internal/controller/cluster/imagebuilder/imagepipeline"
	imagerecipe "github.com/upbound/provider-aws/v2/internal/controller/cluster/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/upbound/provider-aws/v2/internal/controller/cluster/imagebuilder/infrastructureconfiguration"
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

// SetupGated_imagebuilder creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_imagebuilder(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		component.SetupGated,
		containerrecipe.SetupGated,
		distributionconfiguration.SetupGated,
		image.SetupGated,
		imagepipeline.SetupGated,
		imagerecipe.SetupGated,
		infrastructureconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
