package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

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

func BenchmarkGoValidatorLength(b *testing.B) {
	testString := "1234567"
	b.ResetTimer()
	for b.Loop() {
		if !govalidator.StringLength(testString, "7", "7") {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}

func BenchmarkOzzoValidationLength(b *testing.B) {
	testString := "1234567"
	b.ResetTimer()
	for b.Loop() {
		err := validation.Validate(testString, validation.Length(7, 7))
		if err != nil {
			b.Fatal("validation failed:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGookitValidateLength(b *testing.B) {
	testString := "1234567"
	b.ResetTimer()
	for b.Loop() {
		v := validate.New(map[string]any{"name": testString})
		v.StringRule("name", "rune_len:7")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
	b.StopTimer()
}
