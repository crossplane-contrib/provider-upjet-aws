// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package eks

import (
	"testing"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/upbound/provider-aws/apis/cluster/eks/v1beta1"
	"github.com/upbound/provider-aws/apis/cluster/eks/v1beta2"
)

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}

// stringPtrEqual compares two string pointers for equality
func stringPtrEqual(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// TestClusterConversion tests simple Cluster field conversion between v1beta1 and v1beta2
func TestClusterConversion(t *testing.T) {
	t.Run("v1beta1_to_v1beta2_Cluster_conversion", func(t *testing.T) {
		// Create v1beta1 Cluster with individual fields
		src := &v1beta1.Cluster{
			Spec: v1beta1.ClusterSpec{
				ForProvider: v1beta1.ClusterParameters{
					UpgradePolicy: []v1beta1.UpgradePolicyParameters{
						{
							SupportType: stringPtr("STANDARD"),
						},
					},
				},
				InitProvider: v1beta1.ClusterInitParameters{
					UpgradePolicy: []v1beta1.UpgradePolicyInitParameters{
						{
							SupportType: stringPtr("STANDARD"),
						},
					},
				},
			},
			Status: v1beta1.ClusterStatus{
				AtProvider: v1beta1.ClusterObservation{
					UpgradePolicy: []v1beta1.UpgradePolicyObservation{
						{
							SupportType: stringPtr("STANDARD"),
						},
					},
				},
			},
		}

		target := &v1beta2.Cluster{}

		// Apply v1beta1 to v1beta2 conversion
		converter := clusterConverterFromv1beta1Tov1beta2
		err := converter(src, target)
		if err != nil {
			t.Fatalf("Conversion failed: %v", err)
		}

		// Assert the ServerSideEncryptionConfiguration fields were converted correctly
		if !stringPtrEqual(target.Spec.ForProvider.UpgradePolicy.SupportType, stringPtr("STANDARD")) {
			t.Errorf("Expected ForProvider UpgradePolicy SupportType=STANDARD, got %v", target.Spec.ForProvider.UpgradePolicy.SupportType)
		}
		if !stringPtrEqual(target.Spec.InitProvider.UpgradePolicy.SupportType, stringPtr("STANDARD")) {
			t.Errorf("Expected InitProvider UpgradePolicy SupportType=STANDARD, got %v", target.Spec.InitProvider.UpgradePolicy.SupportType)
		}
		if !stringPtrEqual(target.Status.AtProvider.UpgradePolicy.SupportType, stringPtr("STANDARD")) {
			t.Errorf("Expected AtProvider UpgradePolicy SupportType=STANDARD, got %v", target.Status.AtProvider.UpgradePolicy.SupportType)
		}
	})

	t.Run("v1beta2_to_v1beta1_Cluster_conversion", func(t *testing.T) {
		// Create v1beta2 Cluster with individual fields
		src := &v1beta2.Cluster{
			Spec: v1beta2.ClusterSpec{
				ForProvider: v1beta2.ClusterParameters{
					UpgradePolicy: &v1beta2.UpgradePolicyParameters{
						SupportType: stringPtr("STANDARD"),
					},
				},
				InitProvider: v1beta2.ClusterInitParameters{
					UpgradePolicy: &v1beta2.UpgradePolicyInitParameters{
						SupportType: stringPtr("STANDARD"),
					},
				},
			},
			Status: v1beta2.ClusterStatus{
				AtProvider: v1beta2.ClusterObservation{
					UpgradePolicy: &v1beta2.UpgradePolicyObservation{
						SupportType: stringPtr("STANDARD"),
					},
				},
			},
		}

		target := &v1beta1.Cluster{}

		// Apply v1beta2 to v1beta1 conversion
		converter := clusterConverterFromv1beta2Tov1beta1
		err := converter(src, target)
		if err != nil {
			t.Fatalf("Conversion failed: %v", err)
		}

		// Assert the ServerSideEncryptionConfiguration fields were converted correctly
		if len(target.Spec.ForProvider.UpgradePolicy) != 1 {
			t.Errorf("Expected ForProvider UpgradePolicy length=1, got %d", len(target.Spec.ForProvider.UpgradePolicy))
		}
		if !stringPtrEqual(target.Spec.ForProvider.UpgradePolicy[0].SupportType, stringPtr("STANDARD")) {
			t.Errorf("Expected ForProvider UpgradePolicy SupportType=STANDARD, got %v", target.Spec.ForProvider.UpgradePolicy[0].SupportType)
		}
		if len(target.Spec.InitProvider.UpgradePolicy) != 1 {
			t.Errorf("Expected InitProvider UpgradePolicy length=1, got %d", len(target.Spec.InitProvider.UpgradePolicy))
		}
		if !stringPtrEqual(target.Spec.InitProvider.UpgradePolicy[0].SupportType, stringPtr("STANDARD")) {
			t.Errorf("Expected InitProvider UpgradePolicy SupportType=STANDARD, got %v", target.Spec.InitProvider.UpgradePolicy[0].SupportType)
		}
		if len(target.Status.AtProvider.UpgradePolicy) != 1 {
			t.Errorf("Expected AtProvider UpgradePolicy length=1, got %d", len(target.Status.AtProvider.UpgradePolicy))
		}
		if !stringPtrEqual(target.Status.AtProvider.UpgradePolicy[0].SupportType, stringPtr("STANDARD")) {
			t.Errorf("Expected AtProvider UpgradePolicy SupportType=STANDARD, got %v", target.Status.AtProvider.UpgradePolicy[0].SupportType)
		}
	})
}

func TestClusterConversion_v1beta1_to_v1beta2(t *testing.T) {
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
		"ConvertPartialFields": {
			reason: "Should handle partial conversion when only some fields are present",
			args: args{
				src: &v1beta1.Cluster{
					Spec: v1beta1.ClusterSpec{
						ForProvider: v1beta1.ClusterParameters{
							UpgradePolicy: []v1beta1.UpgradePolicyParameters{
								{
									SupportType: stringPtr("STANDARD"),
								},
							},
						},
						// InitProvider is nil
					},
					// Status is nil
				},
				target: &v1beta2.Cluster{},
			},
			want: want{
				target: &v1beta2.Cluster{
					Spec: v1beta2.ClusterSpec{
						ForProvider: v1beta2.ClusterParameters{
							UpgradePolicy: &v1beta2.UpgradePolicyParameters{
								SupportType: stringPtr("STANDARD"),
							},
						},
						// InitProvider should remain nil
					},
					// Status should remain nil
				},
			},
		},
		"ConvertEmptyClusterModeArrays": {
			reason: "Should handle empty ClusterMode arrays gracefully",
			args: args{
				src: &v1beta1.Cluster{
					Spec: v1beta1.ClusterSpec{
						ForProvider: v1beta1.ClusterParameters{
							UpgradePolicy: []v1beta1.UpgradePolicyParameters{},
						},
						InitProvider: v1beta1.ClusterInitParameters{
							UpgradePolicy: []v1beta1.UpgradePolicyInitParameters{},
						},
					},
					Status: v1beta1.ClusterStatus{
						AtProvider: v1beta1.ClusterObservation{
							UpgradePolicy: []v1beta1.UpgradePolicyObservation{},
						},
					},
				},
				target: &v1beta2.Cluster{},
			},
			want: want{
				target: &v1beta2.Cluster{
					// All fields should remain nil/zero
				},
			},
		},
		"ConvertNilClusterModeArrays": {
			reason: "Should handle nil ClusterMode arrays gracefully",
			args: args{
				src: &v1beta1.Cluster{
					Status: v1beta1.ClusterStatus{
						AtProvider: v1beta1.ClusterObservation{},
					},
				},
				target: &v1beta2.Cluster{},
			},
			want: want{
				target: &v1beta2.Cluster{
					// All fields should remain nil/zero
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := clusterConverterFromv1beta1Tov1beta2(tc.args.src, tc.args.target)
			if (err != nil) != (tc.want.err != nil) {
				t.Fatalf("got error: %v, want: %v", err, tc.want.err)
			}
			if err != nil && tc.want.err != nil && err.Error() != tc.want.err.Error() {

				t.Fatalf("got error: %v, want: %v", err, tc.want.err)
			}
			if diff := cmp.Diff(tc.want.target, tc.args.target, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestClusterConversion_v1beta2_to_v1beta1(t *testing.T) {
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
		"ConvertPartialFields": {
			reason: "Should handle partial conversion when only some fields are present",
			args: args{
				src: &v1beta2.Cluster{
					Spec: v1beta2.ClusterSpec{
						ForProvider: v1beta2.ClusterParameters{
							UpgradePolicy: &v1beta2.UpgradePolicyParameters{
								SupportType: stringPtr("STANDARD"),
							},
						},
						// InitProvider is nil
					},
					// Status is nil
				},
				target: &v1beta1.Cluster{},
			},
			want: want{
				target: &v1beta1.Cluster{
					Spec: v1beta1.ClusterSpec{
						ForProvider: v1beta1.ClusterParameters{
							UpgradePolicy: []v1beta1.UpgradePolicyParameters{
								{
									SupportType: stringPtr("STANDARD"),
								},
							},
						},
						// InitProvider should remain nil
					},
					// Status should remain nil
				},
			},
		},
		"ConvertNilClusterMode": {
			reason: "Should handle nil ClusterMode gracefully",
			args: args{
				src: &v1beta2.Cluster{
					Status: v1beta2.ClusterStatus{
						AtProvider: v1beta2.ClusterObservation{},
					},
				},
				target: &v1beta1.Cluster{},
			},
			want: want{
				target: &v1beta1.Cluster{
					Status: v1beta1.ClusterStatus{
						AtProvider: v1beta1.ClusterObservation{
							// All fields should remain nil/zero
						},
					},
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := clusterConverterFromv1beta2Tov1beta1(tc.args.src, tc.args.target)
			if (err != nil) != (tc.want.err != nil) {
				t.Fatalf("got error: %v, want: %v", err, tc.want.err)
			}
			if err != nil && tc.want.err != nil && err.Error() != tc.want.err.Error() {

				t.Fatalf("got error: %v, want: %v", err, tc.want.err)
			}
			if diff := cmp.Diff(tc.want.target, tc.args.target, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}
