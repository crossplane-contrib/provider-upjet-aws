// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sqs

import (
	"github.com/crossplane/upjet/pkg/config"
	awspolicy "github.com/hashicorp/awspolicyequivalence"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the sqs group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_sqs_queue_policy", func(r *config.Resource) {
		r.References["queue_url"] = config.Reference{
			TerraformName: "aws_sqs_queue",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_sqs_queue", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["url"].(string); ok {
				conn["url"] = []byte(a)
			}
			return conn, nil
		}
		// If the key policy is unset on the Queue resource, don't late initialize it, to avoid conflicts with the policy
		// managed by a QueuePolicy resource.
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "name_prefix", "policy")
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Attributes["policy"] == nil || diff.Attributes["policy"].Old == "" || diff.Attributes["policy"].New == "" {
				return diff, nil
			}

			vOld, err := common.RemovePolicyVersion(diff.Attributes["policy"].Old)
			if err != nil {
				return nil, errors.Wrap(err, "failed to remove Version from the old AWS policy document")
			}
			vNew, err := common.RemovePolicyVersion(diff.Attributes["policy"].New)
			if err != nil {
				return nil, errors.Wrap(err, "failed to remove Version from the new AWS policy document")
			}

			ok, err := awspolicy.PoliciesAreEquivalent(vOld, vNew)
			if err != nil {
				return nil, errors.Wrap(err, "failed to compare the old and the new AWS policy documents")
			}
			if ok {
				delete(diff.Attributes, "policy")
			}
			return diff, nil
		}

	})

	p.AddResourceConfigurator("aws_sqs_queue_redrive_policy", func(r *config.Resource) {
		r.References["queue_url"] = config.Reference{
			TerraformName: "aws_sqs_queue",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_sqs_queue_redrive_allow_policy", func(r *config.Resource) {
		r.References["queue_url"] = config.Reference{
			TerraformName: "aws_sqs_queue",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
