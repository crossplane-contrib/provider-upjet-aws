// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	assessmenttarget "github.com/upbound/provider-aws/v2/internal/controller/namespaced/inspector/assessmenttarget"
	assessmenttemplate "github.com/upbound/provider-aws/v2/internal/controller/namespaced/inspector/assessmenttemplate"
	resourcegroup "github.com/upbound/provider-aws/v2/internal/controller/namespaced/inspector/resourcegroup"
)

// Setup_inspector creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_inspector(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		assessmenttarget.Setup,
		assessmenttemplate.Setup,
		resourcegroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_inspector creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_inspector(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		assessmenttarget.SetupGated,
		assessmenttemplate.SetupGated,
		resourcegroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
