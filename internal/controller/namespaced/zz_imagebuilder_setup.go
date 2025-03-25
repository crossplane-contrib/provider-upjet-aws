// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	component "github.com/upbound/provider-aws/internal/controller/namespaced/imagebuilder/component"
	containerrecipe "github.com/upbound/provider-aws/internal/controller/namespaced/imagebuilder/containerrecipe"
	distributionconfiguration "github.com/upbound/provider-aws/internal/controller/namespaced/imagebuilder/distributionconfiguration"
	image "github.com/upbound/provider-aws/internal/controller/namespaced/imagebuilder/image"
	imagepipeline "github.com/upbound/provider-aws/internal/controller/namespaced/imagebuilder/imagepipeline"
	imagerecipe "github.com/upbound/provider-aws/internal/controller/namespaced/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/upbound/provider-aws/internal/controller/namespaced/imagebuilder/infrastructureconfiguration"
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
