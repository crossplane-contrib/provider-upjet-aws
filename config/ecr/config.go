// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package ecr

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the ecr group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecr_repository", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"encryption_configuration.kms_key": {
				Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
				Extractor: common.PathARNExtractor,
			},
		}
		// Deletion takes a while.
		r.UseAsync = true
	})
}
