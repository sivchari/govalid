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
