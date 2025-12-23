package markers_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/sivchari/govalid/internal/analyzers/markers"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	initializer := markers.Initializer()

	a, err := initializer.Init(nil)
	if err != nil {
		t.Fatalf("failed to initialize analyzer: %v", err)
	}

	// Test new format (//govalid:)
	analysistest.Run(t, testdata, a, "a")
}

func TestOldFormat(t *testing.T) {
	testdata := analysistest.TestData()

	initializer := markers.Initializer()

	a, err := initializer.Init(nil)
	if err != nil {
		t.Fatalf("failed to initialize analyzer: %v", err)
	}

	// Test old format (// +govalid:) for backward compatibility
	analysistest.Run(t, testdata, a, "b")
}
