package mwaa

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the sagemaker group.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("aws_mwaa_environment", func(r *config.Resource) {
		r.References = config.References{
			"execution_role_arn": config.Reference{
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
			"network_configuration.subnet_ids": config.Reference{
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetIDRefs",
				SelectorFieldName: "SubnetIDSelector",
			},
			"network_configuration.security_groups": config.Reference{
				TerraformName:     "aws_security_group",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
			"source_bucket_arn": config.Reference{
				TerraformName: "aws_s3_bucket",
			},
		}
		r.UseAsync = true

	})

}
