// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

module github.com/upbound/provider-aws

go 1.23.7

require (
	dario.cat/mergo v1.0.1
	github.com/aws/aws-sdk-go v1.55.6
	github.com/aws/aws-sdk-go-v2 v1.36.3
	github.com/aws/aws-sdk-go-v2/config v1.29.9
	github.com/aws/aws-sdk-go-v2/credentials v1.17.62
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.30
	github.com/aws/aws-sdk-go-v2/service/eks v1.60.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.17
	github.com/aws/smithy-go v1.22.3
	github.com/crossplane/crossplane-runtime v1.17.0
	github.com/crossplane/crossplane-tools v0.0.0-20230925130601-628280f8bf79
	github.com/crossplane/upjet v1.5.0
	github.com/go-ini/ini v1.46.0
	github.com/google/go-cmp v0.7.0
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.62
	github.com/hashicorp/awspolicyequivalence v1.7.0
	github.com/hashicorp/terraform-json v0.24.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.36.1
	github.com/hashicorp/terraform-provider-aws v0.0.0-00010101000000-000000000000
	github.com/json-iterator/go v1.1.12
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.30.0
	k8s.io/apimachinery v0.30.0
	k8s.io/client-go v0.30.0
	k8s.io/utils v0.0.0-20241210054802-24370beab758
	sigs.k8s.io/controller-runtime v0.18.2
	sigs.k8s.io/controller-tools v0.14.0
)

require (
	github.com/ProtonMail/go-crypto v1.1.6 // indirect
	github.com/YakDriver/go-version v0.1.0 // indirect
	github.com/YakDriver/regexache v0.24.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/kingpin/v2 v2.4.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/antchfx/htmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.2.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.66 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.38.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/account v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.39.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/amp v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/amplify v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.37.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appfabric v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appflow v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appintegrations v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.30.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationsignals v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appstream v1.45.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appsync v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.50.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.38.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.52.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/backup v1.41.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/batch v1.51.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/bcmdataexports v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.28.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrockagent v1.39.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/billing v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/budgets v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/chatbot v1.10.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/chime v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkvoice v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.24.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.58.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.45.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.48.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.34.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.55.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecatalyst v1.17.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeconnections v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codegurureviewer v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.29.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.51.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.42.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/configservice v1.52.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/connect v1.127.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/connectcases v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/controltower v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.48.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/costoptimizationhub v1.13.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/customerprofiles v1.45.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.51.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/databrew v1.34.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dataexchange v1.34.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/datasync v1.47.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/datazone v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dax v1.24.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/detective v1.32.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/devopsguru v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dlm v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdb v1.41.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.15.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/drs v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dsql v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.41.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.209.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.43.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecs v1.54.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/efs v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.45.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.45.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/emr v1.48.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrcontainers v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.37.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/evidently v1.24.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/finspace v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/firehose v1.37.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/fis v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/fms v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/fsx v1.53.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/glacier v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/glue v1.106.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/grafana v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/greengrass v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/groundstation v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.54.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.41.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.36.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.7.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internetmonitor v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/invoicing v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/iot v1.64.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/iotanalytics v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/iotevents v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivs v1.43.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafka v1.39.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kendra v1.56.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kms v1.38.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.40.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v1.70.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/launchwizard v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.50.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/licensemanager v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.43.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/location v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/lookoutmetrics v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/m2 v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.45.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.38.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.69.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/medialive v1.71.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackagev2 v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackagevod v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mgn v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mq v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.34.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/neptune v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/neptunegraph v1.17.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkmonitor v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/oam v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.27.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/organizations v1.38.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/osis v1.15.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/outposts v1.49.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/paymentcryptography v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/pcaconnectorad v1.11.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/pcs v1.3.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2 v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/pipes v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/polly v1.47.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/pricing v1.34.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/qbusiness v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/qldb v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.84.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ram v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rbin v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rds v1.94.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshift v1.54.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.33.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.46.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53 v1.49.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53profiles v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/rum v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.78.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3control v1.54.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3outposts v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3tables v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.180.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.13.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/schemas v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.35.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.57.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ses v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.43.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sfn v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/signer v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.34.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.38.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.57.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmquicksetup v1.4.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmsap v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/swf v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/taxsettings v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb v1.10.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamquery v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.44.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.57.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.22.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.14.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/waf v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.58.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.35.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/worklink v1.23.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.55.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspacesweb v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/xray v1.31.1 // indirect
	github.com/beevik/etree v1.5.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cedar-policy/cedar-go v0.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.6.0 // indirect
	github.com/dave/jennifer v1.4.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch v5.9.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.9.0 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-logr/zapr v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gobuffalo/flect v1.0.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.23.0 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2 v2.0.0-beta.63 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.4.1 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.6.3 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/hc-install v0.9.1 // indirect
	github.com/hashicorp/hcl/v2 v2.23.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.22.0 // indirect
	github.com/hashicorp/terraform-plugin-framework v1.14.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-jsontypes v0.2.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timeouts v0.5.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timetypes v0.5.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.17.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.26.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/hashicorp/terraform-plugin-mux v0.18.0 // indirect
	github.com/hashicorp/terraform-plugin-testing v1.11.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.4 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.2 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
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
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/muvaf/typewriter v0.0.0-20210910160850-80e49fe1eb32 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tmccombs/hcl2json v0.3.3 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	github.com/zclconf/go-cty v1.16.2 // indirect
	github.com/zclconf/go-cty-yaml v1.0.3 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.60.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/oauth2 v0.25.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/term v0.30.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.31.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250303144028-a0af3efb3deb // indirect
	google.golang.org/grpc v1.71.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apiextensions-apiserver v0.30.0 // indirect
	k8s.io/component-base v0.30.0 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

// copied from the Terraform provider
replace github.com/hashicorp/terraform-plugin-log => github.com/gdavison/terraform-plugin-log v0.0.0-20230928191232-6c653d8ef8fb

replace github.com/hashicorp/terraform-provider-aws => ./upstream
