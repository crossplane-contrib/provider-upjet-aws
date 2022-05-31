/*
Copyright 2021 Upbound Inc.
*/

package kms

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for kms group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kms_key", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
	})
}
