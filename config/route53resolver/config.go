/*
Copyright 2022 Upbound Inc.
*/

package route53resolver

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for the route53resolver group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53_resolver_query_log_config", func(r *config.Resource) {
		delete(r.References, "destination_arn")
	})
}
