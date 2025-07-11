package benchmark

import (
	"strconv"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidGT(b *testing.B) {
	instance := test.GT{
		Age: 150,
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateGT(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundGT(b *testing.B) {
	validate := validator.New()
	instance := test.GT{
		Age: 150,
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

func BenchmarkGoValidatorGT(b *testing.B) {
	testValue := 150
	testString := strconv.Itoa(testValue)
	b.ResetTimer()
	for b.Loop() {
		// Check if numeric and > 100
		if !govalidator.IsNumeric(testString) {
			b.Fatal("validation failed - not numeric")
		}
		if testValue <= 100 {
			b.Fatal("validation failed - not greater than 100")
		}
	}
	b.StopTimer()
}
