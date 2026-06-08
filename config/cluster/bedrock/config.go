// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bedrock

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_bedrock_guardrail", func(r *config.Resource) {
		r.References["kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		// aws_bedrock_guardrail is a Terraform Plugin Framework resource, so the
		// MaxItems==1 constraints on its single-nested config blocks are not
		// visible to the SingletonListEmbedder schema traverser. The following
		// blocks are enforced as singletons upstream (listvalidator.SizeAtMost(1)
		// on the wrappers, and ListOfObjectsAttribute(ctx, 1, ...) on tier_config)
		// and are flattened to embedded objects. The remaining nested blocks
		// (filters_config, pii_entities_config, regexes_config, topics_config,
		// managed_word_lists_config, words_config) are genuinely repeatable and
		// intentionally retain their list representation.
		r.AddSingletonListConversion("content_policy_config", "contentPolicyConfig")
		r.AddSingletonListConversion("content_policy_config[*].tier_config", "contentPolicyConfig[*].tierConfig")
		r.AddSingletonListConversion("contextual_grounding_policy_config", "contextualGroundingPolicyConfig")
		r.AddSingletonListConversion("cross_region_config", "crossRegionConfig")
		r.AddSingletonListConversion("sensitive_information_policy_config", "sensitiveInformationPolicyConfig")
		r.AddSingletonListConversion("topic_policy_config", "topicPolicyConfig")
		r.AddSingletonListConversion("topic_policy_config[*].tier_config", "topicPolicyConfig[*].tierConfig")
		r.AddSingletonListConversion("word_policy_config", "wordPolicyConfig")
	})
}
