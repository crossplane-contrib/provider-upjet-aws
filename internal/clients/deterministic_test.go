// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"testing"
)

func TestIsDeterministicExternalName(t *testing.T) {
	cases := map[string]struct {
		resourceType string
		want         bool
	}{
		"S3BucketVersioning": {
			resourceType: "aws_s3_bucket_versioning",
			want:         true,
		},
		"S3BucketPublicAccessBlock": {
			resourceType: "aws_s3_bucket_public_access_block",
			want:         true,
		},
		"S3BucketPolicy": {
			resourceType: "aws_s3_bucket_policy",
			want:         true,
		},
		"NonDeterministicInstance": {
			resourceType: "aws_instance",
			want:         false,
		},
		"NonDeterministicVpc": {
			resourceType: "aws_vpc",
			want:         false,
		},
		"UnknownResource": {
			resourceType: "aws_unknown_resource",
			want:         false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := IsDeterministicExternalName(tc.resourceType)
			if got != tc.want {
				t.Errorf("IsDeterministicExternalName(%q) = %v, want %v", tc.resourceType, got, tc.want)
			}
		})
	}
}
