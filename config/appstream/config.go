/*
Copyright 2022 Upbound Inc.
*/

package appstream

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for appstream group.
func Configure(p *config.Provider) { // nolint:gocyclo
	p.AddResourceConfigurator("aws_appstream_fleet", func(r *config.Resource) {
		r.References = config.References{
			"vpc_config.subnet_ids": {
				Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
				RefFieldName:      "SubnetIDRefs",
				SelectorFieldName: "SubnetIDSelector",
			},
		}
		r.UseAsync = true
	})
}
