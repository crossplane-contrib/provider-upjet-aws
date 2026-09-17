// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package fsx

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

// Configure adds configurations for the fsx group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_fsx_windows_file_system", func(r *config.Resource) {
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_fsx_ontap_file_system", func(r *config.Resource) {
		r.LateInitializer.IgnoredFields = []string{"disk_iops_configuration"}
	})

	p.AddResourceConfigurator("aws_fsx_lustre_file_system", func(r *config.Resource) {
		r.TerraformCustomDiff = common.RemoveDiffIfEmpty([]string{"final_backup_tags.%"})
	})

	p.AddResourceConfigurator("aws_fsx_ontap_volume", func(r *config.Resource) {
		r.TerraformCustomDiff = common.RemoveDiffIfEmpty([]string{"final_backup_tags.%"})
	})

	p.AddResourceConfigurator("aws_fsx_openzfs_file_system", func(r *config.Resource) {
		r.TerraformCustomDiff = common.RemoveDiffIfEmpty([]string{"final_backup_tags.%"})
	})

	p.AddResourceConfigurator("aws_fsx_windows_file_system", func(r *config.Resource) {
		r.TerraformCustomDiff = common.RemoveDiffIfEmpty([]string{"final_backup_tags.%"})
	})
}
