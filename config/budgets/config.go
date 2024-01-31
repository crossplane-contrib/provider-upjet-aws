/*
Copyright 2021 Upbound Inc.
*/

package budgets

import (
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"

	"github.com/upbound/provider-aws/apis/budgets/v1beta1"
	"github.com/upbound/provider-aws/apis/budgets/v1beta2"
)

// Configure adds configurations for the budgets group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_budgets_budget_action", func(r *config.Resource) {
		r.References["definition.iam_action_definition.aws_iam_role.example.name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
		}
	})
	p.AddResourceConfigurator("aws_budgets_budget", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.Conversions = append(r.Conversions,
			conversion.NewCustomConverter("v1beta1", "v1beta2", budgetConverterFromv1beta1Tov1beta2),
			conversion.NewCustomConverter("v1beta2", "v1beta1", budgetConverterFromv1beta2Tov1beta1),
		)
	})
}

func budgetConverterFromv1beta1Tov1beta2(src, target resource.Managed) error { //nolint:gocyclo
	srcTyped := src.(*v1beta1.Budget)
	targetTyped := target.(*v1beta2.Budget)
	for k, v := range srcTyped.Spec.ForProvider.CostFilters {
		key := k
		cfp := v1beta2.CostFilterParameters{
			Name:   &key,
			Values: []*string{v},
		}
		targetTyped.Spec.ForProvider.CostFilter = append(targetTyped.Spec.ForProvider.CostFilter, cfp)
	}
	for k, v := range srcTyped.Spec.InitProvider.CostFilters {
		key := k
		cfp := v1beta2.CostFilterInitParameters{
			Name:   &key,
			Values: []*string{v},
		}
		targetTyped.Spec.InitProvider.CostFilter = append(targetTyped.Spec.InitProvider.CostFilter, cfp)
	}
	for k, v := range srcTyped.Status.AtProvider.CostFilters {
		key := k
		cfp := v1beta2.CostFilterObservation{
			Name:   &key,
			Values: []*string{v},
		}
		targetTyped.Status.AtProvider.CostFilter = append(targetTyped.Status.AtProvider.CostFilter, cfp)
	}
	return nil
}

func budgetConverterFromv1beta2Tov1beta1(src, target resource.Managed) error { //nolint:gocyclo
	srcTyped := src.(*v1beta2.Budget)
	targetTyped := target.(*v1beta1.Budget)
	for _, e := range srcTyped.Spec.ForProvider.CostFilter {
		if targetTyped.Spec.ForProvider.CostFilters == nil {
			targetTyped.Spec.ForProvider.CostFilters = map[string]*string{}
		}
		if e.Name != nil && e.Values != nil {
			targetTyped.Spec.ForProvider.CostFilters[*e.Name] = e.Values[0]
		}
	}
	for _, e := range srcTyped.Spec.InitProvider.CostFilter {
		if targetTyped.Spec.InitProvider.CostFilters == nil {
			targetTyped.Spec.InitProvider.CostFilters = map[string]*string{}
		}
		if e.Name != nil && e.Values != nil {
			targetTyped.Spec.InitProvider.CostFilters[*e.Name] = e.Values[0]
		}
	}
	for _, e := range srcTyped.Status.AtProvider.CostFilter {
		if targetTyped.Status.AtProvider.CostFilters == nil {
			targetTyped.Status.AtProvider.CostFilters = map[string]*string{}
		}
		if e.Name != nil && e.Values != nil {
			targetTyped.Status.AtProvider.CostFilters[*e.Name] = e.Values[0]
		}
	}
	return nil
}
