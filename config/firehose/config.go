package firehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the firehose group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kinesis_firehose_delivery_stream", func(r *config.Resource) {
		r.TerraformResource.Schema["splunk_configuration"].Elem.(*schema.Resource).Schema["hec_token"].Sensitive = true

		r.References["extended_s3_configuration.role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["extended_s3_configuration.bucket_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
			Extractor: common.PathARNExtractor,
		}

		r.References["s3_configuration.role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["s3_configuration.bucket_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
			Extractor: common.PathARNExtractor,
		}

		r.References["redshift_configuration.s3_backup_configuration.bucket_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)`,
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"server_side_encryption",
				"version_id",
			},
		}

		delete(r.References, "extended_s3_configuration.data_format_conversion_configuration.schema_configuration.database_name")
	})
}
