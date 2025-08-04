package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidAlpha(b *testing.B) {
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	for b.Loop() {
		err := test.ValidateAlpha(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundAlpha(b *testing.B) {
	validate := validator.New()
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkAsaskevichGovalidatorAlpha(b *testing.B) {
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	for b.Loop() {
		if !govalidator.IsAlpha(instance.FirstName) && !govalidator.IsAlpha(instance.LastName) && !govalidator.IsAlpha(instance.CountryCode) {
			b.Fatal("validation failed")
		}
	}
}

func BenchmarkGookitValidateAlpha(b *testing.B) {
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	v := validate.Struct(instance)
	for b.Loop() {
		if !v.Validate() {
			b.Fatal("validation failed")
		}
	}
}
