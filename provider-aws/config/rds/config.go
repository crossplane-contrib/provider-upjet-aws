/*
Copyright 2021 Upbound Inc.
*/

package rds

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for rds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_rds_cluster", func(r *config.Resource) {
		// Mutually exclusive with aws_rds_cluster_role_association
		config.MoveToStatus(r.TerraformResource, "iam_roles")
		r.References = config.References{
			"s3_import.bucket_name": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket",
			},
			"vpc_security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "VpcSecurityGroupIdRefs",
				SelectorFieldName: "VpcSecurityGroupIdSelector",
			},
			"restore_to_point_in_time.source_cluster_identifier": {
				Type: "Cluster",
			},
			"db_subnet_group_name": {
				Type: "SubnetGroup",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_instance", func(r *config.Resource) {
		r.References = config.References{
			"restore_to_point_in_time.source_db_instance_identifier": {
				Type: "Instance",
			},
			"s3_import.bucket_name": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket",
			},
			"kms_key_id": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
			},
			"performance_insights_kms_key_id": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
			},
			"restore_to_point_in_time.source_cluster_identifier": {
				Type: "Cluster",
			},
			"security_group_names": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupNameRefs",
				SelectorFieldName: "SecurityGroupNameSelector",
			},
			"vpc_security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "VpcSecurityGroupIdRefs",
				SelectorFieldName: "VpcSecurityGroupIdSelector",
			},
			"parameter_group_name": {
				Type: "ParameterGroup",
			},
			"db_subnet_group_name": {
				Type: "SubnetGroup",
			},
		}
		r.UseAsync = true
	})
}
