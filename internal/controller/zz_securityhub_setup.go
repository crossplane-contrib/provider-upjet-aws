// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/securityhub/account"
	actiontarget "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/securityhub/actiontarget"
	findingaggregator "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/securityhub/findingaggregator"
	insight "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/securityhub/insight"
	inviteaccepter "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/securityhub/inviteaccepter"
	member "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/securityhub/member"
	productsubscription "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/securityhub/productsubscription"
	standardssubscription "github.com/crossplane-contrib/provider-upjet-aws/internal/controller/securityhub/standardssubscription"
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
