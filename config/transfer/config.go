// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package transfer

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-upjet-aws/config/common"
)

// Configure adds configurations for the transfer group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_transfer_user", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "aws_transfer_server",
		}
		r.References["role"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
	})
}
