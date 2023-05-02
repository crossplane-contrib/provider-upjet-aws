/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	budget "github.com/upbound/provider-aws/internal/controller/budgets/budget"
	budgetaction "github.com/upbound/provider-aws/internal/controller/budgets/budgetaction"
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
