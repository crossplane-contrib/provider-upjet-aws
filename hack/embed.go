// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package hack

import _ "embed"

// MainTemplate is populated with provider main program template.
//
//go:embed main.go.tmpl
var MainTemplate string
