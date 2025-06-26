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

package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func BenchmarkGovalidValidatorRequired(b *testing.B) {
	required := Required{
		Name:  "",
		Age:   0,
		Items: []string{},
	}
	b.ResetTimer()
	for b.Loop() {
		err := ValidateRequired(&required)
		if err == nil {
			b.Fatal("expected validation error, got nil")
		}
	}
	b.StopTimer()
}

func BenchmarkPlaygroundValidatorRequired(b *testing.B) {
	required := Required{
		Name:  "",
		Age:   0,
		Items: []string{},
	}
	validate := validator.New()
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(required)
		if err == nil {
			b.Fatal("expected validation error, got nil")
		}
	}
	b.StopTimer()
}
