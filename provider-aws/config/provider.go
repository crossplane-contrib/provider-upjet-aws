/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/acm"
	"github.com/upbound/official-providers/provider-aws/config/acmpca"
	"github.com/upbound/official-providers/provider-aws/config/autoscaling"
	"github.com/upbound/official-providers/provider-aws/config/cloudfront"
	"github.com/upbound/official-providers/provider-aws/config/ebs"
	"github.com/upbound/official-providers/provider-aws/config/ec2"
	"github.com/upbound/official-providers/provider-aws/config/ecr"
	"github.com/upbound/official-providers/provider-aws/config/ecrpublic"
	"github.com/upbound/official-providers/provider-aws/config/ecs"
	"github.com/upbound/official-providers/provider-aws/config/eks"
	"github.com/upbound/official-providers/provider-aws/config/elasticache"
	"github.com/upbound/official-providers/provider-aws/config/elasticloadbalancing"
	"github.com/upbound/official-providers/provider-aws/config/globalaccelerator"
	"github.com/upbound/official-providers/provider-aws/config/iam"
	"github.com/upbound/official-providers/provider-aws/config/mq"
	"github.com/upbound/official-providers/provider-aws/config/neptune"
	"github.com/upbound/official-providers/provider-aws/config/rds"
	"github.com/upbound/official-providers/provider-aws/config/route53"
	"github.com/upbound/official-providers/provider-aws/config/s3"
)

//go:embed schema.json
var providerSchema string

// IncludedResources lists all resource patterns included in small set release.
var IncludedResources = []string{
	// Elastic Load Balancing v2 (ALB/NLB)
	"aws_lb$",
	"aws_lb_listener$",
	"aws_lb_target_group$",
	"aws_lb_target_group_attachment$",

	// ecr
	"aws_ecr_*.",

	// ecrpublic
	"aws_ecrpublic_*.$",

	// RDS
	"aws_rds_cluster$",
	"aws_db_instance$",
	"aws_db_parameter_group$",
	"aws_db_subnet_group$",

	// S3
	"aws_s3_bucket.*",
	"aws_s3_object.*",

	// Elasticache
	"aws_elasticache_cluster$",
	"aws_elasticache_subnet_group$",
	"aws_elasticache_parameter_group$",
	"aws_elasticache_replication_group$",
	"aws_elasticache_user$",
	"aws_elasticache_user_group$",

	// ECS
	"aws_ecs_cluster$",
	"aws_ecs_service$",
	"aws_ecs_capacity_provider$",
	"aws_ecs_tag$",
	"aws_ecs_task_definition$",

	// Autoscaling
	"aws_autoscaling_group$",
	"aws_autoscaling_attachment$",

	// EC2
	"aws_instance$",
	"aws_eip$",
	"aws_launch_template$",
	"aws_ec2_transit_gateway$",
	"aws_ec2_transit_gateway_route$",
	"aws_ec2_transit_gateway_route_table$",
	"aws_ec2_transit_gateway_route_table_association$",
	"aws_ec2_transit_gateway_vpc_attachment$",
	"aws_ec2_transit_gateway_vpc_attachment_accepter$",
	"aws_ec2_transit_gateway_route_table_propagation$",
	"aws_vpc$",
	"aws_security_group$",
	"aws_security_group_rule$",
	"aws_subnet$",
	"aws_network_interface$",
	"aws_route$",
	"aws_route_table$",
	"aws_vpc_endpoint$",
	"aws_vpc_ipv4_cidr_block_association$",
	"aws_vpc_peering_connection$",
	"aws_route_table_association$",
	"aws_internet_gateway$",

	// IAM
	"aws_iam_access_key$",
	"aws_iam_group$",
	"aws_iam_group_policy$",
	"aws_iam_group_policy_attachment$",
	"aws_iam_instance_profile$",
	"aws_iam_policy$",
	"aws_iam_policy_attachment$",
	"aws_iam_role$",
	"aws_iam_role_policy$",
	"aws_iam_role_policy_attachment$",
	"aws_iam_user$",
	"aws_iam_user_group_membership$",
	"aws_iam_user_policy$",
	"aws_iam_user_policy_attachment$",
	"aws_iam_openid_connect_provider$",

	// EKS
	"aws_eks_addon$",
	"aws_eks_cluster$",
	"aws_eks_fargate_profile$",
	"aws_eks_node_group$",
	"aws_eks_identity_provider_config$",

	// KMS
	"aws_kms_key$",

	// EBS
	"aws_ebs_volume$",

	// Route53
	"aws_route53_.*",

	// Neptune
	"aws_neptune_cluster$",
	"aws_neptune_cluster_endpoint$",
	"aws_neptune_cluster_instance$",
	"aws_neptune_cluster_parameter_group$",
	"aws_neptune_cluster_snapshot$",
	"aws_neptune_event_subscription",
	"aws_neptune_parameter_group$",
	"aws_neptune_subnet_group$",

	// MQ
	"aws_mq_broker$",
	"aws_mq_configuration$",

	// Global Accelerator
	"aws_globalaccelerator_accelerator",
	"aws_globalaccelerator_endpoint_group",
	"aws_globalaccelerator_listener",

	// ACM (Certificate Manager)
	"aws_acm_.+",

	// ACM PCA (Certificate Manager Private Certificate Authority )
	"aws_acmpca_.+",

	// Cloudfront
	"aws_cloudfront.*",
}

var skipList = []string{
	"aws_waf_rule_group$",              // Too big CRD schema
	"aws_wafregional_rule_group$",      // Too big CRD schema
	"aws_glue_connection$",             // See https://github.com/crossplane-contrib/terrajet/issues/100
	"aws_mwaa_environment$",            // See https://github.com/crossplane-contrib/terrajet/issues/100
	"aws_ecs_tag$",                     // tags are already managed by ecs resources.
	"aws_alb$",                         // identical with aws_lb
	"aws_alb_target_group_attachment$", // identical with aws_lb_target_group_attachment
	"aws_iam_policy_attachment$",       // identical with aws_iam_*_policy_attachment resources.
	"aws_iam_group_policy$",            // identical with aws_iam_*_policy_attachment resources.
	"aws_iam_role_policy$",             // identical with aws_iam_*_policy_attachment resources.
	"aws_iam_user_policy$",             // identical with aws_iam_*_policy_attachment resources.
	"aws_alb$",                         // identical with aws_lb.
	"aws_alb_listener$",                // identical with aws_lb_listener.
	"aws_alb_target_group$",            // identical with aws_lb_target_group.
	"aws_alb_target_group_attachment$", // identical with aws_lb_target_group_attachment.
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), "aws",
		"github.com/upbound/official-providers/provider-aws", "",
		tjconfig.WithShortName("aws"),
		tjconfig.WithRootGroup("aws.upbound.io"),
		tjconfig.WithIncludeList(IncludedResources),
		tjconfig.WithSkipList(skipList),
		tjconfig.WithDefaultResourceOptions(
			GroupKindOverrides(),
			KindOverrides(),
			RegionAddition(),
			TagsAllRemoval(),
			IdentifierAssignedByAWS(),
			NamePrefixRemoval(),
			KnownReferencers(),
			AddExternalTagsField(),
			ExternalNameConfigurations(),
		),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		acm.Configure,
		acmpca.Configure,
		autoscaling.Configure,
		ebs.Configure,
		ec2.Configure,
		ecr.Configure,
		ecrpublic.Configure,
		ecs.Configure,
		eks.Configure,
		elasticache.Configure,
		elasticloadbalancing.Configure,
		globalaccelerator.Configure,
		iam.Configure,
		rds.Configure,
		s3.Configure,
		route53.Configure,
		neptune.Configure,
		mq.Configure,
		cloudfront.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
