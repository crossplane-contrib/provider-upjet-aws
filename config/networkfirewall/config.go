/*
Copyright 2022 Upbound Inc.
*/

package networkfirewall

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for networkfirewall group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_networkfirewall_firewall_policy", func(r *config.Resource) {
		r.References = config.References{
			"firewall_policy.stateless_rule_group_reference.resource_arn": {
				Type: "RuleGroup",
			},
			"firewall_policy.stateful_rule_group_reference.resource_arn": {
				Type: "RuleGroup",
			},
		}
	})
}
