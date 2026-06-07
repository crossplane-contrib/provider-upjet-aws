// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package guardduty

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the grafana group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_guardduty_malware_protection_plan", func(r *config.Resource) {
		r.AddSingletonListConversion("actions", "actions")
		r.AddSingletonListConversion("protected_resource", "protectedResource")
		r.AddSingletonListConversion("protected_resource[*].s3_bucket", "protectedResource[*].s3Bucket")
	})
}
