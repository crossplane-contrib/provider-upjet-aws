// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

type fakeCredentialsProvider struct {
	creds  aws.Credentials
	err    error
	called int32
}

func (f *fakeCredentialsProvider) Retrieve(context.Context) (aws.Credentials, error) {
	atomic.AddInt32(&f.called, 1)
	if f.err != nil {
		return aws.Credentials{}, f.err
	}
	return f.creds, nil
}

func newClusterProviderConfig(source xpv1.CredentialsSource, uid string, generation int64) *v1beta1.ClusterProviderConfig {
	return &v1beta1.ClusterProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			UID:        types.UID(uid),
			Generation: generation,
		},
		Spec: v1beta1.ProviderConfigSpec{
			Credentials: v1beta1.ProviderCredentials{
				Source: source,
			},
		},
	}
}

func TestAWSCredentialsProviderCacheRetrieveCredentialsIRSACachesAccountID(t *testing.T) {
	tokenFile := filepath.Join(t.TempDir(), "web-identity-token")
	if err := os.WriteFile(tokenFile, []byte("token"), 0o600); err != nil {
		t.Fatalf("os.WriteFile(tokenFile): %v", err)
	}
	t.Setenv("AWS_WEB_IDENTITY_TOKEN_FILE", tokenFile)
	t.Setenv("AWS_ROLE_ARN", "arn:aws:iam::111111111111:role/test")

	provider := &fakeCredentialsProvider{
		creds: aws.Credentials{
			AccessKeyID:     "access-key",
			SecretAccessKey: "secret-key",
			SessionToken:    "session-token",
		},
	}
	credCache := aws.NewCredentialsCache(provider)
	cache := NewAWSCredentialsProviderCache()
	pc := newClusterProviderConfig(xpv1.CredentialsSource(authKeyIRSA), "pc-irsa", 1)

	accountIDCalls := 0
	accountIDFn := func(context.Context) (string, error) {
		accountIDCalls++
		return "111111111111", nil
	}

	first, err := cache.RetrieveCredentials(context.Background(), pc, "us-west-2", credCache, accountIDFn)
	if err != nil {
		t.Fatalf("RetrieveCredentials(first): %v", err)
	}
	second, err := cache.RetrieveCredentials(context.Background(), pc, "us-west-2", credCache, accountIDFn)
	if err != nil {
		t.Fatalf("RetrieveCredentials(second): %v", err)
	}

	if first.accountID != "111111111111" || second.accountID != "111111111111" {
		t.Fatalf("expected cached account ID to be preserved, got first=%q second=%q", first.accountID, second.accountID)
	}
	if accountIDCalls != 1 {
		t.Fatalf("accountIDFn calls = %d, want 1", accountIDCalls)
	}
	if got := atomic.LoadInt32(&provider.called); got != 1 {
		t.Fatalf("credentials provider retrieve calls = %d, want 1", got)
	}
	if len(cache.cache) != 1 {
		t.Fatalf("cache size = %d, want 1", len(cache.cache))
	}
}

func TestAWSCredentialsProviderCacheRetrieveCredentialsNonIRSASkipsAccountIDFn(t *testing.T) {
	provider := &fakeCredentialsProvider{
		creds: aws.Credentials{
			AccessKeyID:     "access-key",
			SecretAccessKey: "secret-key",
			SessionToken:    "session-token",
		},
	}
	cache := NewAWSCredentialsProviderCache()
	pc := newClusterProviderConfig(xpv1.CredentialsSourceSecret, "pc-secret", 1)

	accountIDCalled := false
	got, err := cache.RetrieveCredentials(context.Background(), pc, "us-west-2", provider, func(context.Context) (string, error) {
		accountIDCalled = true
		return "111111111111", nil
	})
	if err != nil {
		t.Fatalf("RetrieveCredentials(non-irsa): %v", err)
	}

	if accountIDCalled {
		t.Fatalf("accountIDFn should not be called for non-IRSA credentials")
	}
	if got.accountID != "" {
		t.Fatalf("account ID = %q, want empty", got.accountID)
	}
	if len(cache.cache) != 0 {
		t.Fatalf("cache size = %d, want 0", len(cache.cache))
	}
}

func TestAWSCredentialsProviderCacheRetrieveCredentialsIRSAWithoutCredentialCacheSkipsAccountIDFn(t *testing.T) {
	provider := &fakeCredentialsProvider{
		creds: aws.Credentials{
			AccessKeyID:     "access-key",
			SecretAccessKey: "secret-key",
			SessionToken:    "session-token",
		},
	}
	cache := NewAWSCredentialsProviderCache()
	pc := newClusterProviderConfig(xpv1.CredentialsSource(authKeyIRSA), "pc-irsa-no-cache", 1)

	accountIDCalled := false
	got, err := cache.RetrieveCredentials(context.Background(), pc, "us-west-2", provider, func(context.Context) (string, error) {
		accountIDCalled = true
		return "111111111111", nil
	})
	if err != nil {
		t.Fatalf("RetrieveCredentials(irsa-no-cache): %v", err)
	}

	if accountIDCalled {
		t.Fatalf("accountIDFn should not be called when credentials provider is not aws.CredentialsCache")
	}
	if got.accountID != "" {
		t.Fatalf("account ID = %q, want empty", got.accountID)
	}
	if len(cache.cache) != 0 {
		t.Fatalf("cache size = %d, want 0", len(cache.cache))
	}
}
