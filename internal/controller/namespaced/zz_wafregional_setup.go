// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bytematchset "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/bytematchset"
	geomatchset "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/geomatchset"
	ipset "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/ipset"
	ratebasedrule "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/ratebasedrule"
	regexmatchset "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/regexmatchset"
	regexpatternset "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/regexpatternset"
	rule "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/rule"
	sizeconstraintset "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/sizeconstraintset"
	sqlinjectionmatchset "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/sqlinjectionmatchset"
	webacl "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/webacl"
	xssmatchset "github.com/upbound/provider-aws/internal/controller/namespaced/wafregional/xssmatchset"
)

// Setup_wafregional creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_wafregional(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bytematchset.Setup,
		geomatchset.Setup,
		ipset.Setup,
		ratebasedrule.Setup,
		regexmatchset.Setup,
		regexpatternset.Setup,
		rule.Setup,
		sizeconstraintset.Setup,
		sqlinjectionmatchset.Setup,
		webacl.Setup,
		xssmatchset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_wafregional creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_wafregional(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bytematchset.SetupGated,
		geomatchset.SetupGated,
		ipset.SetupGated,
		ratebasedrule.SetupGated,
		regexmatchset.SetupGated,
		regexpatternset.SetupGated,
		rule.SetupGated,
		sizeconstraintset.SetupGated,
		sqlinjectionmatchset.SetupGated,
		webacl.SetupGated,
		xssmatchset.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
