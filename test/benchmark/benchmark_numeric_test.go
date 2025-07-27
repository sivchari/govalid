package benchmark

import (
    "testing"

    "github.com/go-playground/validator/v10"
    "github.com/sivchari/govalid/test"
)

func BenchmarkGoValidNumeric(b *testing.B) {
    instance := test.Numeric{Number: "123456"}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        err := test.ValidateNumeric(&instance)
        if err != nil {
            b.Fatal("unexpected error:", err)
        }
    }
    b.StopTimer()
}

func BenchmarkGoPlaygroundNumeric(b *testing.B) {
    validate := validator.New()
    validate.RegisterValidation("numeric", func(fl validator.FieldLevel) bool {
        str := fl.Field().String()
        for i := 0; i < len(str); i++ {
            if str[i] < '0' || str[i] > '9' {
                return false
            }
        }
        return true
    })

    instance := struct {
        Number string `validate:"numeric"`
    }{Number: "123456"}

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        err := validate.Struct(&instance)
        if err != nil {
            b.Fatal("unexpected error:", err)
        }
    }
    b.StopTimer()
}

