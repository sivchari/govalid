package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidRequired(b *testing.B) {
	instance := test.Required{
		Name:  "test",
		Age:   1,
		Items: []string{"test"},
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateRequired(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundRequired(b *testing.B) {
	validate := validator.New()
	instance := test.Required{
		Name:  "test",
		Age:   1,
		Items: []string{"test"},
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

func BenchmarkGoValidatorRequired(b *testing.B) {
	testString := "test"
	b.ResetTimer()
	for b.Loop() {
		// Check if string is not empty
		if len(testString) == 0 {
			b.Fatal("validation failed")
		}
		// Or use govalidator.IsNull for empty check
		if govalidator.IsNull(testString) {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}

func BenchmarkOzzoValidationRequired(b *testing.B) {
	testString := "test"
	b.ResetTimer()
	for b.Loop() {
		err := validation.Validate(testString, validation.Required)
		if err != nil {
			b.Fatal("validation failed:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGookitValidateRequired(b *testing.B) {
	testString := "test"
	b.ResetTimer()
	for b.Loop() {
		v := validate.New(map[string]any{"test": testString})
		v.StringRule("test", "required")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
	b.StopTimer()
}
