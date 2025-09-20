// Package scaffold provides utilities to generate files from templates.
package scaffold

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// CreateValidator creates a new validator scaffold file with the given marker name.
// It checks if the file already exists to prevent overwriting.
func CreateValidator(markerName, validatorTemplate, rulesDir string) error {
	if markerName == "" {
		return errors.New("marker name cannot be empty")
	}

	// Convert marker name to various forms
	markerLower := strings.ToLower(markerName)

	structName := cases.Title(language.English).String(markerLower)

	// Check if validator file already exists
	validatorPath := filepath.Join(rulesDir, markerLower+".go")
	if err := checkFileExists(validatorPath); err != nil {
		return err
	}

	// Use a map for template data since we need more fields than ValidatorInfo
	data := map[string]string{
		"MarkerName":     markerLower,
		"StructName":     structName,
		"FunctionName":   "Validate" + structName,
		"MarkerConstant": "GoValidMarker" + structName,
	}

	// Create validator rule file
	if err := generateFromTemplate(validatorTemplate, data, validatorPath); err != nil {
		return fmt.Errorf("failed to generate validator: %w", err)
	}

	fmt.Printf("✓ Created validator scaffold: %s\n", validatorPath)

	return nil
}

// checkFileExists checks if a file already exists and returns an error if it does.
func checkFileExists(path string) error {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("validator file already exists: %s", path)
	}

	return nil
}

// CreateGovalidTest creates a new govalid test file for the given marker.
func CreateGovalidTest(markerName, testTemplate, testDir string) error {
	if markerName == "" {
		return errors.New("marker name cannot be empty")
	}

	// Convert marker name to various forms
	markerLower := strings.ToLower(markerName)

	titleCaseName := cases.Title(language.English).String(markerLower)

	// Generate test file path
	testPath := filepath.Join(testDir, markerLower+"_test.go")

	// Use a map for template data
	data := map[string]string{
		"Name":          markerLower,
		"TitleCaseName": titleCaseName,
	}

	// Create test file
	if err := generateFromTemplate(testTemplate, data, testPath); err != nil {
		return fmt.Errorf("failed to generate test file: %w", err)
	}

	fmt.Printf("✓ Created test file: %s\n", testPath)

	return nil
}
