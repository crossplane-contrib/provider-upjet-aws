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

	// Access Analyzer Analyzers can be imported using the analyzer_name
	"aws_accessanalyzer_analyzer": config.ParameterAsIdentifier("analyzer_name"),

	// account

	// the Alternate Contact can be imported using the account_id and alternate_contact_type separated by a forward slash (/)
	"aws_account_alternate_contact": FormattedIdentifierUserDefined("/", "account_id", "alternate_contact_type"),

	// amp

	// The prometheus alert manager definition can be imported using the workspace identifier
	"aws_prometheus_alert_manager_definition": config.IdentifierFromProvider,
	// The prometheus rule group namespace can be imported using the arn
	"aws_prometheus_rule_group_namespace": config.IdentifierFromProvider,
	// AMP Workspaces can be imported using the identifier
	"aws_prometheus_workspace": config.IdentifierFromProvider,

	// amplify

	// Amplify App can be imported using Amplify App ID (appId)
	"aws_amplify_app": config.IdentifierFromProvider,
	// Amplify backend environment can be imported using app_id and environment_name
	"aws_amplify_backend_environment": FormattedIdentifierUserDefined("/", "app_id", "environment_name"),
	// Amplify branch can be imported using app_id and branch_name
	"aws_amplify_branch": FormattedIdentifierUserDefined("/", "app_id", "branch_name"),
	// Amplify domain association can be imported using app_id and domain_name
	"aws_amplify_domain_association": FormattedIdentifierUserDefined("/", "app_id", "domain_name"),
	// Amplify webhook can be imported using a webhook ID
	"aws_amplify_webhook": config.IdentifierFromProvider,

	// appautoscaling

	// Application AutoScaling Policy can be imported using the service-namespace , resource-id, scalable-dimension and policy-name separated by /
	"aws_appautoscaling_policy": FormattedIdentifierUserDefined("/", "service_namespace", "resource_id", "scalable_dimension", "name"),
	// No import
	"aws_appautoscaling_scheduled_action": config.IdentifierFromProvider,
	// Application AutoScaling Target can be imported using the service-namespace , resource-id and scalable-dimension separated by /
	"aws_appautoscaling_target": FormattedIdentifierUserDefined("/", "service_namespace", "resource_id", "scalable_dimension"),

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
	// AppFlow flows can be imported using the arn
	"aws_appflow_flow": config.IdentifierFromProvider,
}
