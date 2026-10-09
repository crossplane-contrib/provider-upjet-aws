// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/v2/pkg/test"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

// stubTokenRetriever is an stscreds.IdentityTokenRetriever serving a canned
// token.
type stubTokenRetriever struct {
	token []byte
	err   error
}

func (s stubTokenRetriever) GetIdentityToken() ([]byte, error) {
	return s.token, s.err
}

// unsetEnv removes an environment variable for the duration of the test.
// t.Setenv is called first, purely to register its restoring cleanup, which
// handles a previously unset variable correctly.
func unsetEnv(t *testing.T, key string) {
	t.Helper()
	t.Setenv(key, "")
	if err := os.Unsetenv(key); err != nil {
		t.Fatalf("cannot unset the environment variable %s: %v", key, err)
	}
}

// isolateAWSEnv points the AWS SDK's shared configuration at files that do not
// exist and clears the environment that would otherwise leak the developer's or
// the CI runner's own AWS setup into config.LoadDefaultConfig.
func isolateAWSEnv(t *testing.T) {
	t.Helper()
	missing := filepath.Join(t.TempDir(), "does-not-exist")
	t.Setenv("AWS_CONFIG_FILE", missing)
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", missing)
	for _, k := range []string{"AWS_PROFILE", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN", envWebIdentityTokenFile, envWebIdentityRoleARN} {
		unsetEnv(t, k)
	}
}

// writeTokenFile writes a web identity token to a file in a fresh temporary
// directory and returns its path.
func writeTokenFile(t *testing.T, token string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(path, []byte(token), 0o600); err != nil {
		t.Fatalf("cannot write the token file: %v", err)
	}
	return path
}

// tokenServingClient returns a client.Client that serves the supplied token
// from every Secret it is asked for, under the given key. It honors context
// cancellation, like a real client does.
func tokenServingClient(token []byte, key string) *test.MockClient {
	return &test.MockClient{
		MockGet: func(ctx context.Context, _ client.ObjectKey, obj client.Object) error {
			if err := ctx.Err(); err != nil {
				return err
			}
			s, ok := obj.(*corev1.Secret)
			if !ok {
				return fmt.Errorf("unexpected object type %T", obj)
			}
			s.Data = map[string][]byte{key: token}
			return nil
		},
	}
}

func webIdentitySpec(tokenConfig *v1beta1.WebIdentityTokenConfig) *v1beta1.ProviderConfigSpec {
	return &v1beta1.ProviderConfigSpec{
		Credentials: v1beta1.ProviderCredentials{
			Source: authKeyWebIdentity,
			WebIdentity: &v1beta1.AssumeRoleWithWebIdentityOptions{
				RoleARN:     ptr.To("arn:aws:iam::123456789012:role/web-identity"),
				TokenConfig: tokenConfig,
			},
		},
	}
}

func TestFingerprintStaticCreds(t *testing.T) {
	base := aws.Credentials{AccessKeyID: "AKID", SecretAccessKey: "secret", SessionToken: "token"}

	t.Run("IsStableAndDoesNotLeakTheMaterial", func(t *testing.T) {
		got := fingerprintStaticCreds(base)
		if got != fingerprintStaticCreds(base) {
			t.Error("fingerprintStaticCreds(...) is not stable for the same credentials")
		}
		if len(got) != 2*sha256.Size {
			t.Errorf("fingerprintStaticCreds(...): want a hex sha256 digest, got %q", got)
		}
		for _, s := range []string{base.AccessKeyID, base.SecretAccessKey, base.SessionToken} {
			if strings.Contains(got, s) {
				t.Errorf("fingerprintStaticCreds(...) leaks the credential material %q", s)
			}
		}
	})

	t.Run("IsKeyedToThisProcess", func(t *testing.T) {
		// the digest must not be reproducible from the credential material
		// alone, so that it cannot be confirmed by someone holding a candidate
		// credential and reading the cache keys off the debug logs.
		h := sha256.New()
		for _, s := range []string{fpDomainStaticCreds, base.AccessKeyID, base.SecretAccessKey, base.SessionToken} {
			_ = binary.Write(h, binary.BigEndian, uint64(len(s)))
			_, _ = h.Write([]byte(s))
		}
		if unkeyed := fmt.Sprintf("%x", h.Sum(nil)); fingerprintStaticCreds(base) == unkeyed {
			t.Error("fingerprintStaticCreds(...) is an unkeyed digest of the credential material")
		}
	})

	t.Run("DiffersForEveryField", func(t *testing.T) {
		cases := map[string]aws.Credentials{
			"AccessKeyID":     {AccessKeyID: "other", SecretAccessKey: base.SecretAccessKey, SessionToken: base.SessionToken},
			"SecretAccessKey": {AccessKeyID: base.AccessKeyID, SecretAccessKey: "other", SessionToken: base.SessionToken},
			"SessionToken":    {AccessKeyID: base.AccessKeyID, SecretAccessKey: base.SecretAccessKey, SessionToken: "other"},
			// the fields are separated while hashing, so shifting a field
			// boundary must change the digest
			"ShiftedBoundary": {AccessKeyID: base.AccessKeyID + base.SecretAccessKey, SecretAccessKey: "", SessionToken: base.SessionToken},
		}
		want := fingerprintStaticCreds(base)
		for name, creds := range cases {
			t.Run(name, func(t *testing.T) {
				if got := fingerprintStaticCreds(creds); got == want {
					t.Errorf("fingerprintStaticCreds(...): want a digest different from %q, got the same", want)
				}
			})
		}
	})

	t.Run("FramingIsUnambiguous", func(t *testing.T) {
		// AWS credentials never contain a NUL byte, but the framing must not
		// depend on that: these two collided when the fields were merely
		// suffixed with a NUL separator instead of length-prefixed.
		cases := map[string][2]aws.Credentials{
			"NULInsideAField": {
				{AccessKeyID: "a\x00b", SecretAccessKey: "c"},
				{AccessKeyID: "a", SecretAccessKey: "b\x00c"},
			},
			"EmptyVersusAbsent": {
				{AccessKeyID: "a", SecretAccessKey: "", SessionToken: "b"},
				{AccessKeyID: "a", SecretAccessKey: "b", SessionToken: ""},
			},
		}
		for name, pair := range cases {
			t.Run(name, func(t *testing.T) {
				if a, b := fingerprintStaticCreds(pair[0]), fingerprintStaticCreds(pair[1]); a == b {
					t.Errorf("fingerprintStaticCreds(...): distinct credentials digested to the same value %q", a)
				}
			})
		}
	})
}

func TestFingerprintIdentityToken(t *testing.T) {
	const token = "a-web-identity-token"

	t.Run("IsStableAndDoesNotLeakTheToken", func(t *testing.T) {
		got, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte(token)})
		if err != nil {
			t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
		}
		again, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte(token)})
		if err != nil {
			t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
		}
		if got != again {
			t.Error("fingerprintIdentityToken(...) is not stable for the same token")
		}
		if len(got) != 2*sha256.Size {
			t.Errorf("fingerprintIdentityToken(...): want a hex sha256 digest, got %q", got)
		}
		if strings.Contains(got, token) {
			t.Errorf("fingerprintIdentityToken(...) leaks the token %q", token)
		}
	})

	t.Run("IsKeyedToThisProcess", func(t *testing.T) {
		// a web identity token is a bearer credential, so its digest must not be
		// confirmable by someone holding a candidate token and reading the cache
		// keys off the debug logs.
		h := sha256.New()
		for _, s := range []string{fpDomainWebIdentityToken, token} {
			_ = binary.Write(h, binary.BigEndian, uint64(len(s)))
			_, _ = h.Write([]byte(s))
		}
		got, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte(token)})
		if err != nil {
			t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
		}
		if unkeyed := fmt.Sprintf("%x", h.Sum(nil)); got == unkeyed {
			t.Error("fingerprintIdentityToken(...) is an unkeyed digest of the token")
		}
	})

	t.Run("DiffersForARotatedToken", func(t *testing.T) {
		got, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte(token)})
		if err != nil {
			t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
		}
		rotated, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte(token + "-rotated")})
		if err != nil {
			t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
		}
		if got == rotated {
			t.Errorf("fingerprintIdentityToken(...): want a digest different from %q for a rotated token, got the same", got)
		}
	})

	t.Run("IsDomainSeparatedFromTheOtherSources", func(t *testing.T) {
		// every credential source contributes its fingerprint to the very same
		// cache key component, so digests must not collide across sources even
		// when the digested bytes are identical.
		got, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte(token)})
		if err != nil {
			t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
		}
		if other := fingerprintStaticCreds(aws.Credentials{AccessKeyID: token}); got == other {
			t.Error("fingerprintIdentityToken(...) collides with fingerprintStaticCreds(...) for the same material")
		}
		if other := fingerprintMaterial(fpDomainIRSA, []byte(token)); got == other {
			t.Error("fingerprintIdentityToken(...) collides with the IRSA domain for the same material")
		}
	})

	t.Run("RetrievalFailure", func(t *testing.T) {
		if _, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{err: errBoom}); err == nil {
			t.Error("fingerprintIdentityToken(...): want an error when the token cannot be read, got none")
		}
	})
}

func TestFingerprintIRSACreds(t *testing.T) {
	const token = "an-irsa-token"

	t.Run("DiffersForEveryInput", func(t *testing.T) {
		// unlike the WebIdentity source, IRSA takes the token path and the role
		// ARN from the environment rather than from the ProviderConfig spec, so
		// neither is covered by the spec's generation and both must contribute
		// to the digest.
		tokenFile := writeTokenFile(t, token)
		otherFile := writeTokenFile(t, token)
		t.Setenv(envWebIdentityTokenFile, tokenFile)
		t.Setenv(envWebIdentityRoleARN, "arn:aws:iam::123456789012:role/irsa")

		want, err := fingerprintIRSACreds()
		if err != nil {
			t.Fatalf("fingerprintIRSACreds(): unexpected error: %v", err)
		}
		if strings.Contains(want, token) {
			t.Errorf("fingerprintIRSACreds() leaks the token %q", token)
		}

		cases := map[string]func(){
			"RotatedToken": func() {
				if err := os.WriteFile(tokenFile, []byte(token+"-rotated"), 0o600); err != nil {
					t.Fatalf("cannot write the token file: %v", err)
				}
			},
			// the same token content served from a different path is a different
			// identity as far as IRSA is concerned, because the path selects it.
			"DifferentTokenPath": func() { t.Setenv(envWebIdentityTokenFile, otherFile) },
			"DifferentRoleARN":   func() { t.Setenv(envWebIdentityRoleARN, "arn:aws:iam::123456789012:role/other") },
		}
		for name, mutate := range cases {
			t.Run(name, func(t *testing.T) {
				mutate()
				got, err := fingerprintIRSACreds()
				if err != nil {
					t.Fatalf("fingerprintIRSACreds(): unexpected error: %v", err)
				}
				if got == want {
					t.Errorf("fingerprintIRSACreds(): want a digest different from %q, got the same", want)
				}
			})
		}
	})

	t.Run("MissingTokenFileEnvVar", func(t *testing.T) {
		unsetEnv(t, envWebIdentityTokenFile)
		if _, err := fingerprintIRSACreds(); err == nil {
			t.Errorf("fingerprintIRSACreds(): want an error when %s is unset, got none", envWebIdentityTokenFile)
		}
	})

	t.Run("UnreadableTokenFile", func(t *testing.T) {
		t.Setenv(envWebIdentityTokenFile, filepath.Join(t.TempDir(), "does-not-exist"))
		if _, err := fingerprintIRSACreds(); err == nil {
			t.Error("fingerprintIRSACreds(): want an error when the token file cannot be read, got none")
		}
	})
}

func TestWebIdentityTokenRetriever(t *testing.T) {
	const token = "a-web-identity-token"

	t.Run("Filesystem", func(t *testing.T) {
		isolateAWSEnv(t)
		path := writeTokenFile(t, token)
		r, err := webIdentityTokenRetriever(context.Background(), webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
			Source: xpv2.CredentialsSourceFilesystem,
			Fs:     &xpv2.FsSelector{Path: path},
		}).Credentials.WebIdentity, nil)
		if err != nil {
			t.Fatalf("webIdentityTokenRetriever(...): unexpected error: %v", err)
		}
		got, err := r.GetIdentityToken()
		if err != nil {
			t.Fatalf("GetIdentityToken(): unexpected error: %v", err)
		}
		if string(got) != token {
			t.Errorf("GetIdentityToken(): want %q, got %q", token, string(got))
		}
	})

	t.Run("Secret", func(t *testing.T) {
		isolateAWSEnv(t)
		r, err := webIdentityTokenRetriever(context.Background(), webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
			Source:    xpv2.CredentialsSourceSecret,
			SecretRef: &xpv2.SecretKeySelector{Key: "token"},
		}).Credentials.WebIdentity, tokenServingClient([]byte(token), "token"))
		if err != nil {
			t.Fatalf("webIdentityTokenRetriever(...): unexpected error: %v", err)
		}
		got, err := r.GetIdentityToken()
		if err != nil {
			t.Fatalf("GetIdentityToken(): unexpected error: %v", err)
		}
		if string(got) != token {
			t.Errorf("GetIdentityToken(): want %q, got %q", token, string(got))
		}
	})

	t.Run("LegacyEnvTokenFile", func(t *testing.T) {
		isolateAWSEnv(t)
		t.Setenv(envWebIdentityTokenFile, writeTokenFile(t, token))
		t.Setenv(envWebIdentityRoleARN, "arn:aws:iam::123456789012:role/irsa")
		r, err := webIdentityTokenRetriever(context.Background(), webIdentitySpec(nil).Credentials.WebIdentity, nil)
		if err != nil {
			t.Fatalf("webIdentityTokenRetriever(...): unexpected error: %v", err)
		}
		got, err := r.GetIdentityToken()
		if err != nil {
			t.Fatalf("GetIdentityToken(): unexpected error: %v", err)
		}
		if string(got) != token {
			t.Errorf("GetIdentityToken(): want %q, got %q", token, string(got))
		}
	})

	t.Run("LegacyWithoutTheTokenFileEnvVar", func(t *testing.T) {
		isolateAWSEnv(t)
		if _, err := webIdentityTokenRetriever(context.Background(), webIdentitySpec(nil).Credentials.WebIdentity, nil); err == nil {
			t.Errorf("webIdentityTokenRetriever(...): want an error when neither tokenConfig nor %s is configured, got none", envWebIdentityTokenFile)
		}
	})

	t.Run("PartialIRSAEnvironment", func(t *testing.T) {
		// the AWS SDK refuses to build a default config when the token file env
		// var is set without the role ARN one, so this combination is rejected
		// with a dedicated error.
		isolateAWSEnv(t)
		t.Setenv(envWebIdentityTokenFile, writeTokenFile(t, token))
		_, err := webIdentityTokenRetriever(context.Background(), webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
			Source:    xpv2.CredentialsSourceSecret,
			SecretRef: &xpv2.SecretKeySelector{Key: "token"},
		}).Credentials.WebIdentity, tokenServingClient([]byte(token), "token"))
		if err == nil {
			t.Errorf("webIdentityTokenRetriever(...): want an error when %s is set without %s, got none", envWebIdentityTokenFile, envWebIdentityRoleARN)
		}
	})

	t.Run("SurvivesTheReconcileThatBuiltIt", func(t *testing.T) {
		// The credential cache keeps the credential provider owning this
		// retriever alive across reconciliations, so the retriever must not hold
		// on to the context of the reconcile that constructed it. Otherwise
		// every credential refresh after that reconcile returns fails with
		// "context canceled", for the whole lifetime of the cache entry and
		// without any chance of self-healing, as the cache key does not change.
		isolateAWSEnv(t)
		ctx, cancel := context.WithCancel(context.Background())
		r, err := webIdentityTokenRetriever(ctx, webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
			Source:    xpv2.CredentialsSourceSecret,
			SecretRef: &xpv2.SecretKeySelector{Key: "token"},
		}).Credentials.WebIdentity, tokenServingClient([]byte(token), "token"))
		if err != nil {
			t.Fatalf("webIdentityTokenRetriever(...): unexpected error: %v", err)
		}
		cancel()

		got, err := r.GetIdentityToken()
		if err != nil {
			t.Fatalf("GetIdentityToken() after the constructing context was canceled: unexpected error: %v", err)
		}
		if string(got) != token {
			t.Errorf("GetIdentityToken(): want %q, got %q", token, string(got))
		}
	})
}

func TestUseWebIdentityToken(t *testing.T) {
	const token = "a-web-identity-token"
	roleARN := "arn:aws:iam::123456789012:role/web-identity"

	// the fingerprint digests the token only. Its location is part of the
	// ProviderConfig spec, hence covered by the spec's generation in the cache
	// key, so the same token served from any source fingerprints identically.
	wantFingerprint, err := fingerprintIdentityToken(context.Background(), stubTokenRetriever{token: []byte(token)})
	if err != nil {
		t.Fatalf("fingerprintIdentityToken(...): unexpected error: %v", err)
	}

	cases := map[string]struct {
		reason string
		spec   func(t *testing.T) *v1beta1.ProviderConfigSpec
		kube   client.Client
	}{
		"Filesystem": {
			reason: "A token read from the filesystem should be fingerprinted.",
			spec: func(t *testing.T) *v1beta1.ProviderConfigSpec {
				return webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
					Source: xpv2.CredentialsSourceFilesystem,
					Fs:     &xpv2.FsSelector{Path: writeTokenFile(t, token)},
				})
			},
		},
		"Secret": {
			reason: "A token read from a Secret should be fingerprinted.",
			spec: func(*testing.T) *v1beta1.ProviderConfigSpec {
				return webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
					Source:    xpv2.CredentialsSourceSecret,
					SecretRef: &xpv2.SecretKeySelector{Key: "token"},
				})
			},
			kube: tokenServingClient([]byte(token), "token"),
		},
		"LegacyEnvTokenFile": {
			reason: "A token read from the deprecated environment variable configuration should be fingerprinted.",
			spec: func(t *testing.T) *v1beta1.ProviderConfigSpec {
				t.Setenv(envWebIdentityTokenFile, writeTokenFile(t, token))
				t.Setenv(envWebIdentityRoleARN, roleARN)
				return webIdentitySpec(nil)
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			isolateAWSEnv(t)
			cfg, fingerprint, err := UseWebIdentityToken(context.Background(), "us-east-1", tc.spec(t), tc.kube)
			if err != nil {
				t.Fatalf("\n%s\nUseWebIdentityToken(...): unexpected error: %v", tc.reason, err)
			}
			if cfg == nil {
				t.Fatalf("\n%s\nUseWebIdentityToken(...): want a config, got nil", tc.reason)
			}
			// the credential cache only accepts an *aws.CredentialsCache, as
			// keeping the SDK's own refreshing credential cache alive across
			// reconciliations is the whole point of caching it.
			if _, ok := cfg.Credentials.(*aws.CredentialsCache); !ok {
				t.Errorf("\n%s\nUseWebIdentityToken(...): want an *aws.CredentialsCache, got %T", tc.reason, cfg.Credentials)
			}
			if fingerprint != wantFingerprint {
				t.Errorf("\n%s\nUseWebIdentityToken(...): want the fingerprint %q, got %q", tc.reason, wantFingerprint, fingerprint)
			}
		})
	}

	t.Run("RotatedTokenChangesTheFingerprint", func(t *testing.T) {
		isolateAWSEnv(t)
		path := writeTokenFile(t, token)
		spec := webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
			Source: xpv2.CredentialsSourceFilesystem,
			Fs:     &xpv2.FsSelector{Path: path},
		})
		_, before, err := UseWebIdentityToken(context.Background(), "us-east-1", spec, nil)
		if err != nil {
			t.Fatalf("UseWebIdentityToken(...): unexpected error: %v", err)
		}
		if err := os.WriteFile(path, []byte(token+"-rotated"), 0o600); err != nil {
			t.Fatalf("cannot write the token file: %v", err)
		}
		_, after, err := UseWebIdentityToken(context.Background(), "us-east-1", spec, nil)
		if err != nil {
			t.Fatalf("UseWebIdentityToken(...): unexpected error: %v", err)
		}
		if before == after {
			t.Errorf("UseWebIdentityToken(...): want a fingerprint different from %q after the token rotated, got the same", before)
		}
	})

	t.Run("UnreadableToken", func(t *testing.T) {
		// a token that cannot be read fails the whole config construction: the
		// credentials could not have been retrieved in this reconcile either
		// way, and continuing with an empty fingerprint would silently give up
		// on caching.
		isolateAWSEnv(t)
		spec := webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
			Source: xpv2.CredentialsSourceFilesystem,
			Fs:     &xpv2.FsSelector{Path: filepath.Join(t.TempDir(), "does-not-exist")},
		})
		if _, _, err := UseWebIdentityToken(context.Background(), "us-east-1", spec, nil); err == nil {
			t.Error("UseWebIdentityToken(...): want an error when the token cannot be read, got none")
		}
	})

	t.Run("MissingWebIdentityConfig", func(t *testing.T) {
		isolateAWSEnv(t)
		spec := &v1beta1.ProviderConfigSpec{
			Credentials: v1beta1.ProviderCredentials{Source: authKeyWebIdentity},
		}
		if _, _, err := UseWebIdentityToken(context.Background(), "us-east-1", spec, nil); err == nil {
			t.Error("UseWebIdentityToken(...): want an error when spec.credentials.webIdentity is nil, got none")
		}
	})
}

func TestWebIdentityTokenReadTimeout(t *testing.T) {
	const tokenKey = "token"

	// secretRetriever returns the retriever serving the token from a Secret
	// read through the supplied client, detached from the supplied context just
	// like a retriever handed to the AWS SDK is.
	secretRetriever := func(t *testing.T, ctx context.Context, kube client.Client) *xpWebIdentityTokenRetriever {
		t.Helper()
		isolateAWSEnv(t)
		r, err := webIdentityTokenRetriever(ctx, webIdentitySpec(&v1beta1.WebIdentityTokenConfig{
			Source:    xpv2.CredentialsSourceSecret,
			SecretRef: &xpv2.SecretKeySelector{Key: tokenKey},
		}).Credentials.WebIdentity, kube)
		if err != nil {
			t.Fatalf("webIdentityTokenRetriever(...): unexpected error: %v", err)
		}
		xr, ok := r.(*xpWebIdentityTokenRetriever)
		if !ok {
			t.Fatalf("webIdentityTokenRetriever(...): want a *xpWebIdentityTokenRetriever, got %T", r)
		}
		return xr
	}

	// blockingClient never serves the token, it only returns once the context
	// it is given is done, like a client whose API server hangs.
	blockingClient := &test.MockClient{
		MockGet: func(ctx context.Context, _ client.ObjectKey, _ client.Object) error {
			<-ctx.Done()
			return ctx.Err()
		},
	}

	t.Run("DetachedReadTimesOut", func(t *testing.T) {
		r := secretRetriever(t, context.Background(), blockingClient)
		r.readTimeout = 50 * time.Millisecond

		done := make(chan error, 1)
		go func() {
			_, err := r.GetIdentityToken()
			done <- err
		}()
		select {
		case err := <-done:
			if !errors.Is(err, context.DeadlineExceeded) {
				t.Errorf("GetIdentityToken(): want an error wrapping context.DeadlineExceeded, got %v", err)
			}
		case <-time.After(30 * time.Second):
			t.Fatal("GetIdentityToken(): want the read to be bounded by the timeout, it blocked instead")
		}
	})

	t.Run("DefaultTimeoutWhenUnset", func(t *testing.T) {
		if got := (&xpWebIdentityTokenRetriever{}).readTimeoutOrDefault(); got != webIdentityTokenReadTimeout {
			t.Errorf("readTimeoutOrDefault(): want the default %v, got %v", webIdentityTokenReadTimeout, got)
		}
		if got := (&xpWebIdentityTokenRetriever{readTimeout: time.Second}).readTimeoutOrDefault(); got != time.Second {
			t.Errorf("readTimeoutOrDefault(): want the override %v, got %v", time.Second, got)
		}
	})

	// the caller supplying the context owns the read's bound, so the bound
	// retriever must not impose a deadline of its own on top of it.
	t.Run("DeadlineOnlyOnTheDetachedPath", func(t *testing.T) {
		const token = "a-web-identity-token"
		var sawDeadline bool
		recordingClient := &test.MockClient{
			MockGet: func(ctx context.Context, _ client.ObjectKey, obj client.Object) error {
				_, sawDeadline = ctx.Deadline()
				s, ok := obj.(*corev1.Secret)
				if !ok {
					return fmt.Errorf("unexpected object type %T", obj)
				}
				s.Data = map[string][]byte{tokenKey: []byte(token)}
				return nil
			},
		}
		r := secretRetriever(t, context.Background(), recordingClient)

		got, err := r.GetIdentityToken()
		if err != nil {
			t.Fatalf("GetIdentityToken(): unexpected error: %v", err)
		}
		if string(got) != token {
			t.Errorf("GetIdentityToken(): want %q, got %q", token, string(got))
		}
		if !sawDeadline {
			t.Error("GetIdentityToken(): want the detached read to carry a deadline, it carried none")
		}

		got, err = r.withContext(context.Background()).GetIdentityToken()
		if err != nil {
			t.Fatalf("withContext(...).GetIdentityToken(): unexpected error: %v", err)
		}
		if string(got) != token {
			t.Errorf("withContext(...).GetIdentityToken(): want %q, got %q", token, string(got))
		}
		if sawDeadline {
			t.Error("withContext(...).GetIdentityToken(): want the bound read to carry no deadline of its own, it carried one")
		}
	})

	t.Run("BoundReadHonorsItsContext", func(t *testing.T) {
		r := secretRetriever(t, context.Background(), blockingClient)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		done := make(chan error, 1)
		go func() {
			_, err := r.withContext(ctx).GetIdentityToken()
			done <- err
		}()
		select {
		case err := <-done:
			if !errors.Is(err, context.Canceled) {
				t.Errorf("withContext(...).GetIdentityToken(): want an error wrapping context.Canceled, got %v", err)
			}
		case <-time.After(30 * time.Second):
			t.Fatal("withContext(...).GetIdentityToken(): want the read to observe the cancellation, it blocked instead")
		}
	})
}

// make sure the retrievers we hand to the AWS SDK satisfy its interface.
var (
	_ stscreds.IdentityTokenRetriever = &xpWebIdentityTokenRetriever{}
	_ stscreds.IdentityTokenRetriever = boundIdentityTokenRetriever{}
	_ stscreds.IdentityTokenRetriever = stubTokenRetriever{}
)
