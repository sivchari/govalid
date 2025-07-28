package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidEmail(b *testing.B) {
	instance := test.Email{Email: "user@example.com"}
	for b.Loop() {
		err := test.ValidateEmail(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundEmail(b *testing.B) {
	validate := validator.New()
	instance := test.Email{Email: "user@example.com"}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoValidatorEmail(b *testing.B) {
	testEmail := "user@example.com"
	for b.Loop() {
		if !govalidator.IsEmail(testEmail) {
			b.Fatal("validation failed")
		}
	}
}

func BenchmarkOzzoValidationEmail(b *testing.B) {
	testEmail := "user@example.com"
	for b.Loop() {
		err := validation.Validate(testEmail, is.Email)
		if err != nil {
			b.Fatal("validation failed:", err)
		}
	}
}

func BenchmarkGookitValidateEmail(b *testing.B) {
	testEmail := "user@example.com"
	for b.Loop() {
		v := validate.New(map[string]any{"email": testEmail})
		v.StringRule("email", "email")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
}
