// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sns

import (
	"github.com/crossplane/upjet/pkg/config"
	awspolicy "github.com/hashicorp/awspolicyequivalence"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the sns group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_sns_topic_subscription", func(r *config.Resource) {
		r.References["endpoint"] = config.Reference{
			TerraformName: "aws_sqs_queue",
			Extractor:     common.PathARNExtractor,
		}
		r.References["topic_arn"] = config.Reference{
			TerraformName: "aws_sns_topic",
			Extractor:     common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_sns_topic", func(r *config.Resource) {
		// If the topic policy is unset on the Topic resource, don't late initialize it, to avoid conflicts with the
		// policy managed by a TopicPolicy resource.
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "policy")
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
		// Add the topicARN into the secret
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			if a, ok := attr["arn"].(string); ok {
				conn["arn"] = []byte(a)
			}

			return conn, nil
		}
	})
}
