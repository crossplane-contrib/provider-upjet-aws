// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"

	namespacedv1beta1 "github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

// TestConfigureNoForkAWSClientCredentialsProvider asserts that the Terraform
// AWS client uses the supplied refreshing credentials provider, see
// https://github.com/crossplane-contrib/provider-upjet-aws/issues/2292.
func TestConfigureNoForkAWSClientCredentialsProvider(t *testing.T) {
	snapshot := aws.Credentials{AccessKeyID: "snapshot", SecretAccessKey: "secret", SessionToken: "token"}
	refreshed := aws.Credentials{AccessKeyID: "refreshed", SecretAccessKey: "secret", SessionToken: "token"}

	cases := map[string]struct {
		reason        string
		credsProvider func() aws.CredentialsProvider
		wantAccessKey string
	}{
		"RefreshingProvider": {
			reason:        "The client must retrieve its credentials from the supplied provider, so that they are refreshed once the snapshot expires.",
			credsProvider: func() aws.CredentialsProvider { c, _ := newCountingCredsCache(refreshed); return c },
			wantAccessKey: "refreshed",
		},
		"NoProvider": {
			reason:        "Without a provider, the client must fall back to the supplied static credentials.",
			credsProvider: func() aws.CredentialsProvider { return nil },
			wantAccessKey: "snapshot",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			tfProvider := &schema.Provider{}
			tfProvider.SetMeta(&xpprovider.AWSClient{})
			ps := &terraform.Setup{ClientMetadata: map[string]string{keyAccountID: "123456789012"}}
			pc := &namespacedv1beta1.ClusterProviderConfig{}
			credsProvider := tc.credsProvider()

			if err := configureNoForkAWSClient(ctx, ps, &SetupConfig{TerraformProvider: tfProvider}, "us-west-2", snapshot, credsProvider, pc); err != nil {
				t.Fatalf("\n%s\nconfigureNoForkAWSClient(...): unexpected error: %v", tc.reason, err)
			}
			client, ok := ps.Meta.(*xpprovider.AWSClient)
			if !ok {
				t.Fatalf("\n%s\nconfigureNoForkAWSClient(...): want the Terraform setup's meta to be an *xpprovider.AWSClient, got %T", tc.reason, ps.Meta)
			}
			if credsProvider != nil && client.CredentialsProvider(ctx) != credsProvider {
				t.Errorf("\n%s\nconfigureNoForkAWSClient(...): the Terraform AWS client does not use the supplied credentials provider", tc.reason)
			}
			creds, err := client.AwsConfig(ctx).Credentials.Retrieve(ctx)
			if err != nil {
				t.Fatalf("\n%s\nRetrieve(...): unexpected error: %v", tc.reason, err)
			}
			if creds.AccessKeyID != tc.wantAccessKey {
				t.Errorf("\n%s\nRetrieve(...): want access key %q, got %q", tc.reason, tc.wantAccessKey, creds.AccessKeyID)
			}
		})
	}
}
