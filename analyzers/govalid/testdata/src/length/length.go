package length

//go:generate govalid ./length.go

type Length struct {
	// +govalid:length=7
	PostalCode string

	// +govalid:length=10
	PhoneNumber string
}
