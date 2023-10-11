// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servicecatalog

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the servicecatalog group.
func Configure(p *config.Provider) {
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
			Type: "Product",
		}
		r.References["tag_option_id"] = config.Reference{
			Type: "TagOption",
		}
	})

	p.AddResourceConfigurator("aws_servicecatalog_product_portfolio_association", func(r *config.Resource) {
		r.References["product_id"] = config.Reference{
			Type: "Product",
		}
		r.References["portfolio_id"] = config.Reference{
			Type: "Portfolio",
		}
	})

	p.AddResourceConfigurator("aws_servicecatalog_principal_portfolio_association", func(r *config.Resource) {
		r.References["portfolio_id"] = config.Reference{
			Type: "Portfolio",
		}
		r.References["principal_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.User",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_servicecatalog_budget_resource_association", func(r *config.Resource) {
		r.References["resource_id"] = config.Reference{
			Type: "Product",
		}
		r.References["budget_name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/budgets/v1beta1.Budget",
		}
	})
}
