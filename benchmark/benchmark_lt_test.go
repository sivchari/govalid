
package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type LTTestInstance LTBenchData

var ltInstance = LTTestInstance{
	Age: 5,
}

func BenchmarkGoValidLT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ValidateLT((*LT)(&ltInstance))
	}
}

func BenchmarkGoPlaygroundLT(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&ltInstance)
	}
}
