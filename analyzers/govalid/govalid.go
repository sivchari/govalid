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

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"

	"github.com/sivchari/govalid/analyzers/markers"
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

	markersInitializer := markers.Initializer()

	markersAnalyzer, err := markersInitializer.Init(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize markers analyzer: %w", err)
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      doc,
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer, markersAnalyzer},
	}, nil
}

// run is the main function that runs the govalid analyzer.
func (a *analyzer) run(_ *analysis.Pass) (any, error) {
	return nil, nil //nolint:nilnil
}
