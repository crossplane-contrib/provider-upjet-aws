// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

module github.com/upbound/provider-aws

go 1.23.5

require (
	dario.cat/mergo v1.0.0
	github.com/aws/aws-sdk-go v1.55.5
	github.com/aws/aws-sdk-go-v2 v1.30.5
	github.com/aws/aws-sdk-go-v2/config v1.27.35
	github.com/aws/aws-sdk-go-v2/credentials v1.17.33
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.13
	github.com/aws/aws-sdk-go-v2/service/eks v1.48.5
	github.com/aws/aws-sdk-go-v2/service/sts v1.30.8
	github.com/aws/smithy-go v1.20.4
	github.com/crossplane/crossplane-runtime v1.16.0-rc.2.0.20240510094504-3f697876fa57
	github.com/crossplane/crossplane-tools v0.0.0-20230925130601-628280f8bf79
	github.com/crossplane/upjet v1.4.1-0.20240911184956-3afbb7796d46
	github.com/go-ini/ini v1.46.0
	github.com/google/go-cmp v0.6.0
	github.com/hashicorp/awspolicyequivalence v1.6.0
	github.com/hashicorp/terraform-json v0.22.1
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.34.0
	github.com/hashicorp/terraform-provider-aws v0.0.0-00010101000000-000000000000
	github.com/json-iterator/go v1.1.12
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.29.4
	k8s.io/apimachinery v0.29.4
	k8s.io/client-go v0.29.4
	k8s.io/utils v0.0.0-20240921022957-49e7df575cb6
	sigs.k8s.io/controller-runtime v0.17.3
	sigs.k8s.io/controller-tools v0.14.0
)

require (
	github.com/ProtonMail/go-crypto v1.1.0-alpha.5-proton // indirect
	github.com/YakDriver/go-version v0.1.0 // indirect
	github.com/YakDriver/regexache v0.24.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/kingpin/v2 v2.4.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/antchfx/htmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.2.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.32.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/account v1.19.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v1.28.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.35.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/amp v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/amplify v1.24.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.25.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.22.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.32.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/appfabric v1.9.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/appflow v1.43.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/appintegrations v1.28.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.31.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationsignals v1.4.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.27.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.30.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/appstream v1.38.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/appsync v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.44.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.35.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.43.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.22.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/backup v1.37.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/batch v1.44.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/bcmdataexports v1.5.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.17.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/bedrockagent v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/budgets v1.25.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/chatbot v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/chime v1.32.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines v1.18.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/chimesdkvoice v1.17.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.16.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.20.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.53.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.38.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore v1.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.25.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.42.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.40.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.39.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.30.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.43.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecatalyst v1.15.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.25.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.27.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.22.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/codegurureviewer v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.31.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.25.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.44.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.33.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.37.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/configservice v1.48.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/connect v1.109.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/connectcases v1.19.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/controltower v1.16.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.41.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/costoptimizationhub v1.8.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/customerprofiles v1.39.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.40.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/databrew v1.31.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/dataexchange v1.30.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.23.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/datasync v1.40.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/datazone v1.20.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/dax v1.21.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/detective v1.29.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.26.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/devopsguru v1.32.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.27.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.28.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dlm v1.26.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdb v1.37.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.11.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/drs v1.28.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.34.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.177.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.34.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.25.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecs v1.45.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/efs v1.31.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.40.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.26.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.37.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.30.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.25.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/emr v1.43.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrcontainers v1.31.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.23.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.33.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/evidently v1.21.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/finspace v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/firehose v1.32.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/fis v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/fms v1.35.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/fsx v1.47.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.34.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/glacier v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.27.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/glue v1.96.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/grafana v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/greengrass v1.25.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/groundstation v1.29.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.47.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v1.35.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.25.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector v1.23.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.30.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.9.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/internetmonitor v1.17.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/iot v1.56.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/iotanalytics v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/iotevents v1.25.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivs v1.38.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.14.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafka v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.19.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/kendra v1.52.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.12.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.29.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.23.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.25.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/kms v1.35.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.35.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v1.59.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/launchwizard v1.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.47.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/licensemanager v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/location v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/lookoutmetrics v1.29.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/m2 v1.15.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.41.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.33.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.58.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/medialive v1.59.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.32.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediapackagev2 v1.15.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.22.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.21.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/mq v1.25.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.29.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/neptune v1.33.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/neptunegraph v1.11.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.41.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.29.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkmonitor v1.5.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/oam v1.13.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.39.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.14.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/organizations v1.31.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/osis v1.12.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/outposts v1.42.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/paymentcryptography v1.12.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/pcaconnectorad v1.7.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/pcs v1.0.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.32.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2 v1.12.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/pipes v1.15.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/polly v1.43.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/pricing v1.30.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/qbusiness v1.11.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/qldb v1.23.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.71.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/ram v1.27.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/rbin v1.18.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/rds v1.83.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshift v1.46.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.28.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.21.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.43.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.24.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.12.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.24.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.23.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.14.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53 v1.43.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.25.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53profiles v1.2.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.23.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.19.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.30.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/rum v1.19.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.62.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3control v1.47.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3outposts v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.157.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.10.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/schemas v1.26.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.32.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.52.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.16.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.22.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.30.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.28.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.31.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.23.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/ses v1.26.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.33.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sfn v1.31.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/signer v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.31.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.34.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.53.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.32.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssmsap v1.16.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.22.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.27.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.26.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/swf v1.25.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.39.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.50.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.17.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.10.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/waf v1.23.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.23.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.52.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.32.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/worklink v1.22.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.46.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspacesweb v1.21.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/xray v1.27.8 // indirect
	github.com/beevik/etree v1.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cedar-policy/cedar-go v0.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.4.0 // indirect
	github.com/dave/jennifer v1.4.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.8.0 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.17.0 // indirect
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
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.57 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2 v2.0.0-beta.58 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200723130312-85980079f637 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.6.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/hc-install v0.8.0 // indirect
	github.com/hashicorp/hcl/v2 v2.22.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.21.0 // indirect
	github.com/hashicorp/terraform-plugin-framework v1.11.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-jsontypes v0.2.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-timeouts v0.4.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-timetypes v0.5.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.13.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.23.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/hashicorp/terraform-plugin-mux v0.16.0 // indirect
	github.com/hashicorp/terraform-plugin-testing v1.10.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.3 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20240118010651-0ba75a80ca38 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
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
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
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
	github.com/zclconf/go-cty v1.15.0 // indirect
	github.com/zclconf/go-cty-yaml v1.0.3 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.55.0 // indirect
	go.opentelemetry.io/otel v1.30.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/term v0.27.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.25.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240827150818-7e3bb234dfed // indirect
	google.golang.org/grpc v1.66.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apiextensions-apiserver v0.29.2 // indirect
	k8s.io/component-base v0.29.2 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/kube-openapi v0.0.0-20231010175941-2dd684a91f00 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

// copied from the Terraform provider
replace github.com/hashicorp/terraform-plugin-log => github.com/gdavison/terraform-plugin-log v0.0.0-20230928191232-6c653d8ef8fb

replace github.com/hashicorp/terraform-provider-aws => github.com/upbound/terraform-provider-aws v0.0.0-20241019052313-8fed8454b0e3
