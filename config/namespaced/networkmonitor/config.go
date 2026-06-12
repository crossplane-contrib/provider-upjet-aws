// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package networkmonitor

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/namespaced/common"
)

// Configure adds configurations for the networkmonitor group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_networkmonitor_probe", func(r *config.Resource) {
		r.References["monitor_name"] = config.Reference{
			TerraformName: "aws_networkmonitor_monitor",
		}
		r.References["source_arn"] = config.Reference{
			TerraformName: "aws_subnet",
			Extractor:     common.PathARNExtractor,
		}
	})
}
