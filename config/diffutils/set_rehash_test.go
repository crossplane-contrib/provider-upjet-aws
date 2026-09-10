package diffutils

import (
	"slices"
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// stackSchema mirrors the relevant subset of the aws_appstream_stack
// Terraform schema. Both of its sets are the shape
// SuppressComputedOnlySetRehash exists for: `storage_connectors` carries
// Optional+Computed attributes that AWS populates and an example manifest
// leaves unset, while `user_settings` is a set AWS returns in full whatever
// the configuration declares.
func stackSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"storage_connectors": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"connector_type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"domains": {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"resource_identifier": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
		"user_settings": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"action": {
						Type:     schema.TypeString,
						Required: true,
					},
					"permission": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}

func stackResource() *config.Resource {
	return &config.Resource{
		TerraformResource: &schema.Resource{
			Schema: stackSchema(),
		},
	}
}

// connector is a shorthand for a single `storage_connectors` element. An empty
// domains slice or resourceIdentifier stands for "not set".
func connector(connectorType string, domains []any, resourceIdentifier string) map[string]any {
	return map[string]any{
		"connector_type":      connectorType,
		"domains":             domains,
		"resource_identifier": resourceIdentifier,
	}
}

// userSetting is a shorthand for a single `user_settings` element.
func userSetting(action, permission string) map[string]any {
	return map[string]any{"action": action, "permission": permission}
}

func stackBaseAttributes() map[string]any {
	return map[string]any{
		"name":        "example",
		"description": "stack description",
	}
}

// stackState builds an InstanceState the way the plugin SDK would flatten an
// observation, i.e. with both sets keyed by their element hashes.
func stackState(t *testing.T, connectors, settings []any) *terraform.InstanceState {
	t.Helper()
	d, err := schema.InternalMap(stackSchema()).Data(&terraform.InstanceState{ID: "example"}, nil)
	if err != nil {
		t.Fatalf("cannot construct resource data for the state: %v", err)
	}
	for k, v := range stackBaseAttributes() {
		if err := d.Set(k, v); err != nil {
			t.Fatalf("cannot set %q in the state: %v", k, err)
		}
	}
	if err := d.Set("storage_connectors", connectors); err != nil {
		t.Fatalf("cannot set \"storage_connectors\" in the state: %v", err)
	}
	if err := d.Set("user_settings", settings); err != nil {
		t.Fatalf("cannot set \"user_settings\" in the state: %v", err)
	}
	return d.State()
}

// stackDiff produces a diff between the given state and the given desired
// configuration with the plugin SDK's own diff machinery, exactly as upjet's
// Terraform plugin SDK external client does before handing the diff to a
// config.Resource's TerraformCustomDiff.
func stackDiff(t *testing.T, s *terraform.InstanceState, connectors, settings []any, overrides map[string]any) *terraform.InstanceDiff {
	t.Helper()
	raw := stackBaseAttributes()
	for k, v := range overrides {
		raw[k] = v
	}
	raw["storage_connectors"] = connectors
	raw["user_settings"] = settings
	diff, err := schema.InternalMap(stackSchema()).Diff(t.Context(), s, terraform.NewResourceConfigRaw(raw), nil, nil, false)
	if err != nil {
		t.Fatalf("cannot construct the instance diff: %v", err)
	}
	return diff
}

// attributesWithPrefix selects the diff attributes under a field.
func attributesWithPrefix(attrs map[string]terraform.ResourceAttrDiff, prefix string) map[string]terraform.ResourceAttrDiff {
	selected := make(map[string]terraform.ResourceAttrDiff, len(attrs))
	for k, v := range attrs {
		if strings.HasPrefix(k, prefix) {
			selected[k] = v
		}
	}
	return selected
}

func TestSuppressComputedOnlySetRehash(t *testing.T) {
	// The connector as AWS reports it once created: the Optional+Computed
	// attributes are populated server side.
	observedConnector := connector("HOMEFOLDERS", []any{"example.com"}, "arn:aws:s3:::bucket")
	// The connector as an example manifest declares it.
	desiredConnector := connector("HOMEFOLDERS", []any{}, "")

	allUserSettings := []any{
		userSetting("AUTO_TIME_ZONE_REDIRECTION", "DISABLED"),
		userSetting("CLIPBOARD_COPY_FROM_LOCAL_DEVICE", "ENABLED"),
		userSetting("CLIPBOARD_COPY_TO_LOCAL_DEVICE", "ENABLED"),
		userSetting("FILE_UPLOAD", "ENABLED"),
	}

	type args struct {
		// stateConnectors, stateSettings are the sets as observed in
		// the external resource.
		stateConnectors []any
		stateSettings   []any
		// configConnectors, configSettings are the desired sets.
		configConnectors []any
		configSettings   []any
		// configOverrides are attributes outside the two sets that
		// differ from the state.
		configOverrides map[string]any
		// fields is what the caller asks to be suppressed.
		fields []string
	}
	type want struct {
		// suppressedPrefixes are the diff key prefixes that must be
		// gone from the diff. Every other attribute must be left
		// exactly as the plugin SDK calculated it.
		suppressedPrefixes []string
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"ComputedOnlyRehash": {
			reason: "An element re-keyed only because the configuration leaves its Optional+Computed attributes unset should be suppressed.",
			args: args{
				stateConnectors:  []any{observedConnector},
				stateSettings:    allUserSettings,
				configConnectors: []any{desiredConnector},
				configSettings:   allUserSettings,
				fields:           []string{"storage_connectors"},
			},
			want: want{suppressedPrefixes: []string{"storage_connectors."}},
		},
		"ComputedOnlyRehashAmongMultipleElements": {
			reason: "Every element pairs up, so the whole set diff should be suppressed.",
			args: args{
				stateConnectors: []any{
					observedConnector,
					connector("GOOGLE_DRIVE", []any{"corp.example.com"}, "arn:aws:s3:::other"),
				},
				stateSettings: allUserSettings,
				configConnectors: []any{
					desiredConnector,
					connector("GOOGLE_DRIVE", []any{}, ""),
				},
				configSettings: allUserSettings,
				fields:         []string{"storage_connectors"},
			},
			want: want{suppressedPrefixes: []string{"storage_connectors."}},
		},
		"ComputedOnlyRehashWithOtherAttributeChange": {
			reason: "Suppressing the set must not disturb an unrelated attribute change in the same diff.",
			args: args{
				stateConnectors:  []any{observedConnector},
				stateSettings:    allUserSettings,
				configConnectors: []any{desiredConnector},
				configSettings:   allUserSettings,
				configOverrides:  map[string]any{"description": "changed"},
				fields:           []string{"storage_connectors"},
			},
			want: want{suppressedPrefixes: []string{"storage_connectors."}},
		},
		"RequiredAttributeChanged": {
			reason: "A configuration that changes an attribute it actually specifies is real drift and must survive.",
			args: args{
				stateConnectors:  []any{observedConnector},
				stateSettings:    allUserSettings,
				configConnectors: []any{connector("ONE_DRIVE", []any{}, "")},
				configSettings:   allUserSettings,
				fields:           []string{"storage_connectors"},
			},
			want: want{},
		},
		"ComputedAttributeSetToADifferentValue": {
			reason: "An Optional+Computed attribute the configuration does specify is not free to differ, so the diff must survive.",
			args: args{
				stateConnectors:  []any{observedConnector},
				stateSettings:    allUserSettings,
				configConnectors: []any{connector("HOMEFOLDERS", []any{}, "arn:aws:s3:::desired")},
				configSettings:   allUserSettings,
				fields:           []string{"storage_connectors"},
			},
			want: want{},
		},
		"ComputedListAttributeSetToADifferentValue": {
			reason: "A collection-typed Optional+Computed attribute the configuration specifies must be compared, not waved through.",
			args: args{
				stateConnectors:  []any{observedConnector},
				stateSettings:    allUserSettings,
				configConnectors: []any{connector("HOMEFOLDERS", []any{"desired.example.com"}, "")},
				configSettings:   allUserSettings,
				fields:           []string{"storage_connectors"},
			},
			want: want{},
		},
		"ElementAdded": {
			reason: "A differing element count is a genuine addition and must survive.",
			args: args{
				stateConnectors:  []any{observedConnector},
				stateSettings:    allUserSettings,
				configConnectors: []any{desiredConnector, connector("GOOGLE_DRIVE", []any{}, "")},
				configSettings:   allUserSettings,
				fields:           []string{"storage_connectors"},
			},
			want: want{},
		},
		"ElementRemoved": {
			reason: "A differing element count is a genuine removal and must survive.",
			args: args{
				stateConnectors: []any{
					observedConnector,
					connector("GOOGLE_DRIVE", []any{"corp.example.com"}, "arn:aws:s3:::other"),
				},
				stateSettings:    allUserSettings,
				configConnectors: []any{desiredConnector},
				configSettings:   allUserSettings,
				fields:           []string{"storage_connectors"},
			},
			want: want{},
		},
		"ServerPopulatedSetSubset": {
			reason: "A set the API returns in full while the configuration declares a subset is not a re-hash, and must survive so that the drift stays visible.",
			args: args{
				stateConnectors: []any{observedConnector},
				stateSettings: append(append([]any{}, allUserSettings...),
					userSetting("DOMAIN_PASSWORD_SIGNIN", "ENABLED"),
					userSetting("PRINTING_TO_LOCAL_DEVICE", "ENABLED"),
				),
				configConnectors: []any{desiredConnector},
				configSettings:   allUserSettings,
				fields:           []string{"storage_connectors", "user_settings"},
			},
			want: want{suppressedPrefixes: []string{"storage_connectors."}},
		},
		"MultipleFields": {
			reason: "Every named field is considered independently.",
			args: args{
				stateConnectors:  []any{observedConnector},
				stateSettings:    allUserSettings,
				configConnectors: []any{desiredConnector},
				configSettings:   allUserSettings,
				fields:           []string{"storage_connectors", "user_settings"},
			},
			want: want{suppressedPrefixes: []string{"storage_connectors."}},
		},
		"FieldNotNamed": {
			reason: "A field the caller does not name must be left alone, however it diffs.",
			args: args{
				stateConnectors:  []any{observedConnector},
				stateSettings:    allUserSettings,
				configConnectors: []any{desiredConnector},
				configSettings:   allUserSettings,
				fields:           []string{"user_settings"},
			},
			want: want{},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := stackState(t, tc.args.stateConnectors, tc.args.stateSettings)
			diff := stackDiff(t, s, tc.args.configConnectors, tc.args.configSettings, tc.args.configOverrides)
			before := snapshotAttributes(diff)

			got, err := SuppressComputedOnlySetRehash(stackResource(), diff, s, tc.args.fields...)
			if err != nil {
				t.Fatalf("SuppressComputedOnlySetRehash(...): unexpected error: %v\nReason: %s", err, tc.reason)
			}
			if got == nil {
				t.Fatalf("SuppressComputedOnlySetRehash(...): unexpected nil diff\nReason: %s", tc.reason)
			}
			after := snapshotAttributes(got)

			want := map[string]terraform.ResourceAttrDiff{}
			for k, v := range before {
				want[k] = v
			}
			for _, prefix := range tc.want.suppressedPrefixes {
				if len(attributesWithPrefix(before, prefix)) == 0 {
					t.Fatalf("test case is not exercising diff suppression: the calculated diff has no %q attributes\nReason: %s", prefix, tc.reason)
				}
				for k := range attributesWithPrefix(want, prefix) {
					delete(want, k)
				}
			}
			if d := cmp.Diff(want, after); d != "" {
				t.Errorf("SuppressComputedOnlySetRehash(...): -want diff attributes, +got:\n%s\nReason: %s", d, tc.reason)
			}
		})
	}
}

func TestSuppressComputedOnlySetRehashFieldErrors(t *testing.T) {
	cases := map[string]struct {
		reason string
		field  string
		// wantErr, when set, must appear in the reported error. It pins the
		// cases where more than one thing about the path is wrong, and the
		// error has to name the right one.
		wantErr string
	}{
		"UnknownField": {
			reason: "A mistyped field name is a configuration error and must be reported.",
			field:  "storage_connector",
		},
		"NotASet": {
			reason: "A field that is not a TypeSet cannot be re-hashed and must be reported.",
			field:  "description",
		},
		"EmptyPath": {
			reason: "An empty field path must be reported.",
			field:  "",
		},
		"UnknownNestedField": {
			reason: "A mistyped nested field name must be reported.",
			field:  "storage_connectors.0.nope",
		},
		"TrailingWildcard": {
			reason:  "A path ending in a wildcard names an element rather than the set itself, and must be reported as such.",
			field:   "storage_connectors.*",
			wantErr: "cannot end with *",
		},
		"OnlyWildcard": {
			reason: "A path of nothing but a wildcard names no attribute and must be reported.",
			field:  "*",
		},
		"WildcardIntoANonSet": {
			reason:  "A wildcard does not excuse the final segment from having to name a TypeSet.",
			field:   "storage_connectors.*.connector_type",
			wantErr: "set re-hash suppression only applies to a TypeSet",
		},
		"WildcardIntoAnUnknownNestedField": {
			reason: "A mistyped field behind a wildcard must be reported, not silently expanded to nothing.",
			field:  "storage_connectors.*.nope",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := stackState(t, []any{connector("HOMEFOLDERS", []any{}, "")}, []any{userSetting("FILE_UPLOAD", "ENABLED")})
			diff := stackDiff(t, s, []any{connector("HOMEFOLDERS", []any{}, "")}, []any{userSetting("FILE_UPLOAD", "ENABLED")},
				map[string]any{"description": "changed"})
			_, err := SuppressComputedOnlySetRehash(stackResource(), diff, s, tc.field)
			if err == nil {
				t.Fatalf("SuppressComputedOnlySetRehash(...): expected an error for field %q\nReason: %s", tc.field, tc.reason)
			}
			if tc.wantErr != "" && !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("SuppressComputedOnlySetRehash(...): expected the error for field %q to mention %q, got: %v\nReason: %s",
					tc.field, tc.wantErr, err, tc.reason)
			}
		})
	}
}

func TestSuppressComputedOnlySetRehashNonUpdateDiff(t *testing.T) {
	connectors := []any{connector("HOMEFOLDERS", []any{"example.com"}, "arn:aws:s3:::bucket")}
	settings := []any{userSetting("FILE_UPLOAD", "ENABLED")}

	cases := map[string]struct {
		reason string
		diff   func(t *testing.T, s *terraform.InstanceState) *terraform.InstanceDiff
		state  func(t *testing.T) *terraform.InstanceState
	}{
		"NilDiff": {
			reason: "A nil diff should be returned as is.",
			diff:   func(*testing.T, *terraform.InstanceState) *terraform.InstanceDiff { return nil },
			state:  func(t *testing.T) *terraform.InstanceState { return stackState(t, connectors, settings) },
		},
		"EmptyDiff": {
			reason: "An empty diff should be returned as is.",
			diff: func(*testing.T, *terraform.InstanceState) *terraform.InstanceDiff {
				return &terraform.InstanceDiff{Attributes: map[string]*terraform.ResourceAttrDiff{}}
			},
			state: func(t *testing.T) *terraform.InstanceState { return stackState(t, connectors, settings) },
		},
		"DestroyDiff": {
			reason: "A destroy diff should be returned as is.",
			diff: func(t *testing.T, s *terraform.InstanceState) *terraform.InstanceDiff {
				d := stackDiff(t, s, []any{connector("HOMEFOLDERS", []any{}, "")}, settings, nil)
				d.Destroy = true
				return d
			},
			state: func(t *testing.T) *terraform.InstanceState { return stackState(t, connectors, settings) },
		},
		"NilState": {
			reason: "A creation diff has no state to correlate against and should be returned as is.",
			diff: func(t *testing.T, _ *terraform.InstanceState) *terraform.InstanceDiff {
				s := stackState(t, connectors, settings)
				return stackDiff(t, s, []any{connector("HOMEFOLDERS", []any{}, "")}, settings, nil)
			},
			state: func(*testing.T) *terraform.InstanceState { return nil },
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := tc.state(t)
			diff := tc.diff(t, s)
			before := snapshotAttributes(diff)

			got, err := SuppressComputedOnlySetRehash(stackResource(), diff, s, "storage_connectors")
			if err != nil {
				t.Fatalf("SuppressComputedOnlySetRehash(...): unexpected error: %v\nReason: %s", err, tc.reason)
			}
			if d := cmp.Diff(before, snapshotAttributes(got)); d != "" {
				t.Errorf("SuppressComputedOnlySetRehash(...): -want diff attributes, +got:\n%s\nReason: %s", d, tc.reason)
			}
		})
	}
}

func TestSuppressComputedOnlySetRehashFor(t *testing.T) {
	s := stackState(t, []any{connector("HOMEFOLDERS", []any{"example.com"}, "arn:aws:s3:::bucket")},
		[]any{userSetting("FILE_UPLOAD", "ENABLED")})
	diff := stackDiff(t, s, []any{connector("HOMEFOLDERS", []any{}, "")},
		[]any{userSetting("FILE_UPLOAD", "ENABLED")}, nil)
	if len(attributesWithPrefix(snapshotAttributes(diff), "storage_connectors.")) == 0 {
		t.Fatal("test is not exercising diff suppression: the calculated diff has no \"storage_connectors.*\" attributes")
	}

	got, err := SuppressComputedOnlySetRehashFor(stackResource(), "storage_connectors")(diff, s, nil)
	if err != nil {
		t.Fatalf("SuppressComputedOnlySetRehashFor(...): unexpected error: %v", err)
	}
	if attrs := attributesWithPrefix(snapshotAttributes(got), "storage_connectors."); len(attrs) != 0 {
		t.Errorf("SuppressComputedOnlySetRehashFor(...): expected the \"storage_connectors.*\" attributes to be suppressed, got: %v", attrs)
	}
}

// policySchema mirrors the shape a nested set takes: the re-keyed TypeSet is
// not a top-level attribute but sits several collections deep. `policy_details`
// is a singleton list, its `schedule` is a list of arbitrary length, and every
// schedule carries a `cross_region_copy_rule` set whose `cmk_arn` AWS fills in
// server side. Modelled on aws_dlm_lifecycle_policy: the positions under
// `schedule` are exactly what a static field path cannot name, and what "*"
// exists for.
func policySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"policy_details": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"schedule": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:     schema.TypeString,
									Required: true,
								},
								"cross_region_copy_rule": {
									Type:     schema.TypeSet,
									Optional: true,
									Computed: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"target": {
												Type:     schema.TypeString,
												Required: true,
											},
											"cmk_arn": {
												Type:     schema.TypeString,
												Optional: true,
												Computed: true,
											},
											"copy_tags": {
												Type:     schema.TypeBool,
												Optional: true,
												Computed: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func policyResource() *config.Resource {
	return &config.Resource{
		TerraformResource: &schema.Resource{
			Schema: policySchema(),
		},
	}
}

// copyRule is a shorthand for a single `cross_region_copy_rule` element. An
// empty cmkARN stands for "not set".
func copyRule(target, cmkARN string) map[string]any {
	return map[string]any{
		"target":    target,
		"cmk_arn":   cmkARN,
		"copy_tags": true,
	}
}

// policySchedule is a shorthand for a single `schedule` element.
func policySchedule(name string, rules ...any) map[string]any {
	return map[string]any{"name": name, "cross_region_copy_rule": rules}
}

// policyRaw renders the whole resource configuration around the schedules.
func policyRaw(name string, schedules []any) map[string]any {
	return map[string]any{
		"name":           name,
		"policy_details": []any{map[string]any{"schedule": schedules}},
	}
}

func policyState(t *testing.T, schedules []any) *terraform.InstanceState {
	t.Helper()
	d, err := schema.InternalMap(policySchema()).Data(&terraform.InstanceState{ID: "example"}, nil)
	if err != nil {
		t.Fatalf("cannot construct resource data for the state: %v", err)
	}
	for k, v := range policyRaw("example", schedules) {
		if err := d.Set(k, v); err != nil {
			t.Fatalf("cannot set %q in the state: %v", k, err)
		}
	}
	return d.State()
}

func policyDiff(t *testing.T, s *terraform.InstanceState, name string, schedules []any) *terraform.InstanceDiff {
	t.Helper()
	diff, err := schema.InternalMap(policySchema()).Diff(t.Context(), s,
		terraform.NewResourceConfigRaw(policyRaw(name, schedules)), nil, nil, false)
	if err != nil {
		t.Fatalf("cannot construct the instance diff: %v", err)
	}
	return diff
}

func TestSuppressComputedOnlySetRehashNestedSet(t *testing.T) {
	// The rules as AWS reports them, with the Optional+Computed cmk_arn
	// populated server side, and as an example manifest declares them.
	observedDaily := copyRule("us-west-2", "arn:aws:kms:::key/daily")
	desiredDaily := copyRule("us-west-2", "")
	observedWeekly := copyRule("eu-west-1", "arn:aws:kms:::key/weekly")
	desiredWeekly := copyRule("eu-west-1", "")

	const (
		dailyPrefix  = "policy_details.0.schedule.0.cross_region_copy_rule."
		weeklyPrefix = "policy_details.0.schedule.1.cross_region_copy_rule."
	)

	type args struct {
		stateSchedules  []any
		configName      string
		configSchedules []any
		fields          []string
	}
	cases := map[string]struct {
		reason             string
		args               args
		suppressedPrefixes []string
	}{
		"WildcardExpandsToEveryListIndex": {
			reason: "A \"*\" index must reach every element the diff mentions, not just the first.",
			args: args{
				stateSchedules: []any{
					policySchedule("daily", observedDaily),
					policySchedule("weekly", observedWeekly),
				},
				configName: "example",
				configSchedules: []any{
					policySchedule("daily", desiredDaily),
					policySchedule("weekly", desiredWeekly),
				},
				fields: []string{"policy_details.*.schedule.*.cross_region_copy_rule"},
			},
			suppressedPrefixes: []string{dailyPrefix, weeklyPrefix},
		},
		"LiteralIndexAddressesOneElement": {
			reason: "A path spelled with literal indices must still work, and must reach only the element it names.",
			args: args{
				stateSchedules: []any{
					policySchedule("daily", observedDaily),
					policySchedule("weekly", observedWeekly),
				},
				configName: "example",
				configSchedules: []any{
					policySchedule("daily", desiredDaily),
					policySchedule("weekly", desiredWeekly),
				},
				fields: []string{"policy_details.0.schedule.0.cross_region_copy_rule"},
			},
			suppressedPrefixes: []string{dailyPrefix},
		},
		"EveryExpansionIsConsideredIndependently": {
			reason: "One expansion carrying real drift must not keep another expansion's pure re-hash in the diff.",
			args: args{
				stateSchedules: []any{
					policySchedule("daily", observedDaily),
					policySchedule("weekly", observedWeekly),
				},
				configName: "example",
				configSchedules: []any{
					policySchedule("daily", desiredDaily),
					// A Required attribute the configuration does
					// specify changed: real drift, and it must survive.
					policySchedule("weekly", copyRule("ap-south-1", "")),
				},
				fields: []string{"policy_details.*.schedule.*.cross_region_copy_rule"},
			},
			suppressedPrefixes: []string{dailyPrefix},
		},
		"WildcardMatchesNoDiffAttribute": {
			reason: "A field the diff does not touch expands to nothing, and the rest of the diff must be left alone.",
			args: args{
				stateSchedules:  []any{policySchedule("daily", observedDaily)},
				configName:      "renamed",
				configSchedules: []any{policySchedule("daily", observedDaily)},
				fields:          []string{"policy_details.*.schedule.*.cross_region_copy_rule"},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := policyState(t, tc.args.stateSchedules)
			diff := policyDiff(t, s, tc.args.configName, tc.args.configSchedules)
			before := snapshotAttributes(diff)

			got, err := SuppressComputedOnlySetRehash(policyResource(), diff, s, tc.args.fields...)
			if err != nil {
				t.Fatalf("SuppressComputedOnlySetRehash(...): unexpected error: %v\nReason: %s", err, tc.reason)
			}
			if got == nil {
				t.Fatalf("SuppressComputedOnlySetRehash(...): unexpected nil diff\nReason: %s", tc.reason)
			}

			want := map[string]terraform.ResourceAttrDiff{}
			for k, v := range before {
				want[k] = v
			}
			for _, prefix := range tc.suppressedPrefixes {
				if len(attributesWithPrefix(before, prefix)) == 0 {
					t.Fatalf("test case is not exercising diff suppression: the calculated diff has no %q attributes\nReason: %s", prefix, tc.reason)
				}
				for k := range attributesWithPrefix(want, prefix) {
					delete(want, k)
				}
			}
			if d := cmp.Diff(want, snapshotAttributes(got)); d != "" {
				t.Errorf("SuppressComputedOnlySetRehash(...): -want diff attributes, +got:\n%s\nReason: %s", d, tc.reason)
			}
		})
	}
}

func TestExpandTFPathWildcards(t *testing.T) {
	// A diff that touches two schedules of a singleton policy_details, plus a
	// sibling attribute whose name starts with the same characters as the set
	// it must not be confused with.
	attrs := map[string]*terraform.ResourceAttrDiff{
		"name": {Old: "example", New: "renamed"},
		"policy_details.0.schedule.0.cross_region_copy_rule.111.target":  {Old: "us-west-2", New: ""},
		"policy_details.0.schedule.0.cross_region_copy_rule.222.target":  {Old: "", New: "us-west-2"},
		"policy_details.0.schedule.0.cross_region_copy_rule_extra.#":     {Old: "0", New: "1"},
		"policy_details.0.schedule.1.cross_region_copy_rule.333.cmk_arn": {NewComputed: true},
		"policy_details.0.schedule.1.name":                               {Old: "weekly", New: "monthly"},
	}

	cases := map[string]struct {
		reason string
		path   string
		want   []string
	}{
		"LiteralPathTheDiffTouches": {
			reason: "A path without a wildcard expands to itself when the diff carries keys under it.",
			path:   "policy_details.0.schedule.0.cross_region_copy_rule",
			want:   []string{"policy_details.0.schedule.0.cross_region_copy_rule"},
		},
		"LiteralPathTheDiffDoesNotTouch": {
			reason: "A path the diff does not carry expands to nothing, which is what stands in for a HasChange check.",
			path:   "policy_details.0.schedule.2.cross_region_copy_rule",
		},
		"WildcardExpandsToEveryIndexInTheDiff": {
			reason: "Every index the diff mentions must produce its own concrete path, once.",
			path:   "policy_details.*.schedule.*.cross_region_copy_rule",
			want: []string{
				"policy_details.0.schedule.0.cross_region_copy_rule",
				"policy_details.0.schedule.1.cross_region_copy_rule",
			},
		},
		"WildcardInOnePositionOnly": {
			reason: "A wildcard must only widen the position it occupies.",
			path:   "policy_details.0.schedule.*.name",
			want:   []string{"policy_details.0.schedule.1.name"},
		},
		"SiblingWithASharedNamePrefix": {
			reason: "Matching is by path segment, so a field must not claim a sibling whose name it is a string prefix of.",
			path:   "policy_details.*.schedule.*.cross_region_copy_rule_extra",
			want:   []string{"policy_details.0.schedule.0.cross_region_copy_rule_extra"},
		},
		"PathLongerThanAnyDiffKey": {
			reason: "A path deeper than every key in the diff matches nothing.",
			path:   "policy_details.0.schedule.1.name.0.deeper",
		},
		"TopLevelAttribute": {
			reason: "A top-level attribute expands to itself.",
			path:   "name",
			want:   []string{"name"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := expandTFPathWildcards(tc.path, attrs)
			slices.Sort(got)
			want := tc.want
			slices.Sort(want)
			if d := cmp.Diff(want, got, cmpopts.EquateEmpty()); d != "" {
				t.Errorf("expandTFPathWildcards(%q, ...): -want, +got:\n%s\nReason: %s", tc.path, d, tc.reason)
			}
		})
	}
}
