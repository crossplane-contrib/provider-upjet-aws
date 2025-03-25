// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elasticache

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/comments"

	"github.com/upbound/provider-aws/config/namespaced/common"
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
