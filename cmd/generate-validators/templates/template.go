package templates

import (
	"strings"
	"text/template"
	"unicode/utf8"

	"golang.org/x/text/cases"
)

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
