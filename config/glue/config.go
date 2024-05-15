// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package glue

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the glue group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_glue_catalog_database", func(r *config.Resource) {
		// Required in ID but optional in schema since TF defaults to Account ID.
		// This causes refresh to fail in the first reconcile.
		r.TerraformResource.Schema["catalog_id"].Computed = false
		r.TerraformResource.Schema["catalog_id"].Optional = false
	})

	p.AddResourceConfigurator("aws_glue_catalog_table", func(r *config.Resource) {
		// Required in ID but optional in schema since TF defaults to Account ID.
		// This causes refresh to fail in the first reconcile.
		r.TerraformResource.Schema["catalog_id"].Computed = false
		r.TerraformResource.Schema["catalog_id"].Optional = false

		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "partition_index.#")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("aws_glue_connection", func(r *config.Resource) {
		// Required in ID but optional in schema since TF defaults to Account ID.
		// This causes refresh to fail in the first reconcile.
		r.TerraformResource.Schema["catalog_id"].Required = true
		r.TerraformResource.Schema["catalog_id"].Computed = false
		r.TerraformResource.Schema["catalog_id"].Optional = false
	})

	p.AddResourceConfigurator("aws_glue_user_defined_function", func(r *config.Resource) {
		// Required in ID but optional in schema since TF defaults to Account ID.
		// This causes refresh to fail in the first reconcile.
		r.TerraformResource.Schema["catalog_id"].Computed = false
		r.TerraformResource.Schema["catalog_id"].Optional = false
		delete(r.References, "catalog_id")
	})

	p.AddResourceConfigurator("aws_glue_crawler", func(r *config.Resource) {
		r.References["role"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_glue_data_catalog_encryption_settings", func(r *config.Resource) {
		r.References["data_catalog_encryption_settings.connection_password_encryption.aws_kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["data_catalog_encryption_settings.encryption_at_rest.sse_aws_kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_glue_job", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			//
			IgnoredFields: []string{
				"max_capacity", "number_of_workers", "worker_type",
			},
		}

	})

	p.AddResourceConfigurator("aws_glue_security_configuration", func(r *config.Resource) {
		r.References["encryption_configuration.cloudwatch_encryption.kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		r.References["encryption_configuration.job_bookmarks_encryption.kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		r.References["encryption_configuration.s3_encryption.kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
	})

}
