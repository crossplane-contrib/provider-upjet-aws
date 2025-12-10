// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	budgetresourceassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/budgetresourceassociation"
	constraint "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/constraint"
	portfolio "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/portfolio"
	portfolioshare "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/principalportfolioassociation"
	product "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/product"
	productportfolioassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/productportfolioassociation"
	provisioningartifact "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/provisioningartifact"
	serviceaction "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/serviceaction"
	tagoption "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/upbound/provider-aws/v2/internal/controller/cluster/servicecatalog/tagoptionresourceassociation"
)

// Setup_servicecatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicecatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		budgetresourceassociation.Setup,
		constraint.Setup,
		portfolio.Setup,
		portfolioshare.Setup,
		principalportfolioassociation.Setup,
		product.Setup,
		productportfolioassociation.Setup,
		provisioningartifact.Setup,
		serviceaction.Setup,
		tagoption.Setup,
		tagoptionresourceassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_servicecatalog creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_servicecatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		budgetresourceassociation.SetupGated,
		constraint.SetupGated,
		portfolio.SetupGated,
		portfolioshare.SetupGated,
		principalportfolioassociation.SetupGated,
		product.SetupGated,
		productportfolioassociation.SetupGated,
		provisioningartifact.SetupGated,
		serviceaction.SetupGated,
		tagoption.SetupGated,
		tagoptionresourceassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
