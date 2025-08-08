// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	firewall "github.com/upbound/provider-aws/internal/controller/cluster/networkfirewall/firewall"
	firewallpolicy "github.com/upbound/provider-aws/internal/controller/cluster/networkfirewall/firewallpolicy"
	loggingconfiguration "github.com/upbound/provider-aws/internal/controller/cluster/networkfirewall/loggingconfiguration"
	rulegroup "github.com/upbound/provider-aws/internal/controller/cluster/networkfirewall/rulegroup"
)

// Setup_networkfirewall creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkfirewall(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewall.Setup,
		firewallpolicy.Setup,
		loggingconfiguration.Setup,
		rulegroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networkfirewall creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networkfirewall(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewall.SetupGated,
		firewallpolicy.SetupGated,
		loggingconfiguration.SetupGated,
		rulegroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
