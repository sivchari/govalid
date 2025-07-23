// Package main is a tool to generate Go validators and their initializers
package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sivchari/govalid/cmd/generate-validators/internal/generate"
	"github.com/sivchari/govalid/cmd/generate-validators/internal/scaffold"
)

// Embed templates
//
//go:embed templates/initializer.go.tmpl
var initializerTemplate string

//go:embed templates/all.go.tmpl
var allTemplate string

//go:embed templates/registry_init.go.tmpl
var registryInitTemplate string

//go:embed templates/markers.go.tmpl
var markersTemplate string

//go:embed templates/validator.go.tmpl
var validatorTemplate string

// Fixed paths.
const (
	rulesDir     = "internal/validator/rules"
	outputDir    = "internal/validator/registry/initializers"
	registryFile = "internal/analyzers/govalid/registry_init.go"
	markersFile  = "internal/markers/markers_generated.go"
)


var (
	marker = flag.String("marker", "", "Create a new validator with the given marker name (e.g., 'phoneNumber')")
)


func main() {
	flag.Parse()

	// Find project root
	if err := changeToProjectRoot(); err != nil {
		log.Fatalf("Failed to find project root: %v", err)
	}

	// If marker flag is provided, scaffold a new validator
	if *marker != "" {
		if err := scaffold.CreateValidator(*marker, validatorTemplate, rulesDir); err != nil {
			log.Fatalf("Failed to scaffold validator: %v", err)
		}
	}

	// Generate all files from existing validators
	templates := generate.Templates{
		Initializer:  initializerTemplate,
		All:          allTemplate,
		RegistryInit: registryInitTemplate,
		Markers:      markersTemplate,
	}

	if err := generate.GenerateAll(rulesDir, outputDir, registryFile, markersFile, templates); err != nil {
		log.Fatalf("Failed to generate: %v", err)
	}
}

func changeToProjectRoot() error {
	dir, _ := os.Getwd()

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			if err := os.Chdir(dir); err != nil {
				return fmt.Errorf("failed to change directory to %s: %w", dir, err)
			}

			return nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return errors.New("go.mod not found")
		}

		dir = parent
	}
}


