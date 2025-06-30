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

	"github.com/sivchari/govalid/analyzers/markers"
	govaliderrors "github.com/sivchari/govalid/internal/errors"
	govalidmarkers "github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/rules"
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
	PackageName string
	TypeName    string
	Metadata    []*AnalyzedMetadata
	NeedsUTF8   bool
}

// run is the main function that runs the govalid analyzer.
func (g *generator) run(pass *codegen.Pass) error {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
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

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return
		}

		structType, ok := ts.Type.(*ast.StructType)
		if !ok {
			return
		}

		metadata := analyzeMarker(pass, markersInspect, structType, "")
		if len(metadata) == 0 {
			return
		}

		tmplData := TemplateData{
			PackageName: pass.Pkg.Name(),
			TypeName:    ts.Name.Name,
			Metadata:    metadata,
			NeedsUTF8:   needsUTF8(metadata),
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

func analyzeMarker(pass *codegen.Pass, markersInspect markers.Markers, structType *ast.StructType, parent string) []*AnalyzedMetadata {
	analyzed := make([]*AnalyzedMetadata, 0)

	for _, field := range structType.Fields.List {
		validators := make([]validator.Validator, 0)

		// Apply markers to the field
		markers := markersInspect.FieldMarkers(field)

		// Traverse nested structs
		structType, ok := field.Type.(*ast.StructType)
		if !ok {
			validators = makeValidator(pass, markers, field)
			if len(validators) == 0 {
				continue
			}

			analyzed = append(analyzed, &AnalyzedMetadata{
				Validators: validators,
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
			validators = append(validators, makeValidator(pass, markers, field)...)
		}

		// Add the parent variable name to the analyzed metadata
		var parentVariable string
		if parent != "" {
			parentVariable = fmt.Sprintf("%s.%s", parent, field.Names[0].Name)
		} else {
			parentVariable = field.Names[0].Name
		}

		analyzed = append(analyzed, &AnalyzedMetadata{
			Validators:     validators,
			ParentVariable: parentVariable,
		})

		// Recursively analyze nested structs
		analyzed = append(analyzed, analyzeMarker(pass, markersInspect, structType, parentVariable)...)
	}

	return analyzed
}

func makeValidator(pass *codegen.Pass, markers markers.MarkerSet, field *ast.Field) []validator.Validator {
	validators := make([]validator.Validator, 0)

	for _, marker := range markers {
		var v validator.Validator

		switch marker.Identifier {
		case govalidmarkers.GoValidMarkerRequired:
			v = rules.ValidateRequired(pass, field)
		case govalidmarkers.GoValidMarkerLT:
			v = rules.ValidateLT(pass, field, marker.Expressions)
		case govalidmarkers.GoValidMarkerGT:
			v = rules.ValidateGT(pass, field, marker.Expressions)
		case govalidmarkers.GoValidMarkerMaxLength:
			v = rules.ValidateMaxLength(pass, field, marker.Expressions)
		default:
			continue
		}

		if v == nil {
			continue
		}

		validators = append(validators, v)
	}

	return validators
}

// needsUTF8 checks if any validator requires the unicode/utf8 package
func needsUTF8(metadata []*AnalyzedMetadata) bool {
	for _, meta := range metadata {
		for _, validator := range meta.Validators {
			if strings.Contains(validator.Validate(), "utf8.RuneCountInString") {
				return true
			}
		}
	}
	return false
}

func writeFile(pass *codegen.Pass, ts *ast.TypeSpec, tmplData TemplateData) error {
	t, err := template.New("validator").Funcs(template.FuncMap{
		"trimDots": func(s string) string {
			return strings.ReplaceAll(s, ".", "")
		},
	}).Parse(tmpl)
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

	file, err := os.Create(fileName) //nolint:gosec
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() //nolint:errcheck

	fmt.Fprint(file, string(src)) //nolint:errcheck

	return nil
}
