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
	"fmt"
	"go/ast"

	"github.com/sivchari/govalid/analyzers/markers"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	name = "govalid"
	doc  = "govalid generates type-safe validation code for structs based on markers."
)

// analyzer implements the analysis.Analyzer interface for govalid.
type analyzer struct{}

// newAnalyzer creates a new instance of the govalid analyzer.
func newAnalyzer() (*analysis.Analyzer, error) {
	a := &analyzer{}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      doc,
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer, markers.Analyzer},
	}, nil
}

// run is the main function that runs the govalid analyzer.
func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, fmt.Errorf("could not get inspector from pass result")
	}

	markersInspect, ok := pass.ResultOf[markers.Analyzer].(markers.Markers)
	if !ok {
		return nil, fmt.Errorf("could not get markers from pass result")
	}

	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		field, ok := n.(*ast.Field)
		if !ok {
			return
		}

		markers := markersInspect.FieldMarkers(field)
		for _, marker := range markers {
			fmt.Println(marker)
		}
	})

	return nil, nil //nolint:nilnil
}
