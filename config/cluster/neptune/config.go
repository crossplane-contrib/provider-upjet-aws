// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package neptune

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the neptune group
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_neptune_cluster", func(r *config.Resource) {
		r.UseAsync = true
		r.References["snapshot_identifier"] = config.Reference{
			TerraformName: "aws_neptune_cluster_snapshot",
		}
		r.References["replication_source_identifier"] = config.Reference{
			TerraformName: "aws_neptune_cluster",
		}
		r.References["neptune_subnet_group_name"] = config.Reference{
			TerraformName: "aws_neptune_subnet_group",
		}
		r.References["neptune_cluster_parameter_group_name"] = config.Reference{
			TerraformName: "aws_neptune_cluster_parameter_group",
		}
	})
	p.AddResourceConfigurator("aws_neptune_cluster_endpoint", func(r *config.Resource) {
		r.References["cluster_identifier"] = config.Reference{
			TerraformName: "aws_neptune_cluster",
		}
	})
	p.AddResourceConfigurator("aws_neptune_cluster_instance", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cluster_identifier"] = config.Reference{
			TerraformName: "aws_neptune_cluster",
		}
		r.References["neptune_parameter_group_name"] = config.Reference{
			TerraformName: "aws_neptune_parameter_group",
		}
		r.References["neptune_subnet_group_name"] = config.Reference{
			TerraformName: "aws_neptune_subnet_group",
		}
	})
	p.AddResourceConfigurator("aws_neptune_cluster_snapshot", func(r *config.Resource) {
		r.UseAsync = true
		r.References["db_cluster_identifier"] = config.Reference{
			TerraformName: "aws_neptune_cluster",
		}
	})
}
