module github.com/sivchari/govalid/test

go 1.25.0

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/go-playground/validator/v10 v10.30.2
	github.com/gookit/validate v1.5.7
	github.com/sivchari/govalid v0.0.0-00010101000000-000000000000
)

replace github.com/sivchari/govalid => ../

require (
	github.com/gabriel-vasile/mimetype v1.4.13 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/gookit/filter v1.2.3 // indirect
	github.com/gookit/goutil v0.7.4 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/crypto v0.49.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
)
