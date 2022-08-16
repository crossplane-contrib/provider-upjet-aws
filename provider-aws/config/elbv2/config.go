package elbv2

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for acm group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lb", func(r *config.Resource) {
		r.LateInitializer.IgnoredFields = []string{"access_logs"}
	})
}
