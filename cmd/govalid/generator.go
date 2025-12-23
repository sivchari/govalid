package main

import (
	"fmt"

	"github.com/gostaticanalysis/codegen/singlegenerator"

	"github.com/sivchari/govalid/internal/analyzers/govalid"
	"github.com/sivchari/govalid/internal/analyzers/markers"
	"github.com/sivchari/govalid/internal/analyzers/registry"
)

// runGenerator initializes the analyzers and starts the code generator.
func runGenerator() error {
	reg := registry.NewRegistry(
		registry.AddAnalyzers(markers.Initializer()),
		registry.AddGenerators(govalid.Initializer()),
	)

	if err := reg.Init(nil); err != nil {
		return fmt.Errorf("failed to initialize analyzers: %w", err)
	}

	gen, err := reg.Generator(govalid.Name)
	if err != nil {
		return fmt.Errorf("failed to get govalid generator: %w", err)
	}

	singlegenerator.Main(gen)

	return nil
}
