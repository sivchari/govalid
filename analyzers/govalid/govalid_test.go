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

package govalid_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sivchari/govalid/analyzers/govalid"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	initializer := govalid.Initializer()

	a, err := initializer.Init(nil)
	if err != nil {
		t.Fatalf("failed to initialize analyzer: %v", err)
	}

	results := analysistest.Run(t, testdata, a, "a")
	for _, result := range results {
		fileName := fmt.Sprintf("%s.golden", result.Action.Package.Name)
		file, err := os.ReadFile(filepath.Join(testdata, fileName))
		if err != nil {
			t.Fatalf("failed to read golden file %s: %v", fileName, err)
		}
		if diff := cmp.Diff(string(file), result.Action); diff != "" {
			t.Errorf("output mismatch for package %s:\n%s", result.Action.Package.Name, diff)
		}
		fmt.Println(result.Action.Result)
	}
}
