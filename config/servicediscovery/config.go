// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package servicediscovery

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the servicediscovery group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_service_discovery_private_dns_namespace", func(r *config.Resource) {
		r.References["vpc"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC",
		}
	})
}
