// Package templates provides a set of template functions for use in Go templates.
package templates

import (
	"strings"
	"text/template"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	"title": func(s string) string {
		return cases.Title(language.English).String(s)
	},
}
