/*
Copyright 2021 Upbound Inc.
*/

package eks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for eks group.
func Configure(p *config.Provider) { // nolint:gocyclo
	p.AddResourceConfigurator("aws_eks_cluster", func(r *config.Resource) {
		r.References = config.References{
			"role_arn": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
			"vpc_config.subnet_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
				RefFieldName:      "SubnetIDRefs",
				SelectorFieldName: "SubnetIDSelector",
			},
			"vpc_config.security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupIDRefs",
				SelectorFieldName: "SecurityGroupIDSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_node_group", func(r *config.Resource) {
		r.References["cluster_name"] = config.Reference{
			Type:      "Cluster",
			Extractor: "ExternalNameIfClusterActive()",
		}
		r.References["node_role_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["remote_access.source_security_group_ids"] = config.Reference{
			Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SourceSecurityGroupIDRefs",
			SelectorFieldName: "SourceSecurityGroupIDSelector",
		}
		r.References["subnet_ids"] = config.Reference{
			Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_identity_provider_config", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		// OmittedFields in config.ExternalName works only for the top-level fields.
		delete(r.TerraformResource.Schema["oidc"].Elem.(*schema.Resource).Schema, "identity_provider_config_name")
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
		}
	})

	p.AddResourceConfigurator("aws_eks_fargate_profile", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"pod_execution_role_arn": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
			"subnet_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
				RefFieldName:      "SubnetIDRefs",
				SelectorFieldName: "SubnetIDSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_addon", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"service_account_role_arn": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
		}
		r.UseAsync = true
	})
}
