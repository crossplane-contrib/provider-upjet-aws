// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package lambda

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the lambda group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_lambda_alias", func(r *config.Resource) {
		r.References["function_name"] = config.Reference{
			TerraformName: "aws_lambda_function",
		}
	})

	p.AddResourceConfigurator("aws_lambda_code_signing_config", func(r *config.Resource) {
		r.References["allowed_publishers.signing_profile_version_arns"] = config.Reference{
			TerraformName: "aws_signer_signing_profile",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_lambda_event_source_mapping", func(r *config.Resource) {
		r.References["function_name"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
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
			TerraformName: "aws_s3_bucket",
		}
		r.References["role"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["vpc_config.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
		r.References["vpc_config.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		delete(r.TerraformResource.Schema, "filename")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"source_code_hash"},
		}
		r.MetaResource.ArgumentDocs["source_code_hash"] = "Used to trigger updates. Must be set to " +
			"a base64 encoded SHA256 hash of the package file specified with either filename or s3_key. " +
			"If you have specified this field manually, it should be the actual (computed) hash of the " +
			"underlying lambda function specified in the filename, image_uri, s3_bucket fields."
	})

	p.AddResourceConfigurator("aws_lambda_function_event_invoke_config", func(r *config.Resource) {
		r.References["destination_config.on_failure.destination"] = config.Reference{
			TerraformName: "aws_sqs_queue",
			Extractor:     common.PathARNExtractor,
		}
		r.References["destination_config.on_success.destination"] = config.Reference{
			TerraformName: "aws_sns_topic",
			Extractor:     common.PathARNExtractor,
		}
		delete(r.References, "function_name")
		delete(r.References, "qualifier")
	})

	p.AddResourceConfigurator("aws_lambda_function_url", func(r *config.Resource) {
		r.References["function_name"] = config.Reference{
			TerraformName: "aws_lambda_function",
		}
	})

	p.AddResourceConfigurator("aws_lambda_invocation", func(r *config.Resource) {
		r.References["function_name"] = config.Reference{
			TerraformName: "aws_lambda_function",
		}
	})

	p.AddResourceConfigurator("aws_lambda_permission", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"statement_id_prefix"},
		}
		r.References["function_name"] = config.Reference{
			TerraformName: "aws_lambda_function",
		}
		r.References["qualifier"] = config.Reference{
			TerraformName: "aws_lambda_alias",
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
