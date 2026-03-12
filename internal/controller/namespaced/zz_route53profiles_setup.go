// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	association "github.com/upbound/provider-aws/v2/internal/controller/namespaced/route53profiles/association"
	profile "github.com/upbound/provider-aws/v2/internal/controller/namespaced/route53profiles/profile"
	resourceassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/route53profiles/resourceassociation"
)

// Setup_route53profiles creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53profiles(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		association.Setup,
		profile.Setup,
		resourceassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_route53profiles creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_route53profiles(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		association.SetupGated,
		profile.SetupGated,
		resourceassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
