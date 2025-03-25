// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package s3

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry"
)

// Configure adds configurations for the s3 group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_s3_bucket", func(r *config.Resource) {
		// Mutually exclusive with:
		// aws_s3_bucket_accelerate_configuration
		// aws_s3_bucket_acl
		// aws_s3_bucket_cors_configuration
		// aws_s3_bucket_lifecycle_configuration
		// aws_s3_bucket_logging
		// aws_s3_bucket_object_lock_configuration
		// aws_s3_bucket_policy
		// aws_s3_bucket_replication_configuration
		// aws_s3_bucket_request_payment_configuration
		// aws_s3_bucket_server_side_encryption_configuration
		// aws_s3_bucket_versioning
		// aws_s3_bucket_website_configuration
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["id"] = []byte(a)
			}
			if a, ok := attr["arn"].(string); ok {
				conn["arn"] = []byte(a)
			}
			if a, ok := attr["region"].(string); ok {
				conn["region"] = []byte(a)
			}
			return conn, nil
		}
		config.MoveToStatus(r.TerraformResource, "acceleration_status", "acl", "grant", "cors_rule", "lifecycle_rule",
			"logging", "object_lock_configuration", "policy", "replication_configuration", "request_payer",
			"server_side_encryption_configuration", "versioning", "website", "arn")
		r.MetaResource.ExternalName = registry.RandRFC1123Subdomain
		r.TerraformConfigurationInjector = func(jsonMap map[string]any, params map[string]any) error {
			params["region"] = jsonMap["region"]
			// TODO: added to prevent extra reconciliations due to
			// late-initialization or drift. Has better be implemented
			// via defaulting.
			if _, ok := jsonMap["forceDestroy"]; !ok {
				params["force_destroy"] = false
			}
			return nil
		}
	})

	p.AddResourceConfigurator("aws_s3_bucket_acl", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"acl", "access_control_policy"},
		}
	})

	p.AddResourceConfigurator("aws_s3_bucket_metrics", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_s3_bucket_website_configuration", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"routing_rules", "routing_rule"},
		}
	})

	p.AddResourceConfigurator("aws_s3_object", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"etag", "kms_key_id"},
		}
		r.TerraformConfigurationInjector = func(jsonMap map[string]any, params map[string]any) error {
			// TODO: Has better be implemented via defaulting.
			if _, ok := jsonMap["acl"]; !ok {
				params["acl"] = "private"
			}
			return nil
		}
	})

	p.AddResourceConfigurator("aws_s3_object_copy", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"acl", "grant"},
		}
	})

	p.AddResourceConfigurator("aws_s3_bucket_notification", func(r *config.Resource) {
		// NOTE(muvaf): It causes circular dependency. See https://github.com/crossplane/crossplane-runtime/issues/313
		delete(r.References, "lambda_function.lambda_function_arn")
	})

	p.AddResourceConfigurator("aws_s3_bucket_analytics_configuration", func(r *config.Resource) {
		r.References["storage_class_analysis.data_export.destination.s3_bucket_destination.bucket_arn"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)`,
		}
	})

	p.AddResourceConfigurator("aws_s3_bucket_replication_configuration", func(r *config.Resource) {
		r.References["rule.destination.bucket"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)`,
		}
		r.References["rule.destination.encryption_configuration.replica_kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
		}
	})

	p.AddResourceConfigurator("aws_s3_bucket_inventory", func(r *config.Resource) {
		r.References["destination.bucket.bucket_arn"] = config.Reference{
			TerraformName: "aws_s3_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)`,
		}
	})

	// In tf provider version 5.86.0, this resource was migrated from the tf plugin sdk to the tf plugin framework
	// These customizations are largely necessary to preserve the previous schemas so the change is transparent to users.
	p.AddResourceConfigurator("aws_s3_bucket_lifecycle_configuration", func(r *config.Resource) {
		r.Version = "v1beta3"
		// Explicitly set the reconcile version, since this is the only version that has types compatible with the tf
		// plugin framework client at runtime.
		r.ControllerReconcileVersion = "v1beta3"
		r.SetCRDHubVersion("v1beta3")
		r.SetCRDStorageVersion("v1beta1") // to facilitate downgrades
		r.PreviousVersions = []string {"v1beta1", "v1beta2"}

		// tf sdk v2 serializes optional integers to tfjson as string.
		// tf plugin framework serializes optional integers to tfjson as number, which upjet converts to *float64
		// This provider currently relies on the tfjson to generate the schema, so we're patching the tfjson generated by the
		// framework resource in the tf provider so that when we convert the tfjson schema to the tf plugin sdk schema,
		// we get the same type as we were using before.

		// The following fields are *string in v1beta1 and v1beta2, and *float64 in v1beta3
		// rule[*].filter[0].object_size_greater_than
		// rule[*].filter[0].object_size_less_than
		// rule[*].noncurrent_version_expiration[0].newer_noncurrent_versions
		// rule[*].noncurrent_version_transition[*].newer_noncurrent_versions






		// There's a bug somewhere in upjet that's applying the docstring for prefix to these fields that end in prefix
		r.MetaResource.ArgumentDocs["rule.filter.prefix"] = `- (Optional) Prefix identifying one or more objects to which the rule applies. Defaults to an empty string ("") if not specified.`
		r.MetaResource.ArgumentDocs["rule.filter.and.prefix"] = `- (Optional) Prefix identifying one or more objects to which the rule applies.`
	})
}
