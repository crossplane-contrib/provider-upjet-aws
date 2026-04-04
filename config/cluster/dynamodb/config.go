// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dynamodb

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the dynamodb group.
func Configure(p *config.Provider) { //nolint:gocyclo
	// currently needs an ARN reference for external name
	p.AddResourceConfigurator("aws_dynamodb_contributor_insights", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_kinesis_streaming_destination", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
		}

		r.References["stream_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_table_item", func(r *config.Resource) {
		r.References["table_name"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
		}
		delete(r.References, "hash_key")
	})

	p.AddResourceConfigurator("aws_dynamodb_resource_policy", func(r *config.Resource) {
		r.References["resource_arn"] = config.Reference{
			TerraformName: "aws_dynamodb_table",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_dynamodb_table", func(r *config.Resource) {
		r.References["server_side_encryption.kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.References["replica.kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		// global_secondary_index is a TypeSet. The Terraform default set hash
		// is computed from ALL fields of a set element, including computed-only
		// fields like "warm_throughput" and "key_schema". AWS populates these
		// computed fields after creation, so the state hash (which includes
		// warm_throughput/key_schema values) diverges from the config hash
		// (which does not have them). This causes a perpetual diff where
		// Terraform sees the GSI as "delete old hash + add new hash" on every
		// reconcile, even though the actual user-specified values are identical.
		//
		// To fix this, we define a custom hash function that only uses the
		// user-configurable fields (name, hash_key, range_key, projection_type,
		// non_key_attributes, read_capacity, write_capacity) and compare old vs
		// new sets using that hash. If the sets are equal under the custom hash,
		// we suppress all global_secondary_index diffs.
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if state == nil || state.Empty() || diff == nil || diff.Empty() || diff.Destroy {
				return diff, nil
			}

			// --- GSI TypeSet hash suppression ---
			// (see comment block above for full explanation)
			resourceData, err := schema.InternalMap(r.TerraformResource.Schema).Data(state, diff)
			if err == nil && resourceData.HasChange("global_secondary_index") {
				gsiUserFieldsHashFunc := func(v interface{}) int {
					var buf bytes.Buffer
					tfMap, ok := v.(map[string]interface{})
					if !ok {
						return 0
					}
					if name, ok := tfMap["name"].(string); ok {
						fmt.Fprintf(&buf, "%s-", name)
					}
					if hashKey, ok := tfMap["hash_key"].(string); ok {
						fmt.Fprintf(&buf, "%s-", hashKey)
					}
					if rangeKey, ok := tfMap["range_key"].(string); ok {
						fmt.Fprintf(&buf, "%s-", rangeKey)
					}
					if projType, ok := tfMap["projection_type"].(string); ok {
						fmt.Fprintf(&buf, "%s-", projType)
					}
					// read_capacity and write_capacity are "number" type in the
					// Terraform schema, which maps to float64 in Go (not int).
					if readCap, ok := tfMap["read_capacity"].(float64); ok {
						fmt.Fprintf(&buf, "%g-", readCap)
					}
					if writeCap, ok := tfMap["write_capacity"].(float64); ok {
						fmt.Fprintf(&buf, "%g-", writeCap)
					}
					if nka, ok := tfMap["non_key_attributes"]; ok {
						if nkaSet, ok := nka.(*schema.Set); ok {
							nkaList := make([]string, 0, nkaSet.Len())
							for _, v := range nkaSet.List() {
								if s, ok := v.(string); ok {
									nkaList = append(nkaList, s)
								}
							}
							sort.Strings(nkaList)
							for _, s := range nkaList {
								fmt.Fprintf(&buf, "%s-", s)
							}
						}
					}
					return schema.HashString(buf.String())
				}

				oRaw, nRaw := resourceData.GetChange("global_secondary_index")
				oldGSIs := oRaw.(*schema.Set)
				newGSIs := nRaw.(*schema.Set)

				oldGSIsCustomHash := schema.NewSet(gsiUserFieldsHashFunc, oldGSIs.List())
				newGSIsCustomHash := schema.NewSet(gsiUserFieldsHashFunc, newGSIs.List())

				if oldGSIsCustomHash.HashEqual(newGSIsCustomHash) {
					for dk := range diff.Attributes {
						if strings.HasPrefix(dk, "global_secondary_index") {
							delete(diff.Attributes, dk)
						}
					}
				}
			}

			// --- TTL and SSE spurious diff suppression ---
			// After late-initialization, the observed state contains
			// ttl (attribute_name, enabled) and server_side_encryption
			// (enabled) values populated by AWS. However, these values
			// may not be properly written back to the Terraform config
			// during late-init in v1beta1, causing a diff where the
			// config appears to "remove" these fields on every reconcile:
			//   ttl.0.attribute_name:              Old:"TimeToExist" → New:"" (NewRemoved)
			//   ttl.0.enabled:                     Old:"true"        → New:"false"
			//   server_side_encryption.0.enabled:   Old:"true"        → New:"false" (NewRemoved)
			// This triggers an update that fails (empty attributeName
			// violates AWS API constraints) and loops indefinitely.
			// We suppress these diffs when the state has values but the
			// config is trying to remove them.
			if d, ok := diff.Attributes["ttl.0.attribute_name"]; ok && d.Old != "" && d.NewRemoved {
				delete(diff.Attributes, "ttl.0.attribute_name")
				delete(diff.Attributes, "ttl.0.enabled")
			}
			if d, ok := diff.Attributes["server_side_encryption.0.enabled"]; ok && d.Old == "true" && d.NewRemoved {
				delete(diff.Attributes, "server_side_encryption.0.enabled")
			}

			return diff, nil
		}
	})
}
