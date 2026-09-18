// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package efs

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	awspolicy "github.com/hashicorp/awspolicyequivalence"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the efs group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_efs_mount_target", func(r *config.Resource) {
		r.UseAsync = true
		r.References["file_system_id"] = config.Reference{
			TerraformName: "aws_efs_file_system",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		/*r.MetaResource.Examples[0].Dependencies["aws_efs_file_system.foo"] = `{"creation_token": "my-product-foo", "region": "us-west-1"}`
		if err := r.MetaResource.Examples[0].Dependencies.SetPathValue("aws_subnet.alpha", "availability_zone", "us-west-1b"); err != nil {
			panic(err)
		}*/
	})
	p.AddResourceConfigurator("aws_efs_access_point", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			TerraformName: "aws_efs_file_system",
		}
		// r.MetaResource.Examples[0].Dependencies["aws_efs_file_system.foo"] = `{"creation_token": "my-product-foo", "region": "us-west-1"}`
	})
	p.AddResourceConfigurator("aws_efs_backup_policy", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			TerraformName: "aws_efs_file_system",
		}
	})
	p.AddResourceConfigurator("aws_efs_file_system_policy", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			TerraformName: "aws_efs_file_system",
		}
		r.TerraformCustomDiff = fileSystemPolicyCustomDiff
	})

	p.AddResourceConfigurator("aws_efs_file_system", func(r *config.Resource) {
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		r.UseAsync = true

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["id"] = []byte(a)
			}
			return conn, nil
		}
	})
}

// fileSystemPolicyCustomDiff suppresses spurious "policy" diffs for
// aws_efs_file_system_policy. AWS can return the policy document with a
// different but semantically equivalent JSON representation (e.g. reordered
// statements) than what was submitted. Without this, every observe sees a
// diff, triggering an update on every reconcile and exhausting the EFS API
// rate limit.
func fileSystemPolicyCustomDiff(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
	if diff == nil || diff.Attributes["policy"] == nil || diff.Attributes["policy"].Old == "" || diff.Attributes["policy"].New == "" {
		return diff, nil
	}

	vOld, err := common.RemovePolicyVersion(diff.Attributes["policy"].Old)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove Version from the old AWS policy document")
	}
	vNew, err := common.RemovePolicyVersion(diff.Attributes["policy"].New)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove Version from the new AWS policy document")
	}

	ok, err := awspolicy.PoliciesAreEquivalent(vOld, vNew)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compare the old and the new AWS policy documents")
	}
	if ok {
		delete(diff.Attributes, "policy")
	}
	return diff, nil
}
