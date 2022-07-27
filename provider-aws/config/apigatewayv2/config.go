/*
Copyright 2022 Upbound Inc.
*/

package apigatewayv2

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for autoscaling group.
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
			Type:      "github.com/upbound/official-providers/provider-aws/apis/lambda/v1beta1.Function",
			Extractor: "github.com/upbound/official-providers/provider-aws/config/lambda.LambdaFunctionInvokeARN()",
		}
	})
	p.AddResourceConfigurator("aws_apigatewayv2_domain_name", func(r *config.Resource) {
		r.References["domain_name_configuration.certificate_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/acm/v1beta1.Certificate",
			Extractor: common.PathARNExtractor,
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
	})
}
