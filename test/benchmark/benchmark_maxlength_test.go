package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidMaxLength(b *testing.B) {
	instance := test.MaxLength{
		Name: "This is a test string for maxlength validation",
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateMaxLength(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundMaxLength(b *testing.B) {
	validate := validator.New()
	instance := test.MaxLength{
		Name: "This is a test string for maxlength validation",
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
