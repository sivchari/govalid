package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func BenchmarkGovalidValidatorGT(b *testing.B) {
	gt := GT{
		Age: 100,
	}
	b.ResetTimer()
	for b.Loop() {
		err := ValidateGT(&gt)
		if err == nil {
			b.Fatal("expected validation error, got nil")
		}
	}
	b.StopTimer()
}

func BenchmarkPlaygroundValidatorGT(b *testing.B) {
	gt := GT{
		Age: 100,
	}
	validate := validator.New()
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&gt)
		if err == nil {
			b.Fatal("expected validation error, got nil")
		}
	}
	b.StopTimer()
}
