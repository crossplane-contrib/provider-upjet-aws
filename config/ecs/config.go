/*
Copyright 2021 Upbound Inc.
*/

package ecs

import (
	"context"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the ecs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecs_cluster", func(r *config.Resource) {
		// Mutually exclusive with aws_ecs_cluster_capacity_providers
		config.MoveToStatus(r.TerraformResource, "capacity_providers")

		r.References = config.References{
			"execute_command_configuration.kms_key_id": config.Reference{
				Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			},
			"log_configuration.s3_bucket_name": config.Reference{
				Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
			},
		}

		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_service", func(r *config.Resource) {
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			cl, ok := parameters["cluster"].(string)
			if !ok {
				return "", errors.New("cannot generate id without cluster paramater")
			}
			return filepath.Join(cl, externalName), nil
		}
		r.References = config.References{
			"cluster": config.Reference{
				Type: "Cluster",
			},
			"task_definition": config.Reference{
				Type:      "TaskDefinition",
				Extractor: common.PathARNExtractor,
			},
			"iam_role": config.Reference{
				Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
			"load_balancer.target_group_arn": config.Reference{
				Type: "github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup",
			},
			"network_configuration.subnets": config.Reference{
				Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"network_configuration.security_groups": config.Reference{
				Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
		}
		r.MetaResource.ArgumentDocs["cluster"] = `Name of an ECS cluster.`
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_capacity_provider", func(r *config.Resource) {
		r.References = config.References{
			"auto_scaling_group_provider.auto_scaling_group_arn": config.Reference{
				Type:      "github.com/upbound/provider-aws/apis/autoscaling/v1beta1.AutoscalingGroup",
				Extractor: common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_ecs_task_definition", func(r *config.Resource) {
		r.References = config.References{
			"execution_role_arn": config.Reference{
				Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
