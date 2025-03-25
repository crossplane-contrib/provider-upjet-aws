// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package acmpca

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the acmpca group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_acmpca_certificate_authority", func(r *config.Resource) {
		// NOTE(muvaf): It causes circular dependency. See https://github.com/crossplane/crossplane-runtime/issues/313
		delete(r.References, "revocation_configuration.crl_configuration.s3_bucket_name")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"revocation_configuration"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["certificate"].(string); ok {
				conn["certificate"] = []byte(a)
			}
			if a, ok := attr["certificate_chain"].(string); ok {
				conn["certificate_chain"] = []byte(a)
			}
			if a, ok := attr["certificate_signing_request"].(string); ok {
				conn["certificate_signing_request"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_acmpca_certificate", func(r *config.Resource) {
		r.TerraformResource.Schema["certificate_signing_request"].Sensitive = true
		r.References = map[string]config.Reference{
			"certificate_authority_arn": {
				TerraformName: "aws_acmpca_certificate_authority",
			},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["certificate"].(string); ok {
				conn["certificate"] = []byte(a)
			}
			if a, ok := attr["certificate_chain"].(string); ok {
				conn["certificate_chain"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_acmpca_certificate_authority_certificate", func(r *config.Resource) {
		r.TerraformResource.Schema["certificate"].Sensitive = true
		r.TerraformResource.Schema["certificate_chain"].Sensitive = true
		r.References = map[string]config.Reference{
			"certificate_authority_arn": {
				TerraformName: "aws_acmpca_certificate_authority",
			},
		}
	})
}
