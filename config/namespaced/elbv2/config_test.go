// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elbv2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testARN = "arn:aws:elasticloadbalancing:us-east-1:123456789012:targetgroup/my-tg/abc123"

func configWithForward() *terraform.ResourceConfig {
	return &terraform.ResourceConfig{
		Config: map[string]interface{}{
			"action": []interface{}{
				map[string]interface{}{
					"forward": []interface{}{
						map[string]interface{}{
							"target_group": []interface{}{
								map[string]interface{}{"arn": testARN, "weight": 1},
							},
						},
					},
				},
			},
		},
	}
}

func configWithARNOnly() *terraform.ResourceConfig {
	return &terraform.ResourceConfig{
		Config: map[string]interface{}{
			"action": []interface{}{
				map[string]interface{}{
					"target_group_arn": testARN,
				},
			},
		},
	}
}

func TestLBListenerRuleCustomDiff(t *testing.T) {
	cases := map[string]struct {
		reason   string
		diff     *terraform.InstanceDiff
		state    *terraform.InstanceState
		cfg      *terraform.ResourceConfig
		wantKeys []string
		goneKeys []string
	}{
		"NilDiff": {
			reason: "nil diff passes through unchanged",
			diff:   nil,
		},
		"EmptyDiff": {
			reason: "empty diff passes through unchanged",
			diff:   &terraform.InstanceDiff{},
		},
		"DestroyDiff": {
			reason: "destroy diff is never modified",
			diff: &terraform.InstanceDiff{
				Destroy: true,
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: testARN, New: "", NewRemoved: true},
				},
			},
			wantKeys: []string{"action.0.target_group_arn"},
		},
		// Case A: arn absent from state (Old=""), present in config — late-init artifact
		"CaseA_Suppressed": {
			reason: "spurious arn addition is suppressed when the same action has forward configured",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: "", New: testARN},
				},
			},
			cfg:      configWithForward(),
			goneKeys: []string{"action.0.target_group_arn"},
		},
		"CaseA_NotSuppressed_NoForwardInConfig": {
			reason: "arn addition is kept when action only has target_group_arn (real user intent)",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: "", New: testARN},
				},
			},
			cfg:      configWithARNOnly(),
			wantKeys: []string{"action.0.target_group_arn"},
		},
		"CaseA_NotSuppressed_NilConfig": {
			reason: "arn addition is kept when config is nil",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: "", New: testARN},
				},
			},
			cfg:      nil,
			wantKeys: []string{"action.0.target_group_arn"},
		},
		// Case B: arn present in state (Old=arn), absent from config — null RawPlan path
		"CaseB_Suppressed": {
			reason: "spurious arn removal is suppressed when forward is active in state",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: testARN, New: "", NewRemoved: true},
				},
			},
			state: &terraform.InstanceState{
				Attributes: map[string]string{"action.0.forward.#": "1"},
			},
			goneKeys: []string{"action.0.target_group_arn"},
		},
		"CaseB_NotSuppressed_NoForwardInState": {
			reason: "arn removal is kept when forward count is 0 in state",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: testARN, New: "", NewRemoved: true},
				},
			},
			state: &terraform.InstanceState{
				Attributes: map[string]string{"action.0.forward.#": "0"},
			},
			wantKeys: []string{"action.0.target_group_arn"},
		},
		"CaseB_NotSuppressed_NilState": {
			reason: "arn removal is kept when state is nil",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: testARN, New: "", NewRemoved: true},
				},
			},
			state:    nil,
			wantKeys: []string{"action.0.target_group_arn"},
		},
		"RealArnUpdate_NotSuppressed": {
			reason: "a real arn change (both Old and New non-empty) is never suppressed",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: "arn:old", New: "arn:new"},
				},
			},
			cfg:      configWithForward(),
			wantKeys: []string{"action.0.target_group_arn"},
		},
		"UnrelatedKeysUntouched": {
			reason: "keys that do not match action.*.target_group_arn are left untouched",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.type":            {Old: "", New: "forward"},
					"other.0.target_group_arn": {Old: "", New: testARN},
				},
			},
			cfg:      configWithForward(),
			wantKeys: []string{"action.0.type", "other.0.target_group_arn"},
		},
		"MultipleActions_OnlySuppressesMatchingIndex": {
			reason: "when two actions are present, only the one with forward in config is suppressed",
			diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"action.0.target_group_arn": {Old: "", New: testARN},
					"action.1.target_group_arn": {Old: "", New: testARN},
				},
			},
			cfg: &terraform.ResourceConfig{
				Config: map[string]interface{}{
					"action": []interface{}{
						// action 0: only arn, no forward
						map[string]interface{}{"target_group_arn": testARN},
						// action 1: forward configured
						map[string]interface{}{
							"forward": []interface{}{
								map[string]interface{}{
									"target_group": []interface{}{
										map[string]interface{}{"arn": testARN, "weight": 1},
									},
								},
							},
						},
					},
				},
			},
			wantKeys: []string{"action.0.target_group_arn"},
			goneKeys: []string{"action.1.target_group_arn"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := lbListenerRuleCustomDiff(tc.diff, tc.state, tc.cfg)
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

func TestConfigActionHasForward(t *testing.T) {
	cases := map[string]struct {
		reason string
		cfg    *terraform.ResourceConfig
		idx    string
		want   bool
	}{
		"NilConfig": {
			reason: "nil config returns false",
			cfg:    nil,
			idx:    "0",
			want:   false,
		},
		"NilConfigMap": {
			reason: "config with nil inner map returns false",
			cfg:    &terraform.ResourceConfig{},
			idx:    "0",
			want:   false,
		},
		"NoActionsKey": {
			reason: "config without an action key returns false",
			cfg:    &terraform.ResourceConfig{Config: map[string]interface{}{}},
			idx:    "0",
			want:   false,
		},
		"InvalidIndex": {
			reason: "non-numeric index returns false",
			cfg: &terraform.ResourceConfig{
				Config: map[string]interface{}{
					"action": []interface{}{map[string]interface{}{}},
				},
			},
			idx:  "abc",
			want: false,
		},
		"OutOfRangeIndex": {
			reason: "index beyond action slice length returns false",
			cfg: &terraform.ResourceConfig{
				Config: map[string]interface{}{
					"action": []interface{}{map[string]interface{}{}},
				},
			},
			idx:  "5",
			want: false,
		},
		"NoForward": {
			reason: "action with only target_group_arn returns false",
			cfg: &terraform.ResourceConfig{
				Config: map[string]interface{}{
					"action": []interface{}{
						map[string]interface{}{"target_group_arn": testARN},
					},
				},
			},
			idx:  "0",
			want: false,
		},
		"EmptyForwardSlice": {
			reason: "action with an empty forward slice returns false",
			cfg: &terraform.ResourceConfig{
				Config: map[string]interface{}{
					"action": []interface{}{
						map[string]interface{}{"forward": []interface{}{}},
					},
				},
			},
			idx:  "0",
			want: false,
		},
		"HasForward": {
			reason: "action with a non-empty forward slice returns true",
			cfg:    configWithForward(),
			idx:    "0",
			want:   true,
		},
		"SecondActionHasForward": {
			reason: "checks the correct index when multiple actions exist",
			cfg: &terraform.ResourceConfig{
				Config: map[string]interface{}{
					"action": []interface{}{
						map[string]interface{}{"target_group_arn": testARN},
						map[string]interface{}{
							"forward": []interface{}{
								map[string]interface{}{"target_group": []interface{}{}},
							},
						},
					},
				},
			},
			idx:  "1",
			want: true,
		},
		"FirstActionNoForward_SecondHasForward": {
			reason: "index 0 returns false even when index 1 has forward",
			cfg: &terraform.ResourceConfig{
				Config: map[string]interface{}{
					"action": []interface{}{
						map[string]interface{}{"target_group_arn": testARN},
						map[string]interface{}{
							"forward": []interface{}{
								map[string]interface{}{"target_group": []interface{}{}},
							},
						},
					},
				},
			},
			idx:  "0",
			want: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := configActionHasForward(tc.cfg, tc.idx)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("%s: -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
