package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

type GTTestInstance test.TestGT

var gtInstance = GTTestInstance{
	Age: 150,
}

func BenchmarkGoValidGT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test.ValidateGT((*test.GT)(&gtInstance))
	}
}

func BenchmarkGoPlaygroundGT(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(&gtInstance)
	}
}
