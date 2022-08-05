package firehose

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for firehose group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kinesis_firehose_delivery_stream", func(r *config.Resource) {
		r.References["extended_s3_configuration.role_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["extended_s3_configuration.bucket_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket",
			Extractor: common.PathARNExtractor,
		}

		r.References["s3_configuration.role_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["s3_configuration.bucket_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket",
			Extractor: common.PathARNExtractor,
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"server_side_encryption"},
		}

		delete(r.References, "extended_s3_configuration.data_format_conversion_configuration.schema_configuration.database_name")
	})
}
