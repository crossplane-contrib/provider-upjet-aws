/*
Copyright 2022 Upbound Inc.
*/

package neptune

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for neptune group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_neptune_cluster", func(r *config.Resource) {
		r.UseAsync = true
		r.References["snapshot_identifier"] = config.Reference{
			Type: "ClusterSnapshot",
		}
		r.References["replication_source_identifier"] = config.Reference{
			Type: "Cluster",
		}
		r.References["neptune_subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		r.References["neptune_cluster_parameter_group_name"] = config.Reference{
			Type: "ClusterParameterGroup",
		}
	})
	p.AddResourceConfigurator("aws_neptune_cluster_endpoint", func(r *config.Resource) {
		r.References["cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
	})
	p.AddResourceConfigurator("aws_neptune_cluster_instance", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
		r.References["neptune_parameter_group_name"] = config.Reference{
			Type: "ParameterGroup",
		}
		r.References["neptune_subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
	})
	p.AddResourceConfigurator("aws_neptune_cluster_snapshot", func(r *config.Resource) {
		r.UseAsync = true
		r.References["db_cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
	})
}
