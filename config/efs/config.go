// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package efs

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the efs group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_efs_mount_target", func(r *config.Resource) {
		r.UseAsync = true
		r.References["file_system_id"] = config.Reference{
			TerraformName: "aws_efs_file_system",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		/*r.MetaResource.Examples[0].Dependencies["aws_efs_file_system.foo"] = `{"creation_token": "my-product-foo", "region": "us-west-1"}`
		if err := r.MetaResource.Examples[0].Dependencies.SetPathValue("aws_subnet.alpha", "availability_zone", "us-west-1b"); err != nil {
			panic(err)
		}*/
	})
	p.AddResourceConfigurator("aws_efs_access_point", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			TerraformName: "aws_efs_file_system",
		}
		// r.MetaResource.Examples[0].Dependencies["aws_efs_file_system.foo"] = `{"creation_token": "my-product-foo", "region": "us-west-1"}`
	})
	p.AddResourceConfigurator("aws_efs_backup_policy", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			TerraformName: "aws_efs_file_system",
		}
	})
	p.AddResourceConfigurator("aws_efs_file_system_policy", func(r *config.Resource) {
		r.References["file_system_id"] = config.Reference{
			TerraformName: "aws_efs_file_system",
		}
	})

	p.AddResourceConfigurator("aws_efs_file_system", func(r *config.Resource) {
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		r.UseAsync = true

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["id"] = []byte(a)
			}
			return conn, nil
		}
	})
}
