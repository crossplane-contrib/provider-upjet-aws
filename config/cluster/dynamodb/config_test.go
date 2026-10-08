// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dynamodb

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGSIUserFieldsHash(t *testing.T) {
	gsi := func(readCapacity, writeCapacity int, nonKeyAttributes ...interface{}) map[string]interface{} {
		return map[string]interface{}{
			"name":               "gsi1",
			"hash_key":           "gsi1pk",
			"range_key":          "",
			"projection_type":    "INCLUDE",
			"read_capacity":      readCapacity,
			"write_capacity":     writeCapacity,
			"non_key_attributes": schema.NewSet(schema.HashString, nonKeyAttributes),
		}
	}
	withComputedFields := gsi(5, 5, "a", "b")
	withComputedFields["key_schema"] = []interface{}{map[string]interface{}{"attribute_name": "gsi1pk", "key_type": "HASH"}}
	withComputedFields["warm_throughput"] = []interface{}{map[string]interface{}{"read_units_per_second": 12000, "write_units_per_second": 4000}}

	type args struct {
		old map[string]interface{}
		new map[string]interface{}
	}
	type want struct {
		equal bool
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"Unchanged": {
			reason: "Identical elements must hash equal.",
			args:   args{old: gsi(5, 5, "a", "b"), new: gsi(5, 5, "b", "a")},
			want:   want{equal: true},
		},
		"ComputedFieldsIgnored": {
			reason: "key_schema and warm_throughput populated by AWS must not change the hash.",
			args:   args{old: withComputedFields, new: gsi(5, 5, "a", "b")},
			want:   want{equal: true},
		},
		"ReadCapacityChanged": {
			reason: "A read_capacity change must change the hash, or the diff that applies it is suppressed.",
			args:   args{old: gsi(5, 5, "a", "b"), new: gsi(10, 5, "a", "b")},
			want:   want{equal: false},
		},
		"WriteCapacityChanged": {
			reason: "A write_capacity change must change the hash, or the diff that applies it is suppressed.",
			args:   args{old: gsi(5, 5, "a", "b"), new: gsi(5, 10, "a", "b")},
			want:   want{equal: false},
		},
		"NonKeyAttributesChanged": {
			reason: "A non_key_attributes change must change the hash.",
			args:   args{old: gsi(5, 5, "a", "b"), new: gsi(5, 5, "a", "c")},
			want:   want{equal: false},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := want{equal: gsiUserFieldsHash(tc.args.old) == gsiUserFieldsHash(tc.args.new)}
			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(want{})); diff != "" {
				t.Errorf("%s\ngsiUserFieldsHash(old) == gsiUserFieldsHash(new): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
