// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/google/go-cmp/cmp"
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	namespacedv1beta1 "github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
	"k8s.io/utils/ptr"
)

func TestBuildNoForkAWSConfig(t *testing.T) {
	creds := aws.Credentials{
		AccessKeyID:     "access-key",
		SecretAccessKey: "secret-key",
		SessionToken:    "session-token",
	}
	region := "us-west-2"

	cases := map[string]struct {
		reason         string
		pc             *namespacedv1beta1.ClusterProviderConfig
		wantStaticCred bool
		wantAssumeRole []awsbase.AssumeRole
	}{
		"IRSADynamicWithoutRoleChain": {
			reason: "IRSA auth should keep credentials dynamic so they can be refreshed by Terraform provider.",
			pc: &namespacedv1beta1.ClusterProviderConfig{
				Spec: namespacedv1beta1.ProviderConfigSpec{
					Credentials: namespacedv1beta1.ProviderCredentials{
						Source: xpv1.CredentialsSource(authKeyIRSA),
					},
				},
			},
			wantStaticCred: false,
		},
		"PodIdentityDynamicWithRoleChain": {
			reason: "PodIdentity auth should also use dynamic credentials and pass role chain to Terraform config.",
			pc: &namespacedv1beta1.ClusterProviderConfig{
				Spec: namespacedv1beta1.ProviderConfigSpec{
					Credentials: namespacedv1beta1.ProviderCredentials{
						Source: xpv1.CredentialsSource(authKeyPodIdentity),
					},
					S3UsePathStyle:       true,
					SkipRegionValidation: true,
					AssumeRoleChain: []namespacedv1beta1.AssumeRoleOptions{
						{
							RoleARN:    ptr.To("arn:aws:iam::111111111111:role/target"),
							ExternalID: ptr.To("external-id"),
							Tags: []namespacedv1beta1.Tag{
								{
									Key:   ptr.To("team"),
									Value: ptr.To("platform"),
								},
							},
							TransitiveTagKeys: []string{"team"},
						},
					},
				},
			},
			wantStaticCred: false,
			wantAssumeRole: []awsbase.AssumeRole{
				{
					RoleARN:           "arn:aws:iam::111111111111:role/target",
					ExternalID:        "external-id",
					Tags:              map[string]string{"team": "platform"},
					TransitiveTagKeys: []string{"team"},
				},
			},
		},
		"SecretAuthUsesStaticCredentials": {
			reason: "Secret-based auth should preserve existing static credential behavior.",
			pc: &namespacedv1beta1.ClusterProviderConfig{
				Spec: namespacedv1beta1.ProviderConfigSpec{
					Credentials: namespacedv1beta1.ProviderCredentials{
						Source: xpv1.CredentialsSourceSecret,
					},
				},
			},
			wantStaticCred: true,
		},
		"NilProviderConfigFallsBackToStatic": {
			reason:         "A nil ProviderConfig should defensively fall back to static credentials.",
			pc:             nil,
			wantStaticCred: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := buildNoForkAWSConfig(region, creds, tc.pc)

			if got.Region != region {
				t.Fatalf("%s: got region %q, want %q", tc.reason, got.Region, region)
			}
			if !got.SkipCredsValidation {
				t.Fatalf("%s: SkipCredsValidation should be true", tc.reason)
			}
			if !got.SkipRequestingAccountId {
				t.Fatalf("%s: SkipRequestingAccountId should be true", tc.reason)
			}

			if tc.wantStaticCred {
				if got.AccessKey != creds.AccessKeyID {
					t.Fatalf("%s: got access key %q, want %q", tc.reason, got.AccessKey, creds.AccessKeyID)
				}
				if got.SecretKey != creds.SecretAccessKey {
					t.Fatalf("%s: got secret key %q, want %q", tc.reason, got.SecretKey, creds.SecretAccessKey)
				}
				if got.Token != creds.SessionToken {
					t.Fatalf("%s: got session token %q, want %q", tc.reason, got.Token, creds.SessionToken)
				}
			} else {
				if got.AccessKey != "" || got.SecretKey != "" || got.Token != "" {
					t.Fatalf("%s: expected dynamic credentials configuration without static key material", tc.reason)
				}
			}

			if diff := cmp.Diff(tc.wantAssumeRole, got.AssumeRole); diff != "" {
				t.Fatalf("%s: AssumeRole mismatch (-want,+got):\n%s", tc.reason, diff)
			}

			if tc.pc != nil {
				if got.S3UsePathStyle != tc.pc.Spec.S3UsePathStyle {
					t.Fatalf("%s: got S3UsePathStyle %t, want %t", tc.reason, got.S3UsePathStyle, tc.pc.Spec.S3UsePathStyle)
				}
				if got.SkipRegionValidation != tc.pc.Spec.SkipRegionValidation {
					t.Fatalf("%s: got SkipRegionValidation %t, want %t", tc.reason, got.SkipRegionValidation, tc.pc.Spec.SkipRegionValidation)
				}
			}
		})
	}
}
