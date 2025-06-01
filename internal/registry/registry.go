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

// Package registry implements registry for analyzers.
package registry

import (
	"fmt"

	"github.com/sivchari/govalid/internal/config"
	"golang.org/x/tools/go/analysis"
)

// AnalyzerInitializer is an interface for initializing analyzers.
type AnalyzerInitializer interface {
	// Name returns the name of the initialized analyzer.
	Name() string

	// Init initializes the analyzer.
	Init(*config.GovalidConfig) (*analysis.Analyzer, error)
}

// Registry is an interface for managing a collection of analyzers.
type Registry interface {
	// Analyzers returns a slice of initialized analyzers.
	Analyzers() []string

	// Init initializes all analyzers in the registry.
	Init(*config.GovalidConfig) ([]*analysis.Analyzer, error)
}

// registry is an implementation of the Registry interface.
type registry struct {
	analyzerInitializers []AnalyzerInitializer
}

// NewRegistry creates a new Registry with the provided analyzer initializers.
func NewRegistry(initializers ...AnalyzerInitializer) Registry {
	return &registry{
		analyzerInitializers: initializers,
	}
}

// Analyzers returns a slice of names of all analyzers in the registry.
func (r *registry) Analyzers() []string {
	analyzers := make([]string, len(r.analyzerInitializers))
	for i, initializer := range r.analyzerInitializers {
		analyzers[i] = initializer.Name()
	}

	return analyzers
}

// Init initializes all analyzers in the registry and returns a slice of pointers to analysis.Analyzer.
func (r *registry) Init(config *config.GovalidConfig) ([]*analysis.Analyzer, error) {
	analyzers := make([]*analysis.Analyzer, 0, len(r.analyzerInitializers))

	for _, initializer := range r.analyzerInitializers {
		analyzer, err := initializer.Init(config)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize analyzer %s: %w", initializer.Name(), err)
		}

		analyzers = append(analyzers, analyzer)
	}

	return analyzers, nil
}
