package govalid

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"text/template"

	"github.com/gostaticanalysis/codegen"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/imports"

	"github.com/sivchari/govalid/internal/analyzers/markers"
	govaliderrors "github.com/sivchari/govalid/internal/errors"
	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/registry"
)

const (
	// Name is the name of the govalid generator.
	Name = "govalid"
	// Doc is the documentation for the govalid generator.
	Doc = "govalid generates type-safe validation code for structs based on markers."
)

var (
	// dryRun indicates whether the generator should run in dry-run mode.
	dryRun bool
)

// generator is the main type for the govalid analyzer.
type generator struct{}

// newGenerator creates a new instance of the govalid generator.
func newGenerator() (*codegen.Generator, error) {
	g := &generator{}

	generator := &codegen.Generator{
		Name:     Name,
		Doc:      Doc,
		Run:      g.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer, markers.Analyzer},
	}

	return generator, nil
}

// TemplateData holds the data for the template used to generate validation code.
type TemplateData struct {
	PackageName    string
	TypeName       string
	Metadata       []*AnalyzedMetadata
	ImportPackages map[string]struct{}
}

// run is the main function that runs the govalid analyzer.
func (g *generator) run(pass *codegen.Pass) error {
	inspector, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return govaliderrors.ErrCouldNotGetInspector
	}

	markersInspect, ok := pass.ResultOf[markers.Analyzer].(markers.Markers)
	if !ok {
		return govaliderrors.ErrCouldNotGetInspector
	}

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	tmplList := map[string]TemplateData{}

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			return
		}

		for _, spec := range genDecl.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				return
			}

			typeMarkers := markersInspect.TypeMarkers(ts)

			structType, ok := ts.Type.(*ast.StructType)
			if !ok {
				return
			}

			metadata := analyzeMarker(pass, markersInspect, typeMarkers, structType, "", ts.Name.Name)
			if len(metadata) == 0 {
				return
			}

			tmplData := TemplateData{
				PackageName:    pass.Pkg.Name(),
				TypeName:       ts.Name.Name,
				Metadata:       metadata,
				ImportPackages: collectImportPackages(metadata),
			}

			data, ok := tmplList[ts.Name.Name]
			if ok {
				data.Metadata = append(data.Metadata, tmplData.Metadata...)
			}

			if err := writeFile(pass, ts, tmplData); err != nil {
				panic(fmt.Sprintf("failed to write file for %s: %v", ts.Name.Name, err))
			}
		}
	})

	return nil
}

// AnalyzedMetadata holds the metadata for a field in a struct, including its validators and parent variable name.
type AnalyzedMetadata struct {
	Validators     []validator.Validator
	ParentVariable string
}

// makeValidatorInput contains all the input parameters needed for makeValidator function.
type makeValidatorInput struct {
	Pass       *codegen.Pass
	Markers    []markers.Marker
	Field      *ast.Field
	StructName string
	ParentPath string
}

func analyzeMarker(pass *codegen.Pass, markersInspect markers.Markers, typeMarkers markers.MarkerSet, structType *ast.StructType, parent, structName string) []*AnalyzedMetadata {
	analyzed := make([]*AnalyzedMetadata, 0)

	typeMarkersList := make([]markers.Marker, 0, len(typeMarkers))
	for _, marker := range typeMarkers {
		typeMarkersList = append(typeMarkersList, marker)
	}
	sort.SliceStable(typeMarkersList, func(i, j int) bool {
		return strings.Compare(typeMarkersList[i].Identifier, typeMarkersList[j].Identifier) < 0
	})

	for _, field := range structType.Fields.List {
		validators := make([]validator.Validator, 0)

		// Apply markers to the field
		fieldMarkers := markersInspect.FieldMarkers(field)

		fieldMarkersList := make([]markers.Marker, 0, len(fieldMarkers))
		for _, marker := range fieldMarkers {
			fieldMarkersList = append(fieldMarkersList, marker)
		}
		sort.SliceStable(fieldMarkersList, func(i, j int) bool {
			return strings.Compare(fieldMarkersList[i].Identifier, fieldMarkersList[j].Identifier) < 0
		})

		markersList := append(typeMarkersList, fieldMarkersList...)

		input := makeValidatorInput{
			Pass:       pass,
			Markers:    markersList,
			Field:      field,
			StructName: structName,
			ParentPath: parent,
		}

		// Traverse nested structs
		structType, ok := field.Type.(*ast.StructType)
		if !ok {
			validators = makeValidator(input)
			if len(validators) == 0 {
				continue
			}

			analyzed = append(analyzed, &AnalyzedMetadata{
				Validators:     validators,
				ParentVariable: parent,
			})

			continue
		}

		for _, field := range structType.Fields.List {
			/*
				Propagate parent markers to nested fields

				// +govalid:required
				type Nested struct {
					Name string `json:"name"`
				}
			*/
			input.Field = field
			input.Markers = append(typeMarkersList, fieldMarkersList...)
			validators = append(validators, makeValidator(input)...)
		}

		// Add the parent variable name to the analyzed metadata
		var parentVariable string
		if parent != "" {
			parentVariable = fmt.Sprintf("%s.%s", parent, field.Names[0].Name)
		} else {
			parentVariable = field.Names[0].Name
		}

		if len(validators) > 0 {
			analyzed = append(analyzed, &AnalyzedMetadata{
				Validators:     validators,
				ParentVariable: parentVariable,
			})
		}

		// Recursively analyze nested structs
		analyzed = append(analyzed, analyzeMarker(pass, markersInspect, typeMarkers, structType, parentVariable, structName)...)
	}

	return analyzed
}

func makeValidator(input makeValidatorInput) []validator.Validator {
	validators := make([]validator.Validator, 0)

	for _, marker := range input.Markers {
		factory, err := registry.Validator(marker.Identifier)
		if err != nil {
			// Validator not found, skip
			continue
		}

		ruleName := strings.TrimPrefix(marker.Identifier, "govalid:")

		validatorInput := registry.ValidatorInput{
			Pass:        input.Pass,
			Field:       input.Field,
			Expressions: marker.Expressions,
			StructName:  input.StructName,
			RuleName:    ruleName,
			ParentPath:  input.ParentPath,
		}
		v := factory(validatorInput)

		if v == nil {
			continue
		}

		validators = append(validators, v)
	}

	return validators
}

// collectImportPackages analyzes validators and collects required import packages.
func collectImportPackages(metadata []*AnalyzedMetadata) map[string]struct{} {
	packages := make(map[string]struct{})

	for _, meta := range metadata {
		for _, validator := range meta.Validators {
			for _, pkg := range validator.Imports() {
				packages[pkg] = struct{}{}
			}
		}
	}

	return packages
}

func writeFile(pass *codegen.Pass, ts *ast.TypeSpec, tmplData TemplateData) error {
	t, err := template.New("validator").Funcs(template.FuncMap{
		"trimDots": func(s string) string {
			return strings.ReplaceAll(s, ".", "")
		},
	}).Parse(ValidationTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, tmplData); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Use goimports to format the source code with proper import grouping
	src, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return fmt.Errorf("failed to format source code with imports: %w", err)
	}

	src, err = format.Source(src)
	if err != nil {
		return fmt.Errorf("failed to format source code: %w", err)
	}

	if testing.Testing() || dryRun {
		if _, err := pass.Print(string(src)); err != nil {
			return fmt.Errorf("failed to print source code: %w", err)
		}

		return nil
	}

	originalFilePath := pass.Fset.Position(ts.Pos()).Filename
	fileName := strings.TrimSuffix(filepath.Base(originalFilePath), filepath.Ext(originalFilePath))
	typeName := ts.Name.Name
	fileName = fmt.Sprintf("%s_%s_validator.go", fileName, strings.ToLower(typeName))

	file, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}()

	if _, err := fmt.Fprint(file, string(src)); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
