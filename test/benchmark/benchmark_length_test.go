package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidLength(b *testing.B) {
	instance := test.Length{
		Name: "1234567",
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateLength(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundLength(b *testing.B) {
	validate := validator.New()
	instance := test.Length{
		Name: "1234567",
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
