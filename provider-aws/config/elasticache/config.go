/*
Copyright 2021 Upbound Inc.
*/

package elasticache

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for elasticache group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elasticache_parameter_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
	})

	p.AddResourceConfigurator("aws_elasticache_subnet_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
	})

	p.AddResourceConfigurator("aws_elasticache_cluster", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["cluster_id"] = name
			},
			OmittedFields: []string{
				"cluster_id",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
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
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["replication_group_id"] = name
			},
			OmittedFields: []string{
				"replication_group_id",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"subnet_group_name": config.Reference{
				Type: "SubnetGroup",
			},
			"security_group_ids": config.Reference{
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.SecurityGroup",
				RefFieldName:      "SecurityGroupIdRefs",
				SelectorFieldName: "SecurityGroupIdSelector",
			},
			"kms_key_id": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key",
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

	p.AddResourceConfigurator("aws_elasticache_user", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		if s, ok := r.TerraformResource.Schema["passwords"]; ok {
			s.Sensitive = true
		}
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["user_id"] = name
			},
			OmittedFields: []string{
				"user_id",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
	})

	p.AddResourceConfigurator("aws_elasticache_user_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["user_group_id"] = name
			},
			OmittedFields: []string{
				"user_group_id",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"user_ids": config.Reference{
				Type:              "User",
				RefFieldName:      "UserIdRefs",
				SelectorFieldName: "UserIdSelector",
			},
		}
	})
}
