// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package servicecatalog

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the servicecatalog group.
func Configure(p *config.Provider) { //nolint:gocyclo
	// done for proper normalization as the defaulting is done by Terraform
	// please refer to this resource's external-name configuration for details.
	p.AddResourceConfigurator("aws_servicecatalog_principal_portfolio_association", func(r *config.Resource) {
		r.TerraformResource.Schema["accept_language"].Optional = false
		r.TerraformResource.Schema["accept_language"].Required = true
	})

	// done for proper normalization as the defaulting is done by Terraform
	// please refer to this resource's external-name configuration for details.
	p.AddResourceConfigurator("aws_servicecatalog_product_portfolio_association", func(r *config.Resource) {
		r.TerraformResource.Schema["accept_language"].Optional = false
		r.TerraformResource.Schema["accept_language"].Required = true
	})

	p.AddResourceConfigurator("aws_servicecatalog_tag_option_resource_association", func(r *config.Resource) {
		r.References["resource_id"] = config.Reference{
			TerraformName: "aws_servicecatalog_product",
		}
		r.References["tag_option_id"] = config.Reference{
			TerraformName: "aws_servicecatalog_tag_option",
		}
	})

	p.AddResourceConfigurator("aws_servicecatalog_product_portfolio_association", func(r *config.Resource) {
		r.References["product_id"] = config.Reference{
			TerraformName: "aws_servicecatalog_product",
		}
		r.References["portfolio_id"] = config.Reference{
			TerraformName: "aws_servicecatalog_portfolio",
		}
	})

	p.AddResourceConfigurator("aws_servicecatalog_principal_portfolio_association", func(r *config.Resource) {
		r.References["portfolio_id"] = config.Reference{
			TerraformName: "aws_servicecatalog_portfolio",
		}
		r.References["principal_arn"] = config.Reference{
			TerraformName: "aws_iam_user",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_servicecatalog_budget_resource_association", func(r *config.Resource) {
		r.References["resource_id"] = config.Reference{
			TerraformName: "aws_servicecatalog_product",
		}
		r.References["budget_name"] = config.Reference{
			TerraformName: "aws_budgets_budget",
		}
	})
}
