// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	feature "github.com/upbound/provider-aws/v2/internal/controller/namespaced/evidently/feature"
	project "github.com/upbound/provider-aws/v2/internal/controller/namespaced/evidently/project"
	segment "github.com/upbound/provider-aws/v2/internal/controller/namespaced/evidently/segment"
)

// Setup_evidently creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_evidently(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		feature.Setup,
		project.Setup,
		segment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_evidently creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_evidently(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		feature.SetupGated,
		project.SetupGated,
		segment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
