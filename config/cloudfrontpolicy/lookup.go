// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

// Package cloudfrontpolicy resolves CloudFront policy names to policy IDs for
// the cache, origin request and response headers policy resources.
package cloudfrontpolicy

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// cloudFrontClientProvider matches the Terraform provider meta conns.AWSClient.
// Asserting it instead of importing the provider keeps the provider out of this
// package and its test build.
type cloudFrontClientProvider interface {
	CloudFrontClient(ctx context.Context) *cloudfront.Client
}

// policyIDByName returns the ID of the CloudFront policy with the given name,
// or an empty string if there is none.
type policyIDByName func(ctx context.Context, conn *cloudfront.Client, name string) (string, error)

var cloudFrontPolicyLookups = map[string]policyIDByName{
	"aws_cloudfront_cache_policy":            cachePolicyIDByName,
	"aws_cloudfront_origin_request_policy":   originRequestPolicyIDByName,
	"aws_cloudfront_response_headers_policy": responseHeadersPolicyIDByName,
}

// ResolveNameOnRead lets a CloudFront cache, origin request or response
// headers policy be observed by name.
// CloudFront reads these policies by ID only.
// If the Terraform ID is not a UUID, the read treats it as a policy name
// and looks up the policy ID first.
// The Terraform ID is the external name, or forProvider.name while the external name is empty.
func ResolveNameOnRead(r *config.Resource) {
	lookup, ok := cloudFrontPolicyLookups[r.Name]
	// The generation-time schema carries no CRUD functions.
	if !ok || r.TerraformResource == nil || r.TerraformResource.ReadWithoutTimeout == nil {
		return
	}
	r.TerraformResource.ReadWithoutTimeout = readWithPolicyNameLookup(r.TerraformResource.ReadWithoutTimeout,
		func(ctx context.Context, meta any, name string) (string, error) {
			client, ok := meta.(cloudFrontClientProvider)
			if !ok {
				return "", fmt.Errorf("unexpected Terraform provider meta type %T", meta)
			}
			return lookup(ctx, client.CloudFrontClient(ctx), name)
		})
}

// readWithPolicyNameLookup wraps read. If the resource ID is not a UUID, it
// calls find and reads with the returned policy ID, or reports the resource as
// not found when find returns none. The cluster and namespaced providers share
// the Terraform resource, so read may already be wrapped. The inner wrapper
// then sees the resolved UUID and calls read directly.
func readWithPolicyNameLookup(read schema.ReadContextFunc, find func(ctx context.Context, meta any, name string) (string, error)) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
		if uuid.Validate(d.Id()) == nil {
			return read(ctx, d, meta)
		}
		id, err := find(ctx, meta, d.Id())
		if err != nil {
			return diag.FromErr(fmt.Errorf("resolving CloudFront policy name %q to an ID: %w", d.Id(), err))
		}
		if id == "" {
			d.SetId("")
			return nil
		}
		d.SetId(id)
		return read(ctx, d, meta)
	}
}

// findPolicyID pages through list until it finds the policy named name.
func findPolicyID[T any](name string, list func(marker *string) ([]T, *string, error), nameAndID func(T) (string, string)) (string, error) {
	var marker *string
	for {
		items, next, err := list(marker)
		if err != nil {
			return "", err
		}
		for _, it := range items {
			if n, id := nameAndID(it); n == name {
				return id, nil
			}
		}
		if aws.ToString(next) == "" {
			return "", nil
		}
		marker = next
	}
}

func cachePolicyIDByName(ctx context.Context, conn *cloudfront.Client, name string) (string, error) {
	return findPolicyID(name,
		func(marker *string) ([]types.CachePolicySummary, *string, error) {
			out, err := conn.ListCachePolicies(ctx, &cloudfront.ListCachePoliciesInput{Marker: marker})
			if err != nil || out.CachePolicyList == nil {
				return nil, nil, err
			}
			return out.CachePolicyList.Items, out.CachePolicyList.NextMarker, nil
		},
		func(s types.CachePolicySummary) (string, string) {
			if s.CachePolicy == nil || s.CachePolicy.CachePolicyConfig == nil {
				return "", ""
			}
			return aws.ToString(s.CachePolicy.CachePolicyConfig.Name), aws.ToString(s.CachePolicy.Id)
		})
}

func originRequestPolicyIDByName(ctx context.Context, conn *cloudfront.Client, name string) (string, error) {
	return findPolicyID(name,
		func(marker *string) ([]types.OriginRequestPolicySummary, *string, error) {
			out, err := conn.ListOriginRequestPolicies(ctx, &cloudfront.ListOriginRequestPoliciesInput{Marker: marker})
			if err != nil || out.OriginRequestPolicyList == nil {
				return nil, nil, err
			}
			return out.OriginRequestPolicyList.Items, out.OriginRequestPolicyList.NextMarker, nil
		},
		func(s types.OriginRequestPolicySummary) (string, string) {
			if s.OriginRequestPolicy == nil || s.OriginRequestPolicy.OriginRequestPolicyConfig == nil {
				return "", ""
			}
			return aws.ToString(s.OriginRequestPolicy.OriginRequestPolicyConfig.Name), aws.ToString(s.OriginRequestPolicy.Id)
		})
}

func responseHeadersPolicyIDByName(ctx context.Context, conn *cloudfront.Client, name string) (string, error) {
	return findPolicyID(name,
		func(marker *string) ([]types.ResponseHeadersPolicySummary, *string, error) {
			out, err := conn.ListResponseHeadersPolicies(ctx, &cloudfront.ListResponseHeadersPoliciesInput{Marker: marker})
			if err != nil || out.ResponseHeadersPolicyList == nil {
				return nil, nil, err
			}
			return out.ResponseHeadersPolicyList.Items, out.ResponseHeadersPolicyList.NextMarker, nil
		},
		func(s types.ResponseHeadersPolicySummary) (string, string) {
			if s.ResponseHeadersPolicy == nil || s.ResponseHeadersPolicy.ResponseHeadersPolicyConfig == nil {
				return "", ""
			}
			return aws.ToString(s.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.Name), aws.ToString(s.ResponseHeadersPolicy.Id)
		})
}
