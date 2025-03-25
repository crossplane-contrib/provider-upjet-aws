// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kafka

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the kafka group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_msk_cluster", func(r *config.Resource) {
		r.References["encryption_info.encryption_at_rest_kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["logging_info.broker_logs.s3.bucket"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
		r.References["logging_info.broker_logs.cloudwatch_logs.log_group"] = config.Reference{
			TerraformName: "aws_cloudwatch_log_group",
		}
		r.References["broker_node_group_info.client_subnets"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["broker_node_group_info.security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["configuration_info.arn"] = config.Reference{
			TerraformName: "aws_msk_configuration",
			Extractor:     common.PathARNExtractor,
		}
		r.UseAsync = true

		r.Version = "v1beta2"
	})
	p.AddResourceConfigurator("aws_msk_scram_secret_association", func(r *config.Resource) {
		r.References["secret_arn_list"] = config.Reference{
			TerraformName:     "aws_secretsmanager_secret",
			RefFieldName:      "SecretArnRefs",
			SelectorFieldName: "SecretArnSelector",
		}
		r.MetaResource.ArgumentDocs["secret_arn_list"] = "- (Required) List of all AWS Secrets Manager secret ARNs to associate with the cluster. Secrets not referenced, selected or listed here will be disassociated from the cluster."
	})
	p.AddResourceConfigurator("aws_msk_serverless_cluster", func(r *config.Resource) {
		r.UseAsync = true
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
		r.OverrideFieldNames = map[string]string{
			"ClientAuthenticationParameters":     "ServerlessClusterClientAuthenticationParameters",
			"ClientAuthenticationInitParameters": "ServerlessClusterClientAuthenticationInitParameters",
			"ClientAuthenticationObservation":    "ServerlessClusterClientAuthenticationObservation",
			"SaslParameters":                     "ClientAuthenticationSaslParameters",
			"SaslInitParameters":                 "ClientAuthenticationSaslInitParameters",
			"SaslObservation":                    "ClientAuthenticationSaslObservation",
		}
	})
}
