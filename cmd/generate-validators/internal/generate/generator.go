// Package generate provides functions for discovering and generating validator registry files.
package generate

import (
	"fmt"
	"os"
	"path/filepath"
)

// All generates all registry files from existing validators.
func All(rulesDir, outputDir, registryFile, markersFile string, templates *Templates) error {
	validators, err := DiscoverValidators(rulesDir)
	if err != nil {
		return fmt.Errorf("failed to discover validators: %w", err)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0o750); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate individual initializers
	if err := generateInitializers(validators, outputDir, templates.Initializer); err != nil {
		return fmt.Errorf("failed to generate initializers: %w", err)
	}

	// Generate all.go file
	if err := generateFromTemplate(templates.All, validators, filepath.Join(outputDir, "all.go")); err != nil {
		return fmt.Errorf("failed to generate all.go: %w", err)
	}

	// Generate registry init file
	if err := generateFromTemplate(templates.RegistryInit, validators, registryFile); err != nil {
		return fmt.Errorf("failed to generate registry init: %w", err)
	}

	// Generate markers file
	if err := os.MkdirAll(filepath.Dir(markersFile), 0o750); err != nil {
		return fmt.Errorf("failed to create markers directory: %w", err)
	}

	if err := generateFromTemplate(templates.Markers, validators, markersFile); err != nil {
		return fmt.Errorf("failed to generate markers file: %w", err)
	}

	// Generate test files
	testDir := filepath.Join("internal", "analyzers", "govalid", "tests")

	if templates.GovalidTest != "" {
		// Create test directory if it doesn't exist
		if err := os.MkdirAll(testDir, 0o750); err != nil {
			return fmt.Errorf("failed to create test directory: %w", err)
		}

		if err := generateGovalidTests(validators, testDir, templates.GovalidTest); err != nil {
			return fmt.Errorf("failed to generate test files: %w", err)
		}

		fmt.Printf("✓ Generated test files for %d validators\n", len(validators))
	}

	fmt.Printf("✓ Generated initializers for %d validators\n", len(validators))

	return nil
}

// Templates contains all template strings needed for generation.
type Templates struct {
	Initializer  string
	All          string
	RegistryInit string
	Markers      string
	GovalidTest  string
}
