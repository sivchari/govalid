// Package generate provides functions for discovering and generating validator registry files.
package generate

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

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
			TitleCaseName: cases.Title(language.English).String(validator.MarkerName),
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
