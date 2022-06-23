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

	// ACM
	// Imported using ARN: arn:aws:acm:eu-central-1:123456789012:certificate/7e7a28d2-163f-4b8f-b9cd-822f96c08d6a
	"aws_acm_certificate": config.IdentifierFromProvider,
	// No import documented, but https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate_validation#id
	"aws_acm_certificate_validation": config.IdentifierFromProvider,

	// ACM PCA
	// aws_acmpca_certificate can not be imported at this time.
	"aws_acmpca_certificate": config.IdentifierFromProvider,
	// Imported using ARN: arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
	"aws_acmpca_certificate_authority": config.IdentifierFromProvider,
	// No doc on import, but resource is getting CA ARN: arn:aws:acm-pca:eu-central-1:609897127049:certificate-authority/ba0c7989-9641-4f36-a033-dee60121d595
	"aws_acmpca_certificate_authority_certificate": config.IdentifierFromProvider,

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
	"aws_ecrpublic_repository": ParameterAsExternalName("repository_name"),
	// Imported using the repository name.
	"aws_ecrpublic_repository_policy": config.IdentifierFromProvider,

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
	// It has a complex config, adding empty entry here just to enable it.
	"aws_eks_identity_provider_config": eksOIDCIdentityProvider(),

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
	// df579d9a-6396-410e-ac22-e7ad60cf9e7e
	"aws_route53_traffic_policy_instance": config.IdentifierFromProvider,

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
	// S3 bucket can be imported using the bucket
	"aws_s3_bucket": ParameterAsExternalName("bucket"),
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
	// No import documented
	"aws_s3_object_copy": config.IdentifierFromProvider,

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

func eksOIDCIdentityProvider() config.ExternalName {
	// OmittedFields in config.ExternalName works only for the top-level fields.
	// Hence, omitting is done in individual config override in `eks/config.go`
	return config.ExternalName{
		SetIdentifierArgumentFn: func(base map[string]interface{}, externalName string) {
			if _, ok := base["oidc"]; !ok {
				base["oidc"] = map[string]interface{}{}
			}
			if m, ok := base["oidc"].(map[string]interface{}); ok {
				m["identity_provider_config_name"] = externalName
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
