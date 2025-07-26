package alphanum

//go:generate govalid ./alphanum.go

type Alphanum struct {
	// +govalid:alphanum
	ProductCode string
	// +govalid:alphanum
	SerialNumber string
	// +govalid:alphanum
	Username string
}
