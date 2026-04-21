// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package s3vectors

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// s3vectorsNotFoundDiagnostic treats errors from stub ARN reads (fake account
// 000000000000) as "resource not found" so the reconciler proceeds to create.
func s3vectorsNotFoundDiagnostic(diags []*tfprotov6.Diagnostic) bool {
	for _, d := range diags {
		if d.Severity == tfprotov6.DiagnosticSeverityError &&
			strings.Contains(d.Detail, "No account found") {
			return true
		}
	}
	return false
}

// Configure adds configurations for the s3vectors group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_s3vectors_vector_bucket", func(r *config.Resource) {
		r.ExternalName.IsNotFoundDiagnosticFn = s3vectorsNotFoundDiagnostic
	})
	p.AddResourceConfigurator("aws_s3vectors_index", func(r *config.Resource) {
		r.ExternalName.IsNotFoundDiagnosticFn = s3vectorsNotFoundDiagnostic
		r.References["vector_bucket_name"] = config.Reference{
			TerraformName: "aws_s3vectors_vector_bucket",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("vector_bucket_name",true)`,
		}
	})
	p.AddResourceConfigurator("aws_s3vectors_vector_bucket_policy", func(r *config.Resource) {
		r.ExternalName.IsNotFoundDiagnosticFn = s3vectorsNotFoundDiagnostic
		r.References["vector_bucket_arn"] = config.Reference{
			TerraformName: "aws_s3vectors_vector_bucket",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("vector_bucket_arn",true)`,
		}
	})
}
