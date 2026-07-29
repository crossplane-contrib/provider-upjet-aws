// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package common

import (
	"context"
	"testing"
	"time"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
	"github.com/crossplane/crossplane-runtime/v2/pkg/test"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ujfake "github.com/crossplane/upjet/v2/pkg/resource/fake"
)

var (
	errBoom = errors.New("boom")
)

func TestPasswordGenerator(t *testing.T) {
	type args struct {
		kube               client.Client
		secretRefFieldPath string
		toggleFieldPath    string
		mg                 resource.Managed
	}
	type want struct {
		err error
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"CannotGetSecret": {
			reason: "An error should be returned if the referenced secret cannot be retrieved.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(errBoom),
				},
				secretRefFieldPath: "",
				toggleFieldPath:    "",
				mg: &fake.Managed{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "foo-mgd",
						Namespace: "bar",
					},
				},
			},
			want: want{
				err: errors.Wrap(errBoom, ErrGetPasswordSecret),
			},
		},
		"ClusterScopedMR": {
			reason: "should return an error if the MR has no namespace (cluster-scoped)",
			args: args{
				kube: &test.MockClient{
					MockGet: func(ctx context.Context, key client.ObjectKey, obj client.Object) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						s.Data = map[string][]byte{
							"password": []byte("foo"),
						}
						return nil
					},
				},
				secretRefFieldPath: "parameterizable.parameters.passwordSecretRef",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name: "foo-mgd",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
						},
					},
				},
			},
			want: want{
				err: errors.New(errManagedNotNamespaced),
			},
		},
		"SecretAlreadyFull": {
			reason: "Should be no-op if the Secret already has password.",
			args: args{
				kube: &test.MockClient{
					MockGet: func(ctx context.Context, key client.ObjectKey, obj client.Object) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						s.Data = map[string][]byte{
							"password": []byte("foo"),
						}
						return nil
					},
				},
				secretRefFieldPath: "parameterizable.parameters.passwordSecretRef",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
						},
					},
				},
			},
		},
		"ClusterSecretAlreadyFull": {
			reason: "Should be no-op if the Secret already has password.",
			args: args{
				kube: &test.MockClient{
					MockGet: func(ctx context.Context, key client.ObjectKey, obj client.Object) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						s.Data = map[string][]byte{
							"password": []byte("foo"),
						}
						return nil
					},
				},
				secretRefFieldPath: "parameterizable.parameters.masterPasswordSecretRef",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"masterPasswordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
						},
					},
				},
			},
		},
		"NoSecretReference": {
			reason: "Should be no-op if the secret reference is not given.",
			args: args{
				secretRefFieldPath: "parameterizable.parameters.passwordSecretRef",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"another": "field",
						},
					},
				},
			},
		},
		"NoClusterSecretReference": {
			reason: "Should be no-op if the secret reference is not given.",
			args: args{
				secretRefFieldPath: "parameterizable.parameters.masterPasswordSecretRef",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"another": "field",
						},
					},
				},
			},
		},
		"ToggleNotSet": {
			reason: "Should be no-op if the toggle is not set at all.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(nil),
				},
				secretRefFieldPath: "parameterizable.parameters.passwordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
						},
					},
				},
			},
		},
		"ClusterToggleNotSet": {
			reason: "Should be no-op if the toggle is not set at all.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(nil),
				},
				secretRefFieldPath: "parameterizable.parameters.masterPasswordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"masterPasswordSecretRef": map[string]any{
								"name":      "foo",
								"namespace": "bar",
								"key":       "password",
							},
						},
					},
				},
			},
		},
		"ToggleFalse": {
			reason: "Should be no-op if the toggle is set to false.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(nil),
				},
				secretRefFieldPath: "parameterizable.parameters.passwordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
							"autoGeneratePassword": false,
						},
					},
				},
			},
		},
		"ClusterToggleFalse": {
			reason: "Should be no-op if the toggle is set to false.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(nil),
				},
				secretRefFieldPath: "parameterizable.parameters.masterPasswordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"masterPasswordSecretRef": map[string]any{
								"name":      "foo",
								"namespace": "bar",
								"key":       "password",
							},
							"autoGeneratePassword": false,
						},
					},
				},
			},
		},
		"GenerateAndApply": {
			reason: "Should apply if we generate, set the content of an already existing secret.",
			args: args{
				kube: &test.MockClient{
					MockGet: func(ctx context.Context, key client.ObjectKey, obj client.Object) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						s.CreationTimestamp = metav1.Time{Time: time.Now()}
						return nil
					},
					MockPatch: func(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						if len(s.Data["password"]) == 0 {
							return errors.New("password is not set")
						}
						if len(s.OwnerReferences) != 0 {
							return errors.New("owner references should not be set if secret already exists")
						}
						return nil
					},
				},
				secretRefFieldPath: "parameterizable.parameters.passwordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
							"autoGeneratePassword": true,
						},
					},
				},
			},
		},
		"ClusterSecretGenerateAndApply": {
			reason: "Should apply if we generate, set the content of an already existing secret.",
			args: args{
				kube: &test.MockClient{
					MockGet: func(ctx context.Context, key client.ObjectKey, obj client.Object) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						s.CreationTimestamp = metav1.Time{Time: time.Now()}
						return nil
					},
					MockPatch: func(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						if len(s.Data["password"]) == 0 {
							return errors.New("password is not set")
						}
						if len(s.OwnerReferences) != 0 {
							return errors.New("owner references should not be set if secret already exists")
						}
						return nil
					},
				},
				secretRefFieldPath: "parameterizable.parameters.masterPasswordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"masterPasswordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
							"autoGeneratePassword": true,
						},
					},
				},
			},
		},
		"GenerateAndCreate": {
			reason: "Should create if we generate, set the content and there is no secret in place.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(kerrors.NewNotFound(schema.GroupResource{}, "")),
					MockCreate: func(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						if len(s.Data["password"]) == 0 {
							return errors.New("password is not set")
						}
						if len(s.OwnerReferences) == 1 &&
							s.OwnerReferences[0].Name == "foo-mgd" {
							return nil
						}
						return errors.New("owner references should be set if secret is created")
					},
				},
				secretRefFieldPath: "parameterizable.parameters.passwordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
							"autoGeneratePassword": true,
						},
					},
				},
			},
		},
		"ClusterSecretGenerateAndCreate": {
			reason: "Should create if we generate, set the content and there is no secret in place.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(kerrors.NewNotFound(schema.GroupResource{}, "")),
					MockCreate: func(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
						s, ok := obj.(*corev1.Secret)
						if !ok {
							return errors.New("needs to be secret")
						}
						if len(s.Data["password"]) == 0 {
							return errors.New("password is not set")
						}
						if len(s.OwnerReferences) == 1 &&
							s.OwnerReferences[0].Name == "foo-mgd" {
							return nil
						}
						return errors.New("owner references should be set if secret is created")
					},
				},
				secretRefFieldPath: "parameterizable.parameters.masterPasswordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
					Managed: fake.Managed{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo-mgd",
							Namespace: "bar",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"masterPasswordSecretRef": map[string]any{
								"name": "foo",
								"key":  "password",
							},
							"autoGeneratePassword": true,
						},
					},
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := PasswordGenerator(tc.args.secretRefFieldPath, tc.args.toggleFieldPath)(tc.args.kube).Initialize(context.Background(), tc.args.mg)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("PasswordGenerator(...): -want error, +got error:\n%s", diff)
			}
		})
	}

}

func TestJSONStringNormalizationConversion(t *testing.T) {
	type args struct {
		fieldPath string
		params    map[string]any
		mode      config.Mode
	}
	type want struct {
		params map[string]any
		err    error
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"NoOpOnToTerraform": {
			reason: "Should not modify params when mode is ToTerraform.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"policy": `{"Statement":[{"Effect":"Allow"}]}`},
				mode:      config.ToTerraform,
			},
			want: want{
				params: map[string]any{"policy": `{"Statement":[{"Effect":"Allow"}]}`},
			},
		},
		"NoOpOnMissingField": {
			reason: "Should not modify params when the field is not present.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"other": "value"},
				mode:      config.FromTerraform,
			},
			want: want{
				params: map[string]any{"other": "value"},
			},
		},
		"NoOpOnEmptyString": {
			reason: "Should not modify params when the field is an empty string.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"policy": ""},
				mode:      config.FromTerraform,
			},
			want: want{
				params: map[string]any{"policy": ""},
			},
		},
		"NoOpOnInvalidJSON": {
			reason: "Should not modify params when the field contains invalid JSON.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"policy": "not-json"},
				mode:      config.FromTerraform,
			},
			want: want{
				params: map[string]any{"policy": "not-json"},
			},
		},
		"NoOpOnNonStringField": {
			reason: "Should not modify params when the field is not a string.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"policy": 42},
				mode:      config.FromTerraform,
			},
			want: want{
				params: map[string]any{"policy": 42},
			},
		},
		"SortsTopLevelArray": {
			reason: "Should sort top-level arrays in the JSON string.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"policy": `{"Principal":{"AWS":["arn:aws:iam::use2","arn:aws:iam::euw1"]}}`},
				mode:      config.FromTerraform,
			},
			want: want{
				params: map[string]any{"policy": `{"Principal":{"AWS":["arn:aws:iam::euw1","arn:aws:iam::use2"]}}`},
			},
		},
		"AlreadySorted": {
			reason: "Should produce identical output when arrays are already sorted.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"policy": `{"Principal":{"AWS":["arn:aws:iam::euw1","arn:aws:iam::use2"]}}`},
				mode:      config.FromTerraform,
			},
			want: want{
				params: map[string]any{"policy": `{"Principal":{"AWS":["arn:aws:iam::euw1","arn:aws:iam::use2"]}}`},
			},
		},
		"SortsNestedArrays": {
			reason: "Should recursively sort arrays nested inside objects.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"policy": `{"Statement":[{"Effect":"Allow","Principal":{"AWS":["arn:b","arn:a"]}},{"Effect":"Deny","Principal":{"AWS":["arn:d","arn:c"]}}]}`},
				mode:      config.FromTerraform,
			},
			want: want{
				params: map[string]any{"policy": `{"Statement":[{"Effect":"Allow","Principal":{"AWS":["arn:a","arn:b"]}},{"Effect":"Deny","Principal":{"AWS":["arn:c","arn:d"]}}]}`},
			},
		},
		"DeterministicAcrossOrderings": {
			reason: "Two policies differing only in array order should produce the same normalized output.",
			args: args{
				fieldPath: "policy",
				params:    map[string]any{"policy": `{"AWS":["arn:aws:iam::use2","arn:aws:iam::euw1","arn:aws:iam::apse1"]}`},
				mode:      config.FromTerraform,
			},
			want: want{
				params: map[string]any{"policy": `{"AWS":["arn:aws:iam::apse1","arn:aws:iam::euw1","arn:aws:iam::use2"]}`},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewJSONStringNormalizationConversion(tc.args.fieldPath)
			got, err := c.Convert(tc.args.params, nil, tc.args.mode)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("%s\nConvert(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.params, got); diff != "" {
				t.Errorf("%s\nConvert(...): -want params, +got params:\n%s", tc.reason, diff)
			}
		})
	}
}
