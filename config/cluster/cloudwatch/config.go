// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudwatch

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the cloudwatch group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_cloudwatch_metric_stream", func(r *config.Resource) {
		r.MarkAsRequired("name")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name_prefix"},
		}
	})

	p.AddResourceConfigurator("aws_cloudwatch_event_target", func(r *config.Resource) {
		r.References["arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
