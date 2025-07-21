package main

import (
	"bytes"
	_ "embed"
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
)

// Embed templates
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

// Fixed paths
const (
	rulesDir       = "internal/validator/rules"
	outputDir      = "internal/validator/registry/initializers"
	registryFile   = "internal/analyzers/govalid/registry_init.go"
	markersFile    = "internal/markers/markers_generated.go"
)

var (
	marker = flag.String("marker", "", "Create a new validator with the given marker name (e.g., 'phoneNumber')")
)

// ValidatorInfo contains information about a discovered validator.
type ValidatorInfo struct {
	MarkerName     string // e.g., "required", "maxlength"
	MarkerConstant string // e.g., "GoValidMarkerRequired"
	FunctionName   string // e.g., "ValidateRequired"
	StructName     string // e.g., "Required", "MaxLength"
	IsSpecial      bool   // true for validators like required that don't use expressions
	Description    string // Description from comments
	TypeInfo       string // Type information from comments
	HasParameter   bool   // Whether validator takes a parameter
	ParameterName  string // Name of the parameter field
	ErrorMessage   string // Custom error message
	TypeCheckCode  string // Type checking code
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
		return
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
			return os.Chdir(dir)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}

func scaffoldValidator(markerName string) error {
	// Convert marker name to various forms
	markerLower := strings.ToLower(markerName)
	structName := toPascalCase(markerLower)
	
	info := &ValidatorInfo{
		MarkerName:     markerLower,
		StructName:     structName,
		FunctionName:   "Validate" + structName,
		MarkerConstant: "GoValidMarker" + structName,
		HasParameter:   true, // Most validators have parameters
		ParameterName:  "value",
		TypeCheckCode:  generateTypeCheckCode("string"), // Default to string
	}
	
	// Create validator rule file
	validatorPath := filepath.Join(rulesDir, markerLower+".go")
	if err := generateFromTemplate(validatorTemplate, info, validatorPath); err != nil {
		return fmt.Errorf("failed to generate validator: %w", err)
	}
	
	fmt.Printf("✓ Created validator scaffold: %s\n", validatorPath)
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("1. Implement validation logic in %s\n", validatorPath)
	fmt.Printf("2. Update type checking if needed (currently set for string)\n")
	fmt.Printf("3. Add test cases\n")
	fmt.Printf("4. Run 'go generate ./analyzers/govalid' to update registry\n")
	
	return nil
}

func generateTypeCheckCode(fieldType string) string {
	switch fieldType {
	case "string":
		return `	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}`
	case "numeric":
		return `	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || (basic.Info()&types.IsNumeric) == 0 {
		return nil
	}`
	case "collection":
		return `	switch typ.Underlying().(type) {
	case *types.Slice, *types.Array, *types.Map, *types.Chan:
		// Valid collection type
	default:
		return nil
	}`
	default:
		return `	// TODO: Add appropriate type checking`
	}
}

func generateAll() error {
	validators, err := discoverValidators()
	if err != nil {
		return fmt.Errorf("failed to discover validators: %w", err)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
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
	if err := os.MkdirAll(filepath.Dir(markersFile), 0755); err != nil {
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
		return nil, err
	}

	var validatorType string
	var validateFunc string
	var isSpecial bool

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
					
					// Check if it's a special validator (3 params instead of 4)
					if node.Type.Params != nil && len(node.Type.Params.List) == 3 {
						isSpecial = true
					}
				}
			}
		}
		return true
	})

	if validatorType == "" || validateFunc == "" {
		return nil, nil
	}

	// Extract marker name from validator type
	markerName := strings.TrimSuffix(validatorType, "Validator")
	markerName = strings.ToLower(markerName)
	
	// Convert to struct name
	structName := toPascalCase(markerName)
	
	// Convert to marker constant name
	markerConstant := "GoValidMarker" + structName

	// Get description and type info
	description, typeInfo := getValidatorDescription(markerName)

	return &ValidatorInfo{
		MarkerName:     markerName,
		MarkerConstant: markerConstant,
		FunctionName:   validateFunc,
		StructName:     structName,
		IsSpecial:      isSpecial,
		Description:    description,
		TypeInfo:       typeInfo,
	}, nil
}

func getValidatorDescription(markerName string) (description, typeInfo string) {
	// Default descriptions based on marker names
	switch markerName {
	case "required":
		return "the marker for required fields", ""
	case "lt":
		return "the marker for checking if a value is less than a specified maximum", "numeric types"
	case "gt":
		return "the marker for checking if a value is greater than a specified minimum", "numeric types"
	case "lte":
		return "the marker for checking if a value is less than or equal to a specified maximum", "numeric types"
	case "gte":
		return "the marker for checking if a value is greater than or equal to a specified minimum", "numeric types"
	case "maxlength":
		return "the marker for checking if a string's length is less than or equal to a specified maximum", "string types"
	case "minlength":
		return "the marker for checking if a string's length is greater than or equal to a specified minimum", "string types"
	case "maxitems":
		return "the marker for checking if a collection's length is less than or equal to a specified maximum", "slice, array, map, and channel types"
	case "minitems":
		return "the marker for checking if a collection's length is greater than or equal to a specified minimum", "slice, array, map, and channel types"
	case "enum":
		return "the marker for checking if a field value is within a specified set of allowed values", "string, numeric, and custom types with comparable values"
	case "email":
		return "the marker for checking if a string is a valid email address", "string types"
	case "uuid":
		return "the marker for checking if a string is a valid UUID", "string types"
	case "url":
		return "the marker for checking if a string is a valid URL", "string types"
	case "cel":
		return "the marker for CEL (Common Expression Language) validation", "all types"
	default:
		return fmt.Sprintf("the marker for %s validation", markerName), ""
	}
}

func toPascalCase(s string) string {
	// Handle special cases for consistency with existing constants
	switch s {
	case "maxlength":
		return "MaxLength"
	case "minlength":
		return "MinLength"
	case "maxitems":
		return "MaxItems"
	case "minitems":
		return "MinItems"
	case "lt":
		return "LT"
	case "gt":
		return "GT"
	case "lte":
		return "LTE"
	case "gte":
		return "GTE"
	case "cel":
		return "CEL"
	case "uuid":
		return "UUID"
	case "url":
		return "URL"
	default:
		return strings.Title(s)
	}
}

func generateInitializers(validators []ValidatorInfo) error {
	funcMap := template.FuncMap{
		"firstLetter": func(s string) string {
			if len(s) > 0 {
				return strings.ToLower(s[:1])
			}
			return "x"
		},
	}

	// Generate individual initializer files
	for _, validator := range validators {
		t, err := template.New("initializer").Funcs(funcMap).Parse(initializerTemplate)
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, validator); err != nil {
			return fmt.Errorf("failed to execute template for %s: %w", validator.MarkerName, err)
		}

		filename := filepath.Join(outputDir, validator.MarkerName+".go")
		if err := os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filename, err)
		}
	}

	return nil
}

func generateFromTemplate(tmplContent string, data interface{}, outputPath string) error {
	t, err := template.New("template").Parse(tmplContent)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}

	return os.WriteFile(outputPath, buf.Bytes(), 0644)
}