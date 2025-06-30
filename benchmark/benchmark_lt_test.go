
package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func BenchmarkGovalidValidatorLT(b *testing.B) {
	lt := LT{
		Age: 10,
	}
	b.ResetTimer()
	for b.Loop() {
		err := ValidateLT(&lt)
		if err == nil {
			b.Fatal("expected validation error, got nil")
		}
	}
	b.StopTimer()
}

func BenchmarkPlaygroundValidatorLT(b *testing.B) {
	lt := LT{
		Age: 10,
	}
	validate := validator.New()
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(lt)
		if err == nil {
			b.Fatal("expected validation error, got nil")
		}
	}
	b.StopTimer()
}
