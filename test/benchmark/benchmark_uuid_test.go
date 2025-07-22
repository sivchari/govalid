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

func BenchmarkGoValidUUID(b *testing.B) {
	instance := test.UUID{UUID: "6ba7b811-9dad-41d1-80b4-00c04fd430c8"}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateUUID(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundUUID(b *testing.B) {
	validate := validator.New()
	instance := test.UUID{UUID: "6ba7b811-9dad-41d1-80b4-00c04fd430c8"}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoValidatorUUID(b *testing.B) {
	testUUID := "6ba7b811-9dad-41d1-80b4-00c04fd430c8"
	b.ResetTimer()
	for b.Loop() {
		if !govalidator.IsUUID(testUUID) {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}

func BenchmarkOzzoValidationUUID(b *testing.B) {
	testUUID := "6ba7b811-9dad-41d1-80b4-00c04fd430c8"
	b.ResetTimer()
	for b.Loop() {
		err := validation.Validate(testUUID, is.UUID)
		if err != nil {
			b.Fatal("validation failed:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGookitValidateUUID(b *testing.B) {
	testUUID := "6ba7b811-9dad-41d1-80b4-00c04fd430c8"
	b.ResetTimer()
	for b.Loop() {
		v := validate.New(map[string]any{"uuid": testUUID})
		v.StringRule("uuid", "uuid")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
	b.StopTimer()
}
