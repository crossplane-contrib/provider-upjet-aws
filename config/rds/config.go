/*
Copyright 2021 Upbound Inc.
*/

package rds

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/comments"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the rds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_rds_cluster", func(r *config.Resource) {
		// Mutually exclusive with aws_rds_cluster_role_association
		config.MoveToStatus(r.TerraformResource, "iam_roles")
		r.References["s3_import.bucket_name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["restore_to_point_in_time.source_cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
		r.References["db_subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["endpoint"].(string); ok {
				conn["endpoint"] = []byte(a)
			}
			if a, ok := attr["reader_endpoint"].(string); ok {
				conn["reader_endpoint"] = []byte(a)
			}
			if a, ok := attr["master_username"].(string); ok {
				conn["master_username"] = []byte(a)
			}
			if a, ok := attr["port"]; ok {
				conn["port"] = []byte(fmt.Sprintf("%v", a))
			}
			return conn, nil
		}
		r.OverrideFieldNames = map[string]string{
			"S3ImportParameters":                 "ClusterS3ImportParameters",
			"S3ImportInitParameters":             "ClusterS3ImportInitParameters",
			"S3ImportObservation":                "ClusterS3ImportObservation",
			"RestoreToPointInTimeParameters":     "ClusterRestoreToPointInTimeParameters",
			"RestoreToPointInTimeInitParameters": "ClusterRestoreToPointInTimeInitParameters",
			"RestoreToPointInTimeObservation":    "ClusterRestoreToPointInTimeObservation",
			"MasterUserSecretParameters":         "ClusterMasterUserSecretParameters",
			"MasterUserSecretInitParameters":     "ClusterMasterUserSecretInitParameters",
			"MasterUserSecretObservation":        "ClusterMasterUserSecretObservation",
		}
	})

	p.AddResourceConfigurator("aws_rds_cluster_instance", func(r *config.Resource) {
		r.References["restore_to_point_in_time.source_db_instance_identifier"] = config.Reference{
			Type: "Instance",
		}
		r.References["s3_import.bucket_name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.References["performance_insights_kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.References["restore_to_point_in_time.source_cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
		r.References["security_group_names"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SecurityGroupNameRefs",
			SelectorFieldName: "SecurityGroupNameSelector",
		}
		r.References["db_parameter_group_name"] = config.Reference{
			TerraformName: "aws_db_parameter_group",
		}
		r.References["db_subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		delete(r.References, "engine")
		delete(r.References, "engine_version")
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_db_instance", func(r *config.Resource) {
		r.References["db_subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.UseAsync = true
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name", "db_name"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["endpoint"].(string); ok {
				conn["endpoint"] = []byte(a)
			}
			if a, ok := attr["address"].(string); ok {
				conn["address"] = []byte(a)
				conn["host"] = []byte(a)
			}
			if a, ok := attr["username"].(string); ok {
				conn["username"] = []byte(a)
			}
			if a, ok := attr["port"]; ok {
				conn["port"] = []byte(fmt.Sprintf("%v", a))
			}
			if a, ok := attr["password"].(string); ok {
				conn["password"] = []byte(a)
			}
			return conn, nil
		}
		desc, _ := comments.New("If true, the password will be auto-generated and"+
			" stored in the Secret referenced by the passwordSecretRef field.",
			comments.WithTFTag("-"))
		r.TerraformResource.Schema["auto_generate_password"] = &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Description: desc.String(),
		}
		r.InitializerFns = append(r.InitializerFns,
			common.PasswordGenerator(
				"spec.forProvider.passwordSecretRef",
				"spec.forProvider.autoGeneratePassword",
			))
		r.TerraformResource.Schema["password"].Description = "Password for the " +
			"master DB user. If you set autoGeneratePassword to true, the Secret" +
			" referenced here will be created or updated with generated password" +
			" if it does not already contain one."
		r.MetaResource.ArgumentDocs["engine"] = "- (Required unless a `snapshotIdentifier` or `replicateSourceDb` is provided) The database engine to use. For supported values, see the Engine parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html). Note that for Amazon Aurora instances the engine must match the [DB Cluster](https://marketplace.upbound.io/providers/upbound/provider-aws/latest/resources/rds.aws.upbound.io/Cluster/v1beta1)'s engine'. For information on the difference between the available Aurora MySQL engines see Comparison in the [Amazon RDS Release Notes](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraMySQLReleaseNotes/Welcome.html)."
		r.MetaResource.ArgumentDocs["engine_version"] = "- (Optional) The engine version to use. If `autoMinorVersionUpgrade` is enabled, you can provide a prefix of the version such as 5.7 (for 5.7.10). The actual engine version used is returned in the attribute `status.atProvider.engineVersionActual`. For supported values, see the EngineVersion parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html). Note that for Amazon Aurora instances the engine version must match the [DB Cluster](https://marketplace.upbound.io/providers/upbound/provider-aws/latest/resources/rds.aws.upbound.io/Cluster/v1beta1)'s engine version'."
	})

	p.AddResourceConfigurator("aws_db_proxy", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_proxy_endpoint", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_rds_cluster_activity_stream", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_snapshot", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_option_group", func(r *config.Resource) {
		delete(r.References, "option.option_settings.value")
	})

	p.AddResourceConfigurator("aws_db_proxy_target", func(r *config.Resource) {
		delete(r.References, "target_group_name")
	})

	p.AddResourceConfigurator("aws_rds_cluster_endpoint", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_rds_cluster_role_association", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_snapshot_copy", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_instance_automated_backups_replication", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_cluster_snapshot", func(r *config.Resource) {
		r.UseAsync = true
	})
}
