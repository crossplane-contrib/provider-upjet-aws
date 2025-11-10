// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	domain "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cloudsearch/domain"
	domainserviceaccesspolicy "github.com/upbound/provider-aws/v2/internal/controller/namespaced/cloudsearch/domainserviceaccesspolicy"
)

// Setup_cloudsearch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudsearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domainserviceaccesspolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudsearch creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudsearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.SetupGated,
		domainserviceaccesspolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
