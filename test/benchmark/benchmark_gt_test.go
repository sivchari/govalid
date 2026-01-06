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
		Age:    150,
		Scores: []int{120, 150},
	}
	for b.Loop() {
		err := test.ValidateGT(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundGT(b *testing.B) {
	validate := validator.New()
	instance := test.GT{
		Age:    150,
		Scores: []int{120, 150},
	}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoValidatorGT(b *testing.B) {
	testValue := 150
	testString := strconv.Itoa(testValue)
	scores := []int{120, 150}
	for b.Loop() {
		// Check if numeric and > 100
		if !govalidator.IsNumeric(testString) {
			b.Fatal("validation failed - not numeric")
		}
		if testValue <= 100 {
			b.Fatal("validation failed - not greater than 100")
			// Check each value is > 100
			for _, v := range scores {
				if v <= 100 {
					b.Fatal("validation failed - not greater than 100")
				}
			}
		}
	}
}
