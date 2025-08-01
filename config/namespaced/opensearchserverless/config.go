// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package opensearchserverless

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for the opensearchserverless group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_opensearchserverless_security_config", func(r *config.Resource) {
		r.RemoveSingletonListConversion("saml_options")
		// set the path saml_options as an embedded object to honor
		// its single nested block schema. We need to have it converted
		// into an embedded object but there's no need for
		// the Terraform conversion (it already needs to be treated
		// as an object at the Terraform layer and in the current MR API,
		// it's already an embedded object).
		r.SchemaElementOptions.SetEmbeddedObject("saml_options")
	})
	p.AddResourceConfigurator("aws_opensearchserverless_security_policy", func(r *config.Resource) {
		r.TerraformConfigurationInjector = config.CanonicalizeJSONParameters("policy")
	})
	p.AddResourceConfigurator("aws_opensearchserverless_lifecycle_policy", func(r *config.Resource) {
		r.TerraformConfigurationInjector = config.CanonicalizeJSONParameters("policy")
	})
	p.AddResourceConfigurator("aws_opensearchserverless_access_policy", func(r *config.Resource) {
		r.TerraformConfigurationInjector = config.CanonicalizeJSONParameters("policy")
	})
}
