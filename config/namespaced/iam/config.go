// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package iam

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// Configure adds configurations for the iam group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_iam_access_key", func(r *config.Resource) {
		r.References["user"] = config.Reference{
			TerraformName: "aws_iam_user",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["username"] = []byte(a)
			}
			if a, ok := attr["secret"].(string); ok {
				conn["password"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_iam_role", func(r *config.Resource) {
		r.MetaResource.ArgumentDocs["inline_policy"] = `Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, Crossplane will not manage any inline policies in this resource. Configuring one empty block (i.e., inline_policy {}) will cause Crossplane to remove all inline policies added out of band on apply.`
		r.MetaResource.ArgumentDocs["managed_policy_arns"] = `Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, Crossplane will ignore policy attachments to this resource. When configured, Crossplane will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., managed_policy_arns = []) will cause Crossplane to remove all managed policy attachments.`

		// Both inline and attached policies can either be specified in and managed by the Role resource, or by separate
		// RolePolicy and RolePolicyAttachment resources, so the Role should not late initialize them if they were unset
		// by the user, as that would cause reconciliation conflicts with potential future RolePolicy or
		// RolePolicyAttachment resources. See github issues #933 and #1207
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "managed_policy_arns", "inline_policy")
	})

	p.AddResourceConfigurator("aws_iam_instance_profile", func(r *config.Resource) {
		r.References["role"] = config.Reference{
			TerraformName: "aws_iam_role",
		}
	})

	p.AddResourceConfigurator("aws_iam_role_policy_attachment", func(r *config.Resource) {
		r.References["role"] = config.Reference{
			TerraformName: "aws_iam_role",
		}
		r.References["policy_arn"] = config.Reference{
			TerraformName: "aws_iam_policy",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_iam_user_policy_attachment", func(r *config.Resource) {
		r.References["user"] = config.Reference{
			TerraformName: "aws_iam_user",
		}
		r.References["policy_arn"] = config.Reference{
			TerraformName: "aws_iam_policy",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_iam_group_policy_attachment", func(r *config.Resource) {
		r.References["group"] = config.Reference{
			TerraformName: "aws_iam_group",
		}
		r.References["policy_arn"] = config.Reference{
			TerraformName: "aws_iam_policy",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_iam_user_group_membership", func(r *config.Resource) {
		r.References["user"] = config.Reference{
			TerraformName: "aws_iam_user",
		}
		r.References["groups"] = config.Reference{
			TerraformName:     "aws_iam_group",
			RefFieldName:      "GroupRefs",
			SelectorFieldName: "GroupSelector",
		}
	})

	p.AddResourceConfigurator("aws_iam_group_membership", func(r *config.Resource) {
		r.References["users"] = config.Reference{
			TerraformName:     "aws_iam_user",
			RefFieldName:      "UserRefs",
			SelectorFieldName: "UserSelector",
		}
		r.References["group"] = config.Reference{
			TerraformName: "aws_iam_group",
		}
	})

	p.AddResourceConfigurator("aws_iam_service_specific_credential", func(r *config.Resource) {
		r.References["user_name"] = config.Reference{
			TerraformName: "aws_iam_user",
		}
	})

	p.AddResourceConfigurator("aws_iam_user_login_profile", func(r *config.Resource) {
		r.References["user"] = config.Reference{
			TerraformName: "aws_iam_user",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"password_reset_required", "password_length", "pgp_key"},
		}
	})

	p.AddResourceConfigurator("aws_iam_user_ssh_key", func(r *config.Resource) {
		r.References["username"] = config.Reference{
			TerraformName: "aws_iam_user",
		}
	})

	p.AddResourceConfigurator("aws_iam_policy", func(r *config.Resource) {
		// Otherwise TF assigns a random string.
		config.MarkAsRequired(r.TerraformResource, "name")
	})
}
