// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package eks

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the eks group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_eks_cluster", func(r *config.Resource) {
		r.References = config.References{
			"role_arn": {
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
			"vpc_config.subnet_ids": {
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetIDRefs",
				SelectorFieldName: "SubnetIDSelector",
			},
			"vpc_config.security_group_ids": {
				TerraformName:     "aws_security_group",
				RefFieldName:      "SecurityGroupIDRefs",
				SelectorFieldName: "SecurityGroupIDSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_node_group", func(r *config.Resource) {
		r.References["cluster_name"] = config.Reference{
			TerraformName: "aws_eks_cluster",
			Extractor:     "ExternalNameIfClusterActive()",
		}
		r.References["node_role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["remote_access.source_security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SourceSecurityGroupIDRefs",
			SelectorFieldName: "SourceSecurityGroupIDSelector",
		}
		r.References["subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"release_version",
				"version",
			},
			ConditionalIgnoredFields: []string{
				"scaling_config",
			},
		}
		r.UseAsync = true
		r.MetaResource.ArgumentDocs["launch_template.version"] = `- (Required) EC2 Launch Template version number.`
		r.MetaResource.ArgumentDocs["subnet_ids"] = `- Identifiers of EC2 Subnets to associate with the EKS Node Group. Amazon EKS managed node groups can be launched in both public and private subnets. If you plan to deploy load balancers to a subnet, the private subnet must have tag kubernetes.io/role/internal-elb, the public subnet must have tag kubernetes.io/role/elb.`
	})
	p.AddResourceConfigurator("aws_eks_identity_provider_config", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		// OmittedFields in config.ExternalName works only for the top-level fields.
		r.References = config.References{
			"cluster_name": {
				TerraformName: "aws_eks_cluster",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_eks_fargate_profile", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				TerraformName: "aws_eks_cluster",
			},
			"pod_execution_role_arn": {
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
			"subnet_ids": {
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetIDRefs",
				SelectorFieldName: "SubnetIDSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_addon", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				TerraformName: "aws_eks_cluster",
			},
			"service_account_role_arn": {
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_access_policy_association", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				TerraformName: "aws_eks_cluster",
				// Use the terraform id instead of the external name because the external name is set before the cluster
				// has been created.
				Extractor: common.PathTerraformIDExtractor,
			},
			// Principal Arn can refer to either the ARN of an IAM user or an IAM role, with a strong best-practice
			// recommendation to always use roles. However, the eks Access Policy resource won't do anything unless
			// the principal arn matches a principal with an eks Access Entry defined on the same cluster. By retrieving
			// the principal arn from the Access Entry, we provide an easy means of ordered creation.
			"principal_arn": {
				TerraformName: "aws_eks_access_entry",
				Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("principal_arn",true)`,
			},
		}
	})
	p.AddResourceConfigurator("aws_eks_access_entry", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				TerraformName: "aws_eks_cluster",
				// Use the terraform id instead of the external name because the external name is set before the cluster
				// has been created.
				Extractor: common.PathTerraformIDExtractor,
			},
			"principal_arn": {
				TerraformName:     "aws_iam_role",
				Extractor:         common.PathARNExtractor,
				RefFieldName:      "PrincipalArnFromRoleRef",
				SelectorFieldName: "PrincipalArnFromRoleSelector",
			},
		}
	})
}
