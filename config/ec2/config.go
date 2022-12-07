/*
Copyright 2021 Upbound Inc.
*/

package ec2

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for ec2 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_instance", func(r *config.Resource) {
		r.UseAsync = true
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
		r.References["vpc_security_group_ids"] = config.Reference{
			Type:              "SecurityGroup",
			RefFieldName:      "VPCSecurityGroupIDRefs",
			SelectorFieldName: "VPCSecurityGroupIDSelector",
		}
		r.References["security_groups"] = config.Reference{
			Type:              "SecurityGroup",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["root_block_device.kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.References["network_interface.network_interface_id"] = config.Reference{
			Type: "NetworkInterface",
		}
		r.References["ebs_block_device.kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.LateInitializer = config.LateInitializer{
			// NOTE(muvaf): These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/upbound/upjet/issues/107
			IgnoredFields: []string{
				"subnet_id",
				"network_interface",
				"private_ip",
				"source_dest_check",
				"vpc_security_group_ids",
				"associate_public_ip_address",
			},
		}
		config.MoveToStatus(r.TerraformResource, "security_groups")
	})
	p.AddResourceConfigurator("aws_eip", func(r *config.Resource) {
		r.References["instance"] = config.Reference{
			Type: "Instance",
		}
		r.References["network_interface"] = config.Reference{
			Type: "NetworkInterface",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route", func(r *config.Resource) {
		r.References["transit_gateway_attachment_id"] = config.Reference{
			Type: "TransitGatewayVPCAttachment",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table", func(r *config.Resource) {
		r.References["transit_gateway_id"] = config.Reference{
			Type: "TransitGateway",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table_association", func(r *config.Resource) {
		r.References["transit_gateway_attachment_id"] = config.Reference{
			Type: "TransitGatewayVPCAttachment",
		}
		r.References["transit_gateway_route_table_id"] = config.Reference{
			Type: "TransitGatewayRouteTable",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_vpc_attachment", func(r *config.Resource) {
		r.References["transit_gateway_id"] = config.Reference{
			Type: "TransitGateway",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_vpc_attachment_accepter", func(r *config.Resource) {
		r.References["transit_gateway_attachment_id"] = config.Reference{
			Type: "TransitGatewayVPCAttachment",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_connect", func(r *config.Resource) {
		r.References["subnet_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.References["vpc_id"] = config.Reference{
			Type:              "VPC",
			RefFieldName:      "VPCIDRef",
			SelectorFieldName: "VPCIDSelector",
		}
	})

	p.AddResourceConfigurator("aws_launch_template", func(r *config.Resource) {
		r.References["security_group_names"] = config.Reference{
			Type:              "SecurityGroup",
			RefFieldName:      "SecurityGroupNameRefs",
			SelectorFieldName: "SecurityGroupNameSelector",
		}
		r.References["block_device_mappings.ebs.kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.References["iam_instance_profile.arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.InstanceProfile",
			Extractor: common.PathARNExtractor,
		}
		r.References["iam_instance_profile.name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/iam/v1beta1.InstanceProfile",
		}
		r.References["network_interfaces.network_interface_id"] = config.Reference{
			Type: "NetworkInterface",
		}
		r.References["network_interfaces.security_groups"] = config.Reference{
			Type:              "SecurityGroup",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["network_interfaces.subnet_id"] = config.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("aws_vpc_endpoint", func(r *config.Resource) {
		// Mutually exclusive with:
		// aws_vpc_endpoint_subnet_association
		// aws_vpc_endpoint_route_table_association
		// aws_vpc_endpoint_security_group_association
		config.MoveToStatus(r.TerraformResource, "subnet_ids", "security_group_ids", "route_table_ids")
		delete(r.References, "vpc_endpoint_type")
	})

	p.AddResourceConfigurator("aws_subnet", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			// NOTE(muvaf): Conflicts with AvailabilityZone. See the following
			// for more details: https://github.com/upbound/upjet/issues/107
			IgnoredFields: []string{
				"availability_zone_id",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_network_interface", func(r *config.Resource) {
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
		r.References["security_groups"] = config.Reference{
			Type:              "SecurityGroup",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["attachment.instance"] = config.Reference{
			Type: "Instance",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"interface_type", "private_ip_list", "private_ips",
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
			Type: "SecurityGroup",
		}
		r.References["source_security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"self", "source_security_group_id",
			},
		}
	})

	p.AddResourceConfigurator("aws_vpc_peering_connection", func(r *config.Resource) {
		// Mutually exclusive with aws_vpc_peering_connection_options
		config.MoveToStatus(r.TerraformResource, "accepter", "requester")
		r.References["peer_vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("aws_route", func(r *config.Resource) {
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
		r.References["gateway_id"] = config.Reference{
			Type: "InternetGateway",
		}
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
		r.References["network_interface_id"] = config.Reference{
			Type: "NetworkInterface",
		}
		r.References["transit_gateway_id"] = config.Reference{
			Type: "TransitGateway",
		}
		r.References["vpc_peering_connection_id"] = config.Reference{
			Type: "VPCPeeringConnection",
		}
		r.References["vpc_endpoint_id"] = config.Reference{
			Type: "VPCEndpoint",
		}
		r.References["nat_gateway_id"] = config.Reference{
			Type: "NATGateway",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_route_table", func(r *config.Resource) {
		// These are mutually exclusive with aws_route and aws_vpn_gateway_route_propagation.
		config.MoveToStatus(r.TerraformResource, "route", "propagating_vgws")
	})

	p.AddResourceConfigurator("aws_route_table_association", func(r *config.Resource) {
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("aws_main_route_table_association", func(r *config.Resource) {
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_route_table_propagation", func(r *config.Resource) {
		r.References["transit_gateway_attachment_id"] = config.Reference{
			Type: "TransitGatewayVPCAttachment",
		}
		r.References["transit_gateway_route_table_id"] = config.Reference{
			Type: "TransitGatewayRouteTable",
		}
	})

	p.AddResourceConfigurator("aws_nat_gateway", func(r *config.Resource) {
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
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
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ec2_transit_gateway_multicast_domain", func(r *config.Resource) {
		r.References["transit_gateway_id"] = config.Reference{
			Type: "TransitGateway",
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
	})

	p.AddResourceConfigurator("aws_ec2_traffic_mirror_target", func(r *config.Resource) {
		delete(r.References, "network_load_balancer_arn")
	})

	p.AddResourceConfigurator("aws_vpc_ipam_pool", func(r *config.Resource) {
		r.References["ipam_scope_id"] = config.Reference{
			Type: "VPCIpamScope",
		}
	})

	p.AddResourceConfigurator("aws_vpc_ipam_scope", func(r *config.Resource) {
		r.References["ipam_id"] = config.Reference{
			Type: "VPCIpam",
		}
	})

	p.AddResourceConfigurator("aws_ami", func(r *config.Resource) {
		r.References["ebs_block_device.snapshot_id"] = config.Reference{
			Type: "EBSSnapshot",
		}
	})

	p.AddResourceConfigurator("aws_ami_copy", func(r *config.Resource) {
		r.References["source_ami_id"] = config.Reference{
			Type: "AMI",
		}
	})

	p.AddResourceConfigurator("aws_ami_launch_permission", func(r *config.Resource) {
		r.References["image_id"] = config.Reference{
			Type: "AMI",
		}
	})

	p.AddResourceConfigurator("aws_vpn_connection", func(r *config.Resource) {
		r.References["vpn_gateway_id"] = config.Reference{
			Type: "VPNGateway",
		}
	})
}
