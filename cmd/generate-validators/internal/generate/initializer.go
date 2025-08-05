// Package generate provides functions for discovering and generating validator registry files.
package generate

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/sivchari/govalid/cmd/generate-validators/templates"
)

// generateInitializers generates individual initializer files for each validator.
func generateInitializers(validators []ValidatorInfo, outputDir, initializerTemplate string) error {
	for _, validator := range validators {
		t, err := template.New("initializer").Funcs(templates.FuncMap).Parse(initializerTemplate)
		if err != nil {
			return fmt.Errorf("failed to parse template for %s: %w", validator.MarkerName, err)
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, validator); err != nil {
			return fmt.Errorf("failed to execute template for %s: %w", validator.MarkerName, err)
		}

		filename := filepath.Join(outputDir, strings.ToLower(validator.MarkerName)+".go")
		if err := os.WriteFile(filename, buf.Bytes(), 0o600); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filename, err)
		}
	}

	return nil
}

// generateFromTemplate generates a file from a template string and data.
func generateFromTemplate(tmplContent string, data any, outputPath string) error {
	t, err := template.New("template").Funcs(templates.FuncMap).Parse(tmplContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	if err := os.WriteFile(outputPath, buf.Bytes(), 0o600); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputPath, err)
	}

	return nil
}
