/*
Copyright 2021 Upbound Inc.
*/

package route53

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the route53 group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53_traffic_policy_instance", func(r *config.Resource) {
		r.References["hosted_zone_id"] = config.Reference{
			Type: "Zone",
		}
		r.References["traffic_policy_id"] = config.Reference{
			Type: "TrafficPolicy",
		}
	})
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
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: "github.com/upbound/provider-aws/apis/kms/v1beta1.KMSKeyARN()",
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
		delete(r.References, "alias.name")
		delete(r.References, "alias.zone_id")
	})
	p.AddResourceConfigurator("aws_route53_vpc_association_authorization", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_zone", func(r *config.Resource) {
		r.References["delegation_set_id"] = config.Reference{
			Type: "DelegationSet",
		}
	})
}
