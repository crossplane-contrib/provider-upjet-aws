/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package globalaccelerator

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for ebs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_globalaccelerator_accelerator", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_globalaccelerator_endpoint_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"listener_arn": {
				Type: "Listener",
			},
		}
	})

	p.AddResourceConfigurator("aws_globalaccelerator_listener", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"accelerator_arn": {
				Type: "Accelerator",
			},
		}
	})
}
