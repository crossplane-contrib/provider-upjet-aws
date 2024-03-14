// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package appstream

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the appstream group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_appstream_fleet", func(r *config.Resource) {
		r.References["vpc_config.subnet_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.UseAsync = true
		r.Path = "fleet"
	})
	p.AddResourceConfigurator("aws_appstream_image_builder", func(r *config.Resource) {
		r.References["vpc_config.subnet_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.UseAsync = true
		// Otherwise getting Invalid combination of arguments: "image_name": only one of `image_arn,image_name` can be specified, but `image_arn,image_name` were specified.
		config.MoveToStatus(r.TerraformResource, "image_name")
	})
}
