// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package amp

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the amp group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_prometheus_workspace", func(r *config.Resource) {
		// No special configuration needed for workspace
	})

	p.AddResourceConfigurator("aws_prometheus_alert_manager_definition", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			TerraformName: "aws_prometheus_workspace",
		}
	})

	p.AddResourceConfigurator("aws_prometheus_rule_group_namespace", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			TerraformName: "aws_prometheus_workspace",
		}
	})

	p.AddResourceConfigurator("aws_prometheus_scraper", func(r *config.Resource) {
		// Configure cross-resource references
		r.References["destination.amp.workspace_arn"] = config.Reference{
			TerraformName: "aws_prometheus_workspace",
			Extractor:     common.PathARNExtractor,
		}
		r.References["source.eks.cluster_arn"] = config.Reference{
			TerraformName: "aws_eks_cluster",
			Extractor:     common.PathARNExtractor,
		}
		r.References["source.eks.subnet_ids"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["source.eks.security_group_ids"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.References["role_configuration.source_role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
	})
}
