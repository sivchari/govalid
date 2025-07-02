package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidGTE(b *testing.B) {
	instance := test.GTE{Age: 25}

	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateGTE(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundGTE(b *testing.B) {
	validate := validator.New()
	instance := test.GTE{Age: 25}

	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}