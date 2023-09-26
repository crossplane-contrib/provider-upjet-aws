/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"context"
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/registry/reference"

	"github.com/upbound/provider-aws/config/acm"
	"github.com/upbound/provider-aws/config/acmpca"
	"github.com/upbound/provider-aws/config/apigateway"
	"github.com/upbound/provider-aws/config/apigatewayv2"
	"github.com/upbound/provider-aws/config/apprunner"
	"github.com/upbound/provider-aws/config/appstream"
	"github.com/upbound/provider-aws/config/athena"
	"github.com/upbound/provider-aws/config/autoscaling"
	"github.com/upbound/provider-aws/config/backup"
	"github.com/upbound/provider-aws/config/budgets"
	"github.com/upbound/provider-aws/config/cloudfront"
	"github.com/upbound/provider-aws/config/cloudsearch"
	"github.com/upbound/provider-aws/config/cloudwatch"
	"github.com/upbound/provider-aws/config/cloudwatchevents"
	"github.com/upbound/provider-aws/config/cloudwatchlogs"
	"github.com/upbound/provider-aws/config/cognitoidentity"
	"github.com/upbound/provider-aws/config/cognitoidp"
	"github.com/upbound/provider-aws/config/connect"
	"github.com/upbound/provider-aws/config/cur"
	"github.com/upbound/provider-aws/config/datasync"
	"github.com/upbound/provider-aws/config/dax"
	"github.com/upbound/provider-aws/config/devicefarm"
	"github.com/upbound/provider-aws/config/directconnect"
	"github.com/upbound/provider-aws/config/dms"
	"github.com/upbound/provider-aws/config/docdb"
	"github.com/upbound/provider-aws/config/ds"
	"github.com/upbound/provider-aws/config/dynamodb"
	"github.com/upbound/provider-aws/config/ebs"
	"github.com/upbound/provider-aws/config/ec2"
	"github.com/upbound/provider-aws/config/ecr"
	"github.com/upbound/provider-aws/config/ecrpublic"
	"github.com/upbound/provider-aws/config/ecs"
	"github.com/upbound/provider-aws/config/efs"
	"github.com/upbound/provider-aws/config/eks"
	"github.com/upbound/provider-aws/config/elasticache"
	"github.com/upbound/provider-aws/config/elb"
	"github.com/upbound/provider-aws/config/elbv2"
	"github.com/upbound/provider-aws/config/firehose"
	"github.com/upbound/provider-aws/config/fsx"
	"github.com/upbound/provider-aws/config/gamelift"
	"github.com/upbound/provider-aws/config/globalaccelerator"
	"github.com/upbound/provider-aws/config/glue"
	"github.com/upbound/provider-aws/config/grafana"
	"github.com/upbound/provider-aws/config/iam"
	"github.com/upbound/provider-aws/config/kafka"
	"github.com/upbound/provider-aws/config/kendra"
	"github.com/upbound/provider-aws/config/kinesis"
	"github.com/upbound/provider-aws/config/kinesisanalytics"
	kinesisanalytics2 "github.com/upbound/provider-aws/config/kinesisanalyticsv2"
	"github.com/upbound/provider-aws/config/kms"
	"github.com/upbound/provider-aws/config/lakeformation"
	"github.com/upbound/provider-aws/config/lambda"
	"github.com/upbound/provider-aws/config/licensemanager"
	"github.com/upbound/provider-aws/config/medialive"
	"github.com/upbound/provider-aws/config/memorydb"
	"github.com/upbound/provider-aws/config/mq"
	"github.com/upbound/provider-aws/config/neptune"
	"github.com/upbound/provider-aws/config/networkfirewall"
	"github.com/upbound/provider-aws/config/networkmanager"
	"github.com/upbound/provider-aws/config/opensearch"
	"github.com/upbound/provider-aws/config/opsworks"
	"github.com/upbound/provider-aws/config/organization"
	"github.com/upbound/provider-aws/config/qldb"
	"github.com/upbound/provider-aws/config/ram"
	"github.com/upbound/provider-aws/config/rds"
	"github.com/upbound/provider-aws/config/redshift"
	"github.com/upbound/provider-aws/config/redshiftserverless"
	"github.com/upbound/provider-aws/config/rolesanywhere"
	"github.com/upbound/provider-aws/config/route53"
	"github.com/upbound/provider-aws/config/route53recoverycontrolconfig"
	"github.com/upbound/provider-aws/config/route53resolver"
	"github.com/upbound/provider-aws/config/s3"
	"github.com/upbound/provider-aws/config/sagemaker"
	"github.com/upbound/provider-aws/config/secretsmanager"
	"github.com/upbound/provider-aws/config/servicecatalog"
	"github.com/upbound/provider-aws/config/servicediscovery"
	"github.com/upbound/provider-aws/config/sfn"
	"github.com/upbound/provider-aws/config/sns"
	"github.com/upbound/provider-aws/config/sqs"
	"github.com/upbound/provider-aws/config/transfer"
	"github.com/upbound/provider-aws/hack"
)

var (
	//go:embed provider-metadata.yaml
	providerMetadata []byte
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
	"aws_iam_user_policy$",             // identical with aws_iam_*_policy_attachment resources.
	"aws_alb$",                         // identical with aws_lb.
	"aws_alb_listener$",                // identical with aws_lb_listener.
	"aws_alb_target_group$",            // identical with aws_lb_target_group.
	"aws_alb_target_group_attachment$", // identical with aws_lb_target_group_attachment.
	"aws_iot_authorizer$",              // failure with unknown reason.
	"aws_location_map$",                // failure with unknown reason.
	"aws_appflow_connector_profile$",   // failure with unknown reason.
	"aws_rds_reserved_instance",        // Expense of testing
}

// GetProvider returns provider configuration
func GetProvider(ctx context.Context) (*config.Provider, error) {
	p, err := xpprovider.GetProviderSchema(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get the Terraform provider schema")
	}
	modulePath := "github.com/upbound/provider-aws"
	pc := config.NewProvider(p.ResourcesMap, "aws",
		modulePath, providerMetadata,
		config.WithShortName("aws"),
		config.WithRootGroup("aws.upbound.io"),
		config.WithIncludeList(ResourcesWithExternalNameConfig()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithSkipList(skipList),
		config.WithFeaturesPackage("internal/features"),
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
			iamNoForkClientConfig(),
		),
		config.WithMainTemplate(hack.MainTemplate),
	)
	pc.BasePackages.ControllerMap["internal/controller/eks/clusterauth"] = "eks"
	pc.TerraformProvider = p

	for _, configure := range []func(provider *config.Provider){
		acm.Configure,
		acmpca.Configure,
		apigateway.Configure,
		apigatewayv2.Configure,
		apprunner.Configure,
		appstream.Configure,
		athena.Configure,
		autoscaling.Configure,
		backup.Configure,
		cloudfront.Configure,
		cloudsearch.Configure,
		cloudwatch.Configure,
		cloudwatchlogs.Configure,
		cognitoidentity.Configure,
		cognitoidp.Configure,
		connect.Configure,
		cur.Configure,
		datasync.Configure,
		dax.Configure,
		devicefarm.Configure,
		dms.Configure,
		docdb.Configure,
		dynamodb.Configure,
		ebs.Configure,
		ec2.Configure,
		ecr.Configure,
		ecrpublic.Configure,
		ecs.Configure,
		efs.Configure,
		eks.Configure,
		elasticache.Configure,
		elb.Configure,
		elbv2.Configure,
		firehose.Configure,
		gamelift.Configure,
		globalaccelerator.Configure,
		glue.Configure,
		grafana.Configure,
		iam.Configure,
		kafka.Configure,
		kinesis.Configure,
		kinesisanalytics.Configure,
		kinesisanalytics2.Configure,
		kms.Configure,
		lakeformation.Configure,
		lambda.Configure,
		licensemanager.Configure,
		memorydb.Configure,
		mq.Configure,
		neptune.Configure,
		networkfirewall.Configure,
		opensearch.Configure,
		ram.Configure,
		rds.Configure,
		redshift.Configure,
		rolesanywhere.Configure,
		route53.Configure,
		route53resolver.Configure,
		route53recoverycontrolconfig.Configure,
		s3.Configure,
		secretsmanager.Configure,
		servicecatalog.Configure,
		organization.Configure,
		cloudwatchevents.Configure,
		budgets.Configure,
		servicediscovery.Configure,
		sfn.Configure,
		sns.Configure,
		sqs.Configure,
		transfer.Configure,
		directconnect.Configure,
		ds.Configure,
		qldb.Configure,
		fsx.Configure,
		networkmanager.Configure,
		opsworks.Configure,
		sagemaker.Configure,
		redshiftserverless.Configure,
		kendra.Configure,
		medialive.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
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
