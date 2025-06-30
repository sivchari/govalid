
package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

type RequiredTestInstance test.TestRequired

var requiredInstance = RequiredTestInstance{
	Name:  "test",
	Age:   1,
	Items: []string{"test"},
}

func BenchmarkGoValidRequired(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test.ValidateRequired((*test.Required)(&requiredInstance))
	}
}

func BenchmarkGoPlaygroundRequired(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&requiredInstance)
	}
}
