package backup

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the backup group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_backup_vault", func(r *config.Resource) {
		r.References["kms_key_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_backup_selection", func(r *config.Resource) {
		r.References["iam_role_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/iam/v1beta1.Role",
			Extractor: common.PathARNExtractor,
		}
		r.References["plan_id"] = config.Reference{
			Type: "Plan",
		}
	})

	p.AddResourceConfigurator("aws_backup_vault_notifications", func(r *config.Resource) {
		r.References["sns_topic_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/sns/v1beta1.Topic",
			Extractor: common.PathARNExtractor,
		}
		r.References["backup_vault_name"] = config.Reference{
			Type: "Vault",
		}
	})
	p.AddResourceConfigurator("aws_backup_vault_lock_configuration", func(r *config.Resource) {
		r.References["backup_vault_name"] = config.Reference{
			Type: "Vault",
		}
	})

	p.AddResourceConfigurator("aws_backup_framework", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_backup_plan", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_backup_region_settings", func(r *config.Resource) {
		r.TerraformResource.Schema["resource_type_management_preference"].Description += "\nWARNING: All parameters are required to be given: EFS, DynamoDB"
		r.TerraformResource.Schema["resource_type_opt_in_preference"].Description += "\nWARNING: All parameters are required to be given: " +
			"EFS, DynamoDB, EBS, EC2, FSx, S3, Aurora, RDS, Storage Gateway, VirtualMachine"
	})
}
