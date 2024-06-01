// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package datasync

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-upjet-aws/config/common"
)

// Configure adds configurations for the datasync group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_datasync_task", func(r *config.Resource) {
		r.References["destination_location_arn"] = config.Reference{
			TerraformName: "aws_datasync_location_s3",
		}
		r.References["source_location_arn"] = config.Reference{
			TerraformName: "aws_datasync_location_s3",
		}
		r.References["cloudwatch_log_group_arn"] = config.Reference{
			TerraformName: "aws_cloudwatch_log_group",
			Extractor:     common.PathARNExtractor,
		}
	})
}
