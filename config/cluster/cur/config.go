// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cur

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the cur group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_cur_report_definition", func(r *config.Resource) {
		r.References["s3_bucket"] = config.Reference{
			TerraformName: "aws_s3_bucket",
		}
	})
}
