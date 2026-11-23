// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	tjresource "github.com/crossplane/upjet/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"k8s.io/utils/ptr"

	clusters3v1beta1 "github.com/upbound/provider-aws/v2/apis/cluster/s3/v1beta1"
	clusters3 "github.com/upbound/provider-aws/v2/apis/cluster/s3/v1beta2"
	namespaceds3 "github.com/upbound/provider-aws/v2/apis/namespaced/s3/v1beta1"
	pcv1beta1 "github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

func TestBucketTaggingAPI(t *testing.T) {
	cases := map[string]struct {
		mg      resource.Managed
		want    string
		wantErr bool
	}{
		"Absent":           {mg: &clusters3.Bucket{}, want: "Auto"},
		"Auto":             {mg: &clusters3.Bucket{Spec: clusters3.BucketSpec{ForProvider: clusters3.BucketParameters{TaggingAPI: ptr.To("Auto")}}}, want: "Auto"},
		"ClusterBucket":    {mg: &clusters3.Bucket{Spec: clusters3.BucketSpec{ForProvider: clusters3.BucketParameters{TaggingAPI: ptr.To("S3")}}}, want: "S3"},
		"ClusterBucketV1beta1": {mg: &clusters3v1beta1.Bucket{Spec: clusters3v1beta1.BucketSpec{ForProvider: clusters3v1beta1.BucketParameters{TaggingAPI: ptr.To("S3")}}}, want: "S3"},
		"NamespacedBucket": {mg: &namespaceds3.Bucket{Spec: namespaceds3.BucketSpec{ForProvider: namespaceds3.BucketParameters{TaggingAPI: ptr.To("S3")}}}, want: "S3"},
		"DirectoryBucket":  {mg: &namespaceds3.DirectoryBucket{}, want: "Auto"},
		"BucketACL":        {mg: &clusters3.BucketACL{}, want: "Auto"},
		"Empty":            {mg: &clusters3.Bucket{Spec: clusters3.BucketSpec{ForProvider: clusters3.BucketParameters{TaggingAPI: ptr.To("")}}}, wantErr: true},
		"Unknown":          {mg: &clusters3.Bucket{Spec: clusters3.BucketSpec{ForProvider: clusters3.BucketParameters{TaggingAPI: ptr.To("S3Control")}}}, wantErr: true},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := bucketTaggingAPI(tc.mg)
			if (err != nil) != tc.wantErr || got != tc.want {
				t.Fatalf("bucketTaggingAPI() = %q, %v; want %q, error=%v", got, err, tc.want, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			tr := tc.mg.(tjresource.Terraformed)
			params, err := tr.GetParameters()
			if err != nil {
				t.Fatal(err)
			}
			if _, ok := params["tagging_api"]; ok {
				t.Fatal("synthetic option was serialized to Terraform")
			}
			if _, err := tr.LateInitialize([]byte(`{"tagging_api":"S3"}`)); err != nil {
				t.Fatal(err)
			}
			after, err := bucketTaggingAPI(tc.mg)
			if err != nil || after != got {
				t.Fatalf("late initialization changed synthetic option: %v, %v", after, err)
			}
		})
	}
}

// Exercise the embedded provider's tagging implementation, including writes,
// with S3 Control pointed at an endpoint that rejects every request.
func TestS3TaggingAPIClientIsolation(t *testing.T) {
	isolateAWSEnv(t)
	ctx := context.Background()
	_, provider, err := xpprovider.GetProvider(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var requests []string
	var controlRequests int
	s3Endpoint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v20180820/") {
			controlRequests++
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		requests = append(requests, r.Method+" "+r.URL.Path+"?"+r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/xml")
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`<Tagging xmlns="http://s3.amazonaws.com/doc/2006-03-01/"><TagSet><Tag><Key>environment</Key><Value>test</Value></Tag></TagSet></Tagging>`))
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}))
	defer s3Endpoint.Close()
	pc := &pcv1beta1.ClusterProviderConfig{Spec: pcv1beta1.ProviderConfigSpec{
		SkipReqAccountId:     true,
		SkipMetadataApiCheck: true,
		S3UsePathStyle:       true,
		Endpoint:             &pcv1beta1.EndpointConfig{Services: []string{"s3", "s3control"}, URL: pcv1beta1.URLConfig{Static: ptr.To(s3Endpoint.URL)}},
	}}
	config := &SetupConfig{TerraformProvider: provider}
	newClient := func(taggingAPI string) *xpprovider.AWSClient {
		t.Helper()
		ps := terraform.Setup{ClientMetadata: map[string]string{keyAccountID: "123456789012"}}
		if err := configureNoForkAWSClient(ctx, &ps, config, "us-east-1", aws.Credentials{AccessKeyID: "test", SecretAccessKey: "test"}, pc, taggingAPI); err != nil {
			t.Fatal(err)
		}
		if ps.ClientMetadata[keyAccountID] != "123456789012" {
			t.Fatal("client setup modified identity metadata")
		}
		return ps.Meta.(*xpprovider.AWSClient)
	}
	normal := newClient(bucketTaggingAPIAuto)
	compatible := newClient(bucketTaggingAPIS3)
	after := newClient(bucketTaggingAPIAuto)
	if normal.AccountID(ctx) != "123456789012" || compatible.AccountID(ctx) != "" || after.AccountID(ctx) != "123456789012" {
		t.Fatal("bucket option leaked across clients or ProviderConfig flag changed default behavior")
	}
	tagger, ok := compatible.ServicePackage(ctx, "s3").(interface {
		ListTags(context.Context, any, string, string) error
		UpdateTags(context.Context, any, string, string, any, any) error
	})
	if !ok {
		t.Fatal("S3 service package does not expose tagging")
	}
	// S3 Control prefixes the endpoint hostname with the account ID, which
	// cannot resolve for this fixture. The default client must still attempt
	// ListTagsForResource and return that failure rather than use S3 tagging.
	if err := tagger.ListTags(ctx, normal, "test-bucket", "Bucket"); err == nil || !strings.Contains(err.Error(), "ListTagsForResource") {
		t.Fatalf("default client did not attempt S3 Control: %v", err)
	}
	controlRequests = 0
	if err := tagger.ListTags(ctx, compatible, "test-bucket", "Bucket"); err != nil {
		t.Fatal(err)
	}
	if err := tagger.UpdateTags(ctx, compatible, "test-bucket", "Bucket", map[string]string{}, map[string]string{"environment": "test"}); err != nil {
		t.Fatal(err)
	}
	if err := tagger.UpdateTags(ctx, compatible, "test-bucket", "Bucket", map[string]string{"environment": "test"}, map[string]string{}); err != nil {
		t.Fatal(err)
	}
	// Tag updates first read existing tags to preserve externally managed tags.
	want := []string{"GET /test-bucket?tagging=", "GET /test-bucket?tagging=", "PUT /test-bucket?tagging=", "GET /test-bucket?tagging=", "DELETE /test-bucket?tagging="}
	if !reflect.DeepEqual(requests, want) {
		t.Errorf("tagging requests = %v, want %v", requests, want)
	}
	if controlRequests != 0 {
		t.Errorf("made %d S3 Control requests", controlRequests)
	}
}
