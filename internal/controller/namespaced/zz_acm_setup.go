// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	certificate "github.com/upbound/provider-aws/internal/controller/namespaced/acm/certificate"
	certificatevalidation "github.com/upbound/provider-aws/internal/controller/namespaced/acm/certificatevalidation"
)

// Setup_acm creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_acm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.Setup,
		certificatevalidation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_acm creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_acm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.SetupGated,
		certificatevalidation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
