// Package templates provides a set of template functions for use in Go templates.
package templates

import (
	"strings"
	"text/template"
	"unicode/utf8"

	"golang.org/x/text/cases"
)

// FuncMap provides a set of template functions for use in Go templates.
var FuncMap = template.FuncMap{
	"firstLetter": func(s string) string {
		r, size := utf8.DecodeRuneInString(s)
		if size == 0 {
			return "x"
		}

		return strings.ToLower(string(r))
	},
	"title": cases.Title,
}
