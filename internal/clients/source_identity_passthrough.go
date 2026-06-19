// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"regexp"

	"github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"

	"github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

const (
	// AnnotationSourceIdentity is the well-known annotation a managed resource
	// (or its composing Composition) sets when the referenced ProviderConfig
	// opts in to pass-through SourceIdentity. A single shared ProviderConfig
	// can therefore attribute STS sessions to the originating Claim without
	// templating a ProviderConfig per resource.
	AnnotationSourceIdentity = "aws.upbound.io/source-identity"

	// SentinelFromManagedResource is the opt-in marker that may be set on
	// AssumeRoleOptions.SourceIdentity in a ProviderConfig. When the provider
	// sees this exact value on a chain entry, the effective SourceIdentity for
	// the current reconcile is taken from the managed resource's annotation
	// referenced by AnnotationSourceIdentity. Angle brackets ensure the
	// sentinel cannot collide with a legal AWS SourceIdentity value (which is
	// constrained to [A-Za-z0-9_+=,.@-]).
	SentinelFromManagedResource = "<from-mr>"
)

// sourceIdentityRegexp encodes the AWS server-side constraint for
// SourceIdentity values: 2 to 64 characters of [A-Za-z0-9_+=,.@-].
// See AWS STS AssumeRole API reference.
var sourceIdentityRegexp = regexp.MustCompile(`^[A-Za-z0-9_+=,.@-]{2,64}$`)

const (
	errSourceIdentityAnnotationMissing = "ProviderConfig opted in to pass-through SourceIdentity but the managed resource has no %q annotation"
	errSourceIdentityAnnotationInvalid = "managed resource annotation %q has an invalid SourceIdentity value; must match %s"
)

// applyManagedResourceOverrides performs per-reconcile substitutions on the
// supplied ProviderConfigSpec using values carried on the managed resource.
//
// Currently only AssumeRoleChain[i].SourceIdentity entries that opt in via
// SentinelFromManagedResource are rewritten. The annotation value is
// validated against the AWS SourceIdentity character set and length bounds
// before substitution; an invalid or missing annotation yields an error so
// the reconcile fails closed rather than silently issuing an STS call with
// the literal sentinel.
//
// It returns the slice of values substituted into the chain (in chain
// order). Callers must incorporate these values into any credential cache
// key so that two managed resources sharing a ProviderConfig do not share
// STS sessions when their effective SourceIdentity values differ.
//
// spec is mutated in place. Callers are expected to operate on a per-
// reconcile copy; the on-cluster ProviderConfig is never mutated because
// resolveProviderConfig always returns a freshly constructed value.
func applyManagedResourceOverrides(spec *v1beta1.ProviderConfigSpec, mg xpresource.Managed) ([]string, error) {
	if spec == nil || mg == nil {
		return nil, nil
	}
	var substituted []string
	for i, aro := range spec.AssumeRoleChain {
		if aro.SourceIdentity == nil || *aro.SourceIdentity != SentinelFromManagedResource {
			continue
		}
		v, ok := mg.GetAnnotations()[AnnotationSourceIdentity]
		if !ok || v == "" {
			return nil, errors.Errorf(errSourceIdentityAnnotationMissing, AnnotationSourceIdentity)
		}
		if !sourceIdentityRegexp.MatchString(v) {
			return nil, errors.Errorf(errSourceIdentityAnnotationInvalid, AnnotationSourceIdentity, sourceIdentityRegexp.String())
		}
		// take a fresh pointer so we never alias into the (caller-owned) MR
		// annotation map.
		val := v
		spec.AssumeRoleChain[i].SourceIdentity = &val
		substituted = append(substituted, val)
	}
	return substituted, nil
}
