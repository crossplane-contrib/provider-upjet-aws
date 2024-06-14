// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package codeartifact

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the backup group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_codeartifact_domain", func(r *config.Resource) {
		r.References["encryption_key"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
	})
}
