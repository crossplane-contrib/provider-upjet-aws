// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	domain "github.com/upbound/provider-aws/internal/controller/namespaced/codeartifact/domain"
	domainpermissionspolicy "github.com/upbound/provider-aws/internal/controller/namespaced/codeartifact/domainpermissionspolicy"
	repository "github.com/upbound/provider-aws/internal/controller/namespaced/codeartifact/repository"
	repositorypermissionspolicy "github.com/upbound/provider-aws/internal/controller/namespaced/codeartifact/repositorypermissionspolicy"
)

// Setup_codeartifact creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_codeartifact(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domainpermissionspolicy.Setup,
		repository.Setup,
		repositorypermissionspolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_codeartifact creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_codeartifact(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.SetupGated,
		domainpermissionspolicy.SetupGated,
		repository.SetupGated,
		repositorypermissionspolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
