// Package main is the entry point for the govalid command line tool.
package main

import (
	"fmt"

	"github.com/gostaticanalysis/codegen/singlegenerator"

	"github.com/sivchari/govalid/analyzers/govalid"
	"github.com/sivchari/govalid/analyzers/markers"
	"github.com/sivchari/govalid/internal/registry"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// run initializes the analyzers and starts the unit checker.
func run() error {
	registry := registry.NewRegistry(
		registry.AddAnalyzers(markers.Initializer()),
		registry.AddGenerators(govalid.Initializer()),
	)

	if err := registry.Init(nil); err != nil {
		return fmt.Errorf("failed to initialize analyzers: %w", err)
	}

	govalid, err := registry.Generator(govalid.Name)
	if err != nil {
		return fmt.Errorf("failed to get govalid generator: %w", err)
	}

	singlegenerator.Main(govalid)

	return nil
}
