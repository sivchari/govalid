package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidMaxLength(b *testing.B) {
	instance := test.MaxLength{
		Name: "This is a test string for maxlength validation",
	}
	for b.Loop() {
		err := test.ValidateMaxLength(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundMaxLength(b *testing.B) {
	validate := validator.New()
	instance := test.MaxLength{
		Name: "This is a test string for maxlength validation",
	}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoValidatorMaxLength(b *testing.B) {
	testString := "This is a test string for maxlength validation"
	for b.Loop() {
		// Use StringLength function with max length 50
		if !govalidator.StringLength(testString, "0", "50") {
			b.Fatal("validation failed")
		}
	}
}

func BenchmarkOzzoValidationMaxLength(b *testing.B) {
	testString := "This is a test string for maxlength validation"
	for b.Loop() {
		err := validation.Validate(testString, validation.Length(0, 50))
		if err != nil {
			b.Fatal("validation failed:", err)
		}
	}
}

func BenchmarkGookitValidateMaxLength(b *testing.B) {
	testString := "This is a test string for maxlength validation"
	for b.Loop() {
		v := validate.New(map[string]any{"test": testString})
		v.StringRule("test", "max_len:50")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
}
