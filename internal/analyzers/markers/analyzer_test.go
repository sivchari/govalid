package markers_test

import (
	"testing"

	"github.com/sivchari/govalid/internal/analyzers/markers"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	initializer := markers.Initializer()

	a, err := initializer.Init(nil)
	if err != nil {
		t.Fatalf("failed to initialize analyzer: %v", err)
	}

	analysistest.Run(t, testdata, a, "a")
}
