// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

// Copyright 2022 Crossplane Authors
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

package cloudwatchevents

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the cloudwatchevents group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudwatch_event_permission", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			TerraformName: "aws_cloudwatch_event_bus",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_rule", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			TerraformName: "aws_cloudwatch_event_bus",
		}
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_target", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			TerraformName: "aws_cloudwatch_event_bus",
		}
		delete(r.References, "arn")
	})
}
