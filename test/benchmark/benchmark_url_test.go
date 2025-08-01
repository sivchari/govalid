package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"

	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidURL(b *testing.B) {
	instance := test.URL{
		URL: "https://example.com/path/to/resource?key=value",
	}
	for b.Loop() {
		err := test.ValidateURL(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundURL(b *testing.B) {
	validate := validator.New()
	instance := test.URL{
		URL: "https://example.com/path/to/resource?key=value",
	}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoValidatorURL(b *testing.B) {
	testURL := "https://example.com/path/to/resource?key=value"
	for b.Loop() {
		if !govalidator.IsURL(testURL) {
			b.Fatal("validation failed")
		}
	}
}

func BenchmarkGookitValidateURL(b *testing.B) {
	testURL := "https://example.com/path/to/resource?key=value"
	for b.Loop() {
		v := validate.New(map[string]any{"url": testURL})
		v.StringRule("url", "url")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
}
