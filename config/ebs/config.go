/*
Copyright 2021 Upbound Inc.
*/

package ebs

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for ebs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ebs_volume", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"kms_key_id": {
				Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			},
		}
	})
}
