// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
	// "github.com/upbound/provider-aws/apis/s3/v1beta1"
	"github.com/crossplane/upjet/pkg/controller/conversion"
	"github.com/upbound/provider-aws/apis/s3/v1beta2"
	"github.com/upbound/provider-aws/apis/s3/v1beta3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var (
	//go:embed s3/v1beta2.yaml
	v1beta2Manifest []byte
)

func TestMultiversionCRDConversionsV1Beta2(t *testing.T) {
	type args struct {
		manifest []byte
	}

	type want struct {
		err *error
	}

	// Read yaml manifests from embedded files
	// Unmarshall them to apis/s3/v1beta1/BucketLifecycleConfiguration
	// Call the static upjet.controller.conversion.RoundTrip
	cases := map[string]struct {
		args
		want
	}{
		"Empty": {
			args: args{
				manifest:    v1beta2Manifest,
			},
			want: want{
				err: nil,
			},
		},
	}

	// Setup the things that need configuring
	ctx := context.TODO()
	provider, err := GetProvider(ctx, false, false)
	if err != nil {
		t.Fatalf("failed to instantiate provider: %v", err)
	}
	conversion.RegisterConversions(provider, nil)



	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var initial v1beta2.BucketLifecycleConfiguration
			err := yaml.Unmarshal(tc.args.manifest, &initial)
			if err != nil {
				t.Fatalf("failed to unmarshal spec: %v", err)
			}
			if &initial == nil {
				t.Fatalf("initial object is nil")
			}

			hub := v1beta3.BucketLifecycleConfiguration{
				TypeMeta: v1.TypeMeta{
					Kind: initial.Kind,
					APIVersion: "s3.aws.upbound.io/v1beta3",
				},
			}
			err = initial.ConvertTo(&hub)
			if err != nil {
				t.Fatalf("failed to convert to hub: %v", err)
			}

			final := v1beta2.BucketLifecycleConfiguration{
				TypeMeta: v1.TypeMeta{
					Kind: initial.Kind,
					APIVersion: "s3.aws.upbound.io/v1beta2",
				},
			}
			err = final.ConvertFrom(&hub)
			if err != nil {
				t.Fatalf("failed to convert from hub: %v", err)
			}

			if diff := cmp.Diff(initial.Spec, final.Spec); diff != "" {
				t.Errorf("Round-trip conversion mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
