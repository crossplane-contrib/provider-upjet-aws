/*
Copyright 2021 Upbound Inc.
*/

package iam

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for iam group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_iam_access_key", func(r *config.Resource) {
		r.References = config.References{
			"user": config.Reference{
				Type: "User",
			},
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

	p.AddResourceConfigurator("aws_iam_group_membership", func(r *config.Resource) {
		r.References["users"] = config.Reference{
			Type:              "User",
			RefFieldName:      "UserRefs",
			SelectorFieldName: "UserSelector",
		}
		r.References["group"] = config.Reference{
			Type: "Group",
		}
	})

	p.AddResourceConfigurator("aws_iam_service_specific_credential", func(r *config.Resource) {
		r.References["user_name"] = config.Reference{
			Type: "User",
		}
	})

	p.AddResourceConfigurator("aws_iam_user_login_profile", func(r *config.Resource) {
		r.References["user"] = config.Reference{
			Type: "User",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"password_reset_required", "password_length", "pgp_key"},
		}
	})

	p.AddResourceConfigurator("aws_iam_user_ssh_key", func(r *config.Resource) {
		r.References["username"] = config.Reference{
			Type: "User",
		}
	})

	p.AddResourceConfigurator("aws_iam_policy", func(r *config.Resource) {
		// Otherwise TF assigns a random string.
		config.MarkAsRequired(r.TerraformResource, "name")
	})
}
