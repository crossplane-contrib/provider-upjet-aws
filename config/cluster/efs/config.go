// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package efs

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
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

// fileSystemPolicyCustomDiff removes the policy from the diff when the observed
// and the desired policies differ only in statements whose Resource is the file
// system's own ARN. EFS stores a statement submitted without a Resource with the
// file system ARN filled in, so without this the observed policy never matches
// the desired one and the policy is updated on every reconcile.
func fileSystemPolicyCustomDiff(diff *terraform.InstanceDiff, state *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
	if diff == nil || state == nil {
		return diff, nil
	}
	policy := diff.Attributes["policy"]
	fileSystemID := state.Attributes["file_system_id"]
	if policy == nil || policy.Old == "" || policy.New == "" || fileSystemID == "" {
		return diff, nil
	}

	equivalent, err := equivalentWithoutFileSystemResource(policy.Old, policy.New, fileSystemID)
	if err != nil {
		return nil, err
	}
	if equivalent {
		delete(diff.Attributes, "policy")
	}
	return diff, nil
}

// equivalentWithoutFileSystemResource reports whether the observed and the
// desired policies are equivalent once the Resource of every statement that is
// the ARN of the given file system is removed from both.
func equivalentWithoutFileSystemResource(observed, desired, fileSystemID string) (bool, error) {
	o, err := removeFileSystemResource(observed, fileSystemID)
	if err != nil {
		return false, errors.Wrap(err, "failed to normalize the observed file system policy")
	}
	d, err := removeFileSystemResource(desired, fileSystemID)
	if err != nil {
		return false, errors.Wrap(err, "failed to normalize the desired file system policy")
	}
	equivalent, err := awspolicy.PoliciesAreEquivalent(o, d)
	return equivalent, errors.Wrap(err, "failed to compare the observed and the desired file system policies")
}

// removeFileSystemResource removes the Resource of every policy statement whose
// Resource is the ARN of the given file system.
func removeFileSystemResource(policy, fileSystemID string) (string, error) {
	var doc map[string]any
	d := json.NewDecoder(strings.NewReader(policy))
	d.UseNumber()
	if err := d.Decode(&doc); err != nil {
		return "", err
	}
	statements, ok := doc["Statement"].([]any)
	if !ok {
		statements = []any{doc["Statement"]}
	}
	for _, s := range statements {
		if statement, ok := s.(map[string]any); ok && isFileSystemARN(statement["Resource"], fileSystemID) {
			delete(statement, "Resource")
		}
	}
	out, err := json.Marshal(doc)
	return string(out), err
}

// isFileSystemARN reports whether a statement Resource, a string or a list with
// a single string, is the ARN of the given file system.
func isFileSystemARN(resource any, fileSystemID string) bool {
	if l, ok := resource.([]any); ok && len(l) == 1 {
		resource = l[0]
	}
	s, ok := resource.(string)
	if !ok {
		return false
	}
	a, err := arn.Parse(s)
	return err == nil && a.Service == "elasticfilesystem" && a.Resource == "file-system/"+fileSystemID
}
