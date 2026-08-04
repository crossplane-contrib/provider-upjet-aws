// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

// countingCredentialsProvider is an aws.CredentialsProvider that counts how
// many times it was asked to retrieve the credentials. It stands in for a
// provider that performs an AWS API call, e.g. an
// stscreds.AssumeRoleProvider.
type countingCredentialsProvider struct {
	creds aws.Credentials
	err   error
	calls int
}

func (p *countingCredentialsProvider) Retrieve(context.Context) (aws.Credentials, error) {
	p.calls++
	if p.err != nil {
		return aws.Credentials{}, p.err
	}
	return p.creds, nil
}

// newCountingCredsCache returns an aws.CredentialsCache, i.e. what the AWS SDK
// wraps every configured credential provider with, on top of a
// countingCredentialsProvider.
func newCountingCredsCache(creds aws.Credentials) (*aws.CredentialsCache, *countingCredentialsProvider) {
	p := &countingCredentialsProvider{creds: creds}
	return aws.NewCredentialsCache(p), p
}

func staticCredsProviderConfig(uid string, generation int64, chain bool) *v1beta1.ClusterProviderConfig {
	pc := &v1beta1.ClusterProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "pc",
			UID:        types.UID(uid),
			Generation: generation,
		},
		Spec: v1beta1.ProviderConfigSpec{
			Credentials: v1beta1.ProviderCredentials{
				Source: xpv2.CredentialsSourceSecret,
			},
		},
	}
	if chain {
		pc.Spec.AssumeRoleChain = []v1beta1.AssumeRoleOptions{
			{RoleARN: ptr.To("arn:aws:iam::123456789012:role/chained")},
		}
	}
	return pc
}

func TestCacheableSource(t *testing.T) {
	cases := map[string]struct {
		reason string
		pc     *v1beta1.ClusterProviderConfig
		want   bool
	}{
		"IRSA": {
			reason: "IRSA credentials should be cached.",
			pc: &v1beta1.ClusterProviderConfig{Spec: v1beta1.ProviderConfigSpec{
				Credentials: v1beta1.ProviderCredentials{Source: authKeyIRSA},
			}},
			want: true,
		},
		"WebIdentity": {
			reason: "WebIdentity credentials are not supported by the cache yet.",
			pc: &v1beta1.ClusterProviderConfig{Spec: v1beta1.ProviderConfigSpec{
				Credentials: v1beta1.ProviderCredentials{Source: authKeyWebIdentity},
			}},
			want: false,
		},
		"PodIdentity": {
			reason: "PodIdentity credentials are not supported by the cache yet.",
			pc: &v1beta1.ClusterProviderConfig{Spec: v1beta1.ProviderConfigSpec{
				Credentials: v1beta1.ProviderCredentials{Source: authKeyPodIdentity},
			}},
			want: false,
		},
		"Upbound": {
			reason: "Upbound credentials are not supported by the cache yet.",
			pc: &v1beta1.ClusterProviderConfig{Spec: v1beta1.ProviderConfigSpec{
				Credentials: v1beta1.ProviderCredentials{Source: authKeyUpbound},
			}},
			want: false,
		},
		"StaticWithoutRoleChain": {
			reason: "Retrieving static credentials involves no AWS API call, so there is nothing to cache.",
			pc:     staticCredsProviderConfig("uid", 1, false),
			want:   false,
		},
		"StaticWithRoleChain": {
			reason: "Assuming a role chain on top of static credentials involves an AWS API call, so it should be cached.",
			pc:     staticCredsProviderConfig("uid", 1, true),
			want:   true,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tc.want, cacheableSource(tc.pc)); diff != "" {
				t.Errorf("\n%s\ncacheableSource(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestRetrieveCredentialsStaticWithRoleChain(t *testing.T) {
	creds := aws.Credentials{AccessKeyID: "assumed", SecretAccessKey: "secret", SessionToken: "token"}
	meta := awsConfigProvenanceMeta{credsFingerprint: "fingerprint"}
	pc := staticCredsProviderConfig("uid", 1, true)

	t.Run("CacheMissPopulatesTheCache", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		provider, counter := newCountingCredsCache(creds)
		accountIDCalls := 0
		accountIDFn := func(context.Context) (string, error) {
			accountIDCalls++
			return "123456789012", nil
		}

		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, accountIDFn)
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(Credentials{creds: creds, accountID: "123456789012"}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, accountIDCalls); diff != "" {
			t.Errorf("AccountIDFn calls: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, counter.calls); diff != "" {
			t.Errorf("downstream Retrieve calls: -want, +got:\n%s", diff)
		}
	})

	t.Run("CacheHitDoesNotCallAWS", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		provider, counter := newCountingCredsCache(creds)
		accountIDCalls := 0
		accountIDFn := func(context.Context) (string, error) {
			accountIDCalls++
			return "123456789012", nil
		}

		// prime the cache
		if _, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, accountIDFn); err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		// a subsequent call with an equivalent, freshly built provider should
		// be served from the cache. The fresh provider must not be consulted,
		// as that is the AWS API call we are avoiding.
		fresh, freshCounter := newCountingCredsCache(aws.Credentials{AccessKeyID: "should-not-be-used"})
		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", fresh, meta, accountIDFn)
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(Credentials{creds: creds, accountID: "123456789012"}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, accountIDCalls); diff != "" {
			t.Errorf("AccountIDFn calls: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(0, freshCounter.calls); diff != "" {
			t.Errorf("freshly built provider Retrieve calls: -want, +got:\n%s", diff)
		}
		// the cached aws.CredentialsCache itself caches the credentials of its
		// downstream provider until they expire, so the downstream provider is
		// consulted only once.
		if diff := cmp.Diff(1, counter.calls); diff != "" {
			t.Errorf("downstream Retrieve calls: -want, +got:\n%s", diff)
		}
	})

	t.Run("RotatedCredentialsInvalidateTheEntry", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		accountIDFn := func(context.Context) (string, error) { return "123456789012", nil }
		provider, _ := newCountingCredsCache(creds)
		if _, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, accountIDFn); err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}

		// the ProviderConfig has not changed, but the credential material it
		// points at has, e.g. the referenced Secret was rotated. The stale
		// entry must not be served.
		rotatedCreds := aws.Credentials{AccessKeyID: "rotated"}
		rotated, rotatedCounter := newCountingCredsCache(rotatedCreds)
		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", rotated, awsConfigProvenanceMeta{credsFingerprint: "rotated-fingerprint"}, accountIDFn)
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(Credentials{creds: rotatedCreds, accountID: "123456789012"}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(2, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, rotatedCounter.calls); diff != "" {
			t.Errorf("rotated provider Retrieve calls: -want, +got:\n%s", diff)
		}
	})

	t.Run("DistinctRegionsAndProviderConfigsGetDistinctEntries", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		accountIDFn := func(context.Context) (string, error) { return "123456789012", nil }
		for _, args := range []struct {
			pc     *v1beta1.ClusterProviderConfig
			region string
		}{
			{pc: staticCredsProviderConfig("uid", 1, true), region: "us-east-1"},
			{pc: staticCredsProviderConfig("uid", 1, true), region: "eu-west-1"},
			{pc: staticCredsProviderConfig("uid", 2, true), region: "us-east-1"},
			{pc: staticCredsProviderConfig("other-uid", 1, true), region: "us-east-1"},
		} {
			provider, _ := newCountingCredsCache(creds)
			if _, err := c.RetrieveCredentials(context.Background(), args.pc, args.region, provider, meta, accountIDFn); err != nil {
				t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
			}
		}
		if diff := cmp.Diff(4, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
	})

	t.Run("AccountIDFailureIsNotCached", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		provider, _ := newCountingCredsCache(creds)
		accountIDFn := func(context.Context) (string, error) { return "", errBoom }

		_, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, accountIDFn)
		if diff := cmp.Diff(errors.Wrap(errBoom, errGetAccountID).Error(), err.Error()); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want error, +got error:\n%s", diff)
		}
		if diff := cmp.Diff(0, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
	})

	t.Run("MissingFingerprintSkipsTheCache", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		provider, counter := newCountingCredsCache(creds)
		accountIDCalls := 0
		accountIDFn := func(context.Context) (string, error) {
			accountIDCalls++
			return "123456789012", nil
		}

		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, awsConfigProvenanceMeta{}, accountIDFn)
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		// credentials are still returned, they are just not cached, and the
		// account ID is left to the separate identity cache.
		if diff := cmp.Diff(Credentials{creds: creds}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(0, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(0, accountIDCalls); diff != "" {
			t.Errorf("AccountIDFn calls: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, counter.calls); diff != "" {
			t.Errorf("downstream Retrieve calls: -want, +got:\n%s", diff)
		}
	})
}

func TestRetrieveCredentialsUncached(t *testing.T) {
	creds := aws.Credentials{AccessKeyID: "static", SecretAccessKey: "secret"}
	meta := awsConfigProvenanceMeta{credsFingerprint: "fingerprint"}

	cases := map[string]struct {
		reason   string
		pc       *v1beta1.ClusterProviderConfig
		provider aws.CredentialsProvider
	}{
		"StaticWithoutRoleChain": {
			reason: "Static credentials without a role chain should not be cached.",
			pc:     staticCredsProviderConfig("uid", 1, false),
		},
		"UnsupportedSource": {
			reason: "Credential sources that the cache does not support yet should not be cached.",
			pc: &v1beta1.ClusterProviderConfig{Spec: v1beta1.ProviderConfigSpec{
				Credentials: v1beta1.ProviderCredentials{Source: authKeyWebIdentity},
			}},
		},
		"NotACredentialsCache": {
			reason:   "A provider that is not an aws.CredentialsCache cannot be cached.",
			pc:       staticCredsProviderConfig("uid", 1, true),
			provider: &countingCredentialsProvider{creds: creds},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewAWSCredentialsProviderCache()
			provider := tc.provider
			if provider == nil {
				provider, _ = newCountingCredsCache(creds)
			}
			accountIDCalls := 0
			accountIDFn := func(context.Context) (string, error) {
				accountIDCalls++
				return "123456789012", nil
			}

			got, err := c.RetrieveCredentials(context.Background(), tc.pc, "us-east-1", provider, meta, accountIDFn)
			if err != nil {
				t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
			}
			// the credentials are returned from the downstream provider, but
			// the account ID is left to the separate identity cache.
			if diff := cmp.Diff(Credentials{creds: creds}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
				t.Errorf("\n%s\nRetrieveCredentials(...): -want, +got:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(0, len(c.cache)); diff != "" {
				t.Errorf("\n%s\ncache size: -want, +got:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(0, accountIDCalls); diff != "" {
				t.Errorf("\n%s\nAccountIDFn calls: -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestRetrieveCredentialsIRSA(t *testing.T) {
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("a-web-identity-token"), 0o600); err != nil {
		t.Fatalf("cannot write the token file: %v", err)
	}
	t.Setenv(envWebIdentityTokenFile, tokenFile)
	t.Setenv(envWebIdentityRoleARN, "arn:aws:iam::123456789012:role/irsa")

	creds := aws.Credentials{AccessKeyID: "irsa", SecretAccessKey: "secret"}
	pc := &v1beta1.ClusterProviderConfig{
		ObjectMeta: metav1.ObjectMeta{UID: types.UID("uid"), Generation: 1},
		Spec: v1beta1.ProviderConfigSpec{
			Credentials: v1beta1.ProviderCredentials{Source: authKeyIRSA},
		},
	}

	c := NewAWSCredentialsProviderCache()
	provider, counter := newCountingCredsCache(creds)
	accountIDCalls := 0
	accountIDFn := func(context.Context) (string, error) {
		accountIDCalls++
		return "123456789012", nil
	}

	// IRSA does not need a credentials fingerprint, it derives its key material
	// from the projected token file.
	for i := 0; i < 2; i++ {
		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, awsConfigProvenanceMeta{}, accountIDFn)
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(Credentials{creds: creds, accountID: "123456789012"}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
	}
	if diff := cmp.Diff(1, len(c.cache)); diff != "" {
		t.Errorf("cache size: -want, +got:\n%s", diff)
	}
	if diff := cmp.Diff(1, accountIDCalls); diff != "" {
		t.Errorf("AccountIDFn calls: -want, +got:\n%s", diff)
	}
	if diff := cmp.Diff(1, counter.calls); diff != "" {
		t.Errorf("downstream Retrieve calls: -want, +got:\n%s", diff)
	}

	// a rotated projected token must produce a distinct entry
	if err := os.WriteFile(tokenFile, []byte("a-rotated-web-identity-token"), 0o600); err != nil {
		t.Fatalf("cannot write the token file: %v", err)
	}
	rotated, _ := newCountingCredsCache(creds)
	if _, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", rotated, awsConfigProvenanceMeta{}, accountIDFn); err != nil {
		t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
	}
	if diff := cmp.Diff(2, len(c.cache)); diff != "" {
		t.Errorf("cache size: -want, +got:\n%s", diff)
	}
}

func TestCredentialsProviderCacheMakeRoom(t *testing.T) {
	creds := aws.Credentials{AccessKeyID: "assumed"}
	accountIDFn := func(context.Context) (string, error) { return "123456789012", nil }

	provider, _ := newCountingCredsCache(creds)
	oldest := &awsCredentialsProviderCacheEntry{awsCredCache: provider}
	oldest.accountID.Store("123456789012")
	oldest.accessedAt.Store(time.Now().Add(-time.Hour))
	newest := &awsCredentialsProviderCacheEntry{awsCredCache: provider}
	newest.accountID.Store("123456789012")
	newest.accessedAt.Store(time.Now())

	c := NewAWSCredentialsProviderCache(WithCacheMaxSize(2), WithCacheStore(map[string]*awsCredentialsProviderCacheEntry{
		"oldest": oldest,
		"newest": newest,
	}))

	fresh, _ := newCountingCredsCache(creds)
	if _, err := c.RetrieveCredentials(context.Background(), staticCredsProviderConfig("uid", 1, true), "us-east-1", fresh, awsConfigProvenanceMeta{credsFingerprint: "fingerprint"}, accountIDFn); err != nil {
		t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
	}
	if diff := cmp.Diff(2, len(c.cache)); diff != "" {
		t.Errorf("cache size: -want, +got:\n%s", diff)
	}
	if _, ok := c.cache["oldest"]; ok {
		t.Error("the least recently accessed entry should have been evicted")
	}
	if _, ok := c.cache["newest"]; !ok {
		t.Error("the most recently accessed entry should have been kept")
	}
}
