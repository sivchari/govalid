
package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type RequiredTestInstance RequiredBenchData

var requiredInstance = RequiredTestInstance{
	Name:  "test",
	Age:   1,
	Items: []string{"test"},
}

func BenchmarkGoValidRequired(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ValidateRequired((*Required)(&requiredInstance))
	}
}

func BenchmarkGoPlaygroundRequired(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&requiredInstance)
	}
}
