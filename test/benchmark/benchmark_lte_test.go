package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidLTE(b *testing.B) {
	instance := test.LTE{Age: 75}
	for b.Loop() {
		err := test.ValidateLTE(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundLTE(b *testing.B) {
	validate := validator.New()
	instance := test.LTE{Age: 75}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}
