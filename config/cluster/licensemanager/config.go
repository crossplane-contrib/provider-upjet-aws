// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package licensemanager

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/config/cluster/common"
)

// Configure adds configurations for the licensemanager group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_licensemanager_association", func(r *config.Resource) {
		r.References["license_configuration_arn"] = config.Reference{
			TerraformName: "aws_licensemanager_license_configuration",
			Extractor:     common.PathARNExtractor,
		}
	})
}
