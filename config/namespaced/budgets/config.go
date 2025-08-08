// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package budgets

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the budgets group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_budgets_budget_action", func(r *config.Resource) {
		r.References["definition.iam_action_definition.aws_iam_role.example.name"] = config.Reference{
			TerraformName: "aws_iam_role",
		}
	})
}
