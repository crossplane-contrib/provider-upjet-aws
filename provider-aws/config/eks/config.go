/*
Copyright 2021 Upbound Inc.
*/

package eks

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/pkg/errors"

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
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
			"vpc_config.security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupIdRefs",
				SelectorFieldName: "SecurityGroupIdSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_node_group", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"node_role_arn": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role",
				Extractor: common.PathARNExtractor,
			},
			"remote_access.source_security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SourceSecurityGroupIdRefs",
				SelectorFieldName: "SourceSecurityGroupIdSelector",
			},
			"subnet_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_identity_provider_config", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		// OmittedFields works only for the top-level fields.
		delete(r.TerraformResource.Schema["oidc"].Elem.(*schema.Resource).Schema, "identity_provider_config_name")
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, externalName string) {
				if _, ok := base["oidc"]; !ok {
					base["oidc"] = map[string]interface{}{}
				}
				if m, ok := base["oidc"].(map[string]interface{}); ok {
					m["identity_provider_config_name"] = externalName
				}
			},
			GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
				if id, ok := tfstate["id"]; ok {
					return strings.Split(id.(string), ":")[1], nil
				}
				return "", errors.New("there is no id in tfstate")
			},
			GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
				cl, ok := parameters["cluster_name"]
				if !ok {
					return "", errors.New("cluster_name cannot be empty")
				}
				return fmt.Sprintf("%s:%s", cl.(string), externalName), nil
			},
		}

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
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
		}
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
	})
}
