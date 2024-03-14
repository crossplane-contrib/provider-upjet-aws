// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package budgets

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the budgets group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_budgets_budget_action", func(r *config.Resource) {
		r.References["definition.iam_action_definition.aws_iam_role.example.name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
		}
	})
}
