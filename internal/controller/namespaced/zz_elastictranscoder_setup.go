// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	pipeline "github.com/upbound/provider-aws/internal/controller/namespaced/elastictranscoder/pipeline"
	preset "github.com/upbound/provider-aws/internal/controller/namespaced/elastictranscoder/preset"
)

// Setup_elastictranscoder creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_elastictranscoder(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		pipeline.Setup,
		preset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_elastictranscoder creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_elastictranscoder(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		pipeline.SetupGated,
		preset.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
