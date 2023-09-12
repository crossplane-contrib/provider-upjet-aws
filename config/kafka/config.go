package kafka

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for kafka group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_msk_cluster", func(r *config.Resource) {
		r.References["encryption_info.encryption_at_rest_kms_key_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
		r.References["logging_info.broker_logs.s3.bucket"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["logging_info.broker_logs.cloudwatch_logs.log_group"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1.Group",
		}
		r.References["broker_node_group_info.client_subnets"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
		r.References["broker_node_group_info.security_groups"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
		}
		r.References["configuration_info.arn"] = config.Reference{
			Type: "Configuration",
			Extractor: common.PathARNExtractor,
			SelectorFieldName: "ConfigurationInfoSelector",
			RefFieldName: "ConfigurationInfoRef",
		}
		r.References["configuration_info.revision"] = config.Reference{
			Type: "Configuration",
			Extractor: "GetConfigurationRevision",
			SelectorFieldName: "ConfigurationInfoSelector",
			RefFieldName: "ConfigurationInfoRef",
		}
		r.UseAsync = true
	})
}

// Lovingly ripped off from config/common/ARNExtractor. Is there a better way?
func GetConfigurationRevision() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.revision")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}