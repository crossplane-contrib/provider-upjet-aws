// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bedrock

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// TestInferenceProfileAdditionalConnectionDetails runs the real Configure
// function from this package (not a copy of it) against a bare resource for
// aws_bedrock_inference_profile, then calls the AdditionalConnectionDetailsFn
// it wires up exactly the way upjet's GetConnectionDetails does. No
// Terraform schema / codegen pipeline needed: AdditionalConnectionDetailsFn
// is read at runtime (see sensitive.go in upjet), so none of this depends on
// generated code.
func TestInferenceProfileAdditionalConnectionDetails(t *testing.T) {
	minimalSchema := []byte(`{"format_version":"1.0","provider_schemas":{"registry.terraform.io/hashicorp/aws":{}}}`)
	p := config.NewProvider(minimalSchema, "aws", "github.com/upbound/provider-aws/v2", nil)
	p.Resources["aws_bedrock_inference_profile"] = &config.Resource{Name: "aws_bedrock_inference_profile"}

	Configure(p)
	p.ConfigureResources()

	fn := p.Resources["aws_bedrock_inference_profile"].Sensitive.AdditionalConnectionDetailsFn
	if fn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set on aws_bedrock_inference_profile, got nil")
	}

	cases := map[string]struct {
		attr map[string]any
		want map[string][]byte
	}{
		"ARNPresent": {
			attr: map[string]any{
				"arn":    "arn:aws:bedrock:ap-south-1:719723331442:application-inference-profile/82u9dc4lm637",
				"status": "ACTIVE",
			},
			want: map[string][]byte{
				"arn": []byte("arn:aws:bedrock:ap-south-1:719723331442:application-inference-profile/82u9dc4lm637"),
			},
		},
		"ARNMissing": {
			attr: map[string]any{"status": "ACTIVE"},
			want: map[string][]byte{},
		},
		"ARNWrongType": {
			attr: map[string]any{"arn": 12345},
			want: map[string][]byte{},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := fn(tc.attr)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("got %d connection detail(s), want %d: %v", len(got), len(tc.want), got)
			}
			for k, v := range tc.want {
				gv, ok := got[k]
				if !ok {
					t.Fatalf("missing expected key %q in result %v", k, got)
				}
				if string(gv) != string(v) {
					t.Fatalf("key %q: got %q, want %q", k, gv, v)
				}
			}
		})
	}
}
