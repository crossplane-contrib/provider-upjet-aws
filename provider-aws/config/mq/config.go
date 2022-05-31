/*
Copyright 2022 Upbound Inc.
*/

package mq

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for rds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_mq_broker", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// Due to a terrajet limitation, we cannot use "metedata.name" field as the name of the resource
		// Therefore, "spec.forProvider.brokerName" field is not omitted
		// Details can be found in https://github.com/crossplane/terrajet/issues/280
		r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_mq_configuration", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// Due to a terrajet limitation, we cannot use "metedata.name" field as the name of the resource
		// Therefore, "spec.forProvider.name" field is not omitted
		// Details can be found in https://github.com/crossplane/terrajet/issues/280
		r.ExternalName = config.IdentifierFromProvider
	})

}
