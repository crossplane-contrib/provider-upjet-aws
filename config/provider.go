/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/registry/reference"

	"github.com/upbound/provider-aws/config/acm"
	"github.com/upbound/provider-aws/config/acmpca"
	"github.com/upbound/provider-aws/config/apigateway"
	"github.com/upbound/provider-aws/config/apigatewayv2"
	"github.com/upbound/provider-aws/config/athena"
	"github.com/upbound/provider-aws/config/autoscaling"
	"github.com/upbound/provider-aws/config/backup"
	"github.com/upbound/provider-aws/config/cloudfront"
	"github.com/upbound/provider-aws/config/cloudsearch"
	"github.com/upbound/provider-aws/config/cloudwatch"
	"github.com/upbound/provider-aws/config/cloudwatchlogs"
	"github.com/upbound/provider-aws/config/cognitoidentity"
	"github.com/upbound/provider-aws/config/cognitoidp"
	"github.com/upbound/provider-aws/config/connect"
	"github.com/upbound/provider-aws/config/dax"
	"github.com/upbound/provider-aws/config/docdb"
	"github.com/upbound/provider-aws/config/dynamodb"
	"github.com/upbound/provider-aws/config/ebs"
	"github.com/upbound/provider-aws/config/ec2"
	"github.com/upbound/provider-aws/config/ecr"
	"github.com/upbound/provider-aws/config/ecrpublic"
	"github.com/upbound/provider-aws/config/ecs"
	"github.com/upbound/provider-aws/config/efs"
	"github.com/upbound/provider-aws/config/eks"
	"github.com/upbound/provider-aws/config/elasticache"
	"github.com/upbound/provider-aws/config/elasticloadbalancing"
	"github.com/upbound/provider-aws/config/elb"
	"github.com/upbound/provider-aws/config/elbv2"
	"github.com/upbound/provider-aws/config/firehose"
	"github.com/upbound/provider-aws/config/gamelift"
	"github.com/upbound/provider-aws/config/globalaccelerator"
	"github.com/upbound/provider-aws/config/glue"
	"github.com/upbound/provider-aws/config/grafana"
	"github.com/upbound/provider-aws/config/iam"
	"github.com/upbound/provider-aws/config/kafka"
	"github.com/upbound/provider-aws/config/kinesis"
	"github.com/upbound/provider-aws/config/kinesisanalytics"
	kinesisanalytics2 "github.com/upbound/provider-aws/config/kinesisanalyticsv2"
	"github.com/upbound/provider-aws/config/kms"
	"github.com/upbound/provider-aws/config/lakeformation"
	"github.com/upbound/provider-aws/config/lambda"
	"github.com/upbound/provider-aws/config/licensemanager"
	"github.com/upbound/provider-aws/config/mq"
	"github.com/upbound/provider-aws/config/neptune"
	"github.com/upbound/provider-aws/config/opensearch"
	"github.com/upbound/provider-aws/config/rds"
	"github.com/upbound/provider-aws/config/redshift"
	"github.com/upbound/provider-aws/config/route53"
	"github.com/upbound/provider-aws/config/route53resolver"
	"github.com/upbound/provider-aws/config/s3"
	"github.com/upbound/provider-aws/config/secretsmanager"
	"github.com/upbound/provider-aws/config/servicediscovery"
	"github.com/upbound/provider-aws/config/sfn"
	"github.com/upbound/provider-aws/config/sns"
	"github.com/upbound/provider-aws/config/sqs"
	"github.com/upbound/provider-aws/config/transfer"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

var (
	BasePackages = config.BasePackages{
		APIVersion: []string{
			"apis/v1beta1",
		},
		Controller: []string{
			"internal/controller/providerconfig",
			"internal/controller/eks/clusterauth",
		},
	}
)

var skipList = []string{
	"aws_waf_rule_group$",              // Too big CRD schema
	"aws_wafregional_rule_group$",      // Too big CRD schema
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
	modulePath := "github.com/upbound/provider-aws"
	pc := config.NewProvider([]byte(providerSchema), "aws",
		modulePath, providerMetadata,
		config.WithShortName("aws"),
		config.WithRootGroup("aws.upbound.io"),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector("github.com/upbound/provider-aws")}),
		config.WithIncludeList(ResourcesWithExternalNameConfig()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithSkipList(skipList),
		config.WithBasePackages(BasePackages),
		config.WithDefaultResourceOptions(
			GroupKindOverrides(),
			KindOverrides(),
			RegionAddition(),
			TagsAllRemoval(),
			IdentifierAssignedByAWS(),
			KnownReferencers(),
			AddExternalTagsField(),
			ExternalNameConfigurations(),
			NamePrefixRemoval(),
			DocumentationForTags(),
		),
	)

	for _, configure := range []func(provider *config.Provider){
		acm.Configure,
		acmpca.Configure,
		autoscaling.Configure,
		cognitoidentity.Configure,
		cognitoidp.Configure,
		connect.Configure,
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
		lakeformation.Configure,
		route53resolver.Configure,
		dax.Configure,
		apigatewayv2.Configure,
		cloudsearch.Configure,
		apigateway.Configure,
		elbv2.Configure,
		sqs.Configure,
		opensearch.Configure,
		secretsmanager.Configure,
		cloudwatch.Configure,
		cloudwatchlogs.Configure,
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
