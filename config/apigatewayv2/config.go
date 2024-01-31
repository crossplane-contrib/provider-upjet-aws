/*
Copyright 2022 Upbound Inc.
*/

package apigatewayv2

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the apigatewayv2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_apigatewayv2_api_mapping", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			Type: "API",
		}
		r.References["domain_name"] = config.Reference{
			Type: "DomainName",
		}
		r.References["stage"] = config.Reference{
			Type:      "Stage",
			Extractor: common.PathTerraformIDExtractor,
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_authorizer", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			Type: "API",
		}
		r.References["authorizer_uri"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     "github.com/upbound/provider-aws/config/common/apis/lambda.FunctionInvokeARN()",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_domain_name", func(r *config.Resource) {
		r.References["domain_name_configuration.certificate_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/acm/v1beta1.Certificate",
			Extractor: common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_deployment", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			Type: "API",
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
			Type: "API",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_integration_response", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			Type: "API",
		}
		r.References["integration_id"] = config.Reference{
			Type: "Integration",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_model", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			Type: "API",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_route", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			Type: "API",
		}
		r.References["target"] = config.Reference{
			TerraformName: "aws_apigatewayv2_integration",
			Extractor:     "github.com/upbound/provider-aws/config/common/apis.IntegrationIDPrefixed()",
		}
		r.References["authorizer_id"] = config.Reference{
			Type: "Authorizer",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_route_response", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			Type: "API",
		}
		r.References["route_id"] = config.Reference{
			Type: "Route",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_stage", func(r *config.Resource) {
		r.References["api_id"] = config.Reference{
			Type: "API",
		}
		r.References["deployment_id"] = config.Reference{
			Type: "Deployment",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_vpclink", func(r *config.Resource) {
		r.UseAsync = true
	})
}
