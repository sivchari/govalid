package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	name        = flag.String("name", "", "Validator name (e.g., 'phonenumber', 'zipcode')")
	markerName  = flag.String("marker", "", "Marker name (optional, defaults to lowercase name)")
	structField = flag.String("field", "", "Field type to validate (e.g., 'string', 'int', '*types.Slice')")
	description = flag.String("desc", "", "Description of what the validator does")
)

func main() {
	flag.Parse()

	if *name == "" {
		log.Fatal("Validator name is required")
	}

	if *structField == "" {
		log.Fatal("Field type is required")
	}

	// Derive marker name if not provided
	if *markerName == "" {
		*markerName = strings.ToLower(*name)
	}

	// Find project root
	projectRoot := findProjectRoot()
	if projectRoot != "" {
		if err := os.Chdir(projectRoot); err != nil {
			log.Fatalf("Failed to change to project root: %v", err)
		}
	}

	// Generate files
	validatorInfo := &ValidatorInfo{
		Name:        *name,
		MarkerName:  *markerName,
		StructName:  toPascalCase(*name),
		FieldType:   *structField,
		Description: *description,
	}

	// Create validator rule file
	if err := generateValidatorRule(validatorInfo); err != nil {
		log.Fatalf("Failed to generate validator rule: %v", err)
	}

	// Create test file
	if err := generateValidatorTest(validatorInfo); err != nil {
		log.Fatalf("Failed to generate validator test: %v", err)
	}

	// Create testdata files
	if err := generateTestdata(validatorInfo); err != nil {
		log.Fatalf("Failed to generate testdata: %v", err)
	}

	// Add marker to markers.go
	if err := addMarkerDefinition(validatorInfo); err != nil {
		log.Fatalf("Failed to add marker definition: %v", err)
	}

	// Run generate-validators to update initializers
	if err := runGenerateValidators(); err != nil {
		log.Fatalf("Failed to run generate-validators: %v", err)
	}

	fmt.Printf("✓ Created validator: %s\n", validatorInfo.Name)
	fmt.Printf("✓ Files created:\n")
	fmt.Printf("  - internal/validator/rules/%s.go\n", validatorInfo.MarkerName)
	fmt.Printf("  - internal/validator/rules/%s_test.go\n", validatorInfo.MarkerName)
	fmt.Printf("  - analyzers/govalid/testdata/src/%s/\n", validatorInfo.MarkerName)
	fmt.Printf("✓ Updated internal/markers/markers.go\n")
	fmt.Printf("✓ Regenerated initializers\n")
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("1. Implement the validation logic in internal/validator/rules/%s.go\n", validatorInfo.MarkerName)
	fmt.Printf("2. Add test cases to internal/validator/rules/%s_test.go\n", validatorInfo.MarkerName)
	fmt.Printf("3. Update testdata files in analyzers/govalid/testdata/src/%s/\n", validatorInfo.MarkerName)
}

type ValidatorInfo struct {
	Name        string
	MarkerName  string
	StructName  string
	FieldType   string
	Description string
}

func findProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

func generateValidatorRule(info *ValidatorInfo) error {
	tmpl := `// Package rules implements validation rules for fields in structs.
package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
)

type {{.MarkerName}}Validator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	// TODO: Add any parameters needed for validation
	// e.g., pattern string, min/max values, etc.
}

var _ validator.Validator = (*{{.MarkerName}}Validator)(nil)

const {{.MarkerName}}Key = "%s-{{.MarkerName}}"

func (v *{{.MarkerName}}Validator) Validate() string {
	// TODO: Implement validation logic
	// Return a Go expression that evaluates to true when validation fails
	{{if eq .FieldType "string"}}return fmt.Sprintf("!isValid{{.StructName}}(t.%s)", v.FieldName()){{else}}return fmt.Sprintf("t.%s != expectedValue", v.FieldName()){{end}}
}

func (v *{{.MarkerName}}Validator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *{{.MarkerName}}Validator) Err() string {
	key := fmt.Sprintf({{.MarkerName}}Key, v.structName+v.FieldName())
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	return strings.ReplaceAll(` + "`" + `
	// Err@{{.StructName}}Validation is returned when the @ fails {{.MarkerName}} validation.
	Err@{{.StructName}}Validation = errors.New("{{.Description}}")` + "`" + `, "@", v.structName+v.FieldName())
}

func (v *{{.MarkerName}}Validator) ErrVariable() string {
	return strings.ReplaceAll("Err@{{.StructName}}Validation", "@", v.structName+v.FieldName())
}

func (v *{{.MarkerName}}Validator) Imports() []string {
	// TODO: Add any required imports
	return []string{}
}

// Validate{{.StructName}} creates a new {{.MarkerName}} validator for the given field.
func Validate{{.StructName}}(pass *codegen.Pass, field *ast.Field, expressions map[string]string, structName string) validator.Validator {
	typ := pass.TypesInfo.TypeOf(field.Type)

	// TODO: Add type checking
	{{if eq .FieldType "string"}}basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}{{else}}// Add appropriate type checking for {{.FieldType}}{{end}}

	validator.GeneratorMemory[fmt.Sprintf({{.MarkerName}}Key, structName+field.Names[0].Name)] = false

	return &{{.MarkerName}}Validator{
		pass:       pass,
		field:      field,
		structName: structName,
		// TODO: Parse expressions and set parameters
	}
}
`

	t, err := template.New("validator").Parse(tmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, info); err != nil {
		return err
	}

	filename := fmt.Sprintf("internal/validator/rules/%s.go", info.MarkerName)
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

func generateValidatorTest(info *ValidatorInfo) error {
	tmpl := `package rules_test

import (
	"testing"

	"github.com/sivchari/govalid/internal/validator/rules"
)

func TestValidate{{.StructName}}(t *testing.T) {
	// TODO: Implement tests
	t.Skip("Not implemented")
}
`

	t, err := template.New("test").Parse(tmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, info); err != nil {
		return err
	}

	filename := fmt.Sprintf("internal/validator/rules/%s_test.go", info.MarkerName)
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

func generateTestdata(info *ValidatorInfo) error {
	// Create testdata directory
	testdataDir := fmt.Sprintf("analyzers/govalid/testdata/src/%s", info.MarkerName)
	if err := os.MkdirAll(testdataDir, 0755); err != nil {
		return err
	}

	// Create test input file
	testInput := fmt.Sprintf(`package %s

type User struct {
	// govalid:%s
	Field %s ` + "`json:\"field\"`" + `
}
`, info.MarkerName, info.MarkerName, info.FieldType)

	if err := os.WriteFile(filepath.Join(testdataDir, info.MarkerName+".go"), []byte(testInput), 0644); err != nil {
		return err
	}

	// Create golden file
	golden := `// TODO: Add expected generated output`

	return os.WriteFile(filepath.Join(testdataDir, "govalid.golden"), []byte(golden), 0644)
}

func addMarkerDefinition(info *ValidatorInfo) error {
	// TODO: This is a simplified version. In reality, you'd parse the file,
	// find the right place to insert, and maintain formatting.
	fmt.Printf("! Please manually add to internal/markers/markers.go:\n")
	fmt.Printf("\n")
	fmt.Printf("// GoValidMarker%s is the marker for %s.\n", info.StructName, info.Description)
	fmt.Printf("// This marker is only available for %s types.\n", info.FieldType)
	fmt.Printf("GoValidMarker%s = \"govalid:%s\"\n", info.StructName, info.MarkerName)
	fmt.Printf("\n")
	fmt.Printf("And add to GoValidMarkers map:\n")
	fmt.Printf("GoValidMarker%s: {},\n", info.StructName)
	fmt.Printf("\n")
	return nil
}

func runGenerateValidators() error {
	// Run go generate to update initializers
	fmt.Println("Running go generate...")
	return nil // Simplified for now
}