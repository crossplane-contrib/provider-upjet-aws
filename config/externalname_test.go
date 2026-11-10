// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	clustercognitoidp "github.com/upbound/provider-aws/v2/config/cluster/cognitoidp"
	clustercommon "github.com/upbound/provider-aws/v2/config/cluster/common"
	namespacedcognitoidp "github.com/upbound/provider-aws/v2/config/namespaced/cognitoidp"
	namespacedcommon "github.com/upbound/provider-aws/v2/config/namespaced/common"
)

func TestCognitoLogDeliveryConfigurationReferences(t *testing.T) {
	cases := map[string]struct {
		configure func(*ujconfig.Provider)
		extractor string
	}{
		"Cluster": {
			configure: clustercognitoidp.Configure,
			extractor: clustercommon.PathARNExtractor,
		},
		"Namespaced": {
			configure: namespacedcognitoidp.Configure,
			extractor: namespacedcommon.PathARNExtractor,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			const resourceName = "aws_cognito_log_delivery_configuration"
			p := ujconfig.NewProvider([]byte(providerSchema), "aws", "github.com/upbound/provider-aws/v2", providerMetadata,
				ujconfig.WithIncludeList([]string{resourceName + "$"}),
			)
			tc.configure(p)
			p.ConfigureResources()
			r, ok := p.Resources[resourceName]
			if !ok {
				t.Fatal("log delivery configuration resource was not included")
			}
			want := ujconfig.References{
				"user_pool_id": {
					TerraformName: "aws_cognito_user_pool",
				},
				"log_configurations.cloud_watch_logs_configuration.log_group_arn": {
					TerraformName: "aws_cloudwatch_log_group",
					Extractor:     tc.extractor,
				},
				"log_configurations.firehose_configuration.stream_arn": {
					TerraformName: "aws_kinesis_firehose_delivery_stream",
					Extractor:     tc.extractor,
				},
				"log_configurations.s3_configuration.bucket_arn": {
					TerraformName: "aws_s3_bucket",
					Extractor:     tc.extractor,
				},
			}
			if diff := cmp.Diff(want, r.References); diff != "" {
				t.Errorf("References mismatch (-want +got):\n%s", diff)
			}
			wantTFPaths := []string{
				"log_configurations[*].cloud_watch_logs_configuration",
				"log_configurations[*].firehose_configuration",
				"log_configurations[*].s3_configuration",
			}
			sortPaths := cmpopts.SortSlices(func(a, b string) bool { return a < b })
			if diff := cmp.Diff(wantTFPaths, r.TFListConversionPaths(), sortPaths); diff != "" {
				t.Errorf("Terraform singleton conversion paths mismatch (-want +got):\n%s", diff)
			}
			wantCRDPaths := []string{
				"logConfigurations[*].cloudWatchLogsConfiguration",
				"logConfigurations[*].firehoseConfiguration",
				"logConfigurations[*].s3Configuration",
			}
			if diff := cmp.Diff(wantCRDPaths, r.CRDListConversionPaths(), sortPaths); diff != "" {
				t.Errorf("CRD singleton conversion paths mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCognitoLogDeliveryConfigurationExternalName(t *testing.T) {
	e := TerraformPluginFrameworkExternalNameConfigs["aws_cognito_log_delivery_configuration"]
	if len(e.OmittedFields) != 0 {
		t.Fatalf("user_pool_id must remain configurable, got omitted fields %v", e.OmittedFields)
	}
	const poolID = "us-west-2_example123"
	t.Run("CreateWithReference", func(t *testing.T) {
		parameters := map[string]any{"user_pool_id": poolID}
		e.SetIdentifierArgumentFn(parameters, "")
		if parameters["user_pool_id"] != poolID {
			t.Fatalf("SetIdentifierArgumentFn() changed the referenced user_pool_id: %v", parameters)
		}
	})
	t.Run("ImportWithExternalName", func(t *testing.T) {
		parameters := map[string]any{}
		e.SetIdentifierArgumentFn(parameters, poolID)
		if parameters["user_pool_id"] != poolID {
			t.Fatalf("SetIdentifierArgumentFn() = %v, want user_pool_id %q", parameters, poolID)
		}
	})
	t.Run("PreserveConfiguredPool", func(t *testing.T) {
		parameters := map[string]any{"user_pool_id": poolID}
		e.SetIdentifierArgumentFn(parameters, "us-west-2_other")
		if parameters["user_pool_id"] != poolID {
			t.Fatalf("SetIdentifierArgumentFn() overwrote the configured pool: %v", parameters)
		}
	})
	t.Run("ReadExternalName", func(t *testing.T) {
		got, err := e.GetExternalNameFn(map[string]any{"user_pool_id": poolID})
		if err != nil || got != poolID {
			t.Fatalf("GetExternalNameFn() = %q, %v, want %q, nil", got, err, poolID)
		}
	})
	t.Run("MissingPoolInState", func(t *testing.T) {
		if _, err := e.GetExternalNameFn(map[string]any{}); err == nil {
			t.Fatal("GetExternalNameFn() must report a missing user_pool_id")
		}
	})
	t.Run("NoSyntheticID", func(t *testing.T) {
		got, err := e.GetIDFn(context.Background(), poolID, map[string]any{"user_pool_id": poolID}, nil)
		if err != nil || got != "" {
			t.Fatalf("GetIDFn() = %q, %v, want empty ID for a resource without an id attribute", got, err)
		}
	})
}

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
