// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudwatchlogs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for cloudwatchlogs group.
func Configure(p *config.Provider) {
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
		config.MarkAsRequired(r.TerraformResource, "name")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name_prefix"},
		}
	})

}
