// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ecs

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

// Configure adds configurations for the ecs group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_ecs_cluster", func(r *config.Resource) {
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			// expected id format: arn:aws:ecs:us-west-2:123456789123:cluster/example-cluster
			w := strings.Split(tfstate["id"].(string), "/")
			if len(w) != 2 {
				return "", errors.New("terraform ID should be the ARN of the cluster")
			}
			return w[len(w)-1], nil
		}

		// Mutually exclusive with aws_ecs_cluster_capacity_providers
		config.MoveToStatus(r.TerraformResource, "capacity_providers")

		r.References = config.References{
			"execute_command_configuration.kms_key_id": config.Reference{
				TerraformName: "aws_kms_key",
			},
			"log_configuration.s3_bucket_name": config.Reference{
				TerraformName: "aws_s3_bucket",
			},
		}

		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_service", func(r *config.Resource) {
		r.References = config.References{
			"cluster": config.Reference{
				TerraformName: "aws_ecs_cluster",
			},
			"task_definition": config.Reference{
				TerraformName: "aws_ecs_task_definition",
				Extractor:     common.PathARNExtractor,
			},
			"iam_role": config.Reference{
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
			"load_balancer.target_group_arn": config.Reference{
				TerraformName: "aws_lb_target_group",
			},
			"network_configuration.subnets": config.Reference{
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"network_configuration.security_groups": config.Reference{
				TerraformName:     "aws_security_group",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
		}
		r.MetaResource.ArgumentDocs["cluster"] = `Name of an ECS cluster.`
		r.UseAsync = true

		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			td, ok := diff.Attributes["task_definition"]
			if !ok {
				return diff, nil
			}
			if td.Old == td.New {
				return diff, nil
			}
			// example td.New = arn:aws:ecs:us-west-1:<account_id>:task-definition/sampleservice:36
			tdParts := strings.Split(td.New, "/")
			if td.Old == tdParts[1] {
				delete(diff.Attributes, "task_definition")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("aws_ecs_capacity_provider", func(r *config.Resource) {
		r.References = config.References{
			"auto_scaling_group_provider.auto_scaling_group_arn": config.Reference{
				TerraformName: "aws_autoscaling_group",
				Extractor:     common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_ecs_task_definition", func(r *config.Resource) {
		r.References = config.References{
			"execution_role_arn": config.Reference{
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
		}
	})
}
