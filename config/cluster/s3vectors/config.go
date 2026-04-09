// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package s3vectors

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the s3vectors group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_s3vectors_vector_bucket", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("aws_s3vectors_index", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("aws_s3vectors_vector_bucket_policy", func(r *config.Resource) {
		r.References["vector_bucket_arn"] = config.Reference{
			TerraformName: "aws_s3vectors_vector_bucket",
		}
	})
}
