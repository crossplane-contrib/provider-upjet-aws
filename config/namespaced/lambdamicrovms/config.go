// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package lambdamicrovms

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

// Configure adds configurations for the lambdamicrovms group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lambdamicrovms_image", func(r *config.Resource) {
		r.References["build_role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["code_artifact.uri"] = config.Reference{
			TerraformName: "aws_s3_object",
			Extractor:     common.PathS3ObjectURIExtractor,
		}
		r.References["egress_network_connectors"] = config.Reference{
			TerraformName: "aws_lambdacore_network_connector",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("code_artifact", "codeArtifact")
	})

	p.AddResourceConfigurator("aws_lambdamicrovms_microvm", func(r *config.Resource) {
		r.References["egress_network_connectors"] = config.Reference{
			TerraformName: "aws_lambdacore_network_connector",
			Extractor:     common.PathARNExtractor,
		}
		r.References["execution_role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["image_arn"] = config.Reference{
			TerraformName: "aws_lambdamicrovms_image",
			Extractor:     common.PathARNExtractor,
		}
		r.References["logging.cloudwatch.log_group"] = config.Reference{
			TerraformName: "aws_cloudwatch_log_group",
		}
		r.AddSingletonListConversion("idle_policy", "idlePolicy")
		r.AddSingletonListConversion("logging", "logging")
		r.AddSingletonListConversion("logging[*].cloudwatch", "logging[*].cloudwatch")
		r.AddSingletonListConversion("logging[*].disabled", "logging[*].disabled")
	})
}
