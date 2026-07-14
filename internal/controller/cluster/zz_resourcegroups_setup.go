// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	group "github.com/upbound/provider-aws/v2/internal/controller/cluster/resourcegroups/group"
)

// Setup_resourcegroups creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_resourcegroups(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_resourcegroups creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_resourcegroups(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_resourcegroups registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_resourcegroups(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		group.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
