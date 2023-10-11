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

package datasync

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the datasync group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_datasync_task", func(r *config.Resource) {
		r.References["destination_location_arn"] = config.Reference{
			Type: "LocationS3",
		}
		r.References["source_location_arn"] = config.Reference{
			Type: "LocationS3",
		}
		r.References["cloudwatch_log_group_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1.Group",
			Extractor: common.PathARNExtractor,
		}
	})
}
