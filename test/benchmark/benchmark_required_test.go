package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidRequired(b *testing.B) {
	instance := test.Required{
		Name:  "test",
		Age:   1,
		Items: []string{"test"},
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateRequired(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundRequired(b *testing.B) {
	validate := validator.New()
	instance := test.Required{
		Name:  "test",
		Age:   1,
		Items: []string{"test"},
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
