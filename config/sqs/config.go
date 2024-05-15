// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sqs

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the sqs group.
func Configure(p *config.Provider) {
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
