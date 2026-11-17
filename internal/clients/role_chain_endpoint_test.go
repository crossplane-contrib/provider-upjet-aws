// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/utils/ptr"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

// fakeSTS answers AssumeRole with credentials derived from the requested role
// ARN and records the role ARNs it was asked for, in order.
type fakeSTS struct {
	mu    sync.Mutex
	roles []string
}

func (f *fakeSTS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if r.Form.Get("Action") != "AssumeRole" {
		http.Error(w, "unexpected action "+r.Form.Get("Action"), http.StatusBadRequest)
		return
	}
	role := r.Form.Get("RoleArn")
	f.mu.Lock()
	f.roles = append(f.roles, role)
	f.mu.Unlock()
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(w, `<AssumeRoleResponse xmlns="https://sts.amazonaws.com/doc/2011-06-15/">
  <AssumeRoleResult>
    <Credentials>
      <AccessKeyId>%s</AccessKeyId>
      <SecretAccessKey>secret</SecretAccessKey>
      <SessionToken>token</SessionToken>
      <Expiration>2099-01-01T00:00:00Z</Expiration>
    </Credentials>
    <AssumedRoleUser>
      <Arn>%s/session</Arn>
      <AssumedRoleId>AROATEST:session</AssumedRoleId>
    </AssumedRoleUser>
  </AssumeRoleResult>
  <ResponseMetadata><RequestId>00000000-0000-0000-0000-000000000000</RequestId></ResponseMetadata>
</AssumeRoleResponse>`, akidFor(role), role)
}

func (f *fakeSTS) seen() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.roles...)
}

func akidFor(role string) string {
	return "ASIA" + strings.ToUpper(role[strings.LastIndex(role, "/")+1:])
}

// TestRoleChainUsesCustomEndpoint checks that the STS calls made for
// spec.assumeRoleChain go to spec.endpoint, for every link of the chain, like
// every other client built from the ProviderConfig does.
func TestRoleChainUsesCustomEndpoint(t *testing.T) {
	isolateAWSEnv(t)
	// Any request that escapes to a non-loopback host (e.g. the public STS
	// endpoint) fails fast through a dead proxy instead of leaving the machine.
	// Loopback requests, such as those to the fake STS below, are never proxied.
	t.Setenv("HTTPS_PROXY", "http://127.0.0.1:1")
	t.Setenv("HTTP_PROXY", "http://127.0.0.1:1")

	sts := &fakeSTS{}
	srv := httptest.NewServer(sts)
	t.Cleanup(srv.Close)

	first := "arn:aws:iam::111111111111:role/first"
	second := "arn:aws:iam::222222222222:role/second"
	pc := &v1beta1.ClusterProviderConfig{
		Spec: v1beta1.ProviderConfigSpec{
			Credentials: v1beta1.ProviderCredentials{
				Source: xpv2.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv2.CommonCredentialSelectors{
					SecretRef: &xpv2.SecretKeySelector{
						SecretReference: xpv2.SecretReference{Name: "creds", Namespace: "crossplane-system"},
						Key:             "credentials",
					},
				},
			},
			AssumeRoleChain: []v1beta1.AssumeRoleOptions{
				{RoleARN: ptr.To(first)},
				{RoleARN: ptr.To(second)},
			},
			Endpoint: &v1beta1.EndpointConfig{
				URL:               v1beta1.URLConfig{Type: URLConfigTypeStatic, Static: ptr.To(srv.URL)},
				HostnameImmutable: ptr.To(true),
			},
		},
	}
	creds := []byte("[default]\naws_access_key_id = AKIDBASE\naws_secret_access_key = base\n")
	mg := &unstructured.Unstructured{Object: map[string]any{
		"spec": map[string]any{"forProvider": map[string]any{"region": "us-east-1"}},
	}}

	cfg, err := GetAWSConfigWithoutTracking(context.Background(), tokenServingClient(creds, "credentials"), mg, pc)
	if err != nil {
		t.Fatalf("GetAWSConfigWithoutTracking(...): unexpected error: %v", err)
	}
	got, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("Credentials.Retrieve(...): unexpected error (the role chain did not use spec.endpoint): %v", err)
	}
	if want := akidFor(second); got.AccessKeyID != want {
		t.Errorf("Credentials.Retrieve(...): want access key %q from the last link of the chain, got %q", want, got.AccessKeyID)
	}
	if want, seen := []string{first, second}, sts.seen(); strings.Join(seen, ",") != strings.Join(want, ",") {
		t.Errorf("fake STS: want AssumeRole calls %v, got %v", want, seen)
	}
}
