// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package codeartifact

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the codeartifact group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_codeartifact_domain", func(r *config.Resource) {
		r.References["encryption_key"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_codeartifact_domain_permissions_policy", func(r *config.Resource) {
		r.References["domain"] = config.Reference{
			TerraformName: "aws_codeartifact_domain",
			Extractor:     "ExtractDomainName()",
		}
	})

	p.AddResourceConfigurator("aws_codeartifact_repository", func(r *config.Resource) {
		r.References["domain"] = config.Reference{
			TerraformName: "aws_codeartifact_domain",
			Extractor:     "ExtractDomainName()",
		}
	})

	p.AddResourceConfigurator("aws_codeartifact_repository_permissions_policy", func(r *config.Resource) {
		r.References["domain"] = config.Reference{
			TerraformName: "aws_codeartifact_domain",
			Extractor:     "ExtractDomainName()",
		}

		r.References["repository"] = config.Reference{
			TerraformName: "aws_codeartifact_repository",
			Extractor:     "ExtractRepositoryName()",
		}
	})
}
