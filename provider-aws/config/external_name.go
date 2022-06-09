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

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// autoscaling
	//
	"aws_autoscaling_group": config.NameAsIdentifier,
	// No terraform import.
	"aws_autoscaling_attachment": config.IdentifierFromProvider,

	// ebs
	//
	// EBS Volumes can be imported using the id: vol-049df61146c4d7901
	"aws_ebs_volume": config.IdentifierFromProvider,

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
	// Imported using the following format: ROUTETABLEID_DESTINATION
	"aws_route": route(),
	// Imported using id: rtb-4e616f6d69
	"aws_route_table": config.IdentifierFromProvider,
	// Imported using the associated resource ID and Route Table ID separated
	// by a forward slash (/)
	"aws_route_table_association": routeTableAssociation(),
	// No import.
	"aws_main_route_table_association": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Route Table identifier, an
	// underscore, and the EC2 Transit Gateway Attachment identifier:
	// tgw-rtb-12345678_tgw-attach-87654321
	"aws_ec2_transit_gateway_route_table_propagation": FormattedIdentifierFromProvider("_", "transit_gateway_attachment_id", "transit_gateway_route_table_id"),
	// Imported using the id: igw-c0a643a9
	"aws_internet_gateway": config.IdentifierFromProvider,

	// ecr
	//
	"aws_ecr_repository": config.NameAsIdentifier,

	// ecrpublic
	//
	"aws_ecrpublic_repository": ParameterAsExternalName("repository_name"),

	// ecs
	//
	"aws_ecs_cluster":           config.NameAsIdentifier,
	"aws_ecs_service":           config.NameAsIdentifier,
	"aws_ecs_capacity_provider": config.NameAsIdentifier,
	// Imported using ARN: arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123
	"aws_ecs_task_definition": config.IdentifierFromProvider,

	// eks
	//
	"aws_eks_cluster": config.NameAsIdentifier,
	// Imported using the cluster_name and node_group_name separated by a
	// colon (:): my_cluster:my_node_group
	"aws_eks_node_group": FormattedIdentifierUserDefined("node_group_name", ":", "cluster_name"),
	// my_cluster:my_eks_addon
	"aws_eks_addon": FormattedIdentifierUserDefined("addon_name", ":", "cluster_name"),
	// my_cluster:my_fargate_profile
	"aws_eks_fargate_profile": FormattedIdentifierUserDefined("fargate_profile_name", ":", "cluster_name"),

	// elasticache
	//
	"aws_elasticache_parameter_group":   config.NameAsIdentifier,
	"aws_elasticache_subnet_group":      config.NameAsIdentifier,
	"aws_elasticache_cluster":           ParameterAsExternalName("cluster_id"),
	"aws_elasticache_replication_group": ParameterAsExternalName("replication_group_id"),
	"aws_elasticache_user":              ParameterAsExternalName("user_id"),
	"aws_elasticache_user_group":        ParameterAsExternalName("user_group_id"),

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

	// iam
	//
	// AKIA1234567890
	"aws_iam_access_key":       config.IdentifierFromProvider,
	"aws_iam_instance_profile": config.NameAsIdentifier,
	// arn:aws:iam::123456789012:policy/UsersManageOwnCredentials
	"aws_iam_policy": config.IdentifierFromProvider,
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

	// kms
	//
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	"aws_kms_key": config.IdentifierFromProvider,

	// mq
	//
	// a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
	"aws_mq_broker": config.IdentifierFromProvider,
	// c-0187d1eb-88c8-475a-9b79-16ef5a10c94f
	"aws_mq_configuration": config.IdentifierFromProvider,

	// neptune
	//
	"aws_neptune_cluster": ParameterAsExternalName("cluster_identifier"),
	// my_cluster:my_cluster_endpoint
	"aws_neptune_cluster_endpoint":        FormattedIdentifierUserDefined("cluster_endpoint_identifier", ":", "cluster_identifier"),
	"aws_neptune_cluster_instance":        ParameterAsExternalName("identifier"),
	"aws_neptune_cluster_parameter_group": config.NameAsIdentifier,
	"aws_neptune_cluster_snapshot":        ParameterAsExternalName("db_cluster_snapshot_identifier"),
	"aws_neptune_event_subscription":      config.NameAsIdentifier,
	"aws_neptune_parameter_group":         config.NameAsIdentifier,
	"aws_neptune_subnet_group":            config.NameAsIdentifier,

	// rds
	//
	"aws_rds_cluster":        ParameterAsExternalName("cluster_identifier"),
	"aws_db_instance":        ParameterAsExternalName("identifier"),
	"aws_db_parameter_group": config.NameAsIdentifier,
	"aws_db_subnet_group":    config.NameAsIdentifier,

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
	"aws_route53_key_signing_key": FormattedIdentifierUserDefined("name", ",", "hosted_zone_id"),
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	"aws_route53_query_log": config.IdentifierFromProvider,
	// Imported using ID of the record, which is the zone identifier, record
	// name, and record type, separated by underscores (_)
	// Z4KAPRWWNC7JR_dev.example.com_NS
	"aws_route53_record": config.IdentifierFromProvider,
	// Z123456ABCDEFG:vpc-12345678
	"aws_route53_vpc_association_authorization": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	// Z1D633PJN98FT9
	"aws_route53_zone": config.IdentifierFromProvider,
	// Z123456ABCDEFG:vpc-12345678
	"aws_route53_zone_association": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	// Imported using the id and version, e.g.,
	// 01a52019-d16f-422a-ae72-c306d2b6df7e/1
	"aws_route53_traffic_policy": config.IdentifierFromProvider,

	// route53resolver
	//
	// rdsc-be1866ecc1683e95
	"aws_route53_resolver_dnssec_config": config.IdentifierFromProvider,
	// rslvr-in-abcdef01234567890
	"aws_route53_resolver_endpoint": config.IdentifierFromProvider,
	// rdsc-be1866ecc1683e95
	"aws_route53_resolver_firewall_config": config.IdentifierFromProvider,
	// rslvr-fdl-0123456789abcdef
	"aws_route53_resolver_firewall_domain_list": config.IdentifierFromProvider,
	// Imported using the Route 53 Resolver DNS Firewall rule group ID and
	// domain list ID separated by ':', e.g.,
	// rslvr-frg-0123456789abcdef:rslvr-fdl-0123456789abcdef
	"aws_route53_resolver_firewall_rule": config.IdentifierFromProvider,
	// rslvr-frg-0123456789abcdef
	"aws_route53_resolver_firewall_rule_group": config.IdentifierFromProvider,
	// rslvr-frgassoc-0123456789abcdef
	"aws_route53_resolver_firewall_rule_group_association": config.IdentifierFromProvider,
	// rqlc-92edc3b1838248bf
	"aws_route53_resolver_query_log_config": config.IdentifierFromProvider,
	// rqlca-b320624fef3c4d70
	"aws_route53_resolver_query_log_config_association": config.IdentifierFromProvider,
	// rslvr-rr-0123456789abcdef0
	"aws_route53_resolver_rule": config.IdentifierFromProvider,
	// rslvr-rrassoc-97242eaf88example
	"aws_route53_resolver_rule_association": config.IdentifierFromProvider,

	// s3
	//
	"aws_s3_bucket": ParameterAsExternalName("bucket"),
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
		groups, ok := gs.([]string)
		if !ok {
			return "", errors.New("groups field needs to be an array of strings")
		}
		return strings.Join(append([]string{u.(string)}, groups...), "/"), nil
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

// FormattedIdentifierUserDefined is used in cases where the ID is constructed
// using some of the spec fields as well as a field that users use to name the
// resource. For example, vpc_id:cluster_name where vpc_id comes from spec
// but cluster_name is a naming field we can use external name for.
func FormattedIdentifierUserDefined(param, separator string, keys ...string) config.ExternalName {
	e := ParameterAsExternalName(param)
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

// ParameterAsExternalName is a different version of NameAsIdentifier where you
// can define a field name other than "name", such as "cluster_name".
func ParameterAsExternalName(paramName string) config.ExternalName {
	e := config.NameAsIdentifier
	e.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
		base[paramName] = externalName
	}
	e.OmittedFields = []string{
		paramName,
		paramName + "_prefix",
	}
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
