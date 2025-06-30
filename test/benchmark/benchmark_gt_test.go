package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidGT(b *testing.B) {
	instance := test.GT{
		Age: 150,
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateGT(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundGT(b *testing.B) {
	validate := validator.New()
	instance := test.GT{
		Age: 150,
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
