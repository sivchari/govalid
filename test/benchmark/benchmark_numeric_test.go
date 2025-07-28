package benchmark

import (
	"regexp"
	"testing"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidNumeric(b *testing.B) {
	instance := test.Numeric{Number: "123456"}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateNumeric(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundNumeric(b *testing.B) {
	validate := validator.New()
	validate.RegisterValidation("numeric", func(fl validator.FieldLevel) bool {
		str := fl.Field().String()
		for i := 0; i < len(str); i++ {
			if str[i] < '0' || str[i] > '9' {
				return false
			}
		}
		return true
	})

	instance := struct {
		Number string `validate:"numeric"`
	}{Number: "123456"}

	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoValidatorNumeric(b *testing.B) {
	testNumeric := "123456"
	b.ResetTimer()
	for b.Loop() {
		if !govalidator.IsNumeric(testNumeric) {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}

func BenchmarkOzzoValidationNumeric(b *testing.B) {
	testNumeric := "123456"
	re := regexp.MustCompile(`^\d+$`)
	b.ResetTimer()
	for b.Loop() {
		err := validation.Validate(testNumeric, validation.Match(re))
		if err != nil {
			b.Fatal("validation failed:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGookitValidateNumeric(b *testing.B) {
	testNumeric := "123456"
	b.ResetTimer()
	for b.Loop() {
		v := validate.New(map[string]any{"number": testNumeric})
		v.StringRule("number", "numeric")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
	b.StopTimer()
}
