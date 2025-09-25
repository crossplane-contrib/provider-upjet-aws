// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ecr

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/namespaced/common"
)

// configureRepositoryCreationTemplate configures the ECR Repository Creation Template resource
func configureRepositoryCreationTemplate(r *config.Resource) {
	r.ShortGroup = "ecr"
	r.Kind = "RepositoryCreationTemplate"

	// KMS key reference for encryption configuration
	r.References = map[string]config.Reference{
		"encryption_configuration.kms_key": {
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		},
	}

	// External name is the template prefix/name, not ARN
	// Terraform import uses: terraform import aws_ecr_repository_creation_template.example template-name
	r.ExternalName = config.NameAsIdentifier

	// Repository creation templates are relatively quick to create/delete
	r.UseAsync = false
}

// Configure adds configurations for the ecr group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_ecr_repository", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"encryption_configuration.kms_key": {
				TerraformName: "aws_kms_key",
				Extractor:     common.PathARNExtractor,
			},
		}
		// Deletion takes a while.
		r.UseAsync = true
	})

	// Add Repository Creation Template configuration
	p.AddResourceConfigurator("aws_ecr_repository_creation_template", configureRepositoryCreationTemplate)
}
