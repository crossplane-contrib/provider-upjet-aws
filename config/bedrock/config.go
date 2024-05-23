// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bedrock

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the bedrock group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_bedrock_model_invocation_logging_configuration", func(r *config.Resource) {
		r.References["logging_config.s3_config.bucket_name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["logging_config.cloudwatch_config.role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		
	})
}