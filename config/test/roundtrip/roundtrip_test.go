// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// package roundtrip contains API roundtrip tests
// for multi-version managed resource APIs
package roundtrip

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"k8s.io/apimachinery/pkg/runtime"

	clusterapis "github.com/upbound/provider-aws/v2/apis/cluster"
	namespacedapis "github.com/upbound/provider-aws/v2/apis/namespaced"
	"github.com/upbound/provider-aws/v2/config"
)

func TestRoundTrip(t *testing.T) {
	fwProvider, sdkProvider, err := xpprovider.GetProvider(t.Context())
	if err != nil {
		t.Fatalf("GetProviderSchema: %s", err)
	}
	provider, err := config.GetProvider(t.Context(), fwProvider, sdkProvider, false, false)
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}

	providerNamespaced, err := config.GetProviderNamespaced(t.Context(), fwProvider, sdkProvider, false, false)
	if err != nil {
		t.Fatalf("GetNamespacedProvider: %s", err)
	}

	testScheme := runtime.NewScheme()
	if err := clusterapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("cluster-scoped apis AddToScheme: %s", err)
	}
	if err := namespacedapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("namespaced apis AddToScheme: %s", err)
	}

	allCmpOpts := make([]cmp.Option, 0, len(awsCustomCmpOpts)+2)
	allCmpOpts = append(allCmpOpts, roundtrip.EquateEmptyAndSingleZeroSlice(), roundtrip.EquateNilAndZeroValuePtr())
	allCmpOpts = append(allCmpOpts, awsCustomCmpOpts...)
	rt, err := roundtrip.NewRoundTripTest(provider, providerNamespaced, testScheme,
		roundtrip.WithFuzzerConfig(
			roundtrip.FuzzerIterations(10),
			roundtrip.FuzzerNilChance(0)),
		roundtrip.WithFuzzerConfig(
			roundtrip.FuzzerIterations(30),
			roundtrip.FuzzerNilChance(0.3)),
		roundtrip.WithComparisonOptions(
			allCmpOpts...,
		),
		roundtrip.WithExtraFuzzFuncs(awsCustomFuzzers...),
	)
	if err != nil {
		t.Fatalf("NewRoundTripTest: %s", err)
	}

	t.Run("TestSerializationRoundtrip", func(t *testing.T) {
		rt.TestSerializationRoundtrip(t)
	})

	t.Run("TestConversionRoundtrip", func(t *testing.T) {
		rt.TestConversionRoundtrip(t)
	})

}
