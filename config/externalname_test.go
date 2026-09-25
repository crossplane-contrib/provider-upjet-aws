// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	"testing"
)

func TestRdsIntegration(t *testing.T) {
	const stubARN = "arn:aws:rds:us-west-2:000000000000:integration"

	e := rdsIntegration()

	t.Run("SetIdentifierArgumentFn", func(t *testing.T) {
		cases := map[string]struct {
			base         map[string]any
			externalName string
			wantARN      string
		}{
			"ColdStartWithRegion": {
				base:         map[string]any{"region": "us-west-2"},
				externalName: "",
				wantARN:      stubARN,
			},
			"ColdStartWithoutRegion": {
				base:         map[string]any{},
				externalName: "",
				wantARN:      "",
			},
			"WarmStartWithARN": {
				base:         map[string]any{},
				externalName: "arn:aws:rds:us-east-1:123456789012:integration:12345678-1234-1234-1234-123456789012",
				wantARN:      "arn:aws:rds:us-east-1:123456789012:integration:12345678-1234-1234-1234-123456789012",
			},
			"WarmStartARNAlreadyPopulated": {
				base:         map[string]any{"arn": "arn:aws:rds:us-east-1:123456789012:integration:12345678-1234-1234-1234-123456789012"},
				externalName: "some-other-name",
				wantARN:      "arn:aws:rds:us-east-1:123456789012:integration:12345678-1234-1234-1234-123456789012",
			},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				e.SetIdentifierArgumentFn(tc.base, tc.externalName)
				got, _ := tc.base["arn"].(string)
				if got != tc.wantARN {
					t.Errorf("base[\"arn\"] = %q, want %q", got, tc.wantARN)
				}
			})
		}
	})

	t.Run("GetIDFn", func(t *testing.T) {
		cases := map[string]struct {
			externalName string
			parameters   map[string]any
			want         string
		}{
			"ColdStartWithRegion": {
				externalName: "",
				parameters:   map[string]any{"region": "us-west-2"},
				want:         stubARN,
			},
			"WarmStartWithARN": {
				externalName: "arn:aws:rds:us-east-1:123456789012:integration:12345678-1234-1234-1234-123456789012",
				parameters:   map[string]any{"region": "us-east-1"},
				want:         "arn:aws:rds:us-east-1:123456789012:integration:12345678-1234-1234-1234-123456789012",
			},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				got, err := e.GetIDFn(context.Background(), tc.externalName, tc.parameters, nil)
				if err != nil {
					t.Fatalf("GetIDFn returned an error: %v", err)
				}
				if got != tc.want {
					t.Errorf("GetIDFn = %q, want %q", got, tc.want)
				}
			})
		}
	})

	t.Run("GetExternalNameFn", func(t *testing.T) {
		tfState := map[string]any{"arn": "arn:aws:rds:us-east-1:123456789012:integration:12345678-1234-1234-1234-123456789012"}
		got, err := e.GetExternalNameFn(tfState)
		if err != nil {
			t.Fatalf("GetExternalNameFn returned an error: %v", err)
		}
		if want := "arn:aws:rds:us-east-1:123456789012:integration:12345678-1234-1234-1234-123456789012"; got != want {
			t.Errorf("GetExternalNameFn = %q, want %q", got, want)
		}
	})
}

func TestEcsTaskDefinitionSetIdentifierArgumentFn(t *testing.T) {
	e := ecsTaskDefinition()

	cases := map[string]struct {
		base         map[string]any
		externalName string
		wantArn      string
	}{
		"ColdStartWithFullARN": {
			base:         map[string]any{},
			externalName: "arn:aws:ecs:us-east-1:123456789012:task-definition/my-service:7",
			wantArn:      "arn:aws:ecs:us-east-1:123456789012:task-definition/my-service:7",
		},
		"ColdStartWithFamilyRevision": {
			base:         map[string]any{},
			externalName: "my-service:7",
			wantArn:      "",
		},
		"ColdStartWithFamilyOnly": {
			base:         map[string]any{},
			externalName: "my-service",
			wantArn:      "",
		},
		"EmptyExternalName": {
			base:         map[string]any{},
			externalName: "",
			wantArn:      "",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e.SetIdentifierArgumentFn(tc.base, tc.externalName)
			got, _ := tc.base["arn"].(string)
			if got != tc.wantArn {
				t.Errorf("base[\"arn\"] = %q, want %q", got, tc.wantArn)
			}
		})
	}
}
