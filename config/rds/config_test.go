/*
Copyright 2023 Upbound Inc.
*/

package rds

import (
	"context"
	"testing"
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/resource/fake"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ujfake "github.com/crossplane/upjet/pkg/resource/fake"
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
				mg:                 &fake.Managed{},
			},
			want: want{
				err: errors.Wrap(errBoom, errGetPasswordSecret),
			},
		},
		"CannotGetClusterSecret": {
			reason: "An error should be returned if the referenced secret cannot be retrieved.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(errBoom),
				},
				secretRefFieldPath: "",
				toggleFieldPath:    "",
				mg:                 &fake.Managed{},
			},
			want: want{
				err: errors.Wrap(errBoom, errGetPasswordSecret),
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
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name":      "foo",
								"namespace": "bar",
								"key":       "password",
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
		"NoSecretReference": {
			reason: "Should be no-op if the secret reference is not given.",
			args: args{
				secretRefFieldPath: "parameterizable.parameters.passwordSecretRef",
				mg: &ujfake.Terraformed{
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
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name":      "foo",
								"namespace": "bar",
								"key":       "password",
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
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
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
		"ClusterToggleFalse": {
			reason: "Should be no-op if the toggle is set to false.",
			args: args{
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(nil),
				},
				secretRefFieldPath: "parameterizable.parameters.masterPasswordSecretRef",
				toggleFieldPath:    "parameterizable.parameters.autoGeneratePassword",
				mg: &ujfake.Terraformed{
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
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name":      "foo",
								"namespace": "bar",
								"key":       "password",
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
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"masterPasswordSecretRef": map[string]any{
								"name":      "foo",
								"namespace": "bar",
								"key":       "password",
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
							Name: "foo-mgd",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"passwordSecretRef": map[string]any{
								"name":      "foo",
								"namespace": "bar",
								"key":       "password",
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
							Name: "foo-mgd",
						},
					},
					Parameterizable: ujfake.Parameterizable{
						Parameters: map[string]any{
							"masterPasswordSecretRef": map[string]any{
								"name":      "foo",
								"namespace": "bar",
								"key":       "password",
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
