// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

module github.com/upbound/provider-aws/v2

go 1.26.4

tool golang.org/x/tools/cmd/goimports

require (
	dario.cat/mergo v1.0.2
	github.com/alecthomas/kingpin/v2 v2.4.0
	github.com/aws/aws-sdk-go-v2 v1.42.0
	github.com/aws/aws-sdk-go-v2/config v1.32.26
	github.com/aws/aws-sdk-go-v2/credentials v1.19.25
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.29
	github.com/aws/aws-sdk-go-v2/service/eks v1.87.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.43.4
	github.com/aws/smithy-go v1.27.2
	github.com/crossplane/crossplane-runtime/v2 v2.2.0
	github.com/crossplane/crossplane-tools v0.0.0-20250731192036-00d407d8b7ec
	github.com/crossplane/upjet/v2 v2.3.1-0.20260703134037-c73cc3a2d247
	github.com/go-ini/ini v1.46.0
	github.com/google/go-cmp v0.7.0
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.73
	github.com/hashicorp/awspolicyequivalence v1.7.0
	github.com/hashicorp/terraform-json v0.27.2
	github.com/hashicorp/terraform-plugin-framework v1.19.0
	github.com/hashicorp/terraform-plugin-go v0.31.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.40.1
	github.com/hashicorp/terraform-provider-aws v0.0.0-00010101000000-000000000000
	github.com/json-iterator/go v1.1.12
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.79.3
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.35.3
	k8s.io/apiextensions-apiserver v0.35.3
	k8s.io/apimachinery v0.35.3
	k8s.io/client-go v0.35.3
	k8s.io/utils v0.0.0-20260319190234-28399d86e0b5
	sigs.k8s.io/controller-runtime v0.23.3
	sigs.k8s.io/controller-tools v0.20.1
	sigs.k8s.io/randfill v1.0.0
)

require (
	cloud.google.com/go v0.123.0 // indirect
	github.com/ProtonMail/go-crypto v1.4.1 // indirect
	github.com/YakDriver/go-version v0.1.0 // indirect
	github.com/YakDriver/regexache v0.25.0 // indirect
	github.com/YakDriver/smarterr v0.8.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/units v0.0.0-20240927000941-0f3dac36c52b // indirect
	github.com/antchfx/htmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.3.6 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.13 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.22.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.30 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.49.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/account v1.32.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.47.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/amp v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/amplify v1.39.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.35.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.45.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/appfabric v1.17.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/appflow v1.52.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/appintegrations v1.38.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.43.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.35.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationsignals v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.36.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/appstream v1.61.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appsync v1.54.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/arcregionswitch v1.9.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/arczonalshift v1.23.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.58.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.47.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.67.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.31.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/backup v1.57.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/batch v1.66.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/bcmdataexports v1.16.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.64.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrockagent v1.56.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrockagentcorecontrol v1.45.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/billing v1.11.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/budgets v1.44.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/chatbot v1.15.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/chime v1.42.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines v1.27.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkvoice v1.29.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.45.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.34.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.30.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.72.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.65.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore v1.13.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.35.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.33.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.56.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.60.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.78.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.39.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.69.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecatalyst v1.22.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.34.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeconnections v1.11.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.36.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.30.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/codegurureviewer v1.35.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.47.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.36.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.32.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.34.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.62.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.41.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.54.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/configservice v1.64.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/connect v1.177.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/connectcases v1.42.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/controltower v1.29.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.35.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.65.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/costoptimizationhub v1.24.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/customerprofiles v1.62.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.64.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/databrew v1.40.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/dataexchange v1.42.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.31.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/datasync v1.59.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/datazone v1.63.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dax v1.30.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/detective v1.39.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.39.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/devopsagent v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/devopsguru v1.41.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.41.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.39.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/dlm v1.37.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdb v1.49.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.21.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/drs v1.39.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/dsql v1.14.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.59.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.309.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.58.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.39.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecs v1.86.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/efs v1.42.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.54.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.35.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.34.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.55.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.42.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.33.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/emr v1.61.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrcontainers v1.41.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.42.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.46.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/evidently v1.30.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/evs v1.11.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/finspace v1.34.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/firehose v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/fis v1.38.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/fms v1.45.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/fsx v1.66.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.56.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/glacier v1.33.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.36.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/glue v1.147.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/grafana v1.36.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/greengrass v1.33.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/groundstation v1.43.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.80.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.39.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v1.54.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.37.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.56.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector v1.31.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.49.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/interconnect v1.1.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.29 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.30 // indirect
	github.com/aws/aws-sdk-go-v2/service/internetmonitor v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/invoicing v1.11.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/iot v1.75.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivs v1.52.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.22.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafka v1.54.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.31.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/kendra v1.61.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.26.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.44.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.31.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.38.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.34.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/kms v1.53.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.48.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v1.94.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/launchwizard v1.15.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice v1.36.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.62.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/licensemanager v1.38.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.56.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/location v1.52.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/m2 v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.52.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.51.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.93.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/medialive v1.99.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.40.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackagev2 v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackagevod v1.40.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.30.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.34.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/mgn v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mpa v1.8.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/mq v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.41.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/mwaaserverless v1.1.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/neptune v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/neptunegraph v1.22.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.61.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkflowmonitor v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.42.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkmonitor v1.14.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/notifications v1.8.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/notificationscontacts v1.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/oam v1.24.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/observabilityadmin v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/odb v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.72.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.32.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/organizations v1.51.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/osis v1.22.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/outposts v1.62.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/paymentcryptography v1.31.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/pcaconnectorad v1.16.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/pcs v1.21.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.40.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2 v1.30.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/pipes v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/polly v1.58.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/pricing v1.42.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/qbusiness v1.35.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/qldb v1.32.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.115.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ram v1.37.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/rbin v1.28.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/rds v1.119.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/rdsdata v1.33.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshift v1.63.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.35.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.52.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.36.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.25.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.34.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.33.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.23.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53 v1.63.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.36.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53profiles v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.33.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.27.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rum v1.31.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.104.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3control v1.71.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3files v1.1.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3outposts v1.35.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3tables v1.16.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3vectors v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.256.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.33.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.18.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/schemas v1.35.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.42.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.71.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.26.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.31.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.40.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.36.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.35.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/ses v1.35.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.62.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sfn v1.43.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.35.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/signer v1.33.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.40.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.69.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.32.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmquicksetup v1.9.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmsap v1.27.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.31.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.39.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.36.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.44.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/swf v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/taxsettings v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb v1.20.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamquery v1.37.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.56.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.73.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/uxc v1.1.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.34.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.23.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/waf v1.31.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.31.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.74.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/workmail v1.37.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.70.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspacesweb v1.40.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/xray v1.37.4 // indirect
	github.com/beevik/etree v1.6.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cedar-policy/cedar-go v1.8.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.6.3 // indirect
	github.com/dave/jennifer v1.7.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.13.0 // indirect
	github.com/evanphx/json-patch v5.9.11+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.9.11 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/fxamacker/cbor/v2 v2.9.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/getkin/kin-openapi v0.133.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-logr/zapr v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.22.4 // indirect
	github.com/go-openapi/jsonreference v0.21.4 // indirect
	github.com/go-openapi/swag v0.25.4 // indirect
	github.com/go-openapi/swag/cmdutils v0.25.4 // indirect
	github.com/go-openapi/swag/conv v0.25.4 // indirect
	github.com/go-openapi/swag/fileutils v0.25.4 // indirect
	github.com/go-openapi/swag/jsonname v0.25.4 // indirect
	github.com/go-openapi/swag/jsonutils v0.25.4 // indirect
	github.com/go-openapi/swag/loading v0.25.4 // indirect
	github.com/go-openapi/swag/mangling v0.25.4 // indirect
	github.com/go-openapi/swag/netutils v0.25.4 // indirect
	github.com/go-openapi/swag/stringutils v0.25.4 // indirect
	github.com/go-openapi/swag/typeutils v0.25.4 // indirect
	github.com/go-openapi/swag/yamlutils v0.25.4 // indirect
	github.com/gobuffalo/flect v1.0.3 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/gnostic-models v0.7.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.23.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.5.0 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.7.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.8 // indirect
	github.com/hashicorp/go-set/v3 v3.0.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.9.0 // indirect
	github.com/hashicorp/hc-install v0.9.4 // indirect
	github.com/hashicorp/hcl/v2 v2.24.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.25.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-jsontypes v0.2.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timeouts v0.7.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timetypes v0.5.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.19.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.10.0 // indirect
	github.com/hashicorp/terraform-plugin-testing v1.16.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.4.0 // indirect
	github.com/hashicorp/terraform-svchost v0.2.1 // indirect
	github.com/hashicorp/yamux v0.1.2 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20240118010651-0ba75a80ca38 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/muvaf/typewriter v0.0.0-20240614220100-70f9d4a54ea0 // indirect
	github.com/oasdiff/oasdiff v1.11.8 // indirect
	github.com/oasdiff/yaml v0.0.0-20250309154309-f31be36b4037 // indirect
	github.com/oasdiff/yaml3 v0.0.0-20250309153720-d2182401db90 // indirect
	github.com/oklog/run v1.2.0 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.23.2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.67.5 // indirect
	github.com/prometheus/procfs v0.19.2 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tmccombs/hcl2json v0.3.3 // indirect
	github.com/upbound/uptest v0.12.1-0.20260123170102-8965579a213b // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/wI2L/jsondiff v0.7.0 // indirect
	github.com/woodsbury/decimal128 v1.3.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	github.com/yargevad/filepathx v1.0.0 // indirect
	github.com/yuin/goldmark v1.7.16 // indirect
	github.com/zclconf/go-cty v1.18.1 // indirect
	github.com/zclconf/go-cty-yaml v1.0.3 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.69.0 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	go.yaml.in/yaml/v4 v4.0.0-rc.6 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/exp v0.0.0-20260112195511-716be5621a96 // indirect
	golang.org/x/mod v0.37.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/oauth2 v0.34.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/telemetry v0.0.0-20260625142307-59b4966ccb57 // indirect
	golang.org/x/term v0.44.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	golang.org/x/time v0.14.0 // indirect
	golang.org/x/tools v0.47.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.5.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260203192932-546029d2fa20 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/dnaeon/go-vcr.v4 v4.0.7 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.13.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/code-generator v0.35.3 // indirect
	k8s.io/component-base v0.35.3 // indirect
	k8s.io/gengo/v2 v2.0.0-20251215205346-5ee0d033ba5b // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20260127142750-a19766b6e2d4 // indirect
	sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730 // indirect
	sigs.k8s.io/structured-merge-diff/v6 v6.3.2-0.20260122202528-d9cc6641c482 // indirect
	sigs.k8s.io/yaml v1.6.0 // indirect
)

replace github.com/hashicorp/terraform-provider-aws => github.com/upbound/terraform-provider-aws v0.0.0-20260702073434-e25b40151251
