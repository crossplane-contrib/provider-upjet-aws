// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	activereceiptruleset "github.com/upbound/provider-aws/internal/controller/namespaced/ses/activereceiptruleset"
	configurationset "github.com/upbound/provider-aws/internal/controller/namespaced/ses/configurationset"
	domaindkim "github.com/upbound/provider-aws/internal/controller/namespaced/ses/domaindkim"
	domainidentity "github.com/upbound/provider-aws/internal/controller/namespaced/ses/domainidentity"
	domainmailfrom "github.com/upbound/provider-aws/internal/controller/namespaced/ses/domainmailfrom"
	emailidentity "github.com/upbound/provider-aws/internal/controller/namespaced/ses/emailidentity"
	eventdestination "github.com/upbound/provider-aws/internal/controller/namespaced/ses/eventdestination"
	identitynotificationtopic "github.com/upbound/provider-aws/internal/controller/namespaced/ses/identitynotificationtopic"
	identitypolicy "github.com/upbound/provider-aws/internal/controller/namespaced/ses/identitypolicy"
	receiptfilter "github.com/upbound/provider-aws/internal/controller/namespaced/ses/receiptfilter"
	receiptrule "github.com/upbound/provider-aws/internal/controller/namespaced/ses/receiptrule"
	receiptruleset "github.com/upbound/provider-aws/internal/controller/namespaced/ses/receiptruleset"
	template "github.com/upbound/provider-aws/internal/controller/namespaced/ses/template"
)

// Setup_ses creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ses(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activereceiptruleset.Setup,
		configurationset.Setup,
		domaindkim.Setup,
		domainidentity.Setup,
		domainmailfrom.Setup,
		emailidentity.Setup,
		eventdestination.Setup,
		identitynotificationtopic.Setup,
		identitypolicy.Setup,
		receiptfilter.Setup,
		receiptrule.Setup,
		receiptruleset.Setup,
		template.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ses creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ses(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activereceiptruleset.SetupGated,
		configurationset.SetupGated,
		domaindkim.SetupGated,
		domainidentity.SetupGated,
		domainmailfrom.SetupGated,
		emailidentity.SetupGated,
		eventdestination.SetupGated,
		identitynotificationtopic.SetupGated,
		identitypolicy.SetupGated,
		receiptfilter.SetupGated,
		receiptrule.SetupGated,
		receiptruleset.SetupGated,
		template.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
