// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	ipset "github.com/upbound/provider-aws/internal/controller/wafv2/ipset"
	regexpatternset "github.com/upbound/provider-aws/internal/controller/wafv2/regexpatternset"
	rulegroup "github.com/upbound/provider-aws/internal/controller/wafv2/rulegroup"
	webacl "github.com/upbound/provider-aws/internal/controller/wafv2/webacl"
	webaclassociation "github.com/upbound/provider-aws/internal/controller/wafv2/webaclassociation"
	webaclloggingconfiguration "github.com/upbound/provider-aws/internal/controller/wafv2/webaclloggingconfiguration"
)

// Setup_wafv2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_wafv2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ipset.Setup,
		regexpatternset.Setup,
		rulegroup.Setup,
		webacl.Setup,
		webaclassociation.Setup,
		webaclloggingconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
