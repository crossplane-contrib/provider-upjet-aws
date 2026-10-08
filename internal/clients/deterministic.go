// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

// deterministicResourceTypes contains Terraform resource types that are known
// to have deterministic external names even when DisableNameInitializer is true.
// For example, 1-to-1 sub-resources whose identifier is derived deterministically
// from their parent resource (e.g. S3 bucket sub-resources) or spec fields.
var deterministicResourceTypes = map[string]struct{}{
	"aws_s3_bucket_versioning":                          {},
	"aws_s3_bucket_public_access_block":                 {},
	"aws_s3_bucket_ownership_controls":                  {},
	"aws_s3_bucket_policy":                              {},
	"aws_s3_bucket_server_side_encryption_configuration": {},
	"aws_s3_bucket_lifecycle_configuration":             {},
	"aws_s3_bucket_cors_configuration":                  {},
	"aws_s3_bucket_website_configuration":               {},
	"aws_s3_bucket_request_payment_configuration":       {},
	"aws_s3_bucket_accelerate_configuration":            {},
	"aws_s3_bucket_acl":                                 {},
	"aws_s3_bucket_logging":                             {},
	"aws_s3_bucket_notification":                        {},
	"aws_s3_bucket_intelligent_tiering_configuration":   {},
	"aws_s3_bucket_metric":                              {},
	"aws_s3_bucket_inventory":                           {},
	"aws_s3_bucket_abac":                                {},
	"aws_iam_account_alias":                             {},
	"aws_iam_account_password_policy":                   {},
	"aws_ebs_default_kms_key":                           {},
	"aws_ebs_encryption_by_default":                     {},
}

// IsDeterministicExternalName returns true if the given Terraform resource type
// is known to have a deterministic external name even if DisableNameInitializer is true.
func IsDeterministicExternalName(resourceType string) bool {
	_, ok := deterministicResourceTypes[resourceType]
	return ok
}
