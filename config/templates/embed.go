// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package templates

import _ "embed"

// MainTemplate is populated with provider main program template.
//
//go:embed main.go.tmpl
var MainTemplate string

// ControllerTemplate is the template for the MR controller setup files.
//
//go:embed controller.go.tmpl
var ControllerTemplate string

// TerraformedTemplate is populated with conversion methods implementing
// Terraformed interface on CRD structs.
//
//go:embed terraformed.go.tmpl
var TerraformedTemplate string
