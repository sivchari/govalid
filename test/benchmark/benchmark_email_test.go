package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidEmail(b *testing.B) {
	instance := test.Email{Email: "user@example.com"}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateEmail(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundEmail(b *testing.B) {
	validate := validator.New()
	instance := test.Email{Email: "user@example.com"}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}