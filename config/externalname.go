/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/errors"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{

	// ACM
	// Imported using ARN that has a random substring:
	// arn:aws:acm:eu-central-1:123456789012:certificate/7e7a28d2-163f-4b8f-b9cd-822f96c08d6a
	"aws_acm_certificate": config.IdentifierFromProvider,
	// No import documented, but https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate_validation#id
	"aws_acm_certificate_validation": config.IdentifierFromProvider,

	// ACM PCA
	// aws_acmpca_certificate can not be imported at this time.
	"aws_acmpca_certificate": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	//	// arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
	"aws_acmpca_certificate_authority": config.IdentifierFromProvider,
	// No doc on import, but resource is getting CA ARN:
	// arn:aws:acm-pca:eu-central-1:609897127049:certificate-authority/ba0c7989-9641-4f36-a033-dee60121d595
	"aws_acmpca_certificate_authority_certificate": config.IdentifierFromProvider,

	// amp
	//
	// ID is a random UUID.
	"aws_prometheus_workspace":            config.IdentifierFromProvider,
	"aws_prometheus_rule_group_namespace": config.TemplatedStringAsIdentifier("name", "arn:aws:aps:{{ .parameters.region }}:{{ .client_metadata.account_id }}:rulegroupsnamespace/IDstring/{{ .external_name }}"),
	// Uses the ID of workspace, workspace_id parameter.
	"aws_prometheus_alert_manager_definition": config.IdentifierFromProvider,

	// apigatewayv2
	//
	"aws_apigatewayv2_api": config.IdentifierFromProvider,
	// Case4: Imported by using the API mapping identifier and domain name.
	"aws_apigatewayv2_api_mapping": TemplatedStringAsIdentifierWithNoName("{{ .external_name }}/{{ .parameters.domain_name }}"),
	// Case4: Imported by using the API identifier and authorizer identifier.
	"aws_apigatewayv2_authorizer": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier and deployment identifier.
	"aws_apigatewayv2_deployment":  TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	"aws_apigatewayv2_domain_name": config.ParameterAsIdentifier("domain_name"),
	// Case4: Imported by using the API identifier and integration identifier.
	"aws_apigatewayv2_integration": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier, integration identifier and
	// integration response identifier.
	"aws_apigatewayv2_integration_response": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .parameters.integration_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier and model identifier.
	"aws_apigatewayv2_model": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier and route identifier.
	"aws_apigatewayv2_route": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier, route identifier and route
	// response identifier.
	"aws_apigatewayv2_route_response": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .parameters.route_id }}/{{ .external_name }}"),
	// Imported by using the API identifier and stage name.
	"aws_apigatewayv2_stage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}/{{ .external_name }}"),
	// aws_apigatewayv2_vpc_link can be imported by using the VPC Link id
	"aws_apigatewayv2_vpc_link": config.IdentifierFromProvider,

	// autoscaling
	//
	"aws_autoscaling_group": config.NameAsIdentifier,
	// No terraform import.
	"aws_autoscaling_attachment": config.IdentifierFromProvider,

	// DynamoDB Table Items can be imported using the name
	"aws_dynamodb_table_item": config.IdentifierFromProvider,
	// DynamoDB contributor insights
	"aws_dynamodb_contributor_insights": config.IdentifierFromProvider,
	// Dynamodb Kinesis streaming destinations are imported using "table_name,stream_arn"
	"aws_dynamodb_kinesis_streaming_destination": config.IdentifierFromProvider,

	// cloudtrail
	//
	// Cloudtrails can be imported using the name
	"aws_cloudtrail": config.NameAsIdentifier,
	// Event data stores can be imported using their arn
	"aws_cloudtrail_event_data_store": config.IdentifierFromProvider,

	// cognitoidentity
	//
	// us-west-2_abc123
	"aws_cognito_identity_pool": config.IdentifierFromProvider,
	// us-west-2:b64805ad-cb56-40ba-9ffc-f5d8207e6d42
	"aws_cognito_identity_pool_roles_attachment": config.IdentifierFromProvider,
	// us-west-2_abc123:CorpAD
	"aws_cognito_identity_pool_provider_principal_tag": config.IdentifierFromProvider,

	// cognitoidp
	//
	// us-west-2_abc123
	"aws_cognito_user_pool": config.IdentifierFromProvider,
	// us-west-2_abc123/3ho4ek12345678909nh3fmhpko
	"aws_cognito_user_pool_client": config.IdentifierFromProvider,
	// auth.example.org
	"aws_cognito_user_pool_domain": config.IdentifierFromProvider,
	// us-west-2_ZCTarbt5C,12bu4fuk3mlgqa2rtrujgp6egq
	"aws_cognito_user_pool_ui_customization": config.IdentifierFromProvider,
	// Cognito User Groups can be imported using the user_pool_id/name attributes concatenated:
	// us-east-1_vG78M4goG/user-group
	// Following configuration does not work: FormattedIdentifierUserDefinedNameLast("name", "/", "user_pool_id")
	// As it fails with a user group not found sync error
	// TODO: check if this is due to any diff between Terraform import & apply
	// implementations. Currently, the API is not normalized.
	"aws_cognito_user_group": config.IdentifierFromProvider,
	// us-west-2_abc123|https://example.com
	"aws_cognito_resource_server": config.IdentifierFromProvider,
	// us-west-2_abc123:CorpAD
	"aws_cognito_identity_provider": config.IdentifierFromProvider,
	// user_pool_id/name: us-east-1_vG78M4goG/user
	"aws_cognito_user": config.TemplatedStringAsIdentifier("username", "{{ .parameters.user_pool_id }}/{{ .external_name }}"),
	// no doc
	"aws_cognito_user_in_group": config.IdentifierFromProvider,

	// ebs
	//
	// EBS Volumes can be imported using the id: vol-049df61146c4d7901
	"aws_ebs_volume": config.IdentifierFromProvider,
	// EBS Snapshot can be imported using the id
	"aws_ebs_snapshot": config.IdentifierFromProvider,
	// No import
	"aws_ebs_snapshot_copy": config.IdentifierFromProvider,
	// No import
	"aws_ebs_snapshot_import": config.IdentifierFromProvider,

	// ec2
	//
	// Instances can be imported using the id: i-12345678
	"aws_instance": config.IdentifierFromProvider,
	// No terraform import.
	"aws_eip": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway identifier: tgw-12345678
	"aws_ec2_transit_gateway": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Route Table, an underscore,
	// and the destination CIDR: tgw-rtb-12345678_0.0.0.0/0
	"aws_ec2_transit_gateway_route": FormattedIdentifierFromProvider("_", "transit_gateway_route_table_id", "destination_cidr_block"),
	// Imported by using the EC2 Transit Gateway Route Table identifier:
	// tgw-rtb-12345678
	"aws_ec2_transit_gateway_route_table": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Route Table identifier, an
	// underscore, and the EC2 Transit Gateway Attachment identifier, e.g.,
	// tgw-rtb-12345678_tgw-attach-87654321
	"aws_ec2_transit_gateway_route_table_association": FormattedIdentifierFromProvider("_", "transit_gateway_route_table_id", "transit_gateway_attachment_id"),
	// Imported by using the EC2 Transit Gateway Attachment identifier:
	// tgw-attach-12345678
	"aws_ec2_transit_gateway_vpc_attachment": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Attachment identifier: tgw-attach-12345678
	"aws_ec2_transit_gateway_vpc_attachment_accepter": FormattedIdentifierFromProvider("", "transit_gateway_attachment_id"),
	// Imported using the id: lt-12345678
	"aws_launch_template": config.IdentifierFromProvider,
	// Launch configurations can be imported using the name
	"aws_launch_configuration": config.NameAsIdentifier,
	// Imported using the id: vpc-23123
	"aws_vpc": config.IdentifierFromProvider,
	// Imported using the vpc endpoint id: vpce-3ecf2a57
	"aws_vpc_endpoint": config.IdentifierFromProvider,
	// Imported using the subnet id: subnet-9d4a7b6c
	"aws_subnet": config.IdentifierFromProvider,
	// Imported using the id: eni-e5aa89a3
	"aws_network_interface": config.IdentifierFromProvider,
	// Imported using the id: sg-903004f8
	"aws_security_group": config.IdentifierFromProvider,
	// Imported using a very complex format:
	// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule
	"aws_security_group_rule": config.IdentifierFromProvider,
	// Imported by using the VPC CIDR Association ID: vpc-cidr-assoc-xxxxxxxx
	"aws_vpc_ipv4_cidr_block_association": config.IdentifierFromProvider,
	// Imported using the vpc peering id: pcx-111aaa111
	"aws_vpc_peering_connection": config.IdentifierFromProvider,
	// Imported using the vpc peering id: pcx-111aaa111
	"aws_vpc_peering_connection_options": config.IdentifierFromProvider,
	// Imported using the peering connection id: pcx-12345678
	"aws_vpc_peering_connection_accepter": config.IdentifierFromProvider,
	// Imported using the following format: ROUTETABLEID_DESTINATION
	"aws_route": route(),
	// Imported using id: rtb-4e616f6d69
	"aws_route_table": config.IdentifierFromProvider,
	// Imported using the associated resource ID and Route Table ID separated
	// by a forward slash (/)
	"aws_route_table_association": routeTableAssociation(),
	// No import.
	"aws_main_route_table_association": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_group_member": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_group_source": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Route Table identifier, an
	// underscore, and the EC2 Transit Gateway Attachment identifier:
	// tgw-rtb-12345678_tgw-attach-87654321
	"aws_ec2_transit_gateway_route_table_propagation": FormattedIdentifierFromProvider("_", "transit_gateway_attachment_id", "transit_gateway_route_table_id"),
	// Imported using the id: igw-c0a643a9
	"aws_internet_gateway": config.IdentifierFromProvider,
	// NAT Gateways can be imported using the id
	"aws_nat_gateway": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_multicast_domain can be imported by using the EC2 Transit Gateway Multicast Domain identifier
	"aws_ec2_transit_gateway_multicast_domain": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_domain_association": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_peering_attachment can be imported by using the EC2 Transit Gateway Attachment identifier
	"aws_ec2_transit_gateway_peering_attachment": config.IdentifierFromProvider,
	// Prefix List Entries can be imported using the prefix_list_id and cidr separated by a ,
	"aws_ec2_managed_prefix_list_entry": FormattedIdentifierFromProvider(",", "prefix_list_id", "cidr"),
	// Prefix Lists can be imported using the id
	"aws_ec2_managed_prefix_list": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_prefix_list_reference can be imported by using the EC2 Transit Gateway Route Table identifier and EC2 Prefix List identifier, separated by an underscore (_
	"aws_ec2_transit_gateway_prefix_list_reference": FormattedIdentifierFromProvider("_", "transit_gateway_route_table_id", "prefix_list_id"),
	// Egress-only Internet gateways can be imported using the id
	"aws_egress_only_internet_gateway": config.IdentifierFromProvider,
	// EIP Assocations can be imported using their association ID.
	"aws_eip_association": config.IdentifierFromProvider,
	// Flow Logs can be imported using the id
	"aws_flow_log": config.IdentifierFromProvider,
	// Key Pairs can be imported using the key_name
	"aws_key_pair": config.ParameterAsIdentifier("key_name"),
	// Network ACLs can be imported using the id
	"aws_network_acl": config.IdentifierFromProvider,
	// No import
	"aws_network_interface_attachment": config.IdentifierFromProvider,
	// No import
	"aws_network_interface_sg_attachment": config.IdentifierFromProvider,
	// Individual rules can be imported using NETWORK_ACL_ID:RULE_NUMBER:PROTOCOL:EGRESS
	"aws_network_acl_rule": config.IdentifierFromProvider,
	// No import
	"aws_spot_instance_request": config.IdentifierFromProvider,
	// Spot Fleet Requests can be imported using id
	"aws_spot_fleet_request": config.IdentifierFromProvider,
	// EBS Volume Attachments can be imported using DEVICE_NAME:VOLUME_ID:INSTANCE_ID
	"aws_volume_attachment": config.IdentifierFromProvider,
	// VPC DHCP Options can be imported using the dhcp options id
	"aws_vpc_dhcp_options": config.IdentifierFromProvider,
	// DHCP associations can be imported by providing the VPC ID associated with the options
	// terraform import aws_vpc_dhcp_options_association.imported vpc-0f001273ec18911b1
	"aws_vpc_dhcp_options_association": config.IdentifierFromProvider,
	// VPC Endpoint Services can be imported using the VPC endpoint service id
	"aws_vpc_endpoint_service": config.IdentifierFromProvider,
	// VPC Endpoint connection notifications can be imported using the VPC endpoint connection notification id
	"aws_vpc_endpoint_connection_notification": config.IdentifierFromProvider,
	// VPC Endpoint Route Table Associations can be imported using vpc_endpoint_id together with route_table_id
	"aws_vpc_endpoint_route_table_association": FormattedIdentifierFromProvider("/", "vpc_endpoint_id", "route_table_id"),
	// Placement groups can be imported using the name
	"aws_placement_group": config.NameAsIdentifier,
	// A Spot Datafeed Subscription can be imported using the word spot-datafeed-subscription
	"aws_spot_datafeed_subscription": config.IdentifierFromProvider,
	// No import
	"aws_vpc_endpoint_service_allowed_principal": config.IdentifierFromProvider,
	// VPC Endpoint Subnet Associations can be imported using vpc_endpoint_id together with subnet_id
	"aws_vpc_endpoint_subnet_association": FormattedIdentifierFromProvider("/", "vpc_endpoint_id", "subnet_id"),
	// Default VPC route tables can be imported using the vpc_id
	"aws_default_route_table": config.IdentifierFromProvider,
	// Hosts can be imported using the host id
	"aws_ec2_host": config.IdentifierFromProvider,
	// Default VPCs can be imported using the vpc id
	"aws_default_vpc": config.IdentifierFromProvider,
	// Subnets can be imported using the subnet id
	"aws_default_subnet": config.IdentifierFromProvider,
	// VPC DHCP Options can be imported using the dhcp options id
	"aws_default_vpc_dhcp_options": config.IdentifierFromProvider,
	// The EBS default KMS CMK can be imported with the KMS key ARN
	"aws_ebs_default_kms_key": config.IdentifierFromProvider,
	// Default EBS encryption state can be imported
	"aws_ebs_encryption_by_default": config.IdentifierFromProvider,
	// EC2 Availability Zone Groups can be imported using the group name
	"aws_ec2_availability_zone_group": config.ParameterAsIdentifier("group_name"),
	// Capacity Reservations can be imported using the id
	"aws_ec2_capacity_reservation": config.IdentifierFromProvider,
	// aws_ec2_carrier_gateway can be imported using the carrier gateway's ID
	"aws_ec2_carrier_gateway": config.IdentifierFromProvider,
	// Serial console access state can be imported
	"aws_ec2_serial_console_access": config.IdentifierFromProvider,
	// Existing CIDR reservations can be imported using SUBNET_ID:RESERVATION_ID
	"aws_ec2_subnet_cidr_reservation": config.IdentifierFromProvider,
	// Traffic mirror filter can be imported using the id
	"aws_ec2_traffic_mirror_filter": config.IdentifierFromProvider,
	// Traffic mirror rules can be imported using the traffic_mirror_filter_id and id separated by :
	"aws_ec2_traffic_mirror_filter_rule": config.IdentifierFromProvider,
	// Traffic mirror targets can be imported using the id
	"aws_ec2_transit_gateway_connect": config.IdentifierFromProvider,
	// Network Insights Paths can be imported using the id
	"aws_ec2_network_insights_path": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_peering_attachment_accepter can be imported by using the EC2 Transit Gateway Attachment identifier
	"aws_ec2_transit_gateway_peering_attachment_accepter": config.IdentifierFromProvider,
	// No import
	"aws_snapshot_create_volume_permission": config.IdentifierFromProvider,
	// Customer Gateways can be imported using the id
	"aws_customer_gateway": config.IdentifierFromProvider,
	// Default Network ACLs can be imported using the id
	"aws_default_network_acl": config.IdentifierFromProvider,
	// IPAMs can be imported using the ipam pool id
	"aws_vpc_ipam_pool": config.IdentifierFromProvider,
	// IPAMs can be imported using the scope_id
	"aws_vpc_ipam_scope": config.IdentifierFromProvider,
	// IPAMs can be imported using the ipam id
	"aws_vpc_ipam": config.IdentifierFromProvider,
	// IPAMs can be imported using the ipam id
	"aws_vpc_ipam_pool_cidr": config.IdentifierFromProvider,
	// IPAMs can be imported using the allocation id
	"aws_vpc_ipam_pool_cidr_allocation": config.IdentifierFromProvider,
	// aws_ami can be imported using the ID of the AMI
	"aws_ami": config.IdentifierFromProvider,
	// No import
	"aws_ami_copy": config.IdentifierFromProvider,
	// AMI Launch Permissions can be imported using [ACCOUNT-ID|GROUP-NAME|ORGANIZATION-ARN|ORGANIZATIONAL-UNIT-ARN]/IMAGE-ID
	"aws_ami_launch_permission": config.IdentifierFromProvider,
	// VPN Connections can be imported using the vpn connection id
	"aws_vpn_connection": config.IdentifierFromProvider,
	// No import
	"aws_vpn_connection_route": config.IdentifierFromProvider,
	// VPN Gateways can be imported using the vpn gateway id
	"aws_vpn_gateway": config.IdentifierFromProvider,
	// No import
	"aws_vpn_gateway_attachment": config.IdentifierFromProvider,
	// No import
	"aws_vpn_gateway_route_propagation": config.IdentifierFromProvider,
	// Security Groups can be imported using the security group id
	"aws_default_security_group": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_connect_peer can be imported by using the EC2 Transit Gateway Connect Peer identifier
	"aws_ec2_transit_gateway_connect_peer": config.IdentifierFromProvider,

	// ecr
	//
	"aws_ecr_repository": config.NameAsIdentifier,
	// Imported using the name of the repository.
	"aws_ecr_lifecycle_policy": config.IdentifierFromProvider,
	// Use the ecr_repository_prefix to import a Pull Through Cache Rule.
	"aws_ecr_pull_through_cache_rule": config.IdentifierFromProvider,
	// Imported using the registry id, which is not a parameter at all.
	"aws_ecr_registry_policy": config.IdentifierFromProvider,
	// Imported using the registry id, which is not a parameter at all.
	"aws_ecr_registry_scanning_configuration": config.IdentifierFromProvider,
	// Imported using the registry id, which is not a parameter at all.
	"aws_ecr_replication_configuration": config.IdentifierFromProvider,
	// Imported using the parameter called repository but this is not the name
	// of the resource, only a configuration/reference.
	"aws_ecr_repository_policy": config.IdentifierFromProvider,

	// ecrpublic
	//
	"aws_ecrpublic_repository": config.ParameterAsIdentifier("repository_name"),
	// Imported using the repository name.
	"aws_ecrpublic_repository_policy": config.IdentifierFromProvider,

	// ecs
	//
	"aws_ecs_cluster":           config.NameAsIdentifier,
	"aws_ecs_service":           config.NameAsIdentifier,
	"aws_ecs_capacity_provider": config.TemplatedStringAsIdentifier("name", "arn:aws:ecs:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:capacity-provider/{{ .external_name }}"),
	// Imported using ARN that has a random substring, revision at the end:
	// arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123
	"aws_ecs_task_definition": config.IdentifierFromProvider,
	// ECS Account Setting defaults can be imported using the name
	"aws_ecs_account_setting_default": config.IdentifierFromProvider,
	// ECS cluster capacity providers can be imported using the cluster_name attribute
	"aws_ecs_cluster_capacity_providers": config.IdentifierFromProvider,

	// eks
	//
	"aws_eks_cluster": config.NameAsIdentifier,
	// Imported using the cluster_name and node_group_name separated by a
	// colon (:): my_cluster:my_node_group
	"aws_eks_node_group": config.TemplatedStringAsIdentifier("node_group_name", "{{ .parameters.cluster_name }}:{{ .external_name }}"),
	// my_cluster:my_eks_addon
	"aws_eks_addon": FormattedIdentifierUserDefinedNameLast("addon_name", ":", "cluster_name"),
	// my_cluster:my_fargate_profile
	"aws_eks_fargate_profile": FormattedIdentifierUserDefinedNameLast("fargate_profile_name", ":", "cluster_name"),
	// It has a complex config, adding empty entry here just to enable it.
	"aws_eks_identity_provider_config": eksOIDCIdentityProvider(),

	// elasticache
	//
	"aws_elasticache_parameter_group":   config.IdentifierFromProvider,
	"aws_elasticache_subnet_group":      config.NameAsIdentifier,
	"aws_elasticache_cluster":           config.ParameterAsIdentifier("cluster_id"),
	"aws_elasticache_replication_group": config.ParameterAsIdentifier("replication_group_id"),
	"aws_elasticache_user":              config.ParameterAsIdentifier("user_id"),
	"aws_elasticache_user_group":        config.ParameterAsIdentifier("user_group_id"),

	// elasticloadbalancing
	//
	// arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188
	"aws_lb": config.IdentifierFromProvider,
	// arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
	"aws_lb_listener": config.IdentifierFromProvider,
	// arn:aws:elasticloadbalancing:us-west-2:187416307283:targetgroup/app-front-end/20cfe21448b66314
	"aws_lb_target_group": config.IdentifierFromProvider,
	// No import.
	"aws_lb_target_group_attachment": config.IdentifierFromProvider,

	// globalaccelerator
	//
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	"aws_globalaccelerator_accelerator": config.IdentifierFromProvider,
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx
	"aws_globalaccelerator_endpoint_group": config.IdentifierFromProvider,
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx
	"aws_globalaccelerator_listener": config.IdentifierFromProvider,

	// glue
	//
	// Imported using "name".
	"aws_glue_workflow": config.NameAsIdentifier,
	// Imported using arn: arn:aws:glue:us-west-2:123456789012:schema/example/example
	// "aws_glue_schema": config.IdentifierFromProvider,
	// Imported using "name".
	"aws_glue_trigger":               config.NameAsIdentifier,
	"aws_glue_user_defined_function": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .parameters.database_name }}:{{ .external_name }}"),
	// "aws_glue_security_configuration": config.NameAsIdentifier,
	// Imported using the account ID: 12356789012
	"aws_glue_resource_policy":  config.IdentifierFromProvider,
	"aws_glue_catalog_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .external_name }}"),
	"aws_glue_catalog_table":    config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .parameters.database_name }}:{{ .external_name }}"),
	"aws_glue_classifier":       config.NameAsIdentifier,
	// Imported as CATALOG_ID:name 123456789012:MyConnection
	"aws_glue_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .external_name }}"),
	"aws_glue_crawler":    config.NameAsIdentifier,
	// Imported using CATALOG-ID (AWS account ID if not custom), e.g., 123456789012
	"aws_glue_data_catalog_encryption_settings": config.IdentifierFromProvider,
	// "aws_glue_dev_endpoint":                     config.NameAsIdentifier,
	"aws_glue_job": config.NameAsIdentifier,
	// Imported using id, e.g., tfm-c2cafbe83b1c575f49eaca9939220e2fcd58e2d5
	// "aws_glue_ml_transform": config.IdentifierFromProvider,
	// It has no naming argument, imported with their catalog ID (usually
	// AWS account ID), database name, table name and partition values e.g.,
	// 123456789012:MyDatabase:MyTable:val1#val2
	// "aws_glue_partition": config.IdentifierFromProvider,
	// Documentation does not match schema where there are multiple indexes
	// each with their own name.
	// "aws_glue_partition_index": config.IdentifierFromProvider,
	// Imported using ARN: arn:aws:glue:us-west-2:123456789012:registry/example
	"aws_glue_registry": config.TemplatedStringAsIdentifier("registry_name", "arn:aws:glue:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:registry/{{ .external_name }}"),

	// Imported using "name".
	"aws_glue_security_configuration": config.NameAsIdentifier,

	// iam
	//
	// AKIA1234567890
	"aws_iam_access_key":       config.IdentifierFromProvider,
	"aws_iam_instance_profile": config.NameAsIdentifier,
	// arn:aws:iam::123456789012:policy/UsersManageOwnCredentials
	"aws_iam_policy": config.TemplatedStringAsIdentifier("name", "arn:aws:iam::{{ .setup.client_metadata.account_id }}:policy/{{ .external_name }}"),
	"aws_iam_user":   config.NameAsIdentifier,
	"aws_iam_group":  config.NameAsIdentifier,
	"aws_iam_role":   config.NameAsIdentifier,
	// Imported using the role name and policy arn separated by /
	// test-role/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_role_policy_attachment": config.IdentifierFromProvider,
	// Imported using the user name and policy arn separated by /
	// test-user/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_user_policy_attachment": config.IdentifierFromProvider,
	// Imported using the group name and policy arn separated by /
	// test-group/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_group_policy_attachment": config.IdentifierFromProvider,
	// Imported using the user name and group names separated by /
	// user1/group1/group2
	"aws_iam_user_group_membership": iamUserGroupMembership(),
	// arn:aws:iam::123456789012:oidc-provider/accounts.google.com
	"aws_iam_openid_connect_provider": config.IdentifierFromProvider,
	// The current Account Alias can be imported using the account_alias
	"aws_iam_account_alias": config.ParameterAsIdentifier("account_alias"),
	// IAM Account Password Policy can be imported using the word iam-account-password-policy
	"aws_iam_account_password_policy": config.IdentifierFromProvider,
	// No import
	"aws_iam_group_membership": config.IdentifierFromProvider,
	// IAM SAML Providers can be imported using the arn
	"aws_iam_saml_provider": config.TemplatedStringAsIdentifier("name", "arn:aws:iam::{{ .setup.client_metadata.account_id }}:saml-provider/{{ .external_name }}"),
	// IAM Server Certificates can be imported using the name
	"aws_iam_server_certificate": config.NameAsIdentifier,
	// IAM service-linked roles can be imported using role ARN that contains the
	// service name.
	"aws_iam_service_linked_role": config.IdentifierFromProvider,
	// IAM Service Specific Credentials can be imported using the service_name:user_name:service_specific_credential_id
	"aws_iam_service_specific_credential": config.IdentifierFromProvider,
	// IAM Signing Certificates can be imported using the id
	"aws_iam_signing_certificate": config.IdentifierFromProvider,
	// IAM User Login Profiles can be imported without password information support via the IAM User name
	"aws_iam_user_login_profile": config.IdentifierFromProvider,
	// SSH public keys can be imported using the username, ssh_public_key_id, and encoding
	"aws_iam_user_ssh_key": config.IdentifierFromProvider,
	// IAM Virtual MFA Devices can be imported using the arn
	"aws_iam_virtual_mfa_device": config.IdentifierFromProvider,

	// kms
	//
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	"aws_kms_key": config.IdentifierFromProvider,
	// KMS aliases are imported using "alias/" + name
	"aws_kms_alias": kmsAlias(),
	// No import
	"aws_kms_ciphertext": config.IdentifierFromProvider,
	// KMS External Keys can be imported using the id
	"aws_kms_external_key": config.IdentifierFromProvider,
	// KMS Grants can be imported using the Key ID and Grant ID separated by a colon (:)
	"aws_kms_grant": config.IdentifierFromProvider,
	// KMS multi-Region replica keys can be imported using the id
	"aws_kms_replica_external_key": config.IdentifierFromProvider,
	// KMS multi-Region replica keys can be imported using the id
	"aws_kms_replica_key": config.IdentifierFromProvider,

	// mq
	//
	// a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
	"aws_mq_broker": config.IdentifierFromProvider,
	// c-0187d1eb-88c8-475a-9b79-16ef5a10c94f
	"aws_mq_configuration": config.IdentifierFromProvider,

	// neptune
	//
	"aws_neptune_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// my_cluster:my_cluster_endpoint
	"aws_neptune_cluster_endpoint":        FormattedIdentifierUserDefinedNameLast("cluster_endpoint_identifier", ":", "cluster_identifier"),
	"aws_neptune_cluster_instance":        config.ParameterAsIdentifier("identifier"),
	"aws_neptune_cluster_parameter_group": config.NameAsIdentifier,
	"aws_neptune_cluster_snapshot":        config.ParameterAsIdentifier("db_cluster_snapshot_identifier"),
	"aws_neptune_event_subscription":      config.NameAsIdentifier,
	"aws_neptune_parameter_group":         config.NameAsIdentifier,
	"aws_neptune_subnet_group":            config.NameAsIdentifier,

	// rds
	//
	"aws_rds_cluster":        config.ParameterAsIdentifier("cluster_identifier"),
	"aws_db_instance":        config.ParameterAsIdentifier("identifier"),
	"aws_db_parameter_group": config.NameAsIdentifier,
	"aws_db_subnet_group":    config.NameAsIdentifier,
	// aws_db_instance_role_association can be imported using the DB Instance Identifier and IAM Role ARN separated by a comma
	// $ terraform import aws_db_instance_role_association.example my-db-instance,arn:aws:iam::123456789012:role/my-role
	"aws_db_instance_role_association": config.IdentifierFromProvider,
	// DB Option groups can be imported using the name
	"aws_db_option_group": config.NameAsIdentifier,
	// DB proxies can be imported using the name
	"aws_db_proxy": config.NameAsIdentifier,
	// DB proxy default target groups can be imported using the db_proxy_name
	"aws_db_proxy_default_target_group": config.IdentifierFromProvider,
	// DB proxy endpoints can be imported using the DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME
	"aws_db_proxy_endpoint": config.TemplatedStringAsIdentifier("db_proxy_endpoint_name", "{{ .external_name }}/{{ .parameters.db_proxy_name }}"),
	// RDS DB Proxy Targets can be imported using the db_proxy_name, target_group_name, target type (e.g., RDS_INSTANCE or TRACKED_CLUSTER), and resource identifier separated by forward slashes (/)
	"aws_db_proxy_target": config.IdentifierFromProvider,
	// DB Security groups can be imported using the name
	"aws_db_security_group": config.NameAsIdentifier,
	// aws_db_snapshot can be imported by using the snapshot identifier
	"aws_db_snapshot": config.ParameterAsIdentifier("db_snapshot_identifier"),
	// RDS Aurora Cluster Database Activity Streams can be imported using the resource_arn
	"aws_rds_cluster_activity_stream": config.IdentifierFromProvider,
	// RDS Clusters Endpoint can be imported using the cluster_endpoint_identifier
	"aws_rds_cluster_endpoint": config.ParameterAsIdentifier("cluster_endpoint_identifier"),
	// RDS Cluster Instances can be imported using the identifier
	"aws_rds_cluster_instance": config.ParameterAsIdentifier("identifier"),
	// RDS Cluster Parameter Groups can be imported using the name
	"aws_rds_cluster_parameter_group": config.NameAsIdentifier,
	// aws_rds_cluster_role_association can be imported using the DB Cluster Identifier and IAM Role ARN separated by a comma (,)
	// $ terraform import aws_rds_cluster_role_association.example my-db-cluster,arn:aws:iam::123456789012:role/my-role
	"aws_rds_cluster_role_association": FormattedIdentifierFromProvider(",", "db_cluster_identifier", "role_arn"),
	// aws_rds_global_cluster can be imported by using the RDS Global Cluster identifie
	"aws_rds_global_cluster": config.ParameterAsIdentifier("global_cluster_identifier"),
	// aws_db_cluster_snapshot can be imported by using the cluster snapshot identifier
	"aws_db_cluster_snapshot": config.IdentifierFromProvider,
	// DB Event Subscriptions can be imported using the name
	"aws_db_event_subscription": config.NameAsIdentifier,
	// RDS instance automated backups replication can be imported using the arn
	"aws_db_instance_automated_backups_replication": config.IdentifierFromProvider,
	// aws_db_snapshot_copy can be imported by using the snapshot identifier
	"aws_db_snapshot_copy": config.IdentifierFromProvider,

	// route53
	//
	// N1PA6795SAMPLE
	"aws_route53_delegation_set": config.IdentifierFromProvider,
	// abcdef11-2222-3333-4444-555555fedcba
	"aws_route53_health_check": config.IdentifierFromProvider,
	// Z1D633PJN98FT9
	"aws_route53_hosted_zone_dnssec": config.IdentifierFromProvider,
	// Imported by using the Route 53 Hosted Zone identifier and KMS Key
	// identifier, separated by a comma (,), e.g., Z1D633PJN98FT9,example
	// disabled until it's successfully tested
	// "aws_route53_key_signing_key": FormattedIdentifierUserDefinedNameLast("name", ",", "hosted_zone_id"),
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	// disabled until it's successfully tested
	// "aws_route53_query_log": config.IdentifierFromProvider,
	// Imported using ID of the record, which is the zone identifier, record
	// name, and record type, separated by underscores (_)
	// Z4KAPRWWNC7JR_dev.example.com_NS
	"aws_route53_record": config.IdentifierFromProvider,
	// Z123456ABCDEFG:vpc-12345678
	"aws_route53_vpc_association_authorization": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	// Z1D633PJN98FT9
	"aws_route53_zone": config.IdentifierFromProvider,
	// Z123456ABCDEFG:vpc-12345678
	// disabled until it's successfully tested
	// "aws_route53_zone_association": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	// Imported using the id and version, e.g.,
	// 01a52019-d16f-422a-ae72-c306d2b6df7e/1
	"aws_route53_traffic_policy": config.IdentifierFromProvider,
	// df579d9a-6396-410e-ac22-e7ad60cf9e7e
	"aws_route53_traffic_policy_instance": config.IdentifierFromProvider,

	// route53resolver
	//
	// rdsc-be1866ecc1683e95
	// disabled until it's successfully tested
	// "aws_route53_resolver_dnssec_config": config.IdentifierFromProvider,
	// rslvr-in-abcdef01234567890
	"aws_route53_resolver_endpoint": config.IdentifierFromProvider,
	// rdsc-be1866ecc1683e95
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_config": config.IdentifierFromProvider,
	// rslvr-fdl-0123456789abcdef
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_domain_list": config.IdentifierFromProvider,
	// Imported using the Route 53 Resolver DNS Firewall rule group ID and
	// domain list ID separated by ':', e.g.,
	// rslvr-frg-0123456789abcdef:rslvr-fdl-0123456789abcdef
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_rule": config.IdentifierFromProvider,
	// rslvr-frg-0123456789abcdef
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_rule_group": config.IdentifierFromProvider,
	// rslvr-frgassoc-0123456789abcdef
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_rule_group_association": config.IdentifierFromProvider,
	// rqlc-92edc3b1838248bf
	// disabled until it's successfully tested
	// "aws_route53_resolver_query_log_config": config.IdentifierFromProvider,
	// rqlca-b320624fef3c4d70
	// disabled until it's successfully tested
	// "aws_route53_resolver_query_log_config_association": config.IdentifierFromProvider,
	// rslvr-rr-0123456789abcdef0
	"aws_route53_resolver_rule": config.IdentifierFromProvider,
	// rslvr-rrassoc-97242eaf88example
	"aws_route53_resolver_rule_association": config.IdentifierFromProvider,

	// s3
	//
	// S3 bucket can be imported using the bucket
	"aws_s3_bucket": config.ParameterAsIdentifier("bucket"),
	// the S3 bucket accelerate configuration resource should be imported using the bucket
	"aws_s3_bucket_object_lock_configuration": config.IdentifierFromProvider,
	// the S3 bucket accelerate configuration resource should be imported using the bucket
	"aws_s3_bucket_accelerate_configuration": config.IdentifierFromProvider,
	// the S3 bucket ACL resource should be imported using the bucket
	"aws_s3_bucket_acl": config.IdentifierFromProvider,
	// S3 bucket analytics configurations can be imported using bucket:analytics
	"aws_s3_bucket_analytics_configuration": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// The S3 bucket CORS configuration resource should be imported using the bucket
	"aws_s3_bucket_cors_configuration": config.IdentifierFromProvider,
	// S3 bucket intelligent tiering configurations can be imported using bucket:name
	// $ terraform import aws_s3_bucket_intelligent_tiering_configuration.my-bucket-entire-bucket my-bucket:EntireBucket
	"aws_s3_bucket_intelligent_tiering_configuration": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// S3 bucket inventory configurations can be imported using bucket:inventory
	// $ terraform import aws_s3_bucket_inventory.my-bucket-entire-bucket my-bucket:EntireBucket
	"aws_s3_bucket_inventory": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// The S3 bucket lifecycle configuration resource should be imported using the bucket
	"aws_s3_bucket_lifecycle_configuration": config.IdentifierFromProvider,
	// The S3 bucket logging resource should be imported using the bucket
	"aws_s3_bucket_logging": config.IdentifierFromProvider,
	// S3 bucket metric configurations can be imported using bucket:metric
	"aws_s3_bucket_metric": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// S3 bucket notification can be imported using the bucket
	"aws_s3_bucket_notification": config.IdentifierFromProvider,
	// Objects can be imported using the id. The id is the bucket name and the key together
	"aws_s3_bucket_object": config.IdentifierFromProvider,
	// S3 Bucket Ownership Controls can be imported using S3 Bucket name
	"aws_s3_bucket_ownership_controls": config.IdentifierFromProvider,
	// S3 bucket policies can be imported using the bucket name
	"aws_s3_bucket_policy": config.IdentifierFromProvider,
	// aws_s3_bucket_public_access_block can be imported by using the bucket name
	"aws_s3_bucket_public_access_block": config.IdentifierFromProvider,
	// S3 bucket replication configuration can be imported using the bucket
	"aws_s3_bucket_replication_configuration": config.IdentifierFromProvider,
	// The S3 bucket request payment configuration resource should be imported using the bucket
	"aws_s3_bucket_request_payment_configuration": config.IdentifierFromProvider,
	// The S3 server-side encryption configuration resource should be imported using the bucket
	"aws_s3_bucket_server_side_encryption_configuration": config.IdentifierFromProvider,
	// The S3 bucket versioning resource should be imported using the bucket
	"aws_s3_bucket_versioning": config.IdentifierFromProvider,
	// The S3 bucket website configuration resource should be imported using the bucket
	"aws_s3_bucket_website_configuration": config.IdentifierFromProvider,
	// Objects can be imported using the id. The id is the bucket name and the key together
	// $ terraform import aws_s3_object.object some-bucket-name/some/key.txt
	"aws_s3_object": FormattedIdentifierFromProvider("/", "bucket", "key"),

	// cloudfront
	//
	// Cloudfront Cache Policies can be imported using the id
	"aws_cloudfront_cache_policy": config.IdentifierFromProvider,
	// Cloudfront Distributions can be imported using the id
	"aws_cloudfront_distribution": config.IdentifierFromProvider,
	// Cloudfront Field Level Encryption Config can be imported using the id
	"aws_cloudfront_field_level_encryption_config": config.IdentifierFromProvider,
	// Cloudfront Field Level Encryption Profile can be imported using the id
	"aws_cloudfront_field_level_encryption_profile": config.IdentifierFromProvider,
	// CloudFront Functions can be imported using the name
	"aws_cloudfront_function": config.NameAsIdentifier,
	// CloudFront Key Group can be imported using the id
	"aws_cloudfront_key_group": config.IdentifierFromProvider,
	// CloudFront monitoring subscription can be imported using the id
	"aws_cloudfront_monitoring_subscription": config.IdentifierFromProvider,
	// Cloudfront Origin Access Identities can be imported using the id
	"aws_cloudfront_origin_access_identity": config.IdentifierFromProvider,
	// No import documented, but https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_origin_request_policy#name
	"aws_cloudfront_origin_request_policy": config.NameAsIdentifier,
	// CloudFront Public Key can be imported using the id
	"aws_cloudfront_public_key": config.IdentifierFromProvider,
	// CloudFront real-time log configurations can be imported using the ARN,
	// $ terraform import aws_cloudfront_realtime_log_config.example arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig
	"aws_cloudfront_realtime_log_config": config.IdentifierFromProvider,
	// Cloudfront Response Headers Policies can be imported using the id
	"aws_cloudfront_response_headers_policy": config.IdentifierFromProvider,

	// resource groups

	// Resource groups can be imported using the name
	"aws_resourcegroups_group": config.NameAsIdentifier,

	// docdb
	//
	// DocDB Clusters can be imported using the cluster_identifier
	"aws_docdb_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// aws_docdb_global_cluster can be imported by using the Global Cluster id
	"aws_docdb_global_cluster": config.IdentifierFromProvider,
	// DocDB Cluster Instances can be imported using the identifier
	"aws_docdb_cluster_instance": config.ParameterAsIdentifier("identifier"),
	// DocumentDB Subnet groups can be imported using the name
	"aws_docdb_subnet_group": config.NameAsIdentifier,
	// DocumentDB Cluster Parameter Groups can be imported using the name
	"aws_docdb_cluster_parameter_group": config.NameAsIdentifier,
	// aws_docdb_cluster_snapshot can be imported by using the cluster snapshot identifier
	"aws_docdb_cluster_snapshot": config.ParameterAsIdentifier("db_cluster_snapshot_identifier"),
	// DocDB Event Subscriptions can be imported using the name
	"aws_docdb_event_subscription": config.NameAsIdentifier,

	// efs
	//
	// The EFS file systems can be imported using the id
	"aws_efs_file_system": config.IdentifierFromProvider,
	// The EFS mount targets can be imported using the id
	"aws_efs_mount_target": config.IdentifierFromProvider,
	// The EFS access points can be imported using the id
	"aws_efs_access_point": config.IdentifierFromProvider,
	// The EFS backup policies can be imported using the id
	"aws_efs_backup_policy": config.IdentifierFromProvider,
	// The EFS file system policies can be imported using the id
	"aws_efs_file_system_policy": config.IdentifierFromProvider,

	// servicediscovery
	//
	// Service Discovery Private DNS Namespace can be imported using the namespace ID and VPC ID: 0123456789:vpc-123345
	"aws_service_discovery_private_dns_namespace": config.IdentifierFromProvider,
	// Service Discovery Public DNS Namespace can be imported using the namespace ID
	"aws_service_discovery_public_dns_namespace": config.IdentifierFromProvider,
	// Service Discovery HTTP Namespace can be imported using the namespace ID,
	"aws_service_discovery_http_namespace": config.IdentifierFromProvider,
	// Service Discovery Service can be imported using the service ID
	"aws_service_discovery_service": config.IdentifierFromProvider,

	// sqs
	//
	// SQS Queues can be imported using the queue url / id
	"aws_sqs_queue": config.IdentifierFromProvider,
	// SQS Queue Policies can be imported using the queue URL
	// e.g. https://queue.amazonaws.com/0123456789012/myqueue
	"aws_sqs_queue_policy": config.IdentifierFromProvider,

	// secretsmanager
	//
	// It be imported by using the secret Amazon Resource Name (ARN)
	// However, the real ID of the Secret has an Amazon-assigned random suffix,
	// i.e. if you name it with `example`, the real ID is
	// arn:aws:secretsmanager:us-west-1:609897127049:secret:example-VaznFM
	"aws_secretsmanager_secret": config.IdentifierFromProvider,
	// It uses ARN of secret and a randomly assigned ID.
	"aws_secretsmanager_secret_version": config.IdentifierFromProvider,
	// It uses its own secret_id parameter.
	"aws_secretsmanager_secret_rotation": config.IdentifierFromProvider,
	// It uses its own secert_arn parameter.
	"aws_secretsmanager_secret_policy": config.IdentifierFromProvider,

	// transfer
	//
	// Transfer Servers can be imported using the id
	"aws_transfer_server": config.IdentifierFromProvider,
	// Transfer Users can be imported using the server_id and user_name separated by /
	"aws_transfer_user": FormattedIdentifierUserDefinedNameLast("user_name", "/", "server_id"),

	// dynamodb
	//
	// DynamoDB tables can be imported using the name
	"aws_dynamodb_table": config.NameAsIdentifier,
	// DynamoDB Global Tables can be imported using the global table name
	"aws_dynamodb_global_table": config.NameAsIdentifier,
	// aws_dynamodb_tag can be imported by using the DynamoDB resource identifier and key, separated by a comma (,)
	"aws_dynamodb_tag": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_arn }},{{ .parameters.key }}"),

	// sns
	//
	// SNS Topics can be imported using the topic arn
	"aws_sns_topic": config.TemplatedStringAsIdentifier("name", "arn:aws:sns:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:{{ .external_name }}"),
	// SNS Topic Subscriptions can be imported using the subscription arn that
	// contains a random substring in the end.
	"aws_sns_topic_subscription": config.IdentifierFromProvider,

	// backup
	//
	// Backup Framework can be imported using the id which corresponds to the name of the Backup Framework
	"aws_backup_framework": config.IdentifierFromProvider,
	// Backup Global Settings can be imported using the id
	"aws_backup_global_settings": config.IdentifierFromProvider,
	// Backup Plan can be imported using the id
	"aws_backup_plan": config.IdentifierFromProvider,
	// Backup vault can be imported using the name
	"aws_backup_vault": config.NameAsIdentifier,
	// Backup Region Settings can be imported using the region
	"aws_backup_region_settings": config.IdentifierFromProvider,
	// Backup Report Plan can be imported using the id which corresponds to the name of the Backup Report Plan
	"aws_backup_report_plan": config.IdentifierFromProvider,
	// Backup selection can be imported using the role plan_id and id separated by | plan-id|selection-id
	"aws_backup_selection": config.IdentifierFromProvider,
	// Backup vault lock configuration can be imported using the name of the backup vault
	"aws_backup_vault_lock_configuration": config.IdentifierFromProvider,
	// Backup vault notifications can be imported using the name of the backup vault
	"aws_backup_vault_notifications": config.IdentifierFromProvider,
	// Backup vault policy can be imported using the name of the backup vault
	"aws_backup_vault_policy": config.IdentifierFromProvider,

	// grafana
	//
	// Grafana Workspace can be imported using the workspace's id
	"aws_grafana_workspace": config.IdentifierFromProvider,
	// No import
	"aws_grafana_role_association": config.IdentifierFromProvider,
	// Grafana Workspace SAML configuration can be imported using the workspace's id
	"aws_grafana_workspace_saml_configuration": FormattedIdentifierFromProvider("", "workspace_id"),

	// gamelift
	//
	// GameLift Aliases can be imported using the ID
	"aws_gamelift_alias": config.IdentifierFromProvider,
	// GameLift Builds can be imported using the ID
	"aws_gamelift_build": config.IdentifierFromProvider,
	// GameLift Fleets can be imported using the ID
	"aws_gamelift_fleet": config.IdentifierFromProvider,
	// GameLift Game Session Queues can be imported by their name
	"aws_gamelift_game_session_queue": config.NameAsIdentifier,
	// GameLift Scripts can be imported using the ID
	"aws_gamelift_script": config.IdentifierFromProvider,

	// kinesis
	//
	// Even though the documentation says the ID is name, it uses ARN..
	"aws_kinesis_stream": config.TemplatedStringAsIdentifier("name", " arn:aws:kinesis:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:stream/{{ .external_name }}"),
	// Kinesis Stream Consumers can be imported using the Amazon Resource Name (ARN)
	// that has a random substring.
	"aws_kinesis_stream_consumer": config.IdentifierFromProvider,

	// kinesisanalytics
	//
	"aws_kinesis_analytics_application": config.TemplatedStringAsIdentifier("name", "arn:aws:kinesisanalytics:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:application/{{ .external_name }}"),

	// kinesisanalyticsv2
	//
	"aws_kinesisanalyticsv2_application": config.TemplatedStringAsIdentifier("name", "arn:aws:kinesisanalytics:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:application/{{ .external_name }}"),
	// aws_kinesisanalyticsv2_application can be imported by using application_name together with snapshot_name
	// e.g. example-application/example-snapshot
	"aws_kinesisanalyticsv2_application_snapshot": FormattedIdentifierUserDefinedNameLast("snapshot_name", "/", "application_name"),

	// kinesisvideo
	//
	// Kinesis Streams can be imported using the arn that has a random substring
	// in the end.
	// arn:aws:kinesisvideo:us-west-2:123456789012:stream/terraform-kinesis-test/1554978910975
	"aws_kinesis_video_stream": config.IdentifierFromProvider,

	// firehose
	//
	"aws_kinesis_firehose_delivery_stream": config.TemplatedStringAsIdentifier("name", "arn:aws:firehose:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:deliverystream/{{ .external_name }}"),

	// lakeformation
	//
	// No import
	"aws_lakeformation_data_lake_settings": config.IdentifierFromProvider,
	// No import
	"aws_lakeformation_permissions": config.IdentifierFromProvider,
	// No import
	"aws_lakeformation_resource": config.IdentifierFromProvider,

	// lexmodels
	//
	// Bots can be imported using their name.
	"aws_lex_bot": config.NameAsIdentifier,
	// Bot aliases can be imported using an ID with the format bot_name:bot_alias_name
	"aws_lex_bot_alias": FormattedIdentifierUserDefinedNameLast("name", ":", "bot_name"),
	// Intents can be imported using their name.
	"aws_lex_intent": config.NameAsIdentifier,
	// Slot types can be imported using their name.
	"aws_lex_slot_type": config.NameAsIdentifier,

	// licensemanager
	//
	// License configurations can be imported in the form resource_arn,license_configuration_arn
	"aws_licensemanager_association": FormattedIdentifierFromProvider(",", "resource_arn", "license_configuration_arn"),
	// License configurations can be imported using the id
	"aws_licensemanager_license_configuration": config.IdentifierFromProvider,

	// lambda
	//
	// Lambda Function Aliases can be imported using the function_name/alias
	"aws_lambda_alias": config.TemplatedStringAsIdentifier("name", "{{ .parameters.function_name }}/{{ .external_name }}"),
	// Code Signing Configs can be imported using their ARN that has a random
	// substring in the end.
	// arn:aws:lambda:us-west-2:123456789012:code-signing-config:csc-0f6c334abcdea4d8b
	"aws_lambda_code_signing_config": config.IdentifierFromProvider,
	// Lambda event source mappings can be imported using the UUID (event source mapping identifier)
	"aws_lambda_event_source_mapping": config.IdentifierFromProvider,
	// Lambda Functions can be imported using the function_name
	"aws_lambda_function": config.ParameterAsIdentifier("function_name"),
	// Lambda Function Event Invoke Configs can be imported using the
	// fully qualified Function name or Amazon Resource Name (ARN) of the function.
	"aws_lambda_function_event_invoke_config": config.IdentifierFromProvider,
	// Lambda function URLs can be imported using the function_name or function_name/qualifier
	"aws_lambda_function_url": lambdaFunctionURL(),
	// No import"
	"aws_lambda_invocation": config.IdentifierFromProvider,
	// Lambda Layers can be imported using arn that has an assigned version in the
	// end
	"aws_lambda_layer_version": config.IdentifierFromProvider,
	// Lambda Layer Permissions can be imported using layer_name and version_number, separated by a comma (,)
	"aws_lambda_layer_version_permission": config.IdentifierFromProvider,
	// Lambda permission statements can be imported using function_name/statement_id, with an optional qualifier
	"aws_lambda_permission": config.IdentifierFromProvider,
	// Lambda Provisioned Concurrency Configs can be imported using the function_name and qualifier separated by a colon (:)
	"aws_lambda_provisioned_concurrency_config": config.IdentifierFromProvider,

	// signer
	//
	// Signer signing profiles can be imported using the name
	"aws_signer_signing_profile": config.NameAsIdentifier,

	// athena
	//
	// Athena Workgroups can be imported using their name
	"aws_athena_workgroup": config.NameAsIdentifier,
	// Data catalogs can be imported using their name
	"aws_athena_data_catalog": config.NameAsIdentifier,
	// Athena Databases can be imported using their name
	"aws_athena_database": config.NameAsIdentifier,
	// Athena Named Query can be imported using the query ID
	"aws_athena_named_query": config.IdentifierFromProvider,

	// cloudwatchlogs
	//
	// Cloudwatch Log Groups can be imported using the name
	"aws_cloudwatch_log_group": config.NameAsIdentifier,
	// CloudWatch Log Metric Filter can be imported using the log_group_name:name
	"aws_cloudwatch_log_metric_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_group_name }}:{{ .external_name }}"),
	// CloudWatch query definitions can be imported using the query definition ARN.
	"aws_cloudwatch_query_definition": config.IdentifierFromProvider,
	// Cloudwatch Log Stream can be imported using the stream's log_group_name and name
	"aws_cloudwatch_log_stream": config.IdentifierFromProvider,
	// CloudWatch log resource policies can be imported using the policy name
	"aws_cloudwatch_log_resource_policy": config.ParameterAsIdentifier("policy_name"),
	// CloudWatch Logs destinations can be imported using the name
	"aws_cloudwatch_log_destination": config.NameAsIdentifier,
	// CloudWatch Logs destination policies can be imported using the destination_name
	"aws_cloudwatch_log_destination_policy": config.ParameterAsIdentifier("destination_name"),
	// CloudWatch Logs subscription filter can be imported using the log group name and subscription filter name separated by |
	"aws_cloudwatch_log_subscription_filter": config.IdentifierFromProvider,

	// elb
	//
	// ELBs can be imported using the name
	"aws_elb": config.NameAsIdentifier,
	// No import
	"aws_elb_attachment": config.IdentifierFromProvider,
	// Application cookie stickiness policies can be imported using the ELB name, port, and policy name separated by colons (:)
	// my-elb:80:my-policy
	"aws_app_cookie_stickiness_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.load_balancer }}:{{ .parameters.lb_port }}:{{ .external_name }}"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lb_cookie_stickiness_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lb_ssl_negotiation_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_load_balancer_backend_server_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_load_balancer_listener_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_load_balancer_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_proxy_protocol_policy": config.IdentifierFromProvider,

	// iot
	//
	// IoT policies can be imported using the name
	"aws_iot_policy": config.NameAsIdentifier,
	// IOT Things can be imported using the name
	"aws_iot_thing": config.NameAsIdentifier,

	// kafka
	//
	// MSK configurations can be imported using the configuration ARN that has
	// a random substring in the end.
	"aws_msk_configuration": config.IdentifierFromProvider,
	// MSK clusters can be imported using the cluster arn that has a random substring
	// in the end.
	"aws_msk_cluster": config.IdentifierFromProvider,

	// ram
	//
	// Resource shares can be imported using the id
	"aws_ram_resource_share": config.IdentifierFromProvider,

	// redshift
	//
	// Redshift Clusters can be imported using the cluster_identifier
	"aws_redshift_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// Redshift Event Subscriptions can be imported using the name
	"aws_redshift_event_subscription": config.NameAsIdentifier,
	// Redshift Parameter Groups can be imported using the name
	"aws_redshift_parameter_group": config.IdentifierFromProvider,
	// Redshift Scheduled Action can be imported using the name
	"aws_redshift_scheduled_action": config.NameAsIdentifier,
	// Redshift Snapshot Schedule can be imported using the identifier
	"aws_redshift_snapshot_schedule": config.ParameterAsIdentifier("identifier"),
	// Redshift Snapshot Schedule Association can be imported using the <cluster-identifier>/<schedule-identifier>
	"aws_redshift_snapshot_schedule_association": config.IdentifierFromProvider,
	// Redshift subnet groups can be imported using the name
	"aws_redshift_subnet_group": config.NameAsIdentifier,

	// sfn
	//
	"aws_sfn_activity":      config.TemplatedStringAsIdentifier("name", "arn:aws:states:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:activity/{{ .external_name }}"),
	"aws_sfn_state_machine": config.TemplatedStringAsIdentifier("name", "arn:aws:states:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:stateMachine/{{ .external_name }}"),

	// dax
	//
	// DAX Clusters can be imported using the cluster_name
	"aws_dax_cluster": config.ParameterAsIdentifier("cluster_name"),
	// DAX Parameter Group can be imported using the name
	"aws_dax_parameter_group": config.NameAsIdentifier,
	// DAX Subnet Group can be imported using the name
	"aws_dax_subnet_group": config.NameAsIdentifier,

	// cloudsearch
	//
	// CloudSearch Domains can be imported using the name
	"aws_cloudsearch_domain": config.NameAsIdentifier,
	// CloudSearch domain service access policies can be imported using the domain name
	"aws_cloudsearch_domain_service_access_policy": config.IdentifierFromProvider,

	// apigateway
	//
	// API Gateway Keys can be imported using the id
	"aws_api_gateway_api_key": config.IdentifierFromProvider,
	// API Gateway Client Certificates can be imported using the id
	"aws_api_gateway_client_certificate": config.IdentifierFromProvider,
	// aws_api_gateway_rest_api can be imported by using the REST API ID
	"aws_api_gateway_rest_api": config.IdentifierFromProvider,
	// API Gateway documentation_parts can be imported using REST-API-ID/DOC-PART-ID
	"aws_api_gateway_documentation_part": config.IdentifierFromProvider,
	// API Gateway documentation versions can be imported using REST-API-ID/VERSION
	"aws_api_gateway_documentation_version": FormattedIdentifierFromProvider("/", "rest_api_id", "version"),
	// aws_api_gateway_gateway_response can be imported using REST-API-ID/RESPONSE-TYPE
	"aws_api_gateway_gateway_response": FormattedIdentifierFromProvider("/", "rest_api_id", "response_type"),
	// aws_api_gateway_resource can be imported using REST-API-ID/RESOURCE-ID
	"aws_api_gateway_resource": config.IdentifierFromProvider,
	// aws_api_gateway_method can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD
	"aws_api_gateway_method": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method"),
	// aws_api_gateway_method_response can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE
	"aws_api_gateway_method_response": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method", "status_code"),
	// No import
	"aws_api_gateway_deployment": config.IdentifierFromProvider,
	// API Gateway Accounts can be imported using the word api-gateway-account
	"aws_api_gateway_account": config.IdentifierFromProvider,
	// aws_api_gateway_stage can be imported using REST-API-ID/STAGE-NAME
	"aws_api_gateway_stage": FormattedIdentifierFromProvider("/", "rest_api_id", "stage_name"),
	// aws_api_gateway_integration can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD
	"aws_api_gateway_integration": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method"),
	// aws_api_gateway_integration_response can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE
	"aws_api_gateway_integration_response": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method", "status_code"),
	// aws_api_gateway_method_settings can be imported using REST-API-ID/STAGE-NAME/METHOD-PATH
	"aws_api_gateway_method_settings": FormattedIdentifierFromProvider("/", "rest_api_id", "stage_name", "method_path"),
	// aws_api_gateway_model can be imported using REST-API-ID/NAME
	"aws_api_gateway_model": config.IdentifierFromProvider,
	// aws_api_gateway_request_validator can be imported using REST-API-ID/REQUEST-VALIDATOR-ID
	"aws_api_gateway_request_validator": config.IdentifierFromProvider,
	// aws_api_gateway_rest_api_policy can be imported by using the REST API ID
	"aws_api_gateway_rest_api_policy": FormattedIdentifierFromProvider("", "rest_api_id"),
	// AWS API Gateway Authorizer can be imported using the REST-API-ID/AUTHORIZER-ID
	"aws_api_gateway_authorizer": config.IdentifierFromProvider,
	// aws_api_gateway_base_path_mapping can be imported by using the domain name and base path.
	// For empty base_path (e.g., root path (/)): example.com/
	// Otherwise: example.com/base-path
	"aws_api_gateway_base_path_mapping": config.IdentifierFromProvider,
	// API Gateway domain names can be imported using their name
	"aws_api_gateway_domain_name": config.IdentifierFromProvider,
	// AWS API Gateway Usage Plan can be imported using the id
	"aws_api_gateway_usage_plan": config.IdentifierFromProvider,
	// AWS API Gateway Usage Plan Key can be imported using the USAGE-PLAN-ID/USAGE-PLAN-KEY-ID
	"aws_api_gateway_usage_plan_key": config.IdentifierFromProvider,
	// API Gateway VPC Link can be imported using the id
	"aws_api_gateway_vpc_link": config.IdentifierFromProvider,

	// opensearch
	//
	// NOTE(sergen): Parameter as identifier cannot be used, because terraform
	// overrides the id after terraform calls.
	// Please see the following issue in upjet: https://github.com/upbound/upjet/issues/32
	// OpenSearch domains can be imported using the domain_name
	"aws_opensearch_domain": config.IdentifierFromProvider,
	// No imports
	"aws_opensearch_domain_policy": config.IdentifierFromProvider,
	// NOTE(sergen): Parameter as identifier cannot be used, because terraform
	// overrides the id after terraform calls.
	// Please see the following issue in upjet: https://github.com/upbound/upjet/issues/32
	// OpenSearch domains can be imported using the domain_name
	"aws_opensearch_domain_saml_options": config.IdentifierFromProvider,

	// eventbridge
	//
	// Imported using name
	"aws_cloudwatch_event_api_destination": config.NameAsIdentifier,
	// Imported using name
	"aws_cloudwatch_event_archive": config.NameAsIdentifier,
	// Imported using name
	"aws_cloudwatch_event_bus": config.NameAsIdentifier,
	// Imported using event_bus_name
	"aws_cloudwatch_event_bus_policy": config.IdentifierFromProvider,
	// Imported using name
	"aws_cloudwatch_event_connection": config.NameAsIdentifier,
	// Imported using event_bus_name/statement_id
	"aws_cloudwatch_event_permission": FormattedIdentifierFromProvider("/", "event_bus_name", "statement_id"),
	// Imported using event_bus_name/rule_name
	"aws_cloudwatch_event_rule": FormattedIdentifierUserDefinedNameLast("name", "/", "event_bus_name"),
	// Imported using event_bus_name/rule_name/target_id
	"aws_cloudwatch_event_target": FormattedIdentifierFromProvider("/", "event_bus_name", "rule", "target_id"),

	// cloudwatch
	//
	// Use the alarm_name to import a CloudWatch Composite Alarm.
	"aws_cloudwatch_composite_alarm": config.ParameterAsIdentifier("alarm_name"),
	// CloudWatch dashboards can be imported using the dashboard_name
	"aws_cloudwatch_dashboard": config.ParameterAsIdentifier("dashboard_name"),
	// CloudWatch Metric Alarm can be imported using the alarm_name
	"aws_cloudwatch_metric_alarm": config.ParameterAsIdentifier("alarm_name"),
	// CloudWatch metric streams can be imported using the name
	"aws_cloudwatch_metric_stream": config.IdentifierFromProvider,

	// appautoscaling
	//
	// No import
	"aws_appautoscaling_scheduled_action": config.IdentifierFromProvider,
	// Application AutoScaling Policy can be imported using the service-namespace, resource-id, scalable-dimension and policy-name separated by /
	"aws_appautoscaling_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.service_namespace }}/{{ .parameters.resource_id }}/{{ .parameters.scalable_dimension }}/{{ .external_name }}"),
	// Application AutoScaling Target can be imported using the service-namespace , resource-id and scalable-dimension separated by /
	"aws_appautoscaling_target": TemplatedStringAsIdentifierWithNoName("{{ .parameters.service_namespace }}/{{ .parameters.resource_id }}/{{ .parameters.scalable_dimension }}"),

	// codecommit
	//
	// Codecommit repository can be imported using repository name
	"aws_codecommit_repository": config.ParameterAsIdentifier("repository_name"),
	// CodeCommit approval rule templates can be imported using the name
	"aws_codecommit_approval_rule_template": config.NameAsIdentifier,
	// CodeCommit approval rule template associations can be imported using the approval_rule_template_name and repository_name separated by a comma (,)
	"aws_codecommit_approval_rule_template_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.approval_rule_template_name }},{{ .parameters.repository_name }}"),
	// No import
	"aws_codecommit_trigger": config.IdentifierFromProvider,

	// deploy
	//
	// CodeDeploy Applications can be imported using the name
	"aws_codedeploy_app": config.TemplatedStringAsIdentifier("name", "{{ .parameters.application_id }}:{{ .external_name }}"),
	// CodeDeploy Deployment Configurations can be imported using the deployment_config_name
	"aws_codedeploy_deployment_config": config.ParameterAsIdentifier("deployment_config_name"),
	// CodeDeploy Deployment Groups can be imported by their app_name, a colon, and deployment_group_name
	"aws_codedeploy_deployment_group": config.TemplatedStringAsIdentifier("deployment_group_name", "{{ .parameters.app_name }}:{{ .external_name }}"),

	// codepipeline
	//
	// CodePipelines can be imported using the name
	"aws_codepipeline": config.NameAsIdentifier,
	// CodePipeline Webhooks can be imported by their ARN: arn:aws:codepipeline:us-west-2:123456789012:webhook:example
	"aws_codepipeline_webhook": config.TemplatedStringAsIdentifier("name", "arn:aws:codepipeline:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:webhook:{{ .external_name }}"),

	// codestarconnections
	//
	// CodeStar connections can be imported using the ARN
	"aws_codestarconnections_connection": config.IdentifierFromProvider,
	// CodeStar Host can be imported using the ARN
	"aws_codestarconnections_host": config.IdentifierFromProvider,

	// codestarnotifications
	//
	// CodeStar notification rule can be imported using the ARN
	"aws_codestarnotifications_notification_rule": config.IdentifierFromProvider,

	// connect
	//
	// aws_connect_bot_association can be imported by using the Amazon Connect instance ID, Lex (V1) bot name, and Lex (V1) bot region separated by colons (:)
	// TODO: lex_bot.lex_region parameter is not `Required` in TF schema. But we use this field in id construction. So, please mark as required this field while configuration
	"aws_connect_bot_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.instance_id }}:{{ (index .parameters.lex_bot 0).name }}:{{ (index .parameters.lex_bot 0).lex_region }}"),
	// Amazon Connect Contact Flows can be imported using the instance_id and contact_flow_id separated by a colon (:)
	"aws_connect_contact_flow": config.IdentifierFromProvider,
	// Amazon Connect Contact Flow Modules can be imported using the instance_id and contact_flow_module_id separated by a colon (:)
	"aws_connect_contact_flow_module": config.IdentifierFromProvider,
	// Amazon Connect Hours of Operations can be imported using the instance_id and hours_of_operation_id separated by a colon (:)
	"aws_connect_hours_of_operation": config.IdentifierFromProvider,
	// Connect instances can be imported using the id
	"aws_connect_instance": config.IdentifierFromProvider,
	// aws_connect_lambda_function_association can be imported using the instance_id and function_arn separated by a comma (,)
	"aws_connect_lambda_function_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.instance_id }},{{ .parameters.function_arn }}"),
	// Amazon Connect Queues can be imported using the instance_id and queue_id separated by a colon (:)
	"aws_connect_queue": config.IdentifierFromProvider,
	// Amazon Connect Quick Connects can be imported using the instance_id and quick_connect_id separated by a colon (:)
	"aws_connect_quick_connect": config.IdentifierFromProvider,
	// Amazon Connect Routing Profiles can be imported using the instance_id and routing_profile_id separated by a colon (:)
	"aws_connect_routing_profile": config.IdentifierFromProvider,
	// Amazon Connect Security Profiles can be imported using the instance_id and security_profile_id separated by a colon (:)
	"aws_connect_security_profile": config.IdentifierFromProvider,
	// Amazon Connect User Hierarchy Structures can be imported using the instance_id
	"aws_connect_user_hierarchy_structure": config.IdentifierFromProvider,

	// apprunner
	//
	// App Runner AutoScaling Configuration Versions can be imported by using the arn
	"aws_apprunner_auto_scaling_configuration_version": config.IdentifierFromProvider,
	// App Runner Connections can be imported by using the connection_name
	"aws_apprunner_connection": config.ParameterAsIdentifier("connection_name"),
	// App Runner Services can be imported by using the arn
	"aws_apprunner_service": config.IdentifierFromProvider,
	// App Runner vpc connector can be imported by using the arn
	"aws_apprunner_vpc_connector": config.IdentifierFromProvider,

	// appstream
	//
	// aws_appstream_directory_config can be imported using the id
	"aws_appstream_directory_config": config.IdentifierFromProvider,
	// aws_appstream_fleet can be imported using the id
	"aws_appstream_fleet": config.IdentifierFromProvider,
	// AppStream Stack Fleet Association can be imported by using the fleet_name and stack_name separated by a slash (/)
	"aws_appstream_fleet_stack_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.fleet_name }}/{{ .parameters.stack_name}}"),
	// aws_appstream_image_builder can be imported using the name
	"aws_appstream_image_builder": config.NameAsIdentifier,
	// aws_appstream_stack can be imported using the id
	"aws_appstream_stack": config.IdentifierFromProvider,
	// aws_appstream_user can be imported using the user_name and authentication_type separated by a slash (/)
	"aws_appstream_user": config.TemplatedStringAsIdentifier("user_name", "{{ .external_name }}/{{ .parameters.authentication_type }}"),
	// AppStream User Stack Association can be imported by using the user_name, authentication_type, and stack_name, separated by a slash (/)
	"aws_appstream_user_stack_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.user_name }}/{{ .parameters.authentication_type }}/{{ .parameters.stack_name }}/"),

	// appmesh
	//
	// App Mesh service meshes can be imported using the name
	"aws_appmesh_mesh": config.NameAsIdentifier,
	// App Mesh virtual nodes can be imported using mesh_name together with the virtual node's name: simpleapp/serviceBv1
	"aws_appmesh_virtual_node": config.IdentifierFromProvider,
	// App Mesh virtual routers can be imported using mesh_name together with the virtual router's name: simpleapp/serviceB
	"aws_appmesh_virtual_router": config.IdentifierFromProvider,
	// App Mesh virtual gateway can be imported using mesh_name together with the virtual gateway's name: mesh/gw1
	"aws_appmesh_virtual_gateway": config.IdentifierFromProvider,
	// App Mesh virtual services can be imported using mesh_name together with the virtual service's name: simpleapp/servicea.simpleapp.local
	"aws_appmesh_virtual_service": config.IdentifierFromProvider,
	// mesh/gw1/example-gateway-route
	"aws_appmesh_gateway_route": config.IdentifierFromProvider,
	// App Mesh virtual routes can be imported using mesh_name and virtual_router_name together with the route's name, e.g.,
	// simpleapp/serviceB/serviceB-route
	"aws_appmesh_route": config.IdentifierFromProvider,

	// configservice
	//
	// Config Rule can be imported using the name
	"aws_config_config_rule": config.NameAsIdentifier,
	// Configuration Aggregators can be imported using the name
	"aws_config_configuration_aggregator": config.NameAsIdentifier,
	// Configuration Recorder can be imported using the name
	"aws_config_configuration_recorder": config.NameAsIdentifier,
	// Configuration Recorder Status can be imported using the name of the Configuration Recorder
	"aws_config_configuration_recorder_status": config.NameAsIdentifier,
	// Config Conformance Packs can be imported using the name
	"aws_config_conformance_pack": config.NameAsIdentifier,
	// Delivery Channel can be imported using the name
	"aws_config_delivery_channel": config.NameAsIdentifier,
	// Remediation Configurations can be imported using the name config_rule_name
	"aws_config_remediation_configuration": config.ParameterAsIdentifier("config_rule_name"),

	// appsync
	//
	// aws_appsync_api_cache can be imported using the AppSync API ID
	"aws_appsync_api_cache": config.IdentifierFromProvider,
	// aws_appsync_api_key can be imported using the AppSync API ID and key separated by :
	"aws_appsync_api_key": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}:{{ .external_name }}"),
	// aws_appsync_datasource can be imported with their api_id, a hyphen, and name
	"aws_appsync_datasource": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}-{{ .external_name }}"),
	// aws_appsync_function can be imported using the AppSync API ID and Function ID separated by -
	"aws_appsync_function": config.IdentifierFromProvider,
	// AppSync GraphQL API can be imported using the GraphQL API ID
	"aws_appsync_graphql_api": config.IdentifierFromProvider,
	// aws_appsync_resolver can be imported with their api_id, a hyphen, type, a hypen and field
	"aws_appsync_resolver": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_id }}-{{ .parameters.type }}-{{ .parameters.field }}"),

	// accessanalyzer
	//
	// Access Analyzer Analyzers can be imported using the analyzer_name
	"aws_accessanalyzer_analyzer": config.ParameterAsIdentifier("analyzer_name"),

	// account
	//
	// The Alternate Contact for the current account can be imported using the alternate_contact_type
	"aws_account_alternate_contact": config.TemplatedStringAsIdentifier("", "{{ .parameters.alternate_contact_type }}"),

	// amplify
	//
	// Amplify App can be imported using Amplify App ID (appId)
	"aws_amplify_app": config.IdentifierFromProvider,
	// Amplify branch can be imported using app_id and branch_name: d2ypk4k47z8u6/master
	"aws_amplify_branch": config.TemplatedStringAsIdentifier("branch_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify backend environment can be imported using app_id and environment_name: d2ypk4k47z8u6/example
	"aws_amplify_backend_environment": config.TemplatedStringAsIdentifier("environment_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify webhook can be imported using a webhook ID
	"aws_amplify_webhook": config.IdentifierFromProvider,

	// cur
	//
	// Report Definitions can be imported using the report_name
	"aws_cur_report_definition": config.ParameterAsIdentifier("report_name"),

	// dataexchange
	//
	// DataExchange DataSets can be imported by their arn
	"aws_dataexchange_data_set": config.IdentifierFromProvider,
	// DataExchange Revisions can be imported by their data-set-id:revision-id
	"aws_dataexchange_revision": config.IdentifierFromProvider,

	// datapipeline
	//
	// aws_datapipeline_pipeline can be imported by using the id (Pipeline ID)
	"aws_datapipeline_pipeline": config.IdentifierFromProvider,

	// detective
	//
	// aws_detective_graph can be imported using the ARN
	"aws_detective_graph": config.IdentifierFromProvider,
	// aws_detective_member can be imported using the ARN of the graph followed by the account ID of the member account
	"aws_detective_member": config.IdentifierFromProvider,
	// aws_detective_invitation_accepter can be imported using the graph ARN
	"aws_detective_invitation_accepter": config.IdentifierFromProvider,

	// devicefarm
	//
	// DeviceFarm Projects can be imported by their arn
	"aws_devicefarm_project": config.IdentifierFromProvider,
	// DeviceFarm Instance Profiles can be imported by their arn
	"aws_devicefarm_instance_profile": config.IdentifierFromProvider,
	// DeviceFarm Device Pools can be imported by their arn
	"aws_devicefarm_device_pool": config.IdentifierFromProvider,
	// DeviceFarm Network Profiles can be imported by their arn
	"aws_devicefarm_network_profile": config.IdentifierFromProvider,
	// DeviceFarm Uploads can be imported by their arn
	"aws_devicefarm_upload": config.IdentifierFromProvider,
	// DeviceFarm Test Grid Projects can be imported by their arn
	"aws_devicefarm_test_grid_project": config.IdentifierFromProvider,

	// organization
	//
	// imported by using the account id, which is provider-generated
	"aws_organizations_account": config.IdentifierFromProvider,
	// imported by using the account ID and its service principal:
	// 123456789012/config.amazonaws.com
	"aws_organizations_delegated_administrator": FormattedIdentifierFromProvider("/", "account_id", "service_principal"),
	//  imported by using the id, which is a Cloud provider-generated string:
	// o-1234567
	"aws_organizations_organization": config.IdentifierFromProvider,
	// imported by using the id, which is a Cloud provider-generated string:
	// ou-1234567
	"aws_organizations_organizational_unit": config.IdentifierFromProvider,
	// imported by using the policy ID,
	// which is a Cloud provider-generated string:
	// p-12345678
	"aws_organizations_policy": config.IdentifierFromProvider,
	// imported by using the target ID and policy ID
	// 123456789012:p-12345678
	"aws_organizations_policy_attachment": FormattedIdentifierFromProvider(":", "target_id", "policy_id"),

	// batch
	//
	// Batch Scheduling Policy can be imported using the arn: arn:aws:batch:us-east-1:123456789012:scheduling-policy/sample
	"aws_batch_scheduling_policy": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:scheduling-policy/{{ .external_name }}"),

	// budgets
	//
	// Budgets can be imported using AccountID:BudgetName
	"aws_budgets_budget": config.TemplatedStringAsIdentifier("name", "{{ .setup.client_metadata.account_id }}:{{ .external_name }}"),
	// Budgets can be imported using AccountID:ActionID:BudgetName
	"aws_budgets_budget_action": config.IdentifierFromProvider,

	// chime
	//
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector": config.NameAsIdentifier,
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector_group": config.NameAsIdentifier,
	// Chime Voice Connector Logging can be imported using the voice_connector_id
	"aws_chime_voice_connector_logging": config.IdentifierFromProvider,
	// Chime Voice Connector Origination can be imported using the voice_connector_id
	"aws_chime_voice_connector_origination": config.IdentifierFromProvider,
	// Chime Voice Connector Streaming can be imported using the voice_connector_id
	"aws_chime_voice_connector_streaming": config.IdentifierFromProvider,
	// Chime Voice Connector Termination can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination": config.IdentifierFromProvider,
	// Chime Voice Connector Termination Credentials can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination_credentials": config.IdentifierFromProvider,

	// quicksight
	//
	// QuickSight Group can be imported using the aws account id, namespace and group name separated by /
	// 123456789123/default/tf-example
	"aws_quicksight_group": FormattedIdentifierFromProvider("/", "aws_account_id", "namespace", "group_name"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_quicksight_user": config.IdentifierFromProvider,

	// lightsail
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_domain": config.IdentifierFromProvider,
	// Lightsail Instances can be imported using their name
	"aws_lightsail_instance": config.NameAsIdentifier,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_instance_public_ports": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_key_pair": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_static_ip": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_static_ip_attachment": config.IdentifierFromProvider,

	// cloud9
	//
	// No import
	"aws_cloud9_environment_ec2": config.IdentifierFromProvider,
	// Cloud9 environment membership can be imported using the environment-id#user-arn
	"aws_cloud9_environment_membership": config.TemplatedStringAsIdentifier("", "{{ .parameters.environment_id }}#{{ .parameters.user_arn }}"),

	// cloudcontrol
	//
	// No import
	"aws_cloudcontrolapi_resource": config.IdentifierFromProvider,

	// securityhub
	//
	// An existing Security Hub enabled account can be imported using the AWS account ID
	"aws_securityhub_account": config.IdentifierFromProvider,
	// imported using the action target ARN:
	// arn:aws:securityhub:eu-west-1:312940875350:action/custom/a
	// TODO: following configuration assumes the `a` in the above ARN
	// is the security hub custom action identifier
	"aws_securityhub_action_target": config.TemplatedStringAsIdentifier("identifier", "arn:aws:securityhub:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:action/custom/{{ .external_name }}"),
	// imported using the arn that has a random substring:
	// arn:aws:securityhub:eu-west-1:123456789098:finding-aggregator/abcd1234-abcd-1234-1234-abcdef123456
	"aws_securityhub_finding_aggregator": config.IdentifierFromProvider,
	// imported using the ARN that has a random substring:
	// arn:aws:securityhub:us-west-2:1234567890:insight/1234567890/custom/91299ed7-abd0-4e44-a858-d0b15e37141a
	"aws_securityhub_insight": config.IdentifierFromProvider,
	// imported using security hub member account ID
	"aws_securityhub_member": FormattedIdentifierFromProvider("", "account_id"),
	// imported in the form product_arn,arn:
	// arn:aws:securityhub:eu-west-1:733251395267:product/alertlogic/althreatmanagement,arn:aws:securityhub:eu-west-1:123456789012:product-subscription/alertlogic/althreatmanagement
	// looks like it's possible to derive the external-name from
	// the product_arn argument according to the above example
	// (by replacing product by product-subscription), which makes this
	// a special case of FormattedIdentifierFromProvider
	"aws_securityhub_product_subscription": func() config.ExternalName {
		e := config.IdentifierFromProvider
		e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			val, ok := parameters["product_arn"]
			if !ok {
				return "", errors.New("product_arn cannot be empty")
			}
			s, ok := val.(string)
			if !ok {
				return "", errors.New("product_arn needs to be a string")
			}
			return fmt.Sprintf("%s,%s", s, strings.Replace(s, ":product", ":product-subscription", 1)), nil
		}
		return e
	}(),
	// imported using the standards subscription ARN:
	// arn:aws:securityhub:eu-west-1:123456789012:subscription/pci-dss/v/3.2.1
	"aws_securityhub_standards_subscription": FormattedIdentifierFromProvider("", "standards_arn"),
	// imported using the account ID
	"aws_securityhub_invite_accepter": config.IdentifierFromProvider,

	// cloudformation
	//
	// config.NameAsIdentifier did not work, the identifier for the resource turned out to be an ARN
	// arn:aws:cloudformation:us-west-1:123456789123:stack/networking-stack/1e691240-6f2c-11ed-8f91-06094dc221f3
	"aws_cloudformation_stack": TemplatedStringAsIdentifierWithNoName("arn:aws:cloudformation:{{ .parameters.region }}:{{ .client_metadata.account_id }}:stack/{{ .parameters.name }}/{{ .external_name }}"),
	// CloudFormation StackSets can be imported using the name
	"aws_cloudformation_stack_set": config.NameAsIdentifier,

	// autoscaling
	//
	// aws_autoscaling_group_tag can be imported by using the ASG name and key, separated by a comma (,)
	"aws_autoscaling_group_tag": config.IdentifierFromProvider,
	// AutoScaling Lifecycle Hooks can be imported using the role autoscaling_group_name and name separated by /
	"aws_autoscaling_lifecycle_hook": config.TemplatedStringAsIdentifier("name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// No import
	"aws_autoscaling_notification": config.IdentifierFromProvider,
	// AutoScaling scaling policy can be imported using the role autoscaling_group_name and name separated by /
	"aws_autoscaling_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// AutoScaling ScheduledAction can be imported using the auto-scaling-group-name and scheduled-action-name: auto-scaling-group-name/scheduled-action-name
	"aws_autoscaling_schedule": config.TemplatedStringAsIdentifier("scheduled_action_name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),

	// autoscalingplans
	//
	// Auto Scaling scaling plans can be imported using the name
	"aws_autoscalingplans_scaling_plan": config.IdentifierFromProvider,

	// serverlessapplicationrepository
	//
	// imported using the CloudFormation Stack name
	"aws_serverlessapplicationrepository_cloudformation_stack": config.IdentifierFromProvider,

	// directconnect
	//
	// Direct Connect Gateways can be imported using the gateway id
	"aws_dx_gateway": config.IdentifierFromProvider,
	// Direct Connect connections can be imported using the connection id
	"aws_dx_connection": config.IdentifierFromProvider,
	// Direct Connect public virtual interfaces can be imported using the vif id
	"aws_dx_public_virtual_interface": config.IdentifierFromProvider,
	// No import
	"aws_dx_connection_association": config.IdentifierFromProvider,
	// Direct Connect LAGs can be imported using the lag id
	"aws_dx_lag": config.IdentifierFromProvider,
	// Direct Connect transit virtual interfaces can be imported using the vif id
	"aws_dx_transit_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect private virtual interfaces can be imported using the vif id
	"aws_dx_private_virtual_interface": config.IdentifierFromProvider,
	//
	"aws_dx_gateway_association_proposal": config.IdentifierFromProvider,
	// Direct Connect gateway associations can be imported using dx_gateway_id together with associated_gateway_id
	// TODO: associated_gateway_id parameter is not `Required` in TF schema. But we use this field in id construction. So, please mark as required this field while configuration
	"aws_dx_gateway_association": config.IdentifierFromProvider,
	// No import
	"aws_dx_bgp_peer": config.IdentifierFromProvider,
	// Direct Connect hosted private virtual interfaces can be imported using the vif id
	"aws_dx_hosted_private_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted private virtual interfaces can be imported using the vif id
	"aws_dx_hosted_private_virtual_interface_accepter": config.IdentifierFromProvider,
	// Direct Connect hosted public virtual interfaces can be imported using the vif id
	"aws_dx_hosted_public_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted public virtual interfaces can be imported using the vif id
	"aws_dx_hosted_public_virtual_interface_accepter": config.IdentifierFromProvider,
	// Direct Connect hosted transit virtual interfaces can be imported using the vif id
	"aws_dx_hosted_transit_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted transit virtual interfaces can be imported using the vif id
	"aws_dx_hosted_transit_virtual_interface_accepter": config.IdentifierFromProvider,

	// guardduty
	//
	// GuardDuty detectors can be imported using the detector ID
	"aws_guardduty_detector": config.IdentifierFromProvider,
	// GuardDuty filters can be imported using the detector ID and filter's name separated by a colon
	// 00b00fd5aecc0ab60a708659477e9617:MyFilter
	"aws_guardduty_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.detector_id }}:{{ .external_name }}"),
	// GuardDuty members can be imported using the primary GuardDuty detector ID and member AWS account ID
	// 00b00fd5aecc0ab60a708659477e9617:123456789012
	"aws_guardduty_member": config.IdentifierFromProvider,

	// appconfig
	//
	// AppConfig Applications can be imported using their application ID,
	"aws_appconfig_application": config.IdentifierFromProvider,
	// AppConfig Deployment Strategies can be imported by using their deployment strategy ID
	"aws_appconfig_deployment_strategy": config.IdentifierFromProvider,
	// AppConfig Environments can be imported by using the environment ID and application ID separated by a colon (:)
	"aws_appconfig_environment": config.IdentifierFromProvider,
	// AppConfig Configuration Profiles can be imported by using the configuration profile ID and application ID separated by a colon (:)
	"aws_appconfig_configuration_profile": config.IdentifierFromProvider,
	// AppConfig Hosted Configuration Versions can be imported by using the application ID, configuration profile ID, and version number separated by a slash (/)
	"aws_appconfig_hosted_configuration_version": config.IdentifierFromProvider,
	// AppConfig Deployments can be imported by using the application ID, environment ID, and deployment number separated by a slash (/)
	"aws_appconfig_deployment": config.IdentifierFromProvider,

	// appintegrations
	//
	// Amazon AppIntegrations Event Integrations can be imported using the name
	"aws_appintegrations_event_integration": config.NameAsIdentifier,

	// grafana
	//
	// Grafana workspace license association can be imported using the workspace's id
	"aws_grafana_license_association": FormattedIdentifierFromProvider("", "workspace_id"),

	// appflow
	//
	// arn:aws:appflow:us-west-2:123456789012:flow/example-flow
	"aws_appflow_flow": config.TemplatedStringAsIdentifier("name", "arn:aws:appflow:{{ .parameters.region }}:{{ .client_metadata.account_id }}:flow/{{ .external_name }}"),

	// sns
	//
	// SNS platform applications can be imported using the ARN:
	// arn:aws:sns:us-west-2:0123456789012:app/GCM/gcm_application
	"aws_sns_platform_application": config.TemplatedStringAsIdentifier("name", "arn:aws:sns:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:app/GCM/{{ .external_name }}"),
	// no import documentation is provided
	// TODO: we will need to check if normalization is possible
	"aws_sns_sms_preferences": config.IdentifierFromProvider,
	// SNS Topic Policy can be imported using the topic ARN:
	// arn:aws:sns:us-west-2:0123456789012:my-topic
	"aws_sns_topic_policy": FormattedIdentifierFromProvider("", "arn"),

	// servicecatalog
	//
	// imported using the tag option ID, which has provider-generated
	// random parts: tag-pjtvagohlyo3m
	"aws_servicecatalog_tag_option": config.IdentifierFromProvider,
	// imported using the service catalog portfolio id,
	// which has provider-generated random parts:
	// port-12344321
	"aws_servicecatalog_portfolio": config.IdentifierFromProvider,
	// imported using the portfolio share ID: port-12344321:ACCOUNT:123456789012
	// TODO: looks like we can generate the above ID using:
	// portfolio_id:type:principal_id
	// but this has to be validated
	"aws_servicecatalog_portfolio_share": config.IdentifierFromProvider,
	// imported using the accept language, principal ARN, and portfolio ID, separated by a comma:
	// en,arn:aws:iam::123456789012:user/Eleanor,port-68656c6c6f
	// TODO: looks like we can generated the above id using:
	// accept_language,principal_arn,portfolio_id
	// , which lends itself to:
	// FormattedIdentifierFromProvider(",", "accept_language", "principal_arn", "portfolio_id")
	// However, accept_language is optional. We had better make it required as
	// the default is provided by Terraform (and we have no means to default
	// the generated CRD fields as of now)
	"aws_servicecatalog_principal_portfolio_association": config.IdentifierFromProvider,
	// imported using the product ID, which has provider-generated random parts:
	// prod-dnigbtea24ste
	"aws_servicecatalog_product": config.IdentifierFromProvider,
	// imported using the accept language, portfolio ID, and product ID:
	// en:port-68656c6c6f:prod-dnigbtea24ste
	// TODO: looks like we can generated the above id using:
	// accept_language,portfolio_id,product_id
	// , which lends itself to:
	// FormattedIdentifierFromProvider(",", "accept_language", "portfolio_id", "product_id")
	// However, accept_language is optional. We had better make it required as
	// the default is provided by Terraform (and we have no means to default
	// the generated CRD fields as of now)
	"aws_servicecatalog_product_portfolio_association": config.IdentifierFromProvider,
	// imported using the constraint ID, which has random parts
	// generated by the provider: cons-nmdkb6cgxfcrs
	"aws_servicecatalog_constraint": config.IdentifierFromProvider,
	// imported using the tag option ID and resource ID:
	// tag-pjtvyakdlyo3m:prod-dnigbtea24ste
	"aws_servicecatalog_tag_option_resource_association": FormattedIdentifierFromProvider(":", "tag_option_id", "resource_id"),
	// imported using the budget name and resource ID:
	// budget-pjtvyakdlyo3m:prod-dnigbtea24ste
	"aws_servicecatalog_budget_resource_association": config.IdentifierFromProvider,
	// imported using the provisioning artifact ID and product ID separated by a colon:
	// pa-ij2b6lusy6dec:prod-el3an0rma3
	// we could make the product_id attribute the name identifier
	// and concatenate it with the provider-generated provisioning
	// artifact id, but product id does not does not look like to
	// be a good external-name for this resource as this is the
	// provisioning artifact resource.
	"aws_servicecatalog_provisioning_artifact": config.IdentifierFromProvider,
	// imported using the service action ID. which has provider-generated
	// random parts: act-f1w12eperfslh
	"aws_servicecatalog_service_action": config.IdentifierFromProvider,

	// keyspaces
	//
	// Use the name to import a keyspace
	"aws_keyspaces_keyspace": config.NameAsIdentifier,
	// Use the keyspace_name and table_name separated by / to import a table
	// my_keyspace/my_table
	"aws_keyspaces_table": FormattedIdentifierFromProvider("/", "keyspace_name", "table_name"),

	// route53recoveryreadiness
	//
	// Route53 Recovery Readiness recovery groups can be imported via the recovery group name
	"aws_route53recoveryreadiness_recovery_group": config.ParameterAsIdentifier("recovery_group_name"),
	// Route53 Recovery Readiness resource set name can be imported via the resource set name
	"aws_route53recoveryreadiness_resource_set": config.ParameterAsIdentifier("resource_set_name"),
	// Route53 Recovery Readiness cells can be imported via the cell name
	"aws_route53recoveryreadiness_cell": config.ParameterAsIdentifier("cell_name"),
	// Route53 Recovery Readiness readiness checks can be imported via the readiness check name
	"aws_route53recoveryreadiness_readiness_check": config.ParameterAsIdentifier("readiness_check_name"),

	// s3control
	//
	// - For Access Points associated with an AWS Partition S3 Bucket, this resource
	// can be imported using the account_id and name separated by a colon (:)
	// - For Access Points associated with an S3 on Outposts Bucket, this resource
	// can be imported using the Amazon Resource Name (ARN)
	// TODO: There are two different import syntaxes for this resource. For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_s3_access_point": config.IdentifierFromProvider,
	// aws_s3_account_public_access_block can be imported by using the AWS account ID
	"aws_s3_account_public_access_block": config.IdentifierFromProvider,
	// Access Point policies can be imported using the access_point_arn
	// arn:aws:s3:us-west-2:123456789012:accesspoint/example
	"aws_s3control_access_point_policy": config.IdentifierFromProvider,

	// dlm
	//
	// DLM lifecycle policies can be imported by their policy ID
	"aws_dlm_lifecycle_policy": config.IdentifierFromProvider,

	// dms
	//
	// Certificates can be imported using the certificate_id
	"aws_dms_certificate": config.ParameterAsIdentifier("certificate_id"),
	// Endpoints can be imported using the endpoint_id
	"aws_dms_endpoint": config.ParameterAsIdentifier("endpoint_id"),
	// Replication subnet groups can be imported using the replication_subnet_group_id
	"aws_dms_replication_subnet_group": config.ParameterAsIdentifier("replication_subnet_group_id"),

	// ds
	//
	// DirectoryService directories can be imported using the directory id
	"aws_directory_service_directory": config.IdentifierFromProvider,

	// elastictranscoder
	//
	// Elastic Transcoder pipelines can be imported using the id
	"aws_elastictranscoder_pipeline": config.IdentifierFromProvider,
	// Elastic Transcoder presets can be imported using the id
	"aws_elastictranscoder_preset": config.IdentifierFromProvider,

	// schemas
	//
	// EventBridge discoverers can be imported using the id
	"aws_schemas_discoverer": config.IdentifierFromProvider,
	// EventBridge schema registries can be imported using the name
	"aws_schemas_registry": config.NameAsIdentifier,
	// EventBridge schema can be imported using the name and registry_name
	"aws_schemas_schema": FormattedIdentifierFromProvider("/", "name", "registry_name"),

	// mediapackage
	//
	// Media Package Channels can be imported via the channel ID
	"aws_media_package_channel": config.IdentifierFromProvider,

	// mediastore
	//
	// MediaStore Container can be imported using the MediaStore Container Name
	"aws_media_store_container": config.NameAsIdentifier,
	// MediaStore Container Policy can be imported using the MediaStore Container Name
	"aws_media_store_container_policy": FormattedIdentifierFromProvider("", "container_name"),

	// macie2
	//
	// aws_macie2_account can be imported using the id
	"aws_macie2_account": config.IdentifierFromProvider,
	// aws_macie2_classification_job can be imported using the id
	"aws_macie2_classification_job": config.IdentifierFromProvider,
	// aws_macie2_custom_data_identifier can be imported using the id
	"aws_macie2_custom_data_identifier": config.IdentifierFromProvider,
	// aws_macie2_findings_filter can be imported using the id
	"aws_macie2_findings_filter": config.IdentifierFromProvider,
	// aws_macie2_invitation_accepter can be imported using the admin account ID
	"aws_macie2_invitation_accepter": FormattedIdentifierFromProvider("", "administrator_account_id"),
	// aws_macie2_member can be imported using the account ID of the member account
	"aws_macie2_member": FormattedIdentifierFromProvider("", "account_id"),

	// mediaconvert
	//
	// Media Convert Queue can be imported via the queue name
	"aws_media_convert_queue": config.NameAsIdentifier,

	// servicequotas
	//
	// aws_servicequotas_service_quota can be imported by using the service code and quota code, separated by a front slash (/)
	// vpc/L-F678F1CE
	"aws_servicequotas_service_quota": FormattedIdentifierFromProvider("/", "service_code", "quota_code"),

	// pinpoint
	//
	// Pinpoint App can be imported using the application-id
	"aws_pinpoint_app": config.IdentifierFromProvider,
	// Pinpoint SMS Channel can be imported using the application-id
	"aws_pinpoint_sms_channel": FormattedIdentifierFromProvider("", "application_id"),

	// elasticbeanstalk
	//
	// Elastic Beanstalk Applications can be imported using the name
	"aws_elastic_beanstalk_application": config.NameAsIdentifier,
	// No import
	"aws_elastic_beanstalk_configuration_template": config.NameAsIdentifier,

	// ssm
	//
	// AWS SSM Activation can be imported using the id
	"aws_ssm_activation": config.IdentifierFromProvider,
	// SSM associations can be imported using the association_id
	"aws_ssm_association": config.IdentifierFromProvider,
	// SSM Documents can be imported using the name
	"aws_ssm_document": config.NameAsIdentifier,
	// SSM Maintenance Windows can be imported using the maintenance window id
	"aws_ssm_maintenance_window": config.IdentifierFromProvider,
	// SSM Maintenance Window targets can be imported using WINDOW_ID/WINDOW_TARGET_ID
	"aws_ssm_maintenance_window_target": config.IdentifierFromProvider,
	// SSM Patch Baselines can be imported by their baseline ID
	"aws_ssm_patch_baseline": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_ssm_patch_group": config.IdentifierFromProvider,

	// emr
	//
	// EMR Security Configurations can be imported using the name
	"aws_emr_security_configuration": config.NameAsIdentifier,

	// qldb
	//
	// QLDB Ledgers can be imported using the name
	"aws_qldb_ledger": config.NameAsIdentifier,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_qldb_stream": config.IdentifierFromProvider,

	// glacier
	//
	// Glacier Vaults can be imported using the name
	"aws_glacier_vault": config.NameAsIdentifier,
	// Glacier Vault Locks can be imported using the Glacier Vault name
	"aws_glacier_vault_lock": FormattedIdentifierFromProvider("", "vault_name"),

	// iot
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_certificate": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_indexing_configuration": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_logging_options": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_policy_attachment": config.IdentifierFromProvider,
	// IoT fleet provisioning templates can be imported using the name
	"aws_iot_provisioning_template": config.NameAsIdentifier,
	// IOT Role Alias can be imported via the alias
	"aws_iot_role_alias": config.IdentifierFromProvider,
	// IoT Things Groups can be imported using the name
	"aws_iot_thing_group": config.NameAsIdentifier,
	// IoT Thing Group Membership can be imported using the thing group name and thing name
	// thing_group_name/thing_name
	"aws_iot_thing_group_membership": FormattedIdentifierFromProvider("/", "thing_group_name", "thing_name"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_thing_principal_attachment": config.IdentifierFromProvider,
	// IOT Thing Types can be imported using the name
	"aws_iot_thing_type": config.IdentifierFromProvider,
	// IoT Topic Rules can be imported using the name
	"aws_iot_topic_rule": config.NameAsIdentifier,

	// sagemaker
	//
	// SageMaker App Image Configs can be imported using the name
	"aws_sagemaker_app_image_config": config.ParameterAsIdentifier("app_image_config_name"),
	// SageMaker Code Repositories can be imported using the name
	"aws_sagemaker_code_repository": config.ParameterAsIdentifier("code_repository_name"),
	// SageMaker Domains can be imported using the id
	"aws_sagemaker_domain": config.IdentifierFromProvider,
	// Feature Groups can be imported using the name
	"aws_sagemaker_feature_group": config.ParameterAsIdentifier("feature_group_name"),
	// SageMaker Code Images can be imported using the name
	"aws_sagemaker_image": config.ParameterAsIdentifier("image_name"),
	// SageMaker Model Package Groups can be imported using the name
	"aws_sagemaker_model_package_group": config.ParameterAsIdentifier("model_package_group_name"),
	// SageMaker Notebook Instances can be imported using the name
	"aws_sagemaker_notebook_instance": config.NameAsIdentifier,
	// Models can be imported using the name
	"aws_sagemaker_notebook_instance_lifecycle_configuration": config.NameAsIdentifier,
	// SageMaker Studio Lifecycle Configs can be imported using the studio_lifecycle_config_name
	"aws_sagemaker_studio_lifecycle_config": config.ParameterAsIdentifier("studio_lifecycle_config_name"),
	// SageMaker User Profiles can be imported using the arn
	"aws_sagemaker_user_profile": config.IdentifierFromProvider,

	// elbv2
	//
	// Rules can be imported using their ARN
	"aws_lb_listener_rule": config.IdentifierFromProvider,

	// fsx
	//
	// FSx File Systems can be imported using the id
	"aws_fsx_windows_file_system": config.IdentifierFromProvider,
	// FSx File Systems can be imported using the id
	"aws_fsx_lustre_file_system": config.IdentifierFromProvider,
	// FSx Backups can be imported using the id
	"aws_fsx_backup": config.IdentifierFromProvider,
	// FSx Data Repository Associations can be imported using the id
	"aws_fsx_data_repository_association": config.IdentifierFromProvider,

	// route53recoverycontrolconfig
	//
	// Imported using ARN that has a random substring:
	// arn:aws:route53-recovery-control::313517334327:cluster/f9ae13be-a11e-4ec7-8522-94a70468e6ea
	"aws_route53recoverycontrolconfig_cluster": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	// arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8
	"aws_route53recoverycontrolconfig_control_panel": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	// arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8/routingcontrol/d5d90e587870494b
	"aws_route53recoverycontrolconfig_routing_control": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	// arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8/safetyrule/3bacc77003364c0f
	"aws_route53recoverycontrolconfig_safety_rule": config.IdentifierFromProvider,

	// memorydb
	//
	// Use the name to import a parameter group
	"aws_memorydb_parameter_group": config.NameAsIdentifier,
	// Use the name to import a subnet group
	"aws_memorydb_subnet_group": config.NameAsIdentifier,
	// Use the name to import a cluster
	"aws_memorydb_cluster": config.NameAsIdentifier,
	// Use the name to import an ACL
	"aws_memorydb_acl": config.NameAsIdentifier,
	// Use the name to import a snapshot
	"aws_memorydb_snapshot": config.NameAsIdentifier,

	// imagebuilder
	//
	// aws_imagebuilder_container_recipe resources can be imported by using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/example/1.0.0
	"aws_imagebuilder_container_recipe": config.IdentifierFromProvider,
	// aws_imagebuilder_distribution_configurations resources can be imported by using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:distribution-configuration/example
	"aws_imagebuilder_distribution_configuration": config.IdentifierFromProvider,
	// aws_imagebuilder_image resources can be imported using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:image/example/1.0.0/1
	// TODO: Normalize external_name while testing
	"aws_imagebuilder_image": config.IdentifierFromProvider,
	// aws_imagebuilder_image_pipeline resources can be imported using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/example
	"aws_imagebuilder_image_pipeline": config.IdentifierFromProvider,
	// aws_imagebuilder_image_recipe resources can be imported by using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:image-recipe/example/1.0.0
	"aws_imagebuilder_image_recipe": config.IdentifierFromProvider,
	// aws_imagebuilder_infrastructure_configuration can be imported using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:infrastructure-configuration/example
	"aws_imagebuilder_infrastructure_configuration": config.IdentifierFromProvider,

	// inspector
	//
	// Inspector Assessment Targets can be imported via their Amazon Resource Name (ARN)
	// Example: arn:aws:inspector:us-east-1:123456789012:target/0-xxxxxxx
	"aws_inspector_assessment_target": config.IdentifierFromProvider,
	// aws_inspector_assessment_template can be imported by using the template assessment ARN
	// Example: arn:aws:inspector:us-west-2:123456789012:target/0-9IaAzhGR/template/0-WEcjR8CH
	"aws_inspector_assessment_template": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_inspector_resource_group": config.IdentifierFromProvider,

	// ses
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_ses_active_receipt_rule_set": config.IdentifierFromProvider,
	// SES Configuration Sets can be imported using their name
	"aws_ses_configuration_set": config.NameAsIdentifier,
	// DKIM tokens can be imported using the domain attribute
	"aws_ses_domain_dkim": config.ParameterAsIdentifier("domain"),
	// SES domain identities can be imported using the domain name.
	"aws_ses_domain_identity": config.IdentifierFromProvider,
	// MAIL FROM domain can be imported using the domain attribute
	"aws_ses_domain_mail_from": config.IdentifierFromProvider,
	// SES email identities can be imported using the email address.
	"aws_ses_email_identity": config.IdentifierFromProvider,
	// SES event destinations can be imported using configuration_set_name together with the event destination's name
	// Example: some-configuration-set-test/event-destination-sns
	"aws_ses_event_destination": config.TemplatedStringAsIdentifier("name", "{{ .parameters.configuration_set_name }}/{{ .external_name }}"),
	// Identity Notification Topics can be imported using the ID of the record. The ID is made up as IDENTITY|TYPE where IDENTITY is the SES Identity and TYPE is the Notification Type.
	// Example: 'example.com|Bounce'
	"aws_ses_identity_notification_topic": config.IdentifierFromProvider,
	// SES Identity Policies can be imported using the identity and policy name, separated by a pipe character (|)
	// Example: 'example.com|example'
	"aws_ses_identity_policy": config.IdentifierFromProvider,
	// SES Receipt Filter can be imported using their name
	"aws_ses_receipt_filter": config.NameAsIdentifier,
	// SES receipt rules can be imported using the ruleset name and rule name separated by :
	// Example: my_rule_set:my_rule
	"aws_ses_receipt_rule": config.IdentifierFromProvider,
	// SES receipt rule sets can be imported using the rule set name
	"aws_ses_receipt_rule_set": config.IdentifierFromProvider,
	// SES templates can be imported using the template name
	"aws_ses_template": config.NameAsIdentifier,

	// signer
	//
	// Signer signing jobs can be imported using the job_id
	"aws_signer_signing_job": config.IdentifierFromProvider,
	// Signer signing profile permission statements can be imported using profile_name/statement_id
	// Example: prod_profile_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK/ProdAccountStartSigningJobStatementId
	"aws_signer_signing_profile_permission": config.TemplatedStringAsIdentifier("", "{{ .parameters.profile_name }}/{{ .parameters.statement_id }}"),

	// simpledb
	//
	// SimpleDB Domains can be imported using the name
	"aws_simpledb_domain": config.NameAsIdentifier,

	// networkfirewall
	//
	// Network Firewall Policies can be imported using their ARN
	// Example: arn:aws:network-firewall:us-west-1:123456789012:firewall-policy/example
	"aws_networkfirewall_firewall_policy": config.TemplatedStringAsIdentifier("name", "arn:aws:network-firewall:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:firewall-policy/{{ .external_name }}"),
	// Network Firewall Rule Groups can be imported using their ARN
	// Example: arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
	"aws_networkfirewall_rule_group": config.TemplatedStringAsIdentifier("", "arn:aws:network-firewall:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:stateful-rulegroup/{{ .external_name }}"),

	// networkmanager
	//
	// aws_networkmanager_global_network can be imported using the global network ID
	"aws_networkmanager_global_network": config.IdentifierFromProvider,
	// aws_networkmanager_site can be imported using the site ARN
	// Example: arn:aws:networkmanager::123456789012:site/global-network-0d47f6t230mz46dy4/site-444555aaabbb11223
	"aws_networkmanager_site": config.IdentifierFromProvider,
	// aws_networkmanager_link can be imported using the link ARN
	// Example: arn:aws:networkmanager::123456789012:link/global-network-0d47f6t230mz46dy4/link-444555aaabbb11223
	"aws_networkmanager_link": config.IdentifierFromProvider,
	// aws_networkmanager_link_association can be imported using the global network ID, link ID and device ID
	// Example: global-network-0d47f6t230mz46dy4,link-444555aaabbb11223,device-07f6fd08867abc123
	"aws_networkmanager_link_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.global_network_id }},{{ .parameters.link_id }},{{ .parameters.device_id }}"),
	// aws_networkmanager_device can be imported using the device ARN
	// Example: arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/device-07f6fd08867abc123
	"aws_networkmanager_device": config.IdentifierFromProvider,
	// aws_networkmanager_connection can be imported using the connection ARN
	// Example: arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/connection-07f6fd08867abc123
	"aws_networkmanager_connection": config.IdentifierFromProvider,
	// aws_networkmanager_transit_gateway_registration can be imported using the global network ID and transit gateway ARN
	// Example: global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:transit-gateway/tgw-123abc05e04123abc
	"aws_networkmanager_transit_gateway_registration": config.TemplatedStringAsIdentifier("", "{{ .parameters.global_network_id }},{{ .parameters.transit_gateway_arn }}"),
	// aws_networkmanager_transit_gateway_connect_peer_association can be imported using the global network ID and customer gateway ARN
	// Example: global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:transit-gateway-connect-peer/tgw-connect-peer-12345678
	"aws_networkmanager_transit_gateway_connect_peer_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.global_network_id }},arn:aws:ec2:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:transit-gateway-connect-peer/{{ .parameters.transit_gateway_connect_peer_arn }}"),
	// aws_networkmanager_customer_gateway_association can be imported using the global network ID and customer gateway ARN
	// Example: global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:customer-gateway/cgw-123abc05e04123abc
	"aws_networkmanager_customer_gateway_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.global_network_id }},arn:aws:ec2:{{ .parameters.region }}:{{ .setup.client_metadata.account_id }}:customer-gateway/{{ .parameters.customer_gateway_arn }}"),

	// waf
	//
	// WAF Byte Match Set can be imported using the id
	"aws_waf_byte_match_set": config.IdentifierFromProvider,
	// WAF Geo Match Set can be imported using their ID
	"aws_waf_geo_match_set": config.IdentifierFromProvider,
	// WAF IPSets can be imported using their ID
	"aws_waf_ipset": config.IdentifierFromProvider,
	// WAF Rated Based Rule can be imported using the id
	"aws_waf_rate_based_rule": config.IdentifierFromProvider,
	// WAF Regex Match Set can be imported using their ID
	"aws_waf_regex_match_set": config.IdentifierFromProvider,
	// AWS WAF Regex Pattern Set can be imported using their ID
	"aws_waf_regex_pattern_set": config.IdentifierFromProvider,
	// WAF rules can be imported using the id
	"aws_waf_rule": config.IdentifierFromProvider,
	// AWS WAF Size Constraint Set can be imported using their ID
	"aws_waf_size_constraint_set": config.IdentifierFromProvider,
	// AWS WAF SQL Injection Match Set can be imported using their ID
	"aws_waf_sql_injection_match_set": config.IdentifierFromProvider,
	// WAF Web ACL can be imported using the id
	"aws_waf_web_acl": config.IdentifierFromProvider,
	// WAF XSS Match Set can be imported using their ID
	"aws_waf_xss_match_set": config.IdentifierFromProvider,

	// wafregional
	//
	// WAF Regional Byte Match Set can be imported using the id
	"aws_wafregional_byte_match_set": config.IdentifierFromProvider,
	// WAF Regional Geo Match Set can be imported using the id
	"aws_wafregional_geo_match_set": config.IdentifierFromProvider,
	// WAF Regional IPSets can be imported using their ID
	"aws_wafregional_ipset": config.IdentifierFromProvider,
	// WAF Regional Rate Based Rule can be imported using the id
	"aws_wafregional_rate_based_rule": config.IdentifierFromProvider,
	// WAF Regional Regex Match Set can be imported using the id
	"aws_wafregional_regex_match_set": config.IdentifierFromProvider,
	// WAF Regional Regex Pattern Set can be imported using the id
	"aws_wafregional_regex_pattern_set": config.IdentifierFromProvider,
	// WAF Regional Rule can be imported using the id
	"aws_wafregional_rule": config.IdentifierFromProvider,
	// WAF Size Constraint Set can be imported using the id
	"aws_wafregional_size_constraint_set": config.IdentifierFromProvider,
	// WAF Regional Sql Injection Match Set can be imported using the id
	"aws_wafregional_sql_injection_match_set": config.IdentifierFromProvider,
	// WAF Regional Web ACL can be imported using the id
	"aws_wafregional_web_acl": config.IdentifierFromProvider,
	// AWS WAF Regional XSS Match can be imported using the id
	"aws_wafregional_xss_match_set": config.IdentifierFromProvider,

	// swf
	//
	// SWF Domains can be imported using the name
	"aws_swf_domain": config.NameAsIdentifier,

	// timestreamwrite
	//
	// Timestream databases can be imported using the database_name
	"aws_timestreamwrite_database": config.ParameterAsIdentifier("database_name"),
	// Timestream tables can be imported using the table_name and database_name separate by a colon (:)
	// Example: ExampleTable:ExampleDatabase
	"aws_timestreamwrite_table": config.TemplatedStringAsIdentifier("", "{{ .parameters.table_name }}:{{ .parameters.database_name }}"),

	// wafv2
	//
	// WAFv2 IP Sets can be imported using ID/name/scope
	"aws_wafv2_ip_set": config.IdentifierFromProvider,
	// WAFv2 Regex Pattern Sets can be imported using ID/name/scope
	"aws_wafv2_regex_pattern_set": config.IdentifierFromProvider,
	// WAFv2 Rule Group can be imported using ID/name/scope
	"aws_wafv2_rule_group": config.IdentifierFromProvider,
}

func lambdaFunctionURL() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, terraformProviderConfig map[string]interface{}) (string, error) {
		functionName, ok := parameters["function_name"]
		if !ok {
			return "", errors.New("function_name cannot be empty")
		}

		qualifier := parameters["qualifier"]
		if qualifier == nil || qualifier == "" {
			return functionName.(string), nil
		}
		return fmt.Sprintf("%s/%s", functionName.(string), qualifier.(string)), nil
	}
	return e
}

func iamUserGroupMembership() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		u, ok := parameters["user"]
		if !ok {
			return "", errors.New("user cannot be empty")
		}
		gs, ok := parameters["groups"]
		if !ok {
			return "", errors.New("groups cannot be empty")
		}
		var groups []string
		for _, g := range gs.([]interface{}) {
			groups = append(groups, g.(string))
		}
		return strings.Join(append([]string{u.(string)}, groups...), "/"), nil
	}
	return e
}

func kmsAlias() config.ExternalName {
	e := config.NameAsIdentifier
	e.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
		if _, ok := base["name"]; !ok {
			if !strings.HasPrefix(externalName, "alias/") {
				base["name"] = fmt.Sprintf("alias/%s", externalName)
			} else {
				base["name"] = externalName
			}
		}
	}
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id attribute missing from state file")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("value of id needs to be string")
		}

		return strings.TrimPrefix(idStr, "alias/"), nil
	}

	return e
}

func route() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		rtb, ok := parameters["route_table_id"]
		if !ok {
			return "", errors.New("route_table_id cannot be empty")
		}
		switch {
		case parameters["destination_cidr_block"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_cidr_block"].(string)), nil
		case parameters["destination_ipv6_cidr_block"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_ipv6_cidr_block"].(string)), nil
		case parameters["destination_prefix_list_id"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_prefix_list_id"].(string)), nil
		}
		return "", errors.New("destination_cidr_block or destination_ipv6_cidr_block or destination_prefix_list_id has to be given")
	}
	return e
}

func routeTableAssociation() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		rtb, ok := parameters["route_table_id"]
		if !ok {
			return "", errors.New("route_table_id cannot be empty")
		}
		switch {
		case parameters["subnet_id"] != nil:
			return fmt.Sprintf("%s/%s", parameters["subnet_id"].(string), rtb.(string)), nil
		case parameters["gateway_id"] != nil:
			return fmt.Sprintf("%s/%s", parameters["gateway_id"].(string), rtb.(string)), nil
		}
		return "", errors.New("gateway_id or subnet_id has to be given")
	}
	return e
}

func eksOIDCIdentityProvider() config.ExternalName {
	// OmittedFields in config.ExternalName works only for the top-level fields.
	// Hence, omitting is done in individual config override in `eks/config.go`
	return config.ExternalName{
		SetIdentifierArgumentFn: func(base map[string]interface{}, externalName string) {
			if _, ok := base["oidc"]; !ok {
				base["oidc"] = map[string]interface{}{}
			}
			// max length is 1:
			// https://github.com/hashicorp/terraform-provider-aws/blob/7ff39c5b11aafe812e3a4b414aa6d345286b95ec/internal/service/eks/identity_provider_config.go#L58
			if arr, ok := base["oidc"].([]interface{}); ok && len(arr) == 1 {
				if m, ok := arr[0].(map[string]interface{}); ok {
					m["identity_provider_config_name"] = externalName
				}
			}
		},
		GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
			if id, ok := tfstate["id"]; ok {
				return strings.Split(id.(string), ":")[1], nil
			}
			return "", errors.New("there is no id in tfstate")
		},
		GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			cl, ok := parameters["cluster_name"]
			if !ok {
				return "", errors.New("cluster_name cannot be empty")
			}
			return fmt.Sprintf("%s:%s", cl.(string), externalName), nil
		},
	}
}

// FormattedIdentifierFromProvider is a helper function to construct Terraform
// IDs that use elements from the parameters in a certain string format.
// It should be used in cases where all information in the ID is gathered from
// the spec and not user defined like name. For example, zone_id:vpc_id.
func FormattedIdentifierFromProvider(separator string, keys ...string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys))
		for i, key := range keys {
			val, ok := parameters[key]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", key)
			}
			s, ok := val.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be string", key)
			}
			vals[i] = s
		}
		return strings.Join(vals, separator), nil
	}
	return e
}

// FormattedIdentifierUserDefinedNameLast is used in cases where the ID is constructed
// using some of the spec fields as well as a field that users use to name the
// resource. For example, vpc_id:cluster_name where vpc_id comes from spec
// but cluster_name is a naming field we can use external name for.
// This function assumes that the naming field is the LAST component
// in the constructed identifier, which may not always hold
// (e.g., aws_servicecatalog_budget_resource_association).
func FormattedIdentifierUserDefinedNameLast(param, separator string, keys ...string) config.ExternalName {
	e := config.ParameterAsIdentifier(param)
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys)+1)
		for i, k := range keys {
			v, ok := parameters[k]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", k)
			}
			s, ok := v.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be a string", k)
			}
			vals[i] = s
		}
		vals[len(vals)-1] = externalName
		return strings.Join(vals, separator), nil
	}
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id in tfstate cannot be empty")
		}
		s, ok := id.(string)
		if !ok {
			return "", errors.New("value of id needs to be string")
		}
		w := strings.Split(s, separator)
		return w[len(w)-1], nil
	}
	return e
}

// FormattedIdentifierUserDefinedNameFirst is used in cases where the ID is constructed
// using some of the spec fields as well as a field that users use to name the
// resource. For example, budget_name:product_id where product_id comes from spec
// but budget_name is a naming field we can use external name for.
// This function assumes that the naming field is the FIRST component
// in the constructed identifier, which may not always hold
// (e.g., aws_eks_addon).
func FormattedIdentifierUserDefinedNameFirst(param, separator string, keys ...string) config.ExternalName {
	e := config.ParameterAsIdentifier(param)
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys)+1)
		for i, k := range keys {
			v, ok := parameters[k]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", k)
			}
			s, ok := v.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be a string", k)
			}
			vals[i+1] = s
		}
		vals[0] = externalName
		return strings.Join(vals, separator), nil
	}
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id in tfstate cannot be empty")
		}
		s, ok := id.(string)
		if !ok {
			return "", errors.New("value of id needs to be string")
		}
		w := strings.Split(s, separator)
		return w[0], nil
	}
	return e
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.Version = common.VersionV1Beta1
			r.ExternalName = e
		}
	}
}
