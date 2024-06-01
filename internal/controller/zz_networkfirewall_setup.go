// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	firewall "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/networkfirewall/firewall"
	firewallpolicy "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/networkfirewall/firewallpolicy"
	loggingconfiguration "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/networkfirewall/loggingconfiguration"
	rulegroup "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/networkfirewall/rulegroup"
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
