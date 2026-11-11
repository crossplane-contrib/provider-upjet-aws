// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudfrontpolicy

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const testPolicyID = "658327ea-f89d-4fab-a63d-7e88639e58f6"

func TestReadWithPolicyNameLookup(t *testing.T) {
	tests := map[string]struct {
		id         string
		findID     string
		findErr    error
		wantID     string
		wantRead   bool
		wantFind   bool
		wantDiagOK bool
	}{
		"UUIDSkipsLookup": {
			id:         testPolicyID,
			wantID:     testPolicyID,
			wantRead:   true,
			wantDiagOK: true,
		},
		"NameResolvedToID": {
			id:         "Managed-CachingOptimized",
			findID:     testPolicyID,
			wantID:     testPolicyID,
			wantRead:   true,
			wantFind:   true,
			wantDiagOK: true,
		},
		"UnknownNameIsNotFound": {
			id:         "missing",
			wantID:     "",
			wantFind:   true,
			wantDiagOK: true,
		},
		"LookupError": {
			id:       "Managed-CachingOptimized",
			findErr:  errors.New("throttled"),
			wantID:   "Managed-CachingOptimized",
			wantFind: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var readCalled, findCalled bool
			read := func(context.Context, *schema.ResourceData, any) diag.Diagnostics {
				readCalled = true
				return nil
			}
			find := func(_ context.Context, _ any, name string) (string, error) {
				findCalled = true
				if name != tc.id {
					t.Errorf("find called with %q, want %q", name, tc.id)
				}
				return tc.findID, tc.findErr
			}
			d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{}, map[string]any{})
			d.SetId(tc.id)

			diags := readWithPolicyNameLookup(read, find)(context.Background(), d, nil)

			if got := !diags.HasError(); got != tc.wantDiagOK {
				t.Errorf("diags OK = %v, want %v: %v", got, tc.wantDiagOK, diags)
			}
			if d.Id() != tc.wantID {
				t.Errorf("ID = %q, want %q", d.Id(), tc.wantID)
			}
			if readCalled != tc.wantRead {
				t.Errorf("read called = %v, want %v", readCalled, tc.wantRead)
			}
			if findCalled != tc.wantFind {
				t.Errorf("find called = %v, want %v", findCalled, tc.wantFind)
			}
		})
	}
}

func TestFindPolicyID(t *testing.T) {
	type policy struct{ name, id string }
	pages := map[string]struct {
		items []policy
		next  *string
	}{
		"":      {items: []policy{{"a", "id-a"}}, next: aws.String("page2")},
		"page2": {items: []policy{{"b", "id-b"}}},
	}
	list := func(marker *string) ([]policy, *string, error) {
		m := ""
		if marker != nil {
			m = *marker
		}
		p := pages[m]
		return p.items, p.next, nil
	}
	nameAndID := func(p policy) (string, string) { return p.name, p.id }

	tests := map[string]struct {
		name string
		want string
	}{
		"FirstPage":  {name: "a", want: "id-a"},
		"SecondPage": {name: "b", want: "id-b"},
		"NotFound":   {name: "c", want: ""},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := findPolicyID(tc.name, list, nameAndID)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("findPolicyID(%q) = %q, want %q", tc.name, got, tc.want)
			}
		})
	}

	t.Run("ListError", func(t *testing.T) {
		_, err := findPolicyID("a", func(*string) ([]policy, *string, error) {
			return nil, nil, errors.New("access denied")
		}, nameAndID)
		if err == nil {
			t.Fatal("expected an error")
		}
	})
}
