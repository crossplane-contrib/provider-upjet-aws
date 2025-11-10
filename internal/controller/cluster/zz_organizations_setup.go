// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-aws/v2/internal/controller/cluster/organizations/account"
	delegatedadministrator "github.com/upbound/provider-aws/v2/internal/controller/cluster/organizations/delegatedadministrator"
	organization "github.com/upbound/provider-aws/v2/internal/controller/cluster/organizations/organization"
	organizationalunit "github.com/upbound/provider-aws/v2/internal/controller/cluster/organizations/organizationalunit"
	policy "github.com/upbound/provider-aws/v2/internal/controller/cluster/organizations/policy"
	policyattachment "github.com/upbound/provider-aws/v2/internal/controller/cluster/organizations/policyattachment"
)

// Setup_organizations creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_organizations(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		delegatedadministrator.Setup,
		organization.Setup,
		organizationalunit.Setup,
		policy.Setup,
		policyattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_organizations creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_organizations(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		delegatedadministrator.SetupGated,
		organization.SetupGated,
		organizationalunit.SetupGated,
		policy.SetupGated,
		policyattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
