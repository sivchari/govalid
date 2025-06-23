/*
Copyright 2025 sivchari.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	"github.com/sivchari/govalid/analyzers/govalid/validator"
	"github.com/sivchari/govalid/analyzers/markers"
	govalidmarkers "github.com/sivchari/govalid/internal/markers"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	Name = "govalid"
	Doc  = "govalid generates type-safe validation code for structs based on markers."
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

type TemplateData struct {
	PackageName string
	TypeName    string
	Validators  []validator.Validator
}

// run is the main function that runs the govalid analyzer.
func (g *generator) run(pass *codegen.Pass) error {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return fmt.Errorf("could not get inspector from pass result")
	}

	markersInspect, ok := pass.ResultOf[markers.Analyzer].(markers.Markers)
	if !ok {
		return fmt.Errorf("could not get markers from pass result")
	}

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return
		}

		structType, ok := ts.Type.(*ast.StructType)
		if !ok {
			return
		}

		validators := make([]validator.Validator, 0)
		for _, field := range structType.Fields.List {
			markers := markersInspect.FieldMarkers(field)
			for _, marker := range markers {
				var v validator.Validator
				switch marker.Identifier {
				case govalidmarkers.GoValidMarkerRequired:
					v = validator.ValidateRequired(pass, field)
				default:
					continue
				}
				validators = append(validators, v)
			}
		}

		if len(validators) == 0 {
			return
		}

		tmplData := TemplateData{
			PackageName: pass.Pkg.Name(),
			TypeName:    ts.Name.Name,
			Validators:  validators,
		}

		t, err := template.New("validator").Parse(tmpl)
		if err != nil {
			panic(fmt.Errorf("failed to parse template: %w", err))
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, tmplData); err != nil {
			panic(fmt.Errorf("failed to execute template: %w", err))
		}

		src, err := format.Source(buf.Bytes())
		if err != nil {
			panic(fmt.Errorf("failed to format source code: %w", err))
		}

		if testing.Testing() {
			pass.Print(string(src))
			return
		}

		originalFilePath := pass.Fset.Position(ts.Pos()).Filename
		fileName := strings.TrimSuffix(filepath.Base(originalFilePath), filepath.Ext(originalFilePath))
		fileName = fmt.Sprintf("%s_validator.go", fileName)

		file, err := os.Create(fileName)
		if err != nil {
			panic(fmt.Errorf("failed to create file: %w", err))
		}
		defer file.Close()

		fmt.Fprint(file, string(src))

	})

	return nil //nolint:nilnil
}
