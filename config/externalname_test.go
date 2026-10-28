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

func TestGrafanaWorkspaceServiceAccountGetIDFn(t *testing.T) {
	e := grafanaWorkspaceServiceAccount()

	cases := map[string]struct {
		parameters map[string]any
		wantID     string
		wantErr    bool
	}{
		"ColdStartWithPlaceholder": {
			parameters: map[string]any{"workspace_id": "g-abc12345", "service_account_id": "1"},
			wantID:     "g-abc12345,1",
		},
		"CreatedWithRealServiceAccountID": {
			parameters: map[string]any{"workspace_id": "g-abc12345", "service_account_id": "3"},
			wantID:     "g-abc12345,3",
		},
		"MissingWorkspaceID": {
			parameters: map[string]any{"service_account_id": "1"},
			wantErr:    true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := e.GetIDFn(context.Background(), "", tc.parameters, nil)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("GetIDFn() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("GetIDFn() error = %v, want nil", err)
			}
			if got != tc.wantID {
				t.Errorf("GetIDFn() = %q, want %q", got, tc.wantID)
			}
		})
	}
}

func TestGrafanaWorkspaceServiceAccountTokenGetIDFn(t *testing.T) {
	e := grafanaWorkspaceServiceAccountToken()

	cases := map[string]struct {
		parameters map[string]any
		wantID     string
		wantErr    bool
	}{
		"ColdStartWithPlaceholder": {
			parameters: map[string]any{"workspace_id": "g-abc12345", "service_account_id": "1", "service_account_token_id": "1"},
			wantID:     "g-abc12345,1,1",
		},
		"CreatedWithRealTokenID": {
			parameters: map[string]any{"workspace_id": "g-abc12345", "service_account_id": "1", "service_account_token_id": "7"},
			wantID:     "g-abc12345,1,7",
		},
		"MissingServiceAccountID": {
			parameters: map[string]any{"workspace_id": "g-abc12345", "service_account_token_id": "1"},
			wantErr:    true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := e.GetIDFn(context.Background(), "", tc.parameters, nil)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("GetIDFn() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("GetIDFn() error = %v, want nil", err)
			}
			if got != tc.wantID {
				t.Errorf("GetIDFn() = %q, want %q", got, tc.wantID)
			}
		})
	}
}

func TestGlueCatalogTableOptimizerExternalName(t *testing.T) {
	e := TerraformPluginFrameworkExternalNameConfigs["aws_glue_catalog_table_optimizer"]
	tfState := map[string]any{
		"catalog_id":    "123456789012",
		"database_name": "example_db",
		"table_name":    "example_table",
		"type":          "retention",
	}

	cases := map[string]struct {
		tfState  map[string]any
		wantName string
		wantErr  bool
	}{
		"StateWithoutID": {
			tfState:  tfState,
			wantName: "123456789012,example_db,example_table,retention",
		},
		"MissingType": {
			tfState: map[string]any{"catalog_id": "123456789012", "database_name": "example_db", "table_name": "example_table"},
			wantErr: true,
		},
		"EmptyTableName": {
			tfState: map[string]any{"catalog_id": "123456789012", "database_name": "example_db", "table_name": "", "type": "retention"},
			wantErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := e.GetExternalNameFn(tc.tfState)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("GetExternalNameFn() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("GetExternalNameFn() error = %v, want nil", err)
			}
			if got != tc.wantName {
				t.Errorf("GetExternalNameFn() = %q, want %q", got, tc.wantName)
			}
		})
	}

	t.Run("NoIDInState", func(t *testing.T) {
		id, err := e.GetIDFn(context.Background(), "123456789012,example_db,example_table,retention", tfState, nil)
		if err != nil {
			t.Fatalf("GetIDFn() error = %v, want nil", err)
		}
		if id != "" {
			t.Errorf("GetIDFn() = %q, want empty: the resource has no id attribute", id)
		}
	})
}

func TestOpensearchAuthorizeVPCEndpointAccess(t *testing.T) {
	e := opensearchAuthorizeVPCEndpointAccess()

	t.Run("GetExternalNameFn", func(t *testing.T) {
		cases := map[string]struct {
			tfstate map[string]any
			want    string
			wantErr bool
		}{
			"Composite": {
				tfstate: map[string]any{"domain_name": "my-domain", "account": "123456789012"},
				want:    "my-domain,123456789012",
			},
			"MissingAccount": {
				tfstate: map[string]any{"domain_name": "my-domain"},
				wantErr: true,
			},
			"EmptyDomainName": {
				tfstate: map[string]any{"domain_name": "", "account": "123456789012"},
				wantErr: true,
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				got, err := e.GetExternalNameFn(tc.tfstate)
				if tc.wantErr {
					if err == nil {
						t.Fatalf("GetExternalNameFn() error = nil, want error")
					}
					return
				}
				if err != nil {
					t.Fatalf("GetExternalNameFn() error = %v, want nil", err)
				}
				if got != tc.want {
					t.Errorf("GetExternalNameFn() = %q, want %q", got, tc.want)
				}
			})
		}
	})

	t.Run("SetIdentifierArgumentFn", func(t *testing.T) {
		cases := map[string]struct {
			base         map[string]any
			externalName string
			want         map[string]any
		}{
			"FillsMissingKeys": {
				base:         map[string]any{},
				externalName: "my-domain,123456789012",
				want:         map[string]any{"domain_name": "my-domain", "account": "123456789012"},
			},
			"KeepsSpecValues": {
				base:         map[string]any{"domain_name": "spec-domain"},
				externalName: "my-domain,123456789012",
				want:         map[string]any{"domain_name": "spec-domain", "account": "123456789012"},
			},
			"EmptyExternalName": {
				base:         map[string]any{},
				externalName: "",
				want:         map[string]any{},
			},
			"MalformedExternalName": {
				base:         map[string]any{},
				externalName: "my-domain",
				want:         map[string]any{},
			},
		}
		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				e.SetIdentifierArgumentFn(tc.base, tc.externalName)
				if len(tc.base) != len(tc.want) {
					t.Fatalf("base = %v, want %v", tc.base, tc.want)
				}
				for k, v := range tc.want {
					if tc.base[k] != v {
						t.Errorf("base[%q] = %v, want %v", k, tc.base[k], v)
					}
				}
			})
		}
	})
}

func TestCloudFrontPolicyIdentifierGetIDFn(t *testing.T) {
	tests := map[string]struct {
		externalName string
		parameters   map[string]any
		want         string
	}{
		"ExternalNameWins": {
			externalName: "658327ea-f89d-4fab-a63d-7e88639e58f6",
			parameters:   map[string]any{"name": "Managed-CachingOptimized"},
			want:         "658327ea-f89d-4fab-a63d-7e88639e58f6",
		},
		"FallsBackToName": {
			parameters: map[string]any{"name": "Managed-CachingOptimized"},
			want:       "Managed-CachingOptimized",
		},
		"NeitherSet": {
			parameters: map[string]any{},
			want:       "",
		},
	}
	for _, r := range []string{"aws_cloudfront_cache_policy", "aws_cloudfront_origin_request_policy", "aws_cloudfront_response_headers_policy"} {
		e := TerraformPluginSDKExternalNameConfigs[r]
		for name, tc := range tests {
			t.Run(r+"/"+name, func(t *testing.T) {
				got, err := e.GetIDFn(context.Background(), tc.externalName, tc.parameters, nil)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got != tc.want {
					t.Errorf("GetIDFn() = %q, want %q", got, tc.want)
				}
			})
		}
	}
}
