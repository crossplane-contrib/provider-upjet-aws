/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/acm"
	"github.com/upbound/official-providers/provider-aws/config/acmpca"
	"github.com/upbound/official-providers/provider-aws/config/athena"
	"github.com/upbound/official-providers/provider-aws/config/autoscaling"
	"github.com/upbound/official-providers/provider-aws/config/backup"
	"github.com/upbound/official-providers/provider-aws/config/cloudfront"
	"github.com/upbound/official-providers/provider-aws/config/docdb"
	"github.com/upbound/official-providers/provider-aws/config/dynamodb"
	"github.com/upbound/official-providers/provider-aws/config/ebs"
	"github.com/upbound/official-providers/provider-aws/config/ec2"
	"github.com/upbound/official-providers/provider-aws/config/ecr"
	"github.com/upbound/official-providers/provider-aws/config/ecrpublic"
	"github.com/upbound/official-providers/provider-aws/config/ecs"
	"github.com/upbound/official-providers/provider-aws/config/efs"
	"github.com/upbound/official-providers/provider-aws/config/eks"
	"github.com/upbound/official-providers/provider-aws/config/elasticache"
	"github.com/upbound/official-providers/provider-aws/config/elasticloadbalancing"
	"github.com/upbound/official-providers/provider-aws/config/elb"
	"github.com/upbound/official-providers/provider-aws/config/firehose"
	"github.com/upbound/official-providers/provider-aws/config/gamelift"
	"github.com/upbound/official-providers/provider-aws/config/globalaccelerator"
	"github.com/upbound/official-providers/provider-aws/config/glue"
	"github.com/upbound/official-providers/provider-aws/config/grafana"
	"github.com/upbound/official-providers/provider-aws/config/iam"
	"github.com/upbound/official-providers/provider-aws/config/kafka"
	"github.com/upbound/official-providers/provider-aws/config/kinesis"
	"github.com/upbound/official-providers/provider-aws/config/kinesisanalytics"
	kinesisanalytics2 "github.com/upbound/official-providers/provider-aws/config/kinesisanalyticsv2"
	"github.com/upbound/official-providers/provider-aws/config/kms"
	"github.com/upbound/official-providers/provider-aws/config/lambda"
	"github.com/upbound/official-providers/provider-aws/config/licensemanager"
	"github.com/upbound/official-providers/provider-aws/config/mq"
	"github.com/upbound/official-providers/provider-aws/config/neptune"
	"github.com/upbound/official-providers/provider-aws/config/rds"
	"github.com/upbound/official-providers/provider-aws/config/redshift"
	"github.com/upbound/official-providers/provider-aws/config/route53"
	"github.com/upbound/official-providers/provider-aws/config/s3"
	"github.com/upbound/official-providers/provider-aws/config/servicediscovery"
	"github.com/upbound/official-providers/provider-aws/config/sfn"
	"github.com/upbound/official-providers/provider-aws/config/sns"
	"github.com/upbound/official-providers/provider-aws/config/transfer"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

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
	"aws_iot_authorizer$",              // failure with unknown reason.
	"aws_appflow_connector_profile$",   // failure with unknown reason.
	"aws_location_map$",                // failure with unknown reason.
}

// GetProvider returns provider configuration
func GetProvider() *config.Provider {
	pc := config.NewProvider([]byte(providerSchema), "aws",
		"github.com/upbound/official-providers/provider-aws", providerMetadata,
		config.WithShortName("aws"),
		config.WithRootGroup("aws.upbound.io"),
		config.WithIncludeList(ResourcesWithExternalNameConfig()),
		config.WithSkipList(skipList),
		config.WithDefaultResourceOptions(
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

	for _, configure := range []func(provider *config.Provider){
		acm.Configure,
		acmpca.Configure,
		autoscaling.Configure,
		dynamodb.Configure,
		ebs.Configure,
		ec2.Configure,
		ecr.Configure,
		ecrpublic.Configure,
		ecs.Configure,
		eks.Configure,
		elasticache.Configure,
		elasticloadbalancing.Configure,
		globalaccelerator.Configure,
		glue.Configure,
		iam.Configure,
		rds.Configure,
		s3.Configure,
		route53.Configure,
		neptune.Configure,
		mq.Configure,
		cloudfront.Configure,
		servicediscovery.Configure,
		efs.Configure,
		sns.Configure,
		backup.Configure,
		grafana.Configure,
		gamelift.Configure,
		kinesis.Configure,
		kinesisanalytics.Configure,
		kinesisanalytics2.Configure,
		firehose.Configure,
		licensemanager.Configure,
		kms.Configure,
		lambda.Configure,
		athena.Configure,
		docdb.Configure,
		elb.Configure,
		redshift.Configure,
		sfn.Configure,
		transfer.Configure,
		kafka.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// ResourcesWithExternalNameConfig returns the list of resources that have external
// name configured in ExternalNameConfigs table.
func ResourcesWithExternalNameConfig() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}
