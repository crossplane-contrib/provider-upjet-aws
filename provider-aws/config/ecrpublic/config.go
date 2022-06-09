/*
Copyright 2021 Upbound Inc.
*/

package ecrpublic

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for ecrpublic group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecrpublic_repository", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["repository_name"] = name
			},
			OmittedFields: []string{
				"repository_name",
			},
		}
		// Deletion takes a while.
		r.UseAsync = true
	})
}
