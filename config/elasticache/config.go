/*
Copyright 2021 Upbound Inc.
*/

package elasticache

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure adds configurations for the elasticache group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elasticache_cluster", func(r *config.Resource) {
		r.References["parameter_group_name"] = config.Reference{
			TerraformName: "aws_elasticache_parameter_group",
		}
		// log_delivery_configuration.destination can point to either
		// a CloudWatch Logs LogGroup or Kinesis Data Firehose resource.
		delete(r.References, "log_delivery_configuration.destination")
	})

	p.AddResourceConfigurator("aws_elasticache_replication_group", func(r *config.Resource) {
		r.References["subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		r.References["kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.LateInitializer = config.LateInitializer{
			// Conflicting configuration arguments: "number_cache_clusters": conflicts with cluster_mode.0.num_node_groups
			IgnoredFields: []string{
				"cluster_mode",
				"num_node_groups",
				"num_cache_clusters",
				"number_cache_clusters",
				"replication_group_description",
				"description",
			},
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "security_group_names.#")
			}
			return diff, nil
		}
		delete(r.References, "log_delivery_configuration.destination")
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_elasticache_user_group", func(r *config.Resource) {
		r.References["user_ids"] = config.Reference{
			Type:              "User",
			RefFieldName:      "UserIDRefs",
			SelectorFieldName: "UserIDSelector",
		}
	})
}
