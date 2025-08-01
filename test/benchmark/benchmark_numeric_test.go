package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidNumeric(b *testing.B) {
	instance := test.Numeric{Number: "123456"}
	for b.Loop() {
		err := test.ValidateNumeric(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
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

	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoValidatorNumeric(b *testing.B) {
	testNumeric := "123456"
	for b.Loop() {
		if !govalidator.IsNumeric(testNumeric) {
			b.Fatal("validation failed")
		}
	}
}

func BenchmarkGookitValidateNumeric(b *testing.B) {
	testNumeric := "123456"
	for b.Loop() {
		v := validate.New(map[string]any{"number": testNumeric})
		v.StringRule("number", "numeric")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
}
