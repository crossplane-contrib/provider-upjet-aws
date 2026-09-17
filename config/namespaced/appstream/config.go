// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package appstream

import (
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-aws/v2/config/diffutils"
)

// Configure adds configurations for the appstream group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_appstream_fleet", func(r *config.Resource) {
		r.References["vpc_config.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.UseAsync = true
		r.Path = "fleet"
	})
	p.AddResourceConfigurator("aws_appstream_image_builder", func(r *config.Resource) {
		r.References["vpc_config.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.UseAsync = true
		// Otherwise getting Invalid combination of arguments: "image_name": only one of `image_arn,image_name` can be specified, but `image_arn,image_name` were specified.
		config.MoveToStatus(r.TerraformResource, "image_name")
	})

	p.AddResourceConfigurator("aws_appstream_stack", func(r *config.Resource) {
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || len(diff.Attributes) == 0 || diff.Destroy {
				return diff, nil
			}
			diff, err := diffutils.SuppressComputedOnlySetRehash(r, diff, state, "storage_connectors")
			if err != nil {
				return nil, err
			}
			if err := suppressUserSettings(r.TerraformResource.SchemaMap(), diff, state); err != nil {
				return nil, errors.Wrap(err, "cannot check diff for user_settings")
			}
			return diff, nil
		}
	})
}

// suppressUserSettings suppresses the diffs
// for defaulted user_setting elements
// AWS always return all the API-defaulted user_settings
// even when you specify a subset in the desired config,
// which causes a permanent diff.
func suppressUserSettings(sm map[string]*schema.Schema, diff *terraform.InstanceDiff, state *terraform.InstanceState) error {
	d, err := schema.InternalMap(sm).Data(state, diff)
	if err != nil {
		return errors.Wrap(err, "cannot construct diff data")
	}
	if !d.HasChange("user_settings") {
		return nil
	}
	o, n := d.GetChange("user_settings")
	observed, ok := o.(*schema.Set)
	desired, ok2 := n.(*schema.Set)
	if !ok || !ok2 {
		return nil
	}

	// Do not suppress the diff, if all settings are supplied
	// by configuration
	if observed.Len() == desired.Len() || desired.Len() == 0 {
		return nil
	}

	// If the current state is a superset of the desired
	// configuration, suppress the diff.
	if desired.Difference(observed).Len() == 0 {
		for k := range diff.Attributes {
			if strings.HasPrefix(k, "user_settings") {
				delete(diff.Attributes, k)
			}
		}
	}
	return nil
}
