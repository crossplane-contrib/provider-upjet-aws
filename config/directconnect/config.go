// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package directconnect

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the directconnect group.
func Configure(p *config.Provider) {
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
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_dx_gateway_association", func(r *config.Resource) {
		r.TerraformResource.Schema["associated_gateway_id"].Required = true
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"associated_gateway_owner_account_id"},
		}
		r.UseAsync = true
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

	p.AddResourceConfigurator("aws_dx_connection", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"encryption_mode"},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_dx_macsec_key_association", func(r *config.Resource) {
		r.References["connection_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/directconnect/v1beta1.Connection",
		}
		r.References["secret_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_dx_bgp_peer", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_dx_transit_virtual_interface", func(r *config.Resource) {
		r.UseAsync = true
	})
}
