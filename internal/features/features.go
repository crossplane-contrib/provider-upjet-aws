// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package features

import "github.com/crossplane/crossplane-runtime/v2/pkg/feature"

// Feature flags.
const (
	// EnableAlphaExternalSecretStores enables alpha support for
	// External Secret Stores. See the below design for more details.
	// https://github.com/crossplane/crossplane/blob/390ddd/design/design-doc-external-secret-stores.md
	EnableAlphaExternalSecretStores feature.Flag = "EnableAlphaExternalSecretStores"

	// EnableBetaManagementPolicies enables beta support for
	// Management Policies. See the below design for more details.
	// https://github.com/crossplane/crossplane/pull/3531
	EnableBetaManagementPolicies feature.Flag = "EnableBetaManagementPolicies"
)
