// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/utils/ptr"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

func mgWithAnnotations(annotations map[string]string) *fake.Managed {
	return &fake.Managed{ObjectMeta: metav1.ObjectMeta{Annotations: annotations}}
}

func TestApplyManagedResourceOverrides(t *testing.T) {
	const validIdentity = "alice@example.com"
	cases := map[string]struct {
		reason          string
		spec            *v1beta1.ProviderConfigSpec
		mg              *fake.Managed
		wantChain       []v1beta1.AssumeRoleOptions
		wantSubstituted []string
		wantErrSnippet  string
	}{
		"NilSpecNoop": {
			reason: "A nil spec is a defensive no-op.",
			spec:   nil,
			mg:     mgWithAnnotations(nil),
		},
		"NoSentinelNoop": {
			reason: "Entries without the sentinel are left untouched and no extra cache key is produced.",
			spec: &v1beta1.ProviderConfigSpec{AssumeRoleChain: []v1beta1.AssumeRoleOptions{
				{RoleARN: ptr.To("arn:aws:iam::111111111111:role/r"), SourceIdentity: ptr.To("static-value")},
			}},
			mg: mgWithAnnotations(map[string]string{AnnotationSourceIdentity: validIdentity}),
			wantChain: []v1beta1.AssumeRoleOptions{
				{RoleARN: ptr.To("arn:aws:iam::111111111111:role/r"), SourceIdentity: ptr.To("static-value")},
			},
		},
		"SentinelSubstituted": {
			reason: "An entry with the sentinel is rewritten to the annotation value and reported back.",
			spec: &v1beta1.ProviderConfigSpec{AssumeRoleChain: []v1beta1.AssumeRoleOptions{
				{RoleARN: ptr.To("arn:aws:iam::111111111111:role/r"), SourceIdentity: ptr.To(SentinelFromManagedResource)},
			}},
			mg: mgWithAnnotations(map[string]string{AnnotationSourceIdentity: validIdentity}),
			wantChain: []v1beta1.AssumeRoleOptions{
				{RoleARN: ptr.To("arn:aws:iam::111111111111:role/r"), SourceIdentity: ptr.To(validIdentity)},
			},
			wantSubstituted: []string{validIdentity},
		},
		"MixedChain": {
			reason: "Only entries with the sentinel are rewritten; static entries are preserved.",
			spec: &v1beta1.ProviderConfigSpec{AssumeRoleChain: []v1beta1.AssumeRoleOptions{
				{RoleARN: ptr.To("arn:aws:iam::111111111111:role/a"), SourceIdentity: ptr.To(SentinelFromManagedResource)},
				{RoleARN: ptr.To("arn:aws:iam::222222222222:role/b"), SourceIdentity: ptr.To("static")},
				{RoleARN: ptr.To("arn:aws:iam::333333333333:role/c"), SourceIdentity: nil},
			}},
			mg: mgWithAnnotations(map[string]string{AnnotationSourceIdentity: validIdentity}),
			wantChain: []v1beta1.AssumeRoleOptions{
				{RoleARN: ptr.To("arn:aws:iam::111111111111:role/a"), SourceIdentity: ptr.To(validIdentity)},
				{RoleARN: ptr.To("arn:aws:iam::222222222222:role/b"), SourceIdentity: ptr.To("static")},
				{RoleARN: ptr.To("arn:aws:iam::333333333333:role/c"), SourceIdentity: nil},
			},
			wantSubstituted: []string{validIdentity},
		},
		"MissingAnnotationFailsClosed": {
			reason: "Sentinel without an annotation is an error so we never send the literal sentinel to STS.",
			spec: &v1beta1.ProviderConfigSpec{AssumeRoleChain: []v1beta1.AssumeRoleOptions{
				{SourceIdentity: ptr.To(SentinelFromManagedResource)},
			}},
			mg:             mgWithAnnotations(nil),
			wantErrSnippet: "has no",
		},
		"EmptyAnnotationFailsClosed": {
			reason:         "An empty annotation value is treated the same as a missing one.",
			spec:           &v1beta1.ProviderConfigSpec{AssumeRoleChain: []v1beta1.AssumeRoleOptions{{SourceIdentity: ptr.To(SentinelFromManagedResource)}}},
			mg:             mgWithAnnotations(map[string]string{AnnotationSourceIdentity: ""}),
			wantErrSnippet: "has no",
		},
		"InvalidAnnotationFailsClosed": {
			reason: "Values outside the AWS SourceIdentity character set are rejected.",
			spec: &v1beta1.ProviderConfigSpec{AssumeRoleChain: []v1beta1.AssumeRoleOptions{
				{SourceIdentity: ptr.To(SentinelFromManagedResource)},
			}},
			mg:             mgWithAnnotations(map[string]string{AnnotationSourceIdentity: "has spaces"}),
			wantErrSnippet: "invalid SourceIdentity",
		},
		"TooShortAnnotationFailsClosed": {
			reason: "A single-character value violates the 2-64 length bound.",
			spec: &v1beta1.ProviderConfigSpec{AssumeRoleChain: []v1beta1.AssumeRoleOptions{
				{SourceIdentity: ptr.To(SentinelFromManagedResource)},
			}},
			mg:             mgWithAnnotations(map[string]string{AnnotationSourceIdentity: "a"}),
			wantErrSnippet: "invalid SourceIdentity",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := applyManagedResourceOverrides(tc.spec, tc.mg)
			if tc.wantErrSnippet != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErrSnippet) {
					t.Fatalf("%s\nexpected error containing %q, got %v", tc.reason, tc.wantErrSnippet, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("%s\nunexpected error: %v", tc.reason, err)
			}
			if diff := cmp.Diff(tc.wantSubstituted, got); diff != "" {
				t.Errorf("%s\nsubstituted values: -want, +got:\n%s", tc.reason, diff)
			}
			if tc.spec != nil {
				if diff := cmp.Diff(tc.wantChain, tc.spec.AssumeRoleChain); diff != "" {
					t.Errorf("%s\nresulting chain: -want, +got:\n%s", tc.reason, diff)
				}
			}
		})
	}
}

func TestApplyManagedResourceOverridesAliasing(t *testing.T) {
	// Mutating the returned spec must not write through to the MR's
	// annotation map, since callers commonly hold a reference to it.
	annot := map[string]string{AnnotationSourceIdentity: "alice@example.com"}
	mg := mgWithAnnotations(annot)
	spec := &v1beta1.ProviderConfigSpec{AssumeRoleChain: []v1beta1.AssumeRoleOptions{
		{SourceIdentity: ptr.To(SentinelFromManagedResource)},
	}}
	if _, err := applyManagedResourceOverrides(spec, mg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	*spec.AssumeRoleChain[0].SourceIdentity = "mutated"
	if got := annot[AnnotationSourceIdentity]; got != "alice@example.com" {
		t.Errorf("annotation map was aliased into the spec pointer; got %q", got)
	}
}
