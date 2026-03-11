// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bedrockagentcore

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_bedrockagentcore_agent_runtime", func(r *config.Resource) {
		r.AddSingletonListConversion("agent_runtime_artifact", "agentRuntimeArtifact")
		r.AddSingletonListConversion("agent_runtime_artifact[*].code_configuration", "agentRuntimeArtifact[*].codeConfiguration")
		r.AddSingletonListConversion("agent_runtime_artifact[*].code_configuration[*].code", "agentRuntimeArtifact[*].codeConfiguration[*].code")
		r.AddSingletonListConversion("agent_runtime_artifact[*].code_configuration[*].code[*].s3", "agentRuntimeArtifact[*].codeConfiguration[*].code[*].s3")
		r.AddSingletonListConversion("agent_runtime_artifact[*].container_configuration", "agentRuntimeArtifact[*].containerConfiguration")
		r.AddSingletonListConversion("authorizer_configuration", "authorizerConfiguration")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer", "authorizerConfiguration[*].customJwtAuthorizer")
		r.AddSingletonListConversion("lifecycle_configuration", "lifecycleConfiguration")
		r.AddSingletonListConversion("network_configuration", "networkConfiguration")
		r.AddSingletonListConversion("network_configuration[*].network_mode_config", "networkConfiguration[*].networkModeConfig")
		r.AddSingletonListConversion("protocol_configuration", "protocolConfiguration")
		r.AddSingletonListConversion("request_header_configuration", "requestHeaderConfiguration")
	})

	p.AddResourceConfigurator("aws_bedrockagentcore_api_key_credential_provider", func(r *config.Resource) {
		r.TerraformResource.Schema["name"].Computed = true
		r.TerraformResource.Schema["name"].Optional = false
	})

	// aws_bedrockagentcore_browser
	p.AddResourceConfigurator("aws_bedrockagentcore_browser", func(r *config.Resource) {
		r.AddSingletonListConversion("network_configuration", "networkConfiguration")
		r.AddSingletonListConversion("network_configuration[*].vpc_config", "networkConfiguration[*].vpcConfig")
		r.AddSingletonListConversion("recording", "recording")
		r.AddSingletonListConversion("recording[*].s3_location", "recording[*].s3Location")
	})

	// aws_bedrockagentcore_code_interpreter
	p.AddResourceConfigurator("aws_bedrockagentcore_code_interpreter", func(r *config.Resource) {
		r.AddSingletonListConversion("network_configuration", "networkConfiguration")
		r.AddSingletonListConversion("network_configuration[*].vpc_config", "networkConfiguration[*].vpcConfig")
	})
	// aws_bedrockagentcore_gateway
	p.AddResourceConfigurator("aws_bedrockagentcore_gateway", func(r *config.Resource) {
		r.AddSingletonListConversion("authorizer_configuration", "authorizerConfiguration")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer", "authorizerConfiguration[*].customJwtAuthorizer")
		r.AddSingletonListConversion("interceptor_configuration[*].input_configuration", "interceptorConfiguration[*].inputConfiguration")
		r.AddSingletonListConversion("interceptor_configuration[*].interceptor", "interceptorConfiguration[*].interceptor")
		r.AddSingletonListConversion("interceptor_configuration[*].interceptor[*].lambda", "interceptorConfiguration[*].interceptor[*].lambda")
		r.AddSingletonListConversion("protocol_configuration", "protocolConfiguration")
		r.AddSingletonListConversion("protocol_configuration[*].mcp", "protocolConfiguration[*].mcp")
	})
	// aws_bedrockagentcore_gateway_target
	p.AddResourceConfigurator("aws_bedrockagentcore_gateway_target", func(r *config.Resource) {
		r.AddSingletonListConversion("credential_provider_configuration", "credentialProviderConfiguration")
		r.AddSingletonListConversion("credential_provider_configuration[*].api_key", "credentialProviderConfiguration[*].apiKey")
		r.AddSingletonListConversion("credential_provider_configuration[*].gateway_iam_role", "credentialProviderConfiguration[*].gatewayIamRole")
		r.AddSingletonListConversion("credential_provider_configuration[*].oauth", "credentialProviderConfiguration[*].oauth")
		r.AddSingletonListConversion("metadata_configuration", "metadataConfiguration")
		r.AddSingletonListConversion("target_configuration", "targetConfiguration")
		r.AddSingletonListConversion("target_configuration[*].mcp", "targetConfiguration[*].mcp")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda", "targetConfiguration[*].mcp[*].lambda")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema", "targetConfiguration[*].mcp[*].lambda[*].toolSchema")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].input_schema", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].inputSchema")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].input_schema[*].items", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].inputSchema[*].items")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].input_schema[*].items[*].items", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].inputSchema[*].items[*].items")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].input_schema[*].property[*].items", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].inputSchema[*].property[*].items")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].input_schema[*].property[*].items[*].items", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].inputSchema[*].property[*].items[*].items")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].output_schema", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].outputSchema")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].output_schema[*].items", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].outputSchema[*].items")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].output_schema[*].items[*].items", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].outputSchema[*].items[*].items")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].output_schema[*].property[*].items", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].outputSchema[*].property[*].items")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].inline_payload[*].output_schema[*].property[*].items[*].items", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].inlinePayload[*].outputSchema[*].property[*].items[*].items")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].lambda[*].tool_schema[*].s3", "targetConfiguration[*].mcp[*].lambda[*].toolSchema[*].s3")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].mcp_server", "targetConfiguration[*].mcp[*].mcpServer")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].open_api_schema", "targetConfiguration[*].mcp[*].openApiSchema")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].open_api_schema[*].inline_payload", "targetConfiguration[*].mcp[*].openApiSchema[*].inlinePayload")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].open_api_schema[*].s3", "targetConfiguration[*].mcp[*].openApiSchema[*].s3")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].smithy_model", "targetConfiguration[*].mcp[*].smithyModel")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].smithy_model[*].inline_payload", "targetConfiguration[*].mcp[*].smithyModel[*].inlinePayload")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].smithy_model[*].s3", "targetConfiguration[*].mcp[*].smithyModel[*].s3")
	})
	// aws_bedrockagentcore_memory_strategy
	p.AddResourceConfigurator("aws_bedrockagentcore_memory_strategy", func(r *config.Resource) {
		r.AddSingletonListConversion("configuration", "configuration")
		r.AddSingletonListConversion("configuration[*].consolidation", "configuration[*].consolidation")
		r.AddSingletonListConversion("configuration[*].extraction", "configuration[*].extraction")
	})
	// aws_bedrockagentcore_oauth2_credential_provider
	p.AddResourceConfigurator("aws_bedrockagentcore_oauth2_credential_provider", func(r *config.Resource) {
		r.TerraformResource.Schema["name"].Computed = true
		r.TerraformResource.Schema["name"].Optional = false

		r.AddSingletonListConversion("oauth2_provider_config", "oauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].custom_oauth2_provider_config", "oauth2ProviderConfig[*].customOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].custom_oauth2_provider_config[*].oauth_discovery", "oauth2ProviderConfig[*].customOauth2ProviderConfig[*].oauthDiscovery")
		r.AddSingletonListConversion("oauth2_provider_config[*].custom_oauth2_provider_config[*].oauth_discovery[*].authorization_server_metadata", "oauth2ProviderConfig[*].customOauth2ProviderConfig[*].oauthDiscovery[*].authorizationServerMetadata")
		r.AddSingletonListConversion("oauth2_provider_config[*].github_oauth2_provider_config", "oauth2ProviderConfig[*].githubOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].google_oauth2_provider_config", "oauth2ProviderConfig[*].googleOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].microsoft_oauth2_provider_config", "oauth2ProviderConfig[*].microsoftOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].salesforce_oauth2_provider_config", "oauth2ProviderConfig[*].salesforceOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].slack_oauth2_provider_config", "oauth2ProviderConfig[*].slackOauth2ProviderConfig")
	})
	// aws_bedrockagentcore_token_vault_cmk
	p.AddResourceConfigurator("aws_bedrockagentcore_token_vault_cmk", func(r *config.Resource) {
		r.AddSingletonListConversion("kms_configuration", "kmsConfiguration")
	})
	p.AddResourceConfigurator("aws_bedrockagentcore_workload_identity", func(r *config.Resource) {
		r.TerraformResource.Schema["name"].Computed = true
		r.TerraformResource.Schema["name"].Optional = false
	})
}
