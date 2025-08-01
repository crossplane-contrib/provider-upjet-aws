// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ec2

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the ec2 group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_instance", func(r *config.Resource) {
		r.UseAsync = true
		r.References["subnet_id"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["vpc_security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "VPCSecurityGroupIDRefs",
			SelectorFieldName: "VPCSecurityGroupIDSelector",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["root_block_device.kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
		}
		r.References["network_interface.network_interface_id"] = config.Reference{
			TerraformName: "aws_network_interface",
		}
		r.References["ebs_block_device.kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
		}
		r.LateInitializer = config.LateInitializer{
			// NOTE(muvaf): These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/crossplane/upjet/issues/107
			IgnoredFields: []string{
				"subnet_id",
				"network_interface",
				"private_ip",
				"source_dest_check",
				"vpc_security_group_ids",
				"associate_public_ip_address",
				"ipv6_addresses",
				"ipv6_address_count",
				"cpu_core_count",
				"cpu_threads_per_core",
				"cpu_options",
				"root_block_device",
			},
		}
		r.TerraformCustomDiff = common.RemoveDiffIfEmpty([]string{
			"volume_tags.%",
		})
		config.MoveToStatus(r.TerraformResource, "security_groups")
	})
	p.AddResourceConfigurator("aws_eip", func(r *config.Resource) {
		r.References["instance"] = config.Reference{
			TerraformName: "aws_instance",
		}
		r.References["network_interface"] = config.Reference{
			TerraformName: "aws_network_interface",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route", func(r *config.Resource) {
		r.References["transit_gateway_attachment_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway_vpc_attachment",
		}
		r.References["transit_gateway_route_table_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway_route_table",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table", func(r *config.Resource) {
		r.References["transit_gateway_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table_association", func(r *config.Resource) {
		r.References["transit_gateway_attachment_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway_vpc_attachment",
		}
		r.References["transit_gateway_route_table_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway_route_table",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_vpc_attachment", func(r *config.Resource) {
		r.References["transit_gateway_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_vpc_attachment_accepter", func(r *config.Resource) {
		r.References["transit_gateway_attachment_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway_vpc_attachment",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_connect", func(r *config.Resource) {
		r.References["subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName:     "aws_vpc",
			RefFieldName:      "VPCIDRef",
			SelectorFieldName: "VPCIDSelector",
		}
	})

	p.AddResourceConfigurator("aws_launch_template", func(r *config.Resource) {
		r.References["security_group_names"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupNameRefs",
			SelectorFieldName: "SecurityGroupNameSelector",
		}
		r.References["block_device_mappings.ebs.kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["iam_instance_profile.arn"] = config.Reference{
			TerraformName: "aws_iam_instance_profile",
			Extractor:     common.PathARNExtractor,
		}
		r.References["iam_instance_profile.name"] = config.Reference{
			TerraformName: "aws_iam_instance_profile",
		}
		r.References["network_interfaces.network_interface_id"] = config.Reference{
			TerraformName: "aws_network_interface",
		}
		r.References["network_interfaces.security_groups"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["network_interfaces.subnet_id"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"default_version"},
		}

	})

	p.AddResourceConfigurator("aws_vpc_endpoint", func(r *config.Resource) {
		// Mutually exclusive with:
		// aws_vpc_endpoint_subnet_association
		// aws_vpc_endpoint_route_table_association
		// aws_vpc_endpoint_security_group_association
		r.LateInitializer = config.LateInitializer{
			// Conflicts with VPCEndpointSubnetAssociation
			IgnoredFields: []string{
				"subnet_configuration",
			},
		}
		config.MoveToStatus(r.TerraformResource, "subnet_ids", "security_group_ids", "route_table_ids")
		delete(r.References, "vpc_endpoint_type")
	})

	p.AddResourceConfigurator("aws_subnet", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			// NOTE(muvaf): Conflicts with AvailabilityZone. See the following
			// for more details: https://github.com/crossplane/upjet/issues/107
			IgnoredFields: []string{
				"availability_zone_id",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_network_interface", func(r *config.Resource) {
		r.References["subnet_id"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["attachment.instance"] = config.Reference{
			TerraformName: "aws_instance",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"interface_type", "private_ip_list", "private_ips", "ipv6_addresses", "ipv6_address_count",
			},
		}
		// Mutually exclusive with aws_network_interface_attachment
		config.MoveToStatus(r.TerraformResource, "attachment")
	})

	p.AddResourceConfigurator("aws_security_group", func(r *config.Resource) {
		// Mutually exclusive with aws_security_group_rule
		config.MoveToStatus(r.TerraformResource, "ingress", "egress")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"name", "name_prefix",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_security_group_rule", func(r *config.Resource) {
		r.References["security_group_id"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["source_security_group_id"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["prefix_list_ids"] = config.Reference{
			TerraformName:     "aws_ec2_managed_prefix_list",
			RefFieldName:      "PrefixListIDRefs",
			SelectorFieldName: "PrefixListIDSelector",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"cidr_blocks",
				"ipv6_cidr_blocks",
				"self",
				"source_security_group_id",
			},
		}
		r.TerraformConfigurationInjector = func(jsonMap map[string]any, params map[string]any) error {
			// TODO: Has better be implemented via defaulting.
			if _, ok := jsonMap["self"]; !ok {
				params["self"] = false
			}
			return nil
		}
	})

	p.AddResourceConfigurator("aws_vpc_security_group_ingress_rule", func(r *config.Resource) {
		r.Kind = "SecurityGroupIngressRule"
		r.References["security_group_id"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["referenced_security_group_id"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["prefix_list_id"] = config.Reference{
			TerraformName: "aws_ec2_managed_prefix_list",
		}
	})

	p.AddResourceConfigurator("aws_vpc_security_group_egress_rule", func(r *config.Resource) {
		r.Kind = "SecurityGroupEgressRule"
		r.References["security_group_id"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["referenced_security_group_id"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["prefix_list_id"] = config.Reference{
			TerraformName: "aws_ec2_managed_prefix_list",
		}
	})

	p.AddResourceConfigurator("aws_vpc_peering_connection", func(r *config.Resource) {
		// Mutually exclusive with aws_vpc_peering_connection_options
		config.MoveToStatus(r.TerraformResource, "accepter", "requester")
		r.References["peer_vpc_id"] = config.Reference{
			TerraformName: "aws_vpc",
		}
	})

	p.AddResourceConfigurator("aws_route", func(r *config.Resource) {
		r.References["route_table_id"] = config.Reference{
			TerraformName: "aws_route_table",
		}
		r.References["gateway_id"] = config.Reference{
			TerraformName: "aws_internet_gateway",
		}
		r.References["instance_id"] = config.Reference{
			TerraformName: "aws_instance",
		}
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "aws_network_interface",
		}
		r.References["transit_gateway_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway",
		}
		r.References["vpc_peering_connection_id"] = config.Reference{
			TerraformName: "aws_vpc_peering_connection",
		}
		r.References["vpc_endpoint_id"] = config.Reference{
			TerraformName: "aws_vpc_endpoint",
		}
		r.References["nat_gateway_id"] = config.Reference{
			TerraformName: "aws_nat_gateway",
		}
		r.References["destination_prefix_list_id"] = config.Reference{
			TerraformName: "aws_ec2_managed_prefix_list",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_route_table", func(r *config.Resource) {
		// These are mutually exclusive with aws_route and aws_vpn_gateway_route_propagation.
		config.MoveToStatus(r.TerraformResource, "route", "propagating_vgws")
	})

	p.AddResourceConfigurator("aws_route_table_association", func(r *config.Resource) {
		r.References["subnet_id"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["route_table_id"] = config.Reference{
			TerraformName: "aws_route_table",
		}
	})

	p.AddResourceConfigurator("aws_main_route_table_association", func(r *config.Resource) {
		r.References["route_table_id"] = config.Reference{
			TerraformName: "aws_route_table",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table_propagation", func(r *config.Resource) {
		r.References["transit_gateway_attachment_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway_vpc_attachment",
		}
		r.References["transit_gateway_route_table_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway_route_table",
		}
	})

	p.AddResourceConfigurator("aws_nat_gateway", func(r *config.Resource) {
		r.References["subnet_id"] = config.Reference{
			TerraformName: "aws_subnet",
		}
	})

	p.AddResourceConfigurator("aws_network_acl", func(r *config.Resource) {
		// Mutually exclusive with:
		// aws_network_acl_rule
		config.MoveToStatus(r.TerraformResource, "ingress", "egress")
	})

	p.AddResourceConfigurator("aws_vpc_endpoint_service", func(r *config.Resource) {
		// Mutually exclusive with:
		// vpc_endpoint_service_allowed_principal
		config.MoveToStatus(r.TerraformResource, "allowed_principals")
	})

	p.AddResourceConfigurator("aws_flow_log", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"log_format", "log_destination", "log_group_name"},
		}
	})

	p.AddResourceConfigurator("aws_network_acl_rule", func(r *config.Resource) {
		delete(r.References, "cidr_block")
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_peering_attachment", func(r *config.Resource) {
		delete(r.References, "peer_account_id")
	})

	p.AddResourceConfigurator("aws_spot_datafeed_subscription", func(r *config.Resource) {
		delete(r.References, "bucket")
	})

	p.AddResourceConfigurator("aws_vpc", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"ipv6_cidr_block",
				"cidr_block",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_multicast_domain", func(r *config.Resource) {
		r.References["transit_gateway_id"] = config.Reference{
			TerraformName: "aws_ec2_transit_gateway",
		}
	})

	p.AddResourceConfigurator("aws_spot_instance_request", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"valid_until",
				"valid_from",
				"instance_interruption_behavior",
				"source_dest_check",
				"spot_type",
			},
		}

		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "enclave_options.#")
				delete(diff.Attributes, "metadata_options.#")
				delete(diff.Attributes, "maintenance_options.#")
				delete(diff.Attributes, "cpu_options.#")
				delete(diff.Attributes, "network_interface.#")
				delete(diff.Attributes, "capacity_reservation_specification.#")
				delete(diff.Attributes, "ephemeral_block_device.#")
				delete(diff.Attributes, "secondary_private_ips.#")
				delete(diff.Attributes, "private_dns_name_options.#")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("aws_ec2_traffic_mirror_target", func(r *config.Resource) {
		delete(r.References, "network_load_balancer_arn")
	})

	p.AddResourceConfigurator("aws_vpc_ipam_pool", func(r *config.Resource) {
		r.References["ipam_scope_id"] = config.Reference{
			TerraformName: "aws_vpc_ipam_scope",
		}
	})

	p.AddResourceConfigurator("aws_vpc_ipam_scope", func(r *config.Resource) {
		r.References["ipam_id"] = config.Reference{
			TerraformName: "aws_vpc_ipam",
		}
	})

	p.AddResourceConfigurator("aws_ami", func(r *config.Resource) {
		r.References["ebs_block_device.snapshot_id"] = config.Reference{
			TerraformName: "aws_ebs_snapshot",
		}
	})

	p.AddResourceConfigurator("aws_ami_copy", func(r *config.Resource) {
		r.References["source_ami_id"] = config.Reference{
			TerraformName: "aws_ami",
		}
		r.TerraformConfigurationInjector = func(jsonMap map[string]any, params map[string]any) error {
			params["ebs_block_device"] = []any{}
			// TODO: Has better be implemented via defaulting.
			if _, ok := jsonMap["encrypted"]; !ok {
				params["encrypted"] = false
			}
			return nil
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "ebs_block_device.#")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("aws_ami_launch_permission", func(r *config.Resource) {
		r.References["image_id"] = config.Reference{
			TerraformName: "aws_ami",
		}
	})

	p.AddResourceConfigurator("aws_vpn_connection", func(r *config.Resource) {
		r.References["vpn_gateway_id"] = config.Reference{
			TerraformName: "aws_vpn_gateway",
		}
	})

	p.AddResourceConfigurator("aws_ec2_tag", func(r *config.Resource) {
		delete(r.References, "resource_id")
	})
}
