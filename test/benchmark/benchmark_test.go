package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

// Test data with valid values
type RequiredValidData test.TestRequired
type LTValidData test.TestLT
type GTValidData test.TestGT
type MaxLengthValidData test.TestMaxLength

// Test data with invalid values
type RequiredInvalidData test.TestRequired
type LTInvalidData test.TestLT
type GTInvalidData test.TestGT
type MaxLengthInvalidData test.TestMaxLength

var (
	// Valid test data
	requiredValid     = RequiredValidData{Name: "John Doe", Age: 25, Items: []string{"item1", "item2"}}
	ltValid          = LTValidData{Age: 5}
	gtValid          = GTValidData{Age: 150}
	maxLengthValid   = MaxLengthValidData{Name: "Valid Name"}

	// Invalid test data
	requiredInvalid   = RequiredInvalidData{Name: "", Age: 0, Items: nil}
	ltInvalid        = LTInvalidData{Age: 15}
	gtInvalid        = GTInvalidData{Age: 50}
	maxLengthInvalid = MaxLengthInvalidData{Name: "This is a very long name that exceeds the maximum allowed length of 50 characters"}
)

// Required validation benchmarks
func BenchmarkGoValidRequired(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test.ValidateRequired((*test.Required)(&requiredValid))
	}
}

func BenchmarkGoPlaygroundRequired(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&requiredValid)
	}
}

// LT validation benchmarks
func BenchmarkGoValidLT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test.ValidateLT((*test.LT)(&ltValid))
	}
}

func BenchmarkGoPlaygroundLT(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&ltValid)
	}
}

// GT validation benchmarks
func BenchmarkGoValidGT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test.ValidateGT((*test.GT)(&gtValid))
	}
}

func BenchmarkGoPlaygroundGT(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&gtValid)
	}
}

// MaxLength validation benchmarks
func BenchmarkGoValidMaxLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test.ValidateMaxLength((*test.MaxLength)(&maxLengthValid))
	}
}

func BenchmarkGoPlaygroundMaxLength(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&maxLengthValid)
	}
}