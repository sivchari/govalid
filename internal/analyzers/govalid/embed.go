package govalid

import (
	_ "embed"
)

// Embed validation template
//go:embed templates/validation.go.tmpl
var ValidationTemplate string