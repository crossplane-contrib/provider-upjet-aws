// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package opensearch

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the opensearch group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_opensearch_domain", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "access_policies")
		r.References["encrypt_at_rest.kms_key_id"] = config.Reference{
			// its KMS key ARN in AWS API
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		r.References["vpc_options.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}

		r.References["vpc_options.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}

		r.UseAsync = true

		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			asoDiff, ok := diff.Attributes["advanced_security_options.#"]
			if ok && asoDiff.Old == "" && asoDiff.New == "" && asoDiff.NewComputed {
				delete(diff.Attributes, "advanced_security_options.#")
			}
			return diff, nil
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["endpoint"].(string); ok {
				conn["endpoint"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_opensearch_domain_policy", func(r *config.Resource) {
		r.References["domain_name"] = config.Reference{
			TerraformName: "aws_opensearch_domain",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_opensearch_domain_saml_options", func(r *config.Resource) {
		r.UseAsync = true
	})
}
