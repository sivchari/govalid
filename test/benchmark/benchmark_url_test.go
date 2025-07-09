package benchmark_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidURL(b *testing.B) {
	instance := test.URL{
		URL: "https://example.com/path/to/resource?key=value",
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateURL(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundURL(b *testing.B) {
	validate := validator.New()
	instance := test.URL{
		URL: "https://example.com/path/to/resource?key=value",
	}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}
