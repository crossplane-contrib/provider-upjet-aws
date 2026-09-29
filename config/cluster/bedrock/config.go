// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bedrock

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_bedrock_custom_model", func(r *config.Resource) {
		r.References["custom_model_kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["vpc_config.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
		r.References["vpc_config.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.AddSingletonListConversion("output_data_config", "outputDataConfig")
		r.AddSingletonListConversion("training_data_config", "trainingDataConfig")
		r.AddSingletonListConversion("training_metrics", "trainingMetrics")
		r.AddSingletonListConversion("validation_data_config", "validationDataConfig")
		r.AddSingletonListConversion("vpc_config", "vpcConfig")
	})

	p.AddResourceConfigurator("aws_bedrock_evaluation_job", func(r *config.Resource) {
		r.References["customer_encryption_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["inference_config.rag_config.knowledge_base_config.retrieve_and_generate_config.knowledge_base_id"] = config.Reference{
			TerraformName: "aws_bedrockagent_knowledge_base",
		}
		r.References["inference_config.rag_config.knowledge_base_config.retrieve_config.knowledge_base_id"] = config.Reference{
			TerraformName: "aws_bedrockagent_knowledge_base",
		}
		r.References["role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		// bedrock_evaluator_model and rag_config are limited to a single element
		// by the Terraform provider, but are arrays in the AWS API, so they are
		// kept as lists.
		r.AddSingletonListConversion("evaluation_config", "evaluationConfig")
		r.AddSingletonListConversion("evaluation_config[*].automated", "evaluationConfig[*].automated")
		r.AddSingletonListConversion("evaluation_config[*].automated[*].custom_metric_config", "evaluationConfig[*].automated[*].customMetricConfig")
		r.AddSingletonListConversion("evaluation_config[*].automated[*].custom_metric_config[*].custom_metric[*].custom_metric_definition", "evaluationConfig[*].automated[*].customMetricConfig[*].customMetric[*].customMetricDefinition")
		r.AddSingletonListConversion("evaluation_config[*].automated[*].custom_metric_config[*].custom_metric[*].custom_metric_definition[*].rating_scale[*].value", "evaluationConfig[*].automated[*].customMetricConfig[*].customMetric[*].customMetricDefinition[*].ratingScale[*].value")
		r.AddSingletonListConversion("evaluation_config[*].automated[*].custom_metric_config[*].evaluator_model_config", "evaluationConfig[*].automated[*].customMetricConfig[*].evaluatorModelConfig")
		r.AddSingletonListConversion("evaluation_config[*].automated[*].dataset_metric_config[*].dataset", "evaluationConfig[*].automated[*].datasetMetricConfig[*].dataset")
		r.AddSingletonListConversion("evaluation_config[*].automated[*].dataset_metric_config[*].dataset[*].dataset_location", "evaluationConfig[*].automated[*].datasetMetricConfig[*].dataset[*].datasetLocation")
		r.AddSingletonListConversion("evaluation_config[*].automated[*].evaluator_model_config", "evaluationConfig[*].automated[*].evaluatorModelConfig")
		r.AddSingletonListConversion("evaluation_config[*].human", "evaluationConfig[*].human")
		r.AddSingletonListConversion("evaluation_config[*].human[*].dataset_metric_config[*].dataset", "evaluationConfig[*].human[*].datasetMetricConfig[*].dataset")
		r.AddSingletonListConversion("evaluation_config[*].human[*].dataset_metric_config[*].dataset[*].dataset_location", "evaluationConfig[*].human[*].datasetMetricConfig[*].dataset[*].datasetLocation")
		r.AddSingletonListConversion("evaluation_config[*].human[*].human_workflow_config", "evaluationConfig[*].human[*].humanWorkflowConfig")
		r.AddSingletonListConversion("inference_config", "inferenceConfig")
		r.AddSingletonListConversion("inference_config[*].model[*].bedrock_model", "inferenceConfig[*].model[*].bedrockModel")
		r.AddSingletonListConversion("inference_config[*].model[*].bedrock_model[*].performance_config", "inferenceConfig[*].model[*].bedrockModel[*].performanceConfig")
		r.AddSingletonListConversion("inference_config[*].model[*].precomputed_inference_source", "inferenceConfig[*].model[*].precomputedInferenceSource")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].knowledge_base_config", "inferenceConfig[*].ragConfig[*].knowledgeBaseConfig")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].knowledge_base_config[*].retrieve_and_generate_config", "inferenceConfig[*].ragConfig[*].knowledgeBaseConfig[*].retrieveAndGenerateConfig")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].knowledge_base_config[*].retrieve_and_generate_config[*].retrieval_configuration", "inferenceConfig[*].ragConfig[*].knowledgeBaseConfig[*].retrieveAndGenerateConfig[*].retrievalConfiguration")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].knowledge_base_config[*].retrieve_and_generate_config[*].retrieval_configuration[*].vector_search_configuration", "inferenceConfig[*].ragConfig[*].knowledgeBaseConfig[*].retrieveAndGenerateConfig[*].retrievalConfiguration[*].vectorSearchConfiguration")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].knowledge_base_config[*].retrieve_config", "inferenceConfig[*].ragConfig[*].knowledgeBaseConfig[*].retrieveConfig")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].knowledge_base_config[*].retrieve_config[*].knowledge_base_retrieval_configuration", "inferenceConfig[*].ragConfig[*].knowledgeBaseConfig[*].retrieveConfig[*].knowledgeBaseRetrievalConfiguration")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].knowledge_base_config[*].retrieve_config[*].knowledge_base_retrieval_configuration[*].vector_search_configuration", "inferenceConfig[*].ragConfig[*].knowledgeBaseConfig[*].retrieveConfig[*].knowledgeBaseRetrievalConfiguration[*].vectorSearchConfiguration")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].precomputed_rag_source_config", "inferenceConfig[*].ragConfig[*].precomputedRagSourceConfig")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].precomputed_rag_source_config[*].retrieve_and_generate_source_config", "inferenceConfig[*].ragConfig[*].precomputedRagSourceConfig[*].retrieveAndGenerateSourceConfig")
		r.AddSingletonListConversion("inference_config[*].rag_config[*].precomputed_rag_source_config[*].retrieve_source_config", "inferenceConfig[*].ragConfig[*].precomputedRagSourceConfig[*].retrieveSourceConfig")
		r.AddSingletonListConversion("output_data_config", "outputDataConfig")
	})

	p.AddResourceConfigurator("aws_bedrock_guardrail", func(r *config.Resource) {
		r.References["kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("content_policy_config", "contentPolicyConfig")
		r.AddSingletonListConversion("content_policy_config[*].tier_config", "contentPolicyConfig[*].tierConfig")
		r.AddSingletonListConversion("contextual_grounding_policy_config", "contextualGroundingPolicyConfig")
		r.AddSingletonListConversion("cross_region_config", "crossRegionConfig")
		r.AddSingletonListConversion("sensitive_information_policy_config", "sensitiveInformationPolicyConfig")
		r.AddSingletonListConversion("topic_policy_config", "topicPolicyConfig")
		r.AddSingletonListConversion("topic_policy_config[*].tier_config", "topicPolicyConfig[*].tierConfig")
		r.AddSingletonListConversion("word_policy_config", "wordPolicyConfig")
	})

	p.AddResourceConfigurator("aws_bedrock_guardrail_version", func(r *config.Resource) {
		r.References["guardrail_arn"] = config.Reference{
			TerraformName: "aws_bedrock_guardrail",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("guardrail_arn",true)`,
		}
	})

	p.AddResourceConfigurator("aws_bedrock_model_invocation_job", func(r *config.Resource) {
		r.References["output_data_config.s3_output_data_config.s3_encryption_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["vpc_config.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
		r.References["vpc_config.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.AddSingletonListConversion("input_data_config", "inputDataConfig")
		r.AddSingletonListConversion("input_data_config[*].s3_input_data_config", "inputDataConfig[*].s3InputDataConfig")
		r.AddSingletonListConversion("output_data_config", "outputDataConfig")
		r.AddSingletonListConversion("output_data_config[*].s3_output_data_config", "outputDataConfig[*].s3OutputDataConfig")
		r.AddSingletonListConversion("vpc_config", "vpcConfig")
	})

	p.AddResourceConfigurator("aws_bedrock_model_invocation_logging_configuration", func(r *config.Resource) {
		r.References["logging_config.cloudwatch_config.large_data_delivery_s3_config.bucket_name"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
		r.References["logging_config.cloudwatch_config.log_group_name"] = config.Reference{
			TerraformName: "aws_cloudwatch_log_group",
		}
		r.References["logging_config.cloudwatch_config.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["logging_config.s3_config.bucket_name"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
		r.AddSingletonListConversion("logging_config", "loggingConfig")
		r.AddSingletonListConversion("logging_config[*].cloudwatch_config", "loggingConfig[*].cloudwatchConfig")
		r.AddSingletonListConversion("logging_config[*].cloudwatch_config[*].large_data_delivery_s3_config", "loggingConfig[*].cloudwatchConfig[*].largeDataDeliveryS3Config")
		r.AddSingletonListConversion("logging_config[*].s3_config", "loggingConfig[*].s3Config")
	})
}
