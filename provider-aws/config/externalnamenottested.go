/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{
	// accessanalyzer
	//
	// Access Analyzer Analyzers can be imported using the analyzer_name
	"aws_accessanalyzer_analyzer": config.ParameterAsIdentifier("analyzer_name"),

	// account
	//
	// The Alternate Contact for the current account can be imported using the alternate_contact_type
	"aws_account_alternate_contact": config.TemplatedStringAsIdentifier("", "{{ .parameters.alternate_contact_type }}"),

	// amp
	//
	// The prometheus alert manager definition can be imported using the workspace identifier
	"aws_prometheus_alert_manager_definition": config.ParameterAsIdentifier("workspace_id"),
	// The prometheus rule group namespace can be imported using the arn
	"aws_prometheus_rule_group_namespace": config.IdentifierFromProvider,
	// AMP Workspaces can be imported using the identifier
	"aws_prometheus_workspace": config.IdentifierFromProvider,

	// amplify
	//
	// Amplify App can be imported using Amplify App ID (appId)
	"aws_amplify_app": config.IdentifierFromProvider,
	// Amplify backend environment can be imported using app_id and environment_name: d2ypk4k47z8u6/example
	"aws_amplify_backend_environment": config.TemplatedStringAsIdentifier("environment_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify branch can be imported using app_id and branch_name: d2ypk4k47z8u6/master
	"aws_amplify_branch": config.TemplatedStringAsIdentifier("branch_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify domain association can be imported using app_id and domain_name: d2ypk4k47z8u6/example.com
	"aws_amplify_domain_association": config.TemplatedStringAsIdentifier("domain_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify webhook can be imported using a webhook ID
	"aws_amplify_webhook": config.IdentifierFromProvider,

	// appautoscaling
	//
	// Application AutoScaling Policy can be imported using the service-namespace, resource-id, scalable-dimension and policy-name separated by /
	"aws_appautoscaling_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.service_namespace }}/{{ .parameters.resource_id }}/{{ .parameters.scalable_dimension }}/{{ .external_name }}"),
	// No import
	"aws_appautoscaling_scheduled_action": config.IdentifierFromProvider,
	// Application AutoScaling Target can be imported using the service-namespace , resource-id and scalable-dimension separated by /
	"aws_appautoscaling_target": config.TemplatedStringAsIdentifier("", "{{ .parameters.service_namespace }}/{{ .parameters.resource_id }}/{{ .parameters.scalable_dimension }}"),

	// appconfig

	// AppConfig Applications can be imported using their application ID,
	"aws_appconfig_application": config.IdentifierFromProvider,
	// AppConfig Configuration Profiles can be imported by using the configuration profile ID and application ID separated by a colon (:)
	"aws_appconfig_configuration_profile": config.IdentifierFromProvider,
	// AppConfig Deployments can be imported by using the application ID, environment ID, and deployment number separated by a slash (/)
	"aws_appconfig_deployment": config.IdentifierFromProvider,
	// AppConfig Deployment Strategies can be imported by using their deployment strategy ID
	"aws_appconfig_deployment_strategy": config.IdentifierFromProvider,
	// AppConfig Environments can be imported by using the environment ID and application ID separated by a colon (:)
	"aws_appconfig_environment": config.IdentifierFromProvider,
	// AppConfig Hosted Configuration Versions can be imported by using the application ID, configuration profile ID, and version number separated by a slash (/)
	"aws_appconfig_hosted_configuration_version": config.IdentifierFromProvider,

	// appflow
	//
	// AppFlow flows can be imported using the arn
	"aws_appflow_flow": config.IdentifierFromProvider,

	// appintegrations
	//
	// Amazon AppIntegrations Event Integrations can be imported using the name
	"aws_appintegrations_event_integration": config.NameAsIdentifier,

	// appmesh
	//
	// App Mesh gateway routes can be imported using mesh_name and virtual_gateway_name together with the gateway route's name, e.g.,
	// mesh/gw1/example-gateway-route
	"aws_appmesh_gateway_route": config.TemplatedStringAsIdentifier("name", "{{ .parameters.mesh_name }}/{{ .parameters.virtual_gateway_name }}/{{ .external_name }}"),
	// App Mesh service meshes can be imported using the name
	"aws_appmesh_mesh": config.NameAsIdentifier,
	// App Mesh virtual routes can be imported using mesh_name and virtual_router_name together with the route's name, e.g.,
	// simpleapp/serviceB/serviceB-route
	"aws_appmesh_route": config.TemplatedStringAsIdentifier("name", "{{ .parameters.mesh_name }}/{{ .parameters.virtual_router_name }}/{{ .external_name }}"),
	// App Mesh virtual gateway can be imported using mesh_name together with the virtual gateway's name: mesh/gw1
	"aws_appmesh_virtual_gateway": config.TemplatedStringAsIdentifier("name", "{{ .parameters.mesh_name }}/{{ .external_name }}"),
	// App Mesh virtual nodes can be imported using mesh_name together with the virtual node's name: simpleapp/serviceBv1
	"aws_appmesh_virtual_node": config.TemplatedStringAsIdentifier("name", "{{ .parameters.mesh_name }}/{{ .external_name }}"),
	// App Mesh virtual routers can be imported using mesh_name together with the virtual router's name: simpleapp/serviceB
	"aws_appmesh_virtual_router": config.TemplatedStringAsIdentifier("name", "{{ .parameters.mesh_name }}/{{ .external_name }}"),
	// App Mesh virtual services can be imported using mesh_name together with the virtual service's name: simpleapp/servicea.simpleapp.local
	"aws_appmesh_virtual_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.mesh_name }}/{{ .external_name }}"),

	// apprunner
	//
	// App Runner AutoScaling Configuration Versions can be imported by using the arn
	"aws_apprunner_auto_scaling_configuration_version": config.IdentifierFromProvider,
	// App Runner Connections can be imported by using the connection_name
	"aws_apprunner_connection": config.ParameterAsIdentifier("connection_name"),
	// App Runner Custom Domain Associations can be imported by using the domain_name and service_arn separated by a comma (,)
	"aws_apprunner_custom_domain_association": config.TemplatedStringAsIdentifier("domain_name", "{{ .external_name }},{{ .parameters.service_arn }}"),
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
	"aws_appstream_fleet_stack_association": config.TemplatedStringAsIdentifier("stack_name", "{{ .parameters.fleet_name }}/{{ .external_name }}"),
	// aws_appstream_image_builder can be imported using the name
	"aws_appstream_image_builder": config.NameAsIdentifier,
	// aws_appstream_stack can be imported using the id
	"aws_appstream_stack": config.IdentifierFromProvider,
	// aws_appstream_user can be imported using the user_name and authentication_type separated by a slash (/)
	"aws_appstream_user": config.TemplatedStringAsIdentifier("user_name", "{{ .external_name }}/{{ .parameters.authentication_type }}"),
	// AppStream User Stack Association can be imported by using the user_name, authentication_type, and stack_name, separated by a slash (/)
	"aws_appstream_user_stack_association": config.TemplatedStringAsIdentifier("stack_name", "{{ .parameters.user_name }}/{{ .parameters.authentication_type }}/{{ .external_name }}/"),

	// appsync
	//
	// aws_appsync_api_cache can be imported using the AppSync API ID
	"aws_appsync_api_cache": config.IdentifierFromProvider,
	// aws_appsync_api_key can be imported using the AppSync API ID and key separated by :
	"aws_appsync_api_key": config.IdentifierFromProvider,
	// aws_appsync_datasource can be imported with their api_id, a hyphen, and name
	"aws_appsync_datasource": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}-{{ .external_name }}"),
	// aws_appsync_domain_name can be imported using the AppSync domain name
	"aws_appsync_domain_name": config.ParameterAsIdentifier("domain_name"),
	// aws_appsync_domain_name_api_association can be imported using the AppSync domain name
	"aws_appsync_domain_name_api_association": config.ParameterAsIdentifier("domain_name"),
	// aws_appsync_function can be imported using the AppSync API ID and Function ID separated by -
	"aws_appsync_function": config.IdentifierFromProvider,
	// AppSync GraphQL API can be imported using the GraphQL API ID
	"aws_appsync_graphql_api": config.IdentifierFromProvider,
	// aws_appsync_resolver can be imported with their api_id, a hyphen, type, a hypen and field
	"aws_appsync_resolver": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_id }}-{{ .parameters.type }}-{{ .parameters.field }}"),

	// autoscaling
	//
	// aws_autoscaling_group_tag can be imported by using the ASG name and key, separated by a comma (,)
	"aws_autoscaling_group_tag": config.TemplatedStringAsIdentifier("autoscaling_group_name", "{{ .external_name }},{{ .parameters.tag.key }}"),
	// AutoScaling Lifecycle Hooks can be imported using the role autoscaling_group_name and name separated by /
	"aws_autoscaling_lifecycle_hook": config.TemplatedStringAsIdentifier("name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// No import
	"aws_autoscaling_notification": config.IdentifierFromProvider,
	// AutoScaling scaling policy can be imported using the role autoscaling_group_name and name separated by /
	"aws_autoscaling_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// AutoScaling ScheduledAction can be imported using the auto-scaling-group-name and scheduled-action-name: auto-scaling-group-name/scheduled-action-name
	"aws_autoscaling_schedule": config.TemplatedStringAsIdentifier("scheduled_action_name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// Launch configurations can be imported using the name
	"aws_launch_configuration": config.NameAsIdentifier,

	// autoscalingplans
	//
	// Auto Scaling scaling plans can be imported using the name
	"aws_autoscalingplans_scaling_plan": config.NameAsIdentifier,

	// batch
	//
	// AWS Batch compute can be imported using the compute_environment_name
	"aws_batch_compute_environment": config.ParameterAsIdentifier("compute_environment_name"),
	// Batch Job Definition can be imported using the arn: arn:aws:batch:us-east-1:123456789012:job-definition/sample
	"aws_batch_job_definition": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:job-definition/{{ .external_name }}"),
	// Batch Job Queue can be imported using the arn: arn:aws:batch:us-east-1:123456789012:job-queue/sample
	"aws_batch_job_queue": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:job-queue/{{ .external_name }}"),
	// Batch Scheduling Policy can be imported using the arn: arn:aws:batch:us-east-1:123456789012:scheduling-policy/sample
	"aws_batch_scheduling_policy": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:scheduling-policy/{{ .external_name }}"),

	// budgets
	//
	// Budgets can be imported using AccountID:BudgetName
	"aws_budgets_budget": config.TemplatedStringAsIdentifier("name", "{{ .setup.client_metadata.account_id }}:{{ .external_name }}"),
	// Budgets can be imported using AccountID:ActionID:BudgetName
	"aws_budgets_budget_action": config.IdentifierFromProvider,

	// ce
	//
	// aws_ce_cost_category can be imported using the id
	"aws_ce_cost_category": config.IdentifierFromProvider,

	// chime
	//
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector": config.NameAsIdentifier,
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector_group": config.NameAsIdentifier,
	// Chime Voice Connector Logging can be imported using the voice_connector_id
	"aws_chime_voice_connector_logging": config.ParameterAsIdentifier("voice_connector_id"),
	// Chime Voice Connector Origination can be imported using the voice_connector_id
	"aws_chime_voice_connector_origination": config.ParameterAsIdentifier("voice_connector_id"),
	// Chime Voice Connector Streaming can be imported using the voice_connector_id
	"aws_chime_voice_connector_streaming": config.ParameterAsIdentifier("voice_connector_id"),
	// Chime Voice Connector Termination can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination": config.ParameterAsIdentifier("voice_connector_id"),
	// Chime Voice Connector Termination Credentials can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination_credentials": config.ParameterAsIdentifier("voice_connector_id"),

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
	// Cloudformation Stacks can be imported using the name
	"aws_cloudformation_stack": config.NameAsIdentifier,
	// CloudFormation StackSets can be imported using the name
	"aws_cloudformation_stack_set": config.NameAsIdentifier,
	//
	"aws_cloudformation_stack_set_instance": config.IdentifierFromProvider,
	// aws_cloudformation_type can be imported with their type version Amazon Resource Name (ARN)
	"aws_cloudformation_type": config.IdentifierFromProvider,

	// cloudhsmv2
	//
	// CloudHSM v2 Clusters can be imported using the cluster id
	"aws_cloudhsm_v2_cluster": config.IdentifierFromProvider,
	// HSM modules can be imported using their HSM ID
	"aws_cloudhsm_v2_hsm": config.IdentifierFromProvider,

	// cloudtrail
	//
	// Cloudtrails can be imported using the name
	"aws_cloudtrail": config.NameAsIdentifier,
	// Event data stores can be imported using their arn
	"aws_cloudtrail_event_data_store": config.IdentifierFromProvider,

	// cloudwatchlogs
	//
	// CloudWatch Logs destinations can be imported using the name
	"aws_cloudwatch_log_destination": config.NameAsIdentifier,
	// CloudWatch Logs destination policies can be imported using the destination_name
	"aws_cloudwatch_log_destination_policy": config.ParameterAsIdentifier("destination_name"),
	// CloudWatch Log Metric Filter can be imported using the log_group_name:name
	"aws_cloudwatch_log_metric_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_group_name }}:{{ .external_name }}"),
	// CloudWatch log resource policies can be imported using the policy name
	"aws_cloudwatch_log_resource_policy": config.ParameterAsIdentifier("policy_name"),
	// Cloudwatch Log Stream can be imported using the stream's log_group_name and name
	"aws_cloudwatch_log_stream": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_group_name }}:{{ .external_name }}"),
	// CloudWatch Logs subscription filter can be imported using the log group name and subscription filter name separated by |
	"aws_cloudwatch_log_subscription_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_group_name }}|{{ .external_name }}"),
	// CloudWatch query definitions can be imported using the query definition ARN.
	"aws_cloudwatch_query_definition": config.IdentifierFromProvider,

	// codeartifact
	//
	// CodeArtifact Domain can be imported using the CodeArtifact Domain arn
	"aws_codeartifact_domain": config.IdentifierFromProvider,
	// CodeArtifact Domain Permissions Policies can be imported using the CodeArtifact Domain ARN
	"aws_codeartifact_domain_permissions_policy": config.IdentifierFromProvider,
	// CodeArtifact Repository can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository": config.IdentifierFromProvider,
	// CodeArtifact Repository Permissions Policies can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository_permissions_policy": config.IdentifierFromProvider,

	// codebuild
	//
	// CodeBuild Project can be imported using the name
	"aws_codebuild_project": config.NameAsIdentifier,
	// CodeBuild Report Group can be imported using the CodeBuild Report Group arn: arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
	"aws_codebuild_report_group": config.TemplatedStringAsIdentifier("name", "arn:aws:codebuild:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:report-group/{{ .external_name }}"),
	// CodeBuild Resource Policy can be imported using the CodeBuild Resource Policy arn
	"aws_codebuild_resource_policy": config.IdentifierFromProvider,
	// CodeBuild Source Credential can be imported using the CodeBuild Source Credential arn: arn:aws:codebuild:us-west-2:123456789:token:github
	"aws_codebuild_source_credential": config.TemplatedStringAsIdentifier("", "arn:aws:codebuild:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:token:{{ .parameters.token }}"),
	// CodeBuild Webhooks can be imported using the CodeBuild Project name
	"aws_codebuild_webhook": config.ParameterAsIdentifier("project_name"),

	// codecommit
	//
	// CodeCommit approval rule templates can be imported using the name
	"aws_codecommit_approval_rule_template": config.NameAsIdentifier,
	// CodeCommit approval rule template associations can be imported using the approval_rule_template_name and repository_name separated by a comma (,)
	"aws_codecommit_approval_rule_template_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.approval_rule_template_name }},.parameters.repository_name }}"),
	// Codecommit repository can be imported using repository name
	"aws_codecommit_repository": config.NameAsIdentifier,
	// No import
	"aws_codecommit_trigger": config.IdentifierFromProvider,

	// codepipeline
	//
	// CodePipelines can be imported using the name
	"aws_codepipeline": config.NameAsIdentifier,
	// CodePipeline Webhooks can be imported by their ARN: arn:aws:codepipeline:us-west-2:123456789012:webhook:example
	"aws_codepipeline_webhook": config.TemplatedStringAsIdentifier("name", "arn:aws:codepipeline:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:webhook:{{ .external_name }}"),

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

	// cognitoidp
	//
	// Cognito User Groups can be imported using the user_pool_id/name attributes concatenated
	"aws_cognito_user_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.user_pool_id }}/{{ .external_name }}"),
	// No import
	"aws_cognito_user_in_group": config.IdentifierFromProvider,

	// configservice
	//
	// Config aggregate authorizations can be imported using account_id:region
	"aws_config_aggregate_authorization": config.TemplatedStringAsIdentifier("", "{{ .parameters.account_id }}:{{ .parameters.region }}"),
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
	// Config Organization Conformance Packs can be imported using the name
	"aws_config_organization_conformance_pack": config.NameAsIdentifier,
	// Config Organization Custom Rules can be imported using the name
	"aws_config_organization_custom_rule": config.NameAsIdentifier,
	// Config Organization Managed Rules can be imported using the name
	"aws_config_organization_managed_rule": config.NameAsIdentifier,
	// Remediation Configurations can be imported using the name config_rule_name
	"aws_config_remediation_configuration": config.ParameterAsIdentifier("config_rule_name"),
}
