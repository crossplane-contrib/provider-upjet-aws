/*
Copyright 2021 Upbound Inc.
*/

package iam

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for iam group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_iam_access_key", func(r *config.Resource) {
		r.References = config.References{
			"user": config.Reference{
				Type: "User",
			},
		}
	})

	p.AddResourceConfigurator("aws_iam_role", func(r *config.Resource) {
		// Mutually exclusive with:
		// aws_iam_policy_attachment
		// aws_iam_role_policy_attachment
		// aws_iam_role_policy
		common.MutuallyExclusiveFields(r.TerraformResource, "inline_policy", "managed_policy_arns")
	})

	p.AddResourceConfigurator("aws_iam_instance_profile", func(r *config.Resource) {
		r.References = config.References{
			"role": config.Reference{
				Type: "Role",
			},
		}
	})

	p.AddResourceConfigurator("aws_iam_role_policy_attachment", func(r *config.Resource) {
		r.References = config.References{
			"role": config.Reference{
				Type: "Role",
			},
			"policy_arn": config.Reference{
				Type:      "Policy",
				Extractor: common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_iam_user_policy_attachment", func(r *config.Resource) {
		r.References = config.References{
			"user": config.Reference{
				Type: "User",
			},
			"policy_arn": config.Reference{
				Type:      "Policy",
				Extractor: common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_iam_group_policy_attachment", func(r *config.Resource) {
		r.References = config.References{
			"group": config.Reference{
				Type: "Group",
			},
			"policy_arn": config.Reference{
				Type:      "Policy",
				Extractor: common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_iam_user_group_membership", func(r *config.Resource) {
		r.References = config.References{
			"user": config.Reference{
				Type: "User",
			},
			"groups": config.Reference{
				Type:              "Group",
				RefFieldName:      "GroupRefs",
				SelectorFieldName: "GroupSelector",
			},
		}
	})
}
