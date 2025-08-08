// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package organization

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for the organization group.
func Configure(p *config.Provider) { //nolint:gocyclo
	// please see: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/organizations_account
	// If `role_name` is used, it's stated that Terraform will always
	// show a difference. Thus, that argument is removed here.
	// TODO: There may be similar other arguments.
	p.AddResourceConfigurator("aws_organizations_account", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "role_name")
	})
	p.AddResourceConfigurator("aws_organizations_delegated_administrator", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			TerraformName: "aws_organizations_account",
		}
	})
	// We are deleting this reference as we have three different types: Organization Account,
	// Organization Root, Organization Unit.
	p.AddResourceConfigurator("aws_organizations_policy_attachment", func(r *config.Resource) {
		delete(r.References, "target_id")
	})
}
