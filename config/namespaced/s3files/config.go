// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package s3files

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
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
		// posix_user, root_directory and root_directory.creation_permissions are
		// Plugin Framework list-nested blocks backed by a single object
		// (fwtypes.ListNestedObjectValueOf). Framework schemas do not surface
		// max_items, so upjet cannot infer the singleton list on its own.
		//
		// Nested paths must keep the [*] index on every ancestor list segment.
		// The wildcards are stripped for schema generation, but the runtime
		// conversion walks the paths while the parent is still a list, so
		// omitting them fails with "root_directory: not an object".
		r.AddSingletonListConversion("posix_user", "posixUser")
		r.AddSingletonListConversion("root_directory", "rootDirectory")
		r.AddSingletonListConversion("root_directory[*].creation_permissions", "rootDirectory[*].creationPermissions")
	})
	p.AddResourceConfigurator("aws_s3files_mount_target", func(r *config.Resource) {
		r.References["file_system_id"] = fileSystemIDReference
		// security_groups has no "_ids" suffix, so it is not covered by the
		// KnownReferencers defaults.
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
