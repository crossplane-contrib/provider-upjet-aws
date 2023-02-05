/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/registry/reference"
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
			"apis/v1alpha1",
			"apis/v1beta1",
		},
		Controller: []string{
			"internal/controller/providerconfig",
			// "internal/controller/eks/clusterauth",
		},
	}
)

var skipList = []string{
	"aws_waf_rule_group$",              // Too big CRD schema
	"aws_wafregional_rule_group$",      // Too big CRD schema
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
	"aws_location_map$",                // failure with unknown reason.
	"aws_appflow_connector_profile$",   // failure with unknown reason.
}

// GetProvider returns provider configuration
func GetProvider() *config.Provider {
	modulePath := "github.com/dkb-bank/official-provider-aws"
	pc := config.NewProvider([]byte(providerSchema), "aws",
		modulePath, providerMetadata,
		config.WithShortName("aws"),
		config.WithRootGroup("aws.upbound.io"),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector("github.com/dkb-bank/official-provider-aws")}),
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
		// acm.Configure,
		// acmpca.Configure,
		// apigateway.Configure,
		// apigatewayv2.Configure,
		// apprunner.Configure,
		// appstream.Configure,
		// athena.Configure,
		// autoscaling.Configure,
		// backup.Configure,
		// cloudfront.Configure,
		// cloudsearch.Configure,
		// cloudwatch.Configure,
		// cloudwatchlogs.Configure,
		// cognitoidentity.Configure,
		// cognitoidp.Configure,
		// connect.Configure,
		// cur.Configure,
		// dax.Configure,
		// devicefarm.Configure,
		// docdb.Configure,
		// dynamodb.Configure,
		// ebs.Configure,
		// ec2.Configure,
		// ecr.Configure,
		// ecrpublic.Configure,
		// ecs.Configure,
		// efs.Configure,
		// eks.Configure,
		// elasticache.Configure,
		// elasticloadbalancing.Configure,
		// elb.Configure,
		// elbv2.Configure,
		// firehose.Configure,
		// gamelift.Configure,
		// globalaccelerator.Configure,
		// glue.Configure,
		// grafana.Configure,
		// iam.Configure,
		// kafka.Configure,
		// kinesis.Configure,
		// kinesisanalytics.Configure,
		// kinesisanalytics2.Configure,
		// kms.Configure,
		// lakeformation.Configure,
		// lambda.Configure,
		// licensemanager.Configure,
		// mq.Configure,
		// neptune.Configure,
		// opensearch.Configure,
		// rds.Configure,
		// redshift.Configure,
		// route53.Configure,
		// route53resolver.Configure,
		// route53recoverycontrolconfig.Configure,
		// s3.Configure,
		// secretsmanager.Configure,
		// servicecatalog.Configure,
		// organization.Configure,
		// cloudwatchevents.Configure,
		// budgets.Configure,
		// servicediscovery.Configure,
		// sfn.Configure,
		// sns.Configure,
		// sqs.Configure,
		// transfer.Configure,
		// directconnect.Configure,
		// ds.Configure,
		// qldb.Configure,
		// fsx.Configure,
		// networkmanager.Configure,
		// opsworks.Configure,
		// sagemaker.Configure,
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
