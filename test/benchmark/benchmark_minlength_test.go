package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidMinLength(b *testing.B) {
	instance := test.MinLength{Name: "test string with adequate length"}

	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateMinLength(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundMinLength(b *testing.B) {
	validate := validator.New()
	instance := test.MinLength{Name: "test string with adequate length"}

	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoValidatorMinLength(b *testing.B) {
	testString := "test string with adequate length"
	b.ResetTimer()
	for b.Loop() {
		// Use StringLength function with min length 3, max length 50
		if !govalidator.StringLength(testString, "3", "50") {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}
