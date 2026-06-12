package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidAlphanum(b *testing.B) {
	instance := test.Alphanum{ProductCode: "ABC123", SerialNumber: "product2024", Username: "X1"}
	for b.Loop() {
		err := test.ValidateAlphanum(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundAlphanum(b *testing.B) {
	validate := validator.New()
	instance := test.Alphanum{ProductCode: "ABC123", SerialNumber: "product2024", Username: "X1"}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkAsaskevichGovalidatorAlphanum(b *testing.B) {
	instance := test.Alphanum{ProductCode: "ABC123", SerialNumber: "product2024", Username: "X1"}
	for b.Loop() {
		if !govalidator.IsAlphanumeric(instance.ProductCode) && !govalidator.IsAlphanumeric(instance.SerialNumber) && !govalidator.IsAlphanumeric(instance.Username) {
			b.Fatal("validation failed")
		}
	}
}
