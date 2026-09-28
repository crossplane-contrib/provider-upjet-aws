// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bedrockagent

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_bedrockagent_agent", func(r *config.Resource) {
		r.References["customer_encryption_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["agent_resource_role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_bedrockagent_agent_knowledge_base_association", func(r *config.Resource) {
		r.References["agent_id"] = config.Reference{
			TerraformName: "aws_bedrockagent_agent",
		}
		r.References["knowledge_base_id"] = config.Reference{
			TerraformName: "aws_bedrockagent_knowledge_base",
		}
	})

	p.AddResourceConfigurator("aws_bedrockagent_data_source", func(r *config.Resource) {
		r.References["knowledge_base_id"] = config.Reference{
			TerraformName: "aws_bedrockagent_knowledge_base",
		}
		r.References["data_source_configuration.s3_configuration.bucket_arn"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     common.PathARNExtractor,
		}
		r.References["data_source_configuration.confluence_configuration.source_configuration.credentials_secret_arn"] = config.Reference{
			TerraformName: "aws_secretsmanager_secret",
			Extractor:     common.PathARNExtractor,
		}
		r.References["data_source_configuration.salesforce_configuration.source_configuration.credentials_secret_arn"] = config.Reference{
			TerraformName: "aws_secretsmanager_secret",
			Extractor:     common.PathARNExtractor,
		}
		r.References["data_source_configuration.share_point_configuration.source_configuration.credentials_secret_arn"] = config.Reference{
			TerraformName: "aws_secretsmanager_secret",
			Extractor:     common.PathARNExtractor,
		}
		r.References["server_side_encryption_configuration.kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["vector_ingestion_configuration.custom_transformation_configuration.transformation.transformation_function.transformation_lambda_configuration.lambda_arn"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		r.AddSingletonListConversion("data_source_configuration", "dataSourceConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].confluence_configuration", "dataSourceConfiguration[*].confluenceConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].confluence_configuration[*].crawler_configuration", "dataSourceConfiguration[*].confluenceConfiguration[*].crawlerConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].confluence_configuration[*].crawler_configuration[*].filter_configuration", "dataSourceConfiguration[*].confluenceConfiguration[*].crawlerConfiguration[*].filterConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].confluence_configuration[*].source_configuration", "dataSourceConfiguration[*].confluenceConfiguration[*].sourceConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].managed_knowledge_base_connector_configuration", "dataSourceConfiguration[*].managedKnowledgeBaseConnectorConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].managed_knowledge_base_connector_configuration[*].deletion_protection_configuration", "dataSourceConfiguration[*].managedKnowledgeBaseConnectorConfiguration[*].deletionProtectionConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].managed_knowledge_base_connector_configuration[*].media_extraction_configuration", "dataSourceConfiguration[*].managedKnowledgeBaseConnectorConfiguration[*].mediaExtractionConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].managed_knowledge_base_connector_configuration[*].media_extraction_configuration[*].audio_extraction_configuration", "dataSourceConfiguration[*].managedKnowledgeBaseConnectorConfiguration[*].mediaExtractionConfiguration[*].audioExtractionConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].managed_knowledge_base_connector_configuration[*].media_extraction_configuration[*].image_extraction_configuration", "dataSourceConfiguration[*].managedKnowledgeBaseConnectorConfiguration[*].mediaExtractionConfiguration[*].imageExtractionConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].managed_knowledge_base_connector_configuration[*].media_extraction_configuration[*].video_extraction_configuration", "dataSourceConfiguration[*].managedKnowledgeBaseConnectorConfiguration[*].mediaExtractionConfiguration[*].videoExtractionConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].s3_configuration", "dataSourceConfiguration[*].s3Configuration")
		r.AddSingletonListConversion("data_source_configuration[*].salesforce_configuration", "dataSourceConfiguration[*].salesforceConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].salesforce_configuration[*].crawler_configuration", "dataSourceConfiguration[*].salesforceConfiguration[*].crawlerConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].salesforce_configuration[*].crawler_configuration[*].filter_configuration", "dataSourceConfiguration[*].salesforceConfiguration[*].crawlerConfiguration[*].filterConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].salesforce_configuration[*].source_configuration", "dataSourceConfiguration[*].salesforceConfiguration[*].sourceConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].share_point_configuration", "dataSourceConfiguration[*].sharePointConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].share_point_configuration[*].crawler_configuration", "dataSourceConfiguration[*].sharePointConfiguration[*].crawlerConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].share_point_configuration[*].crawler_configuration[*].filter_configuration", "dataSourceConfiguration[*].sharePointConfiguration[*].crawlerConfiguration[*].filterConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].share_point_configuration[*].source_configuration", "dataSourceConfiguration[*].sharePointConfiguration[*].sourceConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].web_configuration", "dataSourceConfiguration[*].webConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].web_configuration[*].crawler_configuration", "dataSourceConfiguration[*].webConfiguration[*].crawlerConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].web_configuration[*].crawler_configuration[*].crawler_limits", "dataSourceConfiguration[*].webConfiguration[*].crawlerConfiguration[*].crawlerLimits")
		r.AddSingletonListConversion("data_source_configuration[*].web_configuration[*].source_configuration", "dataSourceConfiguration[*].webConfiguration[*].sourceConfiguration")
		r.AddSingletonListConversion("data_source_configuration[*].web_configuration[*].source_configuration[*].url_configuration", "dataSourceConfiguration[*].webConfiguration[*].sourceConfiguration[*].urlConfiguration")
		r.AddSingletonListConversion("server_side_encryption_configuration", "serverSideEncryptionConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration", "vectorIngestionConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].chunking_configuration", "vectorIngestionConfiguration[*].chunkingConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].chunking_configuration[*].fixed_size_chunking_configuration", "vectorIngestionConfiguration[*].chunkingConfiguration[*].fixedSizeChunkingConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].chunking_configuration[*].hierarchical_chunking_configuration", "vectorIngestionConfiguration[*].chunkingConfiguration[*].hierarchicalChunkingConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].chunking_configuration[*].semantic_chunking_configuration", "vectorIngestionConfiguration[*].chunkingConfiguration[*].semanticChunkingConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].custom_transformation_configuration", "vectorIngestionConfiguration[*].customTransformationConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].custom_transformation_configuration[*].intermediate_storage", "vectorIngestionConfiguration[*].customTransformationConfiguration[*].intermediateStorage")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].custom_transformation_configuration[*].intermediate_storage[*].s3_location", "vectorIngestionConfiguration[*].customTransformationConfiguration[*].intermediateStorage[*].s3Location")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].custom_transformation_configuration[*].transformation", "vectorIngestionConfiguration[*].customTransformationConfiguration[*].transformation")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].custom_transformation_configuration[*].transformation[*].transformation_function", "vectorIngestionConfiguration[*].customTransformationConfiguration[*].transformation[*].transformationFunction")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].custom_transformation_configuration[*].transformation[*].transformation_function[*].transformation_lambda_configuration", "vectorIngestionConfiguration[*].customTransformationConfiguration[*].transformation[*].transformationFunction[*].transformationLambdaConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].parsing_configuration", "vectorIngestionConfiguration[*].parsingConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].parsing_configuration[*].bedrock_data_automation_configuration", "vectorIngestionConfiguration[*].parsingConfiguration[*].bedrockDataAutomationConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].parsing_configuration[*].bedrock_foundation_model_configuration", "vectorIngestionConfiguration[*].parsingConfiguration[*].bedrockFoundationModelConfiguration")
		r.AddSingletonListConversion("vector_ingestion_configuration[*].parsing_configuration[*].bedrock_foundation_model_configuration[*].parsing_prompt", "vectorIngestionConfiguration[*].parsingConfiguration[*].bedrockFoundationModelConfiguration[*].parsingPrompt")
	})

	p.AddResourceConfigurator("aws_bedrockagent_knowledge_base", func(r *config.Resource) {
		r.References["role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["knowledge_base_configuration.kendra_knowledge_base_configuration.kendra_index_arn"] = config.Reference{
			TerraformName: "aws_kendra_index",
			Extractor:     common.PathARNExtractor,
		}
		r.References["knowledge_base_configuration.managed_knowledge_base_configuration.server_side_encryption_configuration.kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["storage_configuration.opensearch_managed_cluster_configuration.domain_arn"] = config.Reference{
			TerraformName: "aws_opensearch_domain",
			Extractor:     common.PathARNExtractor,
		}
		r.References["storage_configuration.opensearch_serverless_configuration.collection_arn"] = config.Reference{
			TerraformName: "aws_opensearchserverless_collection",
			Extractor:     common.PathARNExtractor,
		}
		r.References["storage_configuration.rds_configuration.resource_arn"] = config.Reference{
			TerraformName: "aws_rds_cluster",
			Extractor:     common.PathARNExtractor,
		}
		r.References["storage_configuration.rds_configuration.credentials_secret_arn"] = config.Reference{
			TerraformName: "aws_secretsmanager_secret",
			Extractor:     common.PathARNExtractor,
		}
		r.References["storage_configuration.s3_vectors_configuration.index_arn"] = config.Reference{
			TerraformName: "aws_s3vectors_index",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("index_arn",true)`,
		}
		r.References["storage_configuration.s3_vectors_configuration.vector_bucket_arn"] = config.Reference{
			TerraformName: "aws_s3vectors_vector_bucket",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("vector_bucket_arn",true)`,
		}
		// injected from the registry example, which interpolates plain values
		// of an aws_redshift_cluster; the database user and name are not
		// resource identities and database_name is shared with Redshift
		// Serverless query engines.
		delete(r.References, "knowledge_base_configuration.sql_knowledge_base_configuration.redshift_configuration.query_engine_configuration.provisioned_configuration.auth_configuration.database_user")
		delete(r.References, "knowledge_base_configuration.sql_knowledge_base_configuration.redshift_configuration.storage_configuration.redshift_configuration.database_name")
		r.AddSingletonListConversion("knowledge_base_configuration", "knowledgeBaseConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].kendra_knowledge_base_configuration", "knowledgeBaseConfiguration[*].kendraKnowledgeBaseConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].managed_knowledge_base_configuration", "knowledgeBaseConfiguration[*].managedKnowledgeBaseConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].managed_knowledge_base_configuration[*].embedding_model_configuration", "knowledgeBaseConfiguration[*].managedKnowledgeBaseConfiguration[*].embeddingModelConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].managed_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration", "knowledgeBaseConfiguration[*].managedKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].managed_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration[*].audio", "knowledgeBaseConfiguration[*].managedKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration[*].audio")
		r.AddSingletonListConversion("knowledge_base_configuration[*].managed_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration[*].audio[*].segmentation_configuration", "knowledgeBaseConfiguration[*].managedKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration[*].audio[*].segmentationConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].managed_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration[*].video", "knowledgeBaseConfiguration[*].managedKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration[*].video")
		r.AddSingletonListConversion("knowledge_base_configuration[*].managed_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration[*].video[*].segmentation_configuration", "knowledgeBaseConfiguration[*].managedKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration[*].video[*].segmentationConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].managed_knowledge_base_configuration[*].server_side_encryption_configuration", "knowledgeBaseConfiguration[*].managedKnowledgeBaseConfiguration[*].serverSideEncryptionConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].query_engine_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].queryEngineConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].query_engine_configuration[*].provisioned_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].queryEngineConfiguration[*].provisionedConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].query_engine_configuration[*].provisioned_configuration[*].auth_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].queryEngineConfiguration[*].provisionedConfiguration[*].authConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].query_engine_configuration[*].serverless_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].queryEngineConfiguration[*].serverlessConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].query_engine_configuration[*].serverless_configuration[*].auth_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].queryEngineConfiguration[*].serverlessConfiguration[*].authConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].query_generation_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].queryGenerationConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].query_generation_configuration[*].generation_context", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].queryGenerationConfiguration[*].generationContext")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].storage_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].storageConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].storage_configuration[*].aws_data_catalog_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].storageConfiguration[*].awsDataCatalogConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].sql_knowledge_base_configuration[*].redshift_configuration[*].storage_configuration[*].redshift_configuration", "knowledgeBaseConfiguration[*].sqlKnowledgeBaseConfiguration[*].redshiftConfiguration[*].storageConfiguration[*].redshiftConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].embedding_model_configuration", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].embeddingModelConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration[*].audio", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration[*].audio")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration[*].audio[*].segmentation_configuration", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration[*].audio[*].segmentationConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration[*].video", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration[*].video")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].embedding_model_configuration[*].bedrock_embedding_model_configuration[*].video[*].segmentation_configuration", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].embeddingModelConfiguration[*].bedrockEmbeddingModelConfiguration[*].video[*].segmentationConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].supplemental_data_storage_configuration", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].supplementalDataStorageConfiguration")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].supplemental_data_storage_configuration[*].storage_location", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].supplementalDataStorageConfiguration[*].storageLocation")
		r.AddSingletonListConversion("knowledge_base_configuration[*].vector_knowledge_base_configuration[*].supplemental_data_storage_configuration[*].storage_location[*].s3_location", "knowledgeBaseConfiguration[*].vectorKnowledgeBaseConfiguration[*].supplementalDataStorageConfiguration[*].storageLocation[*].s3Location")
		r.AddSingletonListConversion("storage_configuration", "storageConfiguration")
		r.AddSingletonListConversion("storage_configuration[*].mongo_db_atlas_configuration", "storageConfiguration[*].mongoDbAtlasConfiguration")
		r.AddSingletonListConversion("storage_configuration[*].mongo_db_atlas_configuration[*].field_mapping", "storageConfiguration[*].mongoDbAtlasConfiguration[*].fieldMapping")
		r.AddSingletonListConversion("storage_configuration[*].neptune_analytics_configuration", "storageConfiguration[*].neptuneAnalyticsConfiguration")
		r.AddSingletonListConversion("storage_configuration[*].neptune_analytics_configuration[*].field_mapping", "storageConfiguration[*].neptuneAnalyticsConfiguration[*].fieldMapping")
		r.AddSingletonListConversion("storage_configuration[*].opensearch_managed_cluster_configuration", "storageConfiguration[*].opensearchManagedClusterConfiguration")
		r.AddSingletonListConversion("storage_configuration[*].opensearch_managed_cluster_configuration[*].field_mapping", "storageConfiguration[*].opensearchManagedClusterConfiguration[*].fieldMapping")
		r.AddSingletonListConversion("storage_configuration[*].opensearch_serverless_configuration", "storageConfiguration[*].opensearchServerlessConfiguration")
		r.AddSingletonListConversion("storage_configuration[*].opensearch_serverless_configuration[*].field_mapping", "storageConfiguration[*].opensearchServerlessConfiguration[*].fieldMapping")
		r.AddSingletonListConversion("storage_configuration[*].pinecone_configuration", "storageConfiguration[*].pineconeConfiguration")
		r.AddSingletonListConversion("storage_configuration[*].pinecone_configuration[*].field_mapping", "storageConfiguration[*].pineconeConfiguration[*].fieldMapping")
		r.AddSingletonListConversion("storage_configuration[*].rds_configuration", "storageConfiguration[*].rdsConfiguration")
		r.AddSingletonListConversion("storage_configuration[*].rds_configuration[*].field_mapping", "storageConfiguration[*].rdsConfiguration[*].fieldMapping")
		r.AddSingletonListConversion("storage_configuration[*].redis_enterprise_cloud_configuration", "storageConfiguration[*].redisEnterpriseCloudConfiguration")
		r.AddSingletonListConversion("storage_configuration[*].redis_enterprise_cloud_configuration[*].field_mapping", "storageConfiguration[*].redisEnterpriseCloudConfiguration[*].fieldMapping")
		r.AddSingletonListConversion("storage_configuration[*].s3_vectors_configuration", "storageConfiguration[*].s3VectorsConfiguration")
	})
}
