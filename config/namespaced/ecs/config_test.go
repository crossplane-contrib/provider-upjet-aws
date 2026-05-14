// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ecs

import (
	"testing"
)

func TestTaskDefinitionSetIdentifierArgumentFn(t *testing.T) {
	setIDFn := func(base map[string]any, externalName string) {
		if arn, _ := base["arn"].(string); arn == "" && externalName != "" {
			base["arn"] = externalName
		}
	}

	cases := map[string]struct {
		base         map[string]any
		externalName string
		wantArn      string
	}{
		"ColdStartObserveWithFullARN": {
			base:         map[string]any{},
			externalName: "arn:aws:ecs:us-east-1:123456789012:task-definition/my-service:7",
			wantArn:      "arn:aws:ecs:us-east-1:123456789012:task-definition/my-service:7",
		},
		"ColdStartObserveWithFamilyRevision": {
			base:         map[string]any{},
			externalName: "my-service:7",
			wantArn:      "my-service:7",
		},
		"WarmStartArnAlreadyPopulated": {
			base:         map[string]any{"arn": "arn:aws:ecs:us-east-1:123456789012:task-definition/my-service:7"},
			externalName: "my-service:7",
			wantArn:      "arn:aws:ecs:us-east-1:123456789012:task-definition/my-service:7",
		},
		"EmptyExternalName": {
			base:         map[string]any{},
			externalName: "",
			wantArn:      "",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			setIDFn(tc.base, tc.externalName)
			got, _ := tc.base["arn"].(string)
			if got != tc.wantArn {
				t.Errorf("base[\"arn\"] = %q, want %q", got, tc.wantArn)
			}
		})
	}
}
