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

package directconnect

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for directconnect group.
func Configure(p *config.Provider) { // nolint:gocyclo
	p.AddResourceConfigurator("aws_dx_public_virtual_interface", func(r *config.Resource) {
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
	})
	p.AddResourceConfigurator("aws_dx_private_virtual_interface", func(r *config.Resource) {
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
		r.References["vpn_gateway_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.VPNGateway",
		}
	})
	p.AddResourceConfigurator("aws_dx_gateway_association", func(r *config.Resource) {
		r.TerraformResource.Schema["associated_gateway_id"].Required = true
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"associated_gateway_owner_account_id"},
		}
	})
	p.AddResourceConfigurator("aws_dx_hosted_transit_virtual_interface", func(r *config.Resource) {
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
	})

	p.AddResourceConfigurator("aws_dx_hosted_public_virtual_interface", func(r *config.Resource) {
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
	})

	p.AddResourceConfigurator("aws_dx_hosted_private_virtual_interface", func(r *config.Resource) {
		r.References["connection_id"] = config.Reference{
			Type: "Connection",
		}
	})

	p.AddResourceConfigurator("aws_dx_hosted_private_virtual_interface_accepter", func(r *config.Resource) {
		r.References["virtual_interface_id"] = config.Reference{
			Type: "HostedPrivateVirtualInterface",
		}
	})

	p.AddResourceConfigurator("aws_dx_hosted_public_virtual_interface_accepter", func(r *config.Resource) {
		r.References["virtual_interface_id"] = config.Reference{
			Type: "HostedPublicVirtualInterface",
		}
	})

	p.AddResourceConfigurator("aws_dx_hosted_transit_virtual_interface_accepter", func(r *config.Resource) {
		r.References["virtual_interface_id"] = config.Reference{
			Type: "HostedTransitVirtualInterface",
		}
	})
}
