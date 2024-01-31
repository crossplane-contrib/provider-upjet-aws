/*
Copyright 2021 Upbound Inc.
*/

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"

	"github.com/upbound/provider-aws/apis/elasticache/v1beta1"
	"github.com/upbound/provider-aws/apis/elasticache/v1beta2"
)

// Configure adds configurations for the elasticache group.
func Configure(p *config.Provider) { //nolint:gocyclo
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

		r.Version = "v1beta2"
		r.Conversions = append(r.Conversions,
			conversion.NewCustomConverter("v1beta1", "v1beta2", func(src, target xpresource.Managed) error {
				srcTyped := src.(*v1beta1.ReplicationGroup)
				targetTyped := target.(*v1beta2.ReplicationGroup)
				if len(srcTyped.Spec.ForProvider.ClusterMode) > 0 {
					if srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups != nil {
						targetTyped.Spec.ForProvider.NumNodeGroups = srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups
					}
					if srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
						targetTyped.Spec.ForProvider.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup
					}
				}
				if len(srcTyped.Spec.InitProvider.ClusterMode) > 0 {
					if srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups != nil {
						targetTyped.Spec.InitProvider.NumNodeGroups = srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups
					}
					if srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
						targetTyped.Spec.InitProvider.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup
					}
				}
				if len(srcTyped.Status.AtProvider.ClusterMode) > 0 {
					if srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups != nil {
						targetTyped.Status.AtProvider.NumNodeGroups = srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups
					}
					if srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
						targetTyped.Status.AtProvider.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup
					}
				}
				return nil
			}),
			conversion.NewCustomConverter("v1beta2", "v1beta1", func(src, target xpresource.Managed) error {
				srcTyped := src.(*v1beta2.ReplicationGroup)
				targetTyped := target.(*v1beta1.ReplicationGroup)
				cm := v1beta1.ClusterModeParameters{}
				if srcTyped.Spec.ForProvider.NumNodeGroups != nil {
					cm.NumNodeGroups = srcTyped.Spec.ForProvider.NumNodeGroups
				}
				if srcTyped.Spec.ForProvider.ReplicasPerNodeGroup != nil {
					cm.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ReplicasPerNodeGroup
				}
				targetTyped.Spec.ForProvider.ClusterMode = []v1beta1.ClusterModeParameters{cm}

				cmi := v1beta1.ClusterModeInitParameters{}
				if srcTyped.Spec.InitProvider.NumNodeGroups != nil {
					cm.NumNodeGroups = srcTyped.Spec.InitProvider.NumNodeGroups
				}
				if srcTyped.Spec.InitProvider.ReplicasPerNodeGroup != nil {
					cm.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ReplicasPerNodeGroup
				}
				targetTyped.Spec.InitProvider.ClusterMode = []v1beta1.ClusterModeInitParameters{cmi}

				cmo := v1beta1.ClusterModeObservation{}
				if srcTyped.Status.AtProvider.NumNodeGroups != nil {
					cm.NumNodeGroups = srcTyped.Status.AtProvider.NumNodeGroups
				}
				if srcTyped.Status.AtProvider.ReplicasPerNodeGroup != nil {
					cm.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ReplicasPerNodeGroup
				}
				targetTyped.Status.AtProvider.ClusterMode = []v1beta1.ClusterModeObservation{cmo}
				return nil
			}),
		)
	})

	p.AddResourceConfigurator("aws_elasticache_user_group", func(r *config.Resource) {
		r.References["user_ids"] = config.Reference{
			Type:              "User",
			RefFieldName:      "UserIDRefs",
			SelectorFieldName: "UserIDSelector",
		}
	})
}
