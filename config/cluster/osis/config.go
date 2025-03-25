// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package osis

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the osis group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_osis_pipeline", func(r *config.Resource) {
		r.References["vpc_options.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}

		r.References["vpc_options.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}

		r.References["encrypt_at_rest.kms_key_arn"] = config.Reference{
			// its KMS key ARN in AWS API
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		r.References["log_publishing_options.cloudwatch_log_destination.log_group"] = config.Reference{
			TerraformName: "aws_cloudwatch_log_group",
		}
	})
}
