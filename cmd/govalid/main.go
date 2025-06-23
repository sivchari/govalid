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
