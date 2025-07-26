package benchmark

import (
	"regexp"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidAlphanum(b *testing.B) {
	instance := test.Alphanum{ProductCode: "product2024", SerialNumber: "ABC123", Username: "username"}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateAlphanum(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundAlphanum(b *testing.B) {
	validate := validator.New()
	instance := test.Alphanum{ProductCode: "product2024", SerialNumber: "ABC123", Username: "username"}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoValidatorAlphanum(b *testing.B) {
	validate := validator.New()
	instance := test.Alphanum{ProductCode: "product2024", SerialNumber: "ABC123", Username: "username"}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkOzzoValidationAlphanum(b *testing.B) {
	instance := test.Alphanum{ProductCode: "product2024", SerialNumber: "ABC123", Username: "username"}
	b.ResetTimer()
	for b.Loop() {
		err := validation.ValidateStruct(&instance,
			validation.Field(&instance.ProductCode, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]+$"))),
			validation.Field(&instance.SerialNumber, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]+$"))),
			validation.Field(&instance.Username, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]+$"))),
		)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGookitValidateAlphanum(b *testing.B) {
	instance := test.Alphanum{ProductCode: "product2024", SerialNumber: "ABC123", Username: "username"}
	b.ResetTimer()
	for b.Loop() {
		v := validate.New(map[string]any{"alphanum": instance})
		v.StringRule("alphanum", "alphanum")
	}
}
