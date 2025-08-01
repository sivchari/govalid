package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidLength(b *testing.B) {
	instance := test.Length{
		Name: "1234567",
	}
	for b.Loop() {
		err := test.ValidateLength(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundLength(b *testing.B) {
	validate := validator.New()
	instance := test.Length{
		Name: "1234567",
	}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoValidatorLength(b *testing.B) {
	testString := "1234567"
	for b.Loop() {
		if !govalidator.StringLength(testString, "7", "7") {
			b.Fatal("validation failed")
		}
	}
}

func BenchmarkGookitValidateLength(b *testing.B) {
	testString := "1234567"
	for b.Loop() {
		v := validate.New(map[string]any{"name": testString})
		v.StringRule("name", "rune_len:7")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
}
