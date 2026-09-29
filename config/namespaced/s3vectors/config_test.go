// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package s3vectors

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestS3vectorsNotFoundDiagnostic(t *testing.T) {
	const accessDenied = "operation error S3Vectors: GetVectorBucket, https response error StatusCode: 403, RequestID: , AccessDeniedException: "

	type args struct {
		diags []*tfprotov6.Diagnostic
	}
	type want struct {
		notFound bool
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"NoDiagnostics": {
			args: args{},
			want: want{notFound: false},
		},
		"StubReadNoAccountFound": {
			args: args{diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "reading S3 Vectors Vector Bucket (arn:aws:s3vectors:us-east-1:000000000000:bucket/xpstub)",
				Detail:   "No account found for the given parameters",
			}}},
			want: want{notFound: true},
		},
		"StubReadAccessDenied": {
			args: args{diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "reading S3 Vectors Vector Bucket (arn:aws:s3vectors:us-east-1:000000000000:bucket/xpstub)",
				Detail:   accessDenied,
			}}},
			want: want{notFound: true},
		},
		"StubIndexReadAccessDenied": {
			args: args{diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "reading S3 Vectors Index (arn:aws:s3vectors:us-east-1:000000000000:bucket/xpstub/index/xpstub)",
				Detail:   "operation error S3Vectors: GetIndex, https response error StatusCode: 403, RequestID: , AccessDeniedException: ",
			}}},
			want: want{notFound: true},
		},
		"RealARNAccessDenied": {
			args: args{diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "reading S3 Vectors Vector Bucket (arn:aws:s3vectors:us-east-1:123456789012:bucket/example)",
				Detail:   accessDenied,
			}}},
			want: want{notFound: false},
		},
		"StubReadOtherError": {
			args: args{diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "reading S3 Vectors Vector Bucket (arn:aws:s3vectors:us-east-1:000000000000:bucket/xpstub)",
				Detail:   "operation error S3Vectors: GetVectorBucket, https response error StatusCode: 500, InternalServerException: ",
			}}},
			want: want{notFound: false},
		},
		"StubReadAccessDeniedWarning": {
			args: args{diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Summary:  "reading S3 Vectors Vector Bucket (arn:aws:s3vectors:us-east-1:000000000000:bucket/xpstub)",
				Detail:   accessDenied,
			}}},
			want: want{notFound: false},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := want{notFound: s3vectorsNotFoundDiagnostic(tc.args.diags)}
			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(want{})); diff != "" {
				t.Errorf("s3vectorsNotFoundDiagnostic(...): -want, +got:\n%s", diff)
			}
		})
	}
}
