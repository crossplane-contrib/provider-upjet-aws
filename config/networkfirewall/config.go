/*
Copyright 2022 Upbound Inc.
*/

package networkfirewall

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the networkfirewall group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_networkfirewall_firewall_policy", func(r *config.Resource) {
		r.References = config.References{
			"firewall_policy.stateless_rule_group_reference.resource_arn": {
				TerraformName: "aws_networkfirewall_rule_group",
				Extractor:     common.PathARNExtractor,
			},
			"firewall_policy.stateful_rule_group_reference.resource_arn": {
				TerraformName: "aws_networkfirewall_rule_group",
				Extractor:     common.PathARNExtractor,
			},
		}
	})

	p.AddResourceConfigurator("aws_networkfirewall_firewall", func(r *config.Resource) {
		r.UseAsync = true
	})
}
