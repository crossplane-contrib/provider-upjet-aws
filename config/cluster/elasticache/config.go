// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elasticache

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/crossplane/upjet/pkg/types/comments"

	"github.com/upbound/provider-aws/apis/cluster/elasticache/v1beta1"
	"github.com/upbound/provider-aws/apis/cluster/elasticache/v1beta2"
	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the elasticache group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_elasticache_cluster", func(r *config.Resource) {
		r.References["parameter_group_name"] = config.Reference{
			TerraformName: "aws_elasticache_parameter_group",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			// This only works for memcached clusters
			if a, ok := attr["cluster_address"].(string); ok {
				conn["cluster_address"] = []byte(a)
			}
			if a, ok := attr["port"]; ok {
				conn["port"] = []byte(fmt.Sprintf("%v", a))
			}
			return conn, nil
		}
		// log_delivery_configuration.destination can point to either
		// a CloudWatch Logs LogGroup or Kinesis Data Firehose resource.
		delete(r.References, "log_delivery_configuration.destination")
	})

	p.AddResourceConfigurator("aws_elasticache_replication_group", func(r *config.Resource) {
		r.References["subnet_group_name"] = config.Reference{
			TerraformName: "aws_elasticache_subnet_group",
		}
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
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
			if a, ok := attr["port"]; ok {
				conn["port"] = []byte(fmt.Sprintf("%v", a))
			}
			return conn, nil
		}

		// Auth token generation
		desc, err := comments.New("If true, the auth token will be auto-generated and"+
			" stored in the Secret referenced by the authTokenSecretRef field.",
			comments.WithTFTag("-"))
		if err != nil {
			panic(errors.Wrap(err, "cannot configure the generated comment for the auto_generate_auth_token argument of the aws_elasticache_replication_group resource"))
		}

		r.TerraformResource.Schema["auto_generate_auth_token"] = &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Description: desc.String(),
		}
		r.InitializerFns = append(r.InitializerFns,
			common.PasswordGenerator(
				"spec.forProvider.authTokenSecretRef",
				"spec.forProvider.autoGenerateAuthToken",
			))
		r.TerraformResource.Schema["auth_token"].Description = "If you set" +
			" autoGenerateAuthToken to true, the Secret referenced here will be" +
			" created or updated with generated auth token if it does not already" +
			" contain one."

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

	p.AddResourceConfigurator("aws_elasticache_serverless_cache", func(r *config.Resource) {
		r.UseAsync = true
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			if endpoints, ok := attr["endpoint"].([]any); ok {
				for i, ep := range endpoints {
					if endpoint, ok := ep.(map[string]any); ok && len(endpoint) > 0 {
						if address, ok := endpoint["address"].(string); ok {
							key := fmt.Sprintf("endpoint_%d_address", i)
							conn[key] = []byte(address)
						}
						if port, ok := endpoint["port"]; ok {
							key := fmt.Sprintf("endpoint_%d_port", i)
							conn[key] = []byte(fmt.Sprintf("%v", port))
						}
					}
				}
			}
			if readerendpoints, ok := attr["reader_endpoint"].([]any); ok {
				for i, rp := range readerendpoints {
					if readerendpoint, ok := rp.(map[string]any); ok && len(readerendpoint) > 0 {
						if address, ok := readerendpoint["address"].(string); ok {
							key := fmt.Sprintf("reader_endpoint_%d_address", i)
							conn[key] = []byte(address)
						}
						if port, ok := readerendpoint["port"]; ok {
							key := fmt.Sprintf("reader_endpoint_%d_port", i)
							conn[key] = []byte(fmt.Sprintf("%v", port))
						}
					}
				}
			}

			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_elasticache_user_group", func(r *config.Resource) {
		r.References["user_ids"] = config.Reference{
			TerraformName:     "aws_elasticache_user",
			RefFieldName:      "UserIDRefs",
			SelectorFieldName: "UserIDSelector",
		}
	})
}
