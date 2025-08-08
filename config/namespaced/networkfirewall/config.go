// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package networkfirewall

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the networkfirewall group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_networkfirewall_firewall_policy", func(r *config.Resource) {
		r.References = config.References{
			"firewall_policy.stateless_rule_group_reference.resource_arn": {
				TerraformName: "aws_networkfirewall_rule_group",
				Extractor:     common.PathARNExtractor,
			},
			"firewall_policy.stateful_rule_group_reference.resource_arn": {
				TerraformName: "aws_networkfirewall_rule_group",
				Extractor:     common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_networkfirewall_firewall", func(r *config.Resource) {
		r.UseAsync = true
		// elements of subnet_mapping set are in the format of (subnet_id,ip_address_type) pairs
		// once created the ip_address_type cannot be changed for a given subnet_id, and the given subnet_id cannot appear
		// in the set with multiple ip_address_type
		// ip_address_type is an optional+computed field which gets assigned if not specified
		// this causes permadiff when only subnet_id is specified
		// due to explanation above, the set can be hashed only based on subnet_ids while diffing,
		// because we are not interested in the changes of ip_address_type for an existing subnet_id
		// if the state and target includes the same set of subnets, we suppress the diff
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if state == nil || state.Empty() || diff == nil || diff.Empty() || diff.Destroy {
				return diff, nil
			}
			resourceData, _ := schema.InternalMap(r.TerraformResource.Schema).Data(state, diff)
			if !resourceData.HasChange("subnet_mapping") {
				return diff, nil
			}

			subnetOnlyHashFunc := func(v interface{}) int {
				var buf bytes.Buffer
				tfMap, ok := v.(map[string]interface{})
				if !ok {
					return 0
				}
				if id, ok := tfMap["subnet_id"].(string); ok {
					buf.WriteString(fmt.Sprintf("%s-", id))
				}
				return schema.HashString(buf.String())
			}

			osmRaw, nsmRaw := resourceData.GetChange("subnet_mapping")
			oldSubnetMappings := osmRaw.(*schema.Set)
			newSubnetMappings := nsmRaw.(*schema.Set)

			oldSubnetMappingsCustomHash := schema.NewSet(subnetOnlyHashFunc, oldSubnetMappings.List())
			newSubnetMappingsCustomHash := schema.NewSet(subnetOnlyHashFunc, newSubnetMappings.List())

			if oldSubnetMappingsCustomHash.HashEqual(newSubnetMappingsCustomHash) {
				for dk := range diff.Attributes {
					if strings.HasPrefix(dk, "subnet_mapping") {
						delete(diff.Attributes, dk)
					}
				}
			}
			if diff.Attributes != nil {
				delete(diff.Attributes, "firewall_status.#")
			}
			return diff, nil
		}
	})
}
