// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package backup

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the backup group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_backup_vault", func(r *config.Resource) {
		r.References["kms_key_arn"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_backup_selection", func(r *config.Resource) {
		r.References["iam_role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["plan_id"] = config.Reference{
			TerraformName: "aws_backup_plan",
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			for key, attrDiff := range diff.Attributes {
				if strings.HasPrefix(key, "condition.") && strings.HasSuffix(key, ".#") {
					if attrDiff.Old == "0" && attrDiff.New == "0" && !attrDiff.NewComputed {
						delete(diff.Attributes, key)
					}
				}
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("aws_backup_vault_notifications", func(r *config.Resource) {
		r.References["sns_topic_arn"] = config.Reference{
			TerraformName: "aws_sns_topic",
			Extractor:     common.PathARNExtractor,
		}
		r.References["backup_vault_name"] = config.Reference{
			TerraformName: "aws_backup_vault",
		}
	})
	p.AddResourceConfigurator("aws_backup_vault_lock_configuration", func(r *config.Resource) {
		r.References["backup_vault_name"] = config.Reference{
			TerraformName: "aws_backup_vault",
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
