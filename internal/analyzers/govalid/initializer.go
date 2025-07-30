package govalid

import (
	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/analyzers/registry"
	"github.com/sivchari/govalid/internal/config"
)

// Initializer returns a new instance of the initializer for the govalid analyzer.
func Initializer() registry.GeneratorInitializer {
	return &initializer{}
}

// initializer is a struct that implements the registry.AnalyzerInitializer interface.
type initializer struct{}

// Init initializes the govalid analyzer with the provided configuration.
func (i *initializer) Init(_ *config.GovalidConfig) (*codegen.Generator, error) {
	return newGenerator()
}

// Name returns the name of the govalid analyzer.
func (i *initializer) Name() string {
	return Name
}
