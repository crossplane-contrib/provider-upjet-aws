// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ecr

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

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
}
