package govalid

import (
	_ "embed"
)

// ValidationTemplate is the template for generating validation code.
//
//go:embed templates/validation.go.tmpl
var ValidationTemplate string
