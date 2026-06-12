// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package timestreaminfluxdb

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the timestreaminfluxdb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_timestreaminfluxdb_db_instance", func(r *config.Resource) {
		r.References["vpc_subnet_ids"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["vpc_security_group_ids"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.AddSingletonListConversion("log_delivery_configuration", "logDeliveryConfiguration")
		r.AddSingletonListConversion("log_delivery_configuration[*].s3_configuration", "logDeliveryConfiguration[*].s3Configuration")
	})
	p.AddResourceConfigurator("aws_timestreaminfluxdb_db_cluster", func(r *config.Resource) {
		r.References["vpc_subnet_ids"] = config.Reference{
			TerraformName: "aws_subnet",
		}
		r.References["vpc_security_group_ids"] = config.Reference{
			TerraformName: "aws_security_group",
		}
		r.AddSingletonListConversion("log_delivery_configuration", "logDeliveryConfiguration")
		r.AddSingletonListConversion("log_delivery_configuration[*].s3_configuration", "logDeliveryConfiguration[*].s3Configuration")
	})
}
