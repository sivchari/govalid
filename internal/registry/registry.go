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

	"github.com/gostaticanalysis/codegen"
	"golang.org/x/tools/go/analysis"

	"github.com/sivchari/govalid/internal/config"
)

// AnalyzerInitializer is an interface for initializing analyzers.
type AnalyzerInitializer interface {
	// Name returns the name of the initialized analyzer.
	Name() string

	// Init initializes the analyzer.
	Init(*config.GovalidConfig) (*analysis.Analyzer, error)
}

// GeneratorInitializer is an interface for initializing generators.
type GeneratorInitializer interface {
	// Name returns the name of the initialized generator.
	Name() string

	// Init initializes the generator.
	Init(*config.GovalidConfig) (*codegen.Generator, error)
}

// Registry is an interface for managing a collection of analyzers and generators.
type Registry interface {
	// Analyzers returns a slice of initialized analyzers.
	Analyzers() []string

	// Analyzer returns an analysis.Analyzer by name.
	Analyzer(name string) (*analysis.Analyzer, error)

	// Generators returns a slice of initialized generators.
	Generators() []string

	// Generator returns a codegen.Generator by name.
	Generator(name string) (*codegen.Generator, error)

	// Init initializes all analyzers and generators.
	Init(*config.GovalidConfig) error
}

// registry is an implementation of the Registry interface.
type registry struct {
	analyzerInitializers  []AnalyzerInitializer
	generatorInitializers []GeneratorInitializer

	initializedAnalyzers  []*analysis.Analyzer
	initializedGenerators []*codegen.Generator
}

// NewRegistry creates a new Registry with the provided analyzer initializers.
func NewRegistry(builders ...Builder) Registry {
	r := &registry{}
	for _, builder := range builders {
		builder(r)
	}

	return r
}

// Builder is a function that takes a Registry and registers analyzers and generators.
type Builder func(registry Registry)

// AddAnalyzers adds a slice of AnalyzerInitializers to the registry.
func AddAnalyzers(analyzers ...AnalyzerInitializer) Builder {
	return func(r Registry) {
		reg, ok := r.(*registry)
		if !ok {
			panic("AddAnalyzers: registry is not of type *registry")
		}

		reg.analyzerInitializers = append(reg.analyzerInitializers, analyzers...)
	}
}

// AddGenerators adds a slice of GeneratorInitializer to the registry.
func AddGenerators(generators ...GeneratorInitializer) Builder {
	return func(r Registry) {
		reg, ok := r.(*registry)
		if !ok {
			panic("AddGenerators: registry is not of type *registry")
		}

		reg.generatorInitializers = append(reg.generatorInitializers, generators...)
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

// Generators returns a slice of names of all generators in the registry.
func (r *registry) Generators() []string {
	generators := make([]string, len(r.generatorInitializers))
	for i, initializer := range r.generatorInitializers {
		generators[i] = initializer.Name()
	}

	return generators
}

func (r *registry) Analyzer(name string) (*analysis.Analyzer, error) {
	for _, analyzer := range r.initializedAnalyzers {
		if analyzer.Name == name {
			return analyzer, nil
		}
	}

	return nil, fmt.Errorf("analyzer %s not found in registry", name)
}

func (r *registry) Generator(name string) (*codegen.Generator, error) {
	for _, generator := range r.initializedGenerators {
		if generator.Name == name {
			return generator, nil
		}
	}

	return nil, fmt.Errorf("generator %s not found in registry", name)
}

// Init initializes all analyzers in the registry and returns a slice of pointers to analysis.Analyzer.
func (r *registry) Init(config *config.GovalidConfig) error {
	for _, initializer := range r.analyzerInitializers {
		analyzer, err := initializer.Init(config)
		if err != nil {
			return fmt.Errorf("failed to initialize analyzer %s: %w", initializer.Name(), err)
		}

		r.initializedAnalyzers = append(r.initializedAnalyzers, analyzer)
	}

	for _, initializer := range r.generatorInitializers {
		generator, err := initializer.Init(config)
		if err != nil {
			return fmt.Errorf("failed to initialize generator %s: %w", initializer.Name(), err)
		}

		r.initializedGenerators = append(r.initializedGenerators, generator)
	}

	return nil
}
