// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sns

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-upjet-aws/config/common"
)

// Configure adds configurations for the sns group.
func Configure(p *config.Provider) {
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
	})
}
