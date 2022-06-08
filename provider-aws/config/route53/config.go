/*
Copyright 2021 Upbound Inc.
*/

package route53

import (
	"strings"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure route53 resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_route53_delegation_set", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("aws_route53_health_check", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("aws_route53_hosted_zone_dnssec", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["hosted_zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_key_signing_key", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["hosted_zone_id"] = config.Reference{
			Type: "Zone",
		}
		r.References["key_management_service_arn"] = config.Reference{
			Type:      "github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key",
			Extractor: "github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.KMSKeyARN()",
		}
	})
	p.AddResourceConfigurator("aws_route53_query_log", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["hosted_zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_record", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
		r.References["health_check_id"] = config.Reference{
			Type: "HealthCheck",
		}
	})
	p.AddResourceConfigurator("aws_route53_vpc_association_authorization", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
			words := strings.Split(externalName, ":")
			if len(words) != 2 {
				return
			}
			base["zone_id"] = words[0]
			base["vpc_id"] = words[1]
		}
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("aws_route53_zone", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"delegation_set_id": config.Reference{
				Type: "DelegationSet",
			},
			"vpc.vpc_id": config.Reference{
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPC",
				RefFieldName:      "VpcIdRef",
				SelectorFieldName: "VpcIdSelector",
			},
		}
	})
	p.AddResourceConfigurator("aws_route53_zone_association", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		// Z123456ABCDEFG:vpc-12345678
		// Z123456ABCDEFG:vpc-12345678:us-east-2
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
			words := strings.Split(externalName, ":")
			if len(words) >= 2 {
				base["zone_id"] = words[0]
				base["vpc_id"] = words[1]
			}
			if len(words) == 3 {
				base["vpc_region"] = words[2]
			}
		}
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
}
