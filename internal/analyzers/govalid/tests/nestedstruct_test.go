package tests

import (
	"path/filepath"
	"testing"

	"github.com/gostaticanalysis/codegen/codegentest"
	"github.com/sivchari/golden"

	"github.com/sivchari/govalid/internal/analyzers/govalid"
	"github.com/sivchari/govalid/internal/analyzers/markers"
	"github.com/sivchari/govalid/internal/analyzers/registry"
)

func TestNestedStruct(t *testing.T) {
	registry := registry.NewRegistry(
		registry.AddAnalyzers(markers.Initializer()),
		registry.AddGenerators(govalid.Initializer()),
	)

	if err := registry.Init(nil); err != nil {
		t.Fatalf("failed to initialize analyzers: %v", err)
	}

	govalid, err := registry.Generator(govalid.Name)
	if err != nil {
		t.Fatalf("failed to get govalid generator: %v", err)
	}

	testCases := []string{
		"nestedstruct/nop",
		"nestedstruct/inside",
		"nestedstruct/partial",
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			results := codegentest.Run(t, codegentest.TestData(), govalid, tc)
			g := golden.New(t, golden.WithUpdate(update))
			name := filepath.Base(tc)

			for _, r := range results {
				g.Assert(name, r.Output.String())
			}
		})
	}
}
