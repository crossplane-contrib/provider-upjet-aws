// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package secretsmanager

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type replica struct {
	region   string
	kmsKeyID string
}

func replicaHash(v any) int {
	m := v.(map[string]any)
	return schema.HashString(fmt.Sprintf("%s-%s-", m["kms_key_id"], m["region"]))
}

func secretResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"replica": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      replicaHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_id":         {Type: schema.TypeString, Optional: true, Computed: true},
						"last_accessed_date": {Type: schema.TypeString, Computed: true},
						"region":             {Type: schema.TypeString, Required: true},
						"status":             {Type: schema.TypeString, Computed: true},
						"status_message":     {Type: schema.TypeString, Computed: true},
					},
				},
			},
		},
	}
}

func stateWith(replicas ...replica) *terraform.InstanceState {
	attributes := map[string]string{"replica.#": fmt.Sprint(len(replicas))}
	for _, r := range replicas {
		h := replicaHash(map[string]any{"kms_key_id": r.kmsKeyID, "region": r.region})
		prefix := fmt.Sprintf("replica.%d.", h)
		attributes[prefix+"region"] = r.region
		attributes[prefix+"kms_key_id"] = r.kmsKeyID
		attributes[prefix+"status"] = "InSync"
		attributes[prefix+"status_message"] = "Replication succeeded"
		attributes[prefix+"last_accessed_date"] = ""
	}
	return &terraform.InstanceState{ID: "arn:aws:secretsmanager:us-west-1:123456789012:secret:example", Attributes: attributes}
}

func configWith(replicas ...replica) *terraform.ResourceConfig {
	list := make([]any, 0, len(replicas))
	for _, r := range replicas {
		entry := map[string]any{"region": r.region}
		if r.kmsKeyID != "" {
			entry["kms_key_id"] = r.kmsKeyID
		}
		list = append(list, entry)
	}
	return terraform.NewResourceConfigRaw(map[string]any{"replica": list})
}

func regionsOf(set *schema.Set) []string {
	regions := make([]string, 0, set.Len())
	for _, v := range set.List() {
		regions = append(regions, v.(map[string]any)["region"].(string))
	}
	sort.Strings(regions)
	return regions
}

func TestCustomizeReplicaDiff(t *testing.T) {
	const (
		kmsA = "arn:aws:kms:us-west-2:123456789012:key/a"
		kmsB = "arn:aws:kms:us-west-2:123456789012:key/b"
		kmsC = "arn:aws:kms:eu-west-1:123456789012:key/c"
		kmsD = "arn:aws:kms:us-east-2:123456789012:key/d"
	)

	type args struct {
		state   []replica
		desired []replica
	}
	type want struct {
		// regions of the replicas in the set the SDK hands to Update
		replicas []string
		// regions of the replicas Update removes from AWS
		removed []string
		// regions of the replicas Update adds to AWS
		added []string
		// whether any change is left in the diff
		changes bool
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"SingleReplicaUnchanged": {
			reason: "A replica with its key chosen by AWS is not a change, which is the single replica case of issue 1079.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}},
				desired: []replica{{region: "us-west-2"}},
			},
			want: want{replicas: []string{"us-west-2"}},
		},
		"TwoReplicasUnchanged": {
			reason: "Replicas with their keys chosen by AWS are not a change.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}, {"eu-west-1", kmsC}},
				desired: []replica{{region: "us-west-2"}, {region: "eu-west-1"}},
			},
			want: want{replicas: []string{"eu-west-1", "us-west-2"}},
		},
		"ExplicitKeyUnchanged": {
			reason: "A replica with the explicit key it already has is not a change.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}},
				desired: []replica{{"us-west-2", kmsA}},
			},
			want: want{replicas: []string{"us-west-2"}},
		},
		"AddSecondRegion": {
			reason: "Adding a region must keep the existing replica, instead of swapping it with the new one (issue 2291).",
			args: args{
				state:   []replica{{"eu-west-1", kmsC}},
				desired: []replica{{region: "eu-west-1"}, {region: "us-west-2"}},
			},
			want: want{
				replicas: []string{"eu-west-1", "us-west-2"},
				added:    []string{"us-west-2"},
				changes:  true,
			},
		},
		"AddThirdRegion": {
			reason: "Adding a region must keep every existing replica.",
			args: args{
				state:   []replica{{"eu-west-1", kmsC}, {"us-west-2", kmsA}},
				desired: []replica{{region: "eu-west-1"}, {region: "us-west-2"}, {region: "us-east-2"}},
			},
			want: want{
				replicas: []string{"eu-west-1", "us-east-2", "us-west-2"},
				added:    []string{"us-east-2"},
				changes:  true,
			},
		},
		"RemoveOneOfTwo": {
			reason: "Removing a region must remove only that replica, and not every replica (issue 2291).",
			args: args{
				state:   []replica{{"us-east-2", kmsD}, {"us-west-2", kmsA}},
				desired: []replica{{region: "us-west-2"}},
			},
			want: want{
				replicas: []string{"us-west-2"},
				removed:  []string{"us-east-2"},
				changes:  true,
			},
		},
		"RemoveOneOfThree": {
			reason: "Removing a region must keep the other replicas.",
			args: args{
				state:   []replica{{"us-east-2", kmsD}, {"us-west-2", kmsA}, {"eu-west-1", kmsC}},
				desired: []replica{{region: "us-west-2"}, {region: "eu-west-1"}},
			},
			want: want{
				replicas: []string{"eu-west-1", "us-west-2"},
				removed:  []string{"us-east-2"},
				changes:  true,
			},
		},
		"ReplaceOneOfTwo": {
			reason: "Replacing a region at the same replica count must keep the replica that stays.",
			args: args{
				state:   []replica{{"eu-west-1", kmsC}, {"us-west-2", kmsA}},
				desired: []replica{{region: "us-west-2"}, {region: "us-east-2"}},
			},
			want: want{
				replicas: []string{"us-east-2", "us-west-2"},
				removed:  []string{"eu-west-1"},
				added:    []string{"us-east-2"},
				changes:  true,
			},
		},
		"ReplaceSingleRegion": {
			reason: "Replacing the only replica removes it and adds the new one.",
			args: args{
				state:   []replica{{"eu-west-1", kmsC}},
				desired: []replica{{region: "us-west-2"}},
			},
			want: want{
				replicas: []string{"us-west-2"},
				removed:  []string{"eu-west-1"},
				added:    []string{"us-west-2"},
				changes:  true,
			},
		},
		"AddRegionWithExplicitKey": {
			reason: "Adding a region with an explicit key must keep the replica that has the same explicit key.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}},
				desired: []replica{{"us-west-2", kmsA}, {"eu-west-1", kmsC}},
			},
			want: want{
				replicas: []string{"eu-west-1", "us-west-2"},
				added:    []string{"eu-west-1"},
				changes:  true,
			},
		},
		"AddExplicitKeyInExistingRegion": {
			reason: "A new explicit key in a region that already has a replica is a new replica.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}},
				desired: []replica{{region: "us-west-2"}, {"us-west-2", kmsB}},
			},
			want: want{
				replicas: []string{"us-west-2", "us-west-2"},
				added:    []string{"us-west-2"},
				changes:  true,
			},
		},
		"ChangeExplicitKey": {
			reason: "Changing the explicit key of a replica removes the replica with the old key and adds one with the new key.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}},
				desired: []replica{{"us-west-2", kmsB}},
			},
			want: want{
				replicas: []string{"us-west-2"},
				removed:  []string{"us-west-2"},
				added:    []string{"us-west-2"},
				changes:  true,
			},
		},
		"DuplicateEntriesUnchanged": {
			reason: "A repeated entry in the desired replicas is the same replica, so nothing changes.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}, {"eu-west-1", kmsC}},
				desired: []replica{{region: "us-west-2"}, {region: "us-west-2"}, {region: "eu-west-1"}},
			},
			want: want{replicas: []string{"eu-west-1", "us-west-2"}},
		},
		"DuplicateExplicitEntriesUnchanged": {
			reason: "A repeated entry with an explicit key is the same replica, so nothing changes.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}, {"eu-west-1", kmsC}},
				desired: []replica{{"us-west-2", kmsA}, {"us-west-2", kmsA}, {region: "eu-west-1"}},
			},
			want: want{replicas: []string{"eu-west-1", "us-west-2"}},
		},
		"DuplicateEntriesAddRegion": {
			reason: "A repeated entry must not hide a region that is added.",
			args: args{
				state:   []replica{{"eu-west-1", kmsC}},
				desired: []replica{{region: "eu-west-1"}, {region: "eu-west-1"}, {region: "us-west-2"}, {region: "us-west-2"}},
			},
			want: want{
				replicas: []string{"eu-west-1", "us-west-2"},
				added:    []string{"us-west-2"},
				changes:  true,
			},
		},
		"RemoveOneOfTwoKeysInRegion": {
			reason: "Two replicas in a region, with a single region-only replica desired, removes only one of them.",
			args: args{
				state:   []replica{{"us-west-2", kmsA}, {"us-west-2", kmsB}},
				desired: []replica{{region: "us-west-2"}},
			},
			want: want{
				replicas: []string{"us-west-2"},
				removed:  []string{"us-west-2"},
				changes:  true,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res := secretResource()
			state := stateWith(tc.args.state...)
			resourceConfig := configWith(tc.args.desired...)

			diff, err := res.Diff(context.Background(), state, resourceConfig, nil)
			if err != nil {
				t.Fatalf("%s: cannot compute the diff of the SDK: %v", tc.reason, err)
			}

			got, err := customizeReplicaDiff(res, diff, state, resourceConfig)
			if err != nil {
				t.Fatalf("%s: customizeReplicaDiff(...): %v", tc.reason, err)
			}

			g := want{changes: got != nil && !got.Empty()}
			if got != nil && !got.Empty() {
				d, err := schema.InternalMap(res.SchemaMap()).Data(state, got)
				if err != nil {
					t.Fatalf("%s: cannot read the diff: %v", tc.reason, err)
				}
				o, n := d.GetChange("replica")
				oldSet, newSet := o.(*schema.Set), n.(*schema.Set)
				g.replicas = regionsOf(newSet)
				g.removed = regionsOf(oldSet.Difference(newSet))
				g.added = regionsOf(newSet.Difference(oldSet))
			} else {
				for _, r := range tc.args.state {
					g.replicas = append(g.replicas, r.region)
				}
				sort.Strings(g.replicas)
			}

			if diff := cmp.Diff(tc.want, g, cmp.AllowUnexported(want{}), cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("%s\ncustomizeReplicaDiff(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
