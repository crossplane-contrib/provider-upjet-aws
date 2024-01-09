/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"context"
	"reflect"
	"unsafe"

	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/xpfwprovider"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"github.com/pkg/errors"

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
	"github.com/upbound/provider-aws/config/identitystore"
	"github.com/upbound/provider-aws/config/iot"
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
	"github.com/upbound/provider-aws/config/ssoadmin"
	"github.com/upbound/provider-aws/config/transfer"
	"github.com/upbound/provider-aws/hack"
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

// workaround for the TF AWS v4.67.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns the provider configuration.
// The `generationProvider` argument specifies whether the provider
// configuration is being read for the code generation pipelines.
// In that case, we will only use the JSON schema for generating
// the CRDs.
func GetProvider(ctx context.Context, generationProvider bool) (*config.Provider, *xpprovider.AWSClient, error) {
	var p *schema.Provider
	var fwProvider fwprovider.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
		fwProvider, _, _ = xpfwprovider.GetProvider(ctx)
	} else {
		// p, err = xpprovider.GetProviderSchema(ctx)
		fwProvider, p, err = xpfwprovider.GetProvider(ctx)
	}
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}
	// we set schema.Provider's meta to nil because p.Configure modifies
	// a singleton pointer. This further assumes that the
	// schema.Provider.Configure calls do not modify the global state
	// represented by the config.Provider.TerraformProvider.
	var awsClient *xpprovider.AWSClient
	if !generationProvider {
		// #nosec G103
		awsClient = (*xpprovider.AWSClient)(unsafe.Pointer(reflect.ValueOf(p.Meta()).Pointer()))
	}

	modulePath := "github.com/upbound/provider-aws"
	pc := config.NewProvider(ctx, []byte(providerSchema), "aws",
		modulePath, providerMetadata,
		config.WithShortName("aws"),
		config.WithRootGroup("aws.upbound.io"),
		config.WithIncludeList(CLIReconciledResourceList()),
		config.WithNoForkIncludeList(NoForkResourceList()),
		config.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithSkipList(skipList),
		config.WithFeaturesPackage("internal/features"),
		config.WithMainTemplate(hack.MainTemplate),
		config.WithTerraformProvider(p),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
		config.WithDefaultResourceOptions(
			GroupKindOverrides(),
			KindOverrides(),
			RegionAddition(),
			TagsAllRemoval(),
			IdentifierAssignedByAWS(),
			KnownReferencers(),
			AddExternalTagsField(),
			ResourceConfigurator(),
			NamePrefixRemoval(),
			DocumentationForTags(),
		),
	)
	p.SetMeta(nil)
	pc.BasePackages.ControllerMap["internal/controller/eks/clusterauth"] = "eks"

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
		ssoadmin.Configure,
		identitystore.Configure,
		iot.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, awsClient, nil
}

// CLIReconciledResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF CLI based architecture.
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// NoForkResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func NoForkResourceList() []string {
	l := make([]string, len(NoForkExternalNameConfigs))
	i := 0
	for name := range NoForkExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

func TerraformPluginFrameworkResourceList() []string {
	l := make([]string, len(TerraformPluginFrameworkExternalNameConfigs))
	i := 0
	for name := range TerraformPluginFrameworkExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}
