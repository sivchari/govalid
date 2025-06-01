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
	"github.com/sivchari/govalid/internal/config"
	"github.com/sivchari/govalid/internal/registry"
	"golang.org/x/tools/go/analysis"
)

// Initializer returns a new instance of the initializer for the govalid analyzer.
func Initializer() registry.AnalyzerInitializer {
	return &initializer{}
}

// initializer is a struct that implements the registry.AnalyzerInitializer interface.
type initializer struct{}

// Init initializes the govalid analyzer with the provided configuration.
func (i *initializer) Init(_ *config.GovalidConfig) (*analysis.Analyzer, error) {
	return newAnalyzer()
}

// Name returns the name of the govalid analyzer.
func (i *initializer) Name() string {
	return name
}
