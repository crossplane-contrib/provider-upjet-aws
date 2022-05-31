/*
Copyright 2021 Upbound Inc.
*/

package ecr

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for ecrs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecr_repository", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.References = map[string]config.Reference{
			"encryption_configuration.kms_key": {
				Type:      "github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key",
				Extractor: common.PathARNExtractor,
			},
		}
		// Deletion takes a while.
		r.UseAsync = true
	})
}
