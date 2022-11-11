package glue

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure glue resources.
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
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_glue_data_catalog_encryption_settings", func(r *config.Resource) {
		r.References["data_catalog_encryption_settings.connection_password_encryption.aws_kms_key_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
		r.References["data_catalog_encryption_settings.encryption_at_rest.sse_aws_kms_key_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_glue_security_configuration", func(r *config.Resource) {
		r.References["encryption_configuration.cloudwatch_encryption.kms_key_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}

		r.References["encryption_configuration.job_bookmarks_encryption.kms_key_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}

		r.References["encryption_configuration.s3_encryption.kms_key_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_glue_connection", func(r *config.Resource) {
		// Catalog ID is a required field but set as optional in schema since TF defaults it to Account ID.
		// We need to mark it as required since it is used as part of "id" field in tfstate.
		r.TerraformResource.Schema["catalog_id"].Required = true
		r.TerraformResource.Schema["catalog_id"].Optional = false
		r.TerraformResource.Schema["catalog_id"].Computed = false
	})

}
