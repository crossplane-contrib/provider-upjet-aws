// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package common

import (
	"context"

	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/password"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/upbound/provider-aws/config/namespaced/common"

	// PathARNExtractor is the golang path to ARNExtractor function
	// in this package.
	PathARNExtractor = SelfPackagePath + ".ARNExtractor()"

	// PathTerraformIDExtractor is the golang path to TerraformID extractor
	// function in this package.
	PathTerraformIDExtractor = SelfPackagePath + ".TerraformID()"

	// VersionV1Beta1 is used for resources that meet the v1beta1 criteria
	// here: https://github.com/upbound/arch/pull/33
	VersionV1Beta1 = "v1beta1"

	// ErrGetPasswordSecret is an error string for failing to get password secret
	ErrGetPasswordSecret = "cannot get password secret"

	// errManagedNotNamespaced is an error string for non-namespaced MRs
	errManagedNotNamespaced = "managed resource is not namespaced"
)

// ARNExtractor extracts ARN of the resources from "status.atProvider.arn" which
// is quite common among all AWS resources.
func ARNExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.arn")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}

// TerraformID returns the Terraform ID string of the resource without any
// manipulation.
func TerraformID() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		return tr.GetID()
	}
}

// PasswordGenerator returns an InitializerFn that will generate a password
// for a resource if the toggle field is set to true and the secret referenced
// by the secretRefFieldPath is not found or does not have content corresponding
// to the password key.
func PasswordGenerator(secretRefFieldPath, toggleFieldPath string) config.NewInitializerFn { //nolint:gocyclo
	// NOTE(muvaf): This function is just 1 point over the cyclo limit but there
	// is no easy way to reduce it without making it harder to read.
	return func(client client.Client) managed.Initializer {
		return managed.InitializerFn(func(ctx context.Context, mg xpresource.Managed) error {
			// this would be a programming/configuration error
			// should be used only for namespaced MRs
			if mg.GetNamespace() == "" {
				return errors.New(errManagedNotNamespaced)
			}

			paved, err := fieldpath.PaveObject(mg)
			if err != nil {
				return errors.Wrap(err, "cannot pave object")
			}
			sel := &v1.LocalSecretKeySelector{}
			if err := paved.GetValueInto(secretRefFieldPath, sel); err != nil {
				return errors.Wrapf(xpresource.Ignore(fieldpath.IsNotFound, err), "cannot unmarshal %s into a secret key selector", secretRefFieldPath)
			}
			s := &corev1.Secret{}
			if err := client.Get(ctx, types.NamespacedName{Namespace: mg.GetNamespace(), Name: sel.Name}, s); xpresource.IgnoreNotFound(err) != nil {
				return errors.Wrap(err, ErrGetPasswordSecret)
			}
			if err == nil && len(s.Data[sel.Key]) != 0 {
				// Password is already set.
				return nil
			}
			// At this point, either the secret doesn't exist, or it doesn't
			// have the password filled.
			if gen, err := paved.GetBool(toggleFieldPath); err != nil || !gen {
				// If there is error, then we return that.
				// If the toggle field is not set to true, then we return nil.
				// Because we don't want to generate a password if the user
				// doesn't want to.
				return errors.Wrapf(xpresource.Ignore(fieldpath.IsNotFound, err), "cannot get the value of %s", toggleFieldPath)
			}
			pw, err := password.Generate()
			if err != nil {
				return errors.Wrap(err, "cannot generate password")
			}
			s.SetName(sel.Name)
			s.SetNamespace(mg.GetNamespace())
			if !meta.WasCreated(s) {
				// We don't want to own the Secret if it is created by someone
				// else, otherwise the deletion of the managed resource will
				// delete the Secret that we didn't create in the first place.
				meta.AddOwnerReference(s, meta.AsOwner(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind())))
			}
			if s.Data == nil {
				s.Data = make(map[string][]byte, 1)
			}
			s.Data[sel.Key] = []byte(pw)
			return errors.Wrap(xpresource.NewAPIPatchingApplicator(client).Apply(ctx, s), "cannot apply password secret")
		})
	}
}

// RemovePolicyVersion removes the "Version" field from a JSON-encoded policy string.
func RemovePolicyVersion(p string) (string, error) {
	var policy any
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(p), &policy); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal the policy from JSON")
	}
	m, ok := policy.(map[string]any)
	if !ok {
		return p, nil
	}
	delete(m, "Version")
	r, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(m)
	return string(r), errors.Wrap(err, "failed to marshal the policy map as JSON")
}

// RemoveDiffIfEmpty removes supplied keys from Terraform diffs, if old and new
// values for the key are empty (""). It is probably safe to delete all such
// keys in the diff. Until we decide to do so, we'll specify the keys, to be on
// the safe side.
func RemoveDiffIfEmpty(keys []string) config.CustomDiff { //nolint:gocyclo // The implementation is pretty straightforward as of this writing.
	return func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
		// Skip diff customization on create
		if state == nil || state.Empty() {
			return diff, nil
		}
		if config == nil {
			return nil, errors.New("resource config cannot be nil")
		}
		// Skip no diff or destroy diffs
		if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
			return diff, nil
		}

		for _, key := range keys {
			if diff.Attributes[key] != nil && diff.Attributes[key].Old == "" && diff.Attributes[key].New == "" {
				delete(diff.Attributes, key)
			}
		}

		return diff, nil
	}
}
