// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"sync/atomic"
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

// cacheKeys returns the keys the supplied cache currently holds, sorted so
// that they are comparable.
func cacheKeys(c *AWSCredentialsProviderCache) []string {
	keys := make([]string, 0, len(c.cache))
	for k := range c.cache {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// entryVersion is the version of the credential material a cache entry was
// built from, i.e. everything that keys an entry but does not appear in its
// slot key.
type entryVersion struct {
	Generation  int64
	Source      string
	Fingerprint string
}

// versionAt returns the version of the entry occupying the supplied slot.
func versionAt(t *testing.T, c *AWSCredentialsProviderCache, slotKey string) entryVersion {
	t.Helper()
	e, ok := c.cache[slotKey]
	if !ok {
		t.Fatalf("cache: want an entry at the slot %q, got the slots %v", slotKey, cacheKeys(c))
	}
	return entryVersion{Generation: e.generation, Source: e.source, Fingerprint: e.credsFingerprint}
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
			reason: "Retrieving WebIdentity credentials calls sts:AssumeRoleWithWebIdentity, so it should be cached.",
			pc: &v1beta1.ClusterProviderConfig{Spec: v1beta1.ProviderConfigSpec{
				Credentials: v1beta1.ProviderCredentials{Source: authKeyWebIdentity},
			}},
			want: true,
		},
		"WebIdentityWithRoleChain": {
			reason: "WebIdentity credentials should be cached with or without a role chain on top.",
			pc: &v1beta1.ClusterProviderConfig{Spec: v1beta1.ProviderConfigSpec{
				Credentials: v1beta1.ProviderCredentials{Source: authKeyWebIdentity},
				AssumeRoleChain: []v1beta1.AssumeRoleOptions{
					{RoleARN: ptr.To("arn:aws:iam::123456789012:role/chained")},
				},
			}},
			want: true,
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
		// the rotated material takes the slot of the material it supersedes,
		// which can never be hit again, instead of accumulating next to it.
		if diff := cmp.Diff([]string{"uid:us-east-1"}, cacheKeys(c)); diff != "" {
			t.Errorf("cache slots: -want, +got:\n%s", diff)
		}
		want := entryVersion{Generation: 1, Source: string(xpv2.CredentialsSourceSecret), Fingerprint: "rotated-fingerprint"}
		if diff := cmp.Diff(want, versionAt(t, c, "uid:us-east-1")); diff != "" {
			t.Errorf("cached version: -want, +got:\n%s", diff)
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
		// the region and the ProviderConfig make up the slot, so those get an
		// entry each. The second generation of "uid" is not a slot of its own:
		// it supersedes the first one and takes over its slot.
		if diff := cmp.Diff([]string{"other-uid:us-east-1", "uid:eu-west-1", "uid:us-east-1"}, cacheKeys(c)); diff != "" {
			t.Errorf("cache slots: -want, +got:\n%s", diff)
		}
		want := entryVersion{Generation: 2, Source: string(xpv2.CredentialsSourceSecret), Fingerprint: "fingerprint"}
		if diff := cmp.Diff(want, versionAt(t, c, "uid:us-east-1")); diff != "" {
			t.Errorf("cached version: -want, +got:\n%s", diff)
		}
	})

	t.Run("AStragglingOlderGenerationDoesNotTakeTheSlot", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		accountIDFn := func(context.Context) (string, error) { return "123456789012", nil }

		// the ProviderConfig was edited, so the entry of the new generation
		// holds the slot.
		edited := staticCredsProviderConfig("uid", 2, true)
		editedProvider, _ := newCountingCredsCache(creds)
		if _, err := c.RetrieveCredentials(context.Background(), edited, "us-east-1", editedProvider, meta, accountIDFn); err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}

		// a reconciliation that still carries the pre-edit ProviderConfig, e.g.
		// one that read it before the edit landed, must be served from its own
		// provider, but must not put the cache back on the superseded
		// generation: the two would otherwise evict each other for as long as
		// the stragglers keep arriving.
		straggler := staticCredsProviderConfig("uid", 1, true)
		stragglerCreds := aws.Credentials{AccessKeyID: "pre-edit"}
		stragglerProvider, stragglerCounter := newCountingCredsCache(stragglerCreds)
		got, err := c.RetrieveCredentials(context.Background(), straggler, "us-east-1", stragglerProvider, meta, accountIDFn)
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(Credentials{creds: stragglerCreds, accountID: "123456789012"}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, stragglerCounter.calls); diff != "" {
			t.Errorf("straggling provider Retrieve calls: -want, +got:\n%s", diff)
		}
		want := entryVersion{Generation: 2, Source: string(xpv2.CredentialsSourceSecret), Fingerprint: "fingerprint"}
		if diff := cmp.Diff(want, versionAt(t, c, "uid:us-east-1")); diff != "" {
			t.Errorf("the superseded generation must not take the slot: -want, +got:\n%s", diff)
		}

		// and the current generation is still served from the cache.
		fresh, freshCounter := newCountingCredsCache(aws.Credentials{AccessKeyID: "should-not-be-used"})
		if _, err := c.RetrieveCredentials(context.Background(), edited, "us-east-1", fresh, meta, accountIDFn); err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(0, freshCounter.calls); diff != "" {
			t.Errorf("freshly built provider Retrieve calls: -want, +got:\n%s", diff)
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
				Credentials: v1beta1.ProviderCredentials{Source: authKeyPodIdentity},
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

	// IRSA derives its key material from the projected token file, which
	// getAWSConfig digests into the fingerprint.
	fingerprint, err := fingerprintIRSACreds()
	if err != nil {
		t.Fatalf("fingerprintIRSACreds(): unexpected error: %v", err)
	}
	for i := 0; i < 2; i++ {
		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, awsConfigProvenanceMeta{credsFingerprint: fingerprint}, accountIDFn)
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
	rotatedFingerprint, err := fingerprintIRSACreds()
	if err != nil {
		t.Fatalf("fingerprintIRSACreds(): unexpected error: %v", err)
	}
	if rotatedFingerprint == fingerprint {
		t.Error("fingerprintIRSACreds(): want a distinct fingerprint after the token rotated, got the same")
	}
	rotated, _ := newCountingCredsCache(creds)
	if _, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", rotated, awsConfigProvenanceMeta{credsFingerprint: rotatedFingerprint}, accountIDFn); err != nil {
		t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
	}
	// the pre-rotation token can never be hit again, so its entry must not be
	// left behind: the projected token rotates on the kubelet's schedule,
	// which would otherwise grow the cache with no bound other than maxSize.
	if diff := cmp.Diff([]string{"uid:us-east-1"}, cacheKeys(c)); diff != "" {
		t.Errorf("cache slots: -want, +got:\n%s", diff)
	}
	if diff := cmp.Diff(entryVersion{Generation: 1, Source: authKeyIRSA, Fingerprint: rotatedFingerprint}, versionAt(t, c, "uid:us-east-1")); diff != "" {
		t.Errorf("cached version: -want, +got:\n%s", diff)
	}
}

func TestRetrieveCredentialsWebIdentity(t *testing.T) {
	creds := aws.Credentials{AccessKeyID: "web-identity", SecretAccessKey: "secret", SessionToken: "token"}
	pc := &v1beta1.ClusterProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: "pc", UID: types.UID("uid"), Generation: 1},
		Spec: v1beta1.ProviderConfigSpec{
			Credentials: v1beta1.ProviderCredentials{
				Source: authKeyWebIdentity,
				WebIdentity: &v1beta1.AssumeRoleWithWebIdentityOptions{
					RoleARN: ptr.To("arn:aws:iam::123456789012:role/web-identity"),
					TokenConfig: &v1beta1.WebIdentityTokenConfig{
						Source:    xpv2.CredentialsSourceSecret,
						SecretRef: &xpv2.SecretKeySelector{Key: "token"},
					},
				},
			},
		},
	}
	fingerprint, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte("a-web-identity-token")})
	if err != nil {
		t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
	}
	meta := awsConfigProvenanceMeta{credsFingerprint: fingerprint}

	t.Run("CacheHitDoesNotCallAWS", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		provider, counter := newCountingCredsCache(creds)
		accountIDCalls := 0
		accountIDFn := func(context.Context) (string, error) {
			accountIDCalls++
			return "123456789012", nil
		}

		// prime the cache, then repeat with an equivalent, freshly built
		// provider. The fresh provider must not be consulted, as its Retrieve is
		// the sts:AssumeRoleWithWebIdentity call we are avoiding.
		if _, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, accountIDFn); err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
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
		if diff := cmp.Diff(0, freshCounter.calls); diff != "" {
			t.Errorf("freshly built provider Retrieve calls: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, counter.calls); diff != "" {
			t.Errorf("downstream Retrieve calls: -want, +got:\n%s", diff)
		}
		// the account ID is memoized on the entry instead of being resolved
		// through the separate identity cache, whose key is the resolved
		// credentials triple and therefore rotates with them.
		if diff := cmp.Diff(1, accountIDCalls); diff != "" {
			t.Errorf("AccountIDFn calls: -want, +got:\n%s", diff)
		}
	})

	t.Run("RotatedTokenInvalidatesTheEntry", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		accountIDFn := func(context.Context) (string, error) { return "123456789012", nil }
		provider, _ := newCountingCredsCache(creds)
		if _, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, accountIDFn); err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}

		// the ProviderConfig has not changed, but the token it points at has,
		// e.g. the projected service account token was rotated or the referenced
		// Secret was updated. The stale entry must not be served.
		rotatedFingerprint, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte("a-rotated-web-identity-token")})
		if err != nil {
			t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
		}
		rotatedCreds := aws.Credentials{AccessKeyID: "rotated"}
		rotated, rotatedCounter := newCountingCredsCache(rotatedCreds)
		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", rotated, awsConfigProvenanceMeta{credsFingerprint: rotatedFingerprint}, accountIDFn)
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(Credentials{creds: rotatedCreds, accountID: "123456789012"}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
		// the rotated token takes the slot of the one it supersedes, which can
		// never be hit again, instead of accumulating next to it.
		if diff := cmp.Diff([]string{"uid:us-east-1"}, cacheKeys(c)); diff != "" {
			t.Errorf("cache slots: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(entryVersion{Generation: 1, Source: authKeyWebIdentity, Fingerprint: rotatedFingerprint}, versionAt(t, c, "uid:us-east-1")); diff != "" {
			t.Errorf("cached version: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, rotatedCounter.calls); diff != "" {
			t.Errorf("rotated provider Retrieve calls: -want, +got:\n%s", diff)
		}
	})

	t.Run("MissingFingerprintSkipsTheCache", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		provider, _ := newCountingCredsCache(creds)
		accountIDFn := func(context.Context) (string, error) { return "123456789012", nil }

		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, awsConfigProvenanceMeta{}, accountIDFn)
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(Credentials{creds: creds}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(0, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
	})
}

func TestRetrieveCredentialsConcurrency(t *testing.T) {
	creds := aws.Credentials{AccessKeyID: "assumed", SecretAccessKey: "secret"}
	meta := awsConfigProvenanceMeta{credsFingerprint: "fingerprint"}

	// blockingAccountIDFn returns an AccountIDFn that signals when it has been
	// entered and then blocks until the returned release function is called,
	// standing in for an sts:GetCallerIdentity that is being throttled.
	blockingAccountIDFn := func() (fn AccountIDFn, entered <-chan struct{}, release func()) {
		in, out := make(chan struct{}), make(chan struct{})
		var once sync.Once
		return func(context.Context) (string, error) {
			close(in)
			<-out
			return "123456789012", nil
		}, in, func() { once.Do(func() { close(out) }) }
	}

	t.Run("AnInitializingEntryDoesNotBlockTheOtherKeys", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		slowAccountIDFn, entered, release := blockingAccountIDFn()
		defer release()

		slowProvider, _ := newCountingCredsCache(creds)
		go func() {
			_, _ = c.RetrieveCredentials(context.Background(), staticCredsProviderConfig("uid-slow", 1, true), "us-east-1", slowProvider, meta, slowAccountIDFn)
		}()
		<-entered

		// the reconciliations of the other ProviderConfigs must not queue
		// behind it: no lock is held across the AWS calls that initialize an
		// entry. Otherwise a single throttled credential source would
		// serialize the whole provider, cache hits included, as a pending
		// writer also locks the readers out.
		otherProvider, _ := newCountingCredsCache(creds)
		done := make(chan error, 1)
		go func() {
			_, err := c.RetrieveCredentials(context.Background(), staticCredsProviderConfig("uid-other", 1, true), "us-east-1", otherProvider, meta, func(context.Context) (string, error) {
				return "123456789012", nil
			})
			done <- err
		}()
		select {
		case err := <-done:
			if err != nil {
				t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
			}
		case <-time.After(30 * time.Second):
			t.Fatal("RetrieveCredentials(...): an entry that is being initialized blocks the reconciliations of the other ProviderConfigs")
		}
	})

	t.Run("ConcurrentMissesOfAKeyShareOneInitialization", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		pc := staticCredsProviderConfig("uid", 1, true)
		provider, counter := newCountingCredsCache(creds)
		var accountIDCalls atomic.Int32
		accountIDFn := func(context.Context) (string, error) {
			accountIDCalls.Add(1)
			return "123456789012", nil
		}

		const reconciliations = 8
		var wg sync.WaitGroup
		got := make([]Credentials, reconciliations)
		errs := make([]error, reconciliations)
		start := make(chan struct{})
		for i := 0; i < reconciliations; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				<-start
				got[i], errs[i] = c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, accountIDFn)
			}(i)
		}
		close(start)
		wg.Wait()

		want := Credentials{creds: creds, accountID: "123456789012"}
		for i := 0; i < reconciliations; i++ {
			if errs[i] != nil {
				t.Fatalf("RetrieveCredentials(...): unexpected error: %v", errs[i])
			}
			if diff := cmp.Diff(want, got[i], cmp.AllowUnexported(Credentials{})); diff != "" {
				t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
			}
		}
		// exactly one of them initializes the entry while the others wait for
		// it, instead of each making the same AWS calls.
		if diff := cmp.Diff(int32(1), accountIDCalls.Load()); diff != "" {
			t.Errorf("AccountIDFn calls: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, counter.calls); diff != "" {
			t.Errorf("downstream Retrieve calls: -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
	})

	t.Run("AWaitingReconciliationObservesItsOwnContext", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		pc := staticCredsProviderConfig("uid", 1, true)
		provider, _ := newCountingCredsCache(creds)
		slowAccountIDFn, entered, release := blockingAccountIDFn()
		defer release()

		go func() {
			_, _ = c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, slowAccountIDFn)
		}()
		<-entered

		// an entry is published before it is initialized, so a reconciliation
		// can find one that is not usable yet. Waiting for it must observe the
		// reconciliation's own context instead of outliving it.
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := c.RetrieveCredentials(ctx, pc, "us-east-1", provider, meta, func(context.Context) (string, error) {
			t.Error("AccountIDFn: a reconciliation waiting for an entry must not initialize it itself")
			return "", nil
		})
		if !errors.Is(err, context.Canceled) {
			t.Errorf("RetrieveCredentials(...): want an error wrapping context.Canceled, got %v", err)
		}
	})

	t.Run("AFailedInitializationIsNotCached", func(t *testing.T) {
		c := NewAWSCredentialsProviderCache()
		pc := staticCredsProviderConfig("uid", 1, true)
		provider, _ := newCountingCredsCache(creds)
		errBoom := errors.New("Throttling: Rate exceeded")

		_, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, func(context.Context) (string, error) {
			return "", errBoom
		})
		if !errors.Is(err, errBoom) {
			t.Fatalf("RetrieveCredentials(...): want an error wrapping %v, got %v", errBoom, err)
		}
		// an entry that failed to initialize must not be left behind, or it
		// would be served to everyone else and the failure would never be
		// retried.
		if diff := cmp.Diff(0, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}

		got, err := c.RetrieveCredentials(context.Background(), pc, "us-east-1", provider, meta, func(context.Context) (string, error) {
			return "123456789012", nil
		})
		if err != nil {
			t.Fatalf("RetrieveCredentials(...): unexpected error: %v", err)
		}
		if diff := cmp.Diff(Credentials{creds: creds, accountID: "123456789012"}, got, cmp.AllowUnexported(Credentials{})); diff != "" {
			t.Errorf("RetrieveCredentials(...): -want, +got:\n%s", diff)
		}
		if diff := cmp.Diff(1, len(c.cache)); diff != "" {
			t.Errorf("cache size: -want, +got:\n%s", diff)
		}
	})
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
