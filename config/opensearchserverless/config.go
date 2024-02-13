/*
Copyright 2024 Upbound Inc.
*/

package opensearchserverless

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for the opensearchserverless group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_opensearchserverless_security_config", func(r *config.Resource) {
		r.SchemaElementOptions.SetEmbeddedObject("saml_options")
	})
}
