package benchmark

import (
	"regexp"
	"testing"

	"github.com/asaskevich/govalidator"
	ozzovalidation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidAlpha(b *testing.B) {
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateAlpha(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundAlpha(b *testing.B) {
	validate := validator.New()
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkOzzoValidationAlpha(b *testing.B) {
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	alphaRegexp := regexp.MustCompile(`^[A-Za-z]+$`)
	b.ResetTimer()
	for b.Loop() {
		err := ozzovalidation.Validate(instance.FirstName, ozzovalidation.Match(alphaRegexp))
		if err != nil {
			b.Fatal("validation failed")
		}
		err = ozzovalidation.Validate(instance.LastName, ozzovalidation.Match(alphaRegexp))
		if err != nil {
			b.Fatal("validation failed")
		}
		err = ozzovalidation.Validate(instance.CountryCode, ozzovalidation.Match(alphaRegexp))
		if err != nil {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}

func BenchmarkAsaskevichGovalidatorAlpha(b *testing.B) {
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	b.ResetTimer()
	for b.Loop() {
		if !govalidator.IsAlpha(instance.FirstName) && !govalidator.IsAlpha(instance.LastName) && !govalidator.IsAlpha(instance.CountryCode) {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}

func BenchmarkGookitValidateAlpha(b *testing.B) {
	instance := test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}
	v := validate.Struct(instance)
	b.ResetTimer()
	for b.Loop() {
		if !v.Validate() {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}
