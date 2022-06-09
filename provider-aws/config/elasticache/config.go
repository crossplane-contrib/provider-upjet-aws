/*
Copyright 2021 Upbound Inc.
*/

package elasticache

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for elasticache group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elasticache_cluster", func(r *config.Resource) {
		r.References = config.References{
			"parameter_group_name": config.Reference{
				Type: "ParameterGroup",
			},
			"subnet_group_name": config.Reference{
				Type: "SubnetGroup",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_elasticache_replication_group", func(r *config.Resource) {
		r.References = config.References{
			"subnet_group_name": config.Reference{
				Type: "SubnetGroup",
			},
			"security_group_ids": config.Reference{
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupIdRefs",
				SelectorFieldName: "SecurityGroupIdSelector",
			},
			"kms_key_id": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
			},
		}
		r.LateInitializer = config.LateInitializer{
			// Conflicting configuration arguments: "number_cache_clusters": conflicts with cluster_mode.0.num_node_groups
			IgnoredFields: []string{
				"cluster_mode",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_elasticache_user_group", func(r *config.Resource) {
		r.References = config.References{
			"user_ids": config.Reference{
				Type:              "User",
				RefFieldName:      "UserIdRefs",
				SelectorFieldName: "UserIdSelector",
			},
		}
	})
}
