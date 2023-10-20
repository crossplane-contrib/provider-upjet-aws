package ssoadmin

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the ssoadmin group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ssoadmin_account_assignment", func(r *config.Resource) {
		r.References["principal_id"] = config.Reference{
			TerraformName:     "aws_identitystore_group",
			Extractor:         `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("group_id",true)`,
			RefFieldName:      "PrincipalGroupRef",
			SelectorFieldName: "PrincipalGroupSelector",
		}
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
	})
	p.AddResourceConfigurator("aws_ssoadmin_permission_set_inline_policy", func(r *config.Resource) {
		r.References["instance_arn"] = config.Reference{
			TerraformName: "aws_ssoadmin_permission_set",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("instance_arn",false)`,
		}
	})
}
