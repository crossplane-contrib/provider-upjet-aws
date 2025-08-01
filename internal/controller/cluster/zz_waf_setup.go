// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bytematchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/bytematchset"
	geomatchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/geomatchset"
	ipset "github.com/upbound/provider-aws/internal/controller/cluster/waf/ipset"
	ratebasedrule "github.com/upbound/provider-aws/internal/controller/cluster/waf/ratebasedrule"
	regexmatchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/regexmatchset"
	regexpatternset "github.com/upbound/provider-aws/internal/controller/cluster/waf/regexpatternset"
	rule "github.com/upbound/provider-aws/internal/controller/cluster/waf/rule"
	sizeconstraintset "github.com/upbound/provider-aws/internal/controller/cluster/waf/sizeconstraintset"
	sqlinjectionmatchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/sqlinjectionmatchset"
	webacl "github.com/upbound/provider-aws/internal/controller/cluster/waf/webacl"
	xssmatchset "github.com/upbound/provider-aws/internal/controller/cluster/waf/xssmatchset"
)

// Setup_waf creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_waf(mgr ctrl.Manager, o controller.Options) error {
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
