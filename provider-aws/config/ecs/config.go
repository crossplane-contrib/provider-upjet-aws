/*
Copyright 2021 Upbound Inc.
*/

package ecs

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for ecs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecs_cluster", func(r *config.Resource) {
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			// expected id format: arn:aws:ecs:us-west-2:123456789123:cluster/example-cluster
			w := strings.Split(tfstate["id"].(string), "/")
			if len(w) != 2 {
				return "", errors.New("terraform ID should be the ARN of the cluster")
			}
			return w[len(w)-1], nil
		}
		r.References = config.References{
			"capacity_providers": config.Reference{
				Type: "CapacityProvider",
			},
			"execute_command_configuration.kms_key_id": config.Reference{
				Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
			},
			"log_configuration.s3_bucket_name": config.Reference{
				Type: "github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_service", func(r *config.Resource) {
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			// expected id format: arn:aws:ecs:us-east-2:123456789123:service/sample-cluster/sample-service
			w := strings.Split(tfstate["id"].(string), "/")
			if len(w) != 3 {
				return "", errors.New("terraform ID should be the ARN of the service")
			}
			return w[len(w)-1], nil
		}
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			cl, ok := parameters["cluster"].(string)
			if !ok {
				return "", errors.New("cannot generate id without cluster paramater")
			}
			return filepath.Join(cl, externalName), nil
		}
		r.References = config.References{
			"cluster": config.Reference{
				Type:      "Cluster",
				Extractor: common.PathARNExtractor,
			},
			"iam_role": config.Reference{
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
			"network_configuration.subnets": config.Reference{
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"network_configuration.security_groups": config.Reference{
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_ecs_capacity_provider", func(r *config.Resource) {
		r.References = config.References{
			"auto_scaling_group_provider.auto_scaling_group_arn": config.Reference{
				Type:      "github.com/upbound/official-providers/provider-aws/apis/autoscaling/v1beta1.AutoscalingGroup",
				Extractor: common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_ecs_task_definition", func(r *config.Resource) {
		r.References = config.References{
			"execution_role_arn": config.Reference{
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
