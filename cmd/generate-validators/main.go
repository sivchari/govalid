// Package main is a tool to generate Go validators and their initializers
package main

import (
	"bytes"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode/utf8"

	"golang.org/x/text/cases"
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

var funcMap = template.FuncMap{
	"firstLetter": func(s string) string {
		r, size := utf8.DecodeRuneInString(s)
		if size == 0 {
			return "x"
		}

		return strings.ToLower(string(r))
	},
	"title": cases.Title,
}

var (
	marker = flag.String("marker", "", "Create a new validator with the given marker name (e.g., 'phoneNumber')")
)

// ValidatorInfo contains information about a discovered validator.
type ValidatorInfo struct {
	MarkerName   string // e.g., "required", "maxlength"
	FunctionName string // e.g., "ValidateRequired"
}

func main() {
	flag.Parse()

	// Find project root
	if err := changeToProjectRoot(); err != nil {
		log.Fatalf("Failed to find project root: %v", err)
	}

	// If marker flag is provided, scaffold a new validator
	if *marker != "" {
		if err := scaffoldValidator(*marker); err != nil {
			log.Fatalf("Failed to scaffold validator: %v", err)
		}
	}

	// Otherwise, generate all files from existing validators
	if err := generateAll(); err != nil {
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

func scaffoldValidator(markerName string) error {
	// Convert marker name to various forms
	markerLower := strings.ToLower(markerName)
	// Simple inline PascalCase conversion
	r, size := utf8.DecodeRuneInString(markerLower)
	if size == 0 {
		return errors.New("marker name cannot be empty")
	}

	structName := strings.ToUpper(string(r)) + markerLower[size:]

	// Use a map for template data since we need more fields than ValidatorInfo
	data := map[string]string{
		"MarkerName":     markerLower,
		"StructName":     structName,
		"FunctionName":   "Validate" + structName,
		"MarkerConstant": "GoValidMarker" + structName,
	}

	// Create validator rule file
	validatorPath := filepath.Join(rulesDir, markerLower+".go")
	if err := generateFromTemplate(validatorTemplate, data, validatorPath); err != nil {
		return fmt.Errorf("failed to generate validator: %w", err)
	}

	fmt.Printf("✓ Created validator scaffold: %s\n", validatorPath)
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("1. Implement validation logic in %s\n", validatorPath)
	fmt.Printf("2. Add test cases\n")
	fmt.Printf("3. Run 'go generate ./internal/analyzers/govalid' to update registry\n")

	return nil
}

func generateAll() error {
	validators, err := discoverValidators()
	if err != nil {
		return fmt.Errorf("failed to discover validators: %w", err)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0o750); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate individual initializers
	if err := generateInitializers(validators); err != nil {
		return fmt.Errorf("failed to generate initializers: %w", err)
	}

	// Generate all.go file
	if err := generateFromTemplate(allTemplate, validators, filepath.Join(outputDir, "all.go")); err != nil {
		return fmt.Errorf("failed to generate all.go: %w", err)
	}

	// Generate registry init file
	if err := generateFromTemplate(registryInitTemplate, validators, registryFile); err != nil {
		return fmt.Errorf("failed to generate registry init: %w", err)
	}

	// Generate markers file
	if err := os.MkdirAll(filepath.Dir(markersFile), 0o750); err != nil {
		return fmt.Errorf("failed to create markers directory: %w", err)
	}

	if err := generateFromTemplate(markersTemplate, validators, markersFile); err != nil {
		return fmt.Errorf("failed to generate markers file: %w", err)
	}

	fmt.Printf("✓ Generated initializers for %d validators\n", len(validators))

	return nil
}

func discoverValidators() ([]ValidatorInfo, error) {
	files, err := filepath.Glob(filepath.Join(rulesDir, "*.go"))
	if err != nil {
		return nil, fmt.Errorf("failed to glob rules directory: %w", err)
	}

	var validators []ValidatorInfo

	for _, file := range files {
		if strings.Contains(file, "_test.go") || strings.Contains(file, "validatorhelper") {
			continue
		}

		info, err := analyzeValidatorFile(file)
		if err != nil {
			log.Printf("Warning: failed to analyze %s: %v", file, err)

			continue
		}

		if info != nil {
			validators = append(validators, *info)
		}
	}

	// Sort for consistent output
	sort.Slice(validators, func(i, j int) bool {
		return validators[i].MarkerName < validators[j].MarkerName
	})

	return validators, nil
}

func analyzeValidatorFile(filepath string) (*ValidatorInfo, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %w", filepath, err)
	}

	var validatorType string

	var validateFunc string

	// Find validator struct and Validate function
	ast.Inspect(node, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.TypeSpec:
			if strings.HasSuffix(node.Name.Name, "Validator") {
				validatorType = node.Name.Name
			}
		case *ast.FuncDecl:
			if node.Name.IsExported() && strings.HasPrefix(node.Name.Name, "Validate") {
				// Check if it returns validator.Validator
				if node.Type.Results != nil && len(node.Type.Results.List) == 1 {
					validateFunc = node.Name.Name
				}
			}
		}

		return true
	})

	if validatorType == "" || validateFunc == "" {
		return nil, nil //nolint:nilnil // No valid validator found
	}

	// Extract marker name from validator type
	markerName := strings.TrimSuffix(validatorType, "Validator")
	markerName = strings.ToLower(markerName)

	return &ValidatorInfo{
		MarkerName:   markerName,
		FunctionName: validateFunc,
	}, nil
}

func generateInitializers(validators []ValidatorInfo) error {
	// Generate individual initializer files
	for _, validator := range validators {
		t, err := template.New("initializer").Funcs(funcMap).Parse(initializerTemplate)
		if err != nil {
			return fmt.Errorf("failed to parse template for %s: %w", validator.MarkerName, err)
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, validator); err != nil {
			return fmt.Errorf("failed to execute template for %s: %w", validator.MarkerName, err)
		}

		filename := filepath.Join(outputDir, validator.MarkerName+".go")
		if err := os.WriteFile(filename, buf.Bytes(), 0o600); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filename, err)
		}
	}

	return nil
}

func generateFromTemplate(tmplContent string, data any, outputPath string) error {
	t, err := template.New("template").Funcs(funcMap).Parse(tmplContent)
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
