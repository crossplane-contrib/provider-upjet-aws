// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package iot

import (
	"fmt"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the iot group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_iot_topic_rule_destination", func(r *config.Resource) {
		r.References["vpc_configuration.security_groups"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["vpc_configuration.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
	})
	// Almost everything in this resource's schema is present twice, once at top level, and once under error_action
	p.AddResourceConfigurator("aws_iot_topic_rule", func(r *config.Resource) {
		var newReferences config.References = map[string]config.Reference{}
		newReferences["cloudwatch_alarm.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["cloudwatch_logs.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["cloudwatch_logs.log_group_name"] = config.Reference{
			TerraformName: "aws_cloudwatch_log_group",
			Extractor:     common.PathARNExtractor,
			// TODO: validate that this is the right extractor
		}
		newReferences["cloudwatch_metric.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["dynamodb.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["dynamodb.table_name"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
			Extractor:     common.PathARNExtractor,
			// TODO: validate that this is the right extractor
		}
		newReferences["dynamodbv2.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["dynamodbv2.put_item.table_name"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
			Extractor:     common.PathARNExtractor,
			// TODO: validate that this is the right extractor
		}
		newReferences["elasticsearch.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["firehose.delivery_stream_name"] = config.Reference{
			TerraformName: "aws_kinesis_firehose_delivery_stream",
			Extractor:     common.PathARNExtractor,
			// TODO: validate that this is the right extractor
		}
		newReferences["firehose.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["iot_analytics.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["iot_events.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["kafka.destination_arn"] = config.Reference{
			TerraformName: "aws_iot_topic_rule_destination",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["kinesis.stream_name"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathARNExtractor,
			// TODO: validate that this is the right extractor
		}
		newReferences["kinesis.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["lambda.function_arn"] = config.Reference{
			TerraformName: "aws_lambda_function",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["republish.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["s3.bucket_name"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     common.PathARNExtractor,
			// TODO: does this work?
		}
		newReferences["s3.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["sns.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["sns.target_arn"] = config.Reference{
			TerraformName: "aws_sns_topic",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["sqs.queue_url"] = config.Reference{
			TerraformName: "aws_sqs_queue",
		}
		newReferences["sqs.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["step_functions.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		newReferences["step_functions.state_machine_name"] = config.Reference{
			TerraformName: "aws_sfn_state_machine",
			Extractor:     common.PathARNExtractor,
			// TODO: or external name?
		}
		newReferences["timestream.database_name"] = config.Reference{
			TerraformName: "aws_timestreamwrite_database",
			// Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)`,
			// Extractor:     common.PathARNExtractor,
		}
		newReferences["timestream.table_name"] = config.Reference{
			TerraformName: "aws_timestreamwrite_table",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("table_name",false)`,
			// Extractor:     common.PathARNExtractor,
		}
		newReferences["timestream.role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		for k, v := range newReferences {
			r.References[k] = v
			r.References[fmt.Sprint("error_action.", k)] = v
		}
	})
}
