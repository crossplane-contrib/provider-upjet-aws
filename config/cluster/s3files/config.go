// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package s3files

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// fileSystemIDReference is the reference configuration for the file_system_id
// field, shared by every resource that hangs off a FileSystem.
var fileSystemIDReference = config.Reference{
	TerraformName: "aws_s3files_file_system",
}

// Configure adds configurations for the s3files group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_s3files_file_system", func(r *config.Resource) {
		// bucket takes the S3 bucket ARN, not its name.
		r.References["bucket"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     common.PathARNExtractor,
		}
	})
	p.AddResourceConfigurator("aws_s3files_access_point", func(r *config.Resource) {
		r.References["file_system_id"] = fileSystemIDReference
		r.AddSingletonListConversion("posix_user", "posixUser")
		r.AddSingletonListConversion("root_directory", "rootDirectory")
		r.AddSingletonListConversion("root_directory[*].creation_permissions", "rootDirectory[*].creationPermissions")
	})
	p.AddResourceConfigurator("aws_s3files_mount_target", func(r *config.Resource) {
		r.References["file_system_id"] = fileSystemIDReference
		r.References["security_groups"] = config.Reference{
			TerraformName: "aws_security_group",
		}
	})
	p.AddResourceConfigurator("aws_s3files_file_system_policy", func(r *config.Resource) {
		r.References["file_system_id"] = fileSystemIDReference
	})
	p.AddResourceConfigurator("aws_s3files_synchronization_configuration", func(r *config.Resource) {
		r.References["file_system_id"] = fileSystemIDReference
	})
}
