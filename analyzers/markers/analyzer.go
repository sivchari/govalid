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

package markers

import (
	"go/ast"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	govaliderrors "github.com/sivchari/govalid/internal/errors"
)

const (
	name = "markers"
	doc  = "markers is a helper for generating govalid validation"
)

// analyzer implements the analysis.Analyzer interface for the markers analyzer.
type analyzer struct{}

// newAnalyzer creates a new instance of the markers analyzer.
func newAnalyzer() *analysis.Analyzer {
	a := &analyzer{}

	return &analysis.Analyzer{
		Name:       name,
		Doc:        doc,
		Run:        a.run,
		Requires:   []*analysis.Analyzer{inspect.Analyzer},
		ResultType: reflect.TypeOf(newMarkers()),
		FactTypes: []analysis.Fact{
			(*MarkerFact)(nil),
		},
	}
}

// run is the main function that runs the markers analyzer.
func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, govaliderrors.ErrCouldNotGetInspector
	}

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	results, ok := newMarkers().(*markers)
	if !ok {
		return nil, govaliderrors.ErrCouldNotCreateMarkers
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			collectTypeSpecMarkers(pass, n, results)
		default:
		}
	})

	return results, nil
}

// collectTypeSpecMarkers collects markers from a TypeSpec node and adds them to the results.
func collectTypeSpecMarkers(pass *analysis.Pass, ts *ast.TypeSpec, result *markers) {
	if ts == nil {
		return
	}

	st, ok := ts.Type.(*ast.StructType)
	if !ok {
		return
	}

	for _, field := range st.Fields.List {
		fieldMarkers(pass, field, result)
	}
}

// fieldMarkers extracts markers from a struct field and adds them to the results.
func fieldMarkers(pass *analysis.Pass, field *ast.Field, results *markers) {
	for _, doc := range field.Doc.List {
		if !strings.HasPrefix(doc.Text, "// +") {
			continue
		}

		markerContent := strings.TrimPrefix(doc.Text, "// +")

		identifier, expressions := extractMarker(markerContent)
		mset := newMarkerSet()
		mset[identifier] = Marker{
			Identifier:  identifier,
			Expressions: expressions,
		}
		results.insertFieldMarkers(field, mset)

		if obj, ok := pass.TypesInfo.Defs[field.Names[0]]; ok {
			pass.ExportObjectFact(obj, &MarkerFact{
				Identifier:  identifier,
				Expressions: expressions,
			})
		}
	}
}

// extractMarker extracts the identifier and expressions from a marker content string.
// It returns the identifier and a map of expressions if applicable.
// If the content does not contain an identifier or expressions, it returns an empty string and nil.
func extractMarker(content string) (string, map[string]string) {
	if strings.Count(content, "=") == 0 {
		return content, nil
	}

	if strings.Count(content, "=") == 1 {
		splits := strings.SplitN(content, "=", 2)
		if len(splits) != 2 {
			return "", nil
		}

		expressions := map[string]string{}
		expressions[splits[0]] = splits[1]

		return splits[0], expressions
	}

	return "", nil
}
