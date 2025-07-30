// Package generate provides functions for discovering and generating validator registry files.
package generate

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode/utf8"

	"github.com/sivchari/govalid/cmd/generate-validators/templates"
)

// TestInfo contains information needed to generate test files.
type TestInfo struct {
	Name          string // e.g., "required", "maxlength"
	TitleCaseName string // e.g., "Required", "Maxlength"
}

// generateGovalidTests generates individual test files for each validator.
func generateGovalidTests(validators []ValidatorInfo, testDir, testTemplate string) error {
	for _, validator := range validators {
		// Convert to TestInfo
		testInfo := TestInfo{
			Name:          validator.MarkerName,
			TitleCaseName: toTitleCase(validator.MarkerName),
		}

		t, err := template.New("test").Funcs(templates.FuncMap).Parse(testTemplate)
		if err != nil {
			return fmt.Errorf("failed to parse test template for %s: %w", validator.MarkerName, err)
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, testInfo); err != nil {
			return fmt.Errorf("failed to execute test template for %s: %w", validator.MarkerName, err)
		}

		filename := filepath.Join(testDir, validator.MarkerName+"_test.go")
		if err := os.WriteFile(filename, buf.Bytes(), 0o600); err != nil {
			return fmt.Errorf("failed to write test file %s: %w", filename, err)
		}
	}

	return nil
}

// toTitleCase converts a marker name to title case.
func toTitleCase(s string) string {
	if s == "" {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	if size == 0 {
		return ""
	}

	return strings.ToUpper(string(r)) + s[size:]
}