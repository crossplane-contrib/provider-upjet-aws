// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bedrockagentcore

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_bedrockagentcore_agent_runtime", func(r *config.Resource) {
		r.References["filesystem_configuration.efs_access_point.access_point_arn"] = config.Reference{
			TerraformName: "aws_efs_access_point",
			Extractor:     common.PathARNExtractor,
		}
		r.References["filesystem_configuration.s3_files_access_point.access_point_arn"] = config.Reference{
			TerraformName: "aws_s3_access_point",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("agent_runtime_artifact", "agentRuntimeArtifact")
		r.AddSingletonListConversion("agent_runtime_artifact[*].code_configuration", "agentRuntimeArtifact[*].codeConfiguration")
		r.AddSingletonListConversion("agent_runtime_artifact[*].code_configuration[*].code", "agentRuntimeArtifact[*].codeConfiguration[*].code")
		r.AddSingletonListConversion("agent_runtime_artifact[*].code_configuration[*].code[*].s3", "agentRuntimeArtifact[*].codeConfiguration[*].code[*].s3")
		r.AddSingletonListConversion("agent_runtime_artifact[*].container_configuration", "agentRuntimeArtifact[*].containerConfiguration")
		r.AddSingletonListConversion("authorizer_configuration", "authorizerConfiguration")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer", "authorizerConfiguration[*].customJwtAuthorizer")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].allowed_workload_configuration", "authorizerConfiguration[*].customJwtAuthorizer[*].allowedWorkloadConfiguration")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].custom_claim[*].authorizing_claim_match_value", "authorizerConfiguration[*].customJwtAuthorizer[*].customClaim[*].authorizingClaimMatchValue")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].custom_claim[*].authorizing_claim_match_value[*].claim_match_value", "authorizerConfiguration[*].customJwtAuthorizer[*].customClaim[*].authorizingClaimMatchValue[*].claimMatchValue")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint[*].managed_vpc_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint[*].self_managed_lattice_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint[*].managed_vpc_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint[*].self_managed_lattice_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("filesystem_configuration[*].efs_access_point", "filesystemConfiguration[*].efsAccessPoint")
		r.AddSingletonListConversion("filesystem_configuration[*].s3_files_access_point", "filesystemConfiguration[*].s3FilesAccessPoint")
		r.AddSingletonListConversion("filesystem_configuration[*].session_storage", "filesystemConfiguration[*].sessionStorage")
		r.AddSingletonListConversion("lifecycle_configuration", "lifecycleConfiguration")
		r.AddSingletonListConversion("network_configuration", "networkConfiguration")
		r.AddSingletonListConversion("network_configuration[*].network_mode_config", "networkConfiguration[*].networkModeConfig")
		r.AddSingletonListConversion("protocol_configuration", "protocolConfiguration")
		r.AddSingletonListConversion("request_header_configuration", "requestHeaderConfiguration")
	})

	p.AddResourceConfigurator("aws_bedrockagentcore_api_key_credential_provider", func(r *config.Resource) {
		r.TerraformResource.Schema["name"].Computed = true
		r.TerraformResource.Schema["name"].Optional = false
		r.AddSingletonListConversion("api_key_secret_config", "apiKeySecretConfig")
	})

	// aws_bedrockagentcore_browser
	p.AddResourceConfigurator("aws_bedrockagentcore_browser", func(r *config.Resource) {
		r.References["certificate.location.secrets_manager.secret_arn"] = config.Reference{
			TerraformName: "aws_secretsmanager_secret",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("browser_signing", "browserSigning")
		r.AddSingletonListConversion("certificate[*].location", "certificate[*].location")
		r.AddSingletonListConversion("certificate[*].location[*].secrets_manager", "certificate[*].location[*].secretsManager")
		r.AddSingletonListConversion("enterprise_policy[*].location", "enterprisePolicy[*].location")
		r.AddSingletonListConversion("enterprise_policy[*].location[*].s3", "enterprisePolicy[*].location[*].s3")
		r.AddSingletonListConversion("network_configuration", "networkConfiguration")
		r.AddSingletonListConversion("network_configuration[*].vpc_config", "networkConfiguration[*].vpcConfig")
		r.AddSingletonListConversion("recording", "recording")
		r.AddSingletonListConversion("recording[*].s3_location", "recording[*].s3Location")
	})

	// aws_bedrockagentcore_code_interpreter
	p.AddResourceConfigurator("aws_bedrockagentcore_code_interpreter", func(r *config.Resource) {
		r.References["certificate.location.secrets_manager.secret_arn"] = config.Reference{
			TerraformName: "aws_secretsmanager_secret",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("certificate[*].location", "certificate[*].location")
		r.AddSingletonListConversion("certificate[*].location[*].secrets_manager", "certificate[*].location[*].secretsManager")
		r.AddSingletonListConversion("network_configuration", "networkConfiguration")
		r.AddSingletonListConversion("network_configuration[*].vpc_config", "networkConfiguration[*].vpcConfig")
	})
	// aws_bedrockagentcore_evaluator
	p.AddResourceConfigurator("aws_bedrockagentcore_evaluator", func(r *config.Resource) {
		r.AddSingletonListConversion("evaluator_config", "evaluatorConfig")
		r.AddSingletonListConversion("evaluator_config[*].code_based", "evaluatorConfig[*].codeBased")
		r.AddSingletonListConversion("evaluator_config[*].code_based[*].lambda_config", "evaluatorConfig[*].codeBased[*].lambdaConfig")
		r.AddSingletonListConversion("evaluator_config[*].llm_as_a_judge", "evaluatorConfig[*].llmAsAJudge")
		r.AddSingletonListConversion("evaluator_config[*].llm_as_a_judge[*].model_config", "evaluatorConfig[*].llmAsAJudge[*].modelConfig")
		r.AddSingletonListConversion("evaluator_config[*].llm_as_a_judge[*].model_config[*].bedrock_evaluator_model_config", "evaluatorConfig[*].llmAsAJudge[*].modelConfig[*].bedrockEvaluatorModelConfig")
		r.AddSingletonListConversion("evaluator_config[*].llm_as_a_judge[*].model_config[*].bedrock_evaluator_model_config[*].inference_config", "evaluatorConfig[*].llmAsAJudge[*].modelConfig[*].bedrockEvaluatorModelConfig[*].inferenceConfig")
		r.AddSingletonListConversion("evaluator_config[*].llm_as_a_judge[*].rating_scale", "evaluatorConfig[*].llmAsAJudge[*].ratingScale")
	})
	// aws_bedrockagentcore_gateway
	p.AddResourceConfigurator("aws_bedrockagentcore_gateway", func(r *config.Resource) {
		r.References["policy_engine_configuration.arn"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_policy_engine",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("policy_engine_arn",true)`,
		}
		r.AddSingletonListConversion("authorizer_configuration", "authorizerConfiguration")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer", "authorizerConfiguration[*].customJwtAuthorizer")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].allowed_workload_configuration", "authorizerConfiguration[*].customJwtAuthorizer[*].allowedWorkloadConfiguration")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].custom_claim[*].authorizing_claim_match_value", "authorizerConfiguration[*].customJwtAuthorizer[*].customClaim[*].authorizingClaimMatchValue")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].custom_claim[*].authorizing_claim_match_value[*].claim_match_value", "authorizerConfiguration[*].customJwtAuthorizer[*].customClaim[*].authorizingClaimMatchValue[*].claimMatchValue")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint[*].managed_vpc_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint[*].self_managed_lattice_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint[*].managed_vpc_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint[*].self_managed_lattice_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("interceptor_configuration[*].input_configuration", "interceptorConfiguration[*].inputConfiguration")
		r.AddSingletonListConversion("interceptor_configuration[*].interceptor", "interceptorConfiguration[*].interceptor")
		r.AddSingletonListConversion("interceptor_configuration[*].interceptor[*].lambda", "interceptorConfiguration[*].interceptor[*].lambda")
		r.AddSingletonListConversion("policy_engine_configuration", "policyEngineConfiguration")
		r.AddSingletonListConversion("protocol_configuration", "protocolConfiguration")
		r.AddSingletonListConversion("protocol_configuration[*].mcp", "protocolConfiguration[*].mcp")
		r.AddSingletonListConversion("protocol_configuration[*].mcp[*].session_configuration", "protocolConfiguration[*].mcp[*].sessionConfiguration")
		r.AddSingletonListConversion("protocol_configuration[*].mcp[*].streaming_configuration", "protocolConfiguration[*].mcp[*].streamingConfiguration")
	})
	// aws_bedrockagentcore_gateway_target
	p.AddResourceConfigurator("aws_bedrockagentcore_gateway_target", func(r *config.Resource) {
		r.References["credential_provider_configuration.oauth.provider_arn"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_oauth2_credential_provider",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("credential_provider_arn",true)`,
		}
		r.References["credential_provider_configuration.api_key.provider_arn"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_api_key_credential_provider",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("credential_provider_arn",true)`,
		}
		r.References["target_configuration.mcp.api_gateway.rest_api_id"] = config.Reference{
			TerraformName: "aws_api_gateway_rest_api",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()`,
		}
		r.AddSingletonListConversion("credential_provider_configuration", "credentialProviderConfiguration")
		r.AddSingletonListConversion("credential_provider_configuration[*].api_key", "credentialProviderConfiguration[*].apiKey")
		r.AddSingletonListConversion("credential_provider_configuration[*].caller_iam_credentials", "credentialProviderConfiguration[*].callerIamCredentials")
		r.AddSingletonListConversion("credential_provider_configuration[*].gateway_iam_role", "credentialProviderConfiguration[*].gatewayIamRole")
		r.AddSingletonListConversion("credential_provider_configuration[*].jwt_passthrough", "credentialProviderConfiguration[*].jwtPassthrough")
		r.AddSingletonListConversion("credential_provider_configuration[*].oauth", "credentialProviderConfiguration[*].oauth")
		r.AddSingletonListConversion("metadata_configuration", "metadataConfiguration")
		r.AddSingletonListConversion("private_endpoint", "privateEndpoint")
		r.AddSingletonListConversion("private_endpoint[*].managed_vpc_resource", "privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("private_endpoint[*].self_managed_lattice_resource", "privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("target_configuration", "targetConfiguration")
		r.AddSingletonListConversion("target_configuration[*].http", "targetConfiguration[*].http")
		r.AddSingletonListConversion("target_configuration[*].http[*].agentcore_runtime", "targetConfiguration[*].http[*].agentcoreRuntime")
		r.AddSingletonListConversion("target_configuration[*].http[*].agentcore_runtime[*].schema", "targetConfiguration[*].http[*].agentcoreRuntime[*].schema")
		r.AddSingletonListConversion("target_configuration[*].http[*].agentcore_runtime[*].schema[*].source", "targetConfiguration[*].http[*].agentcoreRuntime[*].schema[*].source")
		r.AddSingletonListConversion("target_configuration[*].http[*].agentcore_runtime[*].schema[*].source[*].inline_payload", "targetConfiguration[*].http[*].agentcoreRuntime[*].schema[*].source[*].inlinePayload")
		r.AddSingletonListConversion("target_configuration[*].http[*].agentcore_runtime[*].schema[*].source[*].s3", "targetConfiguration[*].http[*].agentcoreRuntime[*].schema[*].source[*].s3")
		r.AddSingletonListConversion("target_configuration[*].http[*].passthrough", "targetConfiguration[*].http[*].passthrough")
		r.AddSingletonListConversion("target_configuration[*].http[*].passthrough[*].schema", "targetConfiguration[*].http[*].passthrough[*].schema")
		r.AddSingletonListConversion("target_configuration[*].http[*].passthrough[*].schema[*].source", "targetConfiguration[*].http[*].passthrough[*].schema[*].source")
		r.AddSingletonListConversion("target_configuration[*].http[*].passthrough[*].schema[*].source[*].inline_payload", "targetConfiguration[*].http[*].passthrough[*].schema[*].source[*].inlinePayload")
		r.AddSingletonListConversion("target_configuration[*].http[*].passthrough[*].schema[*].source[*].s3", "targetConfiguration[*].http[*].passthrough[*].schema[*].source[*].s3")
		r.AddSingletonListConversion("target_configuration[*].http[*].passthrough[*].stickiness_configuration", "targetConfiguration[*].http[*].passthrough[*].stickinessConfiguration")
		r.AddSingletonListConversion("target_configuration[*].inference", "targetConfiguration[*].inference")
		r.AddSingletonListConversion("target_configuration[*].inference[*].connector", "targetConfiguration[*].inference[*].connector")
		r.AddSingletonListConversion("target_configuration[*].inference[*].connector[*].source", "targetConfiguration[*].inference[*].connector[*].source")
		r.AddSingletonListConversion("target_configuration[*].inference[*].provider", "targetConfiguration[*].inference[*].provider")
		r.AddSingletonListConversion("target_configuration[*].inference[*].provider[*].model_mapping", "targetConfiguration[*].inference[*].provider[*].modelMapping")
		r.AddSingletonListConversion("target_configuration[*].inference[*].provider[*].model_mapping[*].provider_prefix", "targetConfiguration[*].inference[*].provider[*].modelMapping[*].providerPrefix")
		r.AddSingletonListConversion("target_configuration[*].mcp", "targetConfiguration[*].mcp")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].api_gateway", "targetConfiguration[*].mcp[*].apiGateway")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].api_gateway[*].api_gateway_tool_configuration", "targetConfiguration[*].mcp[*].apiGateway[*].apiGatewayToolConfiguration")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].connector", "targetConfiguration[*].mcp[*].connector")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].connector[*].source", "targetConfiguration[*].mcp[*].connector[*].source")
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
		r.AddSingletonListConversion("target_configuration[*].mcp[*].mcp_server[*].mcp_tool_schema", "targetConfiguration[*].mcp[*].mcpServer[*].mcpToolSchema")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].mcp_server[*].mcp_tool_schema[*].inline_payload", "targetConfiguration[*].mcp[*].mcpServer[*].mcpToolSchema[*].inlinePayload")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].mcp_server[*].mcp_tool_schema[*].s3", "targetConfiguration[*].mcp[*].mcpServer[*].mcpToolSchema[*].s3")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].open_api_schema", "targetConfiguration[*].mcp[*].openApiSchema")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].open_api_schema[*].inline_payload", "targetConfiguration[*].mcp[*].openApiSchema[*].inlinePayload")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].open_api_schema[*].s3", "targetConfiguration[*].mcp[*].openApiSchema[*].s3")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].smithy_model", "targetConfiguration[*].mcp[*].smithyModel")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].smithy_model[*].inline_payload", "targetConfiguration[*].mcp[*].smithyModel[*].inlinePayload")
		r.AddSingletonListConversion("target_configuration[*].mcp[*].smithy_model[*].s3", "targetConfiguration[*].mcp[*].smithyModel[*].s3")
	})
	// aws_bedrockagentcore_harness
	p.AddResourceConfigurator("aws_bedrockagentcore_harness", func(r *config.Resource) {
		r.References["memory.agentcore_memory_configuration.arn"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_memory",
			Extractor:     common.PathARNExtractor,
		}
		r.References["memory.agentcore_memory_configuration.retrieval_config.strategy_id"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_memory_strategy",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("memory_strategy_id",true)`,
		}
		r.References["tool.config.agentcore_browser.browser_arn"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_browser",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("browser_arn",true)`,
		}
		r.References["tool.config.agentcore_code_interpreter.code_interpreter_arn"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_code_interpreter",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("code_interpreter_arn",true)`,
		}
		r.References["tool.config.agentcore_gateway.gateway_arn"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_gateway",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("gateway_arn",true)`,
		}
		r.References["tool.config.agentcore_gateway.outbound_auth.oauth.provider_arn"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_oauth2_credential_provider",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("credential_provider_arn",true)`,
		}
		r.References["environment.agentcore_runtime_environment.filesystem_configuration.efs_access_point.access_point_arn"] = config.Reference{
			TerraformName: "aws_efs_access_point",
			Extractor:     common.PathARNExtractor,
		}
		r.References["environment.agentcore_runtime_environment.filesystem_configuration.s3_files_access_point.access_point_arn"] = config.Reference{
			TerraformName: "aws_s3_access_point",
			Extractor:     common.PathARNExtractor,
		}
		r.References["model.gemini_model_config.api_key_arn"] = config.Reference{
			TerraformName: "aws_secretsmanager_secret",
			Extractor:     common.PathARNExtractor,
		}
		r.References["model.openai_model_config.api_key_arn"] = config.Reference{
			TerraformName: "aws_secretsmanager_secret",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("authorizer_configuration", "authorizerConfiguration")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer", "authorizerConfiguration[*].customJwtAuthorizer")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].allowed_workload_configuration", "authorizerConfiguration[*].customJwtAuthorizer[*].allowedWorkloadConfiguration")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].custom_claim[*].authorizing_claim_match_value", "authorizerConfiguration[*].customJwtAuthorizer[*].customClaim[*].authorizingClaimMatchValue")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].custom_claim[*].authorizing_claim_match_value[*].claim_match_value", "authorizerConfiguration[*].customJwtAuthorizer[*].customClaim[*].authorizingClaimMatchValue[*].claimMatchValue")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint[*].managed_vpc_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint[*].self_managed_lattice_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint[*].managed_vpc_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint[*].managedVpcResource")
		r.AddSingletonListConversion("authorizer_configuration[*].custom_jwt_authorizer[*].private_endpoint_overrides[*].private_endpoint[*].self_managed_lattice_resource", "authorizerConfiguration[*].customJwtAuthorizer[*].privateEndpointOverrides[*].privateEndpoint[*].selfManagedLatticeResource")
		r.AddSingletonListConversion("environment", "environment")
		// note: conversions below are intentionally commented-out.
		// `environment[*].agentcore_runtime_environment` TF schema changed
		// from nested list attribute to nested block list with SizeAtMost=1
		// validation, during TF v6.55 -> v6.62.
		// We are now eligible for a singleton list conversion, but this
		// will change the MR API. We opt out to not cause a version bump,
		// agentcore_runtime_environment stays as a list in CRD schema.
		// r.AddSingletonListConversion("environment[*].agentcore_runtime_environment", "environment[*].agentcoreRuntimeEnvironment")
		// r.AddSingletonListConversion("environment[*].agentcore_runtime_environment[*].filesystem_configuration[*].efs_access_point", "environment[*].agentcoreRuntimeEnvironment[*].filesystemConfiguration[*].efsAccessPoint")
		// r.AddSingletonListConversion("environment[*].agentcore_runtime_environment[*].filesystem_configuration[*].s3_files_access_point", "environment[*].agentcoreRuntimeEnvironment[*].filesystemConfiguration[*].s3FilesAccessPoint")
		// r.AddSingletonListConversion("environment[*].agentcore_runtime_environment[*].filesystem_configuration[*].session_storage", "environment[*].agentcoreRuntimeEnvironment[*].filesystemConfiguration[*].sessionStorage")
		// r.AddSingletonListConversion("environment[*].agentcore_runtime_environment[*].lifecycle_configuration", "environment[*].agentcoreRuntimeEnvironment[*].lifecycleConfiguration")
		// r.AddSingletonListConversion("environment[*].agentcore_runtime_environment[*].network_configuration", "environment[*].agentcoreRuntimeEnvironment[*].networkConfiguration")
		// r.AddSingletonListConversion("environment[*].agentcore_runtime_environment[*].network_configuration[*].network_mode_config", "environment[*].agentcoreRuntimeEnvironment[*].networkConfiguration[*].networkModeConfig")
		r.AddSingletonListConversion("environment_artifact", "environmentArtifact")
		r.AddSingletonListConversion("environment_artifact[*].container_configuration", "environmentArtifact[*].containerConfiguration")
		r.AddSingletonListConversion("memory", "memory")
		r.AddSingletonListConversion("memory[*].agentcore_memory_configuration", "memory[*].agentcoreMemoryConfiguration")
		r.AddSingletonListConversion("memory[*].agentcore_memory_configuration[*].retrieval_config", "memory[*].agentcoreMemoryConfiguration[*].retrievalConfig")
		r.AddSingletonListConversion("memory[*].disabled", "memory[*].disabled")
		r.AddSingletonListConversion("memory[*].managed_memory_configuration", "memory[*].managedMemoryConfiguration")
		r.AddSingletonListConversion("model", "model")
		r.AddSingletonListConversion("model[*].bedrock_model_config", "model[*].bedrockModelConfig")
		r.AddSingletonListConversion("model[*].gemini_model_config", "model[*].geminiModelConfig")
		r.AddSingletonListConversion("model[*].litellm_model_config", "model[*].litellmModelConfig")
		r.AddSingletonListConversion("model[*].openai_model_config", "model[*].openaiModelConfig")
		r.AddSingletonListConversion("skill[*].aws_skills", "skill[*].awsSkills")
		r.AddSingletonListConversion("skill[*].git", "skill[*].git")
		r.AddSingletonListConversion("skill[*].git[*].auth", "skill[*].git[*].auth")
		r.AddSingletonListConversion("skill[*].s3", "skill[*].s3")
		r.AddSingletonListConversion("tool[*].config", "tool[*].config")
		r.AddSingletonListConversion("tool[*].config[*].agentcore_browser", "tool[*].config[*].agentcoreBrowser")
		r.AddSingletonListConversion("tool[*].config[*].agentcore_code_interpreter", "tool[*].config[*].agentcoreCodeInterpreter")
		r.AddSingletonListConversion("tool[*].config[*].agentcore_gateway", "tool[*].config[*].agentcoreGateway")
		r.AddSingletonListConversion("tool[*].config[*].agentcore_gateway[*].outbound_auth", "tool[*].config[*].agentcoreGateway[*].outboundAuth")
		r.AddSingletonListConversion("tool[*].config[*].agentcore_gateway[*].outbound_auth[*].oauth", "tool[*].config[*].agentcoreGateway[*].outboundAuth[*].oauth")
		r.AddSingletonListConversion("tool[*].config[*].inline_function", "tool[*].config[*].inlineFunction")
		r.AddSingletonListConversion("tool[*].config[*].remote_mcp", "tool[*].config[*].remoteMcp")
		r.AddSingletonListConversion("truncation", "truncation")
	})
	// aws_bedrockagentcore_memory
	p.AddResourceConfigurator("aws_bedrockagentcore_memory", func(r *config.Resource) {
		r.References["stream_delivery_resources.resource.kinesis.data_stream_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("stream_delivery_resources", "streamDeliveryResources")
		r.AddSingletonListConversion("stream_delivery_resources[*].resource", "streamDeliveryResources[*].resource")
		r.AddSingletonListConversion("stream_delivery_resources[*].resource[*].kinesis", "streamDeliveryResources[*].resource[*].kinesis")
		r.AddSingletonListConversion("stream_delivery_resources[*].resource[*].kinesis[*].content_configuration", "streamDeliveryResources[*].resource[*].kinesis[*].contentConfiguration")
	})
	// aws_bedrockagentcore_memory_strategy
	p.AddResourceConfigurator("aws_bedrockagentcore_memory_strategy", func(r *config.Resource) {
		r.AddSingletonListConversion("configuration", "configuration")
		r.AddSingletonListConversion("configuration[*].consolidation", "configuration[*].consolidation")
		r.AddSingletonListConversion("configuration[*].extraction", "configuration[*].extraction")
		r.AddSingletonListConversion("configuration[*].reflection", "configuration[*].reflection")
		r.AddSingletonListConversion("configuration[*].self_managed_configuration", "configuration[*].selfManagedConfiguration")
		r.AddSingletonListConversion("configuration[*].self_managed_configuration[*].invocation_configuration", "configuration[*].selfManagedConfiguration[*].invocationConfiguration")
		r.AddSingletonListConversion("configuration[*].self_managed_configuration[*].trigger_conditions", "configuration[*].selfManagedConfiguration[*].triggerConditions")
		r.AddSingletonListConversion("configuration[*].self_managed_configuration[*].trigger_conditions[*].message_based_trigger", "configuration[*].selfManagedConfiguration[*].triggerConditions[*].messageBasedTrigger")
		r.AddSingletonListConversion("configuration[*].self_managed_configuration[*].trigger_conditions[*].time_based_trigger", "configuration[*].selfManagedConfiguration[*].triggerConditions[*].timeBasedTrigger")
		r.AddSingletonListConversion("configuration[*].self_managed_configuration[*].trigger_conditions[*].token_based_trigger", "configuration[*].selfManagedConfiguration[*].triggerConditions[*].tokenBasedTrigger")
		r.AddSingletonListConversion("memory_record_schema", "memoryRecordSchema")
		r.AddSingletonListConversion("memory_record_schema[*].metadata_schema[*].extraction_config", "memoryRecordSchema[*].metadataSchema[*].extractionConfig")
		r.AddSingletonListConversion("memory_record_schema[*].metadata_schema[*].extraction_config[*].llm_extraction_config", "memoryRecordSchema[*].metadataSchema[*].extractionConfig[*].llmExtractionConfig")
		r.AddSingletonListConversion("memory_record_schema[*].metadata_schema[*].extraction_config[*].llm_extraction_config[*].validation", "memoryRecordSchema[*].metadataSchema[*].extractionConfig[*].llmExtractionConfig[*].validation")
		r.AddSingletonListConversion("memory_record_schema[*].metadata_schema[*].extraction_config[*].llm_extraction_config[*].validation[*].number_validation", "memoryRecordSchema[*].metadataSchema[*].extractionConfig[*].llmExtractionConfig[*].validation[*].numberValidation")
		r.AddSingletonListConversion("memory_record_schema[*].metadata_schema[*].extraction_config[*].llm_extraction_config[*].validation[*].string_list_validation", "memoryRecordSchema[*].metadataSchema[*].extractionConfig[*].llmExtractionConfig[*].validation[*].stringListValidation")
		r.AddSingletonListConversion("memory_record_schema[*].metadata_schema[*].extraction_config[*].llm_extraction_config[*].validation[*].string_validation", "memoryRecordSchema[*].metadataSchema[*].extractionConfig[*].llmExtractionConfig[*].validation[*].stringValidation")
		r.AddSingletonListConversion("reflection_configuration", "reflectionConfiguration")
	})
	// aws_bedrockagentcore_oauth2_credential_provider
	p.AddResourceConfigurator("aws_bedrockagentcore_oauth2_credential_provider", func(r *config.Resource) {
		r.TerraformResource.Schema["name"].Computed = true
		r.TerraformResource.Schema["name"].Optional = false

		r.AddSingletonListConversion("oauth2_provider_config", "oauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].custom_oauth2_provider_config", "oauth2ProviderConfig[*].customOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].custom_oauth2_provider_config[*].client_secret_config", "oauth2ProviderConfig[*].customOauth2ProviderConfig[*].clientSecretConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].custom_oauth2_provider_config[*].oauth_discovery", "oauth2ProviderConfig[*].customOauth2ProviderConfig[*].oauthDiscovery")
		r.AddSingletonListConversion("oauth2_provider_config[*].custom_oauth2_provider_config[*].oauth_discovery[*].authorization_server_metadata", "oauth2ProviderConfig[*].customOauth2ProviderConfig[*].oauthDiscovery[*].authorizationServerMetadata")
		r.AddSingletonListConversion("oauth2_provider_config[*].github_oauth2_provider_config", "oauth2ProviderConfig[*].githubOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].github_oauth2_provider_config[*].client_secret_config", "oauth2ProviderConfig[*].githubOauth2ProviderConfig[*].clientSecretConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].google_oauth2_provider_config", "oauth2ProviderConfig[*].googleOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].google_oauth2_provider_config[*].client_secret_config", "oauth2ProviderConfig[*].googleOauth2ProviderConfig[*].clientSecretConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].microsoft_oauth2_provider_config", "oauth2ProviderConfig[*].microsoftOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].microsoft_oauth2_provider_config[*].client_secret_config", "oauth2ProviderConfig[*].microsoftOauth2ProviderConfig[*].clientSecretConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].salesforce_oauth2_provider_config", "oauth2ProviderConfig[*].salesforceOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].salesforce_oauth2_provider_config[*].client_secret_config", "oauth2ProviderConfig[*].salesforceOauth2ProviderConfig[*].clientSecretConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].slack_oauth2_provider_config", "oauth2ProviderConfig[*].slackOauth2ProviderConfig")
		r.AddSingletonListConversion("oauth2_provider_config[*].slack_oauth2_provider_config[*].client_secret_config", "oauth2ProviderConfig[*].slackOauth2ProviderConfig[*].clientSecretConfig")
	})
	// aws_bedrockagentcore_online_evaluation_config
	p.AddResourceConfigurator("aws_bedrockagentcore_online_evaluation_config", func(r *config.Resource) {
		r.References["evaluator.evaluator_id"] = config.Reference{
			TerraformName: "aws_bedrockagentcore_evaluator",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("evaluator_id",true)`,
		}
		r.AddSingletonListConversion("data_source_config", "dataSourceConfig")
		r.AddSingletonListConversion("data_source_config[*].cloudwatch_logs", "dataSourceConfig[*].cloudwatchLogs")
		r.AddSingletonListConversion("rule", "rule")
		r.AddSingletonListConversion("rule[*].session_config", "rule[*].sessionConfig")
		r.AddSingletonListConversion("rule[*].sampling_config", "rule[*].samplingConfig")
		r.AddSingletonListConversion("rule[*].filter[*].value", "rule[*].filter[*].value")
	})
	// aws_bedrockagentcore_policy
	p.AddResourceConfigurator("aws_bedrockagentcore_policy", func(r *config.Resource) {
		r.AddSingletonListConversion("definition", "definition")
		r.AddSingletonListConversion("definition[*].cedar", "definition[*].cedar")
	})
	// aws_bedrockagentcore_resource_policy
	p.AddResourceConfigurator("aws_bedrockagentcore_resource_policy", func(r *config.Resource) {
		delete(r.References, "resource_arn")
	})
	// aws_bedrockagentcore_token_vault_cmk
	p.AddResourceConfigurator("aws_bedrockagentcore_token_vault_cmk", func(r *config.Resource) {
		r.AddSingletonListConversion("kms_configuration", "kmsConfiguration")
	})
	// aws_bedrockagentcore_workload_identity
	p.AddResourceConfigurator("aws_bedrockagentcore_workload_identity", func(r *config.Resource) {
		r.TerraformResource.Schema["name"].Computed = true
		r.TerraformResource.Schema["name"].Optional = false
	})
}
