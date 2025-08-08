// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-aws/internal/controller/cluster/macie2/account"
	classificationjob "github.com/upbound/provider-aws/internal/controller/cluster/macie2/classificationjob"
	customdataidentifier "github.com/upbound/provider-aws/internal/controller/cluster/macie2/customdataidentifier"
	findingsfilter "github.com/upbound/provider-aws/internal/controller/cluster/macie2/findingsfilter"
	invitationaccepter "github.com/upbound/provider-aws/internal/controller/cluster/macie2/invitationaccepter"
	member "github.com/upbound/provider-aws/internal/controller/cluster/macie2/member"
)

// Setup_macie2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_macie2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		classificationjob.Setup,
		customdataidentifier.Setup,
		findingsfilter.Setup,
		invitationaccepter.Setup,
		member.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_macie2 creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_macie2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		classificationjob.SetupGated,
		customdataidentifier.SetupGated,
		findingsfilter.SetupGated,
		invitationaccepter.SetupGated,
		member.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
