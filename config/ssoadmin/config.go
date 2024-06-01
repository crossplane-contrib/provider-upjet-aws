// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ssoadmin

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-upjet-aws/config/common"
)

// Configure adds configurations for the ssoadmin group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ssoadmin_account_assignment", func(r *config.Resource) {
		r.References["principal_id"] = config.Reference{
			TerraformName:     "aws_identitystore_group",
			RefFieldName:      "PrincipalIDFromGroupRef",
			SelectorFieldName: "PrincipalIDFromGroupSelector",
		}
		r.MetaResource.ArgumentDocs["principal_id"] = "- (Required) An identifier for an object in SSO, such as a " +
			"user or group. PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). This can be " +
			"set to the crossplane external-name of either a Group or User in the identitystore api group, but the " +
			"Ref and Selector fields will only work with a Group."
		r.References["permission_set_arn"] = config.Reference{
			TerraformName: "aws_ssoadmin_permission_set",
			Extractor:     common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_ssoadmin_customer_managed_policy_attachment", func(r *config.Resource) {
		r.References["customer_managed_policy_reference.name"] = config.Reference{
			TerraformName:     "aws_iam_policy",
			RefFieldName:      "PolicyNameRef",
			SelectorFieldName: "PolicyNameSelector",
		}
		//
		delete(r.References, "instance_arn")
	})
	p.AddResourceConfigurator("aws_ssoadmin_permission_set_inline_policy", func(r *config.Resource) {
		delete(r.References, "instance_arn")
	})
	p.AddResourceConfigurator("aws_ssoadmin_permissions_boundary_attachment", func(r *config.Resource) {
		delete(r.References, "instance_arn")
	})

}
