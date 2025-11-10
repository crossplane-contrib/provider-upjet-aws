// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-aws/v2/internal/controller/namespaced/securityhub/account"
	actiontarget "github.com/upbound/provider-aws/v2/internal/controller/namespaced/securityhub/actiontarget"
	findingaggregator "github.com/upbound/provider-aws/v2/internal/controller/namespaced/securityhub/findingaggregator"
	insight "github.com/upbound/provider-aws/v2/internal/controller/namespaced/securityhub/insight"
	inviteaccepter "github.com/upbound/provider-aws/v2/internal/controller/namespaced/securityhub/inviteaccepter"
	member "github.com/upbound/provider-aws/v2/internal/controller/namespaced/securityhub/member"
	productsubscription "github.com/upbound/provider-aws/v2/internal/controller/namespaced/securityhub/productsubscription"
	standardssubscription "github.com/upbound/provider-aws/v2/internal/controller/namespaced/securityhub/standardssubscription"
)

// Setup_securityhub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_securityhub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		actiontarget.Setup,
		findingaggregator.Setup,
		insight.Setup,
		inviteaccepter.Setup,
		member.Setup,
		productsubscription.Setup,
		standardssubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_securityhub creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_securityhub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		actiontarget.SetupGated,
		findingaggregator.SetupGated,
		insight.SetupGated,
		inviteaccepter.SetupGated,
		member.SetupGated,
		productsubscription.SetupGated,
		standardssubscription.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
