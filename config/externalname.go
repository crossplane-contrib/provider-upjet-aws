// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// TerraformPluginFrameworkExternalNameConfigs contains all external
// name configurations belonging to Terraform Plugin Framework
// resources to be reconciled under the no-fork architecture for this
// provider.
var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{

	// ********** When adding new services please keep them alphabetized by their aws go sdk package name **********

	// apigateway
	//
	// API Gateway Accounts can be imported using the word api-gateway-account
	"aws_api_gateway_account": apiGatewayAccount(),

	// appconfig
	//
	// AppConfig Environments can be imported by using the environment ID and application ID separated by a colon (:)
	// terraform-plugin-framework
	"aws_appconfig_environment": appConfigEnvironment(),

	// batch
	// AWS Batch job queue can be imported using the name
	"aws_batch_job_queue": config.TemplatedStringAsIdentifier("name", fullARNTemplate("batch", "job-queue/{{ .external_name }}")),

	// bedrock
	//
	// Bedrock inference profile can be imported using the ID: inference_profile-id-12345678
	"aws_bedrock_inference_profile": identifierFromProviderWithDefaultStub("bedrock12345"),

	// bedrockagent
	//
	// Bedrock Agent can be imported using the agent arn
	"aws_bedrockagent_agent": identifierFromProviderWithDefaultStub("STUB123456"),

	// CodeGuru Profiler
	// Profiling Group can be imported using the the profiling group name
	"aws_codeguruprofiler_profiling_group": config.NameAsIdentifier,

	// cognitoidp
	//
	// us-west-2_abc123/3ho4ek12345678909nh3fmhpko
	"aws_cognito_user_pool_client": cognitoUserPoolClient(),

	// dsql
	//
	// DSQL Cluster can be imported using the identifier
	"aws_dsql_cluster": config.FrameworkResourceWithComputedIdentifier("identifier", "artix3b6dqiognkp7732wzhroi"),
	// DSQL Cluster Peering resource can be imported using the Cluster identifier
	"aws_dsql_cluster_peering": dsqlClusterPeering(),

	// dynamodb
	//
	// DynamoDB table resource policy can be imported using the DynamoDB resource identifier
	"aws_dynamodb_resource_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_arn }}"),

	// ec2
	//
	// Imported by using the id: sgr-02108b27edd666983
	"aws_vpc_security_group_egress_rule": vpcSecurityGroupRule(),
	// Imported by using the id: sgr-02108b27edd666983
	"aws_vpc_security_group_ingress_rule": vpcSecurityGroupRule(),

	// elasticache
	//
	// Imported by using the serverless cache name
	"aws_elasticache_serverless_cache": config.NameAsIdentifier,

	// eks
	//
	// PodIdentityAssociation can be imported using the association ID by passing spec.forProvider.clusterName field
	"aws_eks_pod_identity_association": eksPodIdentityAssociation(),

	// glue
	//
	//
	"aws_glue_catalog_table_optimizer": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .parameters.database_name }}:{{ .external_name }}"),

	// kafka
	//
	// single MSK SCRAM secret associations can be imported using cluster_arn and secret_arn, separated by a comma (,)
	"aws_msk_single_scram_secret_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.cluster_arn }},{{ .parameters.secret_arn }}"),

	// mq
	//
	// admin
	"aws_mq_user": mqUser(),

	// opensearchserverless
	//
	// AccessPolicy can be imported using the policy name
	"aws_opensearchserverless_access_policy": config.NameAsIdentifier,
	// Collection can be imported using the AWS-assigned collection ID. i.e. ch9rq91uv4yd8rff1f39
	"aws_opensearchserverless_collection": opensearchserverlessCollection(),
	// LifecyclePolicy can be imported using the policy name
	"aws_opensearchserverless_lifecycle_policy": config.NameAsIdentifier,
	//  SecurityConfig can be imported using the AWS-assigned security config ID
	"aws_opensearchserverless_security_config": config.TemplatedStringAsIdentifier("name", "{{ .parameters.type }}/{{ .setup.client_metadata.account_id }}/{{ .external_name }}"),
	// SecurityPolicy can be imported using the policy name
	"aws_opensearchserverless_security_policy": config.NameAsIdentifier,
	// VPCEndpoint can be imported using the AWS-assigned VPC Endpoint ID, i.e. vpce-0a957ae9ed5aee308
	"aws_opensearchserverless_vpc_endpoint": opensearchserverlessVpcEndpoint(),

	// osis
	//
	// OSIS Pipeline can be imported using the name
	"aws_osis_pipeline": config.ParameterAsIdentifier("pipeline_name"),

	// amp
	//
	// Prometheus Scraper can be imported using the ARN: arn:aws:aps:us-west-2:123456789012:scraper/s-12345678-1234-1234-1234-123456789012
	// Terraform returns the full ARN as ID, but AWS API expects just the UUID portion (s-UUID).
	// terraform-plugin-framework
	"aws_prometheus_scraper": identifierFromProviderWithDefaultStub("scraper12345"),

	// rds
	//
	// aws_rds_instance_state import format: rdsInstanceId-12345678
	"aws_rds_instance_state": rdsInstanceState(),

	// s3
	//
	// S3 directory bucket can be imported using the full id: [bucket_name]--[azid]--x-s3
	"aws_s3_directory_bucket": config.ParameterAsIdentifier("bucket"),
	// The S3 bucket lifecycle configuration resource should be imported using the bucket
	"aws_s3_bucket_lifecycle_configuration": s3LifecycleConfiguration(),

	// vpclattice
	//
	// VPC Lattice Resource Configuration can be imported using the id
	"aws_vpclattice_resource_configuration": identifierFromProviderWithDefaultStub("rcfg-1234567890abcdef1"),
	// VPC Lattice Resource Gateway can be imported using the id
	"aws_vpclattice_resource_gateway": identifierFromProviderWithDefaultStub("rgw-055b56956a39439ba"),
	// VPC Lattice Service Network Resource Association can be imported using the id
	"aws_vpclattice_service_network_resource_association": identifierFromProviderWithDefaultStub("snra-1234567890abcef12"),

	// ********** When adding new services please keep them alphabetized by their aws go sdk package name **********
}

// TerraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform Plugin SDKv2 resources to be reconciled
// under the no-fork architecture for this provider.
var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{

	// ********** When adding new services please keep them alphabetized by their aws go sdk package name **********

	// accessanalyzer
	//
	// Access Analyzer Analyzers can be imported using the analyzer_name
	"aws_accessanalyzer_analyzer": config.ParameterAsIdentifier("analyzer_name"),
	// AccessAnalyzer ArchiveRule can be imported using the analyzer_name/rule_name
	"aws_accessanalyzer_archive_rule": config.TemplatedStringAsIdentifier("rule_name", "{{ .parameters.analyzer_name }}/{{ .external_name }}"),

	// account
	//
	// The Alternate Contact for the current account can be imported using the alternate_contact_type
	"aws_account_alternate_contact": config.TemplatedStringAsIdentifier("", "{{ .parameters.alternate_contact_type }}"),
	// The account region can be imported using region_name or a comma separated account_id and region_name
	"aws_account_region": config.TemplatedStringAsIdentifier("", "{{ .parameters.region_name }}"),

	// ACM
	// Imported using ARN that has a random substring:
	// arn:aws:acm:eu-central-1:123456789012:certificate/7e7a28d2-163f-4b8f-b9cd-822f96c08d6a
	"aws_acm_certificate": config.IdentifierFromProvider,
	// No import documented, but https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate_validation#id
	"aws_acm_certificate_validation": config.IdentifierFromProvider,

	// ACM PCA
	// aws_acmpca_certificate can not be imported at this time.
	"aws_acmpca_certificate": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	//	// arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
	"aws_acmpca_certificate_authority": config.IdentifierFromProvider,
	// No doc on import, but resource is getting CA ARN:
	// arn:aws:acm-pca:eu-central-1:609897127049:certificate-authority/ba0c7989-9641-4f36-a033-dee60121d595
	"aws_acmpca_certificate_authority_certificate": config.IdentifierFromProvider,
	// No import
	"aws_acmpca_permission": config.IdentifierFromProvider,
	// aws_acmpca_policy can be imported using the resource_arn value
	// Example: arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
	"aws_acmpca_policy": config.IdentifierFromProvider,

	// amp
	//
	// Uses the ID of workspace, workspace_id parameter.
	"aws_prometheus_alert_manager_definition": config.IdentifierFromProvider,
	//
	"aws_prometheus_rule_group_namespace": config.TemplatedStringAsIdentifier("name", fullARNTemplate("aps", "rulegroupsnamespace/{{ .parameters.workspace_id }}/{{ .external_name }}")),
	// ID is a random UUID.
	"aws_prometheus_workspace": config.IdentifierFromProvider,

	// amplify
	//
	// Amplify App can be imported using Amplify App ID (appId)
	"aws_amplify_app": config.IdentifierFromProvider,
	// Amplify backend environment can be imported using app_id and environment_name: d2ypk4k47z8u6/example
	"aws_amplify_backend_environment": config.TemplatedStringAsIdentifier("environment_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify branch can be imported using app_id and branch_name: d2ypk4k47z8u6/master
	"aws_amplify_branch": config.TemplatedStringAsIdentifier("branch_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify webhook can be imported using a webhook ID
	"aws_amplify_webhook": config.IdentifierFromProvider,

	// apigateway
	//
	// API Gateway Keys can be imported using the id
	"aws_api_gateway_api_key": config.IdentifierFromProvider,
	// AWS API Gateway Authorizer can be imported using the REST-API-ID/AUTHORIZER-ID
	"aws_api_gateway_authorizer": config.IdentifierFromProvider,
	// aws_api_gateway_base_path_mapping can be imported by using the domain name and base path.
	// For empty base_path (e.g., root path (/)): example.com/
	// Otherwise: example.com/base-path
	"aws_api_gateway_base_path_mapping": config.IdentifierFromProvider,
	// API Gateway Client Certificates can be imported using the id
	"aws_api_gateway_client_certificate": config.IdentifierFromProvider,
	// No import
	"aws_api_gateway_deployment": config.IdentifierFromProvider,
	// API Gateway documentation_parts can be imported using REST-API-ID/DOC-PART-ID
	"aws_api_gateway_documentation_part": config.IdentifierFromProvider,
	// API Gateway documentation versions can be imported using REST-API-ID/VERSION
	"aws_api_gateway_documentation_version": FormattedIdentifierFromProvider("/", "rest_api_id", "version"),
	// API Gateway domain names can be imported using their name
	"aws_api_gateway_domain_name": config.IdentifierFromProvider,
	// aws_api_gateway_gateway_response can be imported using REST-API-ID/RESPONSE-TYPE
	"aws_api_gateway_gateway_response": FormattedIdentifierFromProvider("/", "rest_api_id", "response_type"),
	// aws_api_gateway_integration can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD
	"aws_api_gateway_integration": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method"),
	// aws_api_gateway_integration_response can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE
	"aws_api_gateway_integration_response": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method", "status_code"),
	// aws_api_gateway_method can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD
	"aws_api_gateway_method": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method"),
	// aws_api_gateway_method_response can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE
	"aws_api_gateway_method_response": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method", "status_code"),
	// aws_api_gateway_method_settings can be imported using REST-API-ID/STAGE-NAME/METHOD-PATH
	"aws_api_gateway_method_settings": FormattedIdentifierFromProvider("/", "rest_api_id", "stage_name", "method_path"),
	// aws_api_gateway_model can be imported using REST-API-ID/NAME
	"aws_api_gateway_model": config.IdentifierFromProvider,
	// aws_api_gateway_request_validator can be imported using REST-API-ID/REQUEST-VALIDATOR-ID
	"aws_api_gateway_request_validator": config.IdentifierFromProvider,
	// aws_api_gateway_resource can be imported using REST-API-ID/RESOURCE-ID
	"aws_api_gateway_resource": config.IdentifierFromProvider,
	// aws_api_gateway_rest_api can be imported by using the REST API ID
	"aws_api_gateway_rest_api": config.IdentifierFromProvider,
	// aws_api_gateway_rest_api_policy can be imported by using the REST API ID
	"aws_api_gateway_rest_api_policy": FormattedIdentifierFromProvider("", "rest_api_id"),
	// aws_api_gateway_stage can be imported using REST-API-ID/STAGE-NAME
	"aws_api_gateway_stage": FormattedIdentifierFromProvider("/", "rest_api_id", "stage_name"),
	// AWS API Gateway Usage Plan can be imported using the id
	"aws_api_gateway_usage_plan": config.IdentifierFromProvider,
	// AWS API Gateway Usage Plan Key can be imported using the USAGE-PLAN-ID/USAGE-PLAN-KEY-ID
	"aws_api_gateway_usage_plan_key": config.IdentifierFromProvider,
	// API Gateway VPC Link can be imported using the id
	"aws_api_gateway_vpc_link": config.IdentifierFromProvider,

	// apigatewayv2
	//
	"aws_apigatewayv2_api": config.IdentifierFromProvider,
	// Case4: Imported by using the API mapping identifier and domain name.
	"aws_apigatewayv2_api_mapping": config.IdentifierFromProvider,
	// Case4: Imported by using the API identifier and authorizer identifier.
	"aws_apigatewayv2_authorizer": config.IdentifierFromProvider,
	// Case4: Imported by using the API identifier and deployment identifier.
	"aws_apigatewayv2_deployment": config.IdentifierFromProvider,
	//
	"aws_apigatewayv2_domain_name": config.ParameterAsIdentifier("domain_name"),
	// Case4: Imported by using the API identifier and integration identifier.
	"aws_apigatewayv2_integration": config.IdentifierFromProvider,
	// Case4: Imported by using the API identifier, integration identifier and
	// integration response identifier.
	"aws_apigatewayv2_integration_response": config.IdentifierFromProvider,
	// Case4: Imported by using the API identifier and model identifier.
	"aws_apigatewayv2_model": config.IdentifierFromProvider,
	// Case4: Imported by using the API identifier and route identifier.
	"aws_apigatewayv2_route": config.IdentifierFromProvider,
	// Case4: Imported by using the API identifier, route identifier and route
	// response identifier.
	"aws_apigatewayv2_route_response": config.IdentifierFromProvider,
	// Imported by using the API identifier and stage name.
	"aws_apigatewayv2_stage": config.NameAsIdentifier,
	// aws_apigatewayv2_vpc_link can be imported by using the VPC Link id
	"aws_apigatewayv2_vpc_link": config.IdentifierFromProvider,

	// appautoscaling
	//
	// Application AutoScaling Policy can be imported using the service-namespace, resource-id, scalable-dimension and policy-name separated by /
	"aws_appautoscaling_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.service_namespace }}/{{ .parameters.resource_id }}/{{ .parameters.scalable_dimension }}/{{ .external_name }}"),
	// No import
	"aws_appautoscaling_scheduled_action": config.IdentifierFromProvider,
	// Application AutoScaling Target can be imported using the service-namespace , resource-id and scalable-dimension separated by /
	"aws_appautoscaling_target": config.IdentifierFromProvider,

	// appconfig
	//
	// AppConfig Applications can be imported using their application ID,
	"aws_appconfig_application": config.IdentifierFromProvider,
	// AppConfig Configuration Profiles can be imported by using the configuration profile ID and application ID separated by a colon (:)
	"aws_appconfig_configuration_profile": config.IdentifierFromProvider,
	// AppConfig Deployments can be imported by using the application ID, environment ID, and deployment number separated by a slash (/)
	"aws_appconfig_deployment": config.IdentifierFromProvider,
	// AppConfig Deployment Strategies can be imported by using their deployment strategy ID
	"aws_appconfig_deployment_strategy": config.IdentifierFromProvider,
	// AppConfig Extensions can be imported using their extension ID
	// ID is a provider-generated
	"aws_appconfig_extension": config.IdentifierFromProvider,
	// AppConfig Extension Associations can be imported using their extension association ID
	// ID is a provider-generated
	"aws_appconfig_extension_association": config.IdentifierFromProvider,
	// AppConfig Hosted Configuration Versions can be imported by using the application ID, configuration profile ID, and version number separated by a slash (/)
	"aws_appconfig_hosted_configuration_version": config.IdentifierFromProvider,

	// appflow
	//
	// arn:aws:appflow:us-west-2:123456789012:flow/example-flow
	"aws_appflow_flow": config.TemplatedStringAsIdentifier("name", fullARNTemplate("appflow", "flow/{{ .external_name }}")),

	// appintegrations
	//
	// Amazon AppIntegrations Event Integrations can be imported using the name
	"aws_appintegrations_event_integration": config.NameAsIdentifier,

	// applicationinsights
	//
	// ApplicationInsights Applications can be imported using the resource_group_name
	"aws_applicationinsights_application": config.ParameterAsIdentifier("resource_group_name"),

	// appmesh
	//
	// mesh/gw1/example-gateway-route
	"aws_appmesh_gateway_route": config.IdentifierFromProvider,
	// App Mesh service meshes can be imported using the name
	"aws_appmesh_mesh": config.NameAsIdentifier,
	// App Mesh virtual routes can be imported using mesh_name and virtual_router_name together with the route's name, e.g.,
	// simpleapp/serviceB/serviceB-route
	"aws_appmesh_route": config.IdentifierFromProvider,
	// App Mesh virtual gateway can be imported using mesh_name together with the virtual gateway's name: mesh/gw1
	"aws_appmesh_virtual_gateway": config.IdentifierFromProvider,
	// App Mesh virtual nodes can be imported using mesh_name together with the virtual node's name: simpleapp/serviceBv1
	"aws_appmesh_virtual_node": config.IdentifierFromProvider,
	// App Mesh virtual routers can be imported using mesh_name together with the virtual router's name: simpleapp/serviceB
	"aws_appmesh_virtual_router": config.IdentifierFromProvider,
	// App Mesh virtual services can be imported using mesh_name together with the virtual service's name: simpleapp/servicea.simpleapp.local
	"aws_appmesh_virtual_service": config.IdentifierFromProvider,

	// apprunner
	//
	// App Runner AutoScaling Configuration Versions can be imported by using the arn
	"aws_apprunner_auto_scaling_configuration_version": config.IdentifierFromProvider,
	// App Runner Connections can be imported by using the connection_name
	"aws_apprunner_connection": config.ParameterAsIdentifier("connection_name"),
	// App Runner Observability Configuration can be imported by using the arn
	// Example: arn:aws:apprunner:us-east-1:1234567890:observabilityconfiguration/example/1/d75bc7ea55b71e724fe5c23452fe22a1
	"aws_apprunner_observability_configuration": config.IdentifierFromProvider,
	// App Runner Services can be imported by using the arn
	"aws_apprunner_service": config.IdentifierFromProvider,
	// App Runner vpc connector can be imported by using the arn
	"aws_apprunner_vpc_connector": config.IdentifierFromProvider,

	// appstream
	//
	// aws_appstream_directory_config can be imported using the id
	"aws_appstream_directory_config": config.IdentifierFromProvider,
	// aws_appstream_fleet can be imported using the id
	"aws_appstream_fleet": config.IdentifierFromProvider,
	// AppStream Stack Fleet Association can be imported by using the fleet_name and stack_name separated by a slash (/)
	"aws_appstream_fleet_stack_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.fleet_name }}/{{ .parameters.stack_name}}"),
	// aws_appstream_image_builder can be imported using the name
	"aws_appstream_image_builder": config.NameAsIdentifier,
	// aws_appstream_stack can be imported using the id
	"aws_appstream_stack": config.IdentifierFromProvider,
	// aws_appstream_user can be imported using the user_name and authentication_type separated by a slash (/)
	"aws_appstream_user": config.TemplatedStringAsIdentifier("user_name", "{{ .external_name }}/{{ .parameters.authentication_type }}"),
	// AppStream User Stack Association can be imported by using the user_name, authentication_type, and stack_name, separated by a slash (/)
	"aws_appstream_user_stack_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.user_name }}/{{ .parameters.authentication_type }}/{{ .parameters.stack_name }}/"),

	// appsync
	//
	// aws_appsync_api_cache can be imported using the AppSync API ID
	"aws_appsync_api_cache": config.IdentifierFromProvider,
	// aws_appsync_api_key can be imported using the AppSync API ID and key separated by :
	"aws_appsync_api_key": TemplatedStringAsProviderDefinedIdentifier("{{ .parameters.api_id }}:{{ .external_name }}"),
	// aws_appsync_datasource can be imported with their api_id, a hyphen, and name
	"aws_appsync_datasource": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}-{{ .external_name }}"),
	// aws_appsync_function can be imported using the AppSync API ID and Function ID separated by -
	"aws_appsync_function": config.IdentifierFromProvider,
	// AppSync GraphQL API can be imported using the GraphQL API ID
	"aws_appsync_graphql_api": config.IdentifierFromProvider,
	// aws_appsync_resolver can be imported with their api_id, a hyphen, type, a hypen and field
	"aws_appsync_resolver": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_id }}-{{ .parameters.type }}-{{ .parameters.field }}"),

	// athena
	//
	// Data catalogs can be imported using their name
	"aws_athena_data_catalog": config.NameAsIdentifier,
	// Athena Databases can be imported using their name
	"aws_athena_database": config.NameAsIdentifier,
	// Athena Named Query can be imported using the query ID
	"aws_athena_named_query": config.IdentifierFromProvider,
	// Athena Workgroups can be imported using their name
	"aws_athena_workgroup": config.NameAsIdentifier,

	// autoscaling
	//
	// No terraform import.
	"aws_autoscaling_attachment": config.IdentifierFromProvider,
	//
	"aws_autoscaling_group": config.NameAsIdentifier,
	// aws_autoscaling_group_tag can be imported by using the ASG name and key, separated by a comma (,)
	"aws_autoscaling_group_tag": config.IdentifierFromProvider,
	// AutoScaling Lifecycle Hooks can be imported using the role autoscaling_group_name and name separated by /
	"aws_autoscaling_lifecycle_hook": config.TemplatedStringAsIdentifier("name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// No import
	"aws_autoscaling_notification": config.IdentifierFromProvider,
	// AutoScaling scaling policy can be imported using the role autoscaling_group_name and name separated by /
	"aws_autoscaling_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// AutoScaling ScheduledAction can be imported using the auto-scaling-group-name and scheduled-action-name: auto-scaling-group-name/scheduled-action-name
	"aws_autoscaling_schedule": config.TemplatedStringAsIdentifier("scheduled_action_name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),

	// autoscalingplans
	//
	// Auto Scaling scaling plans can be imported using the name
	"aws_autoscalingplans_scaling_plan": config.IdentifierFromProvider,

	// backup
	//
	// Backup Framework can be imported using the id which corresponds to the name of the Backup Framework
	"aws_backup_framework": config.IdentifierFromProvider,
	// Backup Global Settings can be imported using the id
	"aws_backup_global_settings": config.IdentifierFromProvider,
	// Backup Plan can be imported using the id
	"aws_backup_plan": config.IdentifierFromProvider,
	// Backup Region Settings can be imported using the region
	"aws_backup_region_settings": config.IdentifierFromProvider,
	// Backup Report Plan can be imported using the id which corresponds to the name of the Backup Report Plan
	"aws_backup_report_plan": config.IdentifierFromProvider,
	// Backup selection can be imported using the role plan_id and id separated by | plan-id|selection-id
	"aws_backup_selection": config.IdentifierFromProvider,
	// Backup vault can be imported using the name
	"aws_backup_vault": config.NameAsIdentifier,
	// Backup vault lock configuration can be imported using the name of the backup vault
	"aws_backup_vault_lock_configuration": config.IdentifierFromProvider,
	// Backup vault notifications can be imported using the name of the backup vault
	"aws_backup_vault_notifications": config.IdentifierFromProvider,
	// Backup vault policy can be imported using the name of the backup vault
	"aws_backup_vault_policy": config.IdentifierFromProvider,

	// batch
	//
	// AWS Batch compute can be imported using the name
	"aws_batch_compute_environment": config.NameAsIdentifier,
	// Batch Job Definition can be imported using ARN that has a random substring, revision at the end:
	// arn:aws:batch:us-east-1:123456789012:job-definition/sample:1
	"aws_batch_job_definition": config.IdentifierFromProvider,
	// Batch Scheduling Policy can be imported using the arn: arn:aws:batch:us-east-1:123456789012:scheduling-policy/sample
	"aws_batch_scheduling_policy": config.TemplatedStringAsIdentifier("name", fullARNTemplate("batch", "scheduling-policy/{{ .external_name }}")),

	// budgets
	//
	// Budgets can be imported using AccountID:BudgetName
	"aws_budgets_budget": config.TemplatedStringAsIdentifier("name", "{{ .setup.client_metadata.account_id }}:{{ .external_name }}"),
	// Budgets can be imported using AccountID:ActionID:BudgetName
	"aws_budgets_budget_action": config.IdentifierFromProvider,

	// ce
	//
	// aws_ce_anomaly_monitor can be imported using the id
	"aws_ce_anomaly_monitor": config.IdentifierFromProvider,

	// chime
	//
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector": config.NameAsIdentifier,
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector_group": config.NameAsIdentifier,
	// Chime Voice Connector Logging can be imported using the voice_connector_id
	"aws_chime_voice_connector_logging": config.IdentifierFromProvider,
	// Chime Voice Connector Origination can be imported using the voice_connector_id
	"aws_chime_voice_connector_origination": config.IdentifierFromProvider,
	// Chime Voice Connector Streaming can be imported using the voice_connector_id
	"aws_chime_voice_connector_streaming": config.IdentifierFromProvider,
	// Chime Voice Connector Termination can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination": config.IdentifierFromProvider,
	// Chime Voice Connector Termination Credentials can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination_credentials": config.IdentifierFromProvider,

	// cloud9
	//
	// No import
	"aws_cloud9_environment_ec2": config.IdentifierFromProvider,
	// Cloud9 environment membership can be imported using the environment-id#user-arn
	"aws_cloud9_environment_membership": config.TemplatedStringAsIdentifier("", "{{ .parameters.environment_id }}#{{ .parameters.user_arn }}"),

	// cloudcontrol
	//
	// No import
	"aws_cloudcontrolapi_resource": config.IdentifierFromProvider,

	// cloudformation
	//
	// config.NameAsIdentifier did not work, the identifier for the resource turned out to be an ARN
	// arn:aws:cloudformation:us-west-1:123456789123:stack/networking-stack/1e691240-6f2c-11ed-8f91-06094dc221f3
	"aws_cloudformation_stack": TemplatedStringAsIdentifierWithNoName(fullARNTemplate("cloudformation", "stack/{{ .parameters.name }}/{{ .external_name }}")),
	// CloudFormation StackSets can be imported using the name
	"aws_cloudformation_stack_set": config.NameAsIdentifier,
	// Cloudformation Stacks Instances imported using the StackSet name, target
	// AWS account ID, and target AWS region separated with commas:
	// example,123456789012,us-east-1
	"aws_cloudformation_stack_set_instance": config.IdentifierFromProvider,

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
	// CloudFront Origin Access Control can be imported using the id
	"aws_cloudfront_origin_access_control": config.IdentifierFromProvider,
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

	// cloudsearch
	//
	// CloudSearch Domains can be imported using the name
	"aws_cloudsearch_domain": config.NameAsIdentifier,
	// CloudSearch domain service access policies can be imported using the domain name
	"aws_cloudsearch_domain_service_access_policy": config.IdentifierFromProvider,

	// cloudtrail
	//
	// Cloudtrails can be imported using the name arn:aws:cloudtrail:us-west-1:153891904029:trail/foobar
	"aws_cloudtrail": config.TemplatedStringAsIdentifier("name", fullARNTemplate("cloudtrail", "trail/{{ .external_name }}")),
	// Event data stores can be imported using their arn
	"aws_cloudtrail_event_data_store": config.IdentifierFromProvider,

	// cloudwatch
	//
	// Use the alarm_name to import a CloudWatch Composite Alarm.
	"aws_cloudwatch_composite_alarm": config.ParameterAsIdentifier("alarm_name"),
	// CloudWatch dashboards can be imported using the dashboard_name
	"aws_cloudwatch_dashboard": config.ParameterAsIdentifier("dashboard_name"),
	// CloudWatch Metric Alarm can be imported using the alarm_name
	"aws_cloudwatch_metric_alarm": config.ParameterAsIdentifier("alarm_name"),
	// CloudWatch metric streams can be imported using the name
	"aws_cloudwatch_metric_stream": config.IdentifierFromProvider,

	// cloudwatchevents
	//
	// Imported using name
	"aws_cloudwatch_event_api_destination": config.NameAsIdentifier,
	// Imported using name
	"aws_cloudwatch_event_archive": config.NameAsIdentifier,
	// Imported using name
	"aws_cloudwatch_event_bus": config.NameAsIdentifier,
	// Imported using event_bus_name
	"aws_cloudwatch_event_bus_policy": config.IdentifierFromProvider,
	// Imported using name
	"aws_cloudwatch_event_connection": config.NameAsIdentifier,
	// Imported using event_bus_name/statement_id
	"aws_cloudwatch_event_permission": FormattedIdentifierFromProvider("/", "event_bus_name", "statement_id"),
	// Imported using event_bus_name/rule_name
	"aws_cloudwatch_event_rule": FormattedIdentifierUserDefinedNameLast("name", "/", "event_bus_name"),
	// Imported using event_bus_name/rule_name/target_id
	"aws_cloudwatch_event_target": FormattedIdentifierFromProvider("/", "event_bus_name", "rule", "target_id"),

	// cloudwatchlogs
	//
	// CloudWatch Logs destinations can be imported using the name
	"aws_cloudwatch_log_destination": config.NameAsIdentifier,
	// CloudWatch Logs destination policies can be imported using the destination_name
	"aws_cloudwatch_log_destination_policy": config.ParameterAsIdentifier("destination_name"),
	// Cloudwatch Log Groups can be imported using the name
	"aws_cloudwatch_log_group": config.NameAsIdentifier,
	// CloudWatch Log Metric Filter can be imported using the log_group_name:name
	"aws_cloudwatch_log_metric_filter": config.NameAsIdentifier,
	// CloudWatch log resource policies can be imported using the policy name
	"aws_cloudwatch_log_resource_policy": config.ParameterAsIdentifier("policy_name"),
	// Cloudwatch Log Stream can be imported using the stream's log_group_name and name
	"aws_cloudwatch_log_stream": config.IdentifierFromProvider,
	// CloudWatch Logs subscription filter can be imported using the log group name and subscription filter name separated by |
	"aws_cloudwatch_log_subscription_filter": config.IdentifierFromProvider,
	// CloudWatch query definitions can be imported using the query definition ARN.
	"aws_cloudwatch_query_definition": config.IdentifierFromProvider,

	// codeartifact
	//
	// CodeArtifact Domain can be imported using the CodeArtifact Domain arn
	"aws_codeartifact_domain": config.TemplatedStringAsIdentifier("", fullARNTemplate("codeartifact", "domain/{{ .external_name }}")),
	// CodeArtifact Domain Permissions Policies can be imported using the CodeArtifact Domain ARN
	"aws_codeartifact_domain_permissions_policy": config.TemplatedStringAsIdentifier("", fullARNTemplate("codeartifact", "domain/{{ .parameters.domain }}")),
	// CodeArtifact Repository can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository": config.TemplatedStringAsIdentifier("", fullARNTemplate("codeartifact", "repository/{{ .parameters.domain }}/{{ .external_name }}")),
	// CodeArtifact Repository Permissions Policies can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository_permissions_policy": config.TemplatedStringAsIdentifier("", fullARNTemplate("codeartifact", "repository/{{ .parameters.domain }}/{{ .parameters.repository }}")),

	// codecommit
	//
	// CodeCommit approval rule templates can be imported using the name
	"aws_codecommit_approval_rule_template": config.NameAsIdentifier,
	// CodeCommit approval rule template associations can be imported using the approval_rule_template_name and repository_name separated by a comma (,)
	"aws_codecommit_approval_rule_template_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.approval_rule_template_name }},{{ .parameters.repository_name }}"),
	// Codecommit repository can be imported using repository name
	"aws_codecommit_repository": config.ParameterAsIdentifier("repository_name"),
	// No import
	"aws_codecommit_trigger": config.IdentifierFromProvider,

	// codepipeline
	//
	// CodePipelines can be imported using the name
	"aws_codepipeline": config.NameAsIdentifier,
	// CodeDeploy CustomActionType can be imported using the id
	"aws_codepipeline_custom_action_type": config.IdentifierFromProvider,
	// CodePipeline Webhooks can be imported by their ARN: arn:aws:codepipeline:us-west-2:123456789012:webhook:example
	"aws_codepipeline_webhook": config.TemplatedStringAsIdentifier("name", fullARNTemplate("codepipeline", "webhook:{{ .external_name }}")),

	// codestarconnections
	//
	// CodeStar connections can be imported using the ARN
	"aws_codestarconnections_connection": config.IdentifierFromProvider,
	// CodeStar Host can be imported using the ARN
	"aws_codestarconnections_host": config.IdentifierFromProvider,

	// codestarnotifications
	//
	// CodeStar notification rule can be imported using the ARN
	"aws_codestarnotifications_notification_rule": config.IdentifierFromProvider,

	// cognitoidentity
	//
	// us-west-2_abc123
	"aws_cognito_identity_pool": config.IdentifierFromProvider,
	// us-west-2_abc123:CorpAD
	"aws_cognito_identity_pool_provider_principal_tag": config.IdentifierFromProvider,
	// us-west-2:b64805ad-cb56-40ba-9ffc-f5d8207e6d42
	"aws_cognito_identity_pool_roles_attachment": config.IdentifierFromProvider,

	// cognitoidp
	//
	// us-west-2_abc123:CorpAD
	"aws_cognito_identity_provider": config.IdentifierFromProvider,
	// us-west-2_abc123|https://example.com
	"aws_cognito_resource_server": config.IdentifierFromProvider,
	// Cognito Risk Configurations can be imported using the id
	"aws_cognito_risk_configuration": config.IdentifierFromProvider,
	// user_pool_id/name: us-east-1_vG78M4goG/user
	"aws_cognito_user": config.TemplatedStringAsIdentifier("username", "{{ .parameters.user_pool_id }}/{{ .external_name }}"),
	// Cognito User Groups can be imported using the user_pool_id/name attributes concatenated:
	// us-east-1_vG78M4goG/user-group
	// Following configuration does not work: FormattedIdentifierUserDefinedNameLast("name", "/", "user_pool_id")
	// As it fails with a user group not found sync error
	// TODO: check if this is due to any diff between Terraform import & apply
	// implementations. Currently, the API is not normalized.
	"aws_cognito_user_group": config.IdentifierFromProvider,
	// no doc
	"aws_cognito_user_in_group": config.IdentifierFromProvider,
	// us-west-2_abc123
	"aws_cognito_user_pool": config.IdentifierFromProvider,
	// auth.example.org
	"aws_cognito_user_pool_domain": config.IdentifierFromProvider,
	// us-west-2_ZCTarbt5C,12bu4fuk3mlgqa2rtrujgp6egq
	"aws_cognito_user_pool_ui_customization": config.IdentifierFromProvider,

	// configservice
	//
	// Config Rule can be imported using the name
	"aws_config_config_rule": config.NameAsIdentifier,
	// Configuration Aggregators can be imported using the name
	"aws_config_configuration_aggregator": config.NameAsIdentifier,
	// Configuration Recorder can be imported using the name
	"aws_config_configuration_recorder": config.NameAsIdentifier,
	// Configuration Recorder Status can be imported using the name of the Configuration Recorder
	"aws_config_configuration_recorder_status": config.NameAsIdentifier,
	// Config Conformance Packs can be imported using the name
	"aws_config_conformance_pack": config.NameAsIdentifier,
	// Delivery Channel can be imported using the name
	"aws_config_delivery_channel": config.NameAsIdentifier,
	// Remediation Configurations can be imported using the name config_rule_name
	"aws_config_remediation_configuration": config.ParameterAsIdentifier("config_rule_name"),

	// connect
	//
	// aws_connect_bot_association can be imported by using the Amazon Connect instance ID, Lex (V1) bot name, and Lex (V1) bot region separated by colons (:)
	// TODO: lex_bot.lex_region parameter is not `Required` in TF schema. But we use this field in id construction. So, please mark as required this field while configuration
	"aws_connect_bot_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.instance_id }}:{{ (index .parameters.lex_bot 0).name }}:{{ (index .parameters.lex_bot 0).lex_region }}"),
	// Amazon Connect Contact Flows can be imported using the instance_id and contact_flow_id separated by a colon (:)
	"aws_connect_contact_flow": config.IdentifierFromProvider,
	// Amazon Connect Contact Flow Modules can be imported using the instance_id and contact_flow_module_id separated by a colon (:)
	"aws_connect_contact_flow_module": config.IdentifierFromProvider,
	// Amazon Connect Hours of Operations can be imported using the instance_id and hours_of_operation_id separated by a colon (:)
	"aws_connect_hours_of_operation": config.IdentifierFromProvider,
	// Connect instances can be imported using the id
	"aws_connect_instance": config.IdentifierFromProvider,
	// Amazon Connect Instance Storage Configs can be imported using the instance_id, association_id, and resource_type separated by a colon (:)
	// Example: f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5:CHAT_TRANSCRIPTS
	"aws_connect_instance_storage_config": config.IdentifierFromProvider,
	// aws_connect_lambda_function_association can be imported using the instance_id and function_arn separated by a comma (,)
	"aws_connect_lambda_function_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.instance_id }},{{ .parameters.function_arn }}"),
	// Amazon Connect Phone Numbers can be imported using its id
	"aws_connect_phone_number": config.IdentifierFromProvider,
	// Amazon Connect Queues can be imported using the instance_id and queue_id separated by a colon (:)
	"aws_connect_queue": config.IdentifierFromProvider,
	// Amazon Connect Quick Connects can be imported using the instance_id and quick_connect_id separated by a colon (:)
	"aws_connect_quick_connect": config.IdentifierFromProvider,
	// Amazon Connect Routing Profiles can be imported using the instance_id and routing_profile_id separated by a colon (:)
	"aws_connect_routing_profile": config.IdentifierFromProvider,
	// Amazon Connect Security Profiles can be imported using the instance_id and security_profile_id separated by a colon (:)
	"aws_connect_security_profile": config.IdentifierFromProvider,
	// Amazon Connect Users can be imported using the instance_id and user_id separated by a colon (:)
	// Example: f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
	"aws_connect_user": config.IdentifierFromProvider,
	// Amazon Connect User Hierarchy Structures can be imported using the instance_id
	"aws_connect_user_hierarchy_structure": config.IdentifierFromProvider,
	// Amazon Connect Vocabularies can be imported using the instance_id and vocabulary_id separated by a colon (:)
	// Example: f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
	"aws_connect_vocabulary": config.IdentifierFromProvider,

	// cur
	//
	// Report Definitions can be imported using the report_name
	"aws_cur_report_definition": config.ParameterAsIdentifier("report_name"),

	// dataexchange
	//
	// DataExchange DataSets can be imported by their arn
	"aws_dataexchange_data_set": config.IdentifierFromProvider,
	// DataExchange Revisions can be imported by their data-set-id:revision-id
	"aws_dataexchange_revision": config.IdentifierFromProvider,

	// datapipeline
	//
	// aws_datapipeline_pipeline can be imported by using the id (Pipeline ID)
	"aws_datapipeline_pipeline": config.IdentifierFromProvider,

	// datasync
	//
	// aws_datasync_location_s3 can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_s3": config.IdentifierFromProvider,
	// aws_datasync_task can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_task": config.IdentifierFromProvider,

	// dax
	//
	// DAX Clusters can be imported using the cluster_name
	"aws_dax_cluster": config.ParameterAsIdentifier("cluster_name"),
	// DAX Parameter Group can be imported using the name
	"aws_dax_parameter_group": config.NameAsIdentifier,
	// DAX Subnet Group can be imported using the name
	"aws_dax_subnet_group": config.NameAsIdentifier,

	// deploy
	//
	// CodeDeploy Applications can be imported using the name
	"aws_codedeploy_app": config.TemplatedStringAsIdentifier("name", "{{ .parameters.application_id }}:{{ .external_name }}"),
	// CodeDeploy Deployment Configurations can be imported using the deployment_config_name
	"aws_codedeploy_deployment_config": config.ParameterAsIdentifier("deployment_config_name"),
	// CodeDeploy Deployment Groups can be imported by their app_name, a colon, and deployment_group_name
	"aws_codedeploy_deployment_group": config.TemplatedStringAsIdentifier("deployment_group_name", "{{ .parameters.app_name }}:{{ .external_name }}"),

	// detective
	//
	// aws_detective_graph can be imported using the ARN
	"aws_detective_graph": config.IdentifierFromProvider,
	// aws_detective_invitation_accepter can be imported using the graph ARN
	"aws_detective_invitation_accepter": config.IdentifierFromProvider,
	// aws_detective_member can be imported using the ARN of the graph followed by the account ID of the member account
	"aws_detective_member": config.IdentifierFromProvider,

	// devicefarm
	//
	// DeviceFarm Device Pools can be imported by their arn
	"aws_devicefarm_device_pool": config.IdentifierFromProvider,
	// DeviceFarm Instance Profiles can be imported by their arn
	"aws_devicefarm_instance_profile": config.IdentifierFromProvider,
	// DeviceFarm Network Profiles can be imported by their arn
	"aws_devicefarm_network_profile": config.IdentifierFromProvider,
	// DeviceFarm Projects can be imported by their arn
	"aws_devicefarm_project": config.IdentifierFromProvider,
	// DeviceFarm Test Grid Projects can be imported by their arn
	"aws_devicefarm_test_grid_project": config.IdentifierFromProvider,
	// DeviceFarm Uploads can be imported by their arn
	"aws_devicefarm_upload": config.IdentifierFromProvider,

	// directconnect
	//
	// No import
	"aws_dx_bgp_peer": config.IdentifierFromProvider,
	// Direct Connect connections can be imported using the connection id
	"aws_dx_connection": config.IdentifierFromProvider,
	// No import
	"aws_dx_connection_association": config.IdentifierFromProvider,
	// Direct Connect Gateways can be imported using the gateway id
	"aws_dx_gateway": config.IdentifierFromProvider,
	// Direct Connect gateway associations can be imported using dx_gateway_id together with associated_gateway_id
	// TODO: associated_gateway_id parameter is not `Required` in TF schema. But we use this field in id construction. So, please mark as required this field while configuration
	"aws_dx_gateway_association": config.IdentifierFromProvider,
	//
	"aws_dx_gateway_association_proposal": config.IdentifierFromProvider,
	// Direct Connect hosted private virtual interfaces can be imported using the vif id
	"aws_dx_hosted_private_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted private virtual interfaces can be imported using the vif id
	"aws_dx_hosted_private_virtual_interface_accepter": config.IdentifierFromProvider,
	// Direct Connect hosted public virtual interfaces can be imported using the vif id
	"aws_dx_hosted_public_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted public virtual interfaces can be imported using the vif id
	"aws_dx_hosted_public_virtual_interface_accepter": config.IdentifierFromProvider,
	// Direct Connect hosted transit virtual interfaces can be imported using the vif id
	"aws_dx_hosted_transit_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted transit virtual interfaces can be imported using the vif id
	"aws_dx_hosted_transit_virtual_interface_accepter": config.IdentifierFromProvider,
	// Direct Connect LAGs can be imported using the lag id
	"aws_dx_lag": config.IdentifierFromProvider,
	// Direct Connect private virtual interfaces can be imported using the vif id
	"aws_dx_private_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect public virtual interfaces can be imported using the vif id
	"aws_dx_public_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect transit virtual interfaces can be imported using the vif id
	"aws_dx_transit_virtual_interface": config.IdentifierFromProvider,

	// dlm
	//
	// DLM lifecycle policies can be imported by their policy ID
	"aws_dlm_lifecycle_policy": config.IdentifierFromProvider,

	// dms
	//
	// Certificates can be imported using the certificate_id
	"aws_dms_certificate": config.ParameterAsIdentifier("certificate_id"),
	// Endpoints can be imported using the endpoint_id
	"aws_dms_endpoint": config.ParameterAsIdentifier("endpoint_id"),
	// Event subscriptions can be imported using the name
	"aws_dms_event_subscription": config.NameAsIdentifier,
	// Replication instances can be imported using the replication_instance_id
	"aws_dms_replication_instance": config.ParameterAsIdentifier("replication_instance_id"),
	// Replication subnet groups can be imported using the replication_subnet_group_id
	"aws_dms_replication_subnet_group": config.ParameterAsIdentifier("replication_subnet_group_id"),
	// Replication tasks can be imported using the replication_task_id
	"aws_dms_replication_task": config.ParameterAsIdentifier("replication_task_id"),
	// Endpoints can be imported using the endpoint_id
	"aws_dms_s3_endpoint": config.ParameterAsIdentifier("endpoint_id"),

	// docdb
	//
	// DocDB Clusters can be imported using the cluster_identifier
	"aws_docdb_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// DocDB Cluster Instances can be imported using the identifier
	"aws_docdb_cluster_instance": config.ParameterAsIdentifier("identifier"),
	// DocumentDB Cluster Parameter Groups can be imported using the name
	"aws_docdb_cluster_parameter_group": config.NameAsIdentifier,
	// aws_docdb_cluster_snapshot can be imported by using the cluster snapshot identifier
	"aws_docdb_cluster_snapshot": config.ParameterAsIdentifier("db_cluster_snapshot_identifier"),
	// DocDB Event Subscriptions can be imported using the name
	"aws_docdb_event_subscription": config.NameAsIdentifier,
	// aws_docdb_global_cluster can be imported by using the Global Cluster id
	"aws_docdb_global_cluster": config.IdentifierFromProvider,
	// DocumentDB Subnet groups can be imported using the name
	"aws_docdb_subnet_group": config.NameAsIdentifier,

	// ds
	//
	// Conditional forwarders can be imported using the directory id and remote_domain_name: d-1234567890:example.com
	"aws_directory_service_conditional_forwarder": config.TemplatedStringAsIdentifier("", "{{ .parameters.directory_id }}:{{ .parameters.remote_domain_name }}"),
	// DirectoryService directories can be imported using the directory id
	"aws_directory_service_directory": config.IdentifierFromProvider,
	// Directory Service Shared Directories can be imported using the owner directory ID/shared directory ID
	// "aws_directory_service_shared_directory": config.TemplatedStringAsIdentifier("", "{{ .parameters.directory_id }}/{{ .external_name }}"),
	"aws_directory_service_shared_directory": config.IdentifierFromProvider,

	// dynamodb
	//
	// DynamoDB contributor insights
	"aws_dynamodb_contributor_insights": config.IdentifierFromProvider,
	// DynamoDB Global Tables can be imported using the global table name
	"aws_dynamodb_global_table": config.NameAsIdentifier,
	// Dynamodb Kinesis streaming destinations are imported using "table_name,stream_arn"
	"aws_dynamodb_kinesis_streaming_destination": config.IdentifierFromProvider,
	// DynamoDB tables can be imported using the name
	"aws_dynamodb_table": config.NameAsIdentifier,
	// DynamoDB Table Items can be imported using the name
	"aws_dynamodb_table_item": config.IdentifierFromProvider,
	// DynamoDB table replicas can be imported using the table-name:main-region
	"aws_dynamodb_table_replica": config.IdentifierFromProvider,
	// aws_dynamodb_tag can be imported by using the DynamoDB resource identifier and key, separated by a comma (,)
	"aws_dynamodb_tag": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_arn }},{{ .parameters.key }}"),

	// ec2
	//
	// aws_ami can be imported using the ID of the AMI
	"aws_ami": config.IdentifierFromProvider,
	// No import
	"aws_ami_copy": config.IdentifierFromProvider,
	// AMI Launch Permissions can be imported using [ACCOUNT-ID|GROUP-NAME|ORGANIZATION-ARN|ORGANIZATIONAL-UNIT-ARN]/IMAGE-ID
	"aws_ami_launch_permission": config.IdentifierFromProvider,
	// Customer Gateways can be imported using the id
	"aws_customer_gateway": config.IdentifierFromProvider,
	// Default Network ACLs can be imported using the id
	"aws_default_network_acl": config.IdentifierFromProvider,
	// Default VPC route tables can be imported using the vpc_id
	"aws_default_route_table": config.IdentifierFromProvider,
	// Security Groups can be imported using the security group id
	"aws_default_security_group": config.IdentifierFromProvider,
	// Subnets can be imported using the subnet id
	"aws_default_subnet": config.IdentifierFromProvider,
	// Default VPCs can be imported using the vpc id
	"aws_default_vpc": config.IdentifierFromProvider,
	// VPC DHCP Options can be imported using the dhcp options id
	"aws_default_vpc_dhcp_options": config.IdentifierFromProvider,
	// The EBS default KMS CMK can be imported with the KMS key ARN
	"aws_ebs_default_kms_key": config.IdentifierFromProvider,
	// Default EBS encryption state can be imported
	"aws_ebs_encryption_by_default": config.IdentifierFromProvider,
	// EBS Snapshot can be imported using the id
	"aws_ebs_snapshot": config.IdentifierFromProvider,
	// No import
	"aws_ebs_snapshot_copy": config.IdentifierFromProvider,
	// No import
	"aws_ebs_snapshot_import": config.IdentifierFromProvider,
	// EBS Volumes can be imported using the id: vol-049df61146c4d7901
	"aws_ebs_volume": config.IdentifierFromProvider,
	// EC2 Availability Zone Groups can be imported using the group name
	"aws_ec2_availability_zone_group": config.ParameterAsIdentifier("group_name"),
	// Capacity Reservations can be imported using the id
	"aws_ec2_capacity_reservation": config.IdentifierFromProvider,
	// aws_ec2_carrier_gateway can be imported using the carrier gateway's ID
	"aws_ec2_carrier_gateway": config.IdentifierFromProvider,
	// aws_ec2_instance_state can be imported by using the instance_id attribute
	"aws_ec2_instance_state": config.IdentifierFromProvider,
	// aws_ec2_fleet can be imported by using the Fleet identifier
	"aws_ec2_fleet": config.IdentifierFromProvider,
	// Network Insights Analyses can be imported using the id
	"aws_ec2_network_insights_analysis": config.IdentifierFromProvider,
	// Prefix Lists can be imported using the id
	"aws_ec2_managed_prefix_list": config.IdentifierFromProvider,
	// Prefix List Entries can be imported using the prefix_list_id and cidr separated by a ,
	"aws_ec2_managed_prefix_list_entry": FormattedIdentifierFromProvider(",", "prefix_list_id", "cidr"),
	// Network Insights Paths can be imported using the id
	"aws_ec2_network_insights_path": config.IdentifierFromProvider,
	// Serial console access state can be imported
	"aws_ec2_serial_console_access": config.IdentifierFromProvider,
	// Existing CIDR reservations can be imported using SUBNET_ID:RESERVATION_ID
	"aws_ec2_subnet_cidr_reservation": config.IdentifierFromProvider,
	// aws_ec2_tag can be imported by using the EC2 resource identifier and key, separated by a comma (,)
	"aws_ec2_tag": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_id }},{{ .parameters.key }}"),
	// Traffic mirror filter can be imported using the id
	"aws_ec2_traffic_mirror_filter": config.IdentifierFromProvider,
	// Traffic mirror rules can be imported using the traffic_mirror_filter_id and id separated by :
	"aws_ec2_traffic_mirror_filter_rule": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway identifier: tgw-12345678
	"aws_ec2_transit_gateway": config.IdentifierFromProvider,
	// Traffic mirror targets can be imported using the id
	"aws_ec2_transit_gateway_connect": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_connect_peer can be imported by using the EC2 Transit Gateway Connect Peer identifier
	"aws_ec2_transit_gateway_connect_peer": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_multicast_domain can be imported by using the EC2 Transit Gateway Multicast Domain identifier
	"aws_ec2_transit_gateway_multicast_domain": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_domain_association": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_group_member": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_group_source": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_peering_attachment can be imported by using the EC2 Transit Gateway Attachment identifier
	"aws_ec2_transit_gateway_peering_attachment": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_peering_attachment_accepter can be imported by using the EC2 Transit Gateway Attachment identifier
	"aws_ec2_transit_gateway_peering_attachment_accepter": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_policy_table can be imported by using the EC2 Transit Gateway Policy Table identifier
	"aws_ec2_transit_gateway_policy_table": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_prefix_list_reference can be imported by using the EC2 Transit Gateway Route Table identifier and EC2 Prefix List identifier, separated by an underscore (_
	"aws_ec2_transit_gateway_prefix_list_reference": FormattedIdentifierFromProvider("_", "transit_gateway_route_table_id", "prefix_list_id"),
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
	// Imported by using the EC2 Transit Gateway Route Table identifier, an
	// underscore, and the EC2 Transit Gateway Attachment identifier:
	// tgw-rtb-12345678_tgw-attach-87654321
	"aws_ec2_transit_gateway_route_table_propagation": FormattedIdentifierFromProvider("_", "transit_gateway_route_table_id", "transit_gateway_attachment_id"),
	// Imported by using the EC2 Transit Gateway Attachment identifier:
	// tgw-attach-12345678
	"aws_ec2_transit_gateway_vpc_attachment": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Attachment identifier: tgw-attach-12345678
	"aws_ec2_transit_gateway_vpc_attachment_accepter": config.IdentifierFromProvider,
	// Egress-only Internet gateways can be imported using the id
	"aws_egress_only_internet_gateway": config.IdentifierFromProvider,
	// No terraform import.
	"aws_eip": config.IdentifierFromProvider,
	// EIP Assocations can be imported using their association ID.
	"aws_eip_association": config.IdentifierFromProvider,
	// Flow Logs can be imported using the id
	"aws_flow_log": config.IdentifierFromProvider,
	// Hosts can be imported using the host id
	"aws_ec2_host": config.IdentifierFromProvider,
	// Instances can be imported using the id: i-12345678
	"aws_instance": config.IdentifierFromProvider,
	// Imported using the id: igw-c0a643a9
	"aws_internet_gateway": config.IdentifierFromProvider,
	// Key Pairs can be imported using the key_name
	"aws_key_pair": config.ParameterAsIdentifier("key_name"),
	// Launch configurations can be imported using the name
	"aws_launch_configuration": config.NameAsIdentifier,
	// Imported using the id: lt-12345678
	"aws_launch_template": config.IdentifierFromProvider,
	// No import.
	"aws_main_route_table_association": config.IdentifierFromProvider,
	// NAT Gateways can be imported using the id
	"aws_nat_gateway": config.IdentifierFromProvider,
	// Network ACLs can be imported using the id
	"aws_network_acl": config.IdentifierFromProvider,
	// Individual rules can be imported using NETWORK_ACL_ID:RULE_NUMBER:PROTOCOL:EGRESS
	"aws_network_acl_rule": config.IdentifierFromProvider,
	// Imported using the id: eni-e5aa89a3
	"aws_network_interface": config.IdentifierFromProvider,
	// No import
	"aws_network_interface_attachment": config.IdentifierFromProvider,
	// No import
	"aws_network_interface_sg_attachment": config.IdentifierFromProvider,
	// Placement groups can be imported using the name
	"aws_placement_group": config.NameAsIdentifier,
	// Imported using the following format: ROUTETABLEID_DESTINATION
	"aws_route": route(),
	// Imported using id: rtb-4e616f6d69
	"aws_route_table": config.IdentifierFromProvider,
	// Imported using the associated resource ID and Route Table ID separated
	// by a forward slash (/)
	"aws_route_table_association": config.IdentifierFromProvider,
	// Imported using the id: sg-903004f8
	"aws_security_group": config.IdentifierFromProvider,
	// Imported using a very complex format:
	// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule
	"aws_security_group_rule": config.IdentifierFromProvider,
	// A Spot Datafeed Subscription can be imported using the word spot-datafeed-subscription
	"aws_spot_datafeed_subscription": config.IdentifierFromProvider,
	// Spot Fleet Requests can be imported using id
	"aws_spot_fleet_request": config.IdentifierFromProvider,
	// No import
	"aws_spot_instance_request": config.IdentifierFromProvider,
	// No import
	"aws_snapshot_create_volume_permission": config.IdentifierFromProvider,
	// Imported using the subnet id: subnet-9d4a7b6c
	"aws_subnet": config.IdentifierFromProvider,
	// Imported using the id: vpc-23123
	"aws_vpc": config.IdentifierFromProvider,
	// VPC DHCP Options can be imported using the dhcp options id
	"aws_vpc_dhcp_options": config.IdentifierFromProvider,
	// DHCP associations can be imported by providing the VPC ID associated with the options
	// terraform import aws_vpc_dhcp_options_association.imported vpc-0f001273ec18911b1
	"aws_vpc_dhcp_options_association": config.IdentifierFromProvider,
	// Imported using the vpc endpoint id: vpce-3ecf2a57
	"aws_vpc_endpoint": config.IdentifierFromProvider,
	// VPC Endpoint Services can be imported using the VPC endpoint service id
	"aws_vpc_endpoint_service": config.IdentifierFromProvider,
	// No import
	"aws_vpc_endpoint_service_allowed_principal": config.IdentifierFromProvider,
	// VPC Endpoint connection notifications can be imported using the VPC endpoint connection notification id
	"aws_vpc_endpoint_connection_notification": config.IdentifierFromProvider,
	// VPC Endpoint Route Table Associations can be imported using vpc_endpoint_id together with route_table_id
	"aws_vpc_endpoint_route_table_association": FormattedIdentifierFromProvider("/", "vpc_endpoint_id", "route_table_id"),
	// VPC Endpoint Subnet Associations can be imported using vpc_endpoint_id together with subnet_id
	"aws_vpc_endpoint_subnet_association": FormattedIdentifierFromProvider("/", "vpc_endpoint_id", "subnet_id"),
	// VPC Endpoint security group Associations can be imported using vpc_endpoint_id together with security_group_id
	"aws_vpc_endpoint_security_group_association": config.IdentifierFromProvider,
	// IPAMs can be imported using the ipam id
	"aws_vpc_ipam": config.IdentifierFromProvider,
	// IPAMs can be imported using the ipam pool id
	"aws_vpc_ipam_pool": config.IdentifierFromProvider,
	// IPAMs can be imported using the ipam id
	"aws_vpc_ipam_pool_cidr": config.IdentifierFromProvider,
	// IPAMs can be imported using the allocation id
	"aws_vpc_ipam_pool_cidr_allocation": config.IdentifierFromProvider,
	// IPAMs can be imported using the scope_id
	"aws_vpc_ipam_scope": config.IdentifierFromProvider,
	// Imported by using the VPC CIDR Association ID: vpc-cidr-assoc-xxxxxxxx
	"aws_vpc_ipv4_cidr_block_association": config.IdentifierFromProvider,
	// Imported using the vpc peering id: pcx-111aaa111
	"aws_vpc_peering_connection": config.IdentifierFromProvider,
	// Imported using the peering connection id: pcx-12345678
	"aws_vpc_peering_connection_accepter": config.IdentifierFromProvider,
	// Imported using the vpc peering id: pcx-111aaa111
	"aws_vpc_peering_connection_options": config.IdentifierFromProvider,
	// EBS Volume Attachments can be imported using DEVICE_NAME:VOLUME_ID:INSTANCE_ID
	"aws_volume_attachment": config.IdentifierFromProvider,
	// VPN Connections can be imported using the vpn connection id
	"aws_vpn_connection": config.IdentifierFromProvider,
	// No import
	"aws_vpn_connection_route": config.IdentifierFromProvider,
	// VPN Gateways can be imported using the vpn gateway id
	"aws_vpn_gateway": config.IdentifierFromProvider,
	// No import
	"aws_vpn_gateway_attachment": config.IdentifierFromProvider,
	// No import
	"aws_vpn_gateway_route_propagation": config.IdentifierFromProvider,

	// verified access
	// VerifiedAccess Endpoint can be imported using the endpoint id (vae): vae-8012925589
	"aws_verifiedaccess_endpoint": config.IdentifierFromProvider,
	// No import
	"aws_verifiedaccess_group": config.IdentifierFromProvider,
	// VerifiedAccess Instance can be imported using the instance id (vai): vae-8012925589
	"aws_verifiedaccess_instance": config.IdentifierFromProvider,
	// VerifiedAccess Instance Logging Configuration can be imported using the instance id (vai): vai-1234567890abcdef0
	"aws_verifiedaccess_instance_logging_configuration": config.IdentifierFromProvider,
	// VerifiedAccess Instance Trust provider Attachment can be imported using the instance/attachment id: vai-1234567890abcdef0/vatp-801292558
	"aws_verifiedaccess_instance_trust_provider_attachment": config.IdentifierFromProvider,
	// VerifiedAccess TrustProvider can be imported using the trust provider id (vatp): vatp-8012925589
	"aws_verifiedaccess_trust_provider": config.IdentifierFromProvider,

	// ecr
	//
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
	//
	"aws_ecr_repository": config.NameAsIdentifier,
	// Imported using the parameter called repository but this is not the name
	// of the resource, only a configuration/reference.
	"aws_ecr_repository_policy": config.IdentifierFromProvider,

	// ecrpublic
	//
	"aws_ecrpublic_repository": config.ParameterAsIdentifier("repository_name"),
	// Imported using the repository name.
	"aws_ecrpublic_repository_policy": config.IdentifierFromProvider,

	// ecs
	//
	// ECS Account Setting defaults can be imported using the name
	"aws_ecs_account_setting_default": config.IdentifierFromProvider,
	//
	"aws_ecs_capacity_provider": config.TemplatedStringAsIdentifier("name", fullARNTemplate("ecs", "capacity-provider/{{ .external_name }}")),
	//
	"aws_ecs_cluster": config.TemplatedStringAsIdentifier("name", fullARNTemplate("ecs", "cluster/{{ .external_name }}")),
	// ECS cluster capacity providers can be imported using the cluster_name attribute
	"aws_ecs_cluster_capacity_providers": config.IdentifierFromProvider,
	//
	"aws_ecs_service": config.NameAsIdentifier,
	// Imported using ARN that has a random substring, revision at the end:
	// arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123
	"aws_ecs_task_definition": config.IdentifierFromProvider,

	// efs
	//
	// The EFS access points can be imported using the id
	"aws_efs_access_point": config.IdentifierFromProvider,
	// The EFS backup policies can be imported using the id
	"aws_efs_backup_policy": config.IdentifierFromProvider,
	// The EFS file systems can be imported using the id
	"aws_efs_file_system": config.IdentifierFromProvider,
	// The EFS file system policies can be imported using the id
	"aws_efs_file_system_policy": config.IdentifierFromProvider,
	// The EFS mount targets can be imported using the id
	"aws_efs_mount_target": config.IdentifierFromProvider,
	// EFS Replication Configurations can be imported using the file system ID of either the source or destination file system
	"aws_efs_replication_configuration": config.IdentifierFromProvider,

	// eks
	//
	// import EKS access entry using the cluster_name and principal_arn separated by a colon (:).
	"aws_eks_access_entry": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_name }}:{{ .parameters.principal_arn }}"),
	// import EKS access entry using the cluster_name principal_arn and policy_arn separated by a (#) which the tf provider docs incorrectly describe as a colon.
	"aws_eks_access_policy_association": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_name }}#{{ .parameters.principal_arn }}#{{ .parameters.policy_arn }}"),
	// "aws_eks_addon": config.TemplatedStringAsIdentifier("addon_name", "{{ .parameters.cluster_name }}:{{ .external_name }}"),
	// my_cluster:my_eks_addon
	"aws_eks_addon": FormattedIdentifierFromProvider(":", "cluster_name", "addon_name"),
	// import EKS cluster using the name.
	"aws_eks_cluster": config.NameAsIdentifier,
	// my_cluster:my_fargate_profile
	"aws_eks_fargate_profile": FormattedIdentifierUserDefinedNameLast("fargate_profile_name", ":", "cluster_name"),
	// It has a complex config, adding empty entry here just to enable it.
	"aws_eks_identity_provider_config": eksOIDCIdentityProvider(),
	// Imported using the cluster_name and node_group_name separated by a
	// colon (:): my_cluster:my_node_group
	"aws_eks_node_group": config.TemplatedStringAsIdentifier("node_group_name", "{{ .parameters.cluster_name }}:{{ .external_name }}"),

	// elasticache
	//
	"aws_elasticache_cluster": config.ParameterAsIdentifier("cluster_id"),
	// ElastiCache Global Replication Groups can be imported using the global_replication_group_id
	"aws_elasticache_global_replication_group": config.IdentifierFromProvider,
	"aws_elasticache_parameter_group":          config.IdentifierFromProvider,
	"aws_elasticache_replication_group":        config.ParameterAsIdentifier("replication_group_id"),
	"aws_elasticache_subnet_group":             config.NameAsIdentifier,
	"aws_elasticache_user":                     config.ParameterAsIdentifier("user_id"),
	"aws_elasticache_user_group":               config.ParameterAsIdentifier("user_group_id"),

	// elasticbeanstalk
	//
	// Elastic Beanstalk Applications can be imported using the name
	"aws_elastic_beanstalk_application": config.NameAsIdentifier,
	// Elastic Beanstalk Applications can be imported using the name
	"aws_elastic_beanstalk_application_version": config.NameAsIdentifier,
	// No import
	"aws_elastic_beanstalk_configuration_template": config.NameAsIdentifier,

	// elasticsearch
	//
	// Elasticsearch domains can be imported using the domain_name
	"aws_elasticsearch_domain": config.TemplatedStringAsIdentifier("domain_name", fullARNTemplate("es", "domain/{{ .external_name }}")),
	// No import
	"aws_elasticsearch_domain_policy": config.IdentifierFromProvider,
	// Elasticsearch domains can be imported using the domain_name
	"aws_elasticsearch_domain_saml_options": config.ParameterAsIdentifier("domain_name"),

	// elastictranscoder
	//
	// Elastic Transcoder pipelines can be imported using the id
	"aws_elastictranscoder_pipeline": config.IdentifierFromProvider,
	// Elastic Transcoder presets can be imported using the id
	"aws_elastictranscoder_preset": config.IdentifierFromProvider,

	// elb
	//
	// Application cookie stickiness policies can be imported using the ELB name, port, and policy name separated by colons (:)
	// my-elb:80:my-policy
	"aws_app_cookie_stickiness_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.load_balancer }}:{{ .parameters.lb_port }}:{{ .external_name }}"),
	// ELBs can be imported using the name
	"aws_elb": config.NameAsIdentifier,
	// No import
	"aws_elb_attachment": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lb_cookie_stickiness_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lb_ssl_negotiation_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_load_balancer_backend_server_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_load_balancer_listener_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_load_balancer_policy": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_proxy_protocol_policy": config.IdentifierFromProvider,

	// elbv2
	//
	// arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188
	"aws_lb": config.IdentifierFromProvider,
	// arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
	"aws_lb_listener": config.IdentifierFromProvider,
	// Listener Certificates can be imported by using the listener arn and certificate arn, separated by an underscore (_)
	// arn:aws:elasticloadbalancing:us-west-2:123456789012:listener/app/test/8e4497da625e2d8a/9ab28ade35828f96/67b3d2d36dd7c26b_arn:aws:iam::123456789012:server-certificate/tf-acc-test-6453083910015726063
	"aws_lb_listener_certificate": config.IdentifierFromProvider,
	// Rules can be imported using their ARN
	"aws_lb_listener_rule": config.IdentifierFromProvider,
	// arn:aws:elasticloadbalancing:us-west-2:187416307283:targetgroup/app-front-end/20cfe21448b66314
	"aws_lb_target_group": config.IdentifierFromProvider,
	// No import.
	"aws_lb_target_group_attachment": config.IdentifierFromProvider,
	// Trust Stores can be imported using their ARN
	"aws_lb_trust_store": config.IdentifierFromProvider,

	// emr
	//
	// EMR Security Configurations can be imported using the name
	"aws_emr_security_configuration": config.NameAsIdentifier,

	// emrserverless
	//
	// EMR Severless applications can be imported using the id
	"aws_emrserverless_application": config.IdentifierFromProvider,

	// evidently
	//
	// CloudWatch Evidently Feature can be imported using the feature name and name or arn of the hosting CloudWatch Evidently Project separated by a :
	// Example: exampleFeatureName:arn:aws:evidently:us-east-1:123456789012:project/example
	"aws_evidently_feature": config.TemplatedStringAsIdentifier("name", "{{ .external_name }}:{{ .parameters.project }}"),
	// CloudWatch Evidently Project can be imported using the arn
	// Example: arn:aws:evidently:us-east-1:123456789012:segment/example
	"aws_evidently_project": config.IdentifierFromProvider,
	// CloudWatch Evidently Segment can be imported using the arn
	// Example: arn:aws:evidently:us-west-2:123456789012:segment/example
	"aws_evidently_segment": config.TemplatedStringAsIdentifier("name", fullARNTemplate("evidently", "segment/{{ .external_name }}")),

	// firehose
	//
	"aws_kinesis_firehose_delivery_stream": config.IdentifierFromProvider,

	// fis
	//
	// FIS Experiment Templates can be imported using the id
	"aws_fis_experiment_template": config.IdentifierFromProvider,

	// fsx
	//
	// FSx Backups can be imported using the id
	"aws_fsx_backup": config.IdentifierFromProvider,
	// FSx Data Repository Associations can be imported using the id
	"aws_fsx_data_repository_association": config.IdentifierFromProvider,
	// FSx File Systems can be imported using the id
	"aws_fsx_lustre_file_system": config.IdentifierFromProvider,
	// FSx File Systems can be imported using the id
	"aws_fsx_ontap_file_system": config.IdentifierFromProvider,
	// FSx Storage Virtual Machine can be imported using the id
	"aws_fsx_ontap_storage_virtual_machine": config.IdentifierFromProvider,
	// FSx File Systems can be imported using the id
	"aws_fsx_windows_file_system": config.IdentifierFromProvider,

	// gamelift
	//
	// GameLift Aliases can be imported using the ID
	"aws_gamelift_alias": config.IdentifierFromProvider,
	// GameLift Builds can be imported using the ID
	"aws_gamelift_build": config.IdentifierFromProvider,
	// GameLift Fleets can be imported using the ID
	"aws_gamelift_fleet": config.IdentifierFromProvider,
	// GameLift Game Session Queues can be imported by their name
	"aws_gamelift_game_session_queue": config.NameAsIdentifier,
	// GameLift Scripts can be imported using the ID
	"aws_gamelift_script": config.IdentifierFromProvider,

	// glacier
	//
	// Glacier Vaults can be imported using the name
	"aws_glacier_vault": config.NameAsIdentifier,
	// Glacier Vault Locks can be imported using the Glacier Vault name
	"aws_glacier_vault_lock": FormattedIdentifierFromProvider("", "vault_name"),

	// globalaccelerator
	//
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	"aws_globalaccelerator_accelerator": config.IdentifierFromProvider,
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx
	"aws_globalaccelerator_endpoint_group": config.IdentifierFromProvider,
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx
	"aws_globalaccelerator_listener": config.IdentifierFromProvider,

	// glue
	//
	"aws_glue_catalog_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .external_name }}"),
	//
	"aws_glue_catalog_table": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .parameters.database_name }}:{{ .external_name }}"),
	//
	"aws_glue_classifier": config.NameAsIdentifier,
	// Imported as CATALOG_ID:name 123456789012:MyConnection
	"aws_glue_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .external_name }}"),
	//
	"aws_glue_crawler": config.NameAsIdentifier,
	// Imported using CATALOG-ID (AWS account ID if not custom), e.g., 123456789012
	"aws_glue_data_catalog_encryption_settings": config.IdentifierFromProvider,
	//
	"aws_glue_job": config.NameAsIdentifier,
	// Imported using ARN: arn:aws:glue:us-west-2:123456789012:registry/example
	"aws_glue_registry": config.TemplatedStringAsIdentifier("registry_name", fullARNTemplate("glue", "registry/{{ .external_name }}")),
	// Imported using the account ID: 12356789012
	"aws_glue_resource_policy": config.IdentifierFromProvider,
	// Glue Registries can be imported using arn
	// Example: arn:aws:glue:us-west-2:123456789012:schema/example/example
	"aws_glue_schema": config.IdentifierFromProvider,
	// Imported using "name".
	"aws_glue_security_configuration": config.NameAsIdentifier,
	// Imported using "name".
	"aws_glue_trigger": config.NameAsIdentifier,
	//
	"aws_glue_user_defined_function": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .parameters.database_name }}:{{ .external_name }}"),
	// Imported using "name".
	"aws_glue_workflow": config.NameAsIdentifier,

	// grafana
	//
	// Grafana Workspace can be imported using the workspace's id
	"aws_grafana_workspace": config.IdentifierFromProvider,
	// No import
	"aws_grafana_role_association": config.IdentifierFromProvider,
	// Grafana Workspace SAML configuration can be imported using the workspace's id
	"aws_grafana_workspace_saml_configuration": FormattedIdentifierFromProvider("", "workspace_id"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_grafana_workspace_api_key": config.IdentifierFromProvider,
	// Grafana workspace license association can be imported using the workspace's id
	"aws_grafana_license_association": FormattedIdentifierFromProvider("", "workspace_id"),

	// guardduty
	//
	// GuardDuty detectors can be imported using the detector ID
	"aws_guardduty_detector": config.IdentifierFromProvider,
	// GuardDuty filters can be imported using the detector ID and filter's name separated by a colon
	// 00b00fd5aecc0ab60a708659477e9617:MyFilter
	"aws_guardduty_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.detector_id }}:{{ .external_name }}"),
	// GuardDuty members can be imported using the primary GuardDuty detector ID and member AWS account ID
	// 00b00fd5aecc0ab60a708659477e9617:123456789012
	"aws_guardduty_member": config.IdentifierFromProvider,

	// iam
	//
	// AKIA1234567890
	"aws_iam_access_key": config.IdentifierFromProvider,
	// The current Account Alias can be imported using the account_alias
	"aws_iam_account_alias": config.ParameterAsIdentifier("account_alias"),
	// IAM Account Password Policy can be imported using the word iam-account-password-policy
	"aws_iam_account_password_policy": config.IdentifierFromProvider,
	//
	"aws_iam_group": config.NameAsIdentifier,
	// No import
	"aws_iam_group_membership": config.IdentifierFromProvider,
	// Imported using the group name and policy arn separated by /
	// test-group/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_group_policy_attachment": config.IdentifierFromProvider,
	//
	"aws_iam_instance_profile": config.NameAsIdentifier,
	// arn:aws:iam::123456789012:oidc-provider/accounts.google.com
	"aws_iam_openid_connect_provider": config.IdentifierFromProvider,
	// arn:aws:iam::123456789012:policy/UsersManageOwnCredentials
	"aws_iam_policy": iamPolicy(),
	//
	"aws_iam_role": config.NameAsIdentifier,
	//
	"aws_iam_role_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.role }}:{{ .external_name }}"),
	// Imported using the role name and policy arn separated by /
	// test-role/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_role_policy_attachment": config.IdentifierFromProvider,
	// IAM SAML Providers can be imported using the arn
	"aws_iam_saml_provider": config.TemplatedStringAsIdentifier("name", regionlessARNTemplate("iam", "saml-provider/{{ .external_name }}")),
	// IAM Server Certificates can be imported using the name
	"aws_iam_server_certificate": config.NameAsIdentifier,
	// IAM service-linked roles can be imported using role ARN that contains the
	// service name.
	"aws_iam_service_linked_role": config.IdentifierFromProvider,
	// IAM Service Specific Credentials can be imported using the service_name:user_name:service_specific_credential_id
	"aws_iam_service_specific_credential": config.IdentifierFromProvider,
	// IAM Signing Certificates can be imported using the id
	"aws_iam_signing_certificate": config.IdentifierFromProvider,
	//
	"aws_iam_user": config.NameAsIdentifier,
	// Imported using the user name and group names separated by /
	// user1/group1/group2
	"aws_iam_user_group_membership": iamUserGroupMembership(),
	// IAM User Login Profiles can be imported without password information support via the IAM User name
	"aws_iam_user_login_profile": config.IdentifierFromProvider,
	// Imported using the user name and policy arn separated by /
	// test-user/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_user_policy_attachment": config.IdentifierFromProvider,
	// SSH public keys can be imported using the username, ssh_public_key_id, and encoding
	"aws_iam_user_ssh_key": config.IdentifierFromProvider,
	// IAM Virtual MFA Devices can be imported using the arn
	"aws_iam_virtual_mfa_device": config.IdentifierFromProvider,

	// identitystore
	//
	// An Identity Store Group can be imported using the combination identity_store_id/group_id
	"aws_identitystore_group": TemplatedStringAsProviderDefinedIdentifier("{{ .parameters.identity_store_id }}/{{ .external_name }}"),
	// aws_identitystore_group_membership can be imported using the identity_store_id/membership_id
	"aws_identitystore_group_membership": TemplatedStringAsProviderDefinedIdentifier("{{ .parameters.identity_store_id }}/{{ .external_name }}"),
	// An Identity Store User can be imported using the combination identity_store_id/user_id
	"aws_identitystore_user": TemplatedStringAsProviderDefinedIdentifier("{{ .parameters.identity_store_id }}/{{ .external_name }}"),

	// imagebuilder
	//
	// aws_imagebuilder_components resources can be imported by using the Amazon Resource Name (ARN)
	"aws_imagebuilder_component": config.IdentifierFromProvider,
	// aws_imagebuilder_container_recipe resources can be imported by using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/example/1.0.0
	"aws_imagebuilder_container_recipe": config.IdentifierFromProvider,
	// aws_imagebuilder_distribution_configurations resources can be imported by using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:distribution-configuration/example
	"aws_imagebuilder_distribution_configuration": config.IdentifierFromProvider,
	// aws_imagebuilder_image resources can be imported using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:image/example/1.0.0/1
	// TODO: Normalize external_name while testing
	"aws_imagebuilder_image": config.IdentifierFromProvider,
	// aws_imagebuilder_image_pipeline resources can be imported using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/example
	"aws_imagebuilder_image_pipeline": config.IdentifierFromProvider,
	// aws_imagebuilder_image_recipe resources can be imported by using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:image-recipe/example/1.0.0
	"aws_imagebuilder_image_recipe": config.IdentifierFromProvider,
	// aws_imagebuilder_infrastructure_configuration can be imported using the Amazon Resource Name (ARN)
	// Example: arn:aws:imagebuilder:us-east-1:123456789012:infrastructure-configuration/example
	"aws_imagebuilder_infrastructure_configuration": config.IdentifierFromProvider,

	// mwaa
	// inspector
	//
	// mwaa_environment can be imported using the name
	"aws_mwaa_environment": config.NameAsIdentifier,

	//
	// Inspector Assessment Targets can be imported via their Amazon Resource Name (ARN)
	// Example: arn:aws:inspector:us-east-1:123456789012:target/0-xxxxxxx
	"aws_inspector_assessment_target": config.IdentifierFromProvider,
	// aws_inspector_assessment_template can be imported by using the template assessment ARN
	// Example: arn:aws:inspector:us-west-2:123456789012:target/0-9IaAzhGR/template/0-WEcjR8CH
	"aws_inspector_assessment_template": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_inspector_resource_group": config.IdentifierFromProvider,

	// inspector2
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	// TODO: Due to testing limitations, not sure if we will be able to test this resource. Do not spend a lot of time for test it.
	"aws_inspector2_enabler": config.IdentifierFromProvider,

	// iot
	//
	// IOT Authorizers can be imported using the name
	"aws_iot_authorizer": config.NameAsIdentifier,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_certificate": config.IdentifierFromProvider,
	// import IoT Domain Configuration using the name.
	"aws_iot_domain_configuration": config.NameAsIdentifier,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_indexing_configuration": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_logging_options": config.IdentifierFromProvider,
	// IoT policies can be imported using the name
	"aws_iot_policy": config.NameAsIdentifier,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_policy_attachment": config.IdentifierFromProvider,
	// IoT fleet provisioning templates can be imported using the name
	"aws_iot_provisioning_template": config.NameAsIdentifier,
	// IOT Role Alias can be imported via the alias
	"aws_iot_role_alias": config.IdentifierFromProvider,
	// IOT Things can be imported using the name
	"aws_iot_thing": config.NameAsIdentifier,
	// IoT Things Groups can be imported using the name
	"aws_iot_thing_group": config.NameAsIdentifier,
	// IoT Thing Group Membership can be imported using the thing group name and thing name
	// thing_group_name/thing_name
	"aws_iot_thing_group_membership": FormattedIdentifierFromProvider("/", "thing_group_name", "thing_name"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_iot_thing_principal_attachment": config.IdentifierFromProvider,
	// IOT Thing Types can be imported using the name
	// NameAsIdentifier would be better, but that would be a breaking api change to remove spec.forProvider.name
	"aws_iot_thing_type": FormattedIdentifierFromProvider("", "name"),
	// IoT Topic Rules can be imported using the name
	"aws_iot_topic_rule": config.NameAsIdentifier,
	// IoT topic rule destinations can be imported using the arn
	// arn:aws:iot:us-west-2:123456789012:ruledestination/vpc/2ce781c8-68a6-4c52-9c62-63fe489ecc60
	"aws_iot_topic_rule_destination": TemplatedStringAsProviderDefinedIdentifier(fullARNTemplate("iot", "ruledestination/vpc/{{ .external_name }}")),

	// ivs
	//
	// IVS (Interactive Video) Channel can be imported using the ARN
	// Example: arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5
	"aws_ivs_channel": config.IdentifierFromProvider,
	// IVS (Interactive Video) Recording Configuration can be imported using the ARN
	// Example: arn:aws:ivs:us-west-2:326937407773:recording-configuration/KAk1sHBl2L47
	"aws_ivs_recording_configuration": config.IdentifierFromProvider,

	// kafka
	//
	// MSK configurations can be imported using the configuration ARN that has
	// a random substring in the end.
	"aws_msk_configuration": config.IdentifierFromProvider,
	// MSK clusters can be imported using the cluster arn that has a random substring
	// in the end.
	"aws_msk_cluster": config.IdentifierFromProvider,
	// MSK relicators can be imported using the relicator arn
	// Example: arn:aws:kafka:us-west-2:123456789012:configuration/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
	"aws_msk_replicator": config.IdentifierFromProvider,
	// The terraform implementation of MSK SCRAM secret associations assume
	// that there is a single aws_msk_scram_secret_association per msk
	// cluster, so the best identifier is the cluster ARN.
	"aws_msk_scram_secret_association": config.IdentifierFromProvider,
	// MSK serverless clusters can be imported using the cluster arn
	// Example: arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
	"aws_msk_serverless_cluster": config.IdentifierFromProvider,
	// Managed Streaming for Kafka Cluster Policy resource can be imported using the cluster_arn
	"aws_msk_cluster_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.cluster_arn }}"),

	// kafkaconnect
	//
	// MSK Connect Connector can be imported using the connector's arn
	// Example: arn:aws:kafkaconnect:eu-central-1:123456789012:connector/example/264edee4-17a3-412e-bd76-6681cfc93805-3
	"aws_mskconnect_connector": TemplatedStringAsProviderDefinedIdentifier(fullARNTemplate("kafkaconnect", "connector/{{ .parameters.name }}/{{ .external_name }}")),
	// MSK Connect Custom Plugin can be imported using the plugin's arn
	// Example: arn:aws:kafkaconnect:eu-central-1:123456789012:custom-plugin/debezium-example/abcdefgh-1234-5678-9abc-defghijklmno-4
	"aws_mskconnect_custom_plugin": TemplatedStringAsProviderDefinedIdentifier(fullARNTemplate("kafkaconnect", "custom-plugin/{{ .parameters.name }}/{{ .external_name }}")),
	// MSK Connect Worker Configuration can be imported using the worker configuration's arn
	// Example: arn:aws:kafkaconnect:eu-central-1:123456789012:worker-configuration/example/8848493b-7fcc-478c-a646-4a52634e3378-4
	"aws_mskconnect_worker_configuration": TemplatedStringAsProviderDefinedIdentifier(fullARNTemplate("kafkaconnect", "worker-configuration/{{ .parameters.name }}/{{ .external_name }}")),

	// kendra
	//
	// Kendra Data Source can be imported using the unique identifiers of the data_source and index separated by a slash (/)
	"aws_kendra_data_source": config.IdentifierFromProvider,
	// Kendra Experience can be imported using the unique identifiers of the experience and index separated by a slash (/)
	"aws_kendra_experience": config.IdentifierFromProvider,
	// Amazon Kendra Indexes can be imported using its id
	// Example: 12345678-1234-5678-9123-123456789123
	// TODO: It seems that ID is autogenerated from provider.
	"aws_kendra_index": config.IdentifierFromProvider,
	// aws_kendra_query_suggestions_block_list can be imported using the unique identifiers of the block list and index separated by a slash (/)
	"aws_kendra_query_suggestions_block_list": config.IdentifierFromProvider,
	// aws_kendra_thesaurus can be imported using the unique identifiers of the thesaurus and index separated by a slash (/)
	"aws_kendra_thesaurus": config.IdentifierFromProvider,

	// keyspaces
	//
	// Use the name to import a keyspace
	"aws_keyspaces_keyspace": config.NameAsIdentifier,
	// Use the keyspace_name and table_name separated by / to import a table
	// my_keyspace/my_table
	"aws_keyspaces_table": FormattedIdentifierFromProvider("/", "keyspace_name", "table_name"),

	// kinesis
	//
	// Even though the documentation says the ID is name, it uses ARN..
	"aws_kinesis_stream": config.TemplatedStringAsIdentifier("name", fullARNTemplate("kinesis", "stream/{{ .external_name }}")),
	// Kinesis Stream Consumers can be imported using the Amazon Resource Name (ARN)
	// that has a random substring.
	"aws_kinesis_stream_consumer": config.IdentifierFromProvider,

	// kinesisanalytics
	//
	"aws_kinesis_analytics_application": config.TemplatedStringAsIdentifier("name", fullARNTemplate("kinesisanalytics", "application/{{ .external_name }}")),

	// kinesisanalyticsv2
	//
	"aws_kinesisanalyticsv2_application": config.TemplatedStringAsIdentifier("name", fullARNTemplate("kinesisanalytics", "application/{{ .external_name }}")),
	// aws_kinesisanalyticsv2_application can be imported by using application_name together with snapshot_name
	// e.g. example-application/example-snapshot
	"aws_kinesisanalyticsv2_application_snapshot": FormattedIdentifierUserDefinedNameLast("snapshot_name", "/", "application_name"),

	// kinesisvideo
	//
	// Kinesis Streams can be imported using the arn that has a random substring
	// in the end.
	// arn:aws:kinesisvideo:us-west-2:123456789012:stream/terraform-kinesis-test/1554978910975
	"aws_kinesis_video_stream": config.IdentifierFromProvider,

	// kms
	//
	// KMS aliases are imported using "alias/" + name
	"aws_kms_alias": kmsAlias(),
	// No import
	"aws_kms_ciphertext": config.IdentifierFromProvider,
	// KMS External Keys can be imported using the id
	"aws_kms_external_key": config.IdentifierFromProvider,
	// KMS Grants can be imported using the Key ID and Grant ID separated by a colon (:)
	"aws_kms_grant": config.IdentifierFromProvider,
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	"aws_kms_key": config.IdentifierFromProvider,
	// KMS multi-Region replica keys can be imported using the id
	"aws_kms_replica_external_key": config.IdentifierFromProvider,
	// KMS multi-Region replica keys can be imported using the id
	"aws_kms_replica_key": config.IdentifierFromProvider,

	// lakeformation
	//
	// No import
	"aws_lakeformation_data_lake_settings": config.IdentifierFromProvider,
	// No import
	"aws_lakeformation_permissions": config.IdentifierFromProvider,
	// No import
	"aws_lakeformation_resource": config.IdentifierFromProvider,

	// lambda
	//
	// Lambda Function Aliases are identified by their ARN, like arn:aws:lambda:eu-west-1:123456789012:function:lambda-function:alias
	"aws_lambda_alias": config.TemplatedStringAsIdentifier("name", fullARNTemplate("lambda", "function:{{ .parameters.function_name }}:{{ .external_name }}")),
	// Code Signing Configs can be imported using their ARN that has a random
	// substring in the end.
	// arn:aws:lambda:us-west-2:123456789012:code-signing-config:csc-0f6c334abcdea4d8b
	"aws_lambda_code_signing_config": config.IdentifierFromProvider,
	// Lambda event source mappings can be imported using the UUID (event source mapping identifier)
	"aws_lambda_event_source_mapping": config.IdentifierFromProvider,
	// Lambda Functions can be imported using the function_name
	"aws_lambda_function": config.ParameterAsIdentifier("function_name"),
	// Lambda Function Event Invoke Configs can be imported using the
	// fully qualified Function name or Amazon Resource Name (ARN) of the function.
	"aws_lambda_function_event_invoke_config": config.IdentifierFromProvider,
	// Lambda function URLs can be imported using the function_name or function_name/qualifier
	"aws_lambda_function_url": lambdaFunctionURL(),
	// No import"
	"aws_lambda_invocation": config.IdentifierFromProvider,
	// Lambda Layers can be imported using arn that has an assigned version in the
	// end
	"aws_lambda_layer_version": config.IdentifierFromProvider,
	// Lambda Layer Permissions can be imported using layer_name and version_number, separated by a comma (,)
	"aws_lambda_layer_version_permission": config.IdentifierFromProvider,
	// Lambda permission statements can be imported using function_name/statement_id, with an optional qualifier
	"aws_lambda_permission": config.IdentifierFromProvider,
	// Lambda Provisioned Concurrency Configs can be imported using the function_name and qualifier separated by a colon (:)
	"aws_lambda_provisioned_concurrency_config": config.IdentifierFromProvider,

	// lexmodels
	//
	// Bots can be imported using their name.
	"aws_lex_bot": config.NameAsIdentifier,
	// Bot aliases can be imported using an ID with the format bot_name:bot_alias_name
	"aws_lex_bot_alias": FormattedIdentifierUserDefinedNameLast("name", ":", "bot_name"),
	// Intents can be imported using their name.
	"aws_lex_intent": config.NameAsIdentifier,
	// Slot types can be imported using their name.
	"aws_lex_slot_type": config.NameAsIdentifier,

	// licensemanager
	//
	// License configurations can be imported in the form resource_arn,license_configuration_arn
	"aws_licensemanager_association": FormattedIdentifierFromProvider(",", "resource_arn", "license_configuration_arn"),
	// License configurations can be imported using the id
	"aws_licensemanager_license_configuration": config.IdentifierFromProvider,

	// lightsail
	//
	// aws_lightsail_bucket can be imported by using the name attribute
	"aws_lightsail_bucket": config.NameAsIdentifier,
	// aws_lightsail_certificate can be imported using the certificate name
	// TODO: Potential bug in documentation. If configuration doesn't work - change to IdentifierFromProvider
	"aws_lightsail_certificate": config.NameAsIdentifier,
	// Lightsail Container Service can be imported using the name
	"aws_lightsail_container_service": config.NameAsIdentifier,
	// aws_lightsail_disk can be imported by using the name attribute
	"aws_lightsail_disk": config.NameAsIdentifier,
	// aws_lightsail_disk can be imported by using the id attribute
	"aws_lightsail_disk_attachment": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_domain": config.IdentifierFromProvider,
	// aws_lightsail_domain_entry can be imported by using the id attribute
	// ID: name_domain_name_type_target
	"aws_lightsail_domain_entry": config.TemplatedStringAsIdentifier("name", "{{ .external_name }}_{{ .parameters.domain_name }}_{{ .parameters.type }}_{{ .parameeters.target }}"),
	// Lightsail Instances can be imported using their name
	"aws_lightsail_instance": config.NameAsIdentifier,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_instance_public_ports": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_key_pair": config.IdentifierFromProvider,
	// aws_lightsail_lb can be imported by using the name attribute
	"aws_lightsail_lb": config.NameAsIdentifier,
	// aws_lightsail_lb_attachment can be imported by using the name attribute
	// ID: lb_name,instance_name
	"aws_lightsail_lb_attachment": config.IdentifierFromProvider,
	// aws_lightsail_lb_certificate can be imported by using the id attribute
	// ID: lb_name,name
	"aws_lightsail_lb_certificate": config.TemplatedStringAsIdentifier("name", "{{ .parameters.lb_name }},{{ .external_name }}"),
	// aws_lightsail_lb_stickiness_policy can be imported by using the lb_name attribute
	"aws_lightsail_lb_stickiness_policy": config.ParameterAsIdentifier("lb_name"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_static_ip": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_lightsail_static_ip_attachment": config.IdentifierFromProvider,

	// location
	//
	// Location Geofence Collection can be imported using the collection_name
	"aws_location_geofence_collection": config.ParameterAsIdentifier("collection_name"),
	// aws_location_place_index resources can be imported using the place index name
	"aws_location_place_index": config.ParameterAsIdentifier("index_name"),
	// aws_location_route_calculator can be imported using the route calculator name
	"aws_location_route_calculator": config.ParameterAsIdentifier("calculator_name"),
	// aws_location_tracker resources can be imported using the tracker name
	"aws_location_tracker": config.ParameterAsIdentifier("tracker_name"),
	// Location Tracker Association can be imported using the tracker_name|consumer_arn
	"aws_location_tracker_association": config.IdentifierFromProvider,

	// macie2
	//
	// aws_macie2_account can be imported using the id
	"aws_macie2_account": config.IdentifierFromProvider,
	// aws_macie2_classification_job can be imported using the id
	"aws_macie2_classification_job": config.IdentifierFromProvider,
	// aws_macie2_custom_data_identifier can be imported using the id
	"aws_macie2_custom_data_identifier": config.IdentifierFromProvider,
	// aws_macie2_findings_filter can be imported using the id
	"aws_macie2_findings_filter": config.IdentifierFromProvider,
	// aws_macie2_invitation_accepter can be imported using the admin account ID
	"aws_macie2_invitation_accepter": FormattedIdentifierFromProvider("", "administrator_account_id"),
	// aws_macie2_member can be imported using the account ID of the member account
	"aws_macie2_member": FormattedIdentifierFromProvider("", "account_id"),

	// mediaconvert
	//
	// Media Convert Queue can be imported via the queue name
	"aws_media_convert_queue": config.NameAsIdentifier,

	// medialive
	//
	// MediaLive Channel can be imported using the channel_id
	"aws_medialive_channel": config.IdentifierFromProvider,
	// MediaLive Input can be imported using the id
	"aws_medialive_input": config.IdentifierFromProvider,
	// MediaLive InputSecurityGroup can be imported using the id
	"aws_medialive_input_security_group": config.IdentifierFromProvider,
	// MediaLive Multiplex can be imported using the id
	"aws_medialive_multiplex": config.IdentifierFromProvider,

	// mediapackage
	//
	// Media Package Channels can be imported via the channel ID
	"aws_media_package_channel": config.IdentifierFromProvider,

	// mediastore
	//
	// MediaStore Container can be imported using the MediaStore Container Name
	"aws_media_store_container": config.NameAsIdentifier,
	// MediaStore Container Policy can be imported using the MediaStore Container Name
	"aws_media_store_container_policy": FormattedIdentifierFromProvider("", "container_name"),

	// memorydb
	//
	// Use the name to import an ACL
	"aws_memorydb_acl": config.NameAsIdentifier,
	// Use the name to import a cluster
	"aws_memorydb_cluster": config.NameAsIdentifier,
	// Use the name to import a parameter group
	"aws_memorydb_parameter_group": config.NameAsIdentifier,
	// Use the name to import a snapshot
	"aws_memorydb_snapshot": config.NameAsIdentifier,
	// Use the name to import a subnet group
	"aws_memorydb_subnet_group": config.NameAsIdentifier,
	// Use the user_name to import a user
	"aws_memorydb_user": config.ParameterAsIdentifier("user_name"),

	// mq
	//
	// a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
	"aws_mq_broker": config.IdentifierFromProvider,
	// c-0187d1eb-88c8-475a-9b79-16ef5a10c94f
	"aws_mq_configuration": config.IdentifierFromProvider,

	// neptune
	//
	//
	"aws_neptune_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// my_cluster:my_cluster_endpoint
	"aws_neptune_cluster_endpoint": FormattedIdentifierUserDefinedNameLast("cluster_endpoint_identifier", ":", "cluster_identifier"),
	//
	"aws_neptune_cluster_instance": config.ParameterAsIdentifier("identifier"),
	//
	"aws_neptune_cluster_parameter_group": config.NameAsIdentifier,
	//
	"aws_neptune_cluster_snapshot": config.ParameterAsIdentifier("db_cluster_snapshot_identifier"),
	//
	"aws_neptune_event_subscription": config.NameAsIdentifier,
	// aws_neptune_global_cluster can be imported by using the Global Cluster identifier
	"aws_neptune_global_cluster": config.ParameterAsIdentifier("global_cluster_identifier"),
	//
	"aws_neptune_parameter_group": config.NameAsIdentifier,
	//
	"aws_neptune_subnet_group": config.NameAsIdentifier,

	// networkfirewall
	//
	// Network Firewall Firewalls can be imported using their ARN
	// Example: arn:aws:network-firewall:us-west-1:123456789012:firewall/example
	// "aws_networkfirewall_firewall": config.TemplatedStringAsIdentifier("name", "arn:aws:network-firewall:{{ .setup.configuration.region }}:{{ .setup.configuration.account_id }}:firewall/{{ .external_name }}"),
	"aws_networkfirewall_firewall": config.IdentifierFromProvider,
	// Network Firewall Policies can be imported using their ARN
	// Example: arn:aws:network-firewall:us-west-1:123456789012:firewall-policy/example
	"aws_networkfirewall_firewall_policy": config.TemplatedStringAsIdentifier("name", fullARNTemplate("network-firewall", "firewall-policy/{{ .external_name }}")),
	// Network Firewall Logging Configurations can be imported using the firewall_arn
	// Example: arn:aws:network-firewall:us-west-1:123456789012:firewall/example
	"aws_networkfirewall_logging_configuration": config.IdentifierFromProvider,
	// Network Firewall Rule Groups can be imported using their ARN
	// Example: arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
	"aws_networkfirewall_rule_group": config.TemplatedStringAsIdentifier("", fullARNTemplate("network-firewall", "{{ .parameters.type | ToLower }}-rulegroup/{{ .external_name }}")),

	// networkmanager
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_networkmanager_attachment_accepter": config.IdentifierFromProvider,
	// aws_networkmanager_connect_attachment can be imported using the attachment ID
	"aws_networkmanager_connect_attachment": config.IdentifierFromProvider,
	// aws_networkmanager_connection can be imported using the connection ARN
	// Example: arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/connection-07f6fd08867abc123
	"aws_networkmanager_connection": config.IdentifierFromProvider,
	// aws_networkmanager_core_network can be imported using the core network ID
	"aws_networkmanager_core_network": config.IdentifierFromProvider,
	// aws_networkmanager_customer_gateway_association can be imported using the global network ID and customer gateway ARN
	// Example: global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:customer-gateway/cgw-123abc05e04123abc
	"aws_networkmanager_customer_gateway_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.global_network_id }},arn:{{ .setup.client_metadata.partition }}:ec2:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:customer-gateway/{{ .parameters.customer_gateway_arn }}"),
	// aws_networkmanager_device can be imported using the device ARN
	// Example: arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/device-07f6fd08867abc123
	"aws_networkmanager_device": config.IdentifierFromProvider,
	// aws_networkmanager_global_network can be imported using the global network ID
	"aws_networkmanager_global_network": config.IdentifierFromProvider,
	// aws_networkmanager_link can be imported using the link ARN
	// Example: arn:aws:networkmanager::123456789012:link/global-network-0d47f6t230mz46dy4/link-444555aaabbb11223
	"aws_networkmanager_link": config.IdentifierFromProvider,
	// aws_networkmanager_link_association can be imported using the global network ID, link ID and device ID
	// Example: global-network-0d47f6t230mz46dy4,link-444555aaabbb11223,device-07f6fd08867abc123
	"aws_networkmanager_link_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.global_network_id }},{{ .parameters.link_id }},{{ .parameters.device_id }}"),
	// aws_networkmanager_site can be imported using the site ARN
	// Example: arn:aws:networkmanager::123456789012:site/global-network-0d47f6t230mz46dy4/site-444555aaabbb11223
	"aws_networkmanager_site": config.IdentifierFromProvider,
	// aws_networkmanager_transit_gateway_connect_peer_association can be imported using the global network ID and customer gateway ARN
	// Example: global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:transit-gateway-connect-peer/tgw-connect-peer-12345678
	"aws_networkmanager_transit_gateway_connect_peer_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.global_network_id }},arn:{{ .setup.client_metadata.partition }}:ec2:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:transit-gateway-connect-peer/{{ .parameters.transit_gateway_connect_peer_arn }}"),
	// aws_networkmanager_transit_gateway_registration can be imported using the global network ID and transit gateway ARN
	// Example: global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:transit-gateway/tgw-123abc05e04123abc
	"aws_networkmanager_transit_gateway_registration": config.TemplatedStringAsIdentifier("", "{{ .parameters.global_network_id }},{{ .parameters.transit_gateway_arn }}"),
	// aws_networkmanager_vpc_attachment can be imported using the attachment ID
	"aws_networkmanager_vpc_attachment": config.IdentifierFromProvider,

	// opensearch
	//
	// NOTE(sergen): Parameter as identifier cannot be used, because terraform
	// overrides the id after terraform calls.
	// Please see the following issue in upjet: https://github.com/crossplane/upjet/issues/32
	// OpenSearch domains can be imported using the domain_name
	"aws_opensearch_domain": config.IdentifierFromProvider,
	// No imports
	"aws_opensearch_domain_policy": config.IdentifierFromProvider,
	// NOTE(sergen): Parameter as identifier cannot be used, because terraform
	// overrides the id after terraform calls.
	// Please see the following issue in upjet: https://github.com/crossplane/upjet/issues/32
	// OpenSearch domains can be imported using the domain_name
	"aws_opensearch_domain_saml_options": config.IdentifierFromProvider,

	// organizations
	//
	// imported by using the account id, which is provider-generated
	"aws_organizations_account": config.IdentifierFromProvider,
	// imported by using the account ID and its service principal:
	// 123456789012/config.amazonaws.com
	"aws_organizations_delegated_administrator": FormattedIdentifierFromProvider("/", "account_id", "service_principal"),
	//  imported by using the id, which is a Cloud provider-generated string:
	// o-1234567
	"aws_organizations_organization": config.IdentifierFromProvider,
	// imported by using the id, which is a Cloud provider-generated string:
	// ou-1234567
	"aws_organizations_organizational_unit": config.IdentifierFromProvider,
	// imported by using the policy ID,
	// which is a Cloud provider-generated string:
	// p-12345678
	"aws_organizations_policy": config.IdentifierFromProvider,
	// imported by using the target ID and policy ID
	// 123456789012:p-12345678
	"aws_organizations_policy_attachment": FormattedIdentifierFromProvider(":", "target_id", "policy_id"),

	// pinpoint
	//
	// Pinpoint App can be imported using the application-id
	"aws_pinpoint_app": config.IdentifierFromProvider,
	// Pinpoint SMS Channel can be imported using the application-id
	"aws_pinpoint_sms_channel": FormattedIdentifierFromProvider("", "application_id"),

	// pipes
	//
	// Pipes can be imported using the name
	"aws_pipes_pipe": config.NameAsIdentifier,

	// qldb
	//
	// QLDB Ledgers can be imported using the name
	"aws_qldb_ledger": config.NameAsIdentifier,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_qldb_stream": config.IdentifierFromProvider,

	// quicksight
	//
	// QuickSight Group can be imported using the aws account id, namespace and group name separated by /
	// 123456789123/default/tf-example
	"aws_quicksight_group": FormattedIdentifierFromProvider("/", "aws_account_id", "namespace", "group_name"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_quicksight_user": config.IdentifierFromProvider,

	// ram
	//
	// RAM Principal Associations can be imported using their Resource Share ARN and the principal separated by a comma:
	// arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,123456789012
	"aws_ram_principal_association": FormattedIdentifierFromProvider(",", "resource_share_arn", "principal"),
	// RAM Resource Associations can be imported using their Resource Share ARN and Resource ARN separated by a comma:
	// arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,arn:aws:ec2:eu-west-1:123456789012:subnet/subnet-12345678
	"aws_ram_resource_association": FormattedIdentifierFromProvider(",", "resource_share_arn", "resource_arn"),
	// Resource shares can be imported using the id
	"aws_ram_resource_share": config.IdentifierFromProvider,
	// Resource share accepters can be imported using the resource share ARN:
	// arn:aws:ram:us-east-1:123456789012:resource-share/c4b56393-e8d9-89d9-6dc9-883752de4767
	"aws_ram_resource_share_accepter": config.IdentifierFromProvider,

	// rds
	//
	// aws_db_cluster_snapshot can be imported by using the cluster snapshot identifier
	"aws_db_cluster_snapshot": config.IdentifierFromProvider,
	// DB Event Subscriptions can be imported using the name
	"aws_db_event_subscription": config.NameAsIdentifier,
	//
	"aws_db_instance": config.IdentifierFromProvider,
	// RDS instance automated backups replication can be imported using the arn
	"aws_db_instance_automated_backups_replication": config.IdentifierFromProvider,
	// aws_db_instance_role_association can be imported using the DB Instance Identifier and IAM Role ARN separated by a comma
	// $ terraform import aws_db_instance_role_association.example my-db-instance,arn:aws:iam::123456789012:role/my-role
	"aws_db_instance_role_association": config.IdentifierFromProvider,
	// DB Option groups can be imported using the name
	"aws_db_option_group": config.NameAsIdentifier,
	//
	"aws_db_parameter_group": config.NameAsIdentifier,
	// DB proxies can be imported using the name
	"aws_db_proxy": config.NameAsIdentifier,
	// DB proxy default target groups can be imported using the db_proxy_name
	"aws_db_proxy_default_target_group": config.IdentifierFromProvider,
	// DB proxy endpoints can be imported using the DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME
	"aws_db_proxy_endpoint": config.TemplatedStringAsIdentifier("db_proxy_endpoint_name", "{{ .external_name }}/{{ .parameters.db_proxy_name }}"),
	// RDS DB Proxy Targets can be imported using the db_proxy_name, target_group_name, target type (e.g., RDS_INSTANCE or TRACKED_CLUSTER), and resource identifier separated by forward slashes (/)
	"aws_db_proxy_target": config.IdentifierFromProvider,
	// NOTE(turkenf): The resource aws_db_security_group is deprecated,
	// Please see: https://github.com/upbound/provider-aws/issues/696
	// aws_db_snapshot can be imported by using the snapshot identifier
	"aws_db_snapshot": config.ParameterAsIdentifier("db_snapshot_identifier"),
	// aws_db_snapshot_copy can be imported by using the snapshot identifier
	"aws_db_snapshot_copy": config.IdentifierFromProvider,
	//
	"aws_db_subnet_group": config.NameAsIdentifier,
	//
	"aws_rds_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// RDS Aurora Cluster Database Activity Streams can be imported using the resource_arn
	"aws_rds_cluster_activity_stream": config.IdentifierFromProvider,
	// RDS Clusters Endpoint can be imported using the cluster_endpoint_identifier
	"aws_rds_cluster_endpoint": config.ParameterAsIdentifier("cluster_endpoint_identifier"),
	// RDS Cluster Instances can be imported using the identifier
	"aws_rds_cluster_instance": config.ParameterAsIdentifier("identifier"),
	// RDS Cluster Parameter Groups can be imported using the name
	"aws_rds_cluster_parameter_group": config.NameAsIdentifier,
	// aws_rds_cluster_role_association can be imported using the DB Cluster Identifier and IAM Role ARN separated by a comma (,)
	// $ terraform import aws_rds_cluster_role_association.example my-db-cluster,arn:aws:iam::123456789012:role/my-role
	"aws_rds_cluster_role_association": FormattedIdentifierFromProvider(",", "db_cluster_identifier", "role_arn"),
	// aws_rds_global_cluster can be imported by using the RDS Global Cluster identifie
	"aws_rds_global_cluster": config.ParameterAsIdentifier("global_cluster_identifier"),

	// redshift
	//
	// Redshift Authentication Profiles support import by authentication_profile_name
	"aws_redshift_authentication_profile": config.ParameterAsIdentifier("authentication_profile_name"),
	// Redshift Clusters can be imported using the cluster_identifier
	"aws_redshift_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// Redshift endpoint access can be imported using the endpoint_name
	"aws_redshift_endpoint_access": config.ParameterAsIdentifier("endpoint_name"),
	// Redshift Event Subscriptions can be imported using the name
	"aws_redshift_event_subscription": config.NameAsIdentifier,
	// Redshift Hsm Client Certificates support import by hsm_client_certificate_identifier
	"aws_redshift_hsm_client_certificate": config.ParameterAsIdentifier("hsm_client_certificate_identifier"),
	// Redshift Hsm Client Certificates support import by hsm_configuration_identifier
	"aws_redshift_hsm_configuration": config.ParameterAsIdentifier("hsm_configuration_identifier"),
	// Redshift Parameter Groups can be imported using the name
	"aws_redshift_parameter_group": config.IdentifierFromProvider,
	// Redshift Scheduled Action can be imported using the name
	"aws_redshift_scheduled_action": config.NameAsIdentifier,
	// Redshift Snapshot Copy Grants support import by name
	"aws_redshift_snapshot_copy_grant": config.IdentifierFromProvider,
	// Redshift Snapshot Schedule can be imported using the identifier
	"aws_redshift_snapshot_schedule": config.ParameterAsIdentifier("identifier"),
	// Redshift Snapshot Schedule Association can be imported using the <cluster-identifier>/<schedule-identifier>
	"aws_redshift_snapshot_schedule_association": config.IdentifierFromProvider,
	// Redshift subnet groups can be imported using the name
	"aws_redshift_subnet_group": config.NameAsIdentifier,
	// Redshift usage limits can be imported using the id
	"aws_redshift_usage_limit": config.IdentifierFromProvider,

	// redshiftserverless
	//
	// Redshift Serverless Endpoint Access can be imported using the endpoint_name
	"aws_redshiftserverless_endpoint_access": config.ParameterAsIdentifier("endpoint_name"),
	// Redshift Serverless Namespaces can be imported using the namespace_name
	"aws_redshiftserverless_namespace": config.ParameterAsIdentifier("namespace_name"),
	// Redshift Serverless Resource Policies can be imported using the resource_arn
	"aws_redshiftserverless_resource_policy": config.IdentifierFromProvider,
	// Redshift Serverless Snapshots can be imported using the snapshot_name
	"aws_redshiftserverless_snapshot": config.ParameterAsIdentifier("snapshot_name"),
	// Redshift Serverless Usage Limits can be imported using the id
	"aws_redshiftserverless_usage_limit": config.IdentifierFromProvider,
	// Redshift Serverless Workgroups can be imported using the workgroup_name
	"aws_redshiftserverless_workgroup": config.ParameterAsIdentifier("workgroup_name"),

	// resource groups
	//
	// Resource groups can be imported using the name
	"aws_resourcegroups_group": config.NameAsIdentifier,

	// rolesanywhere
	//
	// aws_rolesanywhere_profile can be imported using its id
	"aws_rolesanywhere_profile": config.IdentifierFromProvider,

	// route53
	//
	// N1PA6795SAMPLE
	"aws_route53_delegation_set": config.IdentifierFromProvider,
	// abcdef11-2222-3333-4444-555555fedcba
	"aws_route53_health_check": config.IdentifierFromProvider,
	// Z1D633PJN98FT9
	"aws_route53_hosted_zone_dnssec": config.IdentifierFromProvider,
	// Imported using ID of the record, which is the zone identifier, record
	// name, and record type, separated by underscores (_)
	// Z4KAPRWWNC7JR_dev.example.com_NS
	"aws_route53_record": config.IdentifierFromProvider,
	// Imported using the id and version, e.g.,
	// 01a52019-d16f-422a-ae72-c306d2b6df7e/1
	"aws_route53_traffic_policy": config.IdentifierFromProvider,
	// df579d9a-6396-410e-ac22-e7ad60cf9e7e
	"aws_route53_traffic_policy_instance": config.IdentifierFromProvider,
	// Z123456ABCDEFG:vpc-12345678
	"aws_route53_vpc_association_authorization": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	// Z1D633PJN98FT9
	"aws_route53_zone": config.IdentifierFromProvider,
	// Z123456ABCDEFG:vpc-12345678
	// "aws_route53_zone_association": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	"aws_route53_zone_association": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	// Route53 query logging configurations can be imported using their ID
	"aws_route53_query_log": config.IdentifierFromProvider,

	// route53recoverycontrolconfig
	//
	// Imported using ARN that has a random substring:
	// arn:aws:route53-recovery-control::313517334327:cluster/f9ae13be-a11e-4ec7-8522-94a70468e6ea
	"aws_route53recoverycontrolconfig_cluster": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	// arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8
	"aws_route53recoverycontrolconfig_control_panel": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	// arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8/routingcontrol/d5d90e587870494b
	"aws_route53recoverycontrolconfig_routing_control": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	// arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8/safetyrule/3bacc77003364c0f
	"aws_route53recoverycontrolconfig_safety_rule": config.IdentifierFromProvider,

	// route53recoveryreadiness
	//
	// Route53 Recovery Readiness cells can be imported via the cell name
	"aws_route53recoveryreadiness_cell": config.ParameterAsIdentifier("cell_name"),
	// Route53 Recovery Readiness readiness checks can be imported via the readiness check name
	"aws_route53recoveryreadiness_readiness_check": config.ParameterAsIdentifier("readiness_check_name"),
	// Route53 Recovery Readiness recovery groups can be imported via the recovery group name
	"aws_route53recoveryreadiness_recovery_group": config.ParameterAsIdentifier("recovery_group_name"),
	// Route53 Recovery Readiness resource set name can be imported via the resource set name
	"aws_route53recoveryreadiness_resource_set": config.ParameterAsIdentifier("resource_set_name"),

	// route53resolver
	//
	// Route 53 Resolver configs can be imported using the Route 53 Resolver config ID
	"aws_route53_resolver_config": config.IdentifierFromProvider,
	// rslvr-in-abcdef01234567890
	"aws_route53_resolver_endpoint": config.IdentifierFromProvider,
	// rslvr-rr-0123456789abcdef0
	"aws_route53_resolver_rule": config.IdentifierFromProvider,
	// rslvr-rrassoc-97242eaf88example
	"aws_route53_resolver_rule_association": config.IdentifierFromProvider,

	// rum
	//
	// Cloudwatch RUM App Monitor can be imported using the name
	"aws_rum_app_monitor": config.NameAsIdentifier,
	// Cloudwatch RUM Metrics Destination can be imported using the id
	"aws_rum_metrics_destination": config.IdentifierFromProvider,

	// s3
	//
	// S3 bucket can be imported using the bucket
	"aws_s3_bucket": config.ParameterAsIdentifier("bucket"),
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
	// The S3 bucket logging resource should be imported using the bucket
	"aws_s3_bucket_logging": config.IdentifierFromProvider,
	// S3 bucket metric configurations can be imported using bucket:metric
	"aws_s3_bucket_metric": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// S3 bucket notification can be imported using the bucket
	"aws_s3_bucket_notification": config.IdentifierFromProvider,
	// Objects can be imported using the id. The id is the bucket name and the key together
	"aws_s3_bucket_object": config.IdentifierFromProvider,
	// the S3 bucket accelerate configuration resource should be imported using the bucket
	"aws_s3_bucket_object_lock_configuration": config.IdentifierFromProvider,
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
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_s3_object_copy": config.IdentifierFromProvider,

	// s3control
	//
	// S3 Storage Lens configurations can be imported using the account_id and config_id, separated by a colon (:)
	"aws_s3control_storage_lens_configuration": config.IdentifierFromProvider,
	// - For Access Points associated with an AWS Partition S3 Bucket, this resource
	// can be imported using the account_id and name separated by a colon (:)
	// - For Access Points associated with an S3 on Outposts Bucket, this resource
	// can be imported using the Amazon Resource Name (ARN)
	// TODO: There are two different import syntaxes for this resource. For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_s3_access_point": config.IdentifierFromProvider,
	// aws_s3_account_public_access_block can be imported by using the AWS account ID
	"aws_s3_account_public_access_block": config.IdentifierFromProvider,
	// Access Point policies can be imported using the access_point_arn
	// arn:aws:s3:us-west-2:123456789012:accesspoint/example
	"aws_s3control_access_point_policy": config.IdentifierFromProvider,
	// Multi-Region Access Points can be imported using the account_id and name of the Multi-Region Access Point separated by a colon (:)
	// Example: 123456789012:example
	"aws_s3control_multi_region_access_point": config.IdentifierFromProvider,
	// Multi-Region Access Point Policies can be imported using the account_id and name of the Multi-Region Access Point separated by a colon (:)
	// Example: 123456789012:example
	"aws_s3control_multi_region_access_point_policy": config.IdentifierFromProvider,
	// Object Lambda Access Points can be imported using the account_id and name, separated by a colon (:)
	// Example: 123456789012:example
	"aws_s3control_object_lambda_access_point": config.IdentifierFromProvider,
	// Object Lambda Access Point policies can be imported using the account_id and name, separated by a colon (:)
	// Example: 123456789012:example
	"aws_s3control_object_lambda_access_point_policy": config.IdentifierFromProvider,

	// sagemaker
	//
	// SageMaker Apps can be imported using the id
	"aws_sagemaker_app": config.IdentifierFromProvider,
	// SageMaker App Image Configs can be imported using the name
	"aws_sagemaker_app_image_config": config.ParameterAsIdentifier("app_image_config_name"),
	// SageMaker Code Repositories can be imported using the name
	"aws_sagemaker_code_repository": config.ParameterAsIdentifier("code_repository_name"),
	// SageMaker Devices can be imported using the device-fleet-name/device-name
	// my-fleet/my-device
	"aws_sagemaker_device": config.IdentifierFromProvider,
	// SageMaker Device Fleets can be imported using the name
	"aws_sagemaker_device_fleet": config.ParameterAsIdentifier("device_fleet_name"),
	// SageMaker Domains can be imported using the id
	"aws_sagemaker_domain": config.IdentifierFromProvider,
	// Endpoints can be imported using the name
	"aws_sagemaker_endpoint": config.NameAsIdentifier,
	// Endpoint configurations can be imported using the name
	"aws_sagemaker_endpoint_configuration": config.NameAsIdentifier,
	// Feature Groups can be imported using the name
	"aws_sagemaker_feature_group": config.ParameterAsIdentifier("feature_group_name"),
	// SageMaker Code Images can be imported using the name
	"aws_sagemaker_image": config.ParameterAsIdentifier("image_name"),
	// SageMaker Code Images can be imported using the name
	"aws_sagemaker_image_version": config.IdentifierFromProvider,
	// Sagemaker MLFlow tracking server can be imported using the name
	"aws_sagemaker_mlflow_tracking_server": config.ParameterAsIdentifier("tracking_server_name"),
	// Models can be imported using the name
	"aws_sagemaker_model": config.NameAsIdentifier,
	// SageMaker Model Package Groups can be imported using the name
	"aws_sagemaker_model_package_group": config.ParameterAsIdentifier("model_package_group_name"),
	// SageMaker Model Package Groups can be imported using the name
	"aws_sagemaker_model_package_group_policy": config.IdentifierFromProvider,
	// SageMaker Notebook Instances can be imported using the name
	"aws_sagemaker_notebook_instance": config.NameAsIdentifier,
	// Models can be imported using the name
	"aws_sagemaker_notebook_instance_lifecycle_configuration": config.NameAsIdentifier,
	// Models can be imported using the id
	"aws_sagemaker_servicecatalog_portfolio_status": config.IdentifierFromProvider,
	// SageMaker Spaces can be imported using the id
	// Example: arn:aws:sagemaker:us-west-2:123456789012:space/domain-id/space-name
	"aws_sagemaker_space": config.IdentifierFromProvider,
	// SageMaker Studio Lifecycle Configs can be imported using the studio_lifecycle_config_name
	"aws_sagemaker_studio_lifecycle_config": config.ParameterAsIdentifier("studio_lifecycle_config_name"),
	// SageMaker User Profiles can be imported using the arn
	"aws_sagemaker_user_profile": config.IdentifierFromProvider,
	// SageMaker Workforces can be imported using the workforce_name
	"aws_sagemaker_workforce": config.ParameterAsIdentifier("workforce_name"),
	// SageMaker Workteams can be imported using the workteam_name
	"aws_sagemaker_workteam": config.ParameterAsIdentifier("workteam_name"),

	// scheduler
	//
	// Schedules can be imported using the combination group_name/name
	"aws_scheduler_schedule": config.IdentifierFromProvider,
	// Schedule groups can be imported using the name
	"aws_scheduler_schedule_group": config.IdentifierFromProvider,

	// schemas
	//
	// EventBridge discoverers can be imported using the id
	"aws_schemas_discoverer": config.IdentifierFromProvider,
	// EventBridge schema registries can be imported using the name
	"aws_schemas_registry": config.NameAsIdentifier,
	// EventBridge schema can be imported using the name and registry_name
	"aws_schemas_schema": FormattedIdentifierFromProvider("/", "name", "registry_name"),

	// secretsmanager
	//
	// It be imported by using the secret Amazon Resource Name (ARN)
	// However, the real ID of the Secret has an Amazon-assigned random suffix,
	// i.e. if you name it with `example`, the real ID is
	// arn:aws:secretsmanager:us-west-1:609897127049:secret:example-VaznFM
	"aws_secretsmanager_secret": config.IdentifierFromProvider,
	// It uses its own secert_arn parameter.
	"aws_secretsmanager_secret_policy": config.IdentifierFromProvider,
	// It uses its own secret_id parameter.
	"aws_secretsmanager_secret_rotation": config.IdentifierFromProvider,
	// It uses ARN of secret and a randomly assigned ID.
	"aws_secretsmanager_secret_version": config.IdentifierFromProvider,

	// securityhub
	//
	// An existing Security Hub enabled account can be imported using the AWS account ID
	"aws_securityhub_account": config.IdentifierFromProvider,
	// imported using the action target ARN:
	// arn:aws:securityhub:eu-west-1:312940875350:action/custom/a
	// TODO: following configuration assumes the `a` in the above ARN
	// is the security hub custom action identifier
	"aws_securityhub_action_target": config.TemplatedStringAsIdentifier("identifier", fullARNTemplate("securityhub", "action/custom/{{ .external_name }}")),
	// imported using the arn that has a random substring:
	// arn:aws:securityhub:eu-west-1:123456789098:finding-aggregator/abcd1234-abcd-1234-1234-abcdef123456
	"aws_securityhub_finding_aggregator": config.IdentifierFromProvider,
	// imported using the ARN that has a random substring:
	// arn:aws:securityhub:us-west-2:1234567890:insight/1234567890/custom/91299ed7-abd0-4e44-a858-d0b15e37141a
	"aws_securityhub_insight": config.IdentifierFromProvider,
	// imported using the account ID
	"aws_securityhub_invite_accepter": config.IdentifierFromProvider,
	// imported using security hub member account ID
	"aws_securityhub_member": FormattedIdentifierFromProvider("", "account_id"),
	// imported in the form product_arn,arn:
	// arn:aws:securityhub:eu-west-1:733251395267:product/alertlogic/althreatmanagement,arn:aws:securityhub:eu-west-1:123456789012:product-subscription/alertlogic/althreatmanagement
	// looks like it's possible to derive the external-name from
	// the product_arn argument according to the above example
	// (by replacing product by product-subscription), which makes this
	// a special case of FormattedIdentifierFromProvider
	"aws_securityhub_product_subscription": func() config.ExternalName {
		e := config.IdentifierFromProvider
		e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			val, ok := parameters["product_arn"]
			if !ok {
				return "", errors.New("product_arn cannot be empty")
			}
			s, ok := val.(string)
			if !ok {
				return "", errors.New("product_arn needs to be a string")
			}
			return fmt.Sprintf("%s,%s", s, strings.Replace(s, ":product", ":product-subscription", 1)), nil
		}
		return e
	}(),
	// imported using the standards subscription ARN:
	// arn:aws:securityhub:eu-west-1:123456789012:subscription/pci-dss/v/3.2.1
	"aws_securityhub_standards_subscription": FormattedIdentifierFromProvider("", "standards_arn"),

	// serverlessrepo
	//
	// imported using the CloudFormation Stack name
	"aws_serverlessapplicationrepository_cloudformation_stack": config.IdentifierFromProvider,

	// servicecatalog
	//
	// imported using the budget name and resource ID:
	// budget-pjtvyakdlyo3m:prod-dnigbtea24ste
	"aws_servicecatalog_budget_resource_association": config.IdentifierFromProvider,
	// imported using the constraint ID, which has random parts
	// generated by the provider: cons-nmdkb6cgxfcrs
	"aws_servicecatalog_constraint": config.IdentifierFromProvider,
	// imported using the service catalog portfolio id,
	// which has provider-generated random parts:
	// port-12344321
	"aws_servicecatalog_portfolio": config.IdentifierFromProvider,
	// imported using the portfolio share ID: port-12344321:ACCOUNT:123456789012
	// TODO: looks like we can generate the above ID using:
	// portfolio_id:type:principal_id
	// but this has to be validated
	"aws_servicecatalog_portfolio_share": config.IdentifierFromProvider,
	// imported using the accept language, principal ARN, and portfolio ID, separated by a comma:
	// en,arn:aws:iam::123456789012:user/Eleanor,port-68656c6c6f
	// TODO: looks like we can generated the above id using:
	// accept_language,principal_arn,portfolio_id
	// , which lends itself to:
	// FormattedIdentifierFromProvider(",", "accept_language", "principal_arn", "portfolio_id")
	// However, accept_language is optional. We had better make it required as
	// the default is provided by Terraform (and we have no means to default
	// the generated CRD fields as of now)
	"aws_servicecatalog_principal_portfolio_association": config.IdentifierFromProvider,
	// imported using the product ID, which has provider-generated random parts:
	// prod-dnigbtea24ste
	"aws_servicecatalog_product": config.IdentifierFromProvider,
	// imported using the accept language, portfolio ID, and product ID:
	// en:port-68656c6c6f:prod-dnigbtea24ste
	// TODO: looks like we can generated the above id using:
	// accept_language,portfolio_id,product_id
	// , which lends itself to:
	// FormattedIdentifierFromProvider(",", "accept_language", "portfolio_id", "product_id")
	// However, accept_language is optional. We had better make it required as
	// the default is provided by Terraform (and we have no means to default
	// the generated CRD fields as of now)
	"aws_servicecatalog_product_portfolio_association": config.IdentifierFromProvider,
	// imported using the provisioning artifact ID and product ID separated by a colon:
	// pa-ij2b6lusy6dec:prod-el3an0rma3
	// we could make the product_id attribute the name identifier
	// and concatenate it with the provider-generated provisioning
	// artifact id, but product id does not does not look like to
	// be a good external-name for this resource as this is the
	// provisioning artifact resource.
	"aws_servicecatalog_provisioning_artifact": config.IdentifierFromProvider,
	// imported using the service action ID. which has provider-generated
	// random parts: act-f1w12eperfslh
	"aws_servicecatalog_service_action": config.IdentifierFromProvider,
	// imported using the tag option ID, which has provider-generated
	// random parts: tag-pjtvagohlyo3m
	"aws_servicecatalog_tag_option": config.IdentifierFromProvider,
	// imported using the tag option ID and resource ID:
	// tag-pjtvyakdlyo3m:prod-dnigbtea24ste
	"aws_servicecatalog_tag_option_resource_association": FormattedIdentifierFromProvider(":", "tag_option_id", "resource_id"),

	// servicediscovery
	//
	// Service Discovery HTTP Namespace can be imported using the namespace ID,
	"aws_service_discovery_http_namespace": config.IdentifierFromProvider,
	// Service Discovery Private DNS Namespace can be imported using the namespace ID and VPC ID: 0123456789:vpc-123345
	"aws_service_discovery_private_dns_namespace": config.IdentifierFromProvider,
	// Service Discovery Public DNS Namespace can be imported using the namespace ID
	"aws_service_discovery_public_dns_namespace": config.IdentifierFromProvider,
	// Service Discovery Service can be imported using the service ID
	"aws_service_discovery_service": config.IdentifierFromProvider,

	// servicequotas
	//
	// aws_servicequotas_service_quota can be imported by using the service code and quota code, separated by a front slash (/)
	// vpc/L-F678F1CE
	"aws_servicequotas_service_quota": FormattedIdentifierFromProvider("/", "service_code", "quota_code"),

	// ses
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_ses_active_receipt_rule_set": config.IdentifierFromProvider,
	// SES Configuration Sets can be imported using their name
	"aws_ses_configuration_set": config.NameAsIdentifier,
	// DKIM tokens can be imported using the domain attribute
	"aws_ses_domain_dkim": config.ParameterAsIdentifier("domain"),
	// SES domain identities can be imported using the domain name.
	"aws_ses_domain_identity": config.ParameterAsIdentifier("domain"),
	// MAIL FROM domain can be imported using the domain attribute
	"aws_ses_domain_mail_from": config.IdentifierFromProvider,
	// SES email identities can be imported using the email address.
	"aws_ses_email_identity": config.IdentifierFromProvider,
	// SES event destinations can be imported using configuration_set_name together with the event destination's name
	// Example: some-configuration-set-test/event-destination-sns
	"aws_ses_event_destination": config.NameAsIdentifier,
	// Identity Notification Topics can be imported using the ID of the record. The ID is made up as IDENTITY|TYPE where IDENTITY is the SES Identity and TYPE is the Notification Type.
	// Example: 'example.com|Bounce'
	"aws_ses_identity_notification_topic": config.IdentifierFromProvider,
	// SES Identity Policies can be imported using the identity and policy name, separated by a pipe character (|)
	// Example: 'example.com|example'
	"aws_ses_identity_policy": config.IdentifierFromProvider,
	// SES Receipt Filter can be imported using their name
	"aws_ses_receipt_filter": config.NameAsIdentifier,
	// SES receipt rules can be imported using the ruleset name and rule name separated by :
	// Example: my_rule_set:my_rule
	"aws_ses_receipt_rule": config.IdentifierFromProvider,
	// SES receipt rule sets can be imported using the rule set name
	"aws_ses_receipt_rule_set": config.IdentifierFromProvider,
	// SES templates can be imported using the template name
	"aws_ses_template": config.NameAsIdentifier,

	// sesv2
	//
	// SESv2 (Simple Email V2) Configuration Set can be imported using the configuration_set_name
	"aws_sesv2_configuration_set": config.ParameterAsIdentifier("configuration_set_name"),
	// SESv2 (Simple Email V2) Configuration Set Event Destination can be imported using the id (configuration_set_name|event_destination_name)
	"aws_sesv2_configuration_set_event_destination": config.IdentifierFromProvider,
	// SESv2 (Simple Email V2) Dedicated IP Pool can be imported using the pool_name
	"aws_sesv2_dedicated_ip_pool": config.ParameterAsIdentifier("pool_name"),
	// SESv2 (Simple Email V2) Email Identity can be imported using the email_identity
	"aws_sesv2_email_identity": config.ParameterAsIdentifier("email_identity"),
	// SESv2 (Simple Email V2) Email Identity Feedback Attributes can be imported using the email_identity
	"aws_sesv2_email_identity_feedback_attributes": config.ParameterAsIdentifier("email_identity"),
	// SESv2 (Simple Email V2) Email Identity Mail From Attributes can be imported using the email_identity
	"aws_sesv2_email_identity_mail_from_attributes": config.ParameterAsIdentifier("email_identity"),

	// sfn
	//
	"aws_sfn_activity": config.TemplatedStringAsIdentifier("name", fullARNTemplate("states", "activity/{{ .external_name }}")),
	//
	"aws_sfn_state_machine": config.TemplatedStringAsIdentifier("name", fullARNTemplate("states", "stateMachine:{{ .external_name }}")),

	// signer
	//
	// Signer signing jobs can be imported using the job_id
	"aws_signer_signing_job": config.IdentifierFromProvider,
	// Signer signing profiles can be imported using the name
	"aws_signer_signing_profile": config.NameAsIdentifier,
	// Signer signing profile permission statements can be imported using profile_name/statement_id
	// Example: prod_profile_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK/ProdAccountStartSigningJobStatementId
	"aws_signer_signing_profile_permission": config.TemplatedStringAsIdentifier("", "{{ .parameters.profile_name }}/{{ .parameters.statement_id }}"),

	// sns
	//
	// SNS platform applications can be imported using the ARN:
	// arn:aws:sns:us-west-2:0123456789012:app/GCM/gcm_application
	"aws_sns_platform_application": config.TemplatedStringAsIdentifier("name", fullARNTemplate("sns", "app/GCM/{{ .external_name }}")),
	// no import documentation is provided
	// TODO: we will need to check if normalization is possible
	"aws_sns_sms_preferences": config.IdentifierFromProvider,
	// SNS Topics can be imported using the topic arn
	"aws_sns_topic": config.TemplatedStringAsIdentifier("name", fullARNTemplate("sns", "{{ .external_name }}")),
	// SNS Topic Policy can be imported using the topic ARN:
	// arn:aws:sns:us-west-2:0123456789012:my-topic
	"aws_sns_topic_policy": FormattedIdentifierFromProvider("", "arn"),
	// SNS Topic Subscriptions can be imported using the subscription arn that
	// contains a random substring in the end.
	"aws_sns_topic_subscription": config.IdentifierFromProvider,

	// sqs
	//
	// SQS Queues can be imported using the queue url / id
	"aws_sqs_queue": config.IdentifierFromProvider,
	// SQS Queue Policies can be imported using the queue URL
	// e.g. https://queue.amazonaws.com/0123456789012/myqueue
	"aws_sqs_queue_policy": config.IdentifierFromProvider,
	// SQS Queue Redrive Allow Policies can be imported using the queue URL
	"aws_sqs_queue_redrive_allow_policy": config.IdentifierFromProvider,
	// SQS Queue Redrive Policies can be imported using the queue URL
	"aws_sqs_queue_redrive_policy": config.IdentifierFromProvider,

	// ssm
	//
	// AWS SSM Activation can be imported using the id
	"aws_ssm_activation": config.IdentifierFromProvider,
	// SSM associations can be imported using the association_id
	"aws_ssm_association": config.IdentifierFromProvider,
	// The Systems Manager Default Patch Baseline can be imported using the patch baseline ID, patch baseline ARN, or the operating system value
	"aws_ssm_default_patch_baseline": config.IdentifierFromProvider,
	// SSM Documents can be imported using the name
	"aws_ssm_document": config.NameAsIdentifier,
	// SSM Maintenance Windows can be imported using the maintenance window id
	"aws_ssm_maintenance_window": config.IdentifierFromProvider,
	// SSM Maintenance Window targets can be imported using WINDOW_ID/WINDOW_TARGET_ID
	"aws_ssm_maintenance_window_target": config.IdentifierFromProvider,
	// AWS Maintenance Window Task can be imported using the window_id and window_task_id separated by /
	"aws_ssm_maintenance_window_task": config.IdentifierFromProvider,
	// SSM Parameters can be imported using the parameter store name
	"aws_ssm_parameter": config.NameAsIdentifier,
	// SSM Patch Baselines can be imported by their baseline ID
	"aws_ssm_patch_baseline": config.IdentifierFromProvider,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_ssm_patch_group": config.IdentifierFromProvider,
	// SSM resource data sync can be imported using the name
	"aws_ssm_resource_data_sync": config.NameAsIdentifier,
	// AWS SSM Service Setting can be imported using the setting_id
	"aws_ssm_service_setting": config.IdentifierFromProvider,

	// ssoadmin
	//
	// SSO Account Assignments can be imported using the principal_id, principal_type, target_id, target_type, permission_set_arn, instance_arn separated by commas (,)
	// Example: f81d4fae-7dec-11d0-a765-00a0c91e6bf6,GROUP,1234567890,AWS_ACCOUNT,arn:aws:sso:::permissionSet/ssoins-0123456789abcdef/ps-0123456789abcdef,arn:aws:sso:::instance/ssoins-0123456789abcdef
	// This can't really be normalized.
	"aws_ssoadmin_account_assignment": config.TemplatedStringAsIdentifier("", "{{ .parameters.principal_id }},{{ .parameters.principal_type }},{{ .parameters.target_id }},{{ .parameters.target_type }},{{ .parameters.permission_set_arn }},{{ .parameters.instance_arn }}"),
	// SSO Managed Policy Attachments can be imported using the name, path, permission_set_arn, and instance_arn separated by a comma (,)
	// Example: TestPolicy,/,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	// This can't really be normalized.
	"aws_ssoadmin_customer_managed_policy_attachment": config.TemplatedStringAsIdentifier("", "{{  (index .parameters.customer_managed_policy_reference 0).name }},{{ (index .parameters.customer_managed_policy_reference 0).path }},{{ .parameters.permission_set_arn }},{{ .parameters.instance_arn }}"),
	// SSO Instance Access Control Attributes can be imported using the instance_arn
	"aws_ssoadmin_instance_access_control_attributes": config.TemplatedStringAsIdentifier("", "{{ .parameters.instance_arn }}"),
	// SSO Managed Policy Attachments can be imported using the managed_policy_arn, permission_set_arn, and instance_arn separated by a comma (,)
	// Example: arn:aws:iam::aws:policy/AlexaForBusinessDeviceSetup,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	// This can't really be normalized.
	"aws_ssoadmin_managed_policy_attachment": config.TemplatedStringAsIdentifier("", "{{ .parameters.managed_policy_arn }},{{ .parameters.permission_set_arn }},{{ .parameters.instance_arn}}"),
	// SSO Permission Sets can be imported using the arn and instance_arn separated by a comma (,)
	// Example: arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	// TODO: Normalize to the permission set id once breaking changes are acceptable or multiple versions are supported
	"aws_ssoadmin_permission_set": config.IdentifierFromProvider,
	// The best name is the permission set id
	// SSO Permission Set Inline Policies can be imported using the permission_set_arn and instance_arn separated by a comma (,)
	// Example: arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	// TODO: Normalize to the permission set id once breaking changes are acceptable or multiple versions are supported
	"aws_ssoadmin_permission_set_inline_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.permission_set_arn }},{{ .parameters.instance_arn }}"),
	// The best name is the permission set id
	// SSO Admin Permissions Boundary Attachments can be imported using the permission_set_arn and instance_arn, separated by a comma (,)
	// Example: arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
	"aws_ssoadmin_permissions_boundary_attachment": PermissionSetIdAsExternalName(),

	// swf
	//
	// SWF Domains can be imported using the name
	"aws_swf_domain": config.NameAsIdentifier,

	// timestreamwrite
	//
	// Timestream databases can be imported using the database_name
	"aws_timestreamwrite_database": config.ParameterAsIdentifier("database_name"),
	// Timestream tables can be imported using the table_name and database_name separate by a colon (:)
	// Example: ExampleTable:ExampleDatabase
	"aws_timestreamwrite_table": config.TemplatedStringAsIdentifier("", "{{ .parameters.table_name }}:{{ .parameters.database_name }}"),

	// transcribe
	//
	// Transcribe LanguageModel can be imported using the model_name
	"aws_transcribe_language_model": config.ParameterAsIdentifier("model_name"),
	// Transcribe Vocabulary can be imported using the vocabulary_name
	"aws_transcribe_vocabulary": config.ParameterAsIdentifier("vocabulary_name"),
	// Transcribe VocabularyFilter can be imported using the vocabulary_filter_name
	"aws_transcribe_vocabulary_filter": config.ParameterAsIdentifier("vocabulary_filter_name"),

	// transfer
	//
	// Transfer Connector can be imported using the connector_id.
	// Example: c-4221a88afd5f4362a
	"aws_transfer_connector": config.IdentifierFromProvider,
	// Transfer Servers can be imported using the id
	"aws_transfer_server": config.IdentifierFromProvider,
	// Transfer SSH Public Key can be imported using the server_id and user_name and ssh_public_key_id separated by /
	// Example: s-12345678/test-username/key-12345
	"aws_transfer_ssh_key": config.IdentifierFromProvider,
	// aws_transfer_tag can be imported by using the Transfer Family resource identifier and key, separated by a comma (,)
	// Example: arn:aws:transfer:us-east-1:123456789012:server/s-1234567890abcdef0,Name
	"aws_transfer_tag": config.IdentifierFromProvider,
	// Transfer Users can be imported using the server_id and user_name separated by /
	"aws_transfer_user": FormattedIdentifierUserDefinedNameLast("user_name", "/", "server_id"),
	// Transfer Workflows can be imported using the worflow_idgit
	"aws_transfer_workflow": config.IdentifierFromProvider,

	// vpc
	//
	// Note: This resource uses the ec2 go sdk group, but we released it in a package named vpc because we missed it
	// from a list of naming convention exceptions.
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"aws_vpc_network_performance_metric_subscription": config.IdentifierFromProvider,

	// vpclattice
	//
	// VPC Lattice Access Log Subscription can be imported using the id
	"aws_vpclattice_access_log_subscription": config.IdentifierFromProvider,
	// VPC Lattice Auth Policy can be imported using the id
	"aws_vpclattice_auth_policy": config.IdentifierFromProvider,
	// VPC Lattice Listener can be imported using the service_id/listener_id
	"aws_vpclattice_listener": config.IdentifierFromProvider,
	// VPC Lattice Listener Rule can be imported using the id
	"aws_vpclattice_listener_rule": config.IdentifierFromProvider,
	// VPC Lattice Resource Policy can be imported using the id
	"aws_vpclattice_resource_policy": config.IdentifierFromProvider,
	// VPC Lattice Service can be imported using the id
	"aws_vpclattice_service": config.IdentifierFromProvider,
	// VPC Lattice Service Network can be imported using the id
	"aws_vpclattice_service_network": config.IdentifierFromProvider,
	// VPC Lattice Service Network Service Association can be imported using the id
	"aws_vpclattice_service_network_service_association": config.IdentifierFromProvider,
	// VPC Lattice ServiceNetworkVPCAssociation can be imported using the id
	"aws_vpclattice_service_network_vpc_association": config.IdentifierFromProvider,
	// VPC Lattice Target Group can be imported using the id
	"aws_vpclattice_target_group": config.IdentifierFromProvider,
	// No import
	"aws_vpclattice_target_group_attachment": config.IdentifierFromProvider,

	// waf
	//
	// WAF Byte Match Set can be imported using the id
	"aws_waf_byte_match_set": config.IdentifierFromProvider,
	// WAF Geo Match Set can be imported using their ID
	"aws_waf_geo_match_set": config.IdentifierFromProvider,
	// WAF IPSets can be imported using their ID
	"aws_waf_ipset": config.IdentifierFromProvider,
	// WAF Rated Based Rule can be imported using the id
	"aws_waf_rate_based_rule": config.IdentifierFromProvider,
	// WAF Regex Match Set can be imported using their ID
	"aws_waf_regex_match_set": config.IdentifierFromProvider,
	// AWS WAF Regex Pattern Set can be imported using their ID
	"aws_waf_regex_pattern_set": config.IdentifierFromProvider,
	// WAF rules can be imported using the id
	"aws_waf_rule": config.IdentifierFromProvider,
	// AWS WAF Size Constraint Set can be imported using their ID
	"aws_waf_size_constraint_set": config.IdentifierFromProvider,
	// AWS WAF SQL Injection Match Set can be imported using their ID
	"aws_waf_sql_injection_match_set": config.IdentifierFromProvider,
	// WAF Web ACL can be imported using the id
	"aws_waf_web_acl": config.IdentifierFromProvider,
	// WAF XSS Match Set can be imported using their ID
	"aws_waf_xss_match_set": config.IdentifierFromProvider,

	// wafregional
	//
	// WAF Regional Byte Match Set can be imported using the id
	"aws_wafregional_byte_match_set": config.IdentifierFromProvider,
	// WAF Regional Geo Match Set can be imported using the id
	"aws_wafregional_geo_match_set": config.IdentifierFromProvider,
	// WAF Regional IPSets can be imported using their ID
	"aws_wafregional_ipset": config.IdentifierFromProvider,
	// WAF Regional Rate Based Rule can be imported using the id
	"aws_wafregional_rate_based_rule": config.IdentifierFromProvider,
	// WAF Regional Regex Match Set can be imported using the id
	"aws_wafregional_regex_match_set": config.IdentifierFromProvider,
	// WAF Regional Regex Pattern Set can be imported using the id
	"aws_wafregional_regex_pattern_set": config.IdentifierFromProvider,
	// WAF Regional Rule can be imported using the id
	"aws_wafregional_rule": config.IdentifierFromProvider,
	// WAF Size Constraint Set can be imported using the id
	"aws_wafregional_size_constraint_set": config.IdentifierFromProvider,
	// WAF Regional Sql Injection Match Set can be imported using the id
	"aws_wafregional_sql_injection_match_set": config.IdentifierFromProvider,
	// WAF Regional Web ACL can be imported using the id
	"aws_wafregional_web_acl": config.IdentifierFromProvider,
	// AWS WAF Regional XSS Match can be imported using the id
	"aws_wafregional_xss_match_set": config.IdentifierFromProvider,

	// wafv2
	//
	// WAFv2 IP Sets can be imported using ID/name/scope
	"aws_wafv2_ip_set": config.IdentifierFromProvider,
	// WAFv2 Regex Pattern Sets can be imported using ID/name/scope
	"aws_wafv2_regex_pattern_set": config.IdentifierFromProvider,
	// WAFv2 Rule Group can be imported using ID/name/scope
	"aws_wafv2_rule_group": config.IdentifierFromProvider,
	// WAFv2 Web ACL can be imported using ID/name/scope
	"aws_wafv2_web_acl": config.IdentifierFromProvider,
	// WAFv2 Web ACL Association using WEB_ACL_ARN,RESOURCE_ARN
	"aws_wafv2_web_acl_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.web_acl_arn }},{{ .parameters.resource_arn }}"),
	// WAFv2 Web ACL Logging Configurations using the ARN of the WAFv2 Web ACL
	"aws_wafv2_web_acl_logging_configuration": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_arn }}"),

	// workspaces
	//
	// Workspaces directory can be imported using the directory ID
	"aws_workspaces_directory": config.IdentifierFromProvider,
	// WorkSpaces IP groups can be imported using their GroupID
	"aws_workspaces_ip_group": config.IdentifierFromProvider,

	// xray
	//
	// XRay Encryption Config can be imported using the region name
	"aws_xray_encryption_config": config.IdentifierFromProvider,
	// XRay Groups can be imported using the ARN
	// Example: arn:aws:xray:us-west-2:1234567890:group/example-group/TNGX7SW5U6QY36T4ZMOUA3HVLBYCZTWDIOOXY3CJAXTHSS3YCWUA
	"aws_xray_group": config.IdentifierFromProvider,
	// XRay Sampling Rules can be imported using the name
	"aws_xray_sampling_rule": config.ParameterAsIdentifier("rule_name"),

	// ********** When adding new services please keep them alphabetized by their aws go sdk package name **********
}

var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}

// cognitoUserPoolClient
// Note(mbbush) This resource has some unexpected behaviors that make it impossible to write a completely correct
// ExternalName config. Specifically, the terraform id returned in the terraform state is not the same as the
// identifier used to import it. Additionally, if the terraform id set to an empty string, the terraform
// provider passes the empty string through to the aws query during refresh, which returns an api error.
// This could be related to the fact that this resource is implemented using the terraform plugin framework,
// which introduces the concept of a null value as distinct from a zero value.
func cognitoUserPoolClient() config.ExternalName {
	e := config.IdentifierFromProvider
	// TODO: Uncomment when it's acceptable to remove fields from spec.initProvider (major release)
	// e.IdentifierFields = []string{"user_pool_id"}
	e.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, cfg map[string]interface{}) (string, error) {
		if externalName == "" {
			return "invalidnonemptystring", nil
		}
		// Ideally, we'd return parameters.user_pool_id/external_name if this is invoked during a call to terraform import,
		// and the externalName if this is invoked during a call to terraform refresh. But I don't know how to distinguish
		// between them inside this function.
		return externalName, nil
	}
	return e
}

func mqUser() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, cfg map[string]interface{}) (string, error) {
		if externalName == "" {
			return "invalidnonemptystring", nil
		}
		return externalName, nil
	}
	return e
}

func lambdaFunctionURL() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, terraformProviderConfig map[string]interface{}) (string, error) {
		functionName, ok := parameters["function_name"]
		if !ok {
			return "", errors.New("function_name cannot be empty")
		}

		qualifier := parameters["qualifier"]
		if qualifier == nil || qualifier == "" {
			return functionName.(string), nil
		}
		return fmt.Sprintf("%s/%s", functionName.(string), qualifier.(string)), nil
	}
	return e
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
		var groups []string
		for _, g := range gs.([]interface{}) {
			groups = append(groups, g.(string))
		}
		return strings.Join(append([]string{u.(string)}, groups...), "/"), nil
	}
	return e
}

func iamPolicy() config.ExternalName {
	e := config.NameAsIdentifier

	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, setup map[string]any) (string, error) {
		path, ok := parameters["path"]
		if !ok {
			path = "/"
		}
		accountID := setup["client_metadata"].(map[string]string)["account_id"]
		partition := setup["client_metadata"].(map[string]string)["partition"]
		return fmt.Sprintf("arn:%s:iam::%s:policy%s%s", partition, accountID, path, externalName), nil
	}

	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id attribute missing from state file")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("value of id needs to be string")
		}

		arnSlice := strings.Split(idStr, "/")
		return arnSlice[len(arnSlice)-1], nil
	}

	return e
}

func kmsAlias() config.ExternalName {
	e := config.NameAsIdentifier
	e.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
		if _, ok := base["name"]; !ok {
			if !strings.HasPrefix(externalName, "alias/") {
				base["name"] = fmt.Sprintf("alias/%s", externalName)
			} else {
				base["name"] = externalName
			}
		}
	}
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id attribute missing from state file")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("value of id needs to be string")
		}

		return strings.TrimPrefix(idStr, "alias/"), nil
	}

	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]interface{}, _ map[string]interface{}) (string, error) {
		if !strings.HasPrefix(externalName, "alias/") {
			return fmt.Sprintf("alias/%s", externalName), nil
		}
		return externalName, nil
	}
	return e
}

func identifierFromProviderWithDefaultStub(defaultstub string) config.ExternalName {
	// Terraform does not always allow id to be empty.
	// Using a stub value to pass validation.
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if len(externalName) == 0 {
			return defaultstub, nil
		}
		return externalName, nil
	}
	return e
}

func vpcSecurityGroupRule() config.ExternalName {
	// Terraform does not allow security group rule id to be empty.
	// Using a stub value to pass validation.
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if len(externalName) == 0 {
			return "sgr-stub", nil
		}
		return externalName, nil
	}
	return e
}

func appConfigEnvironment() config.ExternalName {
	// Terraform does not allow Environment ID to be empty.
	// Using a stub value to pass validation.
	e := config.IdentifierFromProvider
	e.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
		if _, ok := base["environment_id"]; !ok {
			if externalName == "" {
				// must satisfy regular expression pattern: [a-z0-9]{4,7}
				base["environment_id"] = "tbdeid0"
			}
			if identifiers := strings.Split(externalName, ":"); len(identifiers) == 2 {
				base["environment_id"] = identifiers[0]
			}
		}
	}
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if len(externalName) == 0 {
			return "tbdeid0:tbdeid0", nil
		}
		return externalName, nil
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

func opensearchserverlessVpcEndpoint() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(ctx context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		// must match regex vpce-[0-9a-z]
		if len(externalName) == 0 {
			return "vpce-stubvpcendpoint999999", nil
		}
		return externalName, nil
	}
	return e
}

func opensearchserverlessCollection() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(ctx context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		// [a-z0-9]{3,40}
		if len(externalName) == 0 {
			return "stubcollection9999", nil
		}
		return externalName, nil
	}
	return e
}

// PermissionSetIdAsExternalName uses the id of the permission set (ps-80383020jr9302rk) as the external name, with
// the comma-separated pair permission_set_arn,instance_arn as the terraform id, when both arns are parameters and known
// ahead of time.
// Example: arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
func PermissionSetIdAsExternalName() config.ExternalName {
	return config.ExternalName{
		SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
		IdentifierFields:        []string{"instance_arn", "permission_set_arn"},
		GetExternalNameFn:       getPermissionSetId,
		GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, setup map[string]any) (string, error) {
			if externalName == "" {
				psa, ok := parameters["permission_set_arn"]
				if !ok {
					return "", errors.New("permission_set_arn cannot be empty")
				}
				psaStr, ok := psa.(string)
				if !ok {
					return "", errors.New("value of permission_set_arn needs to be a string")
				}
				externalName = strings.Split(psaStr, "/")[2]
			}
			ia, ok := parameters["instance_arn"]
			if !ok {
				return "", errors.New("instance_arn cannot be empty")
			}

			iaStr, ok := ia.(string)
			if !ok {
				return "", errors.New("value of instance_arn needs to be a string")
			}
			instanceId := strings.Split(iaStr, "/")[1]

			partition := setup["client_metadata"].(map[string]string)["partition"]
			return fmt.Sprintf("arn:%s:sso:::permissionSet/%s/%s,%s", partition, instanceId, externalName, iaStr), nil
		},
		DisableNameInitializer: true,
	}
}

// getPermissionSetId extracts the id of the permission set to use as an external name, from a terraform id formed by
// a comma-separated pair of ARNs, permission_set_arn,instance_arn.
// Example: arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
func getPermissionSetId(tfstate map[string]any) (string, error) {
	id, ok := tfstate["id"]
	if !ok {
		return "", errors.New("id does not exist in tfstate")
	}
	arn := strings.Split(id.(string), ",")[0]
	return strings.Split(arn, "/")[2], nil
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

// FormattedIdentifierUserDefinedNameLast is used in cases where the ID is constructed
// using some of the spec fields as well as a field that users use to name the
// resource. For example, vpc_id:cluster_name where vpc_id comes from spec
// but cluster_name is a naming field we can use external name for.
// This function assumes that the naming field is the LAST component
// in the constructed identifier, which may not always hold
// (e.g., aws_servicecatalog_budget_resource_association).
func FormattedIdentifierUserDefinedNameLast(param, separator string, keys ...string) config.ExternalName {
	e := config.ParameterAsIdentifier(param)
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

// FormattedIdentifierUserDefinedNameFirst is used in cases where the ID is constructed
// using some of the spec fields as well as a field that users use to name the
// resource. For example, budget_name:product_id where product_id comes from spec
// but budget_name is a naming field we can use external name for.
// This function assumes that the naming field is the FIRST component
// in the constructed identifier, which may not always hold
// (e.g., aws_eks_addon).
func FormattedIdentifierUserDefinedNameFirst(param, separator string, keys ...string) config.ExternalName {
	e := config.ParameterAsIdentifier(param)
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
			vals[i+1] = s
		}
		vals[0] = externalName
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
		return w[0], nil
	}
	return e
}

// TemplatedStringAsProviderDefinedIdentifier uses TemplatedStringAsIdentifier but
// without the name initializer, and with a GetIdFn that exits early if the external name is empty.
// This allows it to be used in cases where the ID is constructed with parameters and a provider-defined value, meaning
// no user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsProviderDefinedIdentifier(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	getId := e.GetIDFn
	e.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, cfg map[string]interface{}) (string, error) {
		if externalName == "" {
			return "", nil
		}
		return getId(ctx, externalName, parameters, cfg)
	}
	return e
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
// TODO: This seems to have some problems with handling the initial creation, when
// the parameters in the template are defined but the external name is empty, because
// the provider hasn't assigned its provider-defined identifier yet.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}

// ResourceConfigurator applies all external name configs listed in
// the table TerraformPluginSDKExternalNameConfigs,
// CLIReconciledExternalNameConfigs, and
// TerraformPluginFrameworkExternalNameConfigs and sets the version of
// those resources to v1beta1.
func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// If an external name is configured for multiple architectures,
		// Terraform Plugin Framework takes precedence over Terraform
		// Plugin SDKv2, which takes precedence over CLI architecture.
		e, configured := TerraformPluginFrameworkExternalNameConfigs[r.Name]
		if !configured {
			e, configured = TerraformPluginSDKExternalNameConfigs[r.Name]
			if !configured {
				e, configured = CLIReconciledExternalNameConfigs[r.Name]
			}
		}
		if !configured {
			return
		}
		r.Version = common.VersionV1Beta1
		r.ExternalName = e
		// Note(turkenh): This is special to provider-aws. We had injected
		// region as a parameter for all resources to be consistent with
		// the native aws provider, and now, we need to add manually it to
		// the identifier fields for all resources.
		r.ExternalName.IdentifierFields = append(r.ExternalName.IdentifierFields, "region")
	}
}

func eksOIDCIdentityProvider() config.ExternalName {
	return config.ExternalName{
		SetIdentifierArgumentFn: func(base map[string]interface{}, externalName string) {
			// max length is 1:
			// https://github.com/hashicorp/terraform-provider-aws/blob/7ff39c5b11aafe812e3a4b414aa6d345286b95ec/internal/service/eks/identity_provider_config.go#L58
			if arr, ok := base["oidc"].([]interface{}); ok && len(arr) == 1 {
				if m, ok := arr[0].(map[string]interface{}); ok {
					m["identity_provider_config_name"] = externalName
				}
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
		OmittedFields: []string{
			"oidc.identity_provider_config_name",
			"oidc.identity_provider_config_name_prefix",
		},
	}
}

func eksPodIdentityAssociation() config.ExternalName {
	e := config.IdentifierFromProvider

	// Terraform does not allow Association ID to be empty.
	// Using a stub value to pass validation.
	// Terraform does not use "id" attribute anymore, instead a combination of association_id and cluster_name is used.

	// If the association_id is equal to the stub value, we replace it with the external name.
	// Means that we are probably using a resource status that was not updated after creation due to write conflicts in the k8s API.
	// https://github.com/crossplane-contrib/provider-upjet-aws/issues/1437

	// must be 19 chars long and match regex ^a-[0-9a-z]*$
	stubAssocId := "a-stubassocid123456"
	e.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
		if assocId, ok := base["association_id"]; !ok || assocId == stubAssocId {
			if externalName == "" {
				base["association_id"] = stubAssocId
			} else {
				base["association_id"] = externalName
			}
		}
	}

	return e
}

func apiGatewayAccount() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(ctx context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if len(externalName) == 0 {
			// https://github.com/upbound/terraform-provider-aws/blob/e0753cc85e93ad146356533bf005c26591969299/internal/service/apigateway/account.go#L138-L137
			return "api-gateway-account", nil
		}
		return externalName, nil
	}
	return e
}

// fullARNTemplate builds a templated string for constructing a terraform id component which is an ARN, which includes
// the aws partition, service, region, account id, and resource. This is by far the most common form of ARN.
// e.g. arn:aws:ec2:ap-south-1:123456789012:instance/i-1234567890ab
func fullARNTemplate(service string, resource string) string {
	return genericARNTemplate(service, resource, false)

}

// regionlessARNTemplate builds a templated string for constructing a terraform id component which is an ARN of a
// resource which is regionless, but specific to your account id. It includes the partition, service, account id, and
// resource.
// e.g. arn:aws:iam::123456789012:role/example
func regionlessARNTemplate(service string, resource string) string {
	return genericARNTemplate(service, resource, true)
}

// genericARNTemplate builds a templated string for constructing a terraform id component which is an ARN of any format.
// It always includes the aws partition, service, and resource. Unless you specify to elide them, it will also include
// templates which resolve to the region (from the spec.forProvider) and the account id (calculated from the provider
// config).
func genericARNTemplate(service string, resource string, elideRegion bool) string {
	region := "{{ .setup.configuration.region }}"
	if elideRegion {
		region = ""
	}
	return fmt.Sprintf("arn:{{ .setup.client_metadata.partition }}:%s:%s:{{ .setup.client_metadata.account_id }}:%s", service, region, resource)
}

func rdsInstanceState() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["identifier"]
		if !ok {
			return "", errors.New("identifier field missing from tfstate")
		}
		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("identifier field must be a string")
		}
		return idStr, nil
	}
	return e
}

func s3LifecycleConfiguration() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["bucket"]
		if !ok {
			return "", errors.New("bucket field missing from tfstate")
		}
		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("bucket field must be a string")
		}
		return idStr, nil
	}
	return e
}

func dsqlClusterPeering() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["identifier"]
		if !ok {
			return "", errors.New("identifier field missing from tfstate")
		}
		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("identifier field must be a string")
		}
		return idStr, nil
	}
	return e
}