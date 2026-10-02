// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	"testing"
)

func TestSNSPlatformApplicationImportID(t *testing.T) {
	externalName := TerraformPluginSDKExternalNameConfigs["aws_sns_platform_application"]
	if len(externalName.IdentifierFields) != 0 {
		t.Fatalf("IdentifierFields = %v, want none", externalName.IdentifierFields)
	}
	setup := map[string]any{
		"configuration": map[string]any{
			"region": "eu-west-1",
		},
		"client_metadata": map[string]any{
			"account_id": "123456789012",
			"partition":  "aws",
		},
	}

	tests := map[string]struct {
		parameters  map[string]any
		expectedID  string
		expectedErr string
	}{
		"MissingPlatform": {
			parameters:  map[string]any{},
			expectedErr: "platform is required to build the SNS platform application import id",
		},
		"EmptyPlatform": {
			parameters:  map[string]any{"platform": ""},
			expectedErr: "platform is required to build the SNS platform application import id",
		},
		"NonStringPlatform": {
			parameters:  map[string]any{"platform": 1},
			expectedErr: "platform is required to build the SNS platform application import id",
		},
		"ADM": {
			parameters: map[string]any{"platform": "ADM"},
			expectedID: "arn:aws:sns:eu-west-1:123456789012:app/ADM/example-application",
		},
		"APNS": {
			parameters: map[string]any{"platform": "APNS"},
			expectedID: "arn:aws:sns:eu-west-1:123456789012:app/APNS/example-application",
		},
		"APNS_SANDBOX": {
			parameters: map[string]any{"platform": "APNS_SANDBOX"},
			expectedID: "arn:aws:sns:eu-west-1:123456789012:app/APNS_SANDBOX/example-application",
		},
		"GCM": {
			parameters: map[string]any{"platform": "GCM"},
			expectedID: "arn:aws:sns:eu-west-1:123456789012:app/GCM/example-application",
		},
		"WNS": {
			parameters: map[string]any{"platform": "WNS"},
			expectedID: "arn:aws:sns:eu-west-1:123456789012:app/WNS/example-application",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			id, err := externalName.GetIDFn(
				context.Background(),
				"example-application",
				tc.parameters,
				setup,
			)
			if tc.expectedErr != "" {
				if err == nil || err.Error() != tc.expectedErr {
					t.Fatalf("GetIDFn() error = %v, want %q", err, tc.expectedErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("GetIDFn() error = %v", err)
			}
			if id != tc.expectedID {
				t.Errorf("GetIDFn() = %q, want %q", id, tc.expectedID)
			}
		})
	}
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
