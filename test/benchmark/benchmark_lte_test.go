package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidLTE(b *testing.B) {
	instance := test.LTE{Age: 75}

	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateLTE(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundLTE(b *testing.B) {
	validate := validator.New()
	instance := test.LTE{Age: 75}

	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}