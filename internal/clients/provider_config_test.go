// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	stscredstypesv2 "github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"k8s.io/utils/ptr"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

func TestSetAssumeRoleOptions(t *testing.T) {
	cases := map[string]struct {
		reason string
		aro    v1beta1.AssumeRoleOptions
		want   stscreds.AssumeRoleOptions
	}{
		"Empty": {
			reason: "An empty AssumeRoleOptions should leave the stscreds options at their zero value.",
			aro:    v1beta1.AssumeRoleOptions{},
			want:   stscreds.AssumeRoleOptions{},
		},
		"ExternalIDOnly": {
			reason: "ExternalID should be propagated.",
			aro:    v1beta1.AssumeRoleOptions{ExternalID: ptr.To("ext-1")},
			want:   stscreds.AssumeRoleOptions{ExternalID: ptr.To("ext-1")},
		},
		"SourceIdentityOnly": {
			reason: "SourceIdentity should be propagated.",
			aro:    v1beta1.AssumeRoleOptions{SourceIdentity: ptr.To("alice@example.com")},
			want:   stscreds.AssumeRoleOptions{SourceIdentity: ptr.To("alice@example.com")},
		},
		"SourceIdentityNilStaysNil": {
			reason: "A nil SourceIdentity must remain nil so existing callers see no behavioural change.",
			aro: v1beta1.AssumeRoleOptions{
				ExternalID:     ptr.To("ext-2"),
				SourceIdentity: nil,
			},
			want: stscreds.AssumeRoleOptions{
				ExternalID:     ptr.To("ext-2"),
				SourceIdentity: nil,
			},
		},
		"TagsAndTransitiveTagKeys": {
			reason: "Tags and TransitiveTagKeys should be propagated alongside SourceIdentity.",
			aro: v1beta1.AssumeRoleOptions{
				SourceIdentity: ptr.To("alice@example.com"),
				Tags: []v1beta1.Tag{
					{Key: ptr.To("originating-principal"), Value: ptr.To("alice@example.com")},
				},
				TransitiveTagKeys: []string{"originating-principal"},
			},
			want: stscreds.AssumeRoleOptions{
				SourceIdentity: ptr.To("alice@example.com"),
				Tags: []stscredstypesv2.Tag{
					{Key: ptr.To("originating-principal"), Value: ptr.To("alice@example.com")},
				},
				TransitiveTagKeys: []string{"originating-principal"},
			},
		},
		"AllFields": {
			reason: "All supported fields should be propagated together.",
			aro: v1beta1.AssumeRoleOptions{
				ExternalID:     ptr.To("ext-3"),
				SourceIdentity: ptr.To("bob@example.com"),
				Tags: []v1beta1.Tag{
					{Key: ptr.To("k1"), Value: ptr.To("v1")},
					{Key: ptr.To("k2"), Value: ptr.To("v2")},
				},
				TransitiveTagKeys: []string{"k1"},
			},
			want: stscreds.AssumeRoleOptions{
				ExternalID:     ptr.To("ext-3"),
				SourceIdentity: ptr.To("bob@example.com"),
				Tags: []stscredstypesv2.Tag{
					{Key: ptr.To("k1"), Value: ptr.To("v1")},
					{Key: ptr.To("k2"), Value: ptr.To("v2")},
				},
				TransitiveTagKeys: []string{"k1"},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := stscreds.AssumeRoleOptions{}
			SetAssumeRoleOptions(tc.aro)(&got)
			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(stscreds.AssumeRoleOptions{}, stscredstypesv2.Tag{})); diff != "" {
				t.Errorf("%s\nSetAssumeRoleOptions(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
