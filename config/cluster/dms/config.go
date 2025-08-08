// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dms

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the dms group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_dms_endpoint", func(r *config.Resource) {
		r.References = config.References{
			"secrets_manager_access_role_arn": {
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
			"service_access_role": {
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
			"kms_key_arn": {
				TerraformName: "aws_kms_key",
				Extractor:     common.PathARNExtractor,
			},
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "redshift_settings.#")
			}
			return diff, nil
		}
	})
}
