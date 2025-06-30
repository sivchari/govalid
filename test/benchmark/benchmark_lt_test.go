package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidLT(b *testing.B) {
	instance := test.LT{
		Age: 5,
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateLT(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundLT(b *testing.B) {
	validate := validator.New()
	instance := test.LT{
		Age: 5,
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
