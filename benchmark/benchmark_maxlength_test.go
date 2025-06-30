package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type MaxLengthTestInstance MaxLengthBenchData

var maxLengthInstance = MaxLengthTestInstance{
	Name: "This is a test string for maxlength validation",
}

func BenchmarkGoValidMaxLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ValidateMaxLength((*MaxLength)(&maxLengthInstance))
	}
}

func BenchmarkGoPlaygroundMaxLength(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&maxLengthInstance)
	}
}