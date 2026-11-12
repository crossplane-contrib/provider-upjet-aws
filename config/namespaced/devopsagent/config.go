// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package devopsagent

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the devopsagent group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_devopsagent_asset", func(r *config.Resource) {
		r.AddSingletonListConversion("content", "content")
		r.AddSingletonListConversion("content[*].file", "content[*].file")
		r.AddSingletonListConversion("content[*].file[*].body", "content[*].file[*].body")
		r.AddSingletonListConversion("content[*].source_url", "content[*].sourceUrl")
		r.AddSingletonListConversion("content[*].zip", "content[*].zip")
	})
}
