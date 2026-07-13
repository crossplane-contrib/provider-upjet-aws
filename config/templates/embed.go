// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package templates

import _ "embed"

// TerraformedTemplate is populated with conversion methods implementing
// Terraformed interface on CRD structs.
//
//go:embed terraformed.go.tmpl
var TerraformedTemplate string
