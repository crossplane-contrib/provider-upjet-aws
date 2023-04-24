/*
 Copyright 2022 Upbound Inc.
*/

package features

import "github.com/crossplane/crossplane-runtime/pkg/feature"

// Feature flags.
const (
	// EnableAlphaExternalSecretStores enables alpha support for
	// External Secret Stores. See the below design for more details.
	// https://github.com/crossplane/crossplane/blob/390ddd/design/design-doc-external-secret-stores.md
	EnableAlphaExternalSecretStores feature.Flag = "EnableAlphaExternalSecretStores"

	// EnableAlphaManagementPolicies enables alpha support for
	// Management Policies. See the below design for more details.
	// https://github.com/crossplane/crossplane/pull/3531
	EnableAlphaManagementPolicies feature.Flag = "EnableAlphaManagementPolicies"
)
