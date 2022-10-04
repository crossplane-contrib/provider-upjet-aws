/*
Copyright 2022 Upbound Inc.
*/

package acm

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for acm group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_acm_certificate_validation", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"certificate_arn": {
				Type: "Certificate",
			},
		}
		// Deletion takes a while.
		r.UseAsync = true
	})
}
