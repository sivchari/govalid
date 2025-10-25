package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidIPV4(b *testing.B) {
	instance := test.IPV4{
		IP: "192.168.0.1",
	}
	for b.Loop() {
		err := test.ValidateIPV4(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundIPV4(b *testing.B) {
	validate := validator.New()
	instance := test.IPV4{
		IP: "192.168.0.1",
	}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}
