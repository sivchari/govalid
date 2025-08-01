package govalid

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/gostaticanalysis/codegen"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

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
		(*ast.TypeSpec)(nil),
	}

	tmplList := map[string]TemplateData{}

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return
		}

		structType, ok := ts.Type.(*ast.StructType)
		if !ok {
			return
		}

		metadata := analyzeMarker(pass, markersInspect, structType, "", ts.Name.Name)
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
	})

	return nil
}

// AnalyzedMetadata holds the metadata for a field in a struct, including its validators and parent variable name.
type AnalyzedMetadata struct {
	Validators     []validator.Validator
	ParentVariable string
}

func analyzeMarker(pass *codegen.Pass, markersInspect markers.Markers, structType *ast.StructType, parent, structName string) []*AnalyzedMetadata {
	analyzed := make([]*AnalyzedMetadata, 0)

	for _, field := range structType.Fields.List {
		validators := make([]validator.Validator, 0)

		// Apply markers to the field
		markerSet := markersInspect.FieldMarkers(field)

		// Traverse nested structs
		structType, ok := field.Type.(*ast.StructType)
		if !ok {
			validators = makeValidator(pass, markerSet, field, structName)
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
			validators = append(validators, makeValidator(pass, markerSet, field, structName)...)
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
		analyzed = append(analyzed, analyzeMarker(pass, markersInspect, structType, parentVariable, structName)...)
	}

	return analyzed
}

func makeValidator(pass *codegen.Pass, markers markers.MarkerSet, field *ast.Field, structName string) []validator.Validator {
	validators := make([]validator.Validator, 0)

	for _, marker := range markers {
		factory, err := registry.Validator(marker.Identifier)
		if err != nil {
			// Validator not found, skip
			continue
		}

		v := factory(pass, field, marker.Expressions, structName)
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
		"lower": strings.ToLower,
	}).Parse(ValidationTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, tmplData); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	src, err := format.Source(buf.Bytes())
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
