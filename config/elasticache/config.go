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
			"security_group_ids": config.Reference{
				Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupIDRefs",
				SelectorFieldName: "SecurityGroupIDSelector",
			},
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["cluster_address"].(string); ok {
				conn["cluster_address"] = []byte(a)
			}
			if a, ok := attr["configuration_endpoint"].(string); ok {
				conn["configuration_endpoint"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_elasticache_replication_group", func(r *config.Resource) {
		r.References["subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		r.References["security_group_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
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
		delete(r.References, "log_delivery_configuration.destination")
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["configuration_endpoint_address"].(string); ok {
				conn["configuration_endpoint_address"] = []byte(a)
			}
			if a, ok := attr["primary_endpoint_address"].(string); ok {
				conn["primary_endpoint_address"] = []byte(a)
			}
			if a, ok := attr["reader_endpoint_address"].(string); ok {
				conn["reader_endpoint_address"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_elasticache_user_group", func(r *config.Resource) {
		r.References["user_ids"] = config.Reference{
			Type:              "User",
			RefFieldName:      "UserIDRefs",
			SelectorFieldName: "UserIDSelector",
		}
	})
}
