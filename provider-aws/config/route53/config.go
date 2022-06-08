/*
Copyright 2021 Upbound Inc.
*/

package route53

import (
	"github.com/upbound/official-providers/provider-aws/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure route53 resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53_hosted_zone_dnssec", func(r *config.Resource) {
		r.References["hosted_zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_key_signing_key", func(r *config.Resource) {
		r.References["hosted_zone_id"] = config.Reference{
			Type: "Zone",
		}
		r.References["key_management_service_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key",
			Extractor: "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.KMSKeyARN()",
		}
	})
	p.AddResourceConfigurator("aws_route53_query_log", func(r *config.Resource) {
		r.References["hosted_zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_record", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
		r.References["health_check_id"] = config.Reference{
			Type: "HealthCheck",
		}
	})
	p.AddResourceConfigurator("aws_route53_vpc_association_authorization", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_zone", func(r *config.Resource) {
		// Mutually exclusive with aws_route53_zone_association
		common.MutuallyExclusiveFields(r.TerraformResource, "vpc")
		r.References["delegation_set_id"] = config.Reference{
			Type: "DelegationSet",
		}
	})
	p.AddResourceConfigurator("aws_route53_zone_association", func(r *config.Resource) {
		// Mutually exclusive with existing region field.
		common.MutuallyExclusiveFields(r.TerraformResource, "vpc_region")
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
}
