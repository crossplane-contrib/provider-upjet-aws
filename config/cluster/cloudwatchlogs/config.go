// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudwatchlogs

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the cloudwatchlogs group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_cloudwatch_log_destination", func(r *config.Resource) {
		// the target_arn field is generated together with the associated
		// referencer fields but the auto-generated extractor refers to
		// the `spec` of the target kinesis.Stream, whereas the ARN resides
		// in the `status`. This is caused by the fact that kinesis.Stream has
		// resource configuration that moves the `arn` field from `spec` to
		// `status` but this resource configuration is currently not available to
		// reference injector (because reference injections happens at an earlier
		// stage). For now, we are just replacing the auto-injected reference
		// with a manual reference with the correct extractor. But if
		// we observe more cases of this, we may reconsider making reference
		// injection be able to consume the last state of resource configuration.
		r.References["target_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
	p.AddResourceConfigurator("aws_cloudwatch_log_subscription_filter", func(r *config.Resource) {
		// Please see the comment for aws_cloudwatch_log_destination.target_arn
		// reference configuration
		r.References["destination_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_cloudwatch_log_group", func(r *config.Resource) {
		r.MarkAsRequired("name")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name_prefix"},
		}
		// The kms_key_id field actually references the KMS ARN
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
	})

}
