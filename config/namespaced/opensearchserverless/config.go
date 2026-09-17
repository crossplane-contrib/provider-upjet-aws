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
		// In upstream TF AWS provider 5.x, this was a single-nested block.
		// In XP side, this was handled with an embedded object in CRD with
		// runtime TF conversions removed (TF layer expected an object)
		// At 6.x, https://github.com/hashicorp/terraform-provider-aws/pull/42270
		// switched this to list-nested block, with max 1 elements.
		// Now handle this with a regular singleton-list conversion.
		// No CRD API change: still an embedded object in CRD,
		// with runtime TF conversions to list with 1 element.
		r.AddSingletonListConversion("saml_options", "samlOptions")
		r.AddSingletonListConversion("iam_federation_options", "iamFederationOptions")
		r.AddSingletonListConversion("iam_identity_center_options", "iamIdentityCenterOptions")
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
	p.AddResourceConfigurator("aws_opensearchserverless_collection", func(r *config.Resource) {
		r.AddSingletonListConversion("encryption_config", "encryptionConfig")
		r.AddSingletonListConversion("vector_options", "vectorOptions")
	})
	p.AddResourceConfigurator("aws_opensearchserverless_collection_group", func(r *config.Resource) {
		// capacity_limits is a Plugin Framework list-of-objects attribute with a
		// size-1 validator. Framework attribute validators do not surface as
		// max_items in the provider schema, so upjet cannot infer the singleton
		// list on its own; enforce the embedded-object conversion explicitly.
		r.AddSingletonListConversion("capacity_limits", "capacityLimits")
	})
}
