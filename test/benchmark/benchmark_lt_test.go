
package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

type LTTestInstance test.TestLT

var ltInstance = LTTestInstance{
	Age: 5,
}

func BenchmarkGoValidLT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test.ValidateLT((*test.LT)(&ltInstance))
	}
}

func BenchmarkGoPlaygroundLT(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&ltInstance)
	}
}
