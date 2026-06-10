// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package templates

import _ "embed"

// ControllerTemplate is the template for the MR controller setup files.
//
//go:embed controller.go.tmpl
var ControllerTemplate string
