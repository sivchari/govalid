package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidMinLength(b *testing.B) {
	instance := test.MinLength{
		Name: "valid name",
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateMinLength(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundMinLength(b *testing.B) {
	validate := validator.New()
	instance := test.MinLength{
		Name: "valid name",
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