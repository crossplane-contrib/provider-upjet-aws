// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package diffutils

import (
	"hash/crc32"
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// parameterHash mirrors terraform-provider-aws
// internal/service/rds.parameterHash, the Set hash function of the
// `parameter` attribute of aws_db_parameter_group and
// aws_rds_cluster_parameter_group. It is replicated here because the
// element identity of the set - and therefore the shape of the diff
// this package operates on - depends on it: because `apply_method`
// contributes to the hash, changing only the apply method of a
// parameter shows up as a removal of one set element plus an addition
// of another, rather than as an in-place attribute change.
func parameterHash(v any) int {
	var str strings.Builder
	m := v.(map[string]any)

	// Names are hashed lower case to match how the AWS provider stores
	// them when flattening the API response.
	str.WriteString(strings.ToLower(m["name"].(string)))
	str.WriteRune('-')
	str.WriteString(strings.ToLower(m["apply_method"].(string)))
	str.WriteRune('-')
	str.WriteString(m["value"].(string))
	str.WriteRune('-')

	h := int(crc32.ChecksumIEEE([]byte(str.String())))
	if h < 0 {
		h = -h
	}
	return h
}

// parameterGroupSchema mirrors the relevant subset of the
// aws_db_parameter_group Terraform schema. aws_rds_cluster_parameter_group
// declares the very same `parameter` attribute.
func parameterGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"arn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Default:  "Managed by Terraform",
		},
		"family": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"parameter": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"apply_method": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "immediate",
					},
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Set: parameterHash,
		},
	}
}

func parameterGroupResource() *config.Resource {
	return &config.Resource{
		TerraformResource: &schema.Resource{
			Schema: parameterGroupSchema(),
		},
	}
}

// param is a shorthand for a single element of the `parameter` set.
func param(name, value, applyMethod string) map[string]any {
	return map[string]any{
		"name":         name,
		"value":        value,
		"apply_method": applyMethod,
	}
}

// baseAttributes are the non-`parameter` attributes that are held
// identical between state and config by the test helpers, so that the
// generated diff only contains what a test case is about.
func baseAttributes() map[string]any {
	return map[string]any{
		"arn":         "arn:aws:rds:us-east-1:123456789012:pg:example",
		"description": "Managed by Terraform",
		"family":      "postgres14",
		"name":        "example",
	}
}

// paramGroupState builds an InstanceState the way the Terraform plugin
// SDK would flatten it, i.e. with `parameter` set elements keyed by
// their set hash.
func paramGroupState(t *testing.T, params []any, overrides map[string]any) *terraform.InstanceState {
	t.Helper()
	d, err := schema.InternalMap(parameterGroupSchema()).Data(&terraform.InstanceState{ID: "example"}, nil)
	if err != nil {
		t.Fatalf("cannot construct resource data for the state: %v", err)
	}
	attrs := baseAttributes()
	for k, v := range overrides {
		attrs[k] = v
	}
	for k, v := range attrs {
		if err := d.Set(k, v); err != nil {
			t.Fatalf("cannot set %q in the state: %v", k, err)
		}
	}
	if err := d.Set("parameter", params); err != nil {
		t.Fatalf("cannot set \"parameter\" in the state: %v", err)
	}
	return d.State()
}

// paramGroupDiff produces a diff between the given state and the given
// desired configuration using the plugin SDK's own diff machinery, the
// same way upjet's Terraform plugin SDK external client does before
// handing the diff to a config.Resource's TerraformCustomDiff.
func paramGroupDiff(t *testing.T, s *terraform.InstanceState, params []any, overrides map[string]any) *terraform.InstanceDiff {
	t.Helper()
	raw := baseAttributes()
	for k, v := range overrides {
		raw[k] = v
	}
	raw["parameter"] = params
	diff, err := schema.InternalMap(parameterGroupSchema()).Diff(t.Context(), s, terraform.NewResourceConfigRaw(raw), nil, nil, false)
	if err != nil {
		t.Fatalf("cannot construct the instance diff: %v", err)
	}
	return diff
}

// snapshotAttributes dereferences a diff's attributes so that they can
// be compared after SuppressAWSParameterGroupDiff has mutated the diff
// in place.
func snapshotAttributes(diff *terraform.InstanceDiff) map[string]terraform.ResourceAttrDiff {
	if diff == nil {
		return nil
	}
	snapshot := make(map[string]terraform.ResourceAttrDiff, len(diff.Attributes))
	for k, v := range diff.Attributes {
		snapshot[k] = *v
	}
	return snapshot
}

func filterAttributes(attrs map[string]terraform.ResourceAttrDiff, isParameter bool) map[string]terraform.ResourceAttrDiff {
	filtered := make(map[string]terraform.ResourceAttrDiff, len(attrs))
	for k, v := range attrs {
		if strings.HasPrefix(k, "parameter.") == isParameter {
			filtered[k] = v
		}
	}
	return filtered
}

func TestSuppressAWSParameterGroupDiff(t *testing.T) {
	type args struct {
		// stateParams is the `parameter` set as observed in the
		// external resource.
		stateParams []any
		// configParams is the desired `parameter` set.
		configParams []any
		// configOverrides are non-`parameter` attributes that differ
		// from the state.
		configOverrides map[string]any
	}
	type want struct {
		// suppressed asserts that all `parameter.*` attributes have
		// been dropped from the diff. When false, the `parameter.*`
		// attributes must be left exactly as the plugin SDK
		// calculated them.
		suppressed bool
		// nonParameterAttributes are the diff attribute names that
		// must survive the call untouched.
		nonParameterAttributes []string
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"ApplyMethodOnlyChange": {
			reason: "A diff that only changes the apply method of a parameter should be suppressed.",
			args: args{
				stateParams:  []any{param("max_connections", "100", "immediate")},
				configParams: []any{param("max_connections", "100", "pending-reboot")},
			},
			want: want{suppressed: true},
		},
		"ApplyMethodOnlyChangeAmongMultipleParameters": {
			reason: "A diff that only changes the apply method of a subset of the parameters should be suppressed.",
			args: args{
				stateParams: []any{
					param("max_connections", "100", "immediate"),
					param("log_min_duration_statement", "500", "immediate"),
					param("work_mem", "4096", "pending-reboot"),
				},
				configParams: []any{
					param("max_connections", "100", "pending-reboot"),
					param("log_min_duration_statement", "500", "immediate"),
					param("work_mem", "4096", "immediate"),
				},
			},
			want: want{suppressed: true},
		},
		"ApplyMethodOnlyChangeWithOtherAttributeChange": {
			reason: "Only the parameter diff should be suppressed, the diff of the other attributes should be preserved.",
			args: args{
				stateParams:     []any{param("max_connections", "100", "immediate")},
				configParams:    []any{param("max_connections", "100", "pending-reboot")},
				configOverrides: map[string]any{"description": "Managed by Crossplane"},
			},
			want: want{
				suppressed:             true,
				nonParameterAttributes: []string{"description"},
			},
		},
		"ValueChange": {
			reason: "A parameter value change must not be suppressed, otherwise the update is never applied.",
			args: args{
				stateParams:  []any{param("max_connections", "100", "immediate")},
				configParams: []any{param("max_connections", "200", "immediate")},
			},
			want: want{suppressed: false},
		},
		"ValueAndApplyMethodChange": {
			reason: "A parameter value change must not be suppressed even when its apply method changes as well.",
			args: args{
				stateParams:  []any{param("max_connections", "100", "immediate")},
				configParams: []any{param("max_connections", "200", "pending-reboot")},
			},
			want: want{suppressed: false},
		},
		"ValueChangeOfOneOfManyParameters": {
			reason: "A single parameter value change must not be suppressed together with the apply method only changes.",
			args: args{
				stateParams: []any{
					param("max_connections", "100", "immediate"),
					param("work_mem", "4096", "immediate"),
				},
				configParams: []any{
					param("max_connections", "100", "pending-reboot"),
					param("work_mem", "8192", "immediate"),
				},
			},
			want: want{suppressed: false},
		},
		"ParameterAdded": {
			reason: "Adding a parameter must not be suppressed.",
			args: args{
				stateParams: []any{param("max_connections", "100", "immediate")},
				configParams: []any{
					param("max_connections", "100", "pending-reboot"),
					param("work_mem", "4096", "immediate"),
				},
			},
			want: want{suppressed: false},
		},
		"ParameterRemoved": {
			reason: "Removing a parameter must not be suppressed.",
			args: args{
				stateParams: []any{
					param("max_connections", "100", "immediate"),
					param("work_mem", "4096", "immediate"),
				},
				configParams: []any{param("max_connections", "100", "pending-reboot")},
			},
			want: want{suppressed: false},
		},
		"ParameterReplaced": {
			reason: "Replacing a parameter with another one must not be suppressed although the parameter count is unchanged.",
			args: args{
				stateParams:  []any{param("max_connections", "100", "immediate")},
				configParams: []any{param("work_mem", "100", "immediate")},
			},
			want: want{suppressed: false},
		},
		"DuplicateParameterNamesInConfig": {
			reason: "Duplicate parameter names collapse when flattened, so their apply methods cannot be correlated and the diff must not be suppressed.",
			args: args{
				stateParams: []any{
					param("max_connections", "100", "immediate"),
					param("max_connections", "200", "immediate"),
				},
				configParams: []any{
					param("max_connections", "100", "pending-reboot"),
					param("max_connections", "200", "pending-reboot"),
				},
			},
			want: want{suppressed: false},
		},
		"DuplicateParameterNamesInState": {
			reason: "A duplicate parameter name on only one side of the diff must not be suppressed either.",
			args: args{
				stateParams: []any{
					param("max_connections", "100", "immediate"),
					param("max_connections", "200", "immediate"),
				},
				configParams: []any{
					param("max_connections", "100", "pending-reboot"),
					param("work_mem", "4096", "pending-reboot"),
				},
			},
			want: want{suppressed: false},
		},
		"OnlyOtherAttributeChanged": {
			reason: "A diff without any parameter change should be returned as is.",
			args: args{
				stateParams:     []any{param("max_connections", "100", "immediate")},
				configParams:    []any{param("max_connections", "100", "immediate")},
				configOverrides: map[string]any{"description": "Managed by Crossplane"},
			},
			want: want{
				suppressed:             false,
				nonParameterAttributes: []string{"description"},
			},
		},
		"NoParametersAtAll": {
			reason: "A diff of a parameter group without any parameter should be returned as is.",
			args: args{
				stateParams:     []any{},
				configParams:    []any{},
				configOverrides: map[string]any{"description": "Managed by Crossplane"},
			},
			want: want{
				suppressed:             false,
				nonParameterAttributes: []string{"description"},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := paramGroupState(t, tc.args.stateParams, nil)
			diff := paramGroupDiff(t, s, tc.args.configParams, tc.args.configOverrides)
			before := snapshotAttributes(diff)

			got, err := SuppressAWSParameterGroupDiff(parameterGroupResource(), diff, s)
			if err != nil {
				t.Fatalf("SuppressAWSParameterGroupDiff(...): unexpected error: %v\nReason: %s", err, tc.reason)
			}
			if got == nil {
				t.Fatalf("SuppressAWSParameterGroupDiff(...): unexpected nil diff\nReason: %s", tc.reason)
			}
			after := snapshotAttributes(got)

			// The parameter attributes are either all gone, or all
			// exactly as the plugin SDK calculated them.
			wantParameters := filterAttributes(before, true)
			if tc.want.suppressed {
				if len(wantParameters) == 0 {
					t.Fatalf("test case is not exercising diff suppression: the calculated diff has no \"parameter.*\" attributes\nReason: %s", tc.reason)
				}
				wantParameters = map[string]terraform.ResourceAttrDiff{}
			}
			if d := cmp.Diff(wantParameters, filterAttributes(after, true)); d != "" {
				t.Errorf("SuppressAWSParameterGroupDiff(...): -want \"parameter.*\" diff attributes, +got:\n%s\nReason: %s", d, tc.reason)
			}

			// Attributes other than the parameter set are never touched.
			if d := cmp.Diff(filterAttributes(before, false), filterAttributes(after, false)); d != "" {
				t.Errorf("SuppressAWSParameterGroupDiff(...): -want non-\"parameter.*\" diff attributes, +got:\n%s\nReason: %s", d, tc.reason)
			}
			for _, a := range tc.want.nonParameterAttributes {
				if _, ok := got.Attributes[a]; !ok {
					t.Errorf("SuppressAWSParameterGroupDiff(...): expected attribute %q to be present in the returned diff\nReason: %s", a, tc.reason)
				}
			}
		})
	}
}

func TestSuppressAWSParameterGroupDiffNonUpdateDiff(t *testing.T) {
	// A diff that would be suppressed if it were an update diff, so
	// that the non-update short circuits are what is actually under
	// test.
	applyMethodOnlyDiff := func(t *testing.T) (*terraform.InstanceState, *terraform.InstanceDiff) {
		t.Helper()
		s := paramGroupState(t, []any{param("max_connections", "100", "immediate")}, nil)
		return s, paramGroupDiff(t, s, []any{param("max_connections", "100", "pending-reboot")}, nil)
	}

	cases := map[string]struct {
		reason string
		mutate func(t *testing.T, s *terraform.InstanceState, diff *terraform.InstanceDiff) (*terraform.InstanceState, *terraform.InstanceDiff)
	}{
		"NilDiff": {
			reason: "A nil diff should be returned as is.",
			mutate: func(_ *testing.T, s *terraform.InstanceState, _ *terraform.InstanceDiff) (*terraform.InstanceState, *terraform.InstanceDiff) {
				return s, nil
			},
		},
		"EmptyDiff": {
			reason: "A diff without any attribute should be returned as is.",
			mutate: func(_ *testing.T, s *terraform.InstanceState, _ *terraform.InstanceDiff) (*terraform.InstanceState, *terraform.InstanceDiff) {
				return s, terraform.NewInstanceDiff()
			},
		},
		"DestroyDiff": {
			reason: "A destroy diff should be returned as is.",
			mutate: func(_ *testing.T, s *terraform.InstanceState, diff *terraform.InstanceDiff) (*terraform.InstanceState, *terraform.InstanceDiff) {
				diff.Destroy = true
				return s, diff
			},
		},
		"NilState": {
			reason: "A diff with no state, i.e. a creation diff, should be returned as is.",
			mutate: func(_ *testing.T, _ *terraform.InstanceState, diff *terraform.InstanceDiff) (*terraform.InstanceState, *terraform.InstanceDiff) {
				return nil, diff
			},
		},
		"EmptyState": {
			reason: "A diff with an empty state, i.e. a creation diff, should be returned as is.",
			mutate: func(_ *testing.T, _ *terraform.InstanceState, diff *terraform.InstanceDiff) (*terraform.InstanceState, *terraform.InstanceDiff) {
				return &terraform.InstanceState{}, diff
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s, diff := applyMethodOnlyDiff(t)
			s, diff = tc.mutate(t, s, diff)
			before := snapshotAttributes(diff)

			got, err := SuppressAWSParameterGroupDiff(parameterGroupResource(), diff, s)
			if err != nil {
				t.Fatalf("SuppressAWSParameterGroupDiff(...): unexpected error: %v\nReason: %s", err, tc.reason)
			}
			if got != diff {
				t.Fatalf("SuppressAWSParameterGroupDiff(...): expected the input diff to be returned as is\nReason: %s", tc.reason)
			}
			if d := cmp.Diff(before, snapshotAttributes(got)); d != "" {
				t.Errorf("SuppressAWSParameterGroupDiff(...): -want diff attributes, +got:\n%s\nReason: %s", d, tc.reason)
			}
		})
	}
}

func TestFlattenDBParameterSet(t *testing.T) {
	cases := map[string]struct {
		reason string
		tfList []any
		want   map[string]awsDBParameter
	}{
		"Empty": {
			reason: "An empty list should flatten into an empty map.",
			tfList: []any{},
			want:   map[string]awsDBParameter{},
		},
		"Parameters": {
			reason: "Every parameter should be keyed by its name.",
			tfList: []any{
				param("max_connections", "100", "immediate"),
				param("work_mem", "4096", "pending-reboot"),
			},
			want: map[string]awsDBParameter{
				"max_connections": {Value: "100", ApplyMethod: "immediate"},
				"work_mem":        {Value: "4096", ApplyMethod: "pending-reboot"},
			},
		},
		"ApplyMethodLowerCased": {
			reason: "Apply methods should be normalized to lower case so that they compare case insensitively.",
			tfList: []any{param("max_connections", "100", "Pending-Reboot")},
			want: map[string]awsDBParameter{
				"max_connections": {Value: "100", ApplyMethod: "pending-reboot"},
			},
		},
		"ValueCasePreserved": {
			reason: "Parameter values are case sensitive, hence they should be preserved verbatim.",
			tfList: []any{param("log_statement", "DDL", "immediate")},
			want: map[string]awsDBParameter{
				"log_statement": {Value: "DDL", ApplyMethod: "immediate"},
			},
		},
		"EmptyApplyMethod": {
			reason: "An unset apply method should be left empty rather than defaulted, so that it compares equal to another unset one.",
			tfList: []any{param("max_connections", "100", "")},
			want: map[string]awsDBParameter{
				"max_connections": {Value: "100"},
			},
		},
		"NonMapEntrySkipped": {
			reason: "An entry that is not a map should be skipped instead of panicking.",
			tfList: []any{"max_connections", param("work_mem", "4096", "immediate")},
			want: map[string]awsDBParameter{
				"work_mem": {Value: "4096", ApplyMethod: "immediate"},
			},
		},
		"EmptyNameSkipped": {
			reason: "An entry without a name cannot be correlated, hence it should be skipped.",
			tfList: []any{param("", "100", "immediate"), param("work_mem", "4096", "immediate")},
			want: map[string]awsDBParameter{
				"work_mem": {Value: "4096", ApplyMethod: "immediate"},
			},
		},
		"MissingNameSkipped": {
			reason: "An entry with a non-string name should be skipped instead of panicking.",
			tfList: []any{map[string]any{"value": "100", "apply_method": "immediate"}},
			want:   map[string]awsDBParameter{},
		},
		"MissingValue": {
			reason: "An entry without a value should be flattened with an empty value instead of panicking.",
			tfList: []any{map[string]any{"name": "max_connections", "apply_method": "immediate"}},
			want: map[string]awsDBParameter{
				"max_connections": {ApplyMethod: "immediate"},
			},
		},
		"DuplicateNamesCollapse": {
			reason: "Duplicate parameter names collapse into a single entry, which is what lets the caller detect them by comparing the flattened count with the set length.",
			tfList: []any{
				param("max_connections", "100", "immediate"),
				param("max_connections", "200", "immediate"),
			},
			want: map[string]awsDBParameter{
				"max_connections": {Value: "200", ApplyMethod: "immediate"},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if d := cmp.Diff(tc.want, flattenDBParameterSet(tc.tfList)); d != "" {
				t.Errorf("flattenDBParameterSet(...): -want, +got:\n%s\nReason: %s", d, tc.reason)
			}
		})
	}
}

func TestIsNonUpdateDiff(t *testing.T) {
	attrs := map[string]*terraform.ResourceAttrDiff{"description": {Old: "a", New: "b"}}
	stateAttrs := map[string]string{"description": "a"}

	cases := map[string]struct {
		reason string
		diff   *terraform.InstanceDiff
		state  *terraform.InstanceState
		want   bool
	}{
		"UpdateDiff": {
			reason: "A non-empty diff on a non-empty state is an update diff.",
			diff:   &terraform.InstanceDiff{Attributes: attrs},
			state:  &terraform.InstanceState{ID: "example", Attributes: stateAttrs},
			want:   false,
		},
		"NilDiff": {
			reason: "A nil diff is not an update diff.",
			diff:   nil,
			state:  &terraform.InstanceState{ID: "example", Attributes: stateAttrs},
			want:   true,
		},
		"EmptyDiff": {
			reason: "A diff without attributes is not an update diff.",
			diff:   terraform.NewInstanceDiff(),
			state:  &terraform.InstanceState{ID: "example", Attributes: stateAttrs},
			want:   true,
		},
		"DestroyDiff": {
			reason: "A destroy diff is not an update diff.",
			diff:   &terraform.InstanceDiff{Attributes: attrs, Destroy: true},
			state:  &terraform.InstanceState{ID: "example", Attributes: stateAttrs},
			want:   true,
		},
		"NilState": {
			reason: "A diff without a state is a creation diff, not an update diff.",
			diff:   &terraform.InstanceDiff{Attributes: attrs},
			state:  nil,
			want:   true,
		},
		"EmptyState": {
			reason: "A diff on a state without attributes is a creation diff, not an update diff.",
			diff:   &terraform.InstanceDiff{Attributes: attrs},
			state:  &terraform.InstanceState{},
			want:   true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := isNonUpdateDiff(tc.diff, tc.state); got != tc.want {
				t.Errorf("isNonUpdateDiff(...): want %t, got %t\nReason: %s", tc.want, got, tc.reason)
			}
		})
	}
}
