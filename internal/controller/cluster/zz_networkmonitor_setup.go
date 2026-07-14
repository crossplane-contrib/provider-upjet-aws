// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	monitor "github.com/upbound/provider-aws/v2/internal/controller/cluster/networkmonitor/monitor"
	probe "github.com/upbound/provider-aws/v2/internal/controller/cluster/networkmonitor/probe"
)

// Setup_networkmonitor creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkmonitor(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitor.Setup,
		probe.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networkmonitor creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networkmonitor(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitor.SetupGated,
		probe.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_networkmonitor registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_networkmonitor(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		monitor.SetupWebhookWithManager,
		probe.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
