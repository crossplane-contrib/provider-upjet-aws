/*
Copyright 2021 Upbound Inc.
*/

package eks

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for eks group.
func Configure(p *config.Provider) { // nolint:gocyclo
	p.AddResourceConfigurator("aws_eks_cluster", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.References = config.References{
			"role_arn": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1alpha2.Role",
				Extractor: common.PathARNExtractor,
			},
			"vpc_config.subnet_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.Subnet",
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
			"vpc_config.security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.SecurityGroup",
				RefFieldName:      "SecurityGroupIdRefs",
				SelectorFieldName: "SecurityGroupIdSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_node_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["node_group_name"] = name
			},
			OmittedFields: []string{
				"node_group_name",
				"node_group_name_prefix",
			},
			GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
				cl, ok := parameters["cluster_name"].(string)
				if !ok || cl == "" {
					return "", errors.New("cannot get cluster_name from parameters")
				}
				return fmt.Sprintf("%s:%s", cl, externalName), nil
			},
			GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
				id, ok := tfstate["id"].(string)
				if !ok || id == "" {
					return "", errors.New("cannot get id from tfstate")
				}
				// my_cluster:my_node_group
				w := strings.Split(id, ":")
				if len(w) != 2 {
					return "", errors.New("format of id should be my_cluster:my_node_group")
				}
				return w[len(w)-1], nil
			},
		}
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"node_role_arn": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1alpha2.Role",
				Extractor: common.PathARNExtractor,
			},
			"remote_access.source_security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.SecurityGroup",
				RefFieldName:      "SourceSecurityGroupIdRefs",
				SelectorFieldName: "SourceSecurityGroupIdSelector",
			},
			"subnet_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.Subnet",
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_identity_provider_config", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
		}
	})

	p.AddResourceConfigurator("aws_eks_fargate_profile", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["fargate_profile_name"] = name
			},
			OmittedFields: []string{
				"fargate_profile_name",
			},
			GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
				cl, ok := parameters["cluster_name"].(string)
				if !ok || cl == "" {
					return "", errors.New("cannot get cluster_name from parameters")
				}
				return fmt.Sprintf("%s:%s", cl, externalName), nil
			},
			GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
				id, ok := tfstate["id"].(string)
				if !ok || id == "" {
					return "", errors.New("cannot get id from tfstate")
				}
				// my_cluster:my_fargate_profile
				w := strings.Split(id, ":")
				if len(w) != 2 {
					return "", errors.New("format of id should be my_cluster:my_fargate_profile")
				}
				return w[len(w)-1], nil
			},
		}

		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"pod_execution_role_arn": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1alpha2.Role",
				Extractor: common.PathARNExtractor,
			},
			"subnet_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.Subnet",
				RefFieldName:      "SubnetIdRefs",
				SelectorFieldName: "SubnetIdSelector",
			},
		}
	})
	p.AddResourceConfigurator("aws_eks_addon", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, externalName string) {
				base["addon_name"] = externalName
			},
			OmittedFields: []string{
				"addon_name",
			},
			GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
				cl, ok := parameters["cluster_name"].(string)
				if !ok || cl == "" {
					return "", errors.New("cannot get cluster_name from parameters")
				}
				return fmt.Sprintf("%s:%s", cl, externalName), nil
			},
			GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
				id, ok := tfstate["id"].(string)
				if !ok || id == "" {
					return "", errors.New("cannot get id from tfstate")
				}
				// my_cluster:my_eks_addon
				w := strings.Split(id, ":")
				if len(w) != 2 {
					return "", errors.New("format of id should be my_cluster:my_eks_addon")
				}
				return w[len(w)-1], nil
			},
		}
		r.References = config.References{
			"cluster_name": {
				Type: "Cluster",
			},
			"service_account_role_arn": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/iam/v1alpha2.Role",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
