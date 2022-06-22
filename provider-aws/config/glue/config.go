package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"
)

// Configure glue resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_glue_partition_index", func(r *config.Resource) {
		// OmittedFields works only for the top-level fields.
		// See https://github.com/upbound/upjet/issues/21
		delete(r.TerraformResource.Schema["partition_index"].Elem.(*schema.Resource).Schema, "index_name")
	})

	p.AddResourceConfigurator("aws_glue_data_catalog_encryption_settings", func(r *config.Resource) {
		r.References["data_catalog_encryption_settings.connection_password_encryption.aws_kms_key_id"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
		}
		r.References["data_catalog_encryption_settings.encryption_at_rest.sse_aws_kms_key_id"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
		}
	})
}
