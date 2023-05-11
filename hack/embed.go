package hack

import _ "embed"

// MainTemplate is populated with provider main program template.
//
//go:embed main.go.tmpl
var MainTemplate string
