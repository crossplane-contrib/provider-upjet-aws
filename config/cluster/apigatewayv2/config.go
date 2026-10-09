// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apigatewayv2

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the apigatewayv2 group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_apigatewayv2_api_mapping", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
		r.References["domain_name"] = config.Reference{
			TerraformName: "aws_apigatewayv2_domain_name",
		}
		r.References["stage"] = config.Reference{
			TerraformName: "aws_apigatewayv2_stage",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_authorizer", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
		r.References["authorizer_uri"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     "github.com/upbound/provider-aws/v2/config/cluster/common/apis/lambda.FunctionInvokeARN()",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_domain_name", func(r *config.Resource) {
		r.References["domain_name_configuration.certificate_arn"] = config.Reference{
			TerraformName: "aws_acm_certificate",
			Extractor:     common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_deployment", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
		// Triggers is a meta-argument that has comma-separated list of other resources in the same HCL block that tells
		// terraform to re-create the resource if those in the list changed. Upjet workspaces contain only a single
		// resource, so this is irrelevant.
		delete(r.TerraformResource.Schema, "triggers")
		if err := r.MetaResource.Examples[0].SetPathValue("lifecycle", nil); err != nil {
			panic(err)
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_integration", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_integration_response", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
		r.References["integration_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_integration",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_model", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_route", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
		r.References["target"] = config.Reference{
			TerraformName: "aws_apigatewayv2_integration",
			Extractor:     "github.com/upbound/provider-aws/v2/config/cluster/common/apis.IntegrationIDPrefixed()",
		}
		r.References["authorizer_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_authorizer",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_route_response", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
		r.References["route_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_route",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_routing_rule", func(r *config.Resource) {
		// config/schema.json carries no max_items for Plugin Framework
		// resources, so the nested blocks the source limits with
		// listvalidator.SizeAtMost(1) are converted to embedded objects here.
		// action and condition are SizeAtLeast(1) lists and stay lists.
		r.AddSingletonListConversion("action[*].invoke_api", "action[*].invokeApi")
		r.AddSingletonListConversion("condition[*].match_base_paths", "condition[*].matchBasePaths")
		r.AddSingletonListConversion("condition[*].match_headers", "condition[*].matchHeaders")
		r.AddSingletonListConversion("condition[*].match_headers[*].any_of", "condition[*].matchHeaders[*].anyOf")
		r.References["domain_name"] = config.Reference{
			TerraformName: "aws_apigatewayv2_domain_name",
		}
		// invoke_api targets REST APIs only.
		r.References["action.invoke_api.api_id"] = config.Reference{
			TerraformName: "aws_api_gateway_rest_api",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_stage", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_api",
		}
		r.References["deployment_id"] = config.Reference{
			TerraformName: "aws_apigatewayv2_deployment",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_vpclink", func(r *config.Resource) {
		r.UseAsync = true
	})
}
