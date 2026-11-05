// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudfront

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/upbound/provider-aws/v2/config/cloudfrontpolicy"
	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// embeddedObjectSource gives the wrapped conversion a copy of the source with
// the singleton list at block turned back into an embedded object.
// The singleton list conversion runs first and has already turned block into a
// list in the source.
// The annotation needs the v1beta2 embedded object paths, because those are the
// paths the FromAnnotation conversion writes back.
type embeddedObjectSource struct {
	conversion.PavedConversion
	block string
}

func (c embeddedObjectSource) ConvertPaved(src, target *fieldpath.Paved) (bool, error) {
	if l, err := src.GetValue(c.block); err == nil {
		if items, ok := l.([]any); ok {
			p := fieldpath.Pave(runtime.DeepCopyJSON(src.UnstructuredContent()))
			if len(items) > 0 {
				err = p.SetValue(c.block, items[0])
			} else {
				err = p.DeleteField(c.block)
			}
			if err != nil {
				return false, err
			}
			src = p
		}
	}
	return c.PavedConversion.ConvertPaved(src, target)
}

// Configure adds configurations for the cloudfront group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_cloudfront_distribution", func(r *config.Resource) {
		r.UseAsync = true
		delete(r.References, "origin.domain_name")
		// Extract the policy ID from status.atProvider.id: the external name of an
		// observed policy can still hold the policy name it was resolved from.
		for _, block := range []string{"default_cache_behavior", "ordered_cache_behavior"} {
			r.References[block+".cache_policy_id"] = config.Reference{
				TerraformName: "aws_cloudfront_cache_policy",
				Extractor:     common.PathTerraformIDExtractor,
			}
			r.References[block+".origin_request_policy_id"] = config.Reference{
				TerraformName: "aws_cloudfront_origin_request_policy",
				Extractor:     common.PathTerraformIDExtractor,
			}
			r.References[block+".response_headers_policy_id"] = config.Reference{
				TerraformName: "aws_cloudfront_response_headers_policy",
				Extractor:     common.PathTerraformIDExtractor,
			}
		}
		// The policy references exist in v1beta2 only. Preserve them through
		// an annotation when converting to v1beta1 so that they are not lost.
		for _, spec := range []string{"spec.forProvider", "spec.initProvider"} {
			for _, block := range []string{"defaultCacheBehavior", "orderedCacheBehavior[*]"} {
				for _, field := range []string{"cachePolicyId", "originRequestPolicyId", "responseHeadersPolicyId"} {
					for _, suffix := range []string{"Ref", "Selector"} {
						path := spec + "." + block + "." + field + suffix
						toAnnotation := conversion.NewNewlyIntroducedFieldConversion("v1beta2", "v1beta1", path, conversion.ToAnnotation)
						if block == "defaultCacheBehavior" {
							toAnnotation = embeddedObjectSource{PavedConversion: toAnnotation.(conversion.PavedConversion), block: spec + "." + block}
						}
						r.Conversions = append(r.Conversions,
							conversion.NewNewlyIntroducedFieldConversion("v1beta1", "v1beta2", path, conversion.FromAnnotation),
							toAnnotation)
					}
				}
			}
		}
		// Temporary edit until example generation pipeline fix
		for i, exp := range r.MetaResource.Examples {
			if exp.Name == "s3_distribution" {
				r.MetaResource.Examples = append(r.MetaResource.Examples[:i], r.MetaResource.Examples[i+1:]...)
				break
			}
		}
	})

	// Setting the field as sensitive to be able to pass the content from a k8s secret
	p.AddResourceConfigurator("aws_cloudfront_function", func(r *config.Resource) {
		r.TerraformResource.Schema["code"].Sensitive = true
		// Resolve key_value_store_associations from KeyValueStore ARNs.
		r.References["key_value_store_associations"] = config.Reference{
			TerraformName: "aws_cloudfront_key_value_store",
			Extractor:     common.PathARNExtractor,
		}
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

	p.AddResourceConfigurator("aws_cloudfront_origin_request_policy", func(r *config.Resource) {
		// name is exposed in v1beta2 since the switch to IdentifierFromProvider,
		// but v1beta1 (the storage version) omits it. Preserve it through an
		// annotation when converting to v1beta1 so that it is not lost.
		for _, path := range []string{"spec.forProvider.name", "spec.initProvider.name", "status.atProvider.name"} {
			r.Conversions = append(r.Conversions,
				conversion.NewNewlyIntroducedFieldConversion("v1beta1", "v1beta2", path, conversion.FromAnnotation),
				conversion.NewNewlyIntroducedFieldConversion("v1beta2", "v1beta1", path, conversion.ToAnnotation))
		}
	})

	// Observe policies by name: resolve a non-UUID ID to the policy ID on read.
	for _, name := range []string{"aws_cloudfront_cache_policy", "aws_cloudfront_origin_request_policy", "aws_cloudfront_response_headers_policy"} {
		p.AddResourceConfigurator(name, cloudfrontpolicy.ResolveNameOnRead)
	}

	p.AddResourceConfigurator("aws_cloudfront_vpc_origin", func(r *config.Resource) {
		r.AddSingletonListConversion("vpc_origin_endpoint_config", "vpcOriginEndpointConfig")
		r.AddSingletonListConversion("vpc_origin_endpoint_config[*].origin_ssl_protocols", "vpcOriginEndpointConfig[*].originSslProtocols")
	})
}
