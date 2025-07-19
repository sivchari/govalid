package benchmark

import (
	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
	"testing"
)

func BenchmarkGoValidAlpha(b *testing.B) {
	instance := test.Alpha{FirstName: "test_value"}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateAlpha(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundAlpha(b *testing.B) {
	validate := validator.New()
	instance := test.Alpha{FirstName: "test_value"}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}
