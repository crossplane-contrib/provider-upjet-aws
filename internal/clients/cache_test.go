// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go/middleware"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"
	"k8s.io/utils/ptr"
)

var errBoom = errors.New("a")

func TestGetCallerIdentity(t *testing.T) {
	type args struct {
		creds               aws.Credentials
		getCallerIdentityFn GetCallerIdentityFn
		cache               map[string]*callerIdentityCacheEntry
		maxSize             int
	}
	type want struct {
		id    *sts.GetCallerIdentityOutput
		err   error
		cache map[string]*callerIdentityCacheEntry
	}

	sample := &sts.GetCallerIdentityOutput{
		Account: ptr.To("123456789"),
		Arn:     ptr.To("arn:aws:iam::123456789:role/S3Access"),
	}
	ti := time.Now()
	cases := map[string]struct {
		reason string
		args
		want
	}{
		"NotFoundInCacheAndFail": {
			reason: "It should make the API call if the value is not cached.",
			args: args{
				getCallerIdentityFn: func(_ context.Context, _ aws.Config) (*sts.GetCallerIdentityOutput, error) {
					return nil, errBoom
				},
			},
			want: want{
				err: errors.Wrap(errBoom, errGetCallerIdentityFailed),
			},
		},
		"NotFoundInCacheAndSuccess": {
			reason: "It should make the API call if the value is not cached and return the success result.",
			args: args{
				getCallerIdentityFn: func(_ context.Context, _ aws.Config) (*sts.GetCallerIdentityOutput, error) {
					return sample, nil
				},
			},
			want: want{
				id: sample,
			},
		},
		"FoundInCache": {
			reason: "It should not make the API call if the value is cached.",
			args: args{
				creds: aws.Credentials{
					AccessKeyID:     "sampleaccess",
					SecretAccessKey: "samplesecret",
					SessionToken:    "sampletoken",
				},
				cache: map[string]*callerIdentityCacheEntry{
					"sampleaccess:samplesecret:sampletoken": {
						GetCallerIdentityOutput: sample,
						AccessedAt:              ti,
					},
				},
			},
			want: want{
				id: sample,
			},
		},
		"CleanCache": {
			reason: "It should make sure the size of the cache is within the limits after every call and the dustiest one is deleted.",
			args: args{
				getCallerIdentityFn: func(_ context.Context, _ aws.Config) (*sts.GetCallerIdentityOutput, error) {
					return sample, nil
				},
				creds: aws.Credentials{
					AccessKeyID:     "sampleaccess",
					SecretAccessKey: "samplesecret",
					SessionToken:    "sampletoken3",
				},
				cache: map[string]*callerIdentityCacheEntry{
					"sampleaccess:samplesecret:sampletoken": {
						GetCallerIdentityOutput: sample,
						AccessedAt:              ti.Add(-time.Hour * 1),
					},
					"sampleaccess:samplesecret:sampletoken2": {
						GetCallerIdentityOutput: sample,
						AccessedAt:              ti.Add(-time.Hour * 5), // this should be deleted
					},
				},
				maxSize: 2,
			},
			want: want{
				id: sample,
				cache: map[string]*callerIdentityCacheEntry{
					"sampleaccess:samplesecret:sampletoken": {
						GetCallerIdentityOutput: sample,
					},
					"sampleaccess:samplesecret:sampletoken3": {
						GetCallerIdentityOutput: sample,
					},
				},
			},
		},
	}

	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			opts := []CallerIdentityCacheOption{WithGetCallerIdentityFn(tc.getCallerIdentityFn)}
			if tc.args.cache != nil {
				opts = append(opts, WithCache(tc.args.cache))
			}
			if tc.args.maxSize != 0 {
				opts = append(opts, WithMaxSize(tc.args.maxSize))
			}
			c := NewCallerIdentityCache(opts...)
			id, err := c.GetCallerIdentity(context.TODO(), aws.Config{}, tc.args.creds)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Fatalf("%s: GetCallerIdentity(...): err -want, +got: %s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.id, id,
				cmpopts.IgnoreUnexported(sts.GetCallerIdentityOutput{}, middleware.Metadata{})); diff != "" {
				t.Fatalf("%s: GetCallerIdentity(...): -want, +got: %s", tc.reason, diff)
			}
			if tc.want.cache != nil {
				if diff := cmp.Diff(tc.want.cache, c.cache,
					cmpopts.IgnoreFields(callerIdentityCacheEntry{}, "AccessedAt"),
					cmpopts.IgnoreUnexported(sts.GetCallerIdentityOutput{}, middleware.Metadata{})); diff != "" {
					t.Fatalf("%s: GetCallerIdentity(...): -want, +got: %s", tc.reason, diff)
				}
			}
		})
	}
}
