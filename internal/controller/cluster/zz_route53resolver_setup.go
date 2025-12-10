// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	endpoint "github.com/upbound/provider-aws/v2/internal/controller/cluster/route53resolver/endpoint"
	rule "github.com/upbound/provider-aws/v2/internal/controller/cluster/route53resolver/rule"
	ruleassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/route53resolver/ruleassociation"
)

// Setup_route53resolver creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53resolver(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.Setup,
		rule.Setup,
		ruleassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_route53resolver creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_route53resolver(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.SetupGated,
		rule.SetupGated,
		ruleassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
