// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package grafana

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the grafana group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_grafana_workspace", func(r *config.Resource) {
		r.References["role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_grafana_role_association", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			TerraformName: "aws_grafana_workspace",
		}
	})

	p.AddResourceConfigurator("aws_grafana_workspace_saml_configuration", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			TerraformName: "aws_grafana_workspace",
		}
	})

	p.AddResourceConfigurator("aws_grafana_license_association", func(r *config.Resource) {
		r.UseAsync = true
	})
}
