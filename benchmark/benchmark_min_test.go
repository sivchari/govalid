package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func BenchmarkGovalidValidatorMin(b *testing.B) {
	min := Min{
		Age: 9,
	}
	b.ResetTimer()
	for b.Loop() {
		err := ValidateMin(&min)
		if err == nil {
			b.Fatal("expected validation error, got nil")
		}
	}
	b.StopTimer()
}

func BenchmarkPlaygroundValidatorMin(b *testing.B) {
	min := Min{
		Age: 9,
	}
	validate := validator.New()
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(min)
		if err == nil {
			b.Fatal("expected validation error, got nil")
		}
	}
	b.StopTimer()
}
