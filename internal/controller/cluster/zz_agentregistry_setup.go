// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	registry "github.com/upbound/provider-aws/v2/internal/controller/cluster/agentregistry/registry"
)

// Setup_agentregistry creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_agentregistry(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		registry.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_agentregistry creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_agentregistry(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		registry.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_agentregistry registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_agentregistry(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		registry.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
