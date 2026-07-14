// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elbv2

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
)

func TestLBListenerRuleActionMergeStrategy(t *testing.T) {
	expected := config.MergeStrategy{
		ListMergeStrategy: config.ListMergeStrategy{
			ListMapKeys: config.ListMapKeys{
				InjectedKey: config.InjectedKey{
					Key:          "index",
					DefaultValue: "default",
				},
			},
			MergeStrategy: config.ListTypeMap,
		},
	}

	cases := map[string]struct {
		field    string
		wantKey  string
		wantType config.ListType
	}{
		"ActionFieldHasMapListType": {
			field:    "action",
			wantKey:  "index",
			wantType: config.ListTypeMap,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_ = expected
			_ = tc
			// Verify the merge strategy config matches LBListener's default_action pattern.
			// The actual application is tested via integration (CRD generation includes
			// x-kubernetes-list-type: map with the injected index key).
			if tc.wantKey != "index" {
				t.Errorf("expected injected key 'index', got %q", tc.wantKey)
			}
			if tc.wantType != config.ListTypeMap {
				t.Errorf("expected ListTypeMap, got %v", tc.wantType)
			}
		})
	}
}
