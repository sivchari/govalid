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
	"golang.org/x/text/language"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/imports"

	"github.com/sivchari/govalid/internal/analyzers/markers"
	govaliderrors "github.com/sivchari/govalid/internal/errors"
	"github.com/sivchari/govalid/internal/validator"
	"github.com/sivchari/govalid/internal/validator/expr"
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

	// i18nPath is the path to the YAML i18n configuration, set by the -i18n flag.
	i18nPath string
)

// generator is the main type for the govalid analyzer.
type generator struct {
	// i18n holds the loaded i18n configuration, or nil when -i18n is not set.
	i18n *i18nConfig
}

// newGenerator creates a new instance of the govalid generator.
func newGenerator() (*codegen.Generator, error) {
	g := &generator{}

	generator := &codegen.Generator{
		Name:     Name,
		Doc:      Doc,
		Run:      g.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer, markers.Analyzer},
	}

	generator.Flags.StringVar(&i18nPath, "i18n", "", "path to a YAML file with localized error message templates")

	return generator, nil
}

// TemplateData holds the data for the template used to generate validation code.
type TemplateData struct {
	PackageName     string
	TypeName        string
	Metadata        []*AnalyzedMetadata
	ImportPackages  map[string]struct{}
	ErrDeclarations []validator.ErrDecl
	// I18n is true when at least one error declaration has localized messages,
	// in which case the generated validator localizes errors via govalid.Localize.
	I18n bool
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

	if err := g.ensureI18n(); err != nil {
		return err
	}

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			return
		}

		for _, spec := range genDecl.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			g.generateType(pass, markersInspect, ts)
		}
	})

	return nil
}

// generateType analyzes a single type spec and writes its validator file, if any
// markers are present.
func (g *generator) generateType(pass *codegen.Pass, markersInspect markers.Markers, ts *ast.TypeSpec) {
	structType, ok := ts.Type.(*ast.StructType)
	if !ok {
		return
	}

	typeMarkers := markersInspect.TypeMarkers(ts)

	metadata := analyzeMarker(pass, markersInspect, typeMarkers, structType, "", ts.Name.Name)
	if len(metadata) == 0 {
		return
	}

	errDecls := collectErrDeclarations(metadata, g.i18n)

	tmplData := TemplateData{
		PackageName:     pass.Pkg.Name(),
		TypeName:        ts.Name.Name,
		Metadata:        metadata,
		ImportPackages:  collectImportPackages(metadata),
		ErrDeclarations: errDecls,
		I18n:            hasLocalizedMessages(errDecls),
	}

	if err := writeFile(pass, ts, &tmplData); err != nil {
		panic(fmt.Sprintf("failed to write file for %s: %v", ts.Name.Name, err))
	}
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

//nolint:funlen // This function is complex but cohesive - it handles complete field analysis including nested structs
func analyzeMarker(pass *codegen.Pass, markersInspect markers.Markers, typeMarkers markers.MarkerSet, structType *ast.StructType, parent, structName string) []*AnalyzedMetadata {
	analyzed := make([]*AnalyzedMetadata, 0)

	typeMarkersList := make([]markers.Marker, 0, len(typeMarkers))
	for _, marker := range typeMarkers {
		typeMarkersList = append(typeMarkersList, marker)
	}

	sort.SliceStable(typeMarkersList, func(i, j int) bool {
		return typeMarkersList[i].Identifier < typeMarkersList[j].Identifier
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
			return fieldMarkersList[i].Identifier < fieldMarkersList[j].Identifier
		})

		markersList := make([]markers.Marker, 0, len(typeMarkersList)+len(fieldMarkersList))
		markersList = append(markersList, typeMarkersList...)
		markersList = append(markersList, fieldMarkersList...)

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

				//govalid:required
				type Nested struct {
					Name string `json:"name"`
				}
			*/
			input.Field = field
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
		for _, v := range meta.Validators {
			cond := v.Condition()
			if cond == nil {
				continue
			}

			for _, pkg := range cond.Imports {
				packages[pkg] = struct{}{}
			}
		}
	}

	return packages
}

// collectErrDeclarations collects deduplicated error declarations from all validators.
// When i18n is non-nil, each declaration is populated with localized messages.
func collectErrDeclarations(metadata []*AnalyzedMetadata, i18n *i18nConfig) []validator.ErrDecl {
	seen := make(map[string]struct{})

	var decls []validator.ErrDecl

	for _, meta := range metadata {
		for _, v := range meta.Validators {
			cond := v.Condition()
			if cond == nil {
				continue
			}

			errDecl := v.ErrDecl()

			if _, ok := seen[errDecl.VarName]; ok {
				continue
			}

			seen[errDecl.VarName] = struct{}{}

			if i18n != nil {
				errDecl.Messages = i18n.build(&errDecl)
			}

			decls = append(decls, errDecl)
		}
	}

	return decls
}

// hasLocalizedMessages reports whether any declaration has localized messages.
func hasLocalizedMessages(decls []validator.ErrDecl) bool {
	for i := range decls {
		if len(decls[i].Messages) > 0 {
			return true
		}
	}

	return false
}

// templateFuncMap returns the template.FuncMap used for generating validation code.
func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"trimDots": func(s string) string {
			return strings.ReplaceAll(s, ".", "")
		},
		"hasCondition": func(v validator.Validator) bool {
			return v.Condition() != nil
		},
		"renderCondition": func(v validator.Validator) string {
			cond := v.Condition()
			if cond == nil {
				return ""
			}

			if cond.IfInitStmt != nil {
				return expr.RenderStmt(cond.IfInitStmt) + "; " + expr.Render(cond.Expr)
			}

			return expr.Render(cond.Expr)
		},
		"errVarName": func(v validator.Validator) string {
			return v.ErrDecl().VarName
		},
		"sortedMessages": sortedMessages,
	}
}

// langMessage is a single localized message rendered by the template.
type langMessage struct {
	Lang string
	Msg  string
}

// sortedMessages converts a language-keyed message map into a slice sorted by
// language tag, so generated output is deterministic.
func sortedMessages(m map[language.Tag]string) []langMessage {
	out := make([]langMessage, 0, len(m))
	for tag, msg := range m {
		out = append(out, langMessage{Lang: tag.String(), Msg: msg})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Lang < out[j].Lang
	})

	return out
}

func writeFile(pass *codegen.Pass, ts *ast.TypeSpec, tmplData *TemplateData) error {
	t, err := template.New("validator").Funcs(templateFuncMap()).Parse(ValidationTemplate)
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

	// Get the directory of the original file to place the validator in the same directory
	dir := filepath.Dir(originalFilePath)
	outputPath := filepath.Join(dir, fileName)

	file, err := os.Create(filepath.Clean(outputPath))
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
