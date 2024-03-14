// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package fsx

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the fsx group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_fsx_windows_file_system", func(r *config.Resource) {
		r.References["kms_key_id"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
	})
}
