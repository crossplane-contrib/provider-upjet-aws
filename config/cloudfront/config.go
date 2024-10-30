// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudfront

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the cloudfront group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_cloudfront_distribution", func(r *config.Resource) {
		r.UseAsync = true
		delete(r.References, "origin.domain_name")
	})

	// Setting the field as sensitive to be able to pass the content from a k8s secret
	p.AddResourceConfigurator("aws_cloudfront_function", func(r *config.Resource) {
		r.TerraformResource.Schema["code"].Sensitive = true
	})

	// Setting the field as sensitive to be able to pass the content from a k8s secret
	p.AddResourceConfigurator("aws_cloudfront_public_key", func(r *config.Resource) {
		r.TerraformResource.Schema["encoded_key"].Sensitive = true
	})

	p.AddResourceConfigurator("aws_cloudfront_key_group", func(r *config.Resource) {
		r.References["items"] = config.Reference{
			TerraformName:     "aws_cloudfront_public_key",
			RefFieldName:      "ItemRefs",
			SelectorFieldName: "ItemSelector",
		}
	})

	p.AddResourceConfigurator("aws_cloudfront_realtime_log_config", func(r *config.Resource) {
		r.References["endpoint.kinesis_stream_config.stream_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})
}
