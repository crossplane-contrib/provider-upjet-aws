// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	licenseassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/grafana/licenseassociation"
	roleassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/grafana/roleassociation"
	workspace "github.com/upbound/provider-aws/v2/internal/controller/namespaced/grafana/workspace"
	workspaceapikey "github.com/upbound/provider-aws/v2/internal/controller/namespaced/grafana/workspaceapikey"
	workspacesamlconfiguration "github.com/upbound/provider-aws/v2/internal/controller/namespaced/grafana/workspacesamlconfiguration"
)

// Setup_grafana creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_grafana(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		licenseassociation.Setup,
		roleassociation.Setup,
		workspace.Setup,
		workspaceapikey.Setup,
		workspacesamlconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_grafana creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_grafana(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		licenseassociation.SetupGated,
		roleassociation.SetupGated,
		workspace.SetupGated,
		workspaceapikey.SetupGated,
		workspacesamlconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
