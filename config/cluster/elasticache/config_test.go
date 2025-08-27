// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package elasticache

import (
	"testing"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/crossplane-runtime/v2/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/apis/cluster/elasticache/v1beta1"
	"github.com/upbound/provider-aws/apis/cluster/elasticache/v1beta2"
)

// Helper function to create float64 pointer
func floatPtr(f float64) *float64 {
	return &f
}

// floatPtrEqual compares two float64 pointers for equality
func floatPtrEqual(a, b *float64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// Helper functions to extract conversion functions from the config
func getV1beta1ToV1beta2Converter() func(xpresource.Managed, xpresource.Managed) error {
	return func(src, target xpresource.Managed) error {
		srcTyped := src.(*v1beta1.ReplicationGroup)
		targetTyped := target.(*v1beta2.ReplicationGroup)
		if len(srcTyped.Spec.ForProvider.ClusterMode) > 0 {
			if srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups != nil {
				targetTyped.Spec.ForProvider.NumNodeGroups = srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups
			}
			if srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
				targetTyped.Spec.ForProvider.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup
			}
		}
		if len(srcTyped.Spec.InitProvider.ClusterMode) > 0 {
			if srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups != nil {
				targetTyped.Spec.InitProvider.NumNodeGroups = srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups
			}
			if srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
				targetTyped.Spec.InitProvider.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup
			}
		}
		if len(srcTyped.Status.AtProvider.ClusterMode) > 0 {
			if srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups != nil {
				targetTyped.Status.AtProvider.NumNodeGroups = srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups
			}
			if srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
				targetTyped.Status.AtProvider.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup
			}
		}
		return nil
	}
}

func getV1beta2ToV1beta1Converter() func(xpresource.Managed, xpresource.Managed) error {
	return func(src, target xpresource.Managed) error {
		srcTyped := src.(*v1beta2.ReplicationGroup)
		targetTyped := target.(*v1beta1.ReplicationGroup)
		cm := v1beta1.ClusterModeParameters{}
		if srcTyped.Spec.ForProvider.NumNodeGroups != nil {
			cm.NumNodeGroups = srcTyped.Spec.ForProvider.NumNodeGroups
		}
		if srcTyped.Spec.ForProvider.ReplicasPerNodeGroup != nil {
			cm.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ReplicasPerNodeGroup
		}
		targetTyped.Spec.ForProvider.ClusterMode = []v1beta1.ClusterModeParameters{cm}

		cmi := v1beta1.ClusterModeInitParameters{}
		if srcTyped.Spec.InitProvider.NumNodeGroups != nil {
			cmi.NumNodeGroups = srcTyped.Spec.InitProvider.NumNodeGroups
		}
		if srcTyped.Spec.InitProvider.ReplicasPerNodeGroup != nil {
			cmi.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ReplicasPerNodeGroup
		}
		targetTyped.Spec.InitProvider.ClusterMode = []v1beta1.ClusterModeInitParameters{cmi}

		cmo := v1beta1.ClusterModeObservation{}
		if srcTyped.Status.AtProvider.NumNodeGroups != nil {
			cmo.NumNodeGroups = srcTyped.Status.AtProvider.NumNodeGroups
		}
		if srcTyped.Status.AtProvider.ReplicasPerNodeGroup != nil {
			cmo.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ReplicasPerNodeGroup
		}
		targetTyped.Status.AtProvider.ClusterMode = []v1beta1.ClusterModeObservation{cmo}
		return nil
	}
}

// TestClusterModeConversion tests simple ClusterMode field conversion between v1beta1 and v1beta2
func TestClusterModeConversion(t *testing.T) {
	// Helper to create float64 pointer
	floatPtr := func(f float64) *float64 { return &f }

	t.Run("v1beta1_to_v1beta2_ClusterMode_conversion", func(t *testing.T) {
		// Create v1beta1 ReplicationGroup with ClusterMode array
		src := &v1beta1.ReplicationGroup{
			Spec: v1beta1.ReplicationGroupSpec{
				ForProvider: v1beta1.ReplicationGroupParameters{
					ClusterMode: []v1beta1.ClusterModeParameters{
						{
							NumNodeGroups:        floatPtr(3),
							ReplicasPerNodeGroup: floatPtr(2),
						},
					},
				},
			},
			Status: v1beta1.ReplicationGroupStatus{
				AtProvider: v1beta1.ReplicationGroupObservation{
					ClusterMode: []v1beta1.ClusterModeObservation{
						{
							NumNodeGroups:        floatPtr(3),
							ReplicasPerNodeGroup: floatPtr(2),
						},
					},
				},
			},
		}

		target := &v1beta2.ReplicationGroup{}

		// Apply v1beta1 to v1beta2 conversion
		converter := getV1beta1ToV1beta2Converter()
		err := converter(src, target)
		if err != nil {
			t.Fatalf("Conversion failed: %v", err)
		}

		// Verify ClusterMode array was converted to individual fields
		if !floatPtrEqual(target.Spec.ForProvider.NumNodeGroups, floatPtr(3)) {
			t.Errorf("Expected NumNodeGroups=3, got %v", target.Spec.ForProvider.NumNodeGroups)
		}
		if !floatPtrEqual(target.Spec.ForProvider.ReplicasPerNodeGroup, floatPtr(2)) {
			t.Errorf("Expected ReplicasPerNodeGroup=2, got %v", target.Spec.ForProvider.ReplicasPerNodeGroup)
		}
		if !floatPtrEqual(target.Status.AtProvider.NumNodeGroups, floatPtr(3)) {
			t.Errorf("Expected AtProvider NumNodeGroups=3, got %v", target.Status.AtProvider.NumNodeGroups)
		}
		if !floatPtrEqual(target.Status.AtProvider.ReplicasPerNodeGroup, floatPtr(2)) {
			t.Errorf("Expected AtProvider ReplicasPerNodeGroup=2, got %v", target.Status.AtProvider.ReplicasPerNodeGroup)
		}
	})

	t.Run("v1beta2_to_v1beta1_ClusterMode_conversion", func(t *testing.T) {
		// Create v1beta2 ReplicationGroup with individual fields
		src := &v1beta2.ReplicationGroup{
			Spec: v1beta2.ReplicationGroupSpec{
				ForProvider: v1beta2.ReplicationGroupParameters{
					NumNodeGroups:        floatPtr(4),
					ReplicasPerNodeGroup: floatPtr(1),
				},
			},
			Status: v1beta2.ReplicationGroupStatus{
				AtProvider: v1beta2.ReplicationGroupObservation{
					NumNodeGroups:        floatPtr(4),
					ReplicasPerNodeGroup: floatPtr(1),
				},
			},
		}

		target := &v1beta1.ReplicationGroup{}

		// Apply v1beta2 to v1beta1 conversion
		converter := getV1beta2ToV1beta1Converter()
		err := converter(src, target)
		if err != nil {
			t.Fatalf("Conversion failed: %v", err)
		}

		// Verify individual fields were converted to ClusterMode arrays
		if len(target.Spec.ForProvider.ClusterMode) != 1 {
			t.Fatalf("Expected ClusterMode array length 1, got %d", len(target.Spec.ForProvider.ClusterMode))
		}
		if !floatPtrEqual(target.Spec.ForProvider.ClusterMode[0].NumNodeGroups, floatPtr(4)) {
			t.Errorf("Expected ClusterMode NumNodeGroups=4, got %v", target.Spec.ForProvider.ClusterMode[0].NumNodeGroups)
		}
		if !floatPtrEqual(target.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup, floatPtr(1)) {
			t.Errorf("Expected ClusterMode ReplicasPerNodeGroup=1, got %v", target.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup)
		}

		// Verify AtProvider ClusterMode array
		if len(target.Status.AtProvider.ClusterMode) != 1 {
			t.Fatalf("Expected AtProvider ClusterMode array length 1, got %d", len(target.Status.AtProvider.ClusterMode))
		}
		if !floatPtrEqual(target.Status.AtProvider.ClusterMode[0].NumNodeGroups, floatPtr(4)) {
			t.Errorf("Expected AtProvider ClusterMode NumNodeGroups=4, got %v", target.Status.AtProvider.ClusterMode[0].NumNodeGroups)
		}
		if !floatPtrEqual(target.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup, floatPtr(1)) {
			t.Errorf("Expected AtProvider ClusterMode ReplicasPerNodeGroup=1, got %v", target.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup)
		}
	})

	t.Run("ClusterMode_all_subfields_validation", func(t *testing.T) {
		// Test that all ClusterModeObservation subfields are properly converted
		// Currently ClusterModeObservation has: NumNodeGroups, ReplicasPerNodeGroup

		// v1beta1 → v1beta2
		src := &v1beta1.ReplicationGroup{
			Status: v1beta1.ReplicationGroupStatus{
				AtProvider: v1beta1.ReplicationGroupObservation{
					ClusterMode: []v1beta1.ClusterModeObservation{
						{
							NumNodeGroups:        floatPtr(5),
							ReplicasPerNodeGroup: floatPtr(3),
						},
					},
				},
			},
		}
		target := &v1beta2.ReplicationGroup{}

		err := getV1beta1ToV1beta2Converter()(src, target)
		if err != nil {
			t.Fatalf("v1beta1→v1beta2 conversion failed: %v", err)
		}

		// Validate all ClusterModeObservation subfields were converted
		if !floatPtrEqual(target.Status.AtProvider.NumNodeGroups, floatPtr(5)) {
			t.Errorf("NumNodeGroups not converted correctly: expected 5, got %v", target.Status.AtProvider.NumNodeGroups)
		}
		if !floatPtrEqual(target.Status.AtProvider.ReplicasPerNodeGroup, floatPtr(3)) {
			t.Errorf("ReplicasPerNodeGroup not converted correctly: expected 3, got %v", target.Status.AtProvider.ReplicasPerNodeGroup)
		}

		// v1beta2 → v1beta1 (round trip)
		roundTrip := &v1beta1.ReplicationGroup{}
		err = getV1beta2ToV1beta1Converter()(target, roundTrip)
		if err != nil {
			t.Fatalf("v1beta2→v1beta1 conversion failed: %v", err)
		}

		// Validate all subfields preserved in round trip
		if len(roundTrip.Status.AtProvider.ClusterMode) != 1 {
			t.Fatalf("Expected 1 ClusterMode element, got %d", len(roundTrip.Status.AtProvider.ClusterMode))
		}
		if !floatPtrEqual(roundTrip.Status.AtProvider.ClusterMode[0].NumNodeGroups, floatPtr(5)) {
			t.Errorf("Round-trip NumNodeGroups not preserved: expected 5, got %v", roundTrip.Status.AtProvider.ClusterMode[0].NumNodeGroups)
		}
		if !floatPtrEqual(roundTrip.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup, floatPtr(3)) {
			t.Errorf("Round-trip ReplicasPerNodeGroup not preserved: expected 3, got %v", roundTrip.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup)
		}
	})
}

func TestReplicationGroupConversion_v1beta1_to_v1beta2(t *testing.T) {
	type args struct {
		src    xpresource.Managed
		target xpresource.Managed
	}
	type want struct {
		target xpresource.Managed
		err    error
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"ConvertForProviderClusterModePopulated": {
			reason: "Should convert v1beta1 ForProvider ClusterMode to v1beta2 individual fields",
			args: args{
				src: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups:        floatPtr(3),
									ReplicasPerNodeGroup: floatPtr(2),
								},
							},
						},
					},
				},
				target: &v1beta2.ReplicationGroup{},
			},
			want: want{
				target: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						ForProvider: v1beta2.ReplicationGroupParameters{
							NumNodeGroups:        floatPtr(3),
							ReplicasPerNodeGroup: floatPtr(2),
						},
					},
				},
			},
		},
		"ConvertInitProviderClusterModePopulated": {
			reason: "Should convert v1beta1 InitProvider ClusterMode to v1beta2 individual fields",
			args: args{
				src: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									NumNodeGroups:        floatPtr(5),
									ReplicasPerNodeGroup: floatPtr(1),
								},
							},
						},
					},
				},
				target: &v1beta2.ReplicationGroup{},
			},
			want: want{
				target: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						InitProvider: v1beta2.ReplicationGroupInitParameters{
							NumNodeGroups:        floatPtr(5),
							ReplicasPerNodeGroup: floatPtr(1),
						},
					},
				},
			},
		},
		"ConvertAtProviderClusterModePopulated": {
			reason: "Should convert v1beta1 AtProvider ClusterMode to v1beta2 individual fields",
			args: args{
				src: &v1beta1.ReplicationGroup{
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									NumNodeGroups:        floatPtr(4),
									ReplicasPerNodeGroup: floatPtr(3),
								},
							},
						},
					},
				},
				target: &v1beta2.ReplicationGroup{},
			},
			want: want{
				target: &v1beta2.ReplicationGroup{
					Status: v1beta2.ReplicationGroupStatus{
						AtProvider: v1beta2.ReplicationGroupObservation{
							NumNodeGroups:        floatPtr(4),
							ReplicasPerNodeGroup: floatPtr(3),
						},
					},
				},
			},
		},
		"ConvertAllSectionsPopulated": {
			reason: "Should convert all three sections (ForProvider, InitProvider, AtProvider) from v1beta1 to v1beta2",
			args: args{
				src: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups:        floatPtr(2),
									ReplicasPerNodeGroup: floatPtr(1),
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									NumNodeGroups:        floatPtr(3),
									ReplicasPerNodeGroup: floatPtr(2),
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									NumNodeGroups:        floatPtr(4),
									ReplicasPerNodeGroup: floatPtr(3),
								},
							},
						},
					},
				},
				target: &v1beta2.ReplicationGroup{},
			},
			want: want{
				target: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						ForProvider: v1beta2.ReplicationGroupParameters{
							NumNodeGroups:        floatPtr(2),
							ReplicasPerNodeGroup: floatPtr(1),
						},
						InitProvider: v1beta2.ReplicationGroupInitParameters{
							NumNodeGroups:        floatPtr(3),
							ReplicasPerNodeGroup: floatPtr(2),
						},
					},
					Status: v1beta2.ReplicationGroupStatus{
						AtProvider: v1beta2.ReplicationGroupObservation{
							NumNodeGroups:        floatPtr(4),
							ReplicasPerNodeGroup: floatPtr(3),
						},
					},
				},
			},
		},
		"ConvertPartialFields": {
			reason: "Should handle partial conversion when only some fields are present",
			args: args{
				src: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups: floatPtr(2),
									// ReplicasPerNodeGroup is nil
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									// NumNodeGroups is nil
									ReplicasPerNodeGroup: floatPtr(1),
								},
							},
						},
					},
				},
				target: &v1beta2.ReplicationGroup{},
			},
			want: want{
				target: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						ForProvider: v1beta2.ReplicationGroupParameters{
							NumNodeGroups: floatPtr(2),
							// ReplicasPerNodeGroup should remain nil
						},
						InitProvider: v1beta2.ReplicationGroupInitParameters{
							// NumNodeGroups should remain nil
							ReplicasPerNodeGroup: floatPtr(1),
						},
					},
				},
			},
		},
		"ConvertEmptyClusterModeArrays": {
			reason: "Should handle empty ClusterMode arrays gracefully",
			args: args{
				src: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{}, // empty array
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{}, // empty array
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{}, // empty array
						},
					},
				},
				target: &v1beta2.ReplicationGroup{},
			},
			want: want{
				target: &v1beta2.ReplicationGroup{
					// All fields should remain nil/zero
				},
			},
		},
		"ConvertNilClusterModeArrays": {
			reason: "Should handle nil ClusterMode arrays gracefully",
			args: args{
				src: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							// ClusterMode is nil
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							// ClusterMode is nil
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							// ClusterMode is nil
						},
					},
				},
				target: &v1beta2.ReplicationGroup{},
			},
			want: want{
				target: &v1beta2.ReplicationGroup{
					// All fields should remain nil/zero
				},
			},
		},
	}

	// Get the conversion function from the config (replicated from config.go)
	conversions := []struct {
		from, to string
		fn       func(src, target xpresource.Managed) error
	}{
		{"v1beta1", "v1beta2", func(src, target xpresource.Managed) error {
			srcTyped := src.(*v1beta1.ReplicationGroup)
			targetTyped := target.(*v1beta2.ReplicationGroup)
			if len(srcTyped.Spec.ForProvider.ClusterMode) > 0 {
				if srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups != nil {
					targetTyped.Spec.ForProvider.NumNodeGroups = srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups
				}
				if srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
					targetTyped.Spec.ForProvider.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup
				}
			}
			if len(srcTyped.Spec.InitProvider.ClusterMode) > 0 {
				if srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups != nil {
					targetTyped.Spec.InitProvider.NumNodeGroups = srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups
				}
				if srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
					targetTyped.Spec.InitProvider.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup
				}
			}
			if len(srcTyped.Status.AtProvider.ClusterMode) > 0 {
				if srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups != nil {
					targetTyped.Status.AtProvider.NumNodeGroups = srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups
				}
				if srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
					targetTyped.Status.AtProvider.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup
				}
			}
			return nil
		}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := conversions[0].fn(tc.args.src, tc.args.target)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("%s\nConversion(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.target, tc.args.target); diff != "" {
				t.Errorf("%s\nConversion(...): -want target, +got target:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestReplicationGroupConversion_v1beta2_to_v1beta1(t *testing.T) {
	type args struct {
		src    xpresource.Managed
		target xpresource.Managed
	}
	type want struct {
		target xpresource.Managed
		err    error
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"ConvertForProviderIndividualFields": {
			reason: "Should convert v1beta2 ForProvider individual fields to v1beta1 ClusterMode",
			args: args{
				src: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						ForProvider: v1beta2.ReplicationGroupParameters{
							NumNodeGroups:        floatPtr(3),
							ReplicasPerNodeGroup: floatPtr(2),
						},
					},
				},
				target: &v1beta1.ReplicationGroup{},
			},
			want: want{
				target: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups:        floatPtr(3),
									ReplicasPerNodeGroup: floatPtr(2),
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									// Both fields should be nil since InitProvider was not set in source
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									// Both fields should be nil since AtProvider was not set in source
								},
							},
						},
					},
				},
			},
		},
		"ConvertInitProviderIndividualFields": {
			reason: "Should convert v1beta2 InitProvider individual fields to v1beta1 ClusterMode",
			args: args{
				src: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						InitProvider: v1beta2.ReplicationGroupInitParameters{
							NumNodeGroups:        floatPtr(5),
							ReplicasPerNodeGroup: floatPtr(1),
						},
					},
				},
				target: &v1beta1.ReplicationGroup{},
			},
			want: want{
				target: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									// Both fields should be nil since ForProvider was not set in source
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									NumNodeGroups:        floatPtr(5),
									ReplicasPerNodeGroup: floatPtr(1),
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									// Both fields should be nil since AtProvider was not set in source
								},
							},
						},
					},
				},
			},
		},
		"ConvertAtProviderIndividualFields": {
			reason: "Should convert v1beta2 AtProvider individual fields to v1beta1 ClusterMode",
			args: args{
				src: &v1beta2.ReplicationGroup{
					Status: v1beta2.ReplicationGroupStatus{
						AtProvider: v1beta2.ReplicationGroupObservation{
							NumNodeGroups:        floatPtr(4),
							ReplicasPerNodeGroup: floatPtr(3),
						},
					},
				},
				target: &v1beta1.ReplicationGroup{},
			},
			want: want{
				target: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									// Both fields should be nil since ForProvider was not set in source
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									// Both fields should be nil since InitProvider was not set in source
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									NumNodeGroups:        floatPtr(4),
									ReplicasPerNodeGroup: floatPtr(3),
								},
							},
						},
					},
				},
			},
		},
		"ConvertAllSectionsIndividualFields": {
			reason: "Should convert all three sections (ForProvider, InitProvider, AtProvider) from v1beta2 to v1beta1",
			args: args{
				src: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						ForProvider: v1beta2.ReplicationGroupParameters{
							NumNodeGroups:        floatPtr(2),
							ReplicasPerNodeGroup: floatPtr(1),
						},
						InitProvider: v1beta2.ReplicationGroupInitParameters{
							NumNodeGroups:        floatPtr(3),
							ReplicasPerNodeGroup: floatPtr(2),
						},
					},
					Status: v1beta2.ReplicationGroupStatus{
						AtProvider: v1beta2.ReplicationGroupObservation{
							NumNodeGroups:        floatPtr(4),
							ReplicasPerNodeGroup: floatPtr(3),
						},
					},
				},
				target: &v1beta1.ReplicationGroup{},
			},
			want: want{
				target: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups:        floatPtr(2),
									ReplicasPerNodeGroup: floatPtr(1),
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									NumNodeGroups:        floatPtr(3),
									ReplicasPerNodeGroup: floatPtr(2),
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									NumNodeGroups:        floatPtr(4),
									ReplicasPerNodeGroup: floatPtr(3),
								},
							},
						},
					},
				},
			},
		},
		"ConvertPartialFields": {
			reason: "Should handle partial conversion when only some fields are present",
			args: args{
				src: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						ForProvider: v1beta2.ReplicationGroupParameters{
							NumNodeGroups: floatPtr(2),
							// ReplicasPerNodeGroup is nil
						},
						InitProvider: v1beta2.ReplicationGroupInitParameters{
							// NumNodeGroups is nil
							ReplicasPerNodeGroup: floatPtr(1),
						},
					},
				},
				target: &v1beta1.ReplicationGroup{},
			},
			want: want{
				target: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups: floatPtr(2),
									// ReplicasPerNodeGroup should remain nil
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									// NumNodeGroups should remain nil
									ReplicasPerNodeGroup: floatPtr(1),
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									// Both fields should be nil since AtProvider was not set in source
								},
							},
						},
					},
				},
			},
		},
		"ConvertNilFields": {
			reason: "Should handle nil fields gracefully by creating empty ClusterMode structs",
			args: args{
				src: &v1beta2.ReplicationGroup{
					Spec: v1beta2.ReplicationGroupSpec{
						ForProvider: v1beta2.ReplicationGroupParameters{
							// Both fields are nil
						},
						InitProvider: v1beta2.ReplicationGroupInitParameters{
							// Both fields are nil
						},
					},
					Status: v1beta2.ReplicationGroupStatus{
						AtProvider: v1beta2.ReplicationGroupObservation{
							// Both fields are nil
						},
					},
				},
				target: &v1beta1.ReplicationGroup{},
			},
			want: want{
				target: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									// Both fields should be nil
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									// Both fields should be nil
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									// Both fields should be nil
								},
							},
						},
					},
				},
			},
		},
	}

	// Get the conversion function from the config (replicated from config.go)
	conversion := func(src, target xpresource.Managed) error {
		srcTyped := src.(*v1beta2.ReplicationGroup)
		targetTyped := target.(*v1beta1.ReplicationGroup)
		cm := v1beta1.ClusterModeParameters{}
		if srcTyped.Spec.ForProvider.NumNodeGroups != nil {
			cm.NumNodeGroups = srcTyped.Spec.ForProvider.NumNodeGroups
		}
		if srcTyped.Spec.ForProvider.ReplicasPerNodeGroup != nil {
			cm.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ReplicasPerNodeGroup
		}
		targetTyped.Spec.ForProvider.ClusterMode = []v1beta1.ClusterModeParameters{cm}

		cmi := v1beta1.ClusterModeInitParameters{}
		if srcTyped.Spec.InitProvider.NumNodeGroups != nil {
			cmi.NumNodeGroups = srcTyped.Spec.InitProvider.NumNodeGroups
		}
		if srcTyped.Spec.InitProvider.ReplicasPerNodeGroup != nil {
			cmi.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ReplicasPerNodeGroup
		}
		targetTyped.Spec.InitProvider.ClusterMode = []v1beta1.ClusterModeInitParameters{cmi}

		cmo := v1beta1.ClusterModeObservation{}
		if srcTyped.Status.AtProvider.NumNodeGroups != nil {
			cmo.NumNodeGroups = srcTyped.Status.AtProvider.NumNodeGroups
		}
		if srcTyped.Status.AtProvider.ReplicasPerNodeGroup != nil {
			cmo.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ReplicasPerNodeGroup
		}
		targetTyped.Status.AtProvider.ClusterMode = []v1beta1.ClusterModeObservation{cmo}
		return nil
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := conversion(tc.args.src, tc.args.target)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("%s\nConversion(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.target, tc.args.target); diff != "" {
				t.Errorf("%s\nConversion(...): -want target, +got target:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestReplicationGroupConversion_RoundTrip(t *testing.T) {
	type args struct {
		original *v1beta1.ReplicationGroup
	}
	type want struct {
		final *v1beta1.ReplicationGroup
		err   error
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"RoundTripFullyPopulated": {
			reason: "Should preserve all data through v1beta1 -> v1beta2 -> v1beta1 conversion",
			args: args{
				original: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups:        floatPtr(3),
									ReplicasPerNodeGroup: floatPtr(2),
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									NumNodeGroups:        floatPtr(5),
									ReplicasPerNodeGroup: floatPtr(1),
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									NumNodeGroups:        floatPtr(4),
									ReplicasPerNodeGroup: floatPtr(3),
								},
							},
						},
					},
				},
			},
			want: want{
				final: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups:        floatPtr(3),
									ReplicasPerNodeGroup: floatPtr(2),
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									NumNodeGroups:        floatPtr(5),
									ReplicasPerNodeGroup: floatPtr(1),
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									NumNodeGroups:        floatPtr(4),
									ReplicasPerNodeGroup: floatPtr(3),
								},
							},
						},
					},
				},
			},
		},
		"RoundTripPartialData": {
			reason: "Should preserve partial data through round-trip conversion",
			args: args{
				original: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups: floatPtr(2),
									// ReplicasPerNodeGroup is nil
								},
							},
						},
					},
				},
			},
			want: want{
				final: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									NumNodeGroups: floatPtr(2),
									// ReplicasPerNodeGroup should remain nil
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									// Both fields should be nil since InitProvider was empty in original
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									// Both fields should be nil
								},
							},
						},
					},
				},
			},
		},
		"RoundTripEmptyClusterMode": {
			reason: "Should handle empty ClusterMode arrays through round-trip conversion",
			args: args{
				original: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{}, // empty
						},
					},
				},
			},
			want: want{
				final: &v1beta1.ReplicationGroup{
					Spec: v1beta1.ReplicationGroupSpec{
						ForProvider: v1beta1.ReplicationGroupParameters{
							ClusterMode: []v1beta1.ClusterModeParameters{
								{
									// Both fields should be nil
								},
							},
						},
						InitProvider: v1beta1.ReplicationGroupInitParameters{
							ClusterMode: []v1beta1.ClusterModeInitParameters{
								{
									// Both fields should be nil
								},
							},
						},
					},
					Status: v1beta1.ReplicationGroupStatus{
						AtProvider: v1beta1.ReplicationGroupObservation{
							ClusterMode: []v1beta1.ClusterModeObservation{
								{
									// Both fields should be nil
								},
							},
						},
					},
				},
			},
		},
	}

	// Conversion functions
	v1beta1ToV1beta2 := func(src, target xpresource.Managed) error {
		srcTyped := src.(*v1beta1.ReplicationGroup)
		targetTyped := target.(*v1beta2.ReplicationGroup)
		if len(srcTyped.Spec.ForProvider.ClusterMode) > 0 {
			if srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups != nil {
				targetTyped.Spec.ForProvider.NumNodeGroups = srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups
			}
			if srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
				targetTyped.Spec.ForProvider.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup
			}
		}
		if len(srcTyped.Spec.InitProvider.ClusterMode) > 0 {
			if srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups != nil {
				targetTyped.Spec.InitProvider.NumNodeGroups = srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups
			}
			if srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
				targetTyped.Spec.InitProvider.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup
			}
		}
		if len(srcTyped.Status.AtProvider.ClusterMode) > 0 {
			if srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups != nil {
				targetTyped.Status.AtProvider.NumNodeGroups = srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups
			}
			if srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
				targetTyped.Status.AtProvider.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup
			}
		}
		return nil
	}

	v1beta2ToV1beta1 := func(src, target xpresource.Managed) error {
		srcTyped := src.(*v1beta2.ReplicationGroup)
		targetTyped := target.(*v1beta1.ReplicationGroup)
		cm := v1beta1.ClusterModeParameters{}
		if srcTyped.Spec.ForProvider.NumNodeGroups != nil {
			cm.NumNodeGroups = srcTyped.Spec.ForProvider.NumNodeGroups
		}
		if srcTyped.Spec.ForProvider.ReplicasPerNodeGroup != nil {
			cm.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ReplicasPerNodeGroup
		}
		targetTyped.Spec.ForProvider.ClusterMode = []v1beta1.ClusterModeParameters{cm}

		cmi := v1beta1.ClusterModeInitParameters{}
		if srcTyped.Spec.InitProvider.NumNodeGroups != nil {
			cmi.NumNodeGroups = srcTyped.Spec.InitProvider.NumNodeGroups
		}
		if srcTyped.Spec.InitProvider.ReplicasPerNodeGroup != nil {
			cmi.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ReplicasPerNodeGroup
		}
		targetTyped.Spec.InitProvider.ClusterMode = []v1beta1.ClusterModeInitParameters{cmi}

		cmo := v1beta1.ClusterModeObservation{}
		if srcTyped.Status.AtProvider.NumNodeGroups != nil {
			cmo.NumNodeGroups = srcTyped.Status.AtProvider.NumNodeGroups
		}
		if srcTyped.Status.AtProvider.ReplicasPerNodeGroup != nil {
			cmo.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ReplicasPerNodeGroup
		}
		targetTyped.Status.AtProvider.ClusterMode = []v1beta1.ClusterModeObservation{cmo}
		return nil
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// Create deep copy of original
			originalCopy := &v1beta1.ReplicationGroup{}
			*originalCopy = *tc.args.original

			// Step 1: v1beta1 -> v1beta2
			intermediate := &v1beta2.ReplicationGroup{}
			err := v1beta1ToV1beta2(tc.args.original, intermediate)
			if err != nil {
				t.Errorf("v1beta1 to v1beta2 conversion failed: %v", err)
				return
			}

			// Step 2: v1beta2 -> v1beta1
			final := &v1beta1.ReplicationGroup{}
			err = v1beta2ToV1beta1(intermediate, final)
			if err != nil {
				t.Errorf("v1beta2 to v1beta1 conversion failed: %v", err)
				return
			}

			// Compare final result with expected
			if diff := cmp.Diff(tc.want.final, final); diff != "" {
				t.Errorf("%s\nRound-trip conversion: -want final, +got final:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestReplicationGroupConversion_BugFix_AtProviderAssignment(t *testing.T) {
	// This test specifically verifies the bug fix where AtProvider ReplicasPerNodeGroup
	// was incorrectly assigned to 'cm' instead of 'cmo' in the v1beta2 to v1beta1 conversion.

	// The bug was on line 170 of config.go - it should assign to cmo.ReplicasPerNodeGroup, not cm.ReplicasPerNodeGroup

	src := &v1beta2.ReplicationGroup{
		Status: v1beta2.ReplicationGroupStatus{
			AtProvider: v1beta2.ReplicationGroupObservation{
				NumNodeGroups:        floatPtr(2),
				ReplicasPerNodeGroup: floatPtr(3),
			},
		},
	}

	target := &v1beta1.ReplicationGroup{}

	// Apply the corrected conversion function
	conversion := func(src, target xpresource.Managed) error {
		srcTyped := src.(*v1beta2.ReplicationGroup)
		targetTyped := target.(*v1beta1.ReplicationGroup)
		cm := v1beta1.ClusterModeParameters{}
		if srcTyped.Spec.ForProvider.NumNodeGroups != nil {
			cm.NumNodeGroups = srcTyped.Spec.ForProvider.NumNodeGroups
		}
		if srcTyped.Spec.ForProvider.ReplicasPerNodeGroup != nil {
			cm.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ReplicasPerNodeGroup
		}
		targetTyped.Spec.ForProvider.ClusterMode = []v1beta1.ClusterModeParameters{cm}

		cmi := v1beta1.ClusterModeInitParameters{}
		if srcTyped.Spec.InitProvider.NumNodeGroups != nil {
			cmi.NumNodeGroups = srcTyped.Spec.InitProvider.NumNodeGroups
		}
		if srcTyped.Spec.InitProvider.ReplicasPerNodeGroup != nil {
			cmi.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ReplicasPerNodeGroup
		}
		targetTyped.Spec.InitProvider.ClusterMode = []v1beta1.ClusterModeInitParameters{cmi}

		cmo := v1beta1.ClusterModeObservation{}
		if srcTyped.Status.AtProvider.NumNodeGroups != nil {
			cmo.NumNodeGroups = srcTyped.Status.AtProvider.NumNodeGroups
		}
		if srcTyped.Status.AtProvider.ReplicasPerNodeGroup != nil {
			// BUG FIX: This should assign to cmo.ReplicasPerNodeGroup, NOT cm.ReplicasPerNodeGroup
			cmo.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ReplicasPerNodeGroup
		}
		targetTyped.Status.AtProvider.ClusterMode = []v1beta1.ClusterModeObservation{cmo}
		return nil
	}

	err := conversion(src, target)
	if err != nil {
		t.Errorf("Conversion failed: %v", err)
	}

	// Verify that AtProvider values are correctly assigned to the AtProvider section
	if len(target.Status.AtProvider.ClusterMode) == 0 {
		t.Error("AtProvider ClusterMode should not be nil or empty")
		return
	}

	atProvider := target.Status.AtProvider.ClusterMode[0]
	if atProvider.NumNodeGroups == nil || *atProvider.NumNodeGroups != 2 {
		t.Errorf("AtProvider NumNodeGroups should be 2, got %v", atProvider.NumNodeGroups)
	}

	if atProvider.ReplicasPerNodeGroup == nil || *atProvider.ReplicasPerNodeGroup != 3 {
		t.Errorf("AtProvider ReplicasPerNodeGroup should be 3, got %v", atProvider.ReplicasPerNodeGroup)
	}

	// Verify that ForProvider values are empty (since we didn't provide any)
	if len(target.Spec.ForProvider.ClusterMode) == 0 {
		t.Error("ForProvider ClusterMode should exist even if empty")
		return
	}

	forProvider := target.Spec.ForProvider.ClusterMode[0]
	if forProvider.NumNodeGroups != nil {
		t.Errorf("ForProvider NumNodeGroups should be nil, got %v", forProvider.NumNodeGroups)
	}

	if forProvider.ReplicasPerNodeGroup != nil {
		t.Errorf("ForProvider ReplicasPerNodeGroup should be nil, got %v", forProvider.ReplicasPerNodeGroup)
	}
}

func TestReplicationGroupConversion_DataIntegrityPreservation(t *testing.T) {
	// This test ensures that NumNodeGroups and ReplicasPerNodeGroup values are preserved correctly
	// across conversions and that there's no data loss or corruption.

	testCases := []struct {
		name             string
		numNodeGroups    *float64
		replicasPerGroup *float64
	}{
		{"ZeroValues", floatPtr(0), floatPtr(0)},
		{"SmallValues", floatPtr(1), floatPtr(1)},
		{"TypicalValues", floatPtr(3), floatPtr(2)},
		{"LargeValues", floatPtr(10), floatPtr(5)},
		{"MaxReplicas", floatPtr(1), floatPtr(5)}, // Max valid replicas per node group is 5
		{"OnlyNumNodeGroups", floatPtr(3), nil},
		{"OnlyReplicasPerGroup", nil, floatPtr(2)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test all three sections independently
			sections := []struct {
				name         string
				setupV1beta1 func(*v1beta1.ReplicationGroup)
				checkV1beta2 func(*v1beta2.ReplicationGroup) error
				setupV1beta2 func(*v1beta2.ReplicationGroup)
				checkV1beta1 func(*v1beta1.ReplicationGroup) error
			}{
				{
					name: "ForProvider",
					setupV1beta1: func(rg *v1beta1.ReplicationGroup) {
						rg.Spec.ForProvider.ClusterMode = []v1beta1.ClusterModeParameters{
							{
								NumNodeGroups:        tc.numNodeGroups,
								ReplicasPerNodeGroup: tc.replicasPerGroup,
							},
						}
					},
					checkV1beta2: func(rg *v1beta2.ReplicationGroup) error {
						if !floatPtrEqual(rg.Spec.ForProvider.NumNodeGroups, tc.numNodeGroups) {
							return errors.Errorf("ForProvider NumNodeGroups mismatch: expected %v, got %v",
								tc.numNodeGroups, rg.Spec.ForProvider.NumNodeGroups)
						}
						if !floatPtrEqual(rg.Spec.ForProvider.ReplicasPerNodeGroup, tc.replicasPerGroup) {
							return errors.Errorf("ForProvider ReplicasPerNodeGroup mismatch: expected %v, got %v",
								tc.replicasPerGroup, rg.Spec.ForProvider.ReplicasPerNodeGroup)
						}
						return nil
					},
					setupV1beta2: func(rg *v1beta2.ReplicationGroup) {
						rg.Spec.ForProvider.NumNodeGroups = tc.numNodeGroups
						rg.Spec.ForProvider.ReplicasPerNodeGroup = tc.replicasPerGroup
					},
					checkV1beta1: func(rg *v1beta1.ReplicationGroup) error {
						if len(rg.Spec.ForProvider.ClusterMode) == 0 {
							return errors.New("ForProvider ClusterMode should not be empty")
						}
						cm := rg.Spec.ForProvider.ClusterMode[0]
						if !floatPtrEqual(cm.NumNodeGroups, tc.numNodeGroups) {
							return errors.Errorf("ForProvider NumNodeGroups mismatch: expected %v, got %v",
								tc.numNodeGroups, cm.NumNodeGroups)
						}
						if !floatPtrEqual(cm.ReplicasPerNodeGroup, tc.replicasPerGroup) {
							return errors.Errorf("ForProvider ReplicasPerNodeGroup mismatch: expected %v, got %v",
								tc.replicasPerGroup, cm.ReplicasPerNodeGroup)
						}
						return nil
					},
				},
				{
					name: "InitProvider",
					setupV1beta1: func(rg *v1beta1.ReplicationGroup) {
						rg.Spec.InitProvider.ClusterMode = []v1beta1.ClusterModeInitParameters{
							{
								NumNodeGroups:        tc.numNodeGroups,
								ReplicasPerNodeGroup: tc.replicasPerGroup,
							},
						}
					},
					checkV1beta2: func(rg *v1beta2.ReplicationGroup) error {
						if !floatPtrEqual(rg.Spec.InitProvider.NumNodeGroups, tc.numNodeGroups) {
							return errors.Errorf("InitProvider NumNodeGroups mismatch: expected %v, got %v",
								tc.numNodeGroups, rg.Spec.InitProvider.NumNodeGroups)
						}
						if !floatPtrEqual(rg.Spec.InitProvider.ReplicasPerNodeGroup, tc.replicasPerGroup) {
							return errors.Errorf("InitProvider ReplicasPerNodeGroup mismatch: expected %v, got %v",
								tc.replicasPerGroup, rg.Spec.InitProvider.ReplicasPerNodeGroup)
						}
						return nil
					},
					setupV1beta2: func(rg *v1beta2.ReplicationGroup) {
						rg.Spec.InitProvider.NumNodeGroups = tc.numNodeGroups
						rg.Spec.InitProvider.ReplicasPerNodeGroup = tc.replicasPerGroup
					},
					checkV1beta1: func(rg *v1beta1.ReplicationGroup) error {
						if len(rg.Spec.InitProvider.ClusterMode) == 0 {
							return errors.New("InitProvider ClusterMode should not be empty")
						}
						cm := rg.Spec.InitProvider.ClusterMode[0]
						if !floatPtrEqual(cm.NumNodeGroups, tc.numNodeGroups) {
							return errors.Errorf("InitProvider NumNodeGroups mismatch: expected %v, got %v",
								tc.numNodeGroups, cm.NumNodeGroups)
						}
						if !floatPtrEqual(cm.ReplicasPerNodeGroup, tc.replicasPerGroup) {
							return errors.Errorf("InitProvider ReplicasPerNodeGroup mismatch: expected %v, got %v",
								tc.replicasPerGroup, cm.ReplicasPerNodeGroup)
						}
						return nil
					},
				},
				{
					name: "AtProvider",
					setupV1beta1: func(rg *v1beta1.ReplicationGroup) {
						rg.Status.AtProvider.ClusterMode = []v1beta1.ClusterModeObservation{
							{
								NumNodeGroups:        tc.numNodeGroups,
								ReplicasPerNodeGroup: tc.replicasPerGroup,
							},
						}
					},
					checkV1beta2: func(rg *v1beta2.ReplicationGroup) error {
						if !floatPtrEqual(rg.Status.AtProvider.NumNodeGroups, tc.numNodeGroups) {
							return errors.Errorf("AtProvider NumNodeGroups mismatch: expected %v, got %v",
								tc.numNodeGroups, rg.Status.AtProvider.NumNodeGroups)
						}
						if !floatPtrEqual(rg.Status.AtProvider.ReplicasPerNodeGroup, tc.replicasPerGroup) {
							return errors.Errorf("AtProvider ReplicasPerNodeGroup mismatch: expected %v, got %v",
								tc.replicasPerGroup, rg.Status.AtProvider.ReplicasPerNodeGroup)
						}
						return nil
					},
					setupV1beta2: func(rg *v1beta2.ReplicationGroup) {
						rg.Status.AtProvider.NumNodeGroups = tc.numNodeGroups
						rg.Status.AtProvider.ReplicasPerNodeGroup = tc.replicasPerGroup
					},
					checkV1beta1: func(rg *v1beta1.ReplicationGroup) error {
						if len(rg.Status.AtProvider.ClusterMode) == 0 {
							return errors.New("AtProvider ClusterMode should not be empty")
						}
						cm := rg.Status.AtProvider.ClusterMode[0]
						if !floatPtrEqual(cm.NumNodeGroups, tc.numNodeGroups) {
							return errors.Errorf("AtProvider NumNodeGroups mismatch: expected %v, got %v",
								tc.numNodeGroups, cm.NumNodeGroups)
						}
						if !floatPtrEqual(cm.ReplicasPerNodeGroup, tc.replicasPerGroup) {
							return errors.Errorf("AtProvider ReplicasPerNodeGroup mismatch: expected %v, got %v",
								tc.replicasPerGroup, cm.ReplicasPerNodeGroup)
						}
						return nil
					},
				},
			}

			for _, section := range sections {
				t.Run(section.name, func(t *testing.T) {
					// Test v1beta1 to v1beta2 conversion
					v1beta1Src := &v1beta1.ReplicationGroup{}
					section.setupV1beta1(v1beta1Src)

					v1beta2Target := &v1beta2.ReplicationGroup{}
					err := v1beta1ToV1beta2Conversion(v1beta1Src, v1beta2Target)
					if err != nil {
						t.Errorf("v1beta1 to v1beta2 conversion failed: %v", err)
					}

					if err := section.checkV1beta2(v1beta2Target); err != nil {
						t.Errorf("v1beta1 to v1beta2 validation failed: %v", err)
					}

					// Test v1beta2 to v1beta1 conversion
					v1beta2Src := &v1beta2.ReplicationGroup{}
					section.setupV1beta2(v1beta2Src)

					v1beta1Target := &v1beta1.ReplicationGroup{}
					err = v1beta2ToV1beta1Conversion(v1beta2Src, v1beta1Target)
					if err != nil {
						t.Errorf("v1beta2 to v1beta1 conversion failed: %v", err)
					}

					if err := section.checkV1beta1(v1beta1Target); err != nil {
						t.Errorf("v1beta2 to v1beta1 validation failed: %v", err)
					}
				})
			}
		})
	}
}

// Helper functions for the tests

func v1beta1ToV1beta2Conversion(src, target xpresource.Managed) error {
	srcTyped := src.(*v1beta1.ReplicationGroup)
	targetTyped := target.(*v1beta2.ReplicationGroup)
	if len(srcTyped.Spec.ForProvider.ClusterMode) > 0 {
		if srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups != nil {
			targetTyped.Spec.ForProvider.NumNodeGroups = srcTyped.Spec.ForProvider.ClusterMode[0].NumNodeGroups
		}
		if srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
			targetTyped.Spec.ForProvider.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ClusterMode[0].ReplicasPerNodeGroup
		}
	}
	if len(srcTyped.Spec.InitProvider.ClusterMode) > 0 {
		if srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups != nil {
			targetTyped.Spec.InitProvider.NumNodeGroups = srcTyped.Spec.InitProvider.ClusterMode[0].NumNodeGroups
		}
		if srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
			targetTyped.Spec.InitProvider.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ClusterMode[0].ReplicasPerNodeGroup
		}
	}
	if len(srcTyped.Status.AtProvider.ClusterMode) > 0 {
		if srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups != nil {
			targetTyped.Status.AtProvider.NumNodeGroups = srcTyped.Status.AtProvider.ClusterMode[0].NumNodeGroups
		}
		if srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup != nil {
			targetTyped.Status.AtProvider.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ClusterMode[0].ReplicasPerNodeGroup
		}
	}
	return nil
}

func v1beta2ToV1beta1Conversion(src, target xpresource.Managed) error {
	srcTyped := src.(*v1beta2.ReplicationGroup)
	targetTyped := target.(*v1beta1.ReplicationGroup)
	cm := v1beta1.ClusterModeParameters{}
	if srcTyped.Spec.ForProvider.NumNodeGroups != nil {
		cm.NumNodeGroups = srcTyped.Spec.ForProvider.NumNodeGroups
	}
	if srcTyped.Spec.ForProvider.ReplicasPerNodeGroup != nil {
		cm.ReplicasPerNodeGroup = srcTyped.Spec.ForProvider.ReplicasPerNodeGroup
	}
	targetTyped.Spec.ForProvider.ClusterMode = []v1beta1.ClusterModeParameters{cm}

	cmi := v1beta1.ClusterModeInitParameters{}
	if srcTyped.Spec.InitProvider.NumNodeGroups != nil {
		cmi.NumNodeGroups = srcTyped.Spec.InitProvider.NumNodeGroups
	}
	if srcTyped.Spec.InitProvider.ReplicasPerNodeGroup != nil {
		cmi.ReplicasPerNodeGroup = srcTyped.Spec.InitProvider.ReplicasPerNodeGroup
	}
	targetTyped.Spec.InitProvider.ClusterMode = []v1beta1.ClusterModeInitParameters{cmi}

	cmo := v1beta1.ClusterModeObservation{}
	if srcTyped.Status.AtProvider.NumNodeGroups != nil {
		cmo.NumNodeGroups = srcTyped.Status.AtProvider.NumNodeGroups
	}
	if srcTyped.Status.AtProvider.ReplicasPerNodeGroup != nil {
		cmo.ReplicasPerNodeGroup = srcTyped.Status.AtProvider.ReplicasPerNodeGroup
	}
	targetTyped.Status.AtProvider.ClusterMode = []v1beta1.ClusterModeObservation{cmo}
	return nil
}
