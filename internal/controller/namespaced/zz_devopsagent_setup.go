// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	agentspace "github.com/upbound/provider-aws/v2/internal/controller/namespaced/devopsagent/agentspace"
	asset "github.com/upbound/provider-aws/v2/internal/controller/namespaced/devopsagent/asset"
)

// Setup_devopsagent creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_devopsagent(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentspace.Setup,
		asset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_devopsagent creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_devopsagent(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentspace.SetupGated,
		asset.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_devopsagent registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_devopsagent(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		agentspace.SetupWebhookWithManager,
		asset.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
