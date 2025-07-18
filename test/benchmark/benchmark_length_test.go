package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidLength(b *testing.B) {
	instance := test.Length{
		PostalCode:  "1234567",
		PhoneNumber: "1234567890",
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
		PostalCode:  "1234567",
		PhoneNumber: "1234567890",
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
