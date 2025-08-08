// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	budget "github.com/upbound/provider-aws/internal/controller/cluster/budgets/budget"
	budgetaction "github.com/upbound/provider-aws/internal/controller/cluster/budgets/budgetaction"
)

// Setup_budgets creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_budgets(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		budget.Setup,
		budgetaction.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_budgets creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_budgets(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		budget.SetupGated,
		budgetaction.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
