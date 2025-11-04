// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dsql

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the dsql group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_dsql_cluster", func(r *config.Resource) {
		r.ShortGroup = "dsql"
		r.Kind = "Cluster"
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_dsql_cluster_peering", func(r *config.Resource) {
		r.ShortGroup = "dsql"
		r.Kind = "ClusterPeering"
		r.UseAsync = true
		r.References["clusters"] = config.Reference{
			TerraformName: "aws_dsql_cluster",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath(\"arn\",true)",
		}
	})
}
