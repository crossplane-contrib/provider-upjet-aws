// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	protection "github.com/upbound/provider-aws/v2/internal/controller/namespaced/shield/protection"
	protectionhealthcheckassociation "github.com/upbound/provider-aws/v2/internal/controller/namespaced/shield/protectionhealthcheckassociation"
)

// Setup_shield creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_shield(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		protection.Setup,
		protectionhealthcheckassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_shield creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_shield(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		protection.SetupGated,
		protectionhealthcheckassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_shield registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_shield(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		protection.SetupWebhookWithManager,
		protectionhealthcheckassociation.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
