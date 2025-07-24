package scaffold

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/sivchari/govalid/cmd/generate-validators/templates"
)

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
