// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package efs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	fileSystemID       = "fs-0123456789abcdef0"
	fileSystemARN      = `"arn:aws:elasticfilesystem:eu-central-1:123456789012:file-system/fs-0123456789abcdef0"`
	otherFileSystemARN = `"arn:aws:elasticfilesystem:eu-central-1:123456789012:file-system/fs-0fedcba9876543210"`
)

// statement returns a statement that denies access without TLS, with the
// given JSON as its Resource, or without a Resource when it is empty.
func statement(effect, resource string) string {
	r := ""
	if resource != "" {
		r = `"Resource":` + resource + `,`
	}
	return `{"Effect":"` + effect + `","Principal":{"AWS":"*"},"Action":"*",` + r + `"Condition":{"Bool":{"aws:SecureTransport":"false"}}}`
}

func policy(statements ...string) string {
	s := ""
	for i, st := range statements {
		if i > 0 {
			s += ","
		}
		s += st
	}
	return `{"Version":"2012-10-17","Statement":[` + s + `]}`
}

func TestFileSystemPolicyCustomDiff(t *testing.T) {
	bypass := &terraform.ResourceAttrDiff{Old: "false", New: "true"}

	type args struct {
		attributes map[string]*terraform.ResourceAttrDiff
		state      *terraform.InstanceState
	}
	type want struct {
		attributes map[string]*terraform.ResourceAttrDiff
		err        bool
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"ResourceAddedByEFS": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Deny", ""))},
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{},
			},
		},
		"ResourceAddedByEFSWithOtherChange": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy":                             {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Deny", ""))},
					"bypass_policy_lockout_safety_check": bypass,
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"bypass_policy_lockout_safety_check": bypass,
				},
			},
		},
		"ResourceAddedByEFSToSingleStatement": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {
						Old: `{"Version":"2012-10-17","Statement":` + statement("Deny", fileSystemARN) + `}`,
						New: `{"Version":"2012-10-17","Statement":` + statement("Deny", "") + `}`,
					},
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{},
			},
		},
		"ResourceAddedByEFSToOneOfTwoStatements": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {
						Old: policy(statement("Deny", fileSystemARN), statement("Allow", otherFileSystemARN)),
						New: policy(statement("Deny", ""), statement("Allow", otherFileSystemARN)),
					},
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{},
			},
		},
		"ExplicitResourceAsList": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Deny", "["+fileSystemARN+"]"))},
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{},
			},
		},
		"DifferentResource": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Deny", otherFileSystemARN))},
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Deny", otherFileSystemARN))},
				},
			},
		},
		"ChangedEffect": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Allow", ""))},
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Allow", ""))},
				},
			},
		},
		"AddedStatement": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {
						Old: policy(statement("Deny", fileSystemARN)),
						New: policy(statement("Deny", ""), statement("Allow", otherFileSystemARN)),
					},
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {
						Old: policy(statement("Deny", fileSystemARN)),
						New: policy(statement("Deny", ""), statement("Allow", otherFileSystemARN)),
					},
				},
			},
		},
		"NoPolicyChange": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"bypass_policy_lockout_safety_check": bypass,
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"bypass_policy_lockout_safety_check": bypass,
				},
			},
		},
		"Create": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {New: policy(statement("Deny", ""))},
				},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {New: policy(statement("Deny", ""))},
				},
			},
		},
		"NoState": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Deny", ""))},
				},
				state: &terraform.InstanceState{},
			},
			want: want{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: policy(statement("Deny", ""))},
				},
			},
		},
		"InvalidDesiredPolicy": {
			args: args{
				attributes: map[string]*terraform.ResourceAttrDiff{
					"policy": {Old: policy(statement("Deny", fileSystemARN)), New: "{"},
				},
			},
			want: want{
				err: true,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			state := tc.args.state
			if state == nil {
				state = &terraform.InstanceState{Attributes: map[string]string{"file_system_id": fileSystemID}}
			}
			diff, err := fileSystemPolicyCustomDiff(&terraform.InstanceDiff{Attributes: tc.args.attributes}, state, nil)
			if diff := cmp.Diff(tc.want.err, err != nil); diff != "" {
				t.Fatalf("fileSystemPolicyCustomDiff(...): -want error, +got error:\n%s\nerror: %v", diff, err)
			}
			if err != nil {
				return
			}
			if diff := cmp.Diff(tc.want.attributes, diff.Attributes); diff != "" {
				t.Errorf("fileSystemPolicyCustomDiff(...): -want attributes, +got attributes:\n%s", diff)
			}
		})
	}
}
