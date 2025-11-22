// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package eks

import (
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"

	"github.com/upbound/provider-aws/apis/cluster/eks/v1beta1"
	"github.com/upbound/provider-aws/apis/cluster/eks/v1beta2"
	"github.com/upbound/provider-aws/config/cluster/common"
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
		r.Conversions = append(r.Conversions,
			conversion.NewCustomConverter("v1beta1", "v1beta2", clusterConverterFromv1beta1Tov1beta2),
			conversion.NewCustomConverter("v1beta2", "v1beta1", clusterConverterFromv1beta2Tov1beta1),
		)
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

func clusterConverterFromv1beta1Tov1beta2(src, target xpresource.Managed) error {
	srcTyped := src.(*v1beta1.Cluster)
	targetTyped := target.(*v1beta2.Cluster)

	if len(srcTyped.Spec.ForProvider.UpgradePolicy) > 0 {
		if targetTyped.Spec.ForProvider.UpgradePolicy == nil {
			targetTyped.Spec.ForProvider.UpgradePolicy = &v1beta2.UpgradePolicyParameters{}
		}
		if srcTyped.Spec.ForProvider.UpgradePolicy[0].SupportType != nil {
			targetTyped.Spec.ForProvider.UpgradePolicy.SupportType = srcTyped.Spec.ForProvider.UpgradePolicy[0].SupportType
		}
	}
	if len(srcTyped.Spec.InitProvider.UpgradePolicy) > 0 {
		if targetTyped.Spec.InitProvider.UpgradePolicy == nil {
			targetTyped.Spec.InitProvider.UpgradePolicy = &v1beta2.UpgradePolicyInitParameters{}
		}
		if srcTyped.Spec.InitProvider.UpgradePolicy[0].SupportType != nil {
			targetTyped.Spec.InitProvider.UpgradePolicy.SupportType = srcTyped.Spec.InitProvider.UpgradePolicy[0].SupportType
		}
	}
	if len(srcTyped.Status.AtProvider.UpgradePolicy) > 0 {
		if targetTyped.Status.AtProvider.UpgradePolicy == nil {
			targetTyped.Status.AtProvider.UpgradePolicy = &v1beta2.UpgradePolicyObservation{}
		}
		if srcTyped.Status.AtProvider.UpgradePolicy[0].SupportType != nil {
			targetTyped.Status.AtProvider.UpgradePolicy.SupportType = srcTyped.Status.AtProvider.UpgradePolicy[0].SupportType
		}
	}
	return nil
}

func clusterConverterFromv1beta2Tov1beta1(src, target xpresource.Managed) error {
	srcTyped := src.(*v1beta2.Cluster)
	targetTyped := target.(*v1beta1.Cluster)

	if srcTyped.Spec.ForProvider.UpgradePolicy != nil {
		if len(targetTyped.Spec.ForProvider.UpgradePolicy) == 0 {
			targetTyped.Spec.ForProvider.UpgradePolicy = []v1beta1.UpgradePolicyParameters{{}}
		}
		if srcTyped.Spec.ForProvider.UpgradePolicy.SupportType != nil {
			targetTyped.Spec.ForProvider.UpgradePolicy[0].SupportType = srcTyped.Spec.ForProvider.UpgradePolicy.SupportType
		}
	}
	if srcTyped.Spec.InitProvider.UpgradePolicy != nil {
		if len(targetTyped.Spec.InitProvider.UpgradePolicy) == 0 {
			targetTyped.Spec.InitProvider.UpgradePolicy = []v1beta1.UpgradePolicyInitParameters{{}}
		}
		if srcTyped.Spec.InitProvider.UpgradePolicy.SupportType != nil {
			targetTyped.Spec.InitProvider.UpgradePolicy[0].SupportType = srcTyped.Spec.InitProvider.UpgradePolicy.SupportType
		}
	}
	if srcTyped.Status.AtProvider.UpgradePolicy != nil {
		if len(targetTyped.Status.AtProvider.UpgradePolicy) == 0 {
			targetTyped.Status.AtProvider.UpgradePolicy = []v1beta1.UpgradePolicyObservation{{}}
		}
		if srcTyped.Status.AtProvider.UpgradePolicy.SupportType != nil {
			targetTyped.Status.AtProvider.UpgradePolicy[0].SupportType = srcTyped.Status.AtProvider.UpgradePolicy.SupportType
		}
	}
	return nil
}
