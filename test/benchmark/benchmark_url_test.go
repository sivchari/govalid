package benchmark

import (
	"testing"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidURL(b *testing.B) {
	instance := test.URL{
		URL: "https://example.com/path/to/resource?key=value",
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateURL(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundURL(b *testing.B) {
	validate := validator.New()
	instance := test.URL{
		URL: "https://example.com/path/to/resource?key=value",
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

func BenchmarkGoValidatorURL(b *testing.B) {
	testURL := "https://example.com/path/to/resource?key=value"
	b.ResetTimer()
	for b.Loop() {
		if !govalidator.IsURL(testURL) {
			b.Fatal("validation failed")
		}
	}
	b.StopTimer()
}

func BenchmarkOzzoValidationURL(b *testing.B) {
	testURL := "https://example.com/path/to/resource?key=value"
	b.ResetTimer()
	for b.Loop() {
		err := validation.Validate(testURL, is.URL)
		if err != nil {
			b.Fatal("validation failed:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGookitValidateURL(b *testing.B) {
	testURL := "https://example.com/path/to/resource?key=value"
	b.ResetTimer()
	for b.Loop() {
		v := validate.New(map[string]any{"url": testURL})
		v.StringRule("url", "url")
		if !v.Validate() {
			b.Fatal("validation failed:", v.Errors)
		}
	}
	b.StopTimer()
}
