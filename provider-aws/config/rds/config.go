/*
Copyright 2021 Upbound Inc.
*/

package rds

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for rds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_rds_cluster", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["cluster_identifier"] = name
			},
			OmittedFields: []string{
				"cluster_identifier",
				"cluster_identifier_prefix",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"s3_import.bucket_name": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/s3/v1alpha2.Bucket",
			},
			"vpc_security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.SecurityGroup",
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
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["identifier"] = name
			},
			OmittedFields: []string{
				"identifier",
				"identifier_prefix",
			},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
		r.References = config.References{
			"restore_to_point_in_time.source_db_instance_identifier": {
				Type: "Instance",
			},
			"s3_import.bucket_name": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/s3/v1alpha2.Bucket",
			},
			"kms_key_id": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key",
			},
			"performance_insights_kms_key_id": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key",
			},
			"restore_to_point_in_time.source_cluster_identifier": {
				Type: "Cluster",
			},
			"security_group_names": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.SecurityGroup",
				RefFieldName:      "SecurityGroupNameRefs",
				SelectorFieldName: "SecurityGroupNameSelector",
			},
			"vpc_security_group_ids": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1alpha2.SecurityGroup",
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
	p.AddResourceConfigurator("aws_db_parameter_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("aws_db_subnet_group", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
	})
}
