// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	"github.com/crossplane/upjet/v2/pkg/diffserver"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"
	"k8s.io/utils/ptr"

	ec2v1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/ec2/v1beta1"
	iamv1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/iam/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

// The offline setup exists to make the diff server independent of the host it
// runs on. Two things carry that, and both fail silently when they break: the
// credentials have to be static so that the AWS SDK never consults its
// discovery chain, and the egress guard has to refuse every request so that a
// resource which calls AWS from its CustomizeDiff function fails in a way the
// diff server can report. A host that happens to have credentials hides the
// first, and a network that happens to answer hides the second.

func TestOfflineCredentialsAreStatic(t *testing.T) {
	// aws-sdk-go-base only falls back to environment variables, credentials
	// files and IMDS when no static key is configured. An empty key here would
	// make the diff server compute plans against whatever account the host is
	// logged in to, and fail outright on a host with no credentials at all.
	if offlineAccessKeyID == "" || offlineSecretAccessKey == "" {
		t.Error("the offline placeholder credentials must not be empty, or the AWS SDK falls back to its credential discovery chain and the setup stops being offline")
	}
}

func TestDenyOutboundRequests(t *testing.T) {
	ctx := awsmiddleware.SetServiceID(context.Background(), "SFN")
	ctx = awsmiddleware.SetOperationName(ctx, "ValidateStateMachineDefinition")

	reached := false
	handler := func(stack *middleware.Stack) (any, error) {
		reached = false
		h := middleware.DecorateHandler(middleware.HandlerFunc(
			func(context.Context, any) (any, middleware.Metadata, error) {
				reached = true
				return nil, middleware.Metadata{}, nil
			}), stack)
		_, _, err := h.Handle(ctx, nil)
		return nil, err
	}

	// Without the guard the request reaches the transport. This is here so
	// that the assertions below cannot pass because the harness never sent
	// anything in the first place.
	if _, err := handler(middleware.NewStack("control", func() any { return nil })); err != nil {
		t.Fatalf("the control stack returned an error before the guard was involved: %v", err)
	}
	if !reached {
		t.Fatal("the control stack did not reach the transport, so this test cannot show that the guard is what blocks it")
	}

	stack := middleware.NewStack("test", func() any { return nil })
	if err := denyOutboundRequests(stack); err != nil {
		t.Fatalf("denyOutboundRequests(): cannot install the guard: %v", err)
	}
	_, err := handler(stack)

	if err == nil {
		t.Fatal("denyOutboundRequests(): want the request refused, got no error")
	}
	if reached {
		t.Error("denyOutboundRequests(): the request reached the transport, so nothing was blocked")
	}
	// The diff server turns this error into a FailedPrecondition carrying
	// DIFF_COMPUTATION_NOT_SUPPORTED, which is what tells a client to fall
	// back to a raw diff. An unrecognised error would surface as an Internal
	// one instead, and the client would have nothing to act on.
	if !diffserver.IsDiffComputationNotSupportedError(err) {
		t.Errorf("denyOutboundRequests(): want an error the diff server reports as an unsupported diff, got %v", err)
	}
	for _, want := range []string{"SFN", "ValidateStateMachineDefinition"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("denyOutboundRequests(): want the message to name the refused call %q, got %q", want, err.Error())
		}
	}
}

func TestOfflineRegion(t *testing.T) {
	regional := &ec2v1beta1.VPC{}
	regional.Spec.ForProvider.Region = ptr.To("eu-central-1")
	regional.SetGroupVersionKind(ec2v1beta1.VPC_GroupVersionKind)

	// IAM resources declare no region at all, but the Terraform provider will
	// not configure itself without one.
	global := &iamv1beta1.Role{}
	global.SetGroupVersionKind(iamv1beta1.Role_GroupVersionKind)

	t.Run("FromTheResource", func(t *testing.T) {
		got, err := offlineRegion(regional, &namespacedv1beta1.ClusterProviderConfig{})
		if err != nil {
			t.Fatalf("offlineRegion(): unexpected error: %v", err)
		}
		if diff := cmp.Diff("eu-central-1", got); diff != "" {
			t.Errorf("offlineRegion(): -want, +got:\n%s", diff)
		}
	})

	t.Run("GlobalResourceFallsBack", func(t *testing.T) {
		got, err := offlineRegion(global, &namespacedv1beta1.ClusterProviderConfig{})
		if err != nil {
			t.Fatalf("offlineRegion(): unexpected error: %v", err)
		}
		if got == "" {
			t.Error("offlineRegion(): want a non-empty region for a global resource, because the Terraform provider will not configure itself without one")
		}
	})
}

func TestSetPartitionOffline(t *testing.T) {
	// Resolving the partition is pure computation over the region and the
	// ProviderConfig, so the offline setup does it exactly as the online one
	// does. External name templates put it into Terraform IDs.
	cases := map[string]struct {
		region string
		want   string
	}{
		"Commercial": {region: "us-west-1", want: "aws"},
		"GovCloud":   {region: "us-gov-west-1", want: "aws-us-gov"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ps := terraform.Setup{ClientMetadata: map[string]string{}}
			if err := setPartition(&aws.Config{Region: tc.region}, &namespacedv1beta1.ClusterProviderConfig{}, &ps); err != nil {
				t.Fatalf("setPartition(): unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, ps.ClientMetadata[keyPartition]); diff != "" {
				t.Errorf("setPartition(): -want partition, +got:\n%s", diff)
			}
		})
	}
}
