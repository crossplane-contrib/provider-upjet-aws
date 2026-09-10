// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package diffutils

import (
	"bytes"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// SuppressComputedOnlySetRehash suppresses the perpetual diff a TypeSet block
// produces when its elements are re-keyed only because the desired
// configuration leaves an Optional+Computed attribute of an element unset.
//
// A set element's identity is a hash of its content, so an Optional+Computed
// attribute that the external API fills in - an ARN, a generated ID, a
// server-side default - makes the observed element hash differently from the
// element the configuration produces. The plugin SDK can only express that as
// a removal of one element plus an addition of another, and no apply can
// settle it: the next observation re-populates the attribute and the diff comes
// back. Terraform itself does not hit this, because Terraform core merges prior
// values into null Optional+Computed attributes (objchange.ProposedNew) before
// the legacy differ ever runs; upjet calls the differ directly with the desired
// configuration as-is, so the merge never happens.
//
// This function performs the equivalent correlation after the fact: for each
// named field it pairs every desired element with a distinct observed element
// that could have produced it - equal in every attribute the configuration
// actually specifies, differing only where the configuration leaves an
// Optional+Computed attribute at its zero value - and drops the field's
// attributes from the diff when, and only when, every element pairs up.
//
// It deliberately does not suppress:
//
//   - element additions or removals, i.e. a differing element count. An API
//     that returns more set elements than the configuration declares is a
//     separate problem, and one that plain Terraform reports too. The
//     configuration has to declare them (or drop the block entirely).
//   - a change to any attribute the configuration does specify, including one
//     that is Optional+Computed.
//   - a change nested inside a non-Optional+Computed sub-block. Nested content
//     is compared exactly, so such a difference blocks the pairing and the diff
//     survives. Conservative by design: it errs towards reporting drift.
//
// Known limitation: the plugin SDK cannot distinguish "unset" from "set to the
// zero value" in a diff, so a configuration that deliberately sets an
// Optional+Computed attribute to "", 0 or false is read as leaving it unset,
// and a change to that attribute alone is suppressed. This is the same
// conflation the SDK's own DiffSuppressFunc machinery lives with.
//
// Each field must name a TypeSet whose Elem is a *schema.Resource, addressable
// the way helper/schema.ResourceData.GetChange addresses it - a top-level
// attribute, or a dotted path through the collections that contain it, e.g.
// "policy_details.0.schedule".
//
// A "*" segment stands in for an index that is not known when the resource is
// configured: the position of an element in a list of arbitrary length, or a
// set element hash. It is expanded against the keys the calculated diff
// actually carries, so
// "policy_details.0.schedule.*.cross_region_copy_rule" addresses the nested set
// of every schedule the diff mentions, and each expansion is then considered
// independently - one schedule can be suppressed while another keeps a real
// change. A path may not end in "*", since it has to name the set itself.
//
// An unknown or unsuitable field is a configuration error and is reported as
// one.
func SuppressComputedOnlySetRehash(r *config.Resource, diff *terraform.InstanceDiff, s *terraform.InstanceState, fields ...string) (*terraform.InstanceDiff, error) {
	sm := r.TerraformResource.SchemaMap()
	// Field lookup is validated before the early return, so that a mistyped
	// field name surfaces on the first reconcile rather than on the first
	// reconcile that happens to carry a diff.
	elemSchemas := make([]*schema.Resource, len(fields))
	for i, f := range fields {
		elem, err := setElementSchema(r.TerraformResource, f)
		if err != nil {
			return nil, err
		}
		elemSchemas[i] = elem
	}
	if isNonUpdateDiff(diff, s) {
		return diff, nil
	}
	d, err := schema.InternalMap(sm).Data(s, diff)
	if err != nil {
		return nil, errors.Wrap(err, "cannot construct diff data")
	}

	for i, f := range fields {
		for _, ef := range expandTFPathWildcards(f, diff.Attributes) {
			suppressSetDiff(ef, elemSchemas[i], diff, d)
		}
	}
	return diff, nil
}

// expandTFPathWildcards resolves the "*" segments of a field path against the
// keys the calculated diff carries, returning one concrete path per index the
// diff actually mentions. A path with no wildcard expands to itself when the
// diff touches it, and to nothing when it does not - which is what makes the
// expansion double as the "has this field changed at all" check.
func expandTFPathWildcards(f string, attrs map[string]*terraform.ResourceAttrDiff) []string {
	segments := strings.Split(f, ".")

	matchedKeys := make(map[string]struct{})
	for attrKey := range attrs {
		attrSegments := strings.Split(attrKey, ".")
		if segmentMatches(attrSegments, segments) {
			setPath := strings.Join(attrSegments[:len(segments)], ".")
			matchedKeys[setPath] = struct{}{}
		}
	}
	return slices.Collect(maps.Keys(matchedKeys))
}

// segmentMatches reports whether the field path segments match a prefix of the
// diff attribute key segments, with "*" matching any single segment. Matching
// segment by segment rather than by string prefix keeps a field from claiming a
// sibling whose name it happens to be a prefix of.
func segmentMatches(attrSegments, segments []string) bool {
	if len(segments) > len(attrSegments) {
		return false
	}
	for i, s := range segments {
		if s == "*" {
			continue
		}
		if s != attrSegments[i] {
			return false
		}
	}
	return true
}

// SuppressComputedOnlySetRehashFor returns a config.CustomDiff that applies
// SuppressComputedOnlySetRehash to the given fields, for the common case of a
// resource that needs no other diff customization:
//
//	p.AddResourceConfigurator("aws_appstream_stack", func(r *config.Resource) {
//		r.TerraformCustomDiff = diffutils.SuppressComputedOnlySetRehashFor(r, "storage_connectors")
//	})
//
// Compose SuppressComputedOnlySetRehash directly when the resource already has
// a TerraformCustomDiff.
func SuppressComputedOnlySetRehashFor(r *config.Resource, fields ...string) config.CustomDiff {
	return func(diff *terraform.InstanceDiff, s *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
		return SuppressComputedOnlySetRehash(r, diff, s, fields...)
	}
}

// suppressSetDiff drops the diff attributes of the set at the concrete path f
// when every desired element pairs up with a distinct observed one, i.e. when
// the whole change is a re-hash. It leaves the diff untouched otherwise.
func suppressSetDiff(f string, elemSchema *schema.Resource, diff *terraform.InstanceDiff, d *schema.ResourceData) {
	if !d.HasChange(f) {
		return
	}
	o, n := d.GetChange(f)
	observed, ok := o.(*schema.Set)
	if !ok {
		return
	}
	desired, ok := n.(*schema.Set)
	if !ok {
		return
	}
	// A differing element count is a genuine addition or removal.
	if observed.Len() != desired.Len() || desired.Len() == 0 {
		return
	}
	if !setElementsPairUp(observed.List(), desired.List(), elemSchema) {
		return
	}
	prefix := f + "."
	for k := range diff.Attributes {
		if strings.HasPrefix(k, prefix) {
			delete(diff.Attributes, k)
		}
	}
}

// setElementSchema resolves a dotted field path to the *schema.Resource
// describing the elements of the TypeSet it names.
func setElementSchema(res *schema.Resource, path string) (*schema.Resource, error) { //nolint:gocyclo // easier to follow as a unit
	if path == "" {
		return nil, errors.New("set field path cannot be empty")
	}
	segments := strings.Split(path, ".")
	if segments[len(segments)-1] == "*" {
		return nil, errors.Errorf("set field path %q cannot end with *", path)
	}
	// Index segments address an element of the collection we are already on,
	// which the element schema already describes, and which config.GetSchema
	// does not expect in a field path. A "*" stands for such an index when it
	// is not known statically - a set element hash, or the position of an
	// element in a list of arbitrary length.
	named := make([]string, 0, len(segments))
	for _, segment := range segments {
		if segment == "*" {
			continue
		}
		if _, err := strconv.Atoi(segment); err == nil {
			continue
		}
		named = append(named, segment)
	}
	if len(named) == 0 {
		return nil, errors.Errorf("set field path %q names no attribute", path)
	}
	current := config.GetSchema(res, strings.Join(named, "."))
	if current == nil {
		return nil, errors.Errorf("no such attribute %q in the Terraform schema, cannot suppress set re-hash", path)
	}
	if current.Type != schema.TypeSet {
		return nil, errors.Errorf("attribute %q is a %s, set re-hash suppression only applies to a TypeSet", path, current.Type.String())
	}
	elem, ok := current.Elem.(*schema.Resource)
	if !ok {
		return nil, errors.Errorf("attribute %q is a set of primitives, set re-hash suppression only applies to a set of blocks", path)
	}
	return elem, nil
}

// setElementsPairUp reports whether every desired element can be paired with a
// distinct observed element that could have produced it, i.e. whether the whole
// change is an artifact of Optional+Computed attributes the configuration
// leaves unset. Pairing is greedy, and - as in Terraform's own set element
// correlation - is a heuristic when several elements reduce to the same
// comparison value.
func setElementsPairUp(observed, desired []any, elem *schema.Resource) bool {
	paired := make([]bool, len(observed))
	for _, d := range desired {
		dm, ok := d.(map[string]any)
		if !ok {
			return false
		}
		matched := false
		for i, o := range observed {
			if paired[i] {
				continue
			}
			om, ok := o.(map[string]any)
			if !ok {
				return false
			}
			if couldBePriorElement(om, dm, elem) {
				paired[i] = true
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

// couldBePriorElement reports whether the observed element could be the prior
// state of the desired element: equal everywhere the configuration speaks, and
// free to differ only where the configuration leaves an Optional+Computed
// attribute at its zero value.
func couldBePriorElement(observed, desired map[string]any, elem *schema.Resource) bool {
	for name, s := range elem.SchemaMap() {
		// Attributes that are neither Optional nor Required cannot be set from
		// the configuration, and do not contribute to the element hash either
		// (see helper/schema.SerializeResourceForHash), so they cannot be the
		// reason the element was re-keyed.
		if !s.Optional && !s.Required {
			continue
		}
		if s.Optional && s.Computed && isZeroAttrValue(s, desired[name]) {
			continue
		}
		if serializeAttrValue(s, observed[name]) != serializeAttrValue(s, desired[name]) {
			return false
		}
	}
	return true
}

// isZeroAttrValue reports whether v is the zero value of its schema, which is
// how the plugin SDK represents an attribute the configuration does not set.
func isZeroAttrValue(s *schema.Schema, v any) bool {
	if v == nil {
		return true
	}
	return serializeAttrValue(s, v) == serializeAttrValue(s, s.ZeroValue())
}

// serializeAttrValue renders an attribute value into the plugin SDK's own
// canonical form, the one its set hashes are computed from. Comparing the
// serializations rather than the values keeps nested lists, sets and blocks
// comparable without a type switch per shape.
func serializeAttrValue(s *schema.Schema, v any) string {
	var buf bytes.Buffer
	schema.SerializeValueForHash(&buf, v, s)
	return buf.String()
}
