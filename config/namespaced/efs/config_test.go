// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package efs

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	oldPolicy = `{"Version":"2012-10-17","Statement":[{"Effect":"Deny","Principal":{"AWS":"*"},"Action":"*","Condition":{"Bool":{"aws:SecureTransport":"false"}}}]}`
	// Same policy as oldPolicy, but with keys reordered the way AWS echoes
	// the document back from DescribeFileSystemPolicy.
	equivalentPolicy = `{"Statement":[{"Condition":{"Bool":{"aws:SecureTransport":"false"}},"Action":"*","Principal":{"AWS":"*"},"Effect":"Deny"}],"Version":"2012-10-17"}`
	differentPolicy  = `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":"*"},"Action":"*"}]}`
)

func TestFileSystemPolicyCustomDiff(t *testing.T) {
	cases := map[string]struct {
		reason   string
		diff     *terraform.InstanceDiff
		wantErr  bool
		goneKeys []string
		wantKeys []string
	}{
		"NilDiff": {
			reason: "nil diff passes through unchanged",
			diff:   nil,
		},
		"NoPolicyAttribute": {
			reason: "a diff without a policy attribute is left untouched",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"file_system_id": {Old: "", New: "fs-123"},
				},
			},
			wantKeys: []string{"file_system_id"},
		},
		"CreateDiff_NoOld": {
			reason: "a create diff (no old value) is left untouched",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: "", New: oldPolicy},
				},
			},
			wantKeys: []string{"policy"},
		},
		"EquivalentPolicies_Suppressed": {
			reason: "a semantically equivalent but differently ordered policy is suppressed",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: oldPolicy, New: equivalentPolicy},
				},
			},
			goneKeys: []string{"policy"},
		},
		"DifferentPolicies_NotSuppressed": {
			reason: "a real policy change is kept in the diff",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: oldPolicy, New: differentPolicy},
				},
			},
			wantKeys: []string{"policy"},
		},
		"InvalidJSON_Errors": {
			reason: "an unparseable policy document returns an error instead of panicking",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: oldPolicy, New: "not-json"},
				},
			},
			wantErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := fileSystemPolicyCustomDiff(tc.diff, nil, nil)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("%s: expected an error, got none", tc.reason)
				}
				return
			}
			if err != nil {
				t.Fatalf("%s: unexpected error: %v", tc.reason, err)
			}
			if tc.diff == nil {
				if got != nil {
					t.Errorf("%s: expected nil diff, got non-nil", tc.reason)
				}
				return
			}
			for _, k := range tc.wantKeys {
				if _, ok := got.Attributes[k]; !ok {
					t.Errorf("%s: key %q should be present but is missing", tc.reason, k)
				}
			}
			for _, k := range tc.goneKeys {
				if _, ok := got.Attributes[k]; ok {
					t.Errorf("%s: key %q should have been suppressed but is still present", tc.reason, k)
				}
			}
		})
	}
}
