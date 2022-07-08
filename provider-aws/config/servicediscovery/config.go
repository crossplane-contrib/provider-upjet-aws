package servicediscovery

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for servicediscovery group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_service_discovery_private_dns_namespace", func(r *config.Resource) {
		r.References["vpc"] = config.Reference{
			Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPC",
			RefFieldName:      "VpcIdRef",
			SelectorFieldName: "VpcIdSelector",
		}
	})
}
