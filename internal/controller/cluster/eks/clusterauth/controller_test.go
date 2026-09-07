// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clusterauth

import (
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/google/go-cmp/cmp"
)

func TestSupportedManagementPolicies(t *testing.T) {
	type args struct {
		policies xpv2.ManagementPolicies
	}
	type want struct {
		supported bool
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"Wildcard": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionAll}},
			want: want{supported: true},
		},
		"Paused": {
			args: args{policies: xpv2.ManagementPolicies{}},
			want: want{supported: true},
		},
		"ObserveCreateUpdate": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve, xpv2.ManagementActionCreate, xpv2.ManagementActionUpdate}},
			want: want{supported: true},
		},
		"ObserveCreateUpdateLateInitialize": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve, xpv2.ManagementActionCreate, xpv2.ManagementActionUpdate, xpv2.ManagementActionLateInitialize}},
			want: want{supported: true},
		},
		"ObserveCreateUpdateDelete": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve, xpv2.ManagementActionCreate, xpv2.ManagementActionUpdate, xpv2.ManagementActionDelete}},
			want: want{supported: true},
		},
		"AllActionsExplicit": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve, xpv2.ManagementActionCreate, xpv2.ManagementActionUpdate, xpv2.ManagementActionLateInitialize, xpv2.ManagementActionDelete}},
			want: want{supported: true},
		},
		"ObserveOnlyCannotMintToken": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve}},
			want: want{supported: false},
		},
		"ObserveCreateCannotRefreshToken": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve, xpv2.ManagementActionCreate}},
			want: want{supported: false},
		},
		"ObserveCreateLateInitializeCannotRefreshToken": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve, xpv2.ManagementActionCreate, xpv2.ManagementActionLateInitialize}},
			want: want{supported: false},
		},
		"ObserveCreateDeleteCannotRefreshToken": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve, xpv2.ManagementActionCreate, xpv2.ManagementActionDelete}},
			want: want{supported: false},
		},
		"ObserveUpdateCannotMintToken": {
			args: args{policies: xpv2.ManagementPolicies{xpv2.ManagementActionObserve, xpv2.ManagementActionUpdate}},
			want: want{supported: false},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			resolver := managed.NewManagementPoliciesResolver(true, tc.args.policies, managed.WithSupportedManagementPolicies(supportedManagementPolicies()))
			got := resolver.Validate() == nil
			if diff := cmp.Diff(tc.want.supported, got); diff != "" {
				t.Errorf("Validate() supported: -want, +got:\n%s", diff)
			}
		})
	}
}
