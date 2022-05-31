/*
Copyright 2021 Upbound Inc.
*/

package iam

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for iam group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_iam_access_key", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"user": config.Reference{
				Type: "User",
			},
		}
	})

	p.AddResourceConfigurator("aws_iam_instance_profile", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.References = config.References{
			"role": config.Reference{
				Type: "Role",
			},
		}
	})

	p.AddResourceConfigurator("aws_iam_policy", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// TODO(muvaf): We can reconstruct ARN of this resource in GetIDFn.
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_iam_user", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
	})

	p.AddResourceConfigurator("aws_iam_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
	})

	p.AddResourceConfigurator("aws_iam_role", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
	})

	p.AddResourceConfigurator("aws_iam_role_policy_attachment", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
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
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
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
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
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
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
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

	p.AddResourceConfigurator("aws_iam_openid_connect_provider", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})

}
