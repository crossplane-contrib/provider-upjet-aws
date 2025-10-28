// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package s3

import (
	"testing"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/upbound/provider-aws/apis/cluster/s3/v1beta1"
	"github.com/upbound/provider-aws/apis/cluster/s3/v1beta2"
)

// Helper function to create bool pointer
func boolPtr(b bool) *bool {
	return &b
}

// boolPtrEqual compares two bool pointers for equality
func boolPtrEqual(a, b *bool) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

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

// TestBucketConversion tests simple Bucket field conversion between v1beta1 and v1beta2
func TestBucketConversion(t *testing.T) {
	t.Run("v1beta1_to_v1beta2_Bucket_conversion", func(t *testing.T) {
		// Create v1beta1 Bucket with individual fields
		src := &v1beta1.Bucket{
			Status: v1beta1.BucketStatus{
				AtProvider: v1beta1.BucketObservation{
					ServerSideEncryptionConfiguration: []v1beta1.ServerSideEncryptionConfigurationObservation{
						{
							Rule: []v1beta1.ServerSideEncryptionConfigurationRuleObservation{
								{
									ApplyServerSideEncryptionByDefault: []v1beta1.ApplyServerSideEncryptionByDefaultObservation{
										{
											KMSMasterKeyID: stringPtr("example-key-id"),
											SseAlgorithm:   stringPtr("AES256"),
										},
									},
									BucketKeyEnabled: boolPtr(false),
								},
							},
						},
					},
				},
			},
		}

		target := &v1beta2.Bucket{}

		// Apply v1beta1 to v1beta2 conversion
		converter := bucketConverterFromv1beta1Tov1beta2
		err := converter(src, target)
		if err != nil {
			t.Fatalf("Conversion failed: %v", err)
		}

		// Assert the ServerSideEncryptionConfiguration fields were converted correctly
		if !stringPtrEqual(target.Status.AtProvider.ServerSideEncryptionConfiguration.Rule.ApplyServerSideEncryptionByDefault.KMSMasterKeyID, stringPtr("example-key-id")) {
			t.Errorf("Expected AtProvider ServerSideEncryptionConfiguration KMSMasterKeyID=example-key-id, got %v",
				target.Status.AtProvider.ServerSideEncryptionConfiguration.Rule.ApplyServerSideEncryptionByDefault.KMSMasterKeyID)
		}
		if !stringPtrEqual(target.Status.AtProvider.ServerSideEncryptionConfiguration.Rule.ApplyServerSideEncryptionByDefault.SseAlgorithm, stringPtr("AES256")) {
			t.Errorf("Expected AtProvider ServerSideEncryptionConfiguration SseAlgorithm=AES256, got %v",
				target.Status.AtProvider.ServerSideEncryptionConfiguration.Rule.ApplyServerSideEncryptionByDefault.SseAlgorithm)
		}
		if !boolPtrEqual(target.Status.AtProvider.ServerSideEncryptionConfiguration.Rule.BucketKeyEnabled, boolPtr(false)) {
			t.Errorf("Expected AtProvider ServerSideEncryptionConfiguration BucketKeyEnabled=false, got %v",
				target.Status.AtProvider.ServerSideEncryptionConfiguration.Rule.BucketKeyEnabled)
		}
	})

	t.Run("v1beta2_to_v1beta1_Bucket_conversion", func(t *testing.T) {
		// Create v1beta2 Bucket with individual fields
		src := &v1beta2.Bucket{
			Status: v1beta2.BucketStatus{
				AtProvider: v1beta2.BucketObservation{
					ServerSideEncryptionConfiguration: &v1beta2.ServerSideEncryptionConfigurationObservation{
						Rule: &v1beta2.ServerSideEncryptionConfigurationRuleObservation{
							ApplyServerSideEncryptionByDefault: &v1beta2.ApplyServerSideEncryptionByDefaultObservation{
								KMSMasterKeyID: stringPtr("example-key-id"),
								SseAlgorithm:   stringPtr("AES256"),
							},
							BucketKeyEnabled: boolPtr(false),
						},
					},
				},
			},
		}

		target := &v1beta1.Bucket{}

		// Apply v1beta2 to v1beta1 conversion
		converter := bucketConverterFromv1beta2Tov1beta1
		err := converter(src, target)
		if err != nil {
			t.Fatalf("Conversion failed: %v", err)
		}

		// Assert the ServerSideEncryptionConfiguration fields were converted correctly
		if len(target.Status.AtProvider.ServerSideEncryptionConfiguration) != 1 {
			t.Fatalf("Expected ServerSideEncryptionConfiguration array length 1, got %d", len(target.Status.AtProvider.ServerSideEncryptionConfiguration))
		}
		if len(target.Status.AtProvider.ServerSideEncryptionConfiguration[0].Rule) != 1 {
			t.Fatalf("Expected Rule array length 1, got %d", len(target.Status.AtProvider.ServerSideEncryptionConfiguration[0].Rule))
		}
		rule := target.Status.AtProvider.ServerSideEncryptionConfiguration[0].Rule[0]
		if len(rule.ApplyServerSideEncryptionByDefault) != 1 {
			t.Fatalf("Expected ApplyServerSideEncryptionByDefault array length 1, got %d", len(rule.ApplyServerSideEncryptionByDefault))
		}
		asb := rule.ApplyServerSideEncryptionByDefault[0]
		if !stringPtrEqual(asb.KMSMasterKeyID, stringPtr("example-key-id")) {
			t.Errorf("Expected ApplyServerSideEncryptionByDefault KMSMasterKeyID=example-key-id, got %v", asb.KMSMasterKeyID)
		}
		if !stringPtrEqual(asb.SseAlgorithm, stringPtr("AES256")) {
			t.Errorf("Expected ApplyServerSideEncryptionByDefault SseAlgorithm=AES256, got %v", asb.SseAlgorithm)
		}
		if !boolPtrEqual(rule.BucketKeyEnabled, boolPtr(false)) {
			t.Errorf("Expected Rule BucketKeyEnabled=false, got %v", rule.BucketKeyEnabled)
		}
	})
}

func TestBucketConversion_v1beta1_to_v1beta2(t *testing.T) {
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
				src: &v1beta1.Bucket{
					Status: v1beta1.BucketStatus{
						AtProvider: v1beta1.BucketObservation{
							ServerSideEncryptionConfiguration: []v1beta1.ServerSideEncryptionConfigurationObservation{
								{
									Rule: []v1beta1.ServerSideEncryptionConfigurationRuleObservation{
										{
											ApplyServerSideEncryptionByDefault: []v1beta1.ApplyServerSideEncryptionByDefaultObservation{
												{
													KMSMasterKeyID: stringPtr("example-key-id"),
													// SseAlgorithm should remain nil
												},
											},
											// BucketKeyEnabled should remain nil
										},
									},
								},
							},
						},
					},
				},
				target: &v1beta2.Bucket{},
			},
			want: want{
				target: &v1beta2.Bucket{
					Status: v1beta2.BucketStatus{
						AtProvider: v1beta2.BucketObservation{
							ServerSideEncryptionConfiguration: &v1beta2.ServerSideEncryptionConfigurationObservation{
								Rule: &v1beta2.ServerSideEncryptionConfigurationRuleObservation{
									ApplyServerSideEncryptionByDefault: &v1beta2.ApplyServerSideEncryptionByDefaultObservation{
										KMSMasterKeyID: stringPtr("example-key-id"),
										// SseAlgorithm should remain nil
									},
									// BucketKeyEnabled should remain nil
								},
							},
						},
					},
				},
			},
		},
		"ConvertEmptyClusterModeArrays": {
			reason: "Should handle empty ClusterMode arrays gracefully",
			args: args{
				src: &v1beta1.Bucket{
					Status: v1beta1.BucketStatus{
						AtProvider: v1beta1.BucketObservation{
							ServerSideEncryptionConfiguration: []v1beta1.ServerSideEncryptionConfigurationObservation{}, // empty array
						},
					},
				},
				target: &v1beta2.Bucket{},
			},
			want: want{
				target: &v1beta2.Bucket{
					// All fields should remain nil/zero
				},
			},
		},
		"ConvertNilClusterModeArrays": {
			reason: "Should handle nil ClusterMode arrays gracefully",
			args: args{
				src: &v1beta1.Bucket{
					Status: v1beta1.BucketStatus{
						AtProvider: v1beta1.BucketObservation{},
					},
				},
				target: &v1beta2.Bucket{},
			},
			want: want{
				target: &v1beta2.Bucket{
					// All fields should remain nil/zero
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := bucketConverterFromv1beta1Tov1beta2(tc.args.src, tc.args.target)
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

func TestBucketConversion_v1beta2_to_v1beta1(t *testing.T) {
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
				src: &v1beta2.Bucket{
					Status: v1beta2.BucketStatus{
						AtProvider: v1beta2.BucketObservation{
							ServerSideEncryptionConfiguration: &v1beta2.ServerSideEncryptionConfigurationObservation{
								Rule: &v1beta2.ServerSideEncryptionConfigurationRuleObservation{
									ApplyServerSideEncryptionByDefault: &v1beta2.ApplyServerSideEncryptionByDefaultObservation{
										KMSMasterKeyID: stringPtr("example-key-id"),
										// SseAlgorithm should remain nil
									},
									// BucketKeyEnabled should remain nil
								},
							},
						},
					},
				},
				target: &v1beta1.Bucket{},
			},
			want: want{
				target: &v1beta1.Bucket{
					Status: v1beta1.BucketStatus{
						AtProvider: v1beta1.BucketObservation{
							ServerSideEncryptionConfiguration: []v1beta1.ServerSideEncryptionConfigurationObservation{
								{
									Rule: []v1beta1.ServerSideEncryptionConfigurationRuleObservation{
										{
											ApplyServerSideEncryptionByDefault: []v1beta1.ApplyServerSideEncryptionByDefaultObservation{
												{
													KMSMasterKeyID: stringPtr("example-key-id"),
													// SseAlgorithm should remain nil
												},
											},
											// BucketKeyEnabled should remain nil
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"ConvertNilClusterMode": {
			reason: "Should handle nil ClusterMode gracefully",
			args: args{
				src: &v1beta2.Bucket{
					Status: v1beta2.BucketStatus{
						AtProvider: v1beta2.BucketObservation{},
					},
				},
				target: &v1beta1.Bucket{},
			},
			want: want{
				target: &v1beta1.Bucket{
					Status: v1beta1.BucketStatus{
						AtProvider: v1beta1.BucketObservation{
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
			err := bucketConverterFromv1beta2Tov1beta1(tc.args.src, tc.args.target)
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
