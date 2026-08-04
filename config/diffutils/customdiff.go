// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package diffutils

import (
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type awsDBParameter struct {
	Value       string
	ApplyMethod string
	Hash        string
}

func flattenDBParameterSet(tfList []any) map[string]awsDBParameter {
	dbParams := make(map[string]awsDBParameter, len(tfList))
	for _, tfMapRaw := range tfList {
		tfMap, ok := tfMapRaw.(map[string]any)
		if !ok {
			continue
		}
		name, ok := tfMap["name"].(string)
		if !ok || name == "" {
			continue
		}
		dbParam := awsDBParameter{}
		// Parameter values are case-sensitive, hence stored verbatim.
		if v, ok := tfMap["value"].(string); ok {
			dbParam.Value = v
		}
		if v, ok := tfMap["apply_method"].(string); ok && v != "" {
			dbParam.ApplyMethod = strings.ToLower(v)
		}
		dbParams[name] = dbParam
	}

	return dbParams
}

// SuppressAWSParameterGroupDiff suppresses diff for AWS DB parameter group diff when the diff is
// only on `ApplyMethod` causing perpetual diff.
// See https://github.com/hashicorp/terraform-provider-aws/blob/v6.55.0/website/docs/r/db_parameter_group.html.markdown#problematic-plan-changes
func SuppressAWSParameterGroupDiff(r *config.Resource, diff *terraform.InstanceDiff, s *terraform.InstanceState) (*terraform.InstanceDiff, error) { //nolint:gocyclo // easier to follow as a unit
	if isNonUpdateDiff(diff, s) {
		return diff, nil
	}
	d, err := schema.InternalMap(r.TerraformResource.Schema).Data(s, diff)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot construct diff data")
	}

	if d.HasChange("parameter") {
		o, n := d.GetChange("parameter")
		os, ns := o.(*schema.Set), n.(*schema.Set)

		// no diff suppression when parameters are added or removed
		if os.Len() != ns.Len() {
			return diff, nil
		}

		oldParams := flattenDBParameterSet(os.List())
		newParams := flattenDBParameterSet(ns.List())

		// check if flattening changes the parameter count, meaning
		// there are duplicate parameters with same name and different
		// values/apply-methods. This can be considered as an invalid
		// configuration.
		if len(oldParams) != os.Len() || len(newParams) != ns.Len() {
			// no diff suppression, let other layers reject if needed
			return diff, nil
		}
		for name, newVal := range newParams {
			oldVal, ok := oldParams[name]
			if !ok || oldVal.Value != newVal.Value {
				// Desired config introduced a new parameter,
				// or value changed for an existing parameter
				return diff, nil
			}
		}

		// At this point, we validated that the diff for all parameters
		// in `parameter` attribute is apply-method only.
		// Suppress this diff.
		for k := range diff.Attributes {
			if strings.HasPrefix(k, "parameter.") {
				delete(diff.Attributes, k)
			}
		}
	}
	return diff, nil
}

// isNonUpdateDiff is a utility for checking a TF diff & state pair is a non-update,
// e.g. diff is empty OR is a destroy diff OR creation diff
func isNonUpdateDiff(diff *terraform.InstanceDiff, s *terraform.InstanceState) bool {
	return diff == nil || len(diff.Attributes) == 0 || diff.Destroy || s == nil || len(s.Attributes) == 0
}
