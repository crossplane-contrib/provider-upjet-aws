// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package lambda

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the lambda group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lambda_alias", func(r *config.Resource) {
		r.References["function_name"] = config.Reference{
			Type: "Function",
		}
	})

	p.AddResourceConfigurator("aws_lambda_code_signing_config", func(r *config.Resource) {
		r.References["allowed_publishers.signing_profile_version_arns"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/signer/v1beta1.SigningProfile",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_lambda_event_source_mapping", func(r *config.Resource) {
		r.References["function_name"] = config.Reference{
			Type:      "Function",
			Extractor: common.PathARNExtractor,
		}
		delete(r.References, "event_source_arn")
		// It can be fulfilled by multiple types.
		delete(r.References, "source_access_configuration.uri")
		r.UseAsync = true
	})

	// TODO: Automated test pipeline cannot be run for the lambda group resources.
	// Because many resources of this group need `lambda_function` resource and it
	// has a `filename` field for creation. This field reads a file from local
	// storage and uses this file during provisioning.
	// We may consider adding metadata configuration for the `lambda_function` in
	// a future PR.
	p.AddResourceConfigurator("aws_lambda_function", func(r *config.Resource) {
		r.References["s3_bucket"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["role"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["vpc_config.security_group_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
		r.References["vpc_config.subnet_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		delete(r.TerraformResource.Schema, "filename")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"source_code_hash"},
		}
	})

	p.AddResourceConfigurator("aws_lambda_function_event_invoke_config", func(r *config.Resource) {
		r.References["destination_config.on_failure.destination"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue",
			Extractor: common.PathARNExtractor,
		}
		r.References["destination_config.on_success.destination"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sns/v1beta1.Topic",
			Extractor: common.PathARNExtractor,
		}
		delete(r.References, "function_name")
		delete(r.References, "qualifier")
	})

	p.AddResourceConfigurator("aws_lambda_function_url", func(r *config.Resource) {
		r.References["function_name"] = config.Reference{
			Type: "Function",
		}
	})

	p.AddResourceConfigurator("aws_lambda_invocation", func(r *config.Resource) {
		r.References["function_name"] = config.Reference{
			Type: "Function",
		}
	})

	p.AddResourceConfigurator("aws_lambda_permission", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"statement_id_prefix"},
		}
		r.References["function_name"] = config.Reference{
			Type: "Function",
		}
		r.References["qualifier"] = config.Reference{
			Type: "Alias",
		}
		delete(r.References, "source_arn")
	})

	p.AddResourceConfigurator("aws_lambda_provisioned_concurrency_config", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"provisioned_concurrent_executions"},
		}
		r.UseAsync = true
		delete(r.References, "function_name")
		delete(r.References, "qualifier")
	})
}
