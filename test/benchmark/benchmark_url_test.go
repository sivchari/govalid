package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidURL(b *testing.B) {
	instance := test.URL{URL: "https://example.com/path?query=value"}
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
	instance := test.URL{URL: "https://example.com/path?query=value"}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}